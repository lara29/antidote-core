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
	"time"

	v1 "github.com/nre-learning/antidote-core/pkg/apis/k8s.cni.cncf.io/v1"
	scheme "github.com/nre-learning/antidote-core/pkg/client/clientset/versioned/scheme"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// NetworkAttachmentDefinitionsGetter has a method to return a NetworkAttachmentDefinitionInterface.
// A group's client should implement this interface.
type NetworkAttachmentDefinitionsGetter interface {
	NetworkAttachmentDefinitions(namespace string) NetworkAttachmentDefinitionInterface
}

// NetworkAttachmentDefinitionInterface has methods to work with NetworkAttachmentDefinition resources.
type NetworkAttachmentDefinitionInterface interface {
	Create(*v1.NetworkAttachmentDefinition) (*v1.NetworkAttachmentDefinition, error)
	Update(*v1.NetworkAttachmentDefinition) (*v1.NetworkAttachmentDefinition, error)
	Delete(name string, options *metav1.DeleteOptions) error
	DeleteCollection(options *metav1.DeleteOptions, listOptions metav1.ListOptions) error
	Get(name string, options metav1.GetOptions) (*v1.NetworkAttachmentDefinition, error)
	List(opts metav1.ListOptions) (*v1.NetworkAttachmentDefinitionList, error)
	Watch(opts metav1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1.NetworkAttachmentDefinition, err error)
	NetworkAttachmentDefinitionExpansion
}

// networkAttachmentDefinitions implements NetworkAttachmentDefinitionInterface
type networkAttachmentDefinitions struct {
	client rest.Interface
	ns     string
}

// newNetworkAttachmentDefinitions returns a NetworkAttachmentDefinitions
func newNetworkAttachmentDefinitions(c *K8sV1Client, namespace string) *networkAttachmentDefinitions {
	return &networkAttachmentDefinitions{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the networkAttachmentDefinition, and returns the corresponding networkAttachmentDefinition object, and an error if there is any.
func (c *networkAttachmentDefinitions) Get(name string, options metav1.GetOptions) (result *v1.NetworkAttachmentDefinition, err error) {
	result = &v1.NetworkAttachmentDefinition{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("network-attachment-definitions").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of NetworkAttachmentDefinitions that match those selectors.
func (c *networkAttachmentDefinitions) List(opts metav1.ListOptions) (result *v1.NetworkAttachmentDefinitionList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1.NetworkAttachmentDefinitionList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("network-attachment-definitions").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested networkAttachmentDefinitions.
func (c *networkAttachmentDefinitions) Watch(opts metav1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("network-attachment-definitions").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a networkAttachmentDefinition and creates it.  Returns the server's representation of the networkAttachmentDefinition, and an error, if there is any.
func (c *networkAttachmentDefinitions) Create(networkAttachmentDefinition *v1.NetworkAttachmentDefinition) (result *v1.NetworkAttachmentDefinition, err error) {
	result = &v1.NetworkAttachmentDefinition{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("network-attachment-definitions").
		Body(networkAttachmentDefinition).
		Do().
		Into(result)
	return
}

// Update takes the representation of a networkAttachmentDefinition and updates it. Returns the server's representation of the networkAttachmentDefinition, and an error, if there is any.
func (c *networkAttachmentDefinitions) Update(networkAttachmentDefinition *v1.NetworkAttachmentDefinition) (result *v1.NetworkAttachmentDefinition, err error) {
	result = &v1.NetworkAttachmentDefinition{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("network-attachment-definitions").
		Name(networkAttachmentDefinition.Name).
		Body(networkAttachmentDefinition).
		Do().
		Into(result)
	return
}

// Delete takes name of the networkAttachmentDefinition and deletes it. Returns an error if one occurs.
func (c *networkAttachmentDefinitions) Delete(name string, options *metav1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("network-attachment-definitions").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *networkAttachmentDefinitions) DeleteCollection(options *metav1.DeleteOptions, listOptions metav1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("network-attachment-definitions").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched networkAttachmentDefinition.
func (c *networkAttachmentDefinitions) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1.NetworkAttachmentDefinition, err error) {
	result = &v1.NetworkAttachmentDefinition{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("network-attachment-definitions").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
