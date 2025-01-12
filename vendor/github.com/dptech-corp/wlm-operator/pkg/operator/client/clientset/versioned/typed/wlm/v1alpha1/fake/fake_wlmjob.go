// Copyright (c) 2019 Sylabs, Inc. All rights reserved.
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

// Code generated by main. DO NOT EDIT.

package fake

import (
	v1alpha1 "github.com/dptech-corp/wlm-operator/pkg/operator/apis/wlm/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeWlmJobs implements WlmJobInterface
type FakeWlmJobs struct {
	Fake *FakeWlmV1alpha1
	ns   string
}

var wlmjobsResource = schema.GroupVersionResource{Group: "wlm.sylabs.io", Version: "v1alpha1", Resource: "wlmjobs"}

var wlmjobsKind = schema.GroupVersionKind{Group: "wlm.sylabs.io", Version: "v1alpha1", Kind: "WlmJob"}

// Get takes name of the wlmJob, and returns the corresponding wlmJob object, and an error if there is any.
func (c *FakeWlmJobs) Get(name string, options v1.GetOptions) (result *v1alpha1.WlmJob, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(wlmjobsResource, c.ns, name), &v1alpha1.WlmJob{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.WlmJob), err
}

// List takes label and field selectors, and returns the list of WlmJobs that match those selectors.
func (c *FakeWlmJobs) List(opts v1.ListOptions) (result *v1alpha1.WlmJobList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(wlmjobsResource, wlmjobsKind, c.ns, opts), &v1alpha1.WlmJobList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.WlmJobList{ListMeta: obj.(*v1alpha1.WlmJobList).ListMeta}
	for _, item := range obj.(*v1alpha1.WlmJobList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested wlmJobs.
func (c *FakeWlmJobs) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(wlmjobsResource, c.ns, opts))

}

// Create takes the representation of a wlmJob and creates it.  Returns the server's representation of the wlmJob, and an error, if there is any.
func (c *FakeWlmJobs) Create(wlmJob *v1alpha1.WlmJob) (result *v1alpha1.WlmJob, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(wlmjobsResource, c.ns, wlmJob), &v1alpha1.WlmJob{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.WlmJob), err
}

// Update takes the representation of a wlmJob and updates it. Returns the server's representation of the wlmJob, and an error, if there is any.
func (c *FakeWlmJobs) Update(wlmJob *v1alpha1.WlmJob) (result *v1alpha1.WlmJob, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(wlmjobsResource, c.ns, wlmJob), &v1alpha1.WlmJob{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.WlmJob), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeWlmJobs) UpdateStatus(wlmJob *v1alpha1.WlmJob) (*v1alpha1.WlmJob, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(wlmjobsResource, "status", c.ns, wlmJob), &v1alpha1.WlmJob{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.WlmJob), err
}

// Delete takes name of the wlmJob and deletes it. Returns an error if one occurs.
func (c *FakeWlmJobs) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(wlmjobsResource, c.ns, name), &v1alpha1.WlmJob{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeWlmJobs) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(wlmjobsResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.WlmJobList{})
	return err
}

// Patch applies the patch and returns the patched wlmJob.
func (c *FakeWlmJobs) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.WlmJob, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(wlmjobsResource, c.ns, name, pt, data, subresources...), &v1alpha1.WlmJob{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.WlmJob), err
}
