package client

const (
	RunScriptConfigType             = "runScriptConfig"
	RunScriptConfigFieldCommand     = "command"
	RunScriptConfigFieldEntrypoint  = "entrypoint"
	RunScriptConfigFieldEnv         = "env"
	RunScriptConfigFieldEnvFrom     = "envFrom"
	RunScriptConfigFieldImage       = "image"
	RunScriptConfigFieldIsShell     = "isShell"
	RunScriptConfigFieldPrivileged  = "privileged"
	RunScriptConfigFieldShellScript = "shellScript"
)

type RunScriptConfig struct {
	Command     string            `json:"command,omitempty" yaml:"command,omitempty"`
	Entrypoint  string            `json:"entrypoint,omitempty" yaml:"entrypoint,omitempty"`
	Env         map[string]string `json:"env,omitempty" yaml:"env,omitempty"`
	EnvFrom     []EnvFrom         `json:"envFrom,omitempty" yaml:"envFrom,omitempty"`
	Image       string            `json:"image,omitempty" yaml:"image,omitempty"`
	IsShell     bool              `json:"isShell,omitempty" yaml:"isShell,omitempty"`
	Privileged  bool              `json:"privileged,omitempty" yaml:"privileged,omitempty"`
	ShellScript string            `json:"shellScript,omitempty" yaml:"shellScript,omitempty"`
}
