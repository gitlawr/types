package client

const (
	SearchPipelineInputType                = "searchPipelineInput"
	SearchPipelineInputFieldBranch         = "branch"
	SearchPipelineInputFieldSourceCodeType = "sourceCodeType"
	SearchPipelineInputFieldURL            = "url"
)

type SearchPipelineInput struct {
	Branch         string `json:"branch,omitempty" yaml:"branch,omitempty"`
	SourceCodeType string `json:"sourceCodeType,omitempty" yaml:"sourceCodeType,omitempty"`
	URL            string `json:"url,omitempty" yaml:"url,omitempty"`
}
