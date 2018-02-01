package client

const (
	StepType                        = "step"
	StepFieldPublishImageStepConfig = "publishImageStepConfig"
	StepFieldRunScriptStepConfig    = "runScriptStepConfig"
	StepFieldSourceCodeStepConfig   = "sourceCodeStepConfig"
	StepFieldTimeout                = "timeout"
)

type Step struct {
	PublishImageStepConfig *PublishImageStepConfig `json:"publishImageStepConfig,omitempty"`
	RunScriptStepConfig    *RunScriptStepConfig    `json:"runScriptStepConfig,omitempty"`
	SourceCodeStepConfig   *SourceCodeStepConfig   `json:"sourceCodeStepConfig,omitempty"`
	Timeout                *int64                  `json:"timeout,omitempty"`
}
