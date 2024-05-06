/*
Copyright 2016 The Kubernetes Authors.

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

package rest

import (
	api "github.com/yangsoon/apiserver/pkg/apis/core"
	podstore "github.com/yangsoon/apiserver/pkg/registry/core/pod/storage"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apiserver/pkg/registry/generic"
	"k8s.io/apiserver/pkg/registry/rest"
	genericapiserver "k8s.io/apiserver/pkg/server"
	serverstorage "k8s.io/apiserver/pkg/server/storage"
)

// Config provides information needed to build RESTStorage for core.
type Config struct {
	GenericConfig
}

type legacyProvider struct {
	Config
}

func New(c Config) (*legacyProvider, error) {
	p := &legacyProvider{
		Config: c,
	}
	return p, nil
}

func (c *legacyProvider) NewRESTStorage(apiResourceConfigSource serverstorage.APIResourceConfigSource, restOptionsGetter generic.RESTOptionsGetter) (genericapiserver.APIGroupInfo, error) {
	apiGroupInfo, err := c.GenericConfig.NewRESTStorage(apiResourceConfigSource, restOptionsGetter)
	if err != nil {
		return genericapiserver.APIGroupInfo{}, err
	}

	//persistentVolumeStorage, persistentVolumeStatusStorage, err := pvstore.NewREST(restOptionsGetter)
	//if err != nil {
	//	return genericapiserver.APIGroupInfo{}, err
	//}
	//persistentVolumeClaimStorage, persistentVolumeClaimStatusStorage, err := pvcstore.NewREST(restOptionsGetter)
	//if err != nil {
	//	return genericapiserver.APIGroupInfo{}, err
	//}

	podStorage, err := podstore.NewStorage(
		restOptionsGetter,
	)
	if err != nil {
		return genericapiserver.APIGroupInfo{}, err
	}

	//serviceRESTStorage, serviceStatusStorage, serviceRESTProxy, err := servicestore.NewREST(
	//	restOptionsGetter,
	//	c.primaryServiceClusterIPAllocator.IPFamily(),
	//	c.serviceClusterIPAllocators,
	//	c.serviceNodePortAllocator,
	//	endpointsStorage,
	//	podStorage.Pod,
	//	c.Proxy.Transport)
	//if err != nil {
	//	return genericapiserver.APIGroupInfo{}, err
	//}

	storage := apiGroupInfo.VersionedResourcesStorageMap["v1"]
	if storage == nil {
		storage = map[string]rest.Storage{}
	}

	if resource := "pods"; apiResourceConfigSource.ResourceEnabled(corev1.SchemeGroupVersion.WithResource(resource)) {
		storage[resource] = podStorage.Pod
		storage[resource+"/status"] = podStorage.Status
	}

	if len(storage) > 0 {
		apiGroupInfo.VersionedResourcesStorageMap["v1"] = storage
	}

	return apiGroupInfo, nil
}

func (p *legacyProvider) GroupName() string {
	return api.GroupName
}
