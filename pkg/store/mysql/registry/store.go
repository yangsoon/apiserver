package registry

import (
	"context"
	"fmt"
	"github.com/yangsoon/apiserver/pkg/store/mysql"
	kerrors "k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/api/meta"
	apimeta "k8s.io/apimachinery/pkg/api/meta"
	metainternalversion "k8s.io/apimachinery/pkg/apis/meta/internalversion"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/fields"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/util/uuid"
	"k8s.io/apimachinery/pkg/watch"
	genericapirequest "k8s.io/apiserver/pkg/endpoints/request"
	"k8s.io/apiserver/pkg/registry/generic/registry"
	"k8s.io/apiserver/pkg/registry/rest"
	"k8s.io/apiserver/pkg/storage"
	"strings"
)

var _ rest.StandardStorage = &Store{}
var _ rest.Scoper = &Store{}
var _ rest.Storage = &Store{}
var _ rest.TableConvertor = &Store{}
var _ registry.GenericStore = &Store{}

var _ rest.SingularNameProvider = &Store{}

var _ registry.GenericStore = &Store{}

var cache = make(map[string]runtime.Object)

type Store struct {
	// NewFunc returns a new instance of the type this registry returns for a
	// GET of a single object, e.g.:
	//
	// curl GET /apis/group/version/namespaces/my-ns/myresource/name-of-object
	NewFunc func() runtime.Object

	// NewListFunc returns a new list of the type this registry; it is the
	// type returned when the resource is listed, e.g.:
	//
	// curl GET /apis/group/version/namespaces/my-ns/myresource
	NewListFunc func() runtime.Object

	// DefaultQualifiedResource is the pluralized name of the resource.
	// This field is used if there is no request info present in the context.
	// See qualifiedResourceFromContext for details.
	DefaultQualifiedResource schema.GroupResource

	// SingularQualifiedResource is the singular name of the resource.
	SingularQualifiedResource schema.GroupResource

	// KeyRootFunc returns the root etcd key for this resource; should not
	// include trailing "/".  This is used for operations that work on the
	// entire collection (listing and watching).
	//
	// KeyRootFunc and KeyFunc must be supplied together or not at all.
	KeyRootFunc func(ctx context.Context) string

	// KeyFunc returns the key for a specific object in the collection.
	// KeyFunc is called for Create/Update/Get/Delete. Note that 'namespace'
	// can be gotten from ctx.
	//
	// KeyFunc and KeyRootFunc must be supplied together or not at all.
	KeyFunc func(ctx context.Context, name string) (string, error)

	// ObjectNameFunc returns the name of an object or an error.
	ObjectNameFunc func(obj runtime.Object) (string, error)

	// PredicateFunc returns a matcher corresponding to the provided labels
	// and fields. The SelectionPredicate returned should return true if the
	// object matches the given field and label selectors.
	PredicateFunc func(label labels.Selector, field fields.Selector) storage.SelectionPredicate

	// Decorator is an optional exit hook on an object returned from the
	// underlying storage. The returned object could be an individual object
	// (e.g. Pod) or a list type (e.g. PodList). Decorator is intended for
	// integrations that are above storage and should only be used for
	// specific cases where storage of the value is not appropriate, since
	// they cannot be watched.
	Decorator func(runtime.Object)

	// CreateStrategy implements resource-specific behavior during creation.
	CreateStrategy rest.RESTCreateStrategy
	// BeginCreate is an optional hook that returns a "transaction-like"
	// commit/revert function which will be called at the end of the operation,
	// but before AfterCreate and Decorator, indicating via the argument
	// whether the operation succeeded.  If this returns an error, the function
	// is not called.  Almost nobody should use this hook.
	BeginCreate registry.BeginCreateFunc
	// AfterCreate implements a further operation to run after a resource is
	// created and before it is decorated, optional.
	AfterCreate registry.AfterCreateFunc

	// UpdateStrategy implements resource-specific behavior during updates.
	UpdateStrategy rest.RESTUpdateStrategy
	// BeginUpdate is an optional hook that returns a "transaction-like"
	// commit/revert function which will be called at the end of the operation,
	// but before AfterUpdate and Decorator, indicating via the argument
	// whether the operation succeeded.  If this returns an error, the function
	// is not called.  Almost nobody should use this hook.
	BeginUpdate registry.BeginUpdateFunc
	// AfterUpdate implements a further operation to run after a resource is
	// updated and before it is decorated, optional.
	AfterUpdate registry.AfterUpdateFunc

	// DeleteStrategy implements resource-specific behavior during deletion.
	DeleteStrategy rest.RESTDeleteStrategy
	// AfterDelete implements a further operation to run after a resource is
	// deleted and before it is decorated, optional.
	AfterDelete registry.AfterDeleteFunc

	// TableConvertor is an optional interface for transforming items or lists
	// of items into tabular output. If unset, the default will be used.
	TableConvertor rest.TableConvertor

	// DestroyFunc cleans up clients used by the underlying Storage; optional.
	// If set, DestroyFunc has to be implemented in thread-safe way and
	// be prepared for being called more than once.
	DestroyFunc func()

	KeyFromObjFunc func(object runtime.Object) (string, error)

	ResourcePrefix string
}

