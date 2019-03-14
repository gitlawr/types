package v1alpha1

import (
	"context"

	"github.com/jetstack/cert-manager/pkg/apis/certmanager/v1alpha1"
	"github.com/rancher/norman/controller"
	"github.com/rancher/norman/objectclient"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/apimachinery/pkg/watch"
	"k8s.io/client-go/tools/cache"
)

var (
	IssuerGroupVersionKind = schema.GroupVersionKind{
		Version: Version,
		Group:   GroupName,
		Kind:    "Issuer",
	}
	IssuerResource = metav1.APIResource{
		Name:         "issuers",
		SingularName: "issuer",
		Namespaced:   true,

		Kind: IssuerGroupVersionKind.Kind,
	}
)

func NewIssuer(namespace, name string, obj v1alpha1.Issuer) *v1alpha1.Issuer {
	obj.APIVersion, obj.Kind = IssuerGroupVersionKind.ToAPIVersionAndKind()
	obj.Name = name
	obj.Namespace = namespace
	return &obj
}

type IssuerList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []v1alpha1.Issuer
}

type IssuerHandlerFunc func(key string, obj *v1alpha1.Issuer) (runtime.Object, error)

type IssuerChangeHandlerFunc func(obj *v1alpha1.Issuer) (runtime.Object, error)

type IssuerLister interface {
	List(namespace string, selector labels.Selector) (ret []*v1alpha1.Issuer, err error)
	Get(namespace, name string) (*v1alpha1.Issuer, error)
}

type IssuerController interface {
	Generic() controller.GenericController
	Informer() cache.SharedIndexInformer
	Lister() IssuerLister
	AddHandler(ctx context.Context, name string, handler IssuerHandlerFunc)
	AddClusterScopedHandler(ctx context.Context, name, clusterName string, handler IssuerHandlerFunc)
	Enqueue(namespace, name string)
	Sync(ctx context.Context) error
	Start(ctx context.Context, threadiness int) error
}

type IssuerInterface interface {
	ObjectClient() *objectclient.ObjectClient
	Create(*v1alpha1.Issuer) (*v1alpha1.Issuer, error)
	GetNamespaced(namespace, name string, opts metav1.GetOptions) (*v1alpha1.Issuer, error)
	Get(name string, opts metav1.GetOptions) (*v1alpha1.Issuer, error)
	Update(*v1alpha1.Issuer) (*v1alpha1.Issuer, error)
	Delete(name string, options *metav1.DeleteOptions) error
	DeleteNamespaced(namespace, name string, options *metav1.DeleteOptions) error
	List(opts metav1.ListOptions) (*IssuerList, error)
	Watch(opts metav1.ListOptions) (watch.Interface, error)
	DeleteCollection(deleteOpts *metav1.DeleteOptions, listOpts metav1.ListOptions) error
	Controller() IssuerController
	AddHandler(ctx context.Context, name string, sync IssuerHandlerFunc)
	AddLifecycle(ctx context.Context, name string, lifecycle IssuerLifecycle)
	AddClusterScopedHandler(ctx context.Context, name, clusterName string, sync IssuerHandlerFunc)
	AddClusterScopedLifecycle(ctx context.Context, name, clusterName string, lifecycle IssuerLifecycle)
}

type issuerLister struct {
	controller *issuerController
}

func (l *issuerLister) List(namespace string, selector labels.Selector) (ret []*v1alpha1.Issuer, err error) {
	err = cache.ListAllByNamespace(l.controller.Informer().GetIndexer(), namespace, selector, func(obj interface{}) {
		ret = append(ret, obj.(*v1alpha1.Issuer))
	})
	return
}

func (l *issuerLister) Get(namespace, name string) (*v1alpha1.Issuer, error) {
	var key string
	if namespace != "" {
		key = namespace + "/" + name
	} else {
		key = name
	}
	obj, exists, err := l.controller.Informer().GetIndexer().GetByKey(key)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(schema.GroupResource{
			Group:    IssuerGroupVersionKind.Group,
			Resource: "issuer",
		}, key)
	}
	return obj.(*v1alpha1.Issuer), nil
}

