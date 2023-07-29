// Copyright Aeraki Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by informer-gen. DO NOT EDIT.

package v1alpha1

import (
	"context"
	time "time"

	dubbov1alpha1 "github.com/aeraki-mesh/client-go/pkg/apis/dubbo/v1alpha1"
	versioned "github.com/aeraki-mesh/client-go/pkg/clientset/versioned"
	internalinterfaces "github.com/aeraki-mesh/client-go/pkg/informers/externalversions/internalinterfaces"
	v1alpha1 "github.com/aeraki-mesh/client-go/pkg/listers/dubbo/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// DubboAuthorizationPolicyInformer provides access to a shared informer and lister for
// DubboAuthorizationPolicies.
type DubboAuthorizationPolicyInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1alpha1.DubboAuthorizationPolicyLister
}

type dubboAuthorizationPolicyInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewDubboAuthorizationPolicyInformer constructs a new informer for DubboAuthorizationPolicy type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewDubboAuthorizationPolicyInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredDubboAuthorizationPolicyInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredDubboAuthorizationPolicyInformer constructs a new informer for DubboAuthorizationPolicy type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredDubboAuthorizationPolicyInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.DubboV1alpha1().DubboAuthorizationPolicies(namespace).List(context.TODO(), options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.DubboV1alpha1().DubboAuthorizationPolicies(namespace).Watch(context.TODO(), options)
			},
		},
		&dubbov1alpha1.DubboAuthorizationPolicy{},
		resyncPeriod,
		indexers,
	)
}

func (f *dubboAuthorizationPolicyInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredDubboAuthorizationPolicyInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *dubboAuthorizationPolicyInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&dubbov1alpha1.DubboAuthorizationPolicy{}, f.defaultInformer)
}

func (f *dubboAuthorizationPolicyInformer) Lister() v1alpha1.DubboAuthorizationPolicyLister {
	return v1alpha1.NewDubboAuthorizationPolicyLister(f.Informer().GetIndexer())
}
