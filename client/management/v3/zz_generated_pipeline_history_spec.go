package client

const (
	PipelineHistorySpecType             = "pipelineHistorySpec"
	PipelineHistorySpecFieldDisplayName = "displayName"
	PipelineHistorySpecFieldPipeline    = "pipeline"
	PipelineHistorySpecFieldRunNumber   = "runNumber"
	PipelineHistorySpecFieldTriggerType = "triggerType"
)

type PipelineHistorySpec struct {
	DisplayName string    `json:"displayName,omitempty"`
	Pipeline    *Pipeline `json:"pipeline,omitempty"`
	RunNumber   *int64    `json:"runNumber,omitempty"`
	TriggerType string    `json:"triggerType,omitempty"`
}
