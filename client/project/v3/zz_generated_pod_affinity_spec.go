package client

const (
	PodAffinitySpecType           = "podAffinitySpec"
	PodAffinitySpecFieldPreferred = "preferred"
	PodAffinitySpecFieldRequired  = "required"
)

type PodAffinitySpec struct {
	Preferred []PodAffinityRule `json:"preferred,omitempty" yaml:"preferred,omitempty"`
	Required  []PodAffinityRule `json:"required,omitempty" yaml:"required,omitempty"`
}
