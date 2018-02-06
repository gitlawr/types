package client

const (
	AuthUserInputType                = "authUserInput"
	AuthUserInputFieldCode           = "code"
	AuthUserInputFieldSourceCodeType = "sourceCodeType"
)

type AuthUserInput struct {
	Code           string `json:"code,omitempty"`
	SourceCodeType string `json:"sourceCodeType,omitempty"`
}
