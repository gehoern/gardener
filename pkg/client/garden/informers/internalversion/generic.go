// This file was automatically generated by informer-gen

package internalversion

import (
	"fmt"

	garden "github.com/gardener/gardener/pkg/apis/garden"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	cache "k8s.io/client-go/tools/cache"
)

// GenericInformer is type of SharedIndexInformer which will locate and delegate to other
// sharedInformers based on type
type GenericInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() cache.GenericLister
}

type genericInformer struct {
	informer cache.SharedIndexInformer
	resource schema.GroupResource
}

// Informer returns the SharedIndexInformer.
func (f *genericInformer) Informer() cache.SharedIndexInformer {
	return f.informer
}

// Lister returns the GenericLister.
func (f *genericInformer) Lister() cache.GenericLister {
	return cache.NewGenericLister(f.Informer().GetIndexer(), f.resource)
}

// ForResource gives generic access to a shared informer of the matching type
// TODO extend this to unknown resources with a client pool
func (f *sharedInformerFactory) ForResource(resource schema.GroupVersionResource) (GenericInformer, error) {
	switch resource {
	// Group=garden.sapcloud.io, Version=internalVersion
	case garden.SchemeGroupVersion.WithResource("cloudprofiles"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Garden().InternalVersion().CloudProfiles().Informer()}, nil
	case garden.SchemeGroupVersion.WithResource("crosssecretbindings"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Garden().InternalVersion().CrossSecretBindings().Informer()}, nil
	case garden.SchemeGroupVersion.WithResource("privatesecretbindings"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Garden().InternalVersion().PrivateSecretBindings().Informer()}, nil
	case garden.SchemeGroupVersion.WithResource("quotas"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Garden().InternalVersion().Quotas().Informer()}, nil
	case garden.SchemeGroupVersion.WithResource("seeds"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Garden().InternalVersion().Seeds().Informer()}, nil
	case garden.SchemeGroupVersion.WithResource("shoots"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Garden().InternalVersion().Shoots().Informer()}, nil

	}

	return nil, fmt.Errorf("no informer found for %v", resource)
}
