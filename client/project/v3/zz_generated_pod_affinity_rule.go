package client

const (
	PodAffinityRuleType             = "podAffinityRule"
	PodAffinityRuleFieldNamespaces  = "namespaces"
	PodAffinityRuleFieldRules       = "rules"
	PodAffinityRuleFieldTopologyKey = "topologyKey"
)

type PodAffinityRule struct {
	Namespaces  []string `json:"namespaces,omitempty" yaml:"namespaces,omitempty"`
	Rules       []string `json:"rules,omitempty" yaml:"rules,omitempty"`
	TopologyKey string   `json:"topologyKey,omitempty" yaml:"topologyKey,omitempty"`
}
