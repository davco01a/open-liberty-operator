//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*


Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by controller-gen. DO NOT EDIT.

package v1beta2

import (
	monitoringv1 "github.com/coreos/prometheus-operator/pkg/apis/monitoring/v1"
	routev1 "github.com/openshift/api/route/v1"
	appsv1 "k8s.io/api/apps/v1"
	"k8s.io/api/core/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GithubLogin) DeepCopyInto(out *GithubLogin) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GithubLogin.
func (in *GithubLogin) DeepCopy() *GithubLogin {
	if in == nil {
		return nil
	}
	out := new(GithubLogin)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OAuth2Client) DeepCopyInto(out *OAuth2Client) {
	*out = *in
	if in.AccessTokenRequired != nil {
		in, out := &in.AccessTokenRequired, &out.AccessTokenRequired
		*out = new(bool)
		**out = **in
	}
	if in.AccessTokenSupported != nil {
		in, out := &in.AccessTokenSupported, &out.AccessTokenSupported
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OAuth2Client.
func (in *OAuth2Client) DeepCopy() *OAuth2Client {
	if in == nil {
		return nil
	}
	out := new(OAuth2Client)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OidcClient) DeepCopyInto(out *OidcClient) {
	*out = *in
	if in.UserInfoEndpointEnabled != nil {
		in, out := &in.UserInfoEndpointEnabled, &out.UserInfoEndpointEnabled
		*out = new(bool)
		**out = **in
	}
	if in.HostNameVerificationEnabled != nil {
		in, out := &in.HostNameVerificationEnabled, &out.HostNameVerificationEnabled
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OidcClient.
func (in *OidcClient) DeepCopy() *OidcClient {
	if in == nil {
		return nil
	}
	out := new(OidcClient)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OpenLibertyApplication) DeepCopyInto(out *OpenLibertyApplication) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OpenLibertyApplication.
func (in *OpenLibertyApplication) DeepCopy() *OpenLibertyApplication {
	if in == nil {
		return nil
	}
	out := new(OpenLibertyApplication)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *OpenLibertyApplication) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OpenLibertyApplicationAffinity) DeepCopyInto(out *OpenLibertyApplicationAffinity) {
	*out = *in
	if in.NodeAffinity != nil {
		in, out := &in.NodeAffinity, &out.NodeAffinity
		*out = new(v1.NodeAffinity)
		(*in).DeepCopyInto(*out)
	}
	if in.PodAffinity != nil {
		in, out := &in.PodAffinity, &out.PodAffinity
		*out = new(v1.PodAffinity)
		(*in).DeepCopyInto(*out)
	}
	if in.PodAntiAffinity != nil {
		in, out := &in.PodAntiAffinity, &out.PodAntiAffinity
		*out = new(v1.PodAntiAffinity)
		(*in).DeepCopyInto(*out)
	}
	if in.NodeAffinityLabels != nil {
		in, out := &in.NodeAffinityLabels, &out.NodeAffinityLabels
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Architecture != nil {
		in, out := &in.Architecture, &out.Architecture
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OpenLibertyApplicationAffinity.
func (in *OpenLibertyApplicationAffinity) DeepCopy() *OpenLibertyApplicationAffinity {
	if in == nil {
		return nil
	}
	out := new(OpenLibertyApplicationAffinity)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OpenLibertyApplicationAutoScaling) DeepCopyInto(out *OpenLibertyApplicationAutoScaling) {
	*out = *in
	if in.MinReplicas != nil {
		in, out := &in.MinReplicas, &out.MinReplicas
		*out = new(int32)
		**out = **in
	}
	if in.TargetCPUUtilizationPercentage != nil {
		in, out := &in.TargetCPUUtilizationPercentage, &out.TargetCPUUtilizationPercentage
		*out = new(int32)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OpenLibertyApplicationAutoScaling.
func (in *OpenLibertyApplicationAutoScaling) DeepCopy() *OpenLibertyApplicationAutoScaling {
	if in == nil {
		return nil
	}
	out := new(OpenLibertyApplicationAutoScaling)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OpenLibertyApplicationDeployment) DeepCopyInto(out *OpenLibertyApplicationDeployment) {
	*out = *in
	if in.UpdateStrategy != nil {
		in, out := &in.UpdateStrategy, &out.UpdateStrategy
		*out = new(appsv1.DeploymentStrategy)
		(*in).DeepCopyInto(*out)
	}
	if in.Annotations != nil {
		in, out := &in.Annotations, &out.Annotations
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OpenLibertyApplicationDeployment.
func (in *OpenLibertyApplicationDeployment) DeepCopy() *OpenLibertyApplicationDeployment {
	if in == nil {
		return nil
	}
	out := new(OpenLibertyApplicationDeployment)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OpenLibertyApplicationList) DeepCopyInto(out *OpenLibertyApplicationList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]OpenLibertyApplication, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OpenLibertyApplicationList.
func (in *OpenLibertyApplicationList) DeepCopy() *OpenLibertyApplicationList {
	if in == nil {
		return nil
	}
	out := new(OpenLibertyApplicationList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *OpenLibertyApplicationList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OpenLibertyApplicationMonitoring) DeepCopyInto(out *OpenLibertyApplicationMonitoring) {
	*out = *in
	if in.Labels != nil {
		in, out := &in.Labels, &out.Labels
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Endpoints != nil {
		in, out := &in.Endpoints, &out.Endpoints
		*out = make([]monitoringv1.Endpoint, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OpenLibertyApplicationMonitoring.
func (in *OpenLibertyApplicationMonitoring) DeepCopy() *OpenLibertyApplicationMonitoring {
	if in == nil {
		return nil
	}
	out := new(OpenLibertyApplicationMonitoring)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OpenLibertyApplicationProbes) DeepCopyInto(out *OpenLibertyApplicationProbes) {
	*out = *in
	if in.Liveness != nil {
		in, out := &in.Liveness, &out.Liveness
		*out = new(v1.Probe)
		(*in).DeepCopyInto(*out)
	}
	if in.Readiness != nil {
		in, out := &in.Readiness, &out.Readiness
		*out = new(v1.Probe)
		(*in).DeepCopyInto(*out)
	}
	if in.Startup != nil {
		in, out := &in.Startup, &out.Startup
		*out = new(v1.Probe)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OpenLibertyApplicationProbes.
func (in *OpenLibertyApplicationProbes) DeepCopy() *OpenLibertyApplicationProbes {
	if in == nil {
		return nil
	}
	out := new(OpenLibertyApplicationProbes)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OpenLibertyApplicationRoute) DeepCopyInto(out *OpenLibertyApplicationRoute) {
	*out = *in
	if in.Annotations != nil {
		in, out := &in.Annotations, &out.Annotations
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.CertificateSecretRef != nil {
		in, out := &in.CertificateSecretRef, &out.CertificateSecretRef
		*out = new(string)
		**out = **in
	}
	if in.Termination != nil {
		in, out := &in.Termination, &out.Termination
		*out = new(routev1.TLSTerminationType)
		**out = **in
	}
	if in.InsecureEdgeTerminationPolicy != nil {
		in, out := &in.InsecureEdgeTerminationPolicy, &out.InsecureEdgeTerminationPolicy
		*out = new(routev1.InsecureEdgeTerminationPolicyType)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OpenLibertyApplicationRoute.
func (in *OpenLibertyApplicationRoute) DeepCopy() *OpenLibertyApplicationRoute {
	if in == nil {
		return nil
	}
	out := new(OpenLibertyApplicationRoute)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OpenLibertyApplicationSSO) DeepCopyInto(out *OpenLibertyApplicationSSO) {
	*out = *in
	if in.OIDC != nil {
		in, out := &in.OIDC, &out.OIDC
		*out = make([]OidcClient, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Oauth2 != nil {
		in, out := &in.Oauth2, &out.Oauth2
		*out = make([]OAuth2Client, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Github != nil {
		in, out := &in.Github, &out.Github
		*out = new(GithubLogin)
		**out = **in
	}
	if in.MapToUserRegistry != nil {
		in, out := &in.MapToUserRegistry, &out.MapToUserRegistry
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OpenLibertyApplicationSSO.
func (in *OpenLibertyApplicationSSO) DeepCopy() *OpenLibertyApplicationSSO {
	if in == nil {
		return nil
	}
	out := new(OpenLibertyApplicationSSO)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OpenLibertyApplicationService) DeepCopyInto(out *OpenLibertyApplicationService) {
	*out = *in
	if in.Type != nil {
		in, out := &in.Type, &out.Type
		*out = new(v1.ServiceType)
		**out = **in
	}
	if in.NodePort != nil {
		in, out := &in.NodePort, &out.NodePort
		*out = new(int32)
		**out = **in
	}
	if in.Annotations != nil {
		in, out := &in.Annotations, &out.Annotations
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.TargetPort != nil {
		in, out := &in.TargetPort, &out.TargetPort
		*out = new(int32)
		**out = **in
	}
	if in.CertificateSecretRef != nil {
		in, out := &in.CertificateSecretRef, &out.CertificateSecretRef
		*out = new(string)
		**out = **in
	}
	if in.Ports != nil {
		in, out := &in.Ports, &out.Ports
		*out = make([]v1.ServicePort, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Bindable != nil {
		in, out := &in.Bindable, &out.Bindable
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OpenLibertyApplicationService.
func (in *OpenLibertyApplicationService) DeepCopy() *OpenLibertyApplicationService {
	if in == nil {
		return nil
	}
	out := new(OpenLibertyApplicationService)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OpenLibertyApplicationServiceability) DeepCopyInto(out *OpenLibertyApplicationServiceability) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OpenLibertyApplicationServiceability.
func (in *OpenLibertyApplicationServiceability) DeepCopy() *OpenLibertyApplicationServiceability {
	if in == nil {
		return nil
	}
	out := new(OpenLibertyApplicationServiceability)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OpenLibertyApplicationSpec) DeepCopyInto(out *OpenLibertyApplicationSpec) {
	*out = *in
	if in.PullPolicy != nil {
		in, out := &in.PullPolicy, &out.PullPolicy
		*out = new(v1.PullPolicy)
		**out = **in
	}
	if in.PullSecret != nil {
		in, out := &in.PullSecret, &out.PullSecret
		*out = new(string)
		**out = **in
	}
	if in.ServiceAccountName != nil {
		in, out := &in.ServiceAccountName, &out.ServiceAccountName
		*out = new(string)
		**out = **in
	}
	if in.CreateKnativeService != nil {
		in, out := &in.CreateKnativeService, &out.CreateKnativeService
		*out = new(bool)
		**out = **in
	}
	if in.Expose != nil {
		in, out := &in.Expose, &out.Expose
		*out = new(bool)
		**out = **in
	}
	if in.Replicas != nil {
		in, out := &in.Replicas, &out.Replicas
		*out = new(int32)
		**out = **in
	}
	if in.Autoscaling != nil {
		in, out := &in.Autoscaling, &out.Autoscaling
		*out = new(OpenLibertyApplicationAutoScaling)
		(*in).DeepCopyInto(*out)
	}
	if in.Resources != nil {
		in, out := &in.Resources, &out.Resources
		*out = new(v1.ResourceRequirements)
		(*in).DeepCopyInto(*out)
	}
	if in.Probes != nil {
		in, out := &in.Probes, &out.Probes
		*out = new(OpenLibertyApplicationProbes)
		(*in).DeepCopyInto(*out)
	}
	if in.Deployment != nil {
		in, out := &in.Deployment, &out.Deployment
		*out = new(OpenLibertyApplicationDeployment)
		(*in).DeepCopyInto(*out)
	}
	if in.StatefulSet != nil {
		in, out := &in.StatefulSet, &out.StatefulSet
		*out = new(OpenLibertyApplicationStatefulSet)
		(*in).DeepCopyInto(*out)
	}
	if in.Service != nil {
		in, out := &in.Service, &out.Service
		*out = new(OpenLibertyApplicationService)
		(*in).DeepCopyInto(*out)
	}
	if in.Route != nil {
		in, out := &in.Route, &out.Route
		*out = new(OpenLibertyApplicationRoute)
		(*in).DeepCopyInto(*out)
	}
	if in.Serviceability != nil {
		in, out := &in.Serviceability, &out.Serviceability
		*out = new(OpenLibertyApplicationServiceability)
		**out = **in
	}
	if in.SSO != nil {
		in, out := &in.SSO, &out.SSO
		*out = new(OpenLibertyApplicationSSO)
		(*in).DeepCopyInto(*out)
	}
	if in.Monitoring != nil {
		in, out := &in.Monitoring, &out.Monitoring
		*out = new(OpenLibertyApplicationMonitoring)
		(*in).DeepCopyInto(*out)
	}
	if in.Env != nil {
		in, out := &in.Env, &out.Env
		*out = make([]v1.EnvVar, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.EnvFrom != nil {
		in, out := &in.EnvFrom, &out.EnvFrom
		*out = make([]v1.EnvFromSource, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Volumes != nil {
		in, out := &in.Volumes, &out.Volumes
		*out = make([]v1.Volume, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.VolumeMounts != nil {
		in, out := &in.VolumeMounts, &out.VolumeMounts
		*out = make([]v1.VolumeMount, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.InitContainers != nil {
		in, out := &in.InitContainers, &out.InitContainers
		*out = make([]v1.Container, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.SidecarContainers != nil {
		in, out := &in.SidecarContainers, &out.SidecarContainers
		*out = make([]v1.Container, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Affinity != nil {
		in, out := &in.Affinity, &out.Affinity
		*out = new(OpenLibertyApplicationAffinity)
		(*in).DeepCopyInto(*out)
	}
	if in.SecurityContext != nil {
		in, out := &in.SecurityContext, &out.SecurityContext
		*out = new(v1.SecurityContext)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OpenLibertyApplicationSpec.
func (in *OpenLibertyApplicationSpec) DeepCopy() *OpenLibertyApplicationSpec {
	if in == nil {
		return nil
	}
	out := new(OpenLibertyApplicationSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OpenLibertyApplicationStatefulSet) DeepCopyInto(out *OpenLibertyApplicationStatefulSet) {
	*out = *in
	if in.UpdateStrategy != nil {
		in, out := &in.UpdateStrategy, &out.UpdateStrategy
		*out = new(appsv1.StatefulSetUpdateStrategy)
		(*in).DeepCopyInto(*out)
	}
	if in.Storage != nil {
		in, out := &in.Storage, &out.Storage
		*out = new(OpenLibertyApplicationStorage)
		(*in).DeepCopyInto(*out)
	}
	if in.Annotations != nil {
		in, out := &in.Annotations, &out.Annotations
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OpenLibertyApplicationStatefulSet.
func (in *OpenLibertyApplicationStatefulSet) DeepCopy() *OpenLibertyApplicationStatefulSet {
	if in == nil {
		return nil
	}
	out := new(OpenLibertyApplicationStatefulSet)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OpenLibertyApplicationStatus) DeepCopyInto(out *OpenLibertyApplicationStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]StatusCondition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.RouteAvailable != nil {
		in, out := &in.RouteAvailable, &out.RouteAvailable
		*out = new(bool)
		**out = **in
	}
	if in.Binding != nil {
		in, out := &in.Binding, &out.Binding
		*out = new(v1.LocalObjectReference)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OpenLibertyApplicationStatus.
func (in *OpenLibertyApplicationStatus) DeepCopy() *OpenLibertyApplicationStatus {
	if in == nil {
		return nil
	}
	out := new(OpenLibertyApplicationStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OpenLibertyApplicationStorage) DeepCopyInto(out *OpenLibertyApplicationStorage) {
	*out = *in
	if in.VolumeClaimTemplate != nil {
		in, out := &in.VolumeClaimTemplate, &out.VolumeClaimTemplate
		*out = new(v1.PersistentVolumeClaim)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OpenLibertyApplicationStorage.
func (in *OpenLibertyApplicationStorage) DeepCopy() *OpenLibertyApplicationStorage {
	if in == nil {
		return nil
	}
	out := new(OpenLibertyApplicationStorage)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OpenLibertyDump) DeepCopyInto(out *OpenLibertyDump) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OpenLibertyDump.
func (in *OpenLibertyDump) DeepCopy() *OpenLibertyDump {
	if in == nil {
		return nil
	}
	out := new(OpenLibertyDump)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *OpenLibertyDump) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OpenLibertyDumpList) DeepCopyInto(out *OpenLibertyDumpList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]OpenLibertyDump, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OpenLibertyDumpList.
func (in *OpenLibertyDumpList) DeepCopy() *OpenLibertyDumpList {
	if in == nil {
		return nil
	}
	out := new(OpenLibertyDumpList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *OpenLibertyDumpList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OpenLibertyDumpSpec) DeepCopyInto(out *OpenLibertyDumpSpec) {
	*out = *in
	if in.Include != nil {
		in, out := &in.Include, &out.Include
		*out = make([]OpenLibertyDumpInclude, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OpenLibertyDumpSpec.
func (in *OpenLibertyDumpSpec) DeepCopy() *OpenLibertyDumpSpec {
	if in == nil {
		return nil
	}
	out := new(OpenLibertyDumpSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OpenLibertyDumpStatus) DeepCopyInto(out *OpenLibertyDumpStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]OperationStatusCondition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OpenLibertyDumpStatus.
func (in *OpenLibertyDumpStatus) DeepCopy() *OpenLibertyDumpStatus {
	if in == nil {
		return nil
	}
	out := new(OpenLibertyDumpStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OpenLibertyTrace) DeepCopyInto(out *OpenLibertyTrace) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OpenLibertyTrace.
func (in *OpenLibertyTrace) DeepCopy() *OpenLibertyTrace {
	if in == nil {
		return nil
	}
	out := new(OpenLibertyTrace)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *OpenLibertyTrace) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OpenLibertyTraceList) DeepCopyInto(out *OpenLibertyTraceList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]OpenLibertyTrace, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OpenLibertyTraceList.
func (in *OpenLibertyTraceList) DeepCopy() *OpenLibertyTraceList {
	if in == nil {
		return nil
	}
	out := new(OpenLibertyTraceList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *OpenLibertyTraceList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OpenLibertyTraceSpec) DeepCopyInto(out *OpenLibertyTraceSpec) {
	*out = *in
	if in.MaxFileSize != nil {
		in, out := &in.MaxFileSize, &out.MaxFileSize
		*out = new(int32)
		**out = **in
	}
	if in.MaxFiles != nil {
		in, out := &in.MaxFiles, &out.MaxFiles
		*out = new(int32)
		**out = **in
	}
	if in.Disable != nil {
		in, out := &in.Disable, &out.Disable
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OpenLibertyTraceSpec.
func (in *OpenLibertyTraceSpec) DeepCopy() *OpenLibertyTraceSpec {
	if in == nil {
		return nil
	}
	out := new(OpenLibertyTraceSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OpenLibertyTraceStatus) DeepCopyInto(out *OpenLibertyTraceStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]OperationStatusCondition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	out.OperatedResource = in.OperatedResource
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OpenLibertyTraceStatus.
func (in *OpenLibertyTraceStatus) DeepCopy() *OpenLibertyTraceStatus {
	if in == nil {
		return nil
	}
	out := new(OpenLibertyTraceStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OperatedResource) DeepCopyInto(out *OperatedResource) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OperatedResource.
func (in *OperatedResource) DeepCopy() *OperatedResource {
	if in == nil {
		return nil
	}
	out := new(OperatedResource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OperationStatusCondition) DeepCopyInto(out *OperationStatusCondition) {
	*out = *in
	if in.LastTransitionTime != nil {
		in, out := &in.LastTransitionTime, &out.LastTransitionTime
		*out = (*in).DeepCopy()
	}
	in.LastUpdateTime.DeepCopyInto(&out.LastUpdateTime)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OperationStatusCondition.
func (in *OperationStatusCondition) DeepCopy() *OperationStatusCondition {
	if in == nil {
		return nil
	}
	out := new(OperationStatusCondition)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StatusCondition) DeepCopyInto(out *StatusCondition) {
	*out = *in
	if in.LastTransitionTime != nil {
		in, out := &in.LastTransitionTime, &out.LastTransitionTime
		*out = (*in).DeepCopy()
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StatusCondition.
func (in *StatusCondition) DeepCopy() *StatusCondition {
	if in == nil {
		return nil
	}
	out := new(StatusCondition)
	in.DeepCopyInto(out)
	return out
}
