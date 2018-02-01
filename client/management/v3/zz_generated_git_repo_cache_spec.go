package client

const (
	GitRepoCacheSpecType                   = "gitRepoCacheSpec"
	GitRepoCacheSpecFieldRemoteAccountName = "remoteAccountName"
	GitRepoCacheSpecFieldRepositories      = "repositories"
	GitRepoCacheSpecFieldType              = "type"
	GitRepoCacheSpecFieldUserID            = "userId"
)

type GitRepoCacheSpec struct {
	RemoteAccountName string          `json:"remoteAccountName,omitempty"`
	Repositories      []GitRepository `json:"repositories,omitempty"`
	Type              string          `json:"type,omitempty"`
	UserID            string          `json:"userId,omitempty"`
}
