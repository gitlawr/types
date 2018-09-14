package client

const (
	TargetPipelineType            = "targetPipeline"
	TargetPipelineFieldCondition  = "condition"
	TargetPipelineFieldPipelineID = "pipelineId"
)

type TargetPipeline struct {
	Condition  string `json:"condition,omitempty" yaml:"condition,omitempty"`
	PipelineID string `json:"pipelineId,omitempty" yaml:"pipelineId,omitempty"`
}
