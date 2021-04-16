// Code generated by informer-gen. DO NOT EDIT.

package v1alpha1

import (
	"context"
	time "time"

	workflowv1alpha1 "github.com/argoproj/argo-workflows/v3/pkg/apis/workflow/v1alpha1"
	versioned "github.com/argoproj/argo-workflows/v3/pkg/client/clientset/versioned"
	internalinterfaces "github.com/argoproj/argo-workflows/v3/pkg/client/informers/externalversions/internalinterfaces"
	v1alpha1 "github.com/argoproj/argo-workflows/v3/pkg/client/listers/workflow/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// WorkflowTaskSetInformer provides access to a shared informer and lister for
// WorkflowTaskSets.
type WorkflowTaskSetInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1alpha1.WorkflowTaskSetLister
}

type workflowTaskSetInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewWorkflowTaskSetInformer constructs a new informer for WorkflowTaskSet type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewWorkflowTaskSetInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredWorkflowTaskSetInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredWorkflowTaskSetInformer constructs a new informer for WorkflowTaskSet type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredWorkflowTaskSetInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.ArgoprojV1alpha1().WorkflowTaskSets(namespace).List(context.TODO(), options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.ArgoprojV1alpha1().WorkflowTaskSets(namespace).Watch(context.TODO(), options)
			},
		},
		&workflowv1alpha1.WorkflowTaskSet{},
		resyncPeriod,
		indexers,
	)
}

func (f *workflowTaskSetInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredWorkflowTaskSetInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *workflowTaskSetInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&workflowv1alpha1.WorkflowTaskSet{}, f.defaultInformer)
}

func (f *workflowTaskSetInformer) Lister() v1alpha1.WorkflowTaskSetLister {
	return v1alpha1.NewWorkflowTaskSetLister(f.Informer().GetIndexer())
}
