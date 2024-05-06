/*
Copyright 2015 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package storage

import (
	api "github.com/yangsoon/apiserver/pkg/apis/core"
	"github.com/yangsoon/apiserver/pkg/printers"
	printersinternal "github.com/yangsoon/apiserver/pkg/printers/internalversion"
	printerstorage "github.com/yangsoon/apiserver/pkg/printers/storage"
	"github.com/yangsoon/apiserver/pkg/registry/core/configmap"
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

// NewREST returns a RESTStorage object that will work with ConfigMap objects.
func NewREST(optsGetter generic.RESTOptionsGetter) (*REST, error) {
	store := &mysqlregistry.Store{
		NewFunc:                   func() runtime.Object { return &api.ConfigMap{} },
		NewListFunc:               func() runtime.Object { return &api.ConfigMapList{} },
		PredicateFunc:             configmap.Matcher,
		DefaultQualifiedResource:  api.Resource("configmaps"),
		SingularQualifiedResource: api.Resource("configmap"),

		CreateStrategy: configmap.Strategy,
		UpdateStrategy: configmap.Strategy,
		DeleteStrategy: configmap.Strategy,

		TableConvertor: printerstorage.TableConvertor{TableGenerator: printers.NewTableGenerator().With(printersinternal.AddHandlers)},
	}
	options := &mysql.StoreOptions{
		RESTOptions: optsGetter,
		AttrFunc:    configmap.GetAttrs,
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
	return []string{"cm"}
}