type issuerController struct {
	controller.GenericController
}

func (c *issuerController) Generic() controller.GenericController {
	return c.GenericController
}

func (c *issuerController) Lister() IssuerLister {
	return &issuerLister{
		controller: c,
	}
}

func (c *issuerController) AddHandler(ctx context.Context, name string, handler IssuerHandlerFunc) {
	c.GenericController.AddHandler(ctx, name, func(key string, obj interface{}) (interface{}, error) {
		if obj == nil {
			return handler(key, nil)
		} else if v, ok := obj.(*v1alpha1.Issuer); ok {
			return handler(key, v)
		} else {
			return nil, nil
		}
	})
}

func (c *issuerController) AddClusterScopedHandler(ctx context.Context, name, cluster string, handler IssuerHandlerFunc) {
	c.GenericController.AddHandler(ctx, name, func(key string, obj interface{}) (interface{}, error) {
		if obj == nil {
			return handler(key, nil)
		} else if v, ok := obj.(*v1alpha1.Issuer); ok && controller.ObjectInCluster(cluster, obj) {
			return handler(key, v)
		} else {
			return nil, nil
		}
	})
}

type issuerFactory struct {
}

func (c issuerFactory) Object() runtime.Object {
	return &v1alpha1.Issuer{}
}

func (c issuerFactory) List() runtime.Object {
	return &IssuerList{}
}

func (s *issuerClient) Controller() IssuerController {
	s.client.Lock()
	defer s.client.Unlock()

	c, ok := s.client.issuerControllers[s.ns]
	if ok {
		return c
	}

	genericController := controller.NewGenericController(IssuerGroupVersionKind.Kind+"Controller",
		s.objectClient)

	c = &issuerController{
		GenericController: genericController,
	}

	s.client.issuerControllers[s.ns] = c
	s.client.starters = append(s.client.starters, c)

	return c
}

type issuerClient struct {
	client       *Client
	ns           string
	objectClient *objectclient.ObjectClient
	controller   IssuerController
}

func (s *issuerClient) ObjectClient() *objectclient.ObjectClient {
	return s.objectClient
}

func (s *issuerClient) Create(o *v1alpha1.Issuer) (*v1alpha1.Issuer, error) {
	obj, err := s.objectClient.Create(o)
	return obj.(*v1alpha1.Issuer), err
}

func (s *issuerClient) Get(name string, opts metav1.GetOptions) (*v1alpha1.Issuer, error) {
	obj, err := s.objectClient.Get(name, opts)
	return obj.(*v1alpha1.Issuer), err
}

func (s *issuerClient) GetNamespaced(namespace, name string, opts metav1.GetOptions) (*v1alpha1.Issuer, error) {
	obj, err := s.objectClient.GetNamespaced(namespace, name, opts)
	return obj.(*v1alpha1.Issuer), err
}

func (s *issuerClient) Update(o *v1alpha1.Issuer) (*v1alpha1.Issuer, error) {
	obj, err := s.objectClient.Update(o.Name, o)
	return obj.(*v1alpha1.Issuer), err
}

func (s *issuerClient) Delete(name string, options *metav1.DeleteOptions) error {
	return s.objectClient.Delete(name, options)
}

func (s *issuerClient) DeleteNamespaced(namespace, name string, options *metav1.DeleteOptions) error {
	return s.objectClient.DeleteNamespaced(namespace, name, options)
}

func (s *issuerClient) List(opts metav1.ListOptions) (*IssuerList, error) {
	obj, err := s.objectClient.List(opts)
	return obj.(*IssuerList), err
}

func (s *issuerClient) Watch(opts metav1.ListOptions) (watch.Interface, error) {
	return s.objectClient.Watch(opts)
}

