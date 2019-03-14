package v1alpha1

import (
	"github.com/jetstack/cert-manager/pkg/apis/certmanager/v1alpha1"
	"github.com/rancher/norman/lifecycle"
	"k8s.io/apimachinery/pkg/runtime"
)

type IssuerLifecycle interface {
	Create(obj *v1alpha1.Issuer) (runtime.Object, error)
	Remove(obj *v1alpha1.Issuer) (runtime.Object, error)
	Updated(obj *v1alpha1.Issuer) (runtime.Object, error)
}

type issuerLifecycleAdapter struct {
	lifecycle IssuerLifecycle
}

func (w *issuerLifecycleAdapter) HasCreate() bool {
	o, ok := w.lifecycle.(lifecycle.ObjectLifecycleCondition)
	return !ok || o.HasCreate()
}

func (w *issuerLifecycleAdapter) HasFinalize() bool {
	o, ok := w.lifecycle.(lifecycle.ObjectLifecycleCondition)
	return !ok || o.HasFinalize()
}

func (w *issuerLifecycleAdapter) Create(obj runtime.Object) (runtime.Object, error) {
	o, err := w.lifecycle.Create(obj.(*v1alpha1.Issuer))
	if o == nil {
		return nil, err
	}
	return o, err
}

func (w *issuerLifecycleAdapter) Finalize(obj runtime.Object) (runtime.Object, error) {
	o, err := w.lifecycle.Remove(obj.(*v1alpha1.Issuer))
	if o == nil {
		return nil, err
	}
	return o, err
}

func (w *issuerLifecycleAdapter) Updated(obj runtime.Object) (runtime.Object, error) {
	o, err := w.lifecycle.Updated(obj.(*v1alpha1.Issuer))
	if o == nil {
		return nil, err
	}
	return o, err
}

func NewIssuerLifecycleAdapter(name string, clusterScoped bool, client IssuerInterface, l IssuerLifecycle) IssuerHandlerFunc {
	adapter := &issuerLifecycleAdapter{lifecycle: l}
	syncFn := lifecycle.NewObjectLifecycleAdapter(name, clusterScoped, adapter, client.ObjectClient())
	return func(key string, obj *v1alpha1.Issuer) (runtime.Object, error) {
		newObj, err := syncFn(key, obj)
		if o, ok := newObj.(runtime.Object); ok {
			return o, err
		}
		return nil, err
	}
}
