package install

import (
	"k8s.io/apimachinery/pkg/runtime"
	utilruntime "k8s.io/apimachinery/pkg/util/runtime"

	"github.com/yangsoon/apiserver/pkg/api/legacyscheme"
	"github.com/yangsoon/apiserver/pkg/apis/apps_kruise_io"
	"github.com/yangsoon/apiserver/pkg/apis/apps_kruise_io/v1alpha1"
)

func init() {
	Install(legacyscheme.Scheme)
}

// Install registers the API group and adds types to a scheme
func Install(scheme *runtime.Scheme) {
	utilruntime.Must(apps_kruise_io.AddToScheme(scheme))
	utilruntime.Must(v1alpha1.AddToScheme(scheme))
	utilruntime.Must(scheme.SetVersionPriority(v1alpha1.SchemeGroupVersion))
}
