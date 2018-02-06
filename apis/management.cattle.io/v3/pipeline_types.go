package v3

import (
	"github.com/rancher/norman/condition"
	"github.com/rancher/norman/types"
	"k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

const (
	StepTypeSourceCode = "sourceCode"
	StepTypeRunScript  = "runScript"
	StepTypeBuildImage = "buildImage"
	TriggerTypeCron    = "cron"
	TriggerTypeManual  = "manual"
	TriggerTypeWebhook = "webhook"

	StateWaiting  = "Waiting"
	StateBuilding = "Building"
	StateSuccess  = "Success"
	StateFail     = "Fail"
	StateSkip     = "Skipped"
	StateAbort    = "Abort"
	StatePending  = "Pending"
	StateDenied   = "Denied"
)

type ClusterPipeline struct {
	types.Namespaced

	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ClusterPipelineSpec   `json:"spec"`
	Status ClusterPipelineStatus `json:"status"`
}

type Pipeline struct {
	types.Namespaced

	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   PipelineSpec   `json:"spec"`
	Status PipelineStatus `json:"status" yaml:"-"`
}

type PipelineExecution struct {
	types.Namespaced

	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   PipelineExecutionSpec   `json:"spec"`
	Status PipelineExecutionStatus `json:"status"`
}

type PipelineLog struct {
	types.Namespaced

	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec PipelineLogSpec `json:"spec"`
}

type SourceCodeCredential struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   SourceCodeCredentialSpec   `json:"spec"`
	Status SourceCodeCredentialStatus `json:"status"`
}

type SourceCodeRepository struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   SourceCodeRepositorySpec   `json:"spec"`
	Status SourceCodeRepositoryStatus `json:"status"`
}

type ClusterPipelineSpec struct {
	ClusterName  string        `json:"clusterName" norman:"type=reference[cluster]"`
	Deploy       bool          `json:"deploy"`
	GithubConfig *GithubConfig `json:"githubConfig,omitempty"`
}

type ClusterPipelineStatus struct {
	Conditions []PipelineCondition `json:"conditions,omitempty"`
}

