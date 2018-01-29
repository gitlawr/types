package v3

import (
	"github.com/rancher/norman/condition"
	"github.com/rancher/norman/types"
	"k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type ClusterPipeline struct {
	types.Namespaced

	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ClusterPipelineSpec   `json:"spec"`
	Status ClusterPipelineStatus `json:"status"`
}

type ClusterPipelineSpec struct {
	ClusterName  string        `json:"clusterName" norman:"type=reference[cluster]"`
	GibhubConfig *GibhubConfig `json:"githubConfig,omitempty"`
}

type ClusterPipelineStatus struct {
	Conditions []PipelineCondition `json:"conditions,omitempty"`
}

type GibhubConfig struct {
	Scheme       string `json:"githubConfig,omitempty"`
	Host         string `json:"host,omitempty"`
	ClientId     string `json:"clientId,omitempty"`
	ClientSecret string `json:"clientSecret,omitempty"`
}

var (
	ClusterPipelineConditionInitialized condition.Cond = "Initialized"
	ClusterPipelineConditionProvisioned condition.Cond = "Provisioned"
)

type PipelineCondition struct {
	// Type of cluster condition.
	Type condition.Cond `json:"type"`
	// Status of the condition, one of True, False, Unknown.
	Status v1.ConditionStatus `json:"status"`
	// The last time this condition was updated.
	LastUpdateTime string `json:"lastUpdateTime,omitempty"`
	// Last time the condition transitioned from one status to another.
	LastTransitionTime string `json:"lastTransitionTime,omitempty"`
	// The reason for the condition's last transition.
	Reason string `json:"reason,omitempty"`
	// Human-readable message indicating details about last transition
	Message string `json:"message,omitempty"`
}

type Pipeline struct {
	types.Namespaced

	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   PipelineSpec   `json:"spec"`
	Status PipelineStatus `json:"status"`
}

type PipelineStatus struct {
	NextRunNumber int    `json:"nextRunNumber" yaml:"nextRunNumber,omitempty"`
	LastRunId     string `json:"lastRunId,omitempty" yaml:"lastRunId,omitempty"`
	LastRunStatus string `json:"lastRunStatus,omitempty" yaml:"lastRunStatus,omitempty"`
	LastRunTime   int64  `json:"lastRunTime,omitempty" yaml:"lastRunTime,omitempty"`
	NextRunTime   int64  `json:"nextRunTime,omitempty" yaml:"nextRunTime,omitempty"`
	WebHookId     string `json:"webhookId,omitempty" yaml:"webhookId,omitempty"`
	WebHookToken  string `json:"webhookToken,omitempty" yaml:"webhookToken,omitempty"`
}

type PipelineSpec struct {
	EnableTrigger bool        `json:"enableTrigger,omitempty" norman:"default=true"`
	DisplayName   string      `json:"displayName,omitempty" norman:"required"`
	CronTrigger   CronTrigger `json:"cronTrigger,omitempty" yaml:"cronTrigger,omitempty"`
	Stages        []Stage     `json:"stages,omitempty" yaml:"stages,omitempty" norman:"required"`
}

type CronTrigger struct {
	Timezone        string `json:"timezone,omitempty" yaml:"timezone,omitempty"`
	Spec            string `json:"spec,omitempty" yaml:"spec,omitempty"`
	TriggerOnUpdate bool   `json:"triggerOnUpdate" yaml:"triggerOnUpdate,omitempty"`
}

type Stage struct {
	Name  string `json:"name,omitempty" yaml:"name,omitempty" norman:"required"`
	Steps []Step `json:"steps,omitempty" yaml:"steps,omitempty"`
}

type Step struct {
	Type string `json:"type,omitempty" yaml:"type,omitempty" norman:"required,options=runscript|buildimage,default=runscript"`

	SourceCodeStepConfig *SourceCodeStepConfig
	RunScriptStepConfig  *RunScriptStepConfig
	BuildImageStepConfig *BuildImageStepConfig

	//Step timeout in minutes
	Timeout int `json:"timeout,omitempty" yaml:"timeout,omitempty"`
}