func (s *Store) Get(ctx context.Context, name string, options *metav1.GetOptions) (runtime.Object, error) {
	key, err := s.KeyFunc(ctx, name)
	if err != nil {
		return nil, err
	}
	obj, ok := cache[key]
	if !ok {
		return nil, kerrors.NewNotFound(s.DefaultQualifiedResource, name)
	}
	return obj, nil
}

func (s *Store) NewList() runtime.Object {
	return s.NewListFunc()
}

func (s *Store) List(ctx context.Context, options *metainternalversion.ListOptions) (runtime.Object, error) {
	objectList := s.NewListFunc()
	var prefix string
	ns, ok := genericapirequest.NamespaceFrom(ctx)
	if ok {
		prefix = s.ResourcePrefix + "/" + ns
	}
	var runtimeObjs []runtime.Object
	for key, item := range cache {
		if strings.HasPrefix(key, prefix) {
			runtimeObjs = append(runtimeObjs, item.DeepCopyObject())
		}
	}
	if err := apimeta.SetList(objectList, runtimeObjs); err != nil {
		return nil, err
	}
	return objectList, nil
}

func (s *Store) New() runtime.Object {
	return s.NewFunc()
}

func (s *Store) Create(ctx context.Context, obj runtime.Object, createValidation rest.ValidateObjectFunc, options *metav1.CreateOptions) (runtime.Object, error) {
	key, err := s.KeyFromObjFunc(obj)
	if err != nil {
		return nil, err
	}
	accessor, err := meta.Accessor(obj)
	if err != nil {
		// If no UID can be read, no preconditions are possible
		return nil, err
	}
	accessor.SetUID(uuid.NewUUID())
	cache[key] = obj
	return obj, nil
}

func (s *Store) Update(ctx context.Context, name string, objInfo rest.UpdatedObjectInfo, createValidation rest.ValidateObjectFunc, updateValidation rest.ValidateObjectUpdateFunc, forceAllowCreate bool, options *metav1.UpdateOptions) (runtime.Object, bool, error) {
	key, err := s.KeyFunc(ctx, name)
	if err != nil {
		return nil, false, err
	}
	oldObj := cache[key]
	newOld, err := objInfo.UpdatedObject(ctx, oldObj)
	if err != nil {
		return nil, false, err
	}
	accessor, err := meta.Accessor(newOld)
	if err != nil {
		// If no UID can be read, no preconditions are possible
		return nil, false, err
	}
	accessor.SetUID(uuid.NewUUID())
	cache[key] = newOld
	return newOld, true, nil
}

func (s *Store) Delete(ctx context.Context, name string, deleteValidation rest.ValidateObjectFunc, options *metav1.DeleteOptions) (runtime.Object, bool, error) {
	key, err := s.KeyFunc(ctx, name)
	if err != nil {
		return nil, false, err
	}
	obj := cache[key]
	delete(cache, key)
	return obj, true, nil
}

func (s *Store) DeleteCollection(ctx context.Context, deleteValidation rest.ValidateObjectFunc, options *metav1.DeleteOptions, listOptions *metainternalversion.ListOptions) (runtime.Object, error) {
	return s.New(), nil
}

func (s *Store) Watch(ctx context.Context, options *metainternalversion.ListOptions) (watch.Interface, error) {
	//TODO implement me
	panic("implement me")
}

func (s *Store) Destroy() {
}

func (s *Store) NamespaceScoped() bool {
	if s.CreateStrategy != nil {
		return s.CreateStrategy.NamespaceScoped()
	}
	if s.UpdateStrategy != nil {
		return s.UpdateStrategy.NamespaceScoped()
	}

	panic("programmer error: no CRUD for resource, override NamespaceScoped too")
}

// GetCreateStrategy implements GenericStore.
func (s *Store) GetCreateStrategy() rest.RESTCreateStrategy {
	return s.CreateStrategy
}

// GetUpdateStrategy implements GenericStore.
func (s *Store) GetUpdateStrategy() rest.RESTUpdateStrategy {
	return s.UpdateStrategy
}

// GetDeleteStrategy implements GenericStore.
func (s *Store) GetDeleteStrategy() rest.RESTDeleteStrategy {
	return s.DeleteStrategy
}

