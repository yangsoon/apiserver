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
	appskruiseioapiv1alpha1 "github.com/openkruise/kruise-api/apps/v1alpha1"
	"github.com/yangsoon/apiserver/pkg/api/legacyscheme"
	"github.com/yangsoon/apiserver/pkg/apis/apps_kruise_io"
	clonesetsstore "github.com/yangsoon/apiserver/pkg/registry/apps_kruise_io/cloneset/storage"
	"k8s.io/apiserver/pkg/registry/generic"
	"k8s.io/apiserver/pkg/registry/rest"
	genericapiserver "k8s.io/apiserver/pkg/server"
	serverstorage "k8s.io/apiserver/pkg/server/storage"
)

// StorageProvider is a struct for apps REST storage.
type StorageProvider struct{}

// NewRESTStorage returns APIGroupInfo object.
func (p StorageProvider) NewRESTStorage(apiResourceConfigSource serverstorage.APIResourceConfigSource, restOptionsGetter generic.RESTOptionsGetter) (genericapiserver.APIGroupInfo, error) {
	apiGroupInfo := genericapiserver.NewDefaultAPIGroupInfo(apps_kruise_io.GroupName, legacyscheme.Scheme, legacyscheme.ParameterCodec, legacyscheme.Codecs)
	// If you add a version here, be sure to add an entry in `k8s.io/kubernetes/cmd/kube-apiserver/app/aggregator.go with specific priorities.
	// TODO refactor the plumbing to provide the information in the APIGroupInfo

	if storageMap, err := p.v1alpha1Storage(apiResourceConfigSource, restOptionsGetter); err != nil {
		return genericapiserver.APIGroupInfo{}, err
	} else if len(storageMap) > 0 {
		apiGroupInfo.VersionedResourcesStorageMap[appskruiseioapiv1alpha1.SchemeGroupVersion.Version] = storageMap
	}

	return apiGroupInfo, nil
}

func (p StorageProvider) v1alpha1Storage(apiResourceConfigSource serverstorage.APIResourceConfigSource, restOptionsGetter generic.RESTOptionsGetter) (map[string]rest.Storage, error) {
	storage := map[string]rest.Storage{}

	// deployments
	if resource := "clonesets"; apiResourceConfigSource.ResourceEnabled(appskruiseioapiv1alpha1.SchemeGroupVersion.WithResource(resource)) {
		cloneSetStorage, err := clonesetsstore.NewStorage(restOptionsGetter)
		if err != nil {
			return storage, err
		}
		storage[resource] = cloneSetStorage
		//storage[resource+"/status"] = deploymentStorage.Status
		//storage[resource+"/scale"] = deploymentStorage.Scale
	}
	return storage, nil
}

// GroupName returns name of the group
func (p StorageProvider) GroupName() string {
	return apps_kruise_io.GroupName
}
