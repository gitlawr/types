package v3

import (
	"context"

	"github.com/rancher/norman/clientbase"
	"github.com/rancher/norman/controller"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/watch"
	"k8s.io/client-go/tools/cache"
)

var (
	PipelineLogGroupVersionKind = schema.GroupVersionKind{
		Version: Version,
		Group:   GroupName,
		Kind:    "PipelineLog",
	}
	PipelineLogResource = metav1.APIResource{
		Name:         "pipelinelogs",
		SingularName: "pipelinelog",
		Namespaced:   true,

		Kind: PipelineLogGroupVersionKind.Kind,
	}
)

type PipelineLogList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []PipelineLog
}

type PipelineLogHandlerFunc func(key string, obj *PipelineLog) error

type PipelineLogLister interface {
	List(namespace string, selector labels.Selector) (ret []*PipelineLog, err error)
	Get(namespace, name string) (*PipelineLog, error)
}

type PipelineLogController interface {
	Informer() cache.SharedIndexInformer
	Lister() PipelineLogLister
	AddHandler(name string, handler PipelineLogHandlerFunc)
	AddClusterScopedHandler(name, clusterName string, handler PipelineLogHandlerFunc)
	Enqueue(namespace, name string)
	Sync(ctx context.Context) error
	Start(ctx context.Context, threadiness int) error
}

type PipelineLogInterface interface {
	ObjectClient() *clientbase.ObjectClient
	Create(*PipelineLog) (*PipelineLog, error)
	GetNamespaced(namespace, name string, opts metav1.GetOptions) (*PipelineLog, error)
	Get(name string, opts metav1.GetOptions) (*PipelineLog, error)
	Update(*PipelineLog) (*PipelineLog, error)
	Delete(name string, options *metav1.DeleteOptions) error
	DeleteNamespaced(namespace, name string, options *metav1.DeleteOptions) error
	List(opts metav1.ListOptions) (*PipelineLogList, error)
	Watch(opts metav1.ListOptions) (watch.Interface, error)
	DeleteCollection(deleteOpts *metav1.DeleteOptions, listOpts metav1.ListOptions) error
	Controller() PipelineLogController
	AddHandler(name string, sync PipelineLogHandlerFunc)
	AddLifecycle(name string, lifecycle PipelineLogLifecycle)
	AddClusterScopedHandler(name, clusterName string, sync PipelineLogHandlerFunc)
	AddClusterScopedLifecycle(name, clusterName string, lifecycle PipelineLogLifecycle)
}

type pipelineLogLister struct {
	controller *pipelineLogController
}

func (l *pipelineLogLister) List(namespace string, selector labels.Selector) (ret []*PipelineLog, err error) {
	err = cache.ListAllByNamespace(l.controller.Informer().GetIndexer(), namespace, selector, func(obj interface{}) {
		ret = append(ret, obj.(*PipelineLog))
	})
	return
}

func (l *pipelineLogLister) Get(namespace, name string) (*PipelineLog, error) {
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
			Group:    PipelineLogGroupVersionKind.Group,
			Resource: "pipelineLog",
		}, name)
	}
	return obj.(*PipelineLog), nil
}

type pipelineLogController struct {
	controller.GenericController
}

func (c *pipelineLogController) Lister() PipelineLogLister {
	return &pipelineLogLister{
		controller: c,
	}
}

func (c *pipelineLogController) AddHandler(name string, handler PipelineLogHandlerFunc) {
	c.GenericController.AddHandler(name, func(key string) error {
		obj, exists, err := c.Informer().GetStore().GetByKey(key)
		if err != nil {
			return err
		}
		if !exists {
			return handler(key, nil)
		}
		return handler(key, obj.(*PipelineLog))
	})
}

func (c *pipelineLogController) AddClusterScopedHandler(name, cluster string, handler PipelineLogHandlerFunc) {
	c.GenericController.AddHandler(name, func(key string) error {
		obj, exists, err := c.Informer().GetStore().GetByKey(key)
		if err != nil {
			return err
		}
		if !exists {
			return handler(key, nil)
		}

		if !controller.ObjectInCluster(cluster, obj) {
			return nil
		}

		return handler(key, obj.(*PipelineLog))
	})
}

