/*
Copyright 2014 The Kubernetes Authors.

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

// Package options contains flags and options for initializing an apiserver
package options

import (
	cliflag "k8s.io/component-base/cli/flag"

	controlplaneapiserver "github.com/yangsoon/apiserver/pkg/controlplane/apiserver/options"
	_ "github.com/yangsoon/apiserver/pkg/features" // add the kubernetes feature gates
)

// ServerRunOptions runs a kubernetes api server.
type ServerRunOptions struct {
	*controlplaneapiserver.Options // embedded to avoid noise in existing consumers
}

// NewServerRunOptions creates a new ServerRunOptions object with default parameters
func NewServerRunOptions() *ServerRunOptions {
	s := ServerRunOptions{
		Options: controlplaneapiserver.NewOptions(),
	}
	return &s
}

// Flags returns flags for a specific APIServer by section name
func (s *ServerRunOptions) Flags() (fss cliflag.NamedFlagSets) {
	s.Options.AddFlags(&fss)
	return fss
}
