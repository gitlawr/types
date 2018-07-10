package v3

import (
	reflect "reflect"

	conversion "k8s.io/apimachinery/pkg/conversion"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

func init() {
	SchemeBuilder.Register(RegisterDeepCopies)
}

// RegisterDeepCopies adds deep-copy functions to the given scheme. Public
// to allow building arbitrary schemes.
//
// Deprecated: deepcopy registration will go away when static deepcopy is fully implemented.
func RegisterDeepCopies(scheme *runtime.Scheme) error {
	return scheme.AddGeneratedDeepCopyFuncs(
		conversion.GeneratedDeepCopyFunc{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*App).DeepCopyInto(out.(*App))
			return nil
		}, InType: reflect.TypeOf(&App{})},
		conversion.GeneratedDeepCopyFunc{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*AppCondition).DeepCopyInto(out.(*AppCondition))
			return nil
		}, InType: reflect.TypeOf(&AppCondition{})},
		conversion.GeneratedDeepCopyFunc{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*AppList).DeepCopyInto(out.(*AppList))
			return nil
		}, InType: reflect.TypeOf(&AppList{})},
		conversion.GeneratedDeepCopyFunc{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*AppRevision).DeepCopyInto(out.(*AppRevision))
			return nil
		}, InType: reflect.TypeOf(&AppRevision{})},
		conversion.GeneratedDeepCopyFunc{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*AppRevisionList).DeepCopyInto(out.(*AppRevisionList))
			return nil
		}, InType: reflect.TypeOf(&AppRevisionList{})},
		conversion.GeneratedDeepCopyFunc{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*AppRevisionSpec).DeepCopyInto(out.(*AppRevisionSpec))
			return nil
		}, InType: reflect.TypeOf(&AppRevisionSpec{})},
		conversion.GeneratedDeepCopyFunc{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*AppRevisionStatus).DeepCopyInto(out.(*AppRevisionStatus))
			return nil
		}, InType: reflect.TypeOf(&AppRevisionStatus{})},
		conversion.GeneratedDeepCopyFunc{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*AppSpec).DeepCopyInto(out.(*AppSpec))
			return nil
		}, InType: reflect.TypeOf(&AppSpec{})},
		conversion.GeneratedDeepCopyFunc{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*AppStatus).DeepCopyInto(out.(*AppStatus))
			return nil
		}, InType: reflect.TypeOf(&AppStatus{})},
		conversion.GeneratedDeepCopyFunc{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*AppUpgradeConfig).DeepCopyInto(out.(*AppUpgradeConfig))
			return nil
		}, InType: reflect.TypeOf(&AppUpgradeConfig{})},
		conversion.GeneratedDeepCopyFunc{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*ApplyYamlConfig).DeepCopyInto(out.(*ApplyYamlConfig))
			return nil
		}, InType: reflect.TypeOf(&ApplyYamlConfig{})},
		conversion.GeneratedDeepCopyFunc{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*AuthAppInput).DeepCopyInto(out.(*AuthAppInput))
			return nil
		}, InType: reflect.TypeOf(&AuthAppInput{})},
		conversion.GeneratedDeepCopyFunc{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*AuthUserInput).DeepCopyInto(out.(*AuthUserInput))
			return nil
		}, InType: reflect.TypeOf(&AuthUserInput{})},
		conversion.GeneratedDeepCopyFunc{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*BasicAuth).DeepCopyInto(out.(*BasicAuth))
			return nil
		}, InType: reflect.TypeOf(&BasicAuth{})},
		conversion.GeneratedDeepCopyFunc{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*BasicAuthList).DeepCopyInto(out.(*BasicAuthList))
			return nil
		}, InType: reflect.TypeOf(&BasicAuthList{})},
		conversion.GeneratedDeepCopyFunc{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*Certificate).DeepCopyInto(out.(*Certificate))
			return nil
		}, InType: reflect.TypeOf(&Certificate{})},
		conversion.GeneratedDeepCopyFunc{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*CertificateList).DeepCopyInto(out.(*CertificateList))
			return nil
		}, InType: reflect.TypeOf(&CertificateList{})},
		conversion.GeneratedDeepCopyFunc{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*Constraint).DeepCopyInto(out.(*Constraint))
			return nil
		}, InType: reflect.TypeOf(&Constraint{})},
		conversion.GeneratedDeepCopyFunc{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*Constraints).DeepCopyInto(out.(*Constraints))
			return nil
		}, InType: reflect.TypeOf(&Constraints{})},
		conversion.GeneratedDeepCopyFunc{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*DeploymentRollbackInput).DeepCopyInto(out.(*DeploymentRollbackInput))
			return nil
		}, InType: reflect.TypeOf(&DeploymentRollbackInput{})},
		conversion.GeneratedDeepCopyFunc{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*DockerCredential).DeepCopyInto(out.(*DockerCredential))
			return nil
		}, InType: reflect.TypeOf(&DockerCredential{})},
		conversion.GeneratedDeepCopyFunc{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*DockerCredentialList).DeepCopyInto(out.(*DockerCredentialList))
			return nil
		}, InType: reflect.TypeOf(&DockerCredentialList{})},
		conversion.GeneratedDeepCopyFunc{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*EnvFrom).DeepCopyInto(out.(*EnvFrom))
			return nil
		}, InType: reflect.TypeOf(&EnvFrom{})},
		conversion.GeneratedDeepCopyFunc{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*GithubLoginInput).DeepCopyInto(out.(*GithubLoginInput))
			return nil
		}, InType: reflect.TypeOf(&GithubLoginInput{})},
		conversion.GeneratedDeepCopyFunc{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*GithubPipelineConfig).DeepCopyInto(out.(*GithubPipelineConfig))
			return nil
		}, InType: reflect.TypeOf(&GithubPipelineConfig{})},
		conversion.GeneratedDeepCopyFunc{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*GithubPipelineConfigApplyInput).DeepCopyInto(out.(*GithubPipelineConfigApplyInput))
			return nil
		}, InType: reflect.TypeOf(&GithubPipelineConfigApplyInput{})},
		conversion.GeneratedDeepCopyFunc{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*GithubProvider).DeepCopyInto(out.(*GithubProvider))
			return nil
		}, InType: reflect.TypeOf(&GithubProvider{})},
		conversion.GeneratedDeepCopyFunc{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*GitlabLoginInput).DeepCopyInto(out.(*GitlabLoginInput))
			return nil
		}, InType: reflect.TypeOf(&GitlabLoginInput{})},
		conversion.GeneratedDeepCopyFunc{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*GitlabPipelineConfig).DeepCopyInto(out.(*GitlabPipelineConfig))
			return nil
		}, InType: reflect.TypeOf(&GitlabPipelineConfig{})},
		conversion.GeneratedDeepCopyFunc{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*GitlabPipelineConfigApplyInput).DeepCopyInto(out.(*GitlabPipelineConfigApplyInput))
			return nil
		}, InType: reflect.TypeOf(&GitlabPipelineConfigApplyInput{})},
		conversion.GeneratedDeepCopyFunc{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*GitlabProvider).DeepCopyInto(out.(*GitlabProvider))
			return nil
		}, InType: reflect.TypeOf(&GitlabProvider{})},
		conversion.GeneratedDeepCopyFunc{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*NamespacedBasicAuth).DeepCopyInto(out.(*NamespacedBasicAuth))
			return nil
		}, InType: reflect.TypeOf(&NamespacedBasicAuth{})},
		conversion.GeneratedDeepCopyFunc{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*NamespacedBasicAuthList).DeepCopyInto(out.(*NamespacedBasicAuthList))
			return nil
		}, InType: reflect.TypeOf(&NamespacedBasicAuthList{})},
		conversion.GeneratedDeepCopyFunc{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*NamespacedCertificate).DeepCopyInto(out.(*NamespacedCertificate))
			return nil
		}, InType: reflect.TypeOf(&NamespacedCertificate{})},
		conversion.GeneratedDeepCopyFunc{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*NamespacedCertificateList).DeepCopyInto(out.(*NamespacedCertificateList))
			return nil
		}, InType: reflect.TypeOf(&NamespacedCertificateList{})},
		conversion.GeneratedDeepCopyFunc{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*NamespacedDockerCredential).DeepCopyInto(out.(*NamespacedDockerCredential))
			return nil
		}, InType: reflect.TypeOf(&NamespacedDockerCredential{})},
		conversion.GeneratedDeepCopyFunc{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*NamespacedDockerCredentialList).DeepCopyInto(out.(*NamespacedDockerCredentialList))
			return nil
		}, InType: reflect.TypeOf(&NamespacedDockerCredentialList{})},
		conversion.GeneratedDeepCopyFunc{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*NamespacedSSHAuth).DeepCopyInto(out.(*NamespacedSSHAuth))
			return nil
		}, InType: reflect.TypeOf(&NamespacedSSHAuth{})},
		conversion.GeneratedDeepCopyFunc{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*NamespacedSSHAuthList).DeepCopyInto(out.(*NamespacedSSHAuthList))
			return nil
		}, InType: reflect.TypeOf(&NamespacedSSHAuthList{})},
		conversion.GeneratedDeepCopyFunc{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*NamespacedServiceAccountToken).DeepCopyInto(out.(*NamespacedServiceAccountToken))
			return nil
		}, InType: reflect.TypeOf(&NamespacedServiceAccountToken{})},
		conversion.GeneratedDeepCopyFunc{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*NamespacedServiceAccountTokenList).DeepCopyInto(out.(*NamespacedServiceAccountTokenList))
			return nil
		}, InType: reflect.TypeOf(&NamespacedServiceAccountTokenList{})},
		conversion.GeneratedDeepCopyFunc{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*Pipeline).DeepCopyInto(out.(*Pipeline))
			return nil
		}, InType: reflect.TypeOf(&Pipeline{})},
		conversion.GeneratedDeepCopyFunc{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*PipelineCondition).DeepCopyInto(out.(*PipelineCondition))
			return nil
		}, InType: reflect.TypeOf(&PipelineCondition{})},
		conversion.GeneratedDeepCopyFunc{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*PipelineConfig).DeepCopyInto(out.(*PipelineConfig))
			return nil
		}, InType: reflect.TypeOf(&PipelineConfig{})},
		conversion.GeneratedDeepCopyFunc{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*PipelineExecution).DeepCopyInto(out.(*PipelineExecution))
			return nil
		}, InType: reflect.TypeOf(&PipelineExecution{})},
		conversion.GeneratedDeepCopyFunc{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*PipelineExecutionList).DeepCopyInto(out.(*PipelineExecutionList))
			return nil
		}, InType: reflect.TypeOf(&PipelineExecutionList{})},
		conversion.GeneratedDeepCopyFunc{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*PipelineExecutionSpec).DeepCopyInto(out.(*PipelineExecutionSpec))
			return nil
		}, InType: reflect.TypeOf(&PipelineExecutionSpec{})},
		conversion.GeneratedDeepCopyFunc{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*PipelineExecutionStatus).DeepCopyInto(out.(*PipelineExecutionStatus))
			return nil
		}, InType: reflect.TypeOf(&PipelineExecutionStatus{})},
		conversion.GeneratedDeepCopyFunc{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*PipelineList).DeepCopyInto(out.(*PipelineList))
			return nil
		}, InType: reflect.TypeOf(&PipelineList{})},
		conversion.GeneratedDeepCopyFunc{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*PipelineSetting).DeepCopyInto(out.(*PipelineSetting))
			return nil
		}, InType: reflect.TypeOf(&PipelineSetting{})},
		conversion.GeneratedDeepCopyFunc{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*PipelineSettingList).DeepCopyInto(out.(*PipelineSettingList))
			return nil
		}, InType: reflect.TypeOf(&PipelineSettingList{})},
		conversion.GeneratedDeepCopyFunc{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*PipelineSpec).DeepCopyInto(out.(*PipelineSpec))
			return nil
		}, InType: reflect.TypeOf(&PipelineSpec{})},
		conversion.GeneratedDeepCopyFunc{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*PipelineStatus).DeepCopyInto(out.(*PipelineStatus))
			return nil
		}, InType: reflect.TypeOf(&PipelineStatus{})},
		conversion.GeneratedDeepCopyFunc{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*PipelineSystemImages).DeepCopyInto(out.(*PipelineSystemImages))
			return nil
		}, InType: reflect.TypeOf(&PipelineSystemImages{})},
		conversion.GeneratedDeepCopyFunc{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*PublicEndpoint).DeepCopyInto(out.(*PublicEndpoint))
			return nil
		}, InType: reflect.TypeOf(&PublicEndpoint{})},
		conversion.GeneratedDeepCopyFunc{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*PublishImageConfig).DeepCopyInto(out.(*PublishImageConfig))
			return nil
		}, InType: reflect.TypeOf(&PublishImageConfig{})},
		conversion.GeneratedDeepCopyFunc{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*PushPipelineConfigInput).DeepCopyInto(out.(*PushPipelineConfigInput))
			return nil
		}, InType: reflect.TypeOf(&PushPipelineConfigInput{})},
		conversion.GeneratedDeepCopyFunc{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*RegistryCredential).DeepCopyInto(out.(*RegistryCredential))
			return nil
		}, InType: reflect.TypeOf(&RegistryCredential{})},
		conversion.GeneratedDeepCopyFunc{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*RepoPerm).DeepCopyInto(out.(*RepoPerm))
			return nil
		}, InType: reflect.TypeOf(&RepoPerm{})},
		conversion.GeneratedDeepCopyFunc{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*RollbackRevision).DeepCopyInto(out.(*RollbackRevision))
			return nil
		}, InType: reflect.TypeOf(&RollbackRevision{})},
		conversion.GeneratedDeepCopyFunc{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*RunPipelineInput).DeepCopyInto(out.(*RunPipelineInput))
			return nil
		}, InType: reflect.TypeOf(&RunPipelineInput{})},
		conversion.GeneratedDeepCopyFunc{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*RunScriptConfig).DeepCopyInto(out.(*RunScriptConfig))
			return nil
		}, InType: reflect.TypeOf(&RunScriptConfig{})},
		conversion.GeneratedDeepCopyFunc{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*SSHAuth).DeepCopyInto(out.(*SSHAuth))
			return nil
		}, InType: reflect.TypeOf(&SSHAuth{})},
		conversion.GeneratedDeepCopyFunc{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*SSHAuthList).DeepCopyInto(out.(*SSHAuthList))
			return nil
		}, InType: reflect.TypeOf(&SSHAuthList{})},
		conversion.GeneratedDeepCopyFunc{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*ServiceAccountToken).DeepCopyInto(out.(*ServiceAccountToken))
			return nil
		}, InType: reflect.TypeOf(&ServiceAccountToken{})},
		conversion.GeneratedDeepCopyFunc{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*ServiceAccountTokenList).DeepCopyInto(out.(*ServiceAccountTokenList))
			return nil
		}, InType: reflect.TypeOf(&ServiceAccountTokenList{})},
		conversion.GeneratedDeepCopyFunc{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*SourceCodeConfig).DeepCopyInto(out.(*SourceCodeConfig))
			return nil
		}, InType: reflect.TypeOf(&SourceCodeConfig{})},
		conversion.GeneratedDeepCopyFunc{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*SourceCodeCredential).DeepCopyInto(out.(*SourceCodeCredential))
			return nil
		}, InType: reflect.TypeOf(&SourceCodeCredential{})},
		conversion.GeneratedDeepCopyFunc{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*SourceCodeCredentialList).DeepCopyInto(out.(*SourceCodeCredentialList))
			return nil
		}, InType: reflect.TypeOf(&SourceCodeCredentialList{})},
		conversion.GeneratedDeepCopyFunc{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*SourceCodeCredentialSpec).DeepCopyInto(out.(*SourceCodeCredentialSpec))
			return nil
		}, InType: reflect.TypeOf(&SourceCodeCredentialSpec{})},
		conversion.GeneratedDeepCopyFunc{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*SourceCodeCredentialStatus).DeepCopyInto(out.(*SourceCodeCredentialStatus))
			return nil
		}, InType: reflect.TypeOf(&SourceCodeCredentialStatus{})},
		conversion.GeneratedDeepCopyFunc{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*SourceCodeProvider).DeepCopyInto(out.(*SourceCodeProvider))
			return nil
		}, InType: reflect.TypeOf(&SourceCodeProvider{})},
		conversion.GeneratedDeepCopyFunc{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*SourceCodeProviderConfig).DeepCopyInto(out.(*SourceCodeProviderConfig))
			return nil
		}, InType: reflect.TypeOf(&SourceCodeProviderConfig{})},
		conversion.GeneratedDeepCopyFunc{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*SourceCodeProviderConfigList).DeepCopyInto(out.(*SourceCodeProviderConfigList))
			return nil
		}, InType: reflect.TypeOf(&SourceCodeProviderConfigList{})},
		conversion.GeneratedDeepCopyFunc{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*SourceCodeProviderList).DeepCopyInto(out.(*SourceCodeProviderList))
			return nil
		}, InType: reflect.TypeOf(&SourceCodeProviderList{})},
		conversion.GeneratedDeepCopyFunc{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*SourceCodeRepository).DeepCopyInto(out.(*SourceCodeRepository))
			return nil
		}, InType: reflect.TypeOf(&SourceCodeRepository{})},
		conversion.GeneratedDeepCopyFunc{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*SourceCodeRepositoryList).DeepCopyInto(out.(*SourceCodeRepositoryList))
			return nil
		}, InType: reflect.TypeOf(&SourceCodeRepositoryList{})},
		conversion.GeneratedDeepCopyFunc{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*SourceCodeRepositorySpec).DeepCopyInto(out.(*SourceCodeRepositorySpec))
			return nil
		}, InType: reflect.TypeOf(&SourceCodeRepositorySpec{})},
		conversion.GeneratedDeepCopyFunc{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*SourceCodeRepositoryStatus).DeepCopyInto(out.(*SourceCodeRepositoryStatus))
			return nil
		}, InType: reflect.TypeOf(&SourceCodeRepositoryStatus{})},
		conversion.GeneratedDeepCopyFunc{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*Stage).DeepCopyInto(out.(*Stage))
			return nil
		}, InType: reflect.TypeOf(&Stage{})},
		conversion.GeneratedDeepCopyFunc{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*StageStatus).DeepCopyInto(out.(*StageStatus))
			return nil
		}, InType: reflect.TypeOf(&StageStatus{})},
		conversion.GeneratedDeepCopyFunc{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*Step).DeepCopyInto(out.(*Step))
			return nil
		}, InType: reflect.TypeOf(&Step{})},
		conversion.GeneratedDeepCopyFunc{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*StepStatus).DeepCopyInto(out.(*StepStatus))
			return nil
		}, InType: reflect.TypeOf(&StepStatus{})},
		conversion.GeneratedDeepCopyFunc{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*Workload).DeepCopyInto(out.(*Workload))
			return nil
		}, InType: reflect.TypeOf(&Workload{})},
		conversion.GeneratedDeepCopyFunc{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*WorkloadList).DeepCopyInto(out.(*WorkloadList))
			return nil
		}, InType: reflect.TypeOf(&WorkloadList{})},
	)
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *App) DeepCopyInto(out *App) {
	*out = *in
	out.Namespaced = in.Namespaced
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new App.
func (in *App) DeepCopy() *App {
	if in == nil {
		return nil
	}
	out := new(App)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *App) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	} else {
		return nil
	}
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AppCondition) DeepCopyInto(out *AppCondition) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AppCondition.
func (in *AppCondition) DeepCopy() *AppCondition {
	if in == nil {
		return nil
	}
	out := new(AppCondition)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AppList) DeepCopyInto(out *AppList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]App, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AppList.
