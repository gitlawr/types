package client

const (
	AuthAppInputType              = "authAppInput"
	AuthAppInputFieldClientId     = "clientId"
	AuthAppInputFieldClientSecret = "clientSecret"
	AuthAppInputFieldCode         = "code"
	AuthAppInputFieldHost         = "host"
	AuthAppInputFieldRedirectUrl  = "type"
	AuthAppInputFieldTLS          = "tls"
)

type AuthAppInput struct {
	ClientId     string `json:"clientId,omitempty"`
	ClientSecret string `json:"clientSecret,omitempty"`
	Code         string `json:"code,omitempty"`
	Host         string `json:"host,omitempty"`
	RedirectUrl  string `json:"type,omitempty"`
	TLS          *bool  `json:"tls,omitempty"`
}
