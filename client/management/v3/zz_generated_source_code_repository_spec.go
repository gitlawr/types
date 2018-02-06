package client

const (
	SourceCodeRepositorySpecType                          = "sourceCodeRepositorySpec"
	SourceCodeRepositorySpecFieldLanguage                 = "language"
	SourceCodeRepositorySpecFieldPermissions              = "permissions"
	SourceCodeRepositorySpecFieldSourceCodeCredentialName = "sourceCodeCredentialName"
	SourceCodeRepositorySpecFieldSourceCodeType           = "sourceCodeType"
	SourceCodeRepositorySpecFieldUrl                      = "url"
	SourceCodeRepositorySpecFieldUserId                   = "userId"
)

type SourceCodeRepositorySpec struct {
	Language                 string    `json:"language,omitempty"`
	Permissions              *RepoPerm `json:"permissions,omitempty"`
	SourceCodeCredentialName string    `json:"sourceCodeCredentialName,omitempty"`
	SourceCodeType           string    `json:"sourceCodeType,omitempty"`
	Url                      string    `json:"url,omitempty"`
	UserId                   string    `json:"userId,omitempty"`
}
