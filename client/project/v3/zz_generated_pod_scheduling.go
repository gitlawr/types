package client

const (
	PodSchedulingType              = "podScheduling"
	PodSchedulingFieldAffinity     = "affinity"
	PodSchedulingFieldAntiAffinity = "antiAffinity"
)

type PodScheduling struct {
	Affinity     *PodAffinitySpec `json:"affinity,omitempty" yaml:"affinity,omitempty"`
	AntiAffinity *PodAffinitySpec `json:"antiAffinity,omitempty" yaml:"antiAffinity,omitempty"`
}
