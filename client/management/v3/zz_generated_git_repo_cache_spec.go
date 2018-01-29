package client

const (
	GitRepoCacheSpecType                   = "gitRepoCacheSpec"
	GitRepoCacheSpecFieldRemoteAccountName = "remoteAccountName"
	GitRepoCacheSpecFieldRepositories      = "repositories"
	GitRepoCacheSpecFieldType              = "type"
)

type GitRepoCacheSpec struct {
	RemoteAccountName string          `json:"remoteAccountName,omitempty"`
	Repositories      []GitRepository `json:"repositories,omitempty"`
	Type              string          `json:"type,omitempty"`
}
