package client

const (
	PipelineExecutionStatusType             = "pipelineExecutionStatus"
	PipelineExecutionStatusFieldCommitInfo  = "commitInfo"
	PipelineExecutionStatusFieldEndTime     = "endTime"
	PipelineExecutionStatusFieldEnvVars     = "envVars"
	PipelineExecutionStatusFieldStageStatus = "stageStatus"
	PipelineExecutionStatusFieldStartTime   = "startTime"
	PipelineExecutionStatusFieldState       = "state"
)

type PipelineExecutionStatus struct {
	CommitInfo  string            `json:"commitInfo,omitempty"`
	EndTime     *int64            `json:"endTime,omitempty"`
	EnvVars     map[string]string `json:"envVars,omitempty"`
	StageStatus []StageStatus     `json:"stageStatus,omitempty"`
	StartTime   *int64            `json:"startTime,omitempty"`
	State       string            `json:"state,omitempty"`
}
