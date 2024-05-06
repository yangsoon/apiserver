package storage

import (
	api "github.com/yangsoon/apiserver/pkg/apis/apps_kruise_io"
	"github.com/yangsoon/apiserver/pkg/printers"
	printersinternal "github.com/yangsoon/apiserver/pkg/printers/internalversion"
	printerstorage "github.com/yangsoon/apiserver/pkg/printers/storage"
	"github.com/yangsoon/apiserver/pkg/registry/apps_kruise_io/cloneset"
	"github.com/yangsoon/apiserver/pkg/store/mysql"
	mysqlregistry "github.com/yangsoon/apiserver/pkg/store/mysql/registry"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apiserver/pkg/registry/generic"
	"k8s.io/apiserver/pkg/registry/rest"
)

// REST implements a RESTStorage for ConfigMap
type REST struct {
	*mysqlregistry.Store
}

// NewStorage returns a RESTStorage object that will work with ConfigMap objects.
func NewStorage(optsGetter generic.RESTOptionsGetter) (*REST, error) {
	store := &mysqlregistry.Store{
		NewFunc:                   func() runtime.Object { return &api.CloneSet{} },
		NewListFunc:               func() runtime.Object { return &api.CloneSetList{} },
		PredicateFunc:             cloneset.Matcher,
		DefaultQualifiedResource:  api.Resource("clonesets"),
		SingularQualifiedResource: api.Resource("cloneset"),

		CreateStrategy: cloneset.Strategy,
		UpdateStrategy: cloneset.Strategy,
		DeleteStrategy: cloneset.Strategy,

		TableConvertor: printerstorage.TableConvertor{TableGenerator: printers.NewTableGenerator().With(printersinternal.AddHandlers)},
	}
	options := &mysql.StoreOptions{
		RESTOptions: optsGetter,
		AttrFunc:    cloneset.GetAttrs,
	}
	if err := store.CompleteWithOptions(options); err != nil {
		return nil, err
	}
	return &REST{store}, nil
}

// Implement ShortNamesProvider
var _ rest.ShortNamesProvider = &REST{}

// ShortNames implements the ShortNamesProvider interface. Returns a list of short names for a resource.
func (r *REST) ShortNames() []string {
	return []string{"clone"}
}
