/*
Copyright 2020 The OpenEBS Authors

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

	v1 "github.com/openebs/api/pkg/apis/cstor/v1"
	scheme "github.com/openebs/api/pkg/client/clientset/versioned/scheme"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// CStorVolumeClaimsGetter has a method to return a CStorVolumeClaimInterface.
// A group's client should implement this interface.
type CStorVolumeClaimsGetter interface {
	CStorVolumeClaims(namespace string) CStorVolumeClaimInterface
}

// CStorVolumeClaimInterface has methods to work with CStorVolumeClaim resources.
type CStorVolumeClaimInterface interface {
	Create(*v1.CStorVolumeClaim) (*v1.CStorVolumeClaim, error)
	Update(*v1.CStorVolumeClaim) (*v1.CStorVolumeClaim, error)
	UpdateStatus(*v1.CStorVolumeClaim) (*v1.CStorVolumeClaim, error)
	Delete(name string, options *metav1.DeleteOptions) error
	DeleteCollection(options *metav1.DeleteOptions, listOptions metav1.ListOptions) error
	Get(name string, options metav1.GetOptions) (*v1.CStorVolumeClaim, error)
	List(opts metav1.ListOptions) (*v1.CStorVolumeClaimList, error)
	Watch(opts metav1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1.CStorVolumeClaim, err error)
	CStorVolumeClaimExpansion
}

// cStorVolumeClaims implements CStorVolumeClaimInterface
type cStorVolumeClaims struct {
	client rest.Interface
	ns     string
}

// newCStorVolumeClaims returns a CStorVolumeClaims
func newCStorVolumeClaims(c *CstorV1Client, namespace string) *cStorVolumeClaims {
	return &cStorVolumeClaims{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the cStorVolumeClaim, and returns the corresponding cStorVolumeClaim object, and an error if there is any.
func (c *cStorVolumeClaims) Get(name string, options metav1.GetOptions) (result *v1.CStorVolumeClaim, err error) {
	result = &v1.CStorVolumeClaim{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("cstorvolumeclaims").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of CStorVolumeClaims that match those selectors.
func (c *cStorVolumeClaims) List(opts metav1.ListOptions) (result *v1.CStorVolumeClaimList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1.CStorVolumeClaimList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("cstorvolumeclaims").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested cStorVolumeClaims.
func (c *cStorVolumeClaims) Watch(opts metav1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("cstorvolumeclaims").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a cStorVolumeClaim and creates it.  Returns the server's representation of the cStorVolumeClaim, and an error, if there is any.
func (c *cStorVolumeClaims) Create(cStorVolumeClaim *v1.CStorVolumeClaim) (result *v1.CStorVolumeClaim, err error) {
	result = &v1.CStorVolumeClaim{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("cstorvolumeclaims").
		Body(cStorVolumeClaim).
		Do().
		Into(result)
	return
}

// Update takes the representation of a cStorVolumeClaim and updates it. Returns the server's representation of the cStorVolumeClaim, and an error, if there is any.
func (c *cStorVolumeClaims) Update(cStorVolumeClaim *v1.CStorVolumeClaim) (result *v1.CStorVolumeClaim, err error) {
	result = &v1.CStorVolumeClaim{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("cstorvolumeclaims").
		Name(cStorVolumeClaim.Name).
		Body(cStorVolumeClaim).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *cStorVolumeClaims) UpdateStatus(cStorVolumeClaim *v1.CStorVolumeClaim) (result *v1.CStorVolumeClaim, err error) {
	result = &v1.CStorVolumeClaim{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("cstorvolumeclaims").
		Name(cStorVolumeClaim.Name).
		SubResource("status").
		Body(cStorVolumeClaim).
		Do().
		Into(result)
	return
}

// Delete takes name of the cStorVolumeClaim and deletes it. Returns an error if one occurs.
func (c *cStorVolumeClaims) Delete(name string, options *metav1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("cstorvolumeclaims").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *cStorVolumeClaims) DeleteCollection(options *metav1.DeleteOptions, listOptions metav1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("cstorvolumeclaims").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched cStorVolumeClaim.
func (c *cStorVolumeClaims) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1.CStorVolumeClaim, err error) {
	result = &v1.CStorVolumeClaim{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("cstorvolumeclaims").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
