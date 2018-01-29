package client

const (
	ClusterPipelineSpecType              = "clusterPipelineSpec"
	ClusterPipelineSpecFieldClusterId    = "clusterId"
	ClusterPipelineSpecFieldGibhubConfig = "githubConfig"
)

type ClusterPipelineSpec struct {
	ClusterId    string        `json:"clusterId,omitempty"`
	GibhubConfig *GibhubConfig `json:"githubConfig,omitempty"`
}
