/*
Copyright The Kubernetes Authors.

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

// Code generated by informer-gen. DO NOT EDIT.

package v1

import (
	time "time"

	k8scnicncfiov1 "github.com/nre-learning/antidote-core/pkg/apis/k8s.cni.cncf.io/v1"
	versioned "github.com/nre-learning/antidote-core/pkg/client/clientset/versioned"
	internalinterfaces "github.com/nre-learning/antidote-core/pkg/client/informers/externalversions/internalinterfaces"
	v1 "github.com/nre-learning/antidote-core/pkg/client/listers/k8s.cni.cncf.io/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// NetworkAttachmentDefinitionInformer provides access to a shared informer and lister for
// NetworkAttachmentDefinitions.
type NetworkAttachmentDefinitionInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1.NetworkAttachmentDefinitionLister
}

type networkAttachmentDefinitionInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewNetworkAttachmentDefinitionInformer constructs a new informer for NetworkAttachmentDefinition type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewNetworkAttachmentDefinitionInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredNetworkAttachmentDefinitionInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredNetworkAttachmentDefinitionInformer constructs a new informer for NetworkAttachmentDefinition type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredNetworkAttachmentDefinitionInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options metav1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.K8sV1().NetworkAttachmentDefinitions(namespace).List(options)
			},
			WatchFunc: func(options metav1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.K8sV1().NetworkAttachmentDefinitions(namespace).Watch(options)
			},
		},
		&k8scnicncfiov1.NetworkAttachmentDefinition{},
		resyncPeriod,
		indexers,
	)
}

func (f *networkAttachmentDefinitionInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredNetworkAttachmentDefinitionInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *networkAttachmentDefinitionInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&k8scnicncfiov1.NetworkAttachmentDefinition{}, f.defaultInformer)
}

func (f *networkAttachmentDefinitionInformer) Lister() v1.NetworkAttachmentDefinitionLister {
	return v1.NewNetworkAttachmentDefinitionLister(f.Informer().GetIndexer())
}
