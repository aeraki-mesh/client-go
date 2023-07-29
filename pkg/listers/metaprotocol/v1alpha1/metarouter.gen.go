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

// Code generated by lister-gen. DO NOT EDIT.

package v1alpha1

import (
	v1alpha1 "github.com/aeraki-mesh/client-go/pkg/apis/metaprotocol/v1alpha1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// MetaRouterLister helps list MetaRouters.
// All objects returned here must be treated as read-only.
type MetaRouterLister interface {
	// List lists all MetaRouters in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.MetaRouter, err error)
	// MetaRouters returns an object that can list and get MetaRouters.
	MetaRouters(namespace string) MetaRouterNamespaceLister
	MetaRouterListerExpansion
}

// metaRouterLister implements the MetaRouterLister interface.
type metaRouterLister struct {
	indexer cache.Indexer
}

// NewMetaRouterLister returns a new MetaRouterLister.
func NewMetaRouterLister(indexer cache.Indexer) MetaRouterLister {
	return &metaRouterLister{indexer: indexer}
}

// List lists all MetaRouters in the indexer.
func (s *metaRouterLister) List(selector labels.Selector) (ret []*v1alpha1.MetaRouter, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.MetaRouter))
	})
	return ret, err
}

// MetaRouters returns an object that can list and get MetaRouters.
func (s *metaRouterLister) MetaRouters(namespace string) MetaRouterNamespaceLister {
	return metaRouterNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// MetaRouterNamespaceLister helps list and get MetaRouters.
// All objects returned here must be treated as read-only.
type MetaRouterNamespaceLister interface {
	// List lists all MetaRouters in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.MetaRouter, err error)
	// Get retrieves the MetaRouter from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.MetaRouter, error)
	MetaRouterNamespaceListerExpansion
}

// metaRouterNamespaceLister implements the MetaRouterNamespaceLister
// interface.
type metaRouterNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all MetaRouters in the indexer for a given namespace.
func (s metaRouterNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.MetaRouter, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.MetaRouter))
	})
	return ret, err
}

// Get retrieves the MetaRouter from the indexer for a given namespace and name.
func (s metaRouterNamespaceLister) Get(name string) (*v1alpha1.MetaRouter, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("metarouter"), name)
	}
	return obj.(*v1alpha1.MetaRouter), nil
}
