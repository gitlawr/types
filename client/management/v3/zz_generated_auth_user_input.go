package client

const (
	AuthUserInputType      = "authUserInput"
	AuthUserInputFieldCode = "code"
	AuthUserInputFieldType = "type"
)

type AuthUserInput struct {
	Code string `json:"code,omitempty"`
	Type string `json:"type,omitempty"`
}
