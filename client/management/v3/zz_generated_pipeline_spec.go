package client

const (
	PipelineSpecType                = "pipelineSpec"
	PipelineSpecFieldActive         = "active"
	PipelineSpecFieldProjectId      = "projectId"
	PipelineSpecFieldStages         = "stages"
	PipelineSpecFieldTriggerCron    = "triggerCron"
	PipelineSpecFieldTriggerWebhook = "triggerWebhook"
)

type PipelineSpec struct {
	Active         *bool        `json:"active,omitempty"`
	ProjectId      string       `json:"projectId,omitempty"`
	Stages         []Stage      `json:"stages,omitempty"`
	TriggerCron    *TriggerCron `json:"triggerCron,omitempty"`
	TriggerWebhook *bool        `json:"triggerWebhook,omitempty"`
}
