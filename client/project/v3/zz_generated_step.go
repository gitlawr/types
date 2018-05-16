package client

const (
	StepType                    = "step"
	StepFieldApplyYamlConfig    = "applyYamlConfig"
	StepFieldPublishImageConfig = "publishImageConfig"
	StepFieldRunScriptConfig    = "runScriptConfig"
	StepFieldSourceCodeConfig   = "sourceCodeConfig"
	StepFieldTimeout            = "timeout"
	StepFieldWhen               = "when"
)

type Step struct {
	ApplyYamlConfig    *ApplyYamlConfig    `json:"applyYamlConfig,omitempty" yaml:"applyYamlConfig,omitempty"`
	PublishImageConfig *PublishImageConfig `json:"publishImageConfig,omitempty" yaml:"publishImageConfig,omitempty"`
	RunScriptConfig    *RunScriptConfig    `json:"runScriptConfig,omitempty" yaml:"runScriptConfig,omitempty"`
	SourceCodeConfig   *SourceCodeConfig   `json:"sourceCodeConfig,omitempty" yaml:"sourceCodeConfig,omitempty"`
	Timeout            int64               `json:"timeout,omitempty" yaml:"timeout,omitempty"`
	When               *Constraints        `json:"when,omitempty" yaml:"when,omitempty"`
}
