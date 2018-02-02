package client

const (
	GitRepoCacheSpecType                   = "gitRepoCacheSpec"
	GitRepoCacheSpecFieldRemoteAccountName = "remoteAccountName"
	GitRepoCacheSpecFieldRemoteType        = "remoteType"
	GitRepoCacheSpecFieldRepositories      = "repositories"
	GitRepoCacheSpecFieldUserID            = "userId"
)

type GitRepoCacheSpec struct {
	RemoteAccountName string          `json:"remoteAccountName,omitempty"`
	RemoteType        string          `json:"remoteType,omitempty"`
	Repositories      []GitRepository `json:"repositories,omitempty"`
	UserID            *Required       `json:"userId,omitempty"`
}
