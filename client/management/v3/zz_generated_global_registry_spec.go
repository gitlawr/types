package client

const (
	GlobalRegistrySpecType         = "globalRegistrySpec"
	GlobalRegistrySpecFieldAnswers = "answers"
)

type GlobalRegistrySpec struct {
	Answers map[string]string `json:"answers,omitempty" yaml:"answers,omitempty"`
}