func (s *Store) ConvertToTable(ctx context.Context, object runtime.Object, tableOptions runtime.Object) (*metav1.Table, error) {
	if s.TableConvertor != nil {
		return s.TableConvertor.ConvertToTable(ctx, object, tableOptions)
	}
	return rest.NewDefaultTableConvertor(s.DefaultQualifiedResource).ConvertToTable(ctx, object, tableOptions)
}

func (s *Store) GetSingularName() string {
	return s.SingularQualifiedResource.Resource
}

func (s *Store) CompleteWithOptions(options *mysql.StoreOptions) error {
	if s.DefaultQualifiedResource.Empty() {
		return fmt.Errorf("store %#v must have a non-empty qualified resource", s)
	}
	if s.SingularQualifiedResource.Empty() {
		return fmt.Errorf("store %#v must have a non-empty singular qualified resource", s)
	}
	if s.DefaultQualifiedResource.Group != s.SingularQualifiedResource.Group {
		return fmt.Errorf("store for %#v, singular and plural qualified resource's group name's must match", s)
	}
	if s.NewFunc == nil {
		return fmt.Errorf("store for %s must have NewFunc set", s.DefaultQualifiedResource.String())
	}
	if s.NewListFunc == nil {
		return fmt.Errorf("store for %s must have NewListFunc set", s.DefaultQualifiedResource.String())
	}
	if (s.KeyRootFunc == nil) != (s.KeyFunc == nil) {
		return fmt.Errorf("store for %s must set both KeyRootFunc and KeyFunc or neither", s.DefaultQualifiedResource.String())
	}

	if s.TableConvertor == nil {
		return fmt.Errorf("store for %s must set TableConvertor; rest.NewDefaultTableConvertor(e.DefaultQualifiedResource) can be used to output just name/creation time", s.DefaultQualifiedResource.String())
	}

	var isNamespaced bool
	switch {
	case s.CreateStrategy != nil:
		isNamespaced = s.CreateStrategy.NamespaceScoped()
	case s.UpdateStrategy != nil:
		isNamespaced = s.UpdateStrategy.NamespaceScoped()
	default:
		return fmt.Errorf("store for %s must have CreateStrategy or UpdateStrategy set", s.DefaultQualifiedResource.String())
	}

	if s.DeleteStrategy == nil {
		return fmt.Errorf("store for %s must have DeleteStrategy set", s.DefaultQualifiedResource.String())
	}

	if options.RESTOptions == nil {
		return fmt.Errorf("options for %s must have RESTOptions set", s.DefaultQualifiedResource.String())
	}

	opts, err := options.RESTOptions.GetRESTOptions(s.DefaultQualifiedResource)
	if err != nil {
		return err
	}

	// ResourcePrefix must come from the underlying factory
	prefix := opts.ResourcePrefix
	if !strings.HasPrefix(prefix, "/") {
		prefix = "/" + prefix
	}
	if prefix == "/" {
		return fmt.Errorf("store for %s has an invalid prefix %q", s.DefaultQualifiedResource.String(), opts.ResourcePrefix)
	}
	s.ResourcePrefix = prefix

	// Set the default behavior for storage key generation
	if s.KeyRootFunc == nil && s.KeyFunc == nil {
		if isNamespaced {
			s.KeyRootFunc = func(ctx context.Context) string {
				return registry.NamespaceKeyRootFunc(ctx, prefix)
			}
			s.KeyFunc = func(ctx context.Context, name string) (string, error) {
				return registry.NamespaceKeyFunc(ctx, prefix, name)
			}
		} else {
			s.KeyRootFunc = func(ctx context.Context) string {
				return prefix
			}
			s.KeyFunc = func(ctx context.Context, name string) (string, error) {
				return registry.NoNamespaceKeyFunc(ctx, prefix, name)
			}
		}
	}

	// We adapt the store's keyFunc so that we can use it with the StorageDecorator
	// without making any assumptions about where objects are stored in etcd
	keyFunc := func(obj runtime.Object) (string, error) {
		accessor, err := meta.Accessor(obj)
		if err != nil {
			return "", err
		}

		if isNamespaced {
			return s.KeyFunc(genericapirequest.WithNamespace(genericapirequest.NewContext(), accessor.GetNamespace()), accessor.GetName())
		}

		return s.KeyFunc(genericapirequest.NewContext(), accessor.GetName())
	}
	s.KeyFromObjFunc = keyFunc

	if s.ObjectNameFunc == nil {
		s.ObjectNameFunc = func(obj runtime.Object) (string, error) {
			accessor, err := meta.Accessor(obj)
			if err != nil {
				return "", err
			}
			return accessor.GetName(), nil
		}
	}
	return nil
}
