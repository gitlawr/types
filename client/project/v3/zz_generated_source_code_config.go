package client

const (
	SourceCodeConfigType               = "sourceCodeConfig"
	SourceCodeConfigFieldSkipClone     = "skipClone"
	SourceCodeConfigFieldSkipTLSVerify = "skipTlsVerify"
)

type SourceCodeConfig struct {
	SkipClone     bool `json:"skipClone,omitempty" yaml:"skipClone,omitempty"`
	SkipTLSVerify bool `json:"skipTlsVerify,omitempty" yaml:"skipTlsVerify,omitempty"`
}
