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

// Code generated by client-gen. DO NOT EDIT.

package v1

import (
	"context"
	"time"

	v1 "github.com/parochi/greeting-client/pkg/apis/parochi/v1"
	scheme "github.com/parochi/greeting-client/pkg/client/clientset/versioned/scheme"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// GreetingsGetter has a method to return a GreetingInterface.
// A group's client should implement this interface.
type GreetingsGetter interface {
	Greetings(namespace string) GreetingInterface
}

// GreetingInterface has methods to work with Greeting resources.
type GreetingInterface interface {
	Create(ctx context.Context, greeting *v1.Greeting, opts metav1.CreateOptions) (*v1.Greeting, error)
	Update(ctx context.Context, greeting *v1.Greeting, opts metav1.UpdateOptions) (*v1.Greeting, error)
	UpdateStatus(ctx context.Context, greeting *v1.Greeting, opts metav1.UpdateOptions) (*v1.Greeting, error)
	Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error
	Get(ctx context.Context, name string, opts metav1.GetOptions) (*v1.Greeting, error)
	List(ctx context.Context, opts metav1.ListOptions) (*v1.GreetingList, error)
	Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.Greeting, err error)
	GreetingExpansion
}

// greetings implements GreetingInterface
type greetings struct {
	client rest.Interface
	ns     string
}

// newGreetings returns a Greetings
func newGreetings(c *ParochiV1Client, namespace string) *greetings {
	return &greetings{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the greeting, and returns the corresponding greeting object, and an error if there is any.
func (c *greetings) Get(ctx context.Context, name string, options metav1.GetOptions) (result *v1.Greeting, err error) {
	result = &v1.Greeting{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("greetings").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of Greetings that match those selectors.
func (c *greetings) List(ctx context.Context, opts metav1.ListOptions) (result *v1.GreetingList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1.GreetingList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("greetings").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested greetings.
func (c *greetings) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("greetings").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a greeting and creates it.  Returns the server's representation of the greeting, and an error, if there is any.
func (c *greetings) Create(ctx context.Context, greeting *v1.Greeting, opts metav1.CreateOptions) (result *v1.Greeting, err error) {
	result = &v1.Greeting{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("greetings").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(greeting).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a greeting and updates it. Returns the server's representation of the greeting, and an error, if there is any.
func (c *greetings) Update(ctx context.Context, greeting *v1.Greeting, opts metav1.UpdateOptions) (result *v1.Greeting, err error) {
	result = &v1.Greeting{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("greetings").
		Name(greeting.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(greeting).
		Do(ctx).
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *greetings) UpdateStatus(ctx context.Context, greeting *v1.Greeting, opts metav1.UpdateOptions) (result *v1.Greeting, err error) {
	result = &v1.Greeting{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("greetings").
		Name(greeting.Name).
		SubResource("status").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(greeting).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the greeting and deletes it. Returns an error if one occurs.
func (c *greetings) Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("greetings").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *greetings) DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("greetings").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched greeting.
func (c *greetings) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.Greeting, err error) {
	result = &v1.Greeting{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("greetings").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}