package v1alpha3

import (
	v1alpha3 "github.com/heptio/quartermaster/custom/apis/virtualservice/v1alpha3"
	scheme "github.com/heptio/quartermaster/custom/client/clientset/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// VirtualServicesGetter has a method to return a VirtualServiceInterface.
// A group's client should implement this interface.
type VirtualServicesGetter interface {
	VirtualServices(namespace string) VirtualServiceInterface
}

// VirtualServiceInterface has methods to work with VirtualService resources.
type VirtualServiceInterface interface {
	Create(*v1alpha3.VirtualService) (*v1alpha3.VirtualService, error)
	Update(*v1alpha3.VirtualService) (*v1alpha3.VirtualService, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha3.VirtualService, error)
	List(opts v1.ListOptions) (*v1alpha3.VirtualServiceList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha3.VirtualService, err error)
	VirtualServiceExpansion
}

// virtualServices implements VirtualServiceInterface
type virtualServices struct {
	client rest.Interface
	ns     string
}

// newVirtualServices returns a VirtualServices
func newVirtualServices(c *VirtualserviceV1alpha3Client, namespace string) *virtualServices {
	return &virtualServices{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the virtualService, and returns the corresponding virtualService object, and an error if there is any.
func (c *virtualServices) Get(name string, options v1.GetOptions) (result *v1alpha3.VirtualService, err error) {
	result = &v1alpha3.VirtualService{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("virtualservices").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of VirtualServices that match those selectors.
func (c *virtualServices) List(opts v1.ListOptions) (result *v1alpha3.VirtualServiceList, err error) {
	result = &v1alpha3.VirtualServiceList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("virtualservices").
		VersionedParams(&opts, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested virtualServices.
func (c *virtualServices) Watch(opts v1.ListOptions) (watch.Interface, error) {
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("virtualservices").
		VersionedParams(&opts, scheme.ParameterCodec).
		Watch()
}

// Create takes the representation of a virtualService and creates it.  Returns the server's representation of the virtualService, and an error, if there is any.
func (c *virtualServices) Create(virtualService *v1alpha3.VirtualService) (result *v1alpha3.VirtualService, err error) {
	result = &v1alpha3.VirtualService{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("virtualservices").
		Body(virtualService).
		Do().
		Into(result)
	return
}

// Update takes the representation of a virtualService and updates it. Returns the server's representation of the virtualService, and an error, if there is any.
func (c *virtualServices) Update(virtualService *v1alpha3.VirtualService) (result *v1alpha3.VirtualService, err error) {
	result = &v1alpha3.VirtualService{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("virtualservices").
		Name(virtualService.Name).
		Body(virtualService).
		Do().
		Into(result)
	return
}

// Delete takes name of the virtualService and deletes it. Returns an error if one occurs.
func (c *virtualServices) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("virtualservices").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *virtualServices) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("virtualservices").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched virtualService.
func (c *virtualServices) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha3.VirtualService, err error) {
	result = &v1alpha3.VirtualService{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("virtualservices").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