// Patch applies the patch and returns the patched deployment.
func (s *issuerClient) Patch(o *v1alpha1.Issuer, patchType types.PatchType, data []byte, subresources ...string) (*v1alpha1.Issuer, error) {
	obj, err := s.objectClient.Patch(o.Name, o, patchType, data, subresources...)
	return obj.(*v1alpha1.Issuer), err
}

func (s *issuerClient) DeleteCollection(deleteOpts *metav1.DeleteOptions, listOpts metav1.ListOptions) error {
	return s.objectClient.DeleteCollection(deleteOpts, listOpts)
}

func (s *issuerClient) AddHandler(ctx context.Context, name string, sync IssuerHandlerFunc) {
	s.Controller().AddHandler(ctx, name, sync)
}

func (s *issuerClient) AddLifecycle(ctx context.Context, name string, lifecycle IssuerLifecycle) {
	sync := NewIssuerLifecycleAdapter(name, false, s, lifecycle)
	s.Controller().AddHandler(ctx, name, sync)
}

func (s *issuerClient) AddClusterScopedHandler(ctx context.Context, name, clusterName string, sync IssuerHandlerFunc) {
	s.Controller().AddClusterScopedHandler(ctx, name, clusterName, sync)
}

func (s *issuerClient) AddClusterScopedLifecycle(ctx context.Context, name, clusterName string, lifecycle IssuerLifecycle) {
	sync := NewIssuerLifecycleAdapter(name+"_"+clusterName, true, s, lifecycle)
	s.Controller().AddClusterScopedHandler(ctx, name, clusterName, sync)
}

type IssuerIndexer func(obj *v1alpha1.Issuer) ([]string, error)

type IssuerClientCache interface {
	Get(namespace, name string) (*v1alpha1.Issuer, error)
	List(namespace string, selector labels.Selector) ([]*v1alpha1.Issuer, error)

	Index(name string, indexer IssuerIndexer)
	GetIndexed(name, key string) ([]*v1alpha1.Issuer, error)
}

type IssuerClient interface {
	Create(*v1alpha1.Issuer) (*v1alpha1.Issuer, error)
	Get(namespace, name string, opts metav1.GetOptions) (*v1alpha1.Issuer, error)
	Update(*v1alpha1.Issuer) (*v1alpha1.Issuer, error)
	Delete(namespace, name string, options *metav1.DeleteOptions) error
	List(namespace string, opts metav1.ListOptions) (*IssuerList, error)
	Watch(opts metav1.ListOptions) (watch.Interface, error)

	Cache() IssuerClientCache

	OnCreate(ctx context.Context, name string, sync IssuerChangeHandlerFunc)
	OnChange(ctx context.Context, name string, sync IssuerChangeHandlerFunc)
	OnRemove(ctx context.Context, name string, sync IssuerChangeHandlerFunc)
	Enqueue(namespace, name string)

	Generic() controller.GenericController
	ObjectClient() *objectclient.ObjectClient
	Interface() IssuerInterface
}

type issuerClientCache struct {
	client *issuerClient2
}

type issuerClient2 struct {
	iface      IssuerInterface
	controller IssuerController
}

func (n *issuerClient2) Interface() IssuerInterface {
	return n.iface
}

func (n *issuerClient2) Generic() controller.GenericController {
	return n.iface.Controller().Generic()
}

func (n *issuerClient2) ObjectClient() *objectclient.ObjectClient {
	return n.Interface().ObjectClient()
}

func (n *issuerClient2) Enqueue(namespace, name string) {
	n.iface.Controller().Enqueue(namespace, name)
}

func (n *issuerClient2) Create(obj *v1alpha1.Issuer) (*v1alpha1.Issuer, error) {
	return n.iface.Create(obj)
}

func (n *issuerClient2) Get(namespace, name string, opts metav1.GetOptions) (*v1alpha1.Issuer, error) {
	return n.iface.GetNamespaced(namespace, name, opts)
}