type SourceCodeStepConfig struct {
	Repository string `json:"repository,omitempty" yaml:"repository,omitempty" norman:"required"`
	Branch     string `json:"branch,omitempty" yaml:"branch,omitempty" norman:"required"`
	RemoteUser string `json:"remoteUser,omitempty" yaml:"remoteUser,omitempty" norman:"required,type=reference[remoteuser]"`
}

type RunScriptStepConfig struct {
	Image       string   `json:"image,omitempty" yaml:"image,omitempty"`
	ShellScript string   `json:"shellScript,omitempty" yaml:"shellScript,omitempty"`
	Entrypoint  string   `json:"entrypoint,omitempty" yaml:"enrtypoint,omitempty"`
	Args        string   `json:"args,omitempty" yaml:"args,omitempty"`
	Env         []string `json:"env,omitempty" yaml:"env,omitempty"`
}

type BuildImageStepConfig struct {
	DockerfilePath string `json:"dockerFilePath,omittempty" yaml:"dockerFilePath,omitempty"`
	BuildPath      string `json:"buildPath,omitempty" yaml:"buildPath,omitempty"`
	TargetImage    string `json:"targetImage,omitempty" yaml:"targetImage,omitempty"`
	Push           bool   `json:"push" yaml:"push,omitempty"`
}

type PipelineHistory struct {
	types.Namespaced

	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   PipelineHistorySpec   `json:"spec"`
	Status PipelineHistoryStatus `json:"status"`
}

type PipelineHistorySpec struct {
	DisplayName string   `json:"displayName,omitempty" norman:"required"`
	RunNumber   int      `json:"runNumber,omitempty" norman:"required,min=1"`
	TriggerType string   `json:"triggerType,omitempty" norman:"required,options=manual|cron|webhook"`
	Pipeline    Pipeline `json:"pipeline,omitempty" norman:"required"`
}

type PipelineHistoryStatus struct {
	CommitInfo  string            `json:"commitInfo,omitempty"`
	EnvVars     map[string]string `json:"envVars,omitempty"`
	Status      string            `json:"status,omitempty"`
	StartTime   int64             `json:"startTime,omitempty"`
	EndTime     int64             `json:"endTime,omitempty"`
	StageStatus []StageStatus     `json:"stageStatus,omitempty"`
}

type StageStatus struct {
	Status     string       `json:"status,omitempty"`
	StartTime  int64        `json:"startTime,omitempty"`
	EndTime    int64        `json:"endTime,omitempty"`
	StepStatus []StepStatus `json:"stepStatus,omitempty"`
}

type StepStatus struct {
	Status    string `json:"status,omitempty"`
	StartTime int64  `json:"startTime,omitempty"`
	EndTime   int64  `json:"endTime,omitempty"`
}

type RemoteAccount struct {
	types.Namespaced

	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   RemoteAccountSpec   `json:"spec"`
	Status RemoteAccountStatus `json:"status"`
}

type RemoteAccountSpec struct {
	DisplayName string `json:"displayName,omitempty" norman:"required"`
	Type        string `json:"type,omitempty" norman:"required,options=github"`
	UserName    string `json:"userName,omitempty" norman:"required,type=reference[user]"`
	AvatarURL   string `json:"avatarUrl,omitempty"`
	HTMLURL     string `json:"htmlUrl,omitempty"`
	AccountName string `json:"accountId,omitempty"`
	AccessToken string `json:"accessToken,omitempty"`
}

type RemoteAccountStatus struct {
}

type GitRepoCache struct {
	types.Namespaced

	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   GitRepoCacheSpec   `json:"spec"`
	Status GitRepoCacheStatus `json:"status"`
}

type GitRepoCacheSpec struct {
	Type              string          `json:"type,omitempty" norman:"required,options=github"`
	RemoteAccountName string          `json:"remoteAccountName,omitempty" norman:"required,type=reference[remoteaccount]`
	Repositories      []GitRepository `json:"repositories,omitempty"`
}

type GitRepository struct {
	Name        string   `json:"name,omitempty"`
	CloneURL    string   `json:"cloneUrl,omitempty"`
	Permissions RepoPerm `json:"permissions,omitempty"`
	Language    string   `json:"language,omitempty"`
}

type GitRepoCacheStatus struct {
}

type RepoPerm struct {
	Pull  bool `json:"pull,omitempty"`
	Push  bool `json:"push,omitempty"`
	Admin bool `json:"admin,omitempty"`
}
