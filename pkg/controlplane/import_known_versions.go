package controlplane

import (
	// These imports are the API groups the API server will support.
	_ "github.com/yangsoon/apiserver/pkg/apis/apps/install"
	_ "github.com/yangsoon/apiserver/pkg/apis/apps_kruise_io/install"
	_ "github.com/yangsoon/apiserver/pkg/apis/core/install"
)
