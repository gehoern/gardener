package v1beta1

import (
	v1beta1 "github.com/gardener/gardener/pkg/apis/garden/v1beta1"
	scheme "github.com/gardener/gardener/pkg/client/garden/clientset/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// SeedsGetter has a method to return a SeedInterface.
// A group's client should implement this interface.
type SeedsGetter interface {
	Seeds() SeedInterface
}

// SeedInterface has methods to work with Seed resources.
type SeedInterface interface {
	Create(*v1beta1.Seed) (*v1beta1.Seed, error)
	Update(*v1beta1.Seed) (*v1beta1.Seed, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1beta1.Seed, error)
	List(opts v1.ListOptions) (*v1beta1.SeedList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1beta1.Seed, err error)
	SeedExpansion
}

// seeds implements SeedInterface
type seeds struct {
	client rest.Interface
}

// newSeeds returns a Seeds
func newSeeds(c *GardenV1beta1Client) *seeds {
	return &seeds{
		client: c.RESTClient(),
	}
}

// Get takes name of the seed, and returns the corresponding seed object, and an error if there is any.
func (c *seeds) Get(name string, options v1.GetOptions) (result *v1beta1.Seed, err error) {
	result = &v1beta1.Seed{}
	err = c.client.Get().
		Resource("seeds").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of Seeds that match those selectors.
func (c *seeds) List(opts v1.ListOptions) (result *v1beta1.SeedList, err error) {
	result = &v1beta1.SeedList{}
	err = c.client.Get().
		Resource("seeds").
		VersionedParams(&opts, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested seeds.
func (c *seeds) Watch(opts v1.ListOptions) (watch.Interface, error) {
	opts.Watch = true
	return c.client.Get().
		Resource("seeds").
		VersionedParams(&opts, scheme.ParameterCodec).
		Watch()
}

// Create takes the representation of a seed and creates it.  Returns the server's representation of the seed, and an error, if there is any.
func (c *seeds) Create(seed *v1beta1.Seed) (result *v1beta1.Seed, err error) {
	result = &v1beta1.Seed{}
	err = c.client.Post().
		Resource("seeds").
		Body(seed).
		Do().
		Into(result)
	return
}

// Update takes the representation of a seed and updates it. Returns the server's representation of the seed, and an error, if there is any.
func (c *seeds) Update(seed *v1beta1.Seed) (result *v1beta1.Seed, err error) {
	result = &v1beta1.Seed{}
	err = c.client.Put().
		Resource("seeds").
		Name(seed.Name).
		Body(seed).
		Do().
		Into(result)
	return
}

// Delete takes name of the seed and deletes it. Returns an error if one occurs.
func (c *seeds) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Resource("seeds").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *seeds) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	return c.client.Delete().
		Resource("seeds").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched seed.
func (c *seeds) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1beta1.Seed, err error) {
	result = &v1beta1.Seed{}
	err = c.client.Patch(pt).
		Resource("seeds").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
