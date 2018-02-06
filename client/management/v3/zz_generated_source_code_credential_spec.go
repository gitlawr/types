package client

const (
	SourceCodeCredentialSpecType                = "sourceCodeCredentialSpec"
	SourceCodeCredentialSpecFieldAccessToken    = "accessToken"
	SourceCodeCredentialSpecFieldAvatarURL      = "avatarUrl"
	SourceCodeCredentialSpecFieldClusterId      = "clusterId"
	SourceCodeCredentialSpecFieldDisplayName    = "displayName"
	SourceCodeCredentialSpecFieldHTMLURL        = "htmlUrl"
	SourceCodeCredentialSpecFieldLoginName      = "loginName"
	SourceCodeCredentialSpecFieldRancherUserId  = "rancherUserId"
	SourceCodeCredentialSpecFieldSourceCodeType = "sourceCodeType"
)

type SourceCodeCredentialSpec struct {
	AccessToken    string `json:"accessToken,omitempty"`
	AvatarURL      string `json:"avatarUrl,omitempty"`
	ClusterId      string `json:"clusterId,omitempty"`
	DisplayName    string `json:"displayName,omitempty"`
	HTMLURL        string `json:"htmlUrl,omitempty"`
	LoginName      string `json:"loginName,omitempty"`
	RancherUserId  string `json:"rancherUserId,omitempty"`
	SourceCodeType string `json:"sourceCodeType,omitempty"`
}
