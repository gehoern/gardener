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

// FakeSeeds implements SeedInterface
type FakeSeeds struct {
	Fake *FakeGarden
}

var seedsResource = schema.GroupVersionResource{Group: "garden.sapcloud.io", Version: "", Resource: "seeds"}

var seedsKind = schema.GroupVersionKind{Group: "garden.sapcloud.io", Version: "", Kind: "Seed"}

// Get takes name of the seed, and returns the corresponding seed object, and an error if there is any.
func (c *FakeSeeds) Get(name string, options v1.GetOptions) (result *garden.Seed, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(seedsResource, name), &garden.Seed{})
	if obj == nil {
		return nil, err
	}
	return obj.(*garden.Seed), err
}

// List takes label and field selectors, and returns the list of Seeds that match those selectors.
func (c *FakeSeeds) List(opts v1.ListOptions) (result *garden.SeedList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(seedsResource, seedsKind, opts), &garden.SeedList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &garden.SeedList{}
	for _, item := range obj.(*garden.SeedList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested seeds.
func (c *FakeSeeds) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(seedsResource, opts))
}

// Create takes the representation of a seed and creates it.  Returns the server's representation of the seed, and an error, if there is any.
func (c *FakeSeeds) Create(seed *garden.Seed) (result *garden.Seed, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(seedsResource, seed), &garden.Seed{})
	if obj == nil {
		return nil, err
	}
	return obj.(*garden.Seed), err
}

// Update takes the representation of a seed and updates it. Returns the server's representation of the seed, and an error, if there is any.
func (c *FakeSeeds) Update(seed *garden.Seed) (result *garden.Seed, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(seedsResource, seed), &garden.Seed{})
	if obj == nil {
		return nil, err
	}
	return obj.(*garden.Seed), err
}

// Delete takes name of the seed and deletes it. Returns an error if one occurs.
func (c *FakeSeeds) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteAction(seedsResource, name), &garden.Seed{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeSeeds) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(seedsResource, listOptions)

	_, err := c.Fake.Invokes(action, &garden.SeedList{})
	return err
}

// Patch applies the patch and returns the patched seed.
func (c *FakeSeeds) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *garden.Seed, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(seedsResource, name, data, subresources...), &garden.Seed{})
	if obj == nil {
		return nil, err
	}
	return obj.(*garden.Seed), err
}