type GithubConfig struct {
	tls          bool   `json:"tls,omitempty"`
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

type PipelineStatus struct {
	NextRunNumber int    `json:"nextRunNumber" yaml:"nextRunNumber,omitempty" norman:"default=1,min=1"`
	LastRunId     string `json:"lastRunId,omitempty" yaml:"lastRunId,omitempty"`
	LastRunState  string `json:"lastRunState,omitempty" yaml:"lastRunState,omitempty"`
	LastRunTime   int64  `json:"lastRunTime,omitempty" yaml:"lastRunTime,omitempty"`
	NextRunTime   int64  `json:"nextRunTime,omitempty" yaml:"nextRunTime,omitempty"`
	WebHookId     string `json:"webhookId,omitempty" yaml:"webhookId,omitempty"`
	WebHookToken  string `json:"webhookToken,omitempty" yaml:"webhookToken,omitempty"`
}

type PipelineSpec struct {
	ProjectName string `json:"projectName" yaml:"projectName" norman:"required,type=reference[project]"`

	Active bool `json:"active,omitempty" yaml:"active,omitempty"`
	//DisplayName string   `json:"displayName,omitempty" yaml:"displayName,omitempty" norman:"required"`
	WebhookTrigger *TriggerWebhook `json:"triggerWebhook,omitempty" yaml:"triggerWebhook,omitempty"`
	TriggerCron    *TriggerCron    `json:"triggerCron,omitempty" yaml:"triggerCron,omitempty"`

	Stages []Stage `json:"stages,omitempty" yaml:"stages,omitempty" norman:"required"`
}

type Triggers struct {
	WebhookTrigger *TriggerWebhook `json:"triggerWebhook,omitempty" yaml:"triggerWebhook,omitempty"`
	CronTrigger    *TriggerCron    `json:"triggerCron,omitempty" yaml:"triggerCron,omitempty"`
}

type TriggerWebhook struct {
	Active bool `json:"active,omitempty" yaml:"active,omitempty"`
}

type TriggerCron struct {
	Active   bool   `json:"active,omitempty" yaml:"active,omitempty"`
	Timezone string `json:"timezone,omitempty" yaml:"timezone,omitempty"`
	Spec     string `json:"spec,omitempty" yaml:"spec,omitempty"`
}

type Stage struct {
	Name  string `json:"name,omitempty" yaml:"name,omitempty" norman:"required"`
	Steps []Step `json:"steps,omitempty" yaml:"steps,omitempty"`
}

type Step struct {
	Type string `json:"type,omitempty" yaml:"type,omitempty" norman:"required,options=sourcecode|runscript|buildimage|pushimage,default=runscript"`

	SourceCodeStepConfig   *SourceCodeStepConfig   `json:"sourceCodeConfig,omitempty" yaml:"sourceCodeConfig,omitempty"`
	RunScriptStepConfig    *RunScriptStepConfig    `json:"runScriptConfig,omitempty" yaml:"runScriptConfig,omitempty"`
	PublishImageStepConfig *PublishImageStepConfig `json:"publishImageConfig,omitempty" yaml:"publishImageConfig,omitempty"`
	//Step timeout in minutes
	Timeout int `json:"timeout,omitempty" yaml:"timeout,omitempty"`
}

type SourceCodeStepConfig struct {
	Repository               string `json:"repository,omitempty" yaml:"repository,omitempty" `
	Branch                   string `json:"branch,omitempty" yaml:"branch,omitempty" `
	BranchCondition          string `json:"branchCondition,omitempty" yaml:"branchCondition,omitempty" norman:"options=only|except|all,default=only"`
	SourceCodeCredentialName string `json:"sourceCodeCredentialName,omitempty" yaml:"sourceCodeCredentialName,omitempty" norman:"reference[sourcecodecredential]"`
}

type RunScriptStepConfig struct {
	Image       string   `json:"image,omitempty" yaml:"image,omitempty" norman:"required"`
	ShellScript string   `json:"shellScript,omitempty" yaml:"shellScript,omitempty"`
	Entrypoint  string   `json:"entrypoint,omitempty" yaml:"enrtypoint,omitempty"`
	Command     string   `json:"command,omitempty" yaml:"command,omitempty"`
	Env         []string `json:"env,omitempty" yaml:"env,omitempty"`
}

type PublishImageStepConfig struct {
	DockerfilePath string `json:"dockerfilePath,omittempty" yaml:"dockerfilePath,omitempty" norman:"required,default=./Dockerfile"`
	BuildContext   string `json:"buildContext,omitempty" yaml:"buildContext,omitempty" norman:"required,default=."`
	Tag            string `json:"tag,omitempty" yaml:"tag,omitempty" norman:"required,default=${CICD_GIT_REPOSITORY_NAME}:${CICD_GIT_BRANCH}"`
}

type PipelineExecutionSpec struct {
	ProjectName string `json:"projectName" norman:"required,type=reference[project]"`

	RunNumber       int      `json:"runNumber,omitempty" norman:"required,min=1"`
	TriggeredBy     string   `json:"triggeredBy,omitempty" norman:"required,options=user|cron|webhook"`
	TriggerUserName string   `json:"triggerUserName,omitempty" norman:"type=reference[user]"`
	Pipeline        Pipeline `json:"pipeline,omitempty" norman:"required"`
}

type PipelineExecutionStatus struct {
	CommitInfo  string            `json:"commitInfo,omitempty"`
	EnvVars     map[string]string `json:"envVars,omitempty"`
	State       string            `json:"state,omitempty"`
	StartTime   int64             `json:"startTime,omitempty"`
	EndTime     int64             `json:"endTime,omitempty"`
	StageStatus []StageStatus     `json:"stageStatus,omitempty"`
}

type StageStatus struct {
	State      string       `json:"state,omitempty"`
	StartTime  int64        `json:"startTime,omitempty"`
	EndTime    int64        `json:"endTime,omitempty"`
	StepStatus []StepStatus `json:"stepStatus,omitempty"`
}

type StepStatus struct {
	State     string `json:"state,omitempty"`
	StartTime int64  `json:"startTime,omitempty"`
	EndTime   int64  `json:"endTime,omitempty"`
}

type SourceCodeCredentialSpec struct {
	//DisplayName string `json:"displayName,omitempty" norman:"required"`
	SourceCodeType string `json:"sourceCodeType,omitempty" norman:"required,options=github"`
	UserName       string `json:"userName" norman:"required,type=reference[user]"`
	AvatarURL      string `json:"avatarUrl,omitempty"`
	HTMLURL        string `json:"htmlUrl,omitempty"`
	Login          string `json:"login,omitempty"`
	AccountName    string `json:"accountName,omitempty"`
	AccessToken    string `json:"accessToken,omitempty"`
}

type SourceCodeCredentialStatus struct {
}

type SourceCodeRepositorySpec struct {
	SourceCodeType           string   `json:"sourceCodeType,omitempty" norman:"required,options=github"`
	UserName                 string   `json:"userName" norman:"required,type=reference[user]"`
	SourceCodeCredentialName string   `json:"sourceCodeCredentialName,omitempty" norman:"required,type=reference[sourcecodecredential]`
	Name                     string   `json:"name,omitempty"`
	CloneURL                 string   `json:"cloneUrl,omitempty"`
	Permissions              RepoPerm `json:"permissions,omitempty"`
	Language                 string   `json:"language,omitempty"`
}

type SourceCodeRepositoryStatus struct {
}

type RepoPerm struct {
	Pull  bool `json:"pull,omitempty"`
	Push  bool `json:"push,omitempty"`
	Admin bool `json:"admin,omitempty"`
}

type PipelineLogSpec struct {
	ProjectName string `json:"projectName" yaml:"projectName" norman:"required,type=reference[project]"`

	PipelineExecutionName string `json:"pipelineExecutionName,omitempty" norman:"type=reference[pipelineexecution]`
	Stage                 int    `json:"stage,omitempty" norman:"min=1"`
	Step                  int    `json:"step,omitempty" norman:"min=1"`
	Message               string `json:"message,omitempty"`
}
