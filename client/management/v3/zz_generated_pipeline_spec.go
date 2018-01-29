package client

const (
	PipelineSpecType               = "pipelineSpec"
	PipelineSpecFieldCronTrigger   = "cronTrigger"
	PipelineSpecFieldDisplayName   = "displayName"
	PipelineSpecFieldEnableTrigger = "enableTrigger"
	PipelineSpecFieldStages        = "stages"
)

type PipelineSpec struct {
	CronTrigger   *CronTrigger `json:"cronTrigger,omitempty"`
	DisplayName   string       `json:"displayName,omitempty"`
	EnableTrigger *bool        `json:"enableTrigger,omitempty"`
	Stages        []Stage      `json:"stages,omitempty"`
}
