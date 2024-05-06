package mysql

import (
	"k8s.io/apiserver/pkg/registry/generic"
	"k8s.io/apiserver/pkg/storage"
	"k8s.io/client-go/tools/cache"
)

// StoreOptions is set of configuration options used to complete generic registries.
type StoreOptions struct {
	RESTOptions generic.RESTOptionsGetter
	TriggerFunc storage.IndexerFuncs
	AttrFunc    storage.AttrFunc
	Indexers    *cache.Indexers
}
