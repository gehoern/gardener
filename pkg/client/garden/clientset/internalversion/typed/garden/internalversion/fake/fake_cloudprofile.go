package fake

import (
	garden "github.com/gardener/gardener/pkg/apis/garden"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeCloudProfiles implements CloudProfileInterface
type FakeCloudProfiles struct {
	Fake *FakeGarden
}

var cloudprofilesResource = schema.GroupVersionResource{Group: "garden.sapcloud.io", Version: "", Resource: "cloudprofiles"}

var cloudprofilesKind = schema.GroupVersionKind{Group: "garden.sapcloud.io", Version: "", Kind: "CloudProfile"}

// Get takes name of the cloudProfile, and returns the corresponding cloudProfile object, and an error if there is any.
func (c *FakeCloudProfiles) Get(name string, options v1.GetOptions) (result *garden.CloudProfile, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(cloudprofilesResource, name), &garden.CloudProfile{})
	if obj == nil {
		return nil, err
	}
	return obj.(*garden.CloudProfile), err
}

// List takes label and field selectors, and returns the list of CloudProfiles that match those selectors.
func (c *FakeCloudProfiles) List(opts v1.ListOptions) (result *garden.CloudProfileList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(cloudprofilesResource, cloudprofilesKind, opts), &garden.CloudProfileList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &garden.CloudProfileList{}
	for _, item := range obj.(*garden.CloudProfileList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested cloudProfiles.
func (c *FakeCloudProfiles) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(cloudprofilesResource, opts))
}

// Create takes the representation of a cloudProfile and creates it.  Returns the server's representation of the cloudProfile, and an error, if there is any.
func (c *FakeCloudProfiles) Create(cloudProfile *garden.CloudProfile) (result *garden.CloudProfile, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(cloudprofilesResource, cloudProfile), &garden.CloudProfile{})
	if obj == nil {
		return nil, err
	}
	return obj.(*garden.CloudProfile), err
}

// Update takes the representation of a cloudProfile and updates it. Returns the server's representation of the cloudProfile, and an error, if there is any.
func (c *FakeCloudProfiles) Update(cloudProfile *garden.CloudProfile) (result *garden.CloudProfile, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(cloudprofilesResource, cloudProfile), &garden.CloudProfile{})
	if obj == nil {
		return nil, err
	}
	return obj.(*garden.CloudProfile), err
}

// Delete takes name of the cloudProfile and deletes it. Returns an error if one occurs.
func (c *FakeCloudProfiles) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteAction(cloudprofilesResource, name), &garden.CloudProfile{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeCloudProfiles) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(cloudprofilesResource, listOptions)

	_, err := c.Fake.Invokes(action, &garden.CloudProfileList{})
	return err
}

// Patch applies the patch and returns the patched cloudProfile.
func (c *FakeCloudProfiles) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *garden.CloudProfile, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(cloudprofilesResource, name, data, subresources...), &garden.CloudProfile{})
	if obj == nil {
		return nil, err
	}
	return obj.(*garden.CloudProfile), err
}
