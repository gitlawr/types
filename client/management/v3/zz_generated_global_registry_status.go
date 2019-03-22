package client

const (
	GlobalRegistryStatusType            = "globalRegistryStatus"
	GlobalRegistryStatusFieldConditions = "conditions"
)

type GlobalRegistryStatus struct {
	Conditions []ClusterCondition `json:"conditions,omitempty" yaml:"conditions,omitempty"`
}
