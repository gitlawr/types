package client

const (
	ApplyYamlConfigType           = "applyYamlConfig"
	ApplyYamlConfigFieldContent   = "content"
	ApplyYamlConfigFieldEnv       = "env"
	ApplyYamlConfigFieldEnvFrom   = "envFrom"
	ApplyYamlConfigFieldNamespace = "namespace"
	ApplyYamlConfigFieldPath      = "path"
)

type ApplyYamlConfig struct {
	Content   string            `json:"content,omitempty" yaml:"content,omitempty"`
	Env       map[string]string `json:"env,omitempty" yaml:"env,omitempty"`
	EnvFrom   []EnvFrom         `json:"envFrom,omitempty" yaml:"envFrom,omitempty"`
	Namespace string            `json:"namespace,omitempty" yaml:"namespace,omitempty"`
	Path      string            `json:"path,omitempty" yaml:"path,omitempty"`
}