type pipelineLogFactory struct {
}

func (c pipelineLogFactory) Object() runtime.Object {
	return &PipelineLog{}
}

func (c pipelineLogFactory) List() runtime.Object {
	return &PipelineLogList{}
}

func (s *pipelineLogClient) Controller() PipelineLogController {
	s.client.Lock()
	defer s.client.Unlock()

	c, ok := s.client.pipelineLogControllers[s.ns]
	if ok {
		return c
	}

	genericController := controller.NewGenericController(PipelineLogGroupVersionKind.Kind+"Controller",
		s.objectClient)

	c = &pipelineLogController{
		GenericController: genericController,
	}

	s.client.pipelineLogControllers[s.ns] = c
	s.client.starters = append(s.client.starters, c)

	return c
}

type pipelineLogClient struct {
	client       *Client
	ns           string
	objectClient *clientbase.ObjectClient
	controller   PipelineLogController
}

func (s *pipelineLogClient) ObjectClient() *clientbase.ObjectClient {
	return s.objectClient
}

func (s *pipelineLogClient) Create(o *PipelineLog) (*PipelineLog, error) {
	obj, err := s.objectClient.Create(o)
	return obj.(*PipelineLog), err
}

func (s *pipelineLogClient) Get(name string, opts metav1.GetOptions) (*PipelineLog, error) {
	obj, err := s.objectClient.Get(name, opts)
	return obj.(*PipelineLog), err
}

func (s *pipelineLogClient) GetNamespaced(namespace, name string, opts metav1.GetOptions) (*PipelineLog, error) {
	obj, err := s.objectClient.GetNamespaced(namespace, name, opts)
	return obj.(*PipelineLog), err
}

func (s *pipelineLogClient) Update(o *PipelineLog) (*PipelineLog, error) {
	obj, err := s.objectClient.Update(o.Name, o)
	return obj.(*PipelineLog), err
}

func (s *pipelineLogClient) Delete(name string, options *metav1.DeleteOptions) error {
	return s.objectClient.Delete(name, options)
}

func (s *pipelineLogClient) DeleteNamespaced(namespace, name string, options *metav1.DeleteOptions) error {
	return s.objectClient.DeleteNamespaced(namespace, name, options)
}

func (s *pipelineLogClient) List(opts metav1.ListOptions) (*PipelineLogList, error) {
	obj, err := s.objectClient.List(opts)
	return obj.(*PipelineLogList), err
}

func (s *pipelineLogClient) Watch(opts metav1.ListOptions) (watch.Interface, error) {
	return s.objectClient.Watch(opts)
}

// Patch applies the patch and returns the patched deployment.
func (s *pipelineLogClient) Patch(o *PipelineLog, data []byte, subresources ...string) (*PipelineLog, error) {
	obj, err := s.objectClient.Patch(o.Name, o, data, subresources...)
	return obj.(*PipelineLog), err
}

func (s *pipelineLogClient) DeleteCollection(deleteOpts *metav1.DeleteOptions, listOpts metav1.ListOptions) error {
	return s.objectClient.DeleteCollection(deleteOpts, listOpts)
}

func (s *pipelineLogClient) AddHandler(name string, sync PipelineLogHandlerFunc) {
	s.Controller().AddHandler(name, sync)
}

func (s *pipelineLogClient) AddLifecycle(name string, lifecycle PipelineLogLifecycle) {
	sync := NewPipelineLogLifecycleAdapter(name, false, s, lifecycle)
	s.AddHandler(name, sync)
}

func (s *pipelineLogClient) AddClusterScopedHandler(name, clusterName string, sync PipelineLogHandlerFunc) {
	s.Controller().AddClusterScopedHandler(name, clusterName, sync)
}

func (s *pipelineLogClient) AddClusterScopedLifecycle(name, clusterName string, lifecycle PipelineLogLifecycle) {
	sync := NewPipelineLogLifecycleAdapter(name+"_"+clusterName, true, s, lifecycle)
	s.AddClusterScopedHandler(name, clusterName, sync)
}
