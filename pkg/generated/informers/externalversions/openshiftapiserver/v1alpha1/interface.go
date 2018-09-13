// Code generated by informer-gen. DO NOT EDIT.

package v1alpha1

import (
	internalinterfaces "github.com/openshift/cluster-openshift-apiserver-operator/pkg/generated/informers/externalversions/internalinterfaces"
)

// Interface provides access to all the informers in this group version.
type Interface interface {
	// OpenShiftAPIServerOperatorConfigs returns a OpenShiftAPIServerOperatorConfigInformer.
	OpenShiftAPIServerOperatorConfigs() OpenShiftAPIServerOperatorConfigInformer
}

type version struct {
	factory          internalinterfaces.SharedInformerFactory
	namespace        string
	tweakListOptions internalinterfaces.TweakListOptionsFunc
}

// New returns a new Interface.
func New(f internalinterfaces.SharedInformerFactory, namespace string, tweakListOptions internalinterfaces.TweakListOptionsFunc) Interface {
	return &version{factory: f, namespace: namespace, tweakListOptions: tweakListOptions}
}

// OpenShiftAPIServerOperatorConfigs returns a OpenShiftAPIServerOperatorConfigInformer.
func (v *version) OpenShiftAPIServerOperatorConfigs() OpenShiftAPIServerOperatorConfigInformer {
	return &openShiftAPIServerOperatorConfigInformer{factory: v.factory, tweakListOptions: v.tweakListOptions}
}