func (n *issuerClient2) Update(obj *v1alpha1.Issuer) (*v1alpha1.Issuer, error) {
	return n.iface.Update(obj)
}

func (n *issuerClient2) Delete(namespace, name string, options *metav1.DeleteOptions) error {
	return n.iface.DeleteNamespaced(namespace, name, options)
}

func (n *issuerClient2) List(namespace string, opts metav1.ListOptions) (*IssuerList, error) {
	return n.iface.List(opts)
}

func (n *issuerClient2) Watch(opts metav1.ListOptions) (watch.Interface, error) {
	return n.iface.Watch(opts)
}

func (n *issuerClientCache) Get(namespace, name string) (*v1alpha1.Issuer, error) {
	return n.client.controller.Lister().Get(namespace, name)
}

func (n *issuerClientCache) List(namespace string, selector labels.Selector) ([]*v1alpha1.Issuer, error) {
	return n.client.controller.Lister().List(namespace, selector)
}

func (n *issuerClient2) Cache() IssuerClientCache {
	n.loadController()
	return &issuerClientCache{
		client: n,
	}
}

func (n *issuerClient2) OnCreate(ctx context.Context, name string, sync IssuerChangeHandlerFunc) {
	n.loadController()
	n.iface.AddLifecycle(ctx, name+"-create", &issuerLifecycleDelegate{create: sync})
}

func (n *issuerClient2) OnChange(ctx context.Context, name string, sync IssuerChangeHandlerFunc) {
	n.loadController()
	n.iface.AddLifecycle(ctx, name+"-change", &issuerLifecycleDelegate{update: sync})
}

func (n *issuerClient2) OnRemove(ctx context.Context, name string, sync IssuerChangeHandlerFunc) {
	n.loadController()
	n.iface.AddLifecycle(ctx, name, &issuerLifecycleDelegate{remove: sync})
}

func (n *issuerClientCache) Index(name string, indexer IssuerIndexer) {
	err := n.client.controller.Informer().GetIndexer().AddIndexers(map[string]cache.IndexFunc{
		name: func(obj interface{}) ([]string, error) {
			if v, ok := obj.(*v1alpha1.Issuer); ok {
				return indexer(v)
			}
			return nil, nil
		},
	})

	if err != nil {
		panic(err)
	}
}

func (n *issuerClientCache) GetIndexed(name, key string) ([]*v1alpha1.Issuer, error) {
	var result []*v1alpha1.Issuer
	objs, err := n.client.controller.Informer().GetIndexer().ByIndex(name, key)
	if err != nil {
		return nil, err
	}
	for _, obj := range objs {
		if v, ok := obj.(*v1alpha1.Issuer); ok {
			result = append(result, v)
		}
	}

	return result, nil
}

func (n *issuerClient2) loadController() {
	if n.controller == nil {
		n.controller = n.iface.Controller()
	}
}

type issuerLifecycleDelegate struct {
	create IssuerChangeHandlerFunc
	update IssuerChangeHandlerFunc
	remove IssuerChangeHandlerFunc
}

func (n *issuerLifecycleDelegate) HasCreate() bool {
	return n.create != nil
}

func (n *issuerLifecycleDelegate) Create(obj *v1alpha1.Issuer) (runtime.Object, error) {
	if n.create == nil {
		return obj, nil
	}
	return n.create(obj)
}

func (n *issuerLifecycleDelegate) HasFinalize() bool {
	return n.remove != nil
}

func (n *issuerLifecycleDelegate) Remove(obj *v1alpha1.Issuer) (runtime.Object, error) {
	if n.remove == nil {
		return obj, nil
	}
	return n.remove(obj)
}

func (n *issuerLifecycleDelegate) Updated(obj *v1alpha1.Issuer) (runtime.Object, error) {
	if n.update == nil {
		return obj, nil
	}
	return n.update(obj)
}
