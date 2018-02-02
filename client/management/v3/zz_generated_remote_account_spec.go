package client

const (
	RemoteAccountSpecType             = "remoteAccountSpec"
	RemoteAccountSpecFieldAccessToken = "accessToken"
	RemoteAccountSpecFieldAccountName = "accountName"
	RemoteAccountSpecFieldAvatarURL   = "avatarUrl"
	RemoteAccountSpecFieldDisplayName = "displayName"
	RemoteAccountSpecFieldHTMLURL     = "htmlUrl"
	RemoteAccountSpecFieldUserID      = "userId"
)

type RemoteAccountSpec struct {
	AccessToken string    `json:"accessToken,omitempty"`
	AccountName string    `json:"accountName,omitempty"`
	AvatarURL   string    `json:"avatarUrl,omitempty"`
	DisplayName string    `json:"displayName,omitempty"`
	HTMLURL     string    `json:"htmlUrl,omitempty"`
	UserID      *Required `json:"userId,omitempty"`
}
