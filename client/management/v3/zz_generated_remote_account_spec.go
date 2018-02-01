package client

const (
	RemoteAccountSpecType             = "remoteAccountSpec"
	RemoteAccountSpecFieldAccessToken = "accessToken"
	RemoteAccountSpecFieldAccountName = "accountName"
	RemoteAccountSpecFieldAvatarURL   = "avatarUrl"
	RemoteAccountSpecFieldDisplayName = "displayName"
	RemoteAccountSpecFieldHTMLURL     = "htmlUrl"
	RemoteAccountSpecFieldType        = "type"
	RemoteAccountSpecFieldUserID      = "userId"
)

type RemoteAccountSpec struct {
	AccessToken string `json:"accessToken,omitempty"`
	AccountName string `json:"accountName,omitempty"`
	AvatarURL   string `json:"avatarUrl,omitempty"`
	DisplayName string `json:"displayName,omitempty"`
	HTMLURL     string `json:"htmlUrl,omitempty"`
	Type        string `json:"type,omitempty"`
	UserID      string `json:"userId,omitempty"`
}
