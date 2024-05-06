/*
Copyright 2024 The RSI Authors.

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

// copy from https://github.com/kubernetes/kubernetes/blob/36981002246682ed7dc4de54ccc2a96c1a0cbbdb/pkg/apis/core/doc.go

// +k8s:deepcopy-gen=package

// Package core contains the latest (or "internal") version of the
// Kubernetes API objects. This is the API objects as represented in memory.
// The contract presented to clients is located in the versioned packages,
// which are sub-directories. The first one is "v1". Those packages
// describe how a particular version is serialized to storage/network.
package core // import "github.com/yangsoon/apiserver/pkg/apis/core"
