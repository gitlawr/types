package client

const (
	SourceCodeStepConfigType            = "sourceCodeStepConfig"
	SourceCodeStepConfigFieldBranch     = "branch"
	SourceCodeStepConfigFieldRemoteUser = "remoteUser"
	SourceCodeStepConfigFieldRepository = "repository"
)

type SourceCodeStepConfig struct {
	Branch     string `json:"branch,omitempty"`
	RemoteUser string `json:"remoteUser,omitempty"`
	Repository string `json:"repository,omitempty"`
}
