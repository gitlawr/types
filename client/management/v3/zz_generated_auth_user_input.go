package client

const (
	AuthUserInputType             = "authUserInput"
	AuthUserInputFieldCode        = "code"
	AuthUserInputFieldRedirectUrl = "type"
)

type AuthUserInput struct {
	Code        string `json:"code,omitempty"`
	RedirectUrl string `json:"type,omitempty"`
}
