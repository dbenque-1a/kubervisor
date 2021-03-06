/*
MIT License

Copyright (c) 2018 Kubervisor

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
*/

// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	v1alpha1 "github.com/amadeusitgroup/kubervisor/pkg/api/kubervisor/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeKubervisorServices implements KubervisorServiceInterface
type FakeKubervisorServices struct {
	Fake *FakeKubervisorV1alpha1
	ns   string
}

var kubervisorservicesResource = schema.GroupVersionResource{Group: "kubervisor.k8s.io", Version: "v1alpha1", Resource: "kubervisorservices"}

var kubervisorservicesKind = schema.GroupVersionKind{Group: "kubervisor.k8s.io", Version: "v1alpha1", Kind: "KubervisorService"}

// Get takes name of the kubervisorService, and returns the corresponding kubervisorService object, and an error if there is any.
func (c *FakeKubervisorServices) Get(name string, options v1.GetOptions) (result *v1alpha1.KubervisorService, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(kubervisorservicesResource, c.ns, name), &v1alpha1.KubervisorService{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.KubervisorService), err
}

// List takes label and field selectors, and returns the list of KubervisorServices that match those selectors.
func (c *FakeKubervisorServices) List(opts v1.ListOptions) (result *v1alpha1.KubervisorServiceList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(kubervisorservicesResource, kubervisorservicesKind, c.ns, opts), &v1alpha1.KubervisorServiceList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.KubervisorServiceList{}
	for _, item := range obj.(*v1alpha1.KubervisorServiceList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested kubervisorServices.
func (c *FakeKubervisorServices) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(kubervisorservicesResource, c.ns, opts))

}

// Create takes the representation of a kubervisorService and creates it.  Returns the server's representation of the kubervisorService, and an error, if there is any.
func (c *FakeKubervisorServices) Create(kubervisorService *v1alpha1.KubervisorService) (result *v1alpha1.KubervisorService, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(kubervisorservicesResource, c.ns, kubervisorService), &v1alpha1.KubervisorService{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.KubervisorService), err
}

// Update takes the representation of a kubervisorService and updates it. Returns the server's representation of the kubervisorService, and an error, if there is any.
func (c *FakeKubervisorServices) Update(kubervisorService *v1alpha1.KubervisorService) (result *v1alpha1.KubervisorService, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(kubervisorservicesResource, c.ns, kubervisorService), &v1alpha1.KubervisorService{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.KubervisorService), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeKubervisorServices) UpdateStatus(kubervisorService *v1alpha1.KubervisorService) (*v1alpha1.KubervisorService, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(kubervisorservicesResource, "status", c.ns, kubervisorService), &v1alpha1.KubervisorService{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.KubervisorService), err
}

// Delete takes name of the kubervisorService and deletes it. Returns an error if one occurs.
func (c *FakeKubervisorServices) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(kubervisorservicesResource, c.ns, name), &v1alpha1.KubervisorService{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeKubervisorServices) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(kubervisorservicesResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.KubervisorServiceList{})
	return err
}

// Patch applies the patch and returns the patched kubervisorService.
func (c *FakeKubervisorServices) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.KubervisorService, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(kubervisorservicesResource, c.ns, name, data, subresources...), &v1alpha1.KubervisorService{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.KubervisorService), err
}
