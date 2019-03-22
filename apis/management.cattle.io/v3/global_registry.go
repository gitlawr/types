package v3

import (
	"github.com/rancher/norman/condition"
	"github.com/rancher/norman/types"
	"k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type GlobalRegistryConditionType string

const (
	GlobalRegistryConditionProvisioned condition.Cond = "Provisioned"
	GlobalRegistryConditionReady       condition.Cond = "Ready"
)

type GlobalRegistry struct {
	types.Namespaced

	metav1.TypeMeta `json:",inline"`
	// Standard object’s metadata. More info:
	// https://github.com/kubernetes/community/blob/master/contributors/devel/api-conventions.md#metadata
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   GlobalRegistrySpec   `json:"spec,omitempty"`
	Status GlobalRegistryStatus `json:"status,omitempty"`
}

type GlobalRegistrySpec struct {
	Answers map[string]string `json:"answers,omitempty" norman:"required"`
}

type GlobalRegistryStatus struct {
	Conditions []ClusterCondition `json:"conditions,omitempty"`
}

type GlobalRegistryCondition struct {
	// Type of registry condition.
	Type GlobalRegistryConditionType `json:"type"`
	// Status of the condition, one of True, False, Unknown.
	Status v1.ConditionStatus `json:"status"`
	// The last time this condition was updated.
	LastUpdateTime string `json:"lastUpdateTime,omitempty"`
	// Last time the condition transitioned from one status to another.
	LastTransitionTime string `json:"lastTransitionTime,omitempty"`
	// The reason for the condition's last transition.
	Reason string `json:"reason,omitempty"`
	// Human-readable message indicating details about last transition
	Message string `json:"message,omitempty"`
}