func (in *AppList) DeepCopy() *AppList {
	if in == nil {
		return nil
	}
	out := new(AppList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AppList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	} else {
		return nil
	}
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AppRevision) DeepCopyInto(out *AppRevision) {
	*out = *in
	out.Namespaced = in.Namespaced
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AppRevision.
func (in *AppRevision) DeepCopy() *AppRevision {
	if in == nil {
		return nil
	}
	out := new(AppRevision)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AppRevision) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	} else {
		return nil
	}
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AppRevisionList) DeepCopyInto(out *AppRevisionList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]AppRevision, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AppRevisionList.
func (in *AppRevisionList) DeepCopy() *AppRevisionList {
	if in == nil {
		return nil
	}
	out := new(AppRevisionList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AppRevisionList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	} else {
		return nil
	}
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AppRevisionSpec) DeepCopyInto(out *AppRevisionSpec) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AppRevisionSpec.
func (in *AppRevisionSpec) DeepCopy() *AppRevisionSpec {
	if in == nil {
		return nil
	}
	out := new(AppRevisionSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AppRevisionStatus) DeepCopyInto(out *AppRevisionStatus) {
	*out = *in
	if in.Answers != nil {
		in, out := &in.Answers, &out.Answers
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AppRevisionStatus.
func (in *AppRevisionStatus) DeepCopy() *AppRevisionStatus {
	if in == nil {
		return nil
	}
	out := new(AppRevisionStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AppSpec) DeepCopyInto(out *AppSpec) {
	*out = *in
	if in.Files != nil {
		in, out := &in.Files, &out.Files
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Answers != nil {
		in, out := &in.Answers, &out.Answers
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AppSpec.
func (in *AppSpec) DeepCopy() *AppSpec {
	if in == nil {
		return nil
	}
	out := new(AppSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AppStatus) DeepCopyInto(out *AppStatus) {
	*out = *in
	if in.AppliedFiles != nil {
		in, out := &in.AppliedFiles, &out.AppliedFiles
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]AppCondition, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AppStatus.
func (in *AppStatus) DeepCopy() *AppStatus {
	if in == nil {
		return nil
	}
	out := new(AppStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AppUpgradeConfig) DeepCopyInto(out *AppUpgradeConfig) {
	*out = *in
	if in.Answers != nil {
		in, out := &in.Answers, &out.Answers
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AppUpgradeConfig.
func (in *AppUpgradeConfig) DeepCopy() *AppUpgradeConfig {
	if in == nil {
		return nil
	}
	out := new(AppUpgradeConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ApplyYamlConfig) DeepCopyInto(out *ApplyYamlConfig) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ApplyYamlConfig.
func (in *ApplyYamlConfig) DeepCopy() *ApplyYamlConfig {
	if in == nil {
		return nil
	}
	out := new(ApplyYamlConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AuthAppInput) DeepCopyInto(out *AuthAppInput) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AuthAppInput.
func (in *AuthAppInput) DeepCopy() *AuthAppInput {
	if in == nil {
		return nil
	}
	out := new(AuthAppInput)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AuthUserInput) DeepCopyInto(out *AuthUserInput) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AuthUserInput.
func (in *AuthUserInput) DeepCopy() *AuthUserInput {
	if in == nil {
		return nil
	}
	out := new(AuthUserInput)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BasicAuth) DeepCopyInto(out *BasicAuth) {
	*out = *in
	out.Namespaced = in.Namespaced
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BasicAuth.
func (in *BasicAuth) DeepCopy() *BasicAuth {
	if in == nil {
		return nil
	}
	out := new(BasicAuth)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *BasicAuth) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	} else {
		return nil
	}
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BasicAuthList) DeepCopyInto(out *BasicAuthList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]BasicAuth, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BasicAuthList.
func (in *BasicAuthList) DeepCopy() *BasicAuthList {
	if in == nil {
		return nil
	}
	out := new(BasicAuthList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *BasicAuthList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	} else {
		return nil
	}
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Certificate) DeepCopyInto(out *Certificate) {
	*out = *in
	out.Namespaced = in.Namespaced
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	if in.SubjectAlternativeNames != nil {
		in, out := &in.SubjectAlternativeNames, &out.SubjectAlternativeNames
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Certificate.
func (in *Certificate) DeepCopy() *Certificate {
	if in == nil {
		return nil
	}
	out := new(Certificate)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Certificate) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	} else {
		return nil
	}
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CertificateList) DeepCopyInto(out *CertificateList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Certificate, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CertificateList.
func (in *CertificateList) DeepCopy() *CertificateList {
	if in == nil {
		return nil
	}
	out := new(CertificateList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *CertificateList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	} else {
		return nil
	}
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Constraint) DeepCopyInto(out *Constraint) {
	*out = *in
	if in.Include != nil {
		in, out := &in.Include, &out.Include
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Exclude != nil {
		in, out := &in.Exclude, &out.Exclude
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Constraint.
func (in *Constraint) DeepCopy() *Constraint {
	if in == nil {
		return nil
	}
	out := new(Constraint)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Constraints) DeepCopyInto(out *Constraints) {
	*out = *in
	if in.Branch != nil {
		in, out := &in.Branch, &out.Branch
		if *in == nil {
			*out = nil
		} else {
			*out = new(Constraint)
			(*in).DeepCopyInto(*out)
		}
	}
	if in.Event != nil {
		in, out := &in.Event, &out.Event
		if *in == nil {
			*out = nil
		} else {
			*out = new(Constraint)
			(*in).DeepCopyInto(*out)
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Constraints.
func (in *Constraints) DeepCopy() *Constraints {
	if in == nil {
		return nil
	}
	out := new(Constraints)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DeploymentRollbackInput) DeepCopyInto(out *DeploymentRollbackInput) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DeploymentRollbackInput.
func (in *DeploymentRollbackInput) DeepCopy() *DeploymentRollbackInput {
	if in == nil {
		return nil
	}
	out := new(DeploymentRollbackInput)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DockerCredential) DeepCopyInto(out *DockerCredential) {
	*out = *in
	out.Namespaced = in.Namespaced
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	if in.Registries != nil {
		in, out := &in.Registries, &out.Registries
		*out = make(map[string]RegistryCredential, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DockerCredential.
func (in *DockerCredential) DeepCopy() *DockerCredential {
	if in == nil {
		return nil
	}
	out := new(DockerCredential)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *DockerCredential) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	} else {
		return nil
	}
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DockerCredentialList) DeepCopyInto(out *DockerCredentialList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]DockerCredential, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DockerCredentialList.
func (in *DockerCredentialList) DeepCopy() *DockerCredentialList {
	if in == nil {
		return nil
	}
	out := new(DockerCredentialList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *DockerCredentialList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	} else {
		return nil
	}
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EnvFrom) DeepCopyInto(out *EnvFrom) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EnvFrom.
func (in *EnvFrom) DeepCopy() *EnvFrom {
	if in == nil {
		return nil
	}
	out := new(EnvFrom)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GithubLoginInput) DeepCopyInto(out *GithubLoginInput) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GithubLoginInput.
func (in *GithubLoginInput) DeepCopy() *GithubLoginInput {
	if in == nil {
		return nil
	}
	out := new(GithubLoginInput)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GithubPipelineConfig) DeepCopyInto(out *GithubPipelineConfig) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.SourceCodeProviderConfig.DeepCopyInto(&out.SourceCodeProviderConfig)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GithubPipelineConfig.
func (in *GithubPipelineConfig) DeepCopy() *GithubPipelineConfig {
	if in == nil {
		return nil
	}
	out := new(GithubPipelineConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *GithubPipelineConfig) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	} else {
		return nil
	}
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GithubPipelineConfigApplyInput) DeepCopyInto(out *GithubPipelineConfigApplyInput) {
	*out = *in
	in.GithubConfig.DeepCopyInto(&out.GithubConfig)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GithubPipelineConfigApplyInput.
func (in *GithubPipelineConfigApplyInput) DeepCopy() *GithubPipelineConfigApplyInput {
	if in == nil {
		return nil
	}
	out := new(GithubPipelineConfigApplyInput)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GithubProvider) DeepCopyInto(out *GithubProvider) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.SourceCodeProvider.DeepCopyInto(&out.SourceCodeProvider)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GithubProvider.
func (in *GithubProvider) DeepCopy() *GithubProvider {
	if in == nil {
		return nil
	}
	out := new(GithubProvider)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *GithubProvider) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	} else {
		return nil
	}
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GitlabLoginInput) DeepCopyInto(out *GitlabLoginInput) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GitlabLoginInput.
func (in *GitlabLoginInput) DeepCopy() *GitlabLoginInput {
	if in == nil {
		return nil
	}
	out := new(GitlabLoginInput)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GitlabPipelineConfig) DeepCopyInto(out *GitlabPipelineConfig) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.SourceCodeProviderConfig.DeepCopyInto(&out.SourceCodeProviderConfig)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GitlabPipelineConfig.
func (in *GitlabPipelineConfig) DeepCopy() *GitlabPipelineConfig {
	if in == nil {
		return nil
	}
	out := new(GitlabPipelineConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *GitlabPipelineConfig) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	} else {
		return nil
	}
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GitlabPipelineConfigApplyInput) DeepCopyInto(out *GitlabPipelineConfigApplyInput) {
	*out = *in
	in.GitlabConfig.DeepCopyInto(&out.GitlabConfig)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GitlabPipelineConfigApplyInput.
func (in *GitlabPipelineConfigApplyInput) DeepCopy() *GitlabPipelineConfigApplyInput {
	if in == nil {
		return nil
	}
	out := new(GitlabPipelineConfigApplyInput)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GitlabProvider) DeepCopyInto(out *GitlabProvider) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.SourceCodeProvider.DeepCopyInto(&out.SourceCodeProvider)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GitlabProvider.
func (in *GitlabProvider) DeepCopy() *GitlabProvider {
	if in == nil {
		return nil
	}
	out := new(GitlabProvider)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *GitlabProvider) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	} else {
		return nil
	}
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NamespacedBasicAuth) DeepCopyInto(out *NamespacedBasicAuth) {
	*out = *in
	out.Namespaced = in.Namespaced
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NamespacedBasicAuth.
func (in *NamespacedBasicAuth) DeepCopy() *NamespacedBasicAuth {
	if in == nil {
		return nil
	}
	out := new(NamespacedBasicAuth)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *NamespacedBasicAuth) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	} else {
		return nil
	}
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NamespacedBasicAuthList) DeepCopyInto(out *NamespacedBasicAuthList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]NamespacedBasicAuth, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NamespacedBasicAuthList.
func (in *NamespacedBasicAuthList) DeepCopy() *NamespacedBasicAuthList {
	if in == nil {
		return nil
	}
	out := new(NamespacedBasicAuthList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *NamespacedBasicAuthList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	} else {
		return nil
	}
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NamespacedCertificate) DeepCopyInto(out *NamespacedCertificate) {
	*out = *in
	out.Namespaced = in.Namespaced
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	if in.SubjectAlternativeNames != nil {
		in, out := &in.SubjectAlternativeNames, &out.SubjectAlternativeNames
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NamespacedCertificate.
func (in *NamespacedCertificate) DeepCopy() *NamespacedCertificate {
	if in == nil {
		return nil
	}
	out := new(NamespacedCertificate)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *NamespacedCertificate) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	} else {
		return nil
	}
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NamespacedCertificateList) DeepCopyInto(out *NamespacedCertificateList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]NamespacedCertificate, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NamespacedCertificateList.
func (in *NamespacedCertificateList) DeepCopy() *NamespacedCertificateList {
	if in == nil {
		return nil
	}
	out := new(NamespacedCertificateList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *NamespacedCertificateList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	} else {
		return nil
	}
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NamespacedDockerCredential) DeepCopyInto(out *NamespacedDockerCredential) {
	*out = *in
	out.Namespaced = in.Namespaced
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	if in.Registries != nil {
		in, out := &in.Registries, &out.Registries
		*out = make(map[string]RegistryCredential, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NamespacedDockerCredential.
func (in *NamespacedDockerCredential) DeepCopy() *NamespacedDockerCredential {
	if in == nil {
		return nil
	}
	out := new(NamespacedDockerCredential)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *NamespacedDockerCredential) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	} else {
		return nil
	}
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NamespacedDockerCredentialList) DeepCopyInto(out *NamespacedDockerCredentialList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]NamespacedDockerCredential, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NamespacedDockerCredentialList.
func (in *NamespacedDockerCredentialList) DeepCopy() *NamespacedDockerCredentialList {
	if in == nil {
		return nil
	}
	out := new(NamespacedDockerCredentialList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *NamespacedDockerCredentialList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	} else {
		return nil
	}
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NamespacedSSHAuth) DeepCopyInto(out *NamespacedSSHAuth) {
	*out = *in
	out.Namespaced = in.Namespaced
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NamespacedSSHAuth.
func (in *NamespacedSSHAuth) DeepCopy() *NamespacedSSHAuth {
	if in == nil {
		return nil
	}
	out := new(NamespacedSSHAuth)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *NamespacedSSHAuth) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	} else {
		return nil
	}
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NamespacedSSHAuthList) DeepCopyInto(out *NamespacedSSHAuthList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]NamespacedSSHAuth, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NamespacedSSHAuthList.
func (in *NamespacedSSHAuthList) DeepCopy() *NamespacedSSHAuthList {
	if in == nil {
		return nil
	}
	out := new(NamespacedSSHAuthList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *NamespacedSSHAuthList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	} else {
		return nil
	}
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NamespacedServiceAccountToken) DeepCopyInto(out *NamespacedServiceAccountToken) {
	*out = *in
	out.Namespaced = in.Namespaced
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NamespacedServiceAccountToken.
func (in *NamespacedServiceAccountToken) DeepCopy() *NamespacedServiceAccountToken {
	if in == nil {
		return nil
	}
	out := new(NamespacedServiceAccountToken)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *NamespacedServiceAccountToken) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	} else {
		return nil
	}
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NamespacedServiceAccountTokenList) DeepCopyInto(out *NamespacedServiceAccountTokenList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]NamespacedServiceAccountToken, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NamespacedServiceAccountTokenList.
func (in *NamespacedServiceAccountTokenList) DeepCopy() *NamespacedServiceAccountTokenList {
	if in == nil {
		return nil
	}
	out := new(NamespacedServiceAccountTokenList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *NamespacedServiceAccountTokenList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	} else {
		return nil
	}
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Pipeline) DeepCopyInto(out *Pipeline) {
	*out = *in
	out.Namespaced = in.Namespaced
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Pipeline.
func (in *Pipeline) DeepCopy() *Pipeline {
	if in == nil {
		return nil
	}
	out := new(Pipeline)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Pipeline) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	} else {
		return nil
	}
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PipelineCondition) DeepCopyInto(out *PipelineCondition) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PipelineCondition.
func (in *PipelineCondition) DeepCopy() *PipelineCondition {
	if in == nil {
		return nil
	}
	out := new(PipelineCondition)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PipelineConfig) DeepCopyInto(out *PipelineConfig) {
	*out = *in
	if in.Stages != nil {
		in, out := &in.Stages, &out.Stages
		*out = make([]Stage, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Branch != nil {
		in, out := &in.Branch, &out.Branch
		if *in == nil {
			*out = nil
		} else {
			*out = new(Constraint)
			(*in).DeepCopyInto(*out)
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PipelineConfig.
func (in *PipelineConfig) DeepCopy() *PipelineConfig {
	if in == nil {
		return nil
	}
	out := new(PipelineConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PipelineExecution) DeepCopyInto(out *PipelineExecution) {
	*out = *in
	out.Namespaced = in.Namespaced
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PipelineExecution.
func (in *PipelineExecution) DeepCopy() *PipelineExecution {
	if in == nil {
		return nil
	}
	out := new(PipelineExecution)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *PipelineExecution) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	} else {
		return nil
	}
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PipelineExecutionList) DeepCopyInto(out *PipelineExecutionList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]PipelineExecution, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PipelineExecutionList.
func (in *PipelineExecutionList) DeepCopy() *PipelineExecutionList {
	if in == nil {
		return nil
	}
	out := new(PipelineExecutionList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *PipelineExecutionList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	} else {
		return nil
	}
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PipelineExecutionSpec) DeepCopyInto(out *PipelineExecutionSpec) {
	*out = *in
	in.PipelineConfig.DeepCopyInto(&out.PipelineConfig)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PipelineExecutionSpec.
func (in *PipelineExecutionSpec) DeepCopy() *PipelineExecutionSpec {
	if in == nil {
		return nil
	}
	out := new(PipelineExecutionSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PipelineExecutionStatus) DeepCopyInto(out *PipelineExecutionStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]PipelineCondition, len(*in))
		copy(*out, *in)
	}
	if in.Stages != nil {
		in, out := &in.Stages, &out.Stages
		*out = make([]StageStatus, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PipelineExecutionStatus.
func (in *PipelineExecutionStatus) DeepCopy() *PipelineExecutionStatus {
	if in == nil {
		return nil
	}
	out := new(PipelineExecutionStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PipelineList) DeepCopyInto(out *PipelineList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Pipeline, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PipelineList.
func (in *PipelineList) DeepCopy() *PipelineList {
	if in == nil {
		return nil
	}
	out := new(PipelineList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *PipelineList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	} else {
		return nil
	}
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PipelineSetting) DeepCopyInto(out *PipelineSetting) {
	*out = *in
	out.Namespaced = in.Namespaced
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PipelineSetting.
func (in *PipelineSetting) DeepCopy() *PipelineSetting {
	if in == nil {
		return nil
	}
	out := new(PipelineSetting)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *PipelineSetting) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	} else {
		return nil
	}
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PipelineSettingList) DeepCopyInto(out *PipelineSettingList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]PipelineSetting, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PipelineSettingList.
func (in *PipelineSettingList) DeepCopy() *PipelineSettingList {
	if in == nil {
		return nil
	}
	out := new(PipelineSettingList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *PipelineSettingList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	} else {
		return nil
	}
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PipelineSpec) DeepCopyInto(out *PipelineSpec) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PipelineSpec.
func (in *PipelineSpec) DeepCopy() *PipelineSpec {
	if in == nil {
		return nil
	}
	out := new(PipelineSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PipelineStatus) DeepCopyInto(out *PipelineStatus) {
	*out = *in
	if in.SourceCodeCredential != nil {
		in, out := &in.SourceCodeCredential, &out.SourceCodeCredential
		if *in == nil {
			*out = nil
		} else {
			*out = new(SourceCodeCredential)
			(*in).DeepCopyInto(*out)
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PipelineStatus.
func (in *PipelineStatus) DeepCopy() *PipelineStatus {
	if in == nil {
		return nil
	}
	out := new(PipelineStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PipelineSystemImages) DeepCopyInto(out *PipelineSystemImages) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PipelineSystemImages.
func (in *PipelineSystemImages) DeepCopy() *PipelineSystemImages {
	if in == nil {
		return nil
	}
	out := new(PipelineSystemImages)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PublicEndpoint) DeepCopyInto(out *PublicEndpoint) {
	*out = *in
	if in.Addresses != nil {
		in, out := &in.Addresses, &out.Addresses
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PublicEndpoint.
func (in *PublicEndpoint) DeepCopy() *PublicEndpoint {
	if in == nil {
		return nil
	}
	out := new(PublicEndpoint)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PublishImageConfig) DeepCopyInto(out *PublishImageConfig) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PublishImageConfig.
func (in *PublishImageConfig) DeepCopy() *PublishImageConfig {
	if in == nil {
		return nil
	}
	out := new(PublishImageConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PushPipelineConfigInput) DeepCopyInto(out *PushPipelineConfigInput) {
	*out = *in
	if in.Configs != nil {
		in, out := &in.Configs, &out.Configs
		*out = make(map[string]PipelineConfig, len(*in))
		for key, val := range *in {
			(*out)[key] = *val.DeepCopy()
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PushPipelineConfigInput.
func (in *PushPipelineConfigInput) DeepCopy() *PushPipelineConfigInput {
	if in == nil {
		return nil
	}
	out := new(PushPipelineConfigInput)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RegistryCredential) DeepCopyInto(out *RegistryCredential) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RegistryCredential.
func (in *RegistryCredential) DeepCopy() *RegistryCredential {
	if in == nil {
		return nil
	}
	out := new(RegistryCredential)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RepoPerm) DeepCopyInto(out *RepoPerm) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RepoPerm.
func (in *RepoPerm) DeepCopy() *RepoPerm {
	if in == nil {
		return nil
	}
	out := new(RepoPerm)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RollbackRevision) DeepCopyInto(out *RollbackRevision) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RollbackRevision.
func (in *RollbackRevision) DeepCopy() *RollbackRevision {
	if in == nil {
		return nil
	}
	out := new(RollbackRevision)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RunPipelineInput) DeepCopyInto(out *RunPipelineInput) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RunPipelineInput.
func (in *RunPipelineInput) DeepCopy() *RunPipelineInput {
	if in == nil {
		return nil
	}
	out := new(RunPipelineInput)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RunScriptConfig) DeepCopyInto(out *RunScriptConfig) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RunScriptConfig.
func (in *RunScriptConfig) DeepCopy() *RunScriptConfig {
	if in == nil {
		return nil
	}
	out := new(RunScriptConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SSHAuth) DeepCopyInto(out *SSHAuth) {
	*out = *in
	out.Namespaced = in.Namespaced
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SSHAuth.
func (in *SSHAuth) DeepCopy() *SSHAuth {
	if in == nil {
		return nil
	}
	out := new(SSHAuth)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *SSHAuth) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	} else {
		return nil
	}
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SSHAuthList) DeepCopyInto(out *SSHAuthList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]SSHAuth, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SSHAuthList.
func (in *SSHAuthList) DeepCopy() *SSHAuthList {
	if in == nil {
		return nil
	}
	out := new(SSHAuthList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *SSHAuthList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	} else {
		return nil
	}
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServiceAccountToken) DeepCopyInto(out *ServiceAccountToken) {
	*out = *in
	out.Namespaced = in.Namespaced
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServiceAccountToken.
func (in *ServiceAccountToken) DeepCopy() *ServiceAccountToken {
	if in == nil {
		return nil
	}
	out := new(ServiceAccountToken)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ServiceAccountToken) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	} else {
		return nil
	}
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServiceAccountTokenList) DeepCopyInto(out *ServiceAccountTokenList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ServiceAccountToken, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServiceAccountTokenList.
func (in *ServiceAccountTokenList) DeepCopy() *ServiceAccountTokenList {
	if in == nil {
		return nil
	}
	out := new(ServiceAccountTokenList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ServiceAccountTokenList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	} else {
		return nil
	}
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SourceCodeConfig) DeepCopyInto(out *SourceCodeConfig) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SourceCodeConfig.
func (in *SourceCodeConfig) DeepCopy() *SourceCodeConfig {
	if in == nil {
		return nil
	}
	out := new(SourceCodeConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SourceCodeCredential) DeepCopyInto(out *SourceCodeCredential) {
	*out = *in
	out.Namespaced = in.Namespaced
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	out.Status = in.Status
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SourceCodeCredential.
func (in *SourceCodeCredential) DeepCopy() *SourceCodeCredential {
	if in == nil {
		return nil
	}
	out := new(SourceCodeCredential)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *SourceCodeCredential) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	} else {
		return nil
	}
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SourceCodeCredentialList) DeepCopyInto(out *SourceCodeCredentialList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]SourceCodeCredential, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SourceCodeCredentialList.
func (in *SourceCodeCredentialList) DeepCopy() *SourceCodeCredentialList {
	if in == nil {
		return nil
	}
	out := new(SourceCodeCredentialList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *SourceCodeCredentialList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	} else {
		return nil
	}
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SourceCodeCredentialSpec) DeepCopyInto(out *SourceCodeCredentialSpec) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SourceCodeCredentialSpec.
func (in *SourceCodeCredentialSpec) DeepCopy() *SourceCodeCredentialSpec {
	if in == nil {
		return nil
	}
	out := new(SourceCodeCredentialSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SourceCodeCredentialStatus) DeepCopyInto(out *SourceCodeCredentialStatus) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SourceCodeCredentialStatus.
func (in *SourceCodeCredentialStatus) DeepCopy() *SourceCodeCredentialStatus {
	if in == nil {
		return nil
	}
	out := new(SourceCodeCredentialStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SourceCodeProvider) DeepCopyInto(out *SourceCodeProvider) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SourceCodeProvider.
func (in *SourceCodeProvider) DeepCopy() *SourceCodeProvider {
	if in == nil {
		return nil
	}
	out := new(SourceCodeProvider)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *SourceCodeProvider) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	} else {
		return nil
	}
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SourceCodeProviderConfig) DeepCopyInto(out *SourceCodeProviderConfig) {
	*out = *in
	out.Namespaced = in.Namespaced
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SourceCodeProviderConfig.
func (in *SourceCodeProviderConfig) DeepCopy() *SourceCodeProviderConfig {
	if in == nil {
		return nil
	}
	out := new(SourceCodeProviderConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *SourceCodeProviderConfig) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	} else {
		return nil
	}
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SourceCodeProviderConfigList) DeepCopyInto(out *SourceCodeProviderConfigList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]SourceCodeProviderConfig, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SourceCodeProviderConfigList.
func (in *SourceCodeProviderConfigList) DeepCopy() *SourceCodeProviderConfigList {
	if in == nil {
		return nil
	}
	out := new(SourceCodeProviderConfigList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *SourceCodeProviderConfigList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	} else {
		return nil
	}
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SourceCodeProviderList) DeepCopyInto(out *SourceCodeProviderList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]SourceCodeProvider, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SourceCodeProviderList.
func (in *SourceCodeProviderList) DeepCopy() *SourceCodeProviderList {
	if in == nil {
		return nil
	}
	out := new(SourceCodeProviderList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *SourceCodeProviderList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	} else {
		return nil
	}
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SourceCodeRepository) DeepCopyInto(out *SourceCodeRepository) {
	*out = *in
	out.Namespaced = in.Namespaced
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	out.Status = in.Status
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SourceCodeRepository.
func (in *SourceCodeRepository) DeepCopy() *SourceCodeRepository {
	if in == nil {
		return nil
	}
	out := new(SourceCodeRepository)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *SourceCodeRepository) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	} else {
		return nil
	}
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SourceCodeRepositoryList) DeepCopyInto(out *SourceCodeRepositoryList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]SourceCodeRepository, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SourceCodeRepositoryList.
func (in *SourceCodeRepositoryList) DeepCopy() *SourceCodeRepositoryList {
	if in == nil {
		return nil
	}
	out := new(SourceCodeRepositoryList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *SourceCodeRepositoryList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	} else {
		return nil
	}
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SourceCodeRepositorySpec) DeepCopyInto(out *SourceCodeRepositorySpec) {
	*out = *in
	out.Permissions = in.Permissions
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SourceCodeRepositorySpec.
func (in *SourceCodeRepositorySpec) DeepCopy() *SourceCodeRepositorySpec {
	if in == nil {
		return nil
	}
	out := new(SourceCodeRepositorySpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SourceCodeRepositoryStatus) DeepCopyInto(out *SourceCodeRepositoryStatus) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SourceCodeRepositoryStatus.
func (in *SourceCodeRepositoryStatus) DeepCopy() *SourceCodeRepositoryStatus {
	if in == nil {
		return nil
	}
	out := new(SourceCodeRepositoryStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Stage) DeepCopyInto(out *Stage) {
	*out = *in
	if in.Steps != nil {
		in, out := &in.Steps, &out.Steps
		*out = make([]Step, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.When != nil {
		in, out := &in.When, &out.When
		if *in == nil {
			*out = nil
		} else {
			*out = new(Constraints)
			(*in).DeepCopyInto(*out)
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Stage.
func (in *Stage) DeepCopy() *Stage {
	if in == nil {
		return nil
	}
	out := new(Stage)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StageStatus) DeepCopyInto(out *StageStatus) {
	*out = *in
	if in.Steps != nil {
		in, out := &in.Steps, &out.Steps
		*out = make([]StepStatus, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StageStatus.
func (in *StageStatus) DeepCopy() *StageStatus {
	if in == nil {
		return nil
	}
	out := new(StageStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Step) DeepCopyInto(out *Step) {
	*out = *in
	if in.SourceCodeConfig != nil {
		in, out := &in.SourceCodeConfig, &out.SourceCodeConfig
		if *in == nil {
			*out = nil
		} else {
			*out = new(SourceCodeConfig)
			**out = **in
		}
	}
	if in.RunScriptConfig != nil {
		in, out := &in.RunScriptConfig, &out.RunScriptConfig
		if *in == nil {
			*out = nil
		} else {
			*out = new(RunScriptConfig)
			**out = **in
		}
	}
	if in.PublishImageConfig != nil {
		in, out := &in.PublishImageConfig, &out.PublishImageConfig
		if *in == nil {
			*out = nil
		} else {
			*out = new(PublishImageConfig)
			**out = **in
		}
	}
	if in.ApplyYamlConfig != nil {
		in, out := &in.ApplyYamlConfig, &out.ApplyYamlConfig
		if *in == nil {
			*out = nil
		} else {
			*out = new(ApplyYamlConfig)
			**out = **in
		}
	}
	if in.Env != nil {
		in, out := &in.Env, &out.Env
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.EnvFrom != nil {
		in, out := &in.EnvFrom, &out.EnvFrom
		*out = make([]EnvFrom, len(*in))
		copy(*out, *in)
	}
	if in.When != nil {
		in, out := &in.When, &out.When
		if *in == nil {
			*out = nil
		} else {
			*out = new(Constraints)
			(*in).DeepCopyInto(*out)
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Step.
func (in *Step) DeepCopy() *Step {
	if in == nil {
		return nil
	}
	out := new(Step)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StepStatus) DeepCopyInto(out *StepStatus) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StepStatus.
func (in *StepStatus) DeepCopy() *StepStatus {
	if in == nil {
		return nil
	}
	out := new(StepStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Workload) DeepCopyInto(out *Workload) {
	*out = *in
	out.Namespaced = in.Namespaced
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Workload.
func (in *Workload) DeepCopy() *Workload {
	if in == nil {
		return nil
	}
	out := new(Workload)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Workload) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	} else {
		return nil
	}
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WorkloadList) DeepCopyInto(out *WorkloadList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Workload, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WorkloadList.
func (in *WorkloadList) DeepCopy() *WorkloadList {
	if in == nil {
		return nil
	}
	out := new(WorkloadList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *WorkloadList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	} else {
		return nil
	}
}
