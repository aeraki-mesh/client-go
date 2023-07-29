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

// Code generated by client-gen. DO NOT EDIT.

package v1alpha1

import (
	"context"
	json "encoding/json"
	"fmt"
	"time"

	v1alpha1 "github.com/aeraki-mesh/client-go/pkg/apis/metaprotocol/v1alpha1"
	metaprotocolv1alpha1 "github.com/aeraki-mesh/client-go/pkg/applyconfiguration/metaprotocol/v1alpha1"
	scheme "github.com/aeraki-mesh/client-go/pkg/clientset/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// MetaRoutersGetter has a method to return a MetaRouterInterface.
// A group's client should implement this interface.
type MetaRoutersGetter interface {
	MetaRouters(namespace string) MetaRouterInterface
}

// MetaRouterInterface has methods to work with MetaRouter resources.
type MetaRouterInterface interface {
	Create(ctx context.Context, metaRouter *v1alpha1.MetaRouter, opts v1.CreateOptions) (*v1alpha1.MetaRouter, error)
	Update(ctx context.Context, metaRouter *v1alpha1.MetaRouter, opts v1.UpdateOptions) (*v1alpha1.MetaRouter, error)
	UpdateStatus(ctx context.Context, metaRouter *v1alpha1.MetaRouter, opts v1.UpdateOptions) (*v1alpha1.MetaRouter, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v1alpha1.MetaRouter, error)
	List(ctx context.Context, opts v1.ListOptions) (*v1alpha1.MetaRouterList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.MetaRouter, err error)
	Apply(ctx context.Context, metaRouter *metaprotocolv1alpha1.MetaRouterApplyConfiguration, opts v1.ApplyOptions) (result *v1alpha1.MetaRouter, err error)
	ApplyStatus(ctx context.Context, metaRouter *metaprotocolv1alpha1.MetaRouterApplyConfiguration, opts v1.ApplyOptions) (result *v1alpha1.MetaRouter, err error)
	MetaRouterExpansion
}

// metaRouters implements MetaRouterInterface
type metaRouters struct {
	client rest.Interface
	ns     string
}

// newMetaRouters returns a MetaRouters
func newMetaRouters(c *MetaprotocolV1alpha1Client, namespace string) *metaRouters {
	return &metaRouters{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the metaRouter, and returns the corresponding metaRouter object, and an error if there is any.
func (c *metaRouters) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.MetaRouter, err error) {
	result = &v1alpha1.MetaRouter{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("metarouters").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of MetaRouters that match those selectors.
func (c *metaRouters) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.MetaRouterList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.MetaRouterList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("metarouters").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested metaRouters.
func (c *metaRouters) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("metarouters").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a metaRouter and creates it.  Returns the server's representation of the metaRouter, and an error, if there is any.
func (c *metaRouters) Create(ctx context.Context, metaRouter *v1alpha1.MetaRouter, opts v1.CreateOptions) (result *v1alpha1.MetaRouter, err error) {
	result = &v1alpha1.MetaRouter{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("metarouters").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(metaRouter).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a metaRouter and updates it. Returns the server's representation of the metaRouter, and an error, if there is any.
func (c *metaRouters) Update(ctx context.Context, metaRouter *v1alpha1.MetaRouter, opts v1.UpdateOptions) (result *v1alpha1.MetaRouter, err error) {
	result = &v1alpha1.MetaRouter{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("metarouters").
		Name(metaRouter.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(metaRouter).
		Do(ctx).
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *metaRouters) UpdateStatus(ctx context.Context, metaRouter *v1alpha1.MetaRouter, opts v1.UpdateOptions) (result *v1alpha1.MetaRouter, err error) {
	result = &v1alpha1.MetaRouter{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("metarouters").
		Name(metaRouter.Name).
		SubResource("status").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(metaRouter).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the metaRouter and deletes it. Returns an error if one occurs.
func (c *metaRouters) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("metarouters").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *metaRouters) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("metarouters").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched metaRouter.
func (c *metaRouters) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.MetaRouter, err error) {
	result = &v1alpha1.MetaRouter{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("metarouters").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}

// Apply takes the given apply declarative configuration, applies it and returns the applied metaRouter.
func (c *metaRouters) Apply(ctx context.Context, metaRouter *metaprotocolv1alpha1.MetaRouterApplyConfiguration, opts v1.ApplyOptions) (result *v1alpha1.MetaRouter, err error) {
	if metaRouter == nil {
		return nil, fmt.Errorf("metaRouter provided to Apply must not be nil")
	}
	patchOpts := opts.ToPatchOptions()
	data, err := json.Marshal(metaRouter)
	if err != nil {
		return nil, err
	}
	name := metaRouter.Name
	if name == nil {
		return nil, fmt.Errorf("metaRouter.Name must be provided to Apply")
	}
	result = &v1alpha1.MetaRouter{}
	err = c.client.Patch(types.ApplyPatchType).
		Namespace(c.ns).
		Resource("metarouters").
		Name(*name).
		VersionedParams(&patchOpts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}

// ApplyStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating ApplyStatus().
func (c *metaRouters) ApplyStatus(ctx context.Context, metaRouter *metaprotocolv1alpha1.MetaRouterApplyConfiguration, opts v1.ApplyOptions) (result *v1alpha1.MetaRouter, err error) {
	if metaRouter == nil {
		return nil, fmt.Errorf("metaRouter provided to Apply must not be nil")
	}
	patchOpts := opts.ToPatchOptions()
	data, err := json.Marshal(metaRouter)
	if err != nil {
		return nil, err
	}

	name := metaRouter.Name
	if name == nil {
		return nil, fmt.Errorf("metaRouter.Name must be provided to Apply")
	}

	result = &v1alpha1.MetaRouter{}
	err = c.client.Patch(types.ApplyPatchType).
		Namespace(c.ns).
		Resource("metarouters").
		Name(*name).
		SubResource("status").
		VersionedParams(&patchOpts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}