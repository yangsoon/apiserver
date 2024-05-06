package options

import (
	"k8s.io/apimachinery/pkg/runtime/schema"
	genericregistry "k8s.io/apiserver/pkg/registry/generic"
	"k8s.io/apiserver/pkg/server"
	"k8s.io/apiserver/pkg/storage/storagebackend"
)

type MySQLOptions struct {
	StorageConfig storagebackend.Config
}

var _ genericregistry.RESTOptionsGetter = emptyRESTOptionsGetter{}

func (o *MySQLOptions) ApplyTo(c *server.Config) error {
	c.RESTOptionsGetter = o.CreateRESTOptionsGetter()
	return nil
}

func (o *MySQLOptions) CreateRESTOptionsGetter() genericregistry.RESTOptionsGetter {
	return emptyRESTOptionsGetter{}
}

type emptyRESTOptionsGetter struct {
}

func (e emptyRESTOptionsGetter) GetRESTOptions(resource schema.GroupResource) (genericregistry.RESTOptions, error) {
	return genericregistry.RESTOptions{
		ResourcePrefix: e.ResourcePrefix(resource),
	}, nil
}

func (e emptyRESTOptionsGetter) ResourcePrefix(resource schema.GroupResource) string {
	return resource.Group + "/" + resource.Resource
}
