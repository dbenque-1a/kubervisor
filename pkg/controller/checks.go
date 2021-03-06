package controller

import (
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/labels"

	api "github.com/amadeusitgroup/kubervisor/pkg/api/kubervisor/v1alpha1"
	"github.com/amadeusitgroup/kubervisor/pkg/controller/item"
)

// IsSpecUpdated return true if the the KubervisorService have been updated
func IsSpecUpdated(bc *api.KubervisorService, svc *corev1.Service, bci item.Interface) bool {
	selector := labels.Set(svc.Spec.Selector).AsSelectorPreValidated()
	return bci.CompareWithSpec(&bc.Spec, selector)
}
