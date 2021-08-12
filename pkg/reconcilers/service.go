package reconcilers

import (
	"fmt"
	"reflect"

	"github.com/3scale/apicast-operator/pkg/k8sutils"
	v1 "k8s.io/api/core/v1"
)

func ServicePortMutator(existingObj, desiredObj k8sutils.KubernetesObject) (bool, error) {
	existing, ok := existingObj.(*v1.Service)
	if !ok {
		return false, fmt.Errorf("%T is not a *v1.Service", existingObj)
	}
	desired, ok := desiredObj.(*v1.Service)
	if !ok {
		return false, fmt.Errorf("%T is not a *v1.Service", desiredObj)
	}

	updated := false

	if !reflect.DeepEqual(existing.Spec.Ports, desired.Spec.Ports) {
		updated = true
		existing.Spec.Ports = desired.Spec.Ports
	}

	return updated, nil
}
