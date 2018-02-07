package client

const (
	AuthAppInputType              = "authAppInput"
	AuthAppInputFieldClientId     = "clientId"
	AuthAppInputFieldClientSecret = "clientSecret"
	AuthAppInputFieldCode         = "code"
	AuthAppInputFieldHost         = "host"
	AuthAppInputFieldTLS          = "tls"
	AuthAppInputFieldType         = "type"
)

type AuthAppInput struct {
	ClientId     string `json:"clientId,omitempty"`
	ClientSecret string `json:"clientSecret,omitempty"`
	Code         string `json:"code,omitempty"`
	Host         string `json:"host,omitempty"`
	TLS          *bool  `json:"tls,omitempty"`
	Type         string `json:"type,omitempty"`
}
