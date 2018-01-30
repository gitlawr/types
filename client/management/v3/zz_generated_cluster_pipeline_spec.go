package client

const (
	ClusterPipelineSpecType              = "clusterPipelineSpec"
	ClusterPipelineSpecFieldClusterId    = "clusterId"
	ClusterPipelineSpecFieldGithubConfig = "githubConfig"
)

type ClusterPipelineSpec struct {
	ClusterId    string        `json:"clusterId,omitempty"`
	GithubConfig *GithubConfig `json:"githubConfig,omitempty"`
}
