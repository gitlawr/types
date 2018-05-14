package client

const (
	PublishImageConfigType                = "publishImageConfig"
	PublishImageConfigFieldBuildContext   = "buildContext"
	PublishImageConfigFieldDockerfilePath = "dockerfilePath"
	PublishImageConfigFieldPush           = "push"
	PublishImageConfigFieldTag            = "tag"
)

type PublishImageConfig struct {
	BuildContext   string `json:"buildContext,omitempty" yaml:"buildContext,omitempty"`
	DockerfilePath string `json:"dockerfilePath,omitempty" yaml:"dockerfilePath,omitempty"`
	Push           string `json:"push,omitempty" yaml:"push,omitempty"`
	Tag            string `json:"tag,omitempty" yaml:"tag,omitempty"`
}
