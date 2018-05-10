package client

const (
	ClusterPipelineSpecType                  = "clusterPipelineSpec"
	ClusterPipelineSpecFieldClusterId        = "clusterId"
	ClusterPipelineSpecFieldDeploy           = "deploy"
	ClusterPipelineSpecFieldGeneralGitConfig = "generalGitConfig"
	ClusterPipelineSpecFieldGithubConfig     = "githubConfig"
	ClusterPipelineSpecFieldGitlabConfig     = "gitlabConfig"
)

type ClusterPipelineSpec struct {
	ClusterId        string            `json:"clusterId,omitempty" yaml:"clusterId,omitempty"`
	Deploy           bool              `json:"deploy,omitempty" yaml:"deploy,omitempty"`
	GeneralGitConfig *GeneralGitConfig `json:"generalGitConfig,omitempty" yaml:"generalGitConfig,omitempty"`
	GithubConfig     *GitAppConfig     `json:"githubConfig,omitempty" yaml:"githubConfig,omitempty"`
	GitlabConfig     *GitAppConfig     `json:"gitlabConfig,omitempty" yaml:"gitlabConfig,omitempty"`
}
