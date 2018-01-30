package client

const (
	PipelineSpecType             = "pipelineSpec"
	PipelineSpecFieldActive      = "active"
	PipelineSpecFieldCronTrigger = "cronTrigger"
	PipelineSpecFieldDisplayName = "displayName"
	PipelineSpecFieldProjectId   = "projectId"
	PipelineSpecFieldStages      = "stages"
)

type PipelineSpec struct {
	Active      *bool        `json:"active,omitempty"`
	CronTrigger *CronTrigger `json:"cronTrigger,omitempty"`
	DisplayName string       `json:"displayName,omitempty"`
	ProjectId   string       `json:"projectId,omitempty"`
	Stages      []Stage      `json:"stages,omitempty"`
}
