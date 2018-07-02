// +build !ignore_autogenerated

/*
Copyright 2018 Giant Swarm GmbH.

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

// Code generated by deepcopy-gen. DO NOT EDIT.

package v1alpha1

import (
	net "net"

	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AWSConfig) DeepCopyInto(out *AWSConfig) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AWSConfig.
func (in *AWSConfig) DeepCopy() *AWSConfig {
	if in == nil {
		return nil
	}
	out := new(AWSConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AWSConfig) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AWSConfigList) DeepCopyInto(out *AWSConfigList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]AWSConfig, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AWSConfigList.
func (in *AWSConfigList) DeepCopy() *AWSConfigList {
	if in == nil {
		return nil
	}
	out := new(AWSConfigList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AWSConfigList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AWSConfigSpec) DeepCopyInto(out *AWSConfigSpec) {
	*out = *in
	in.Cluster.DeepCopyInto(&out.Cluster)
	in.AWS.DeepCopyInto(&out.AWS)
	out.VersionBundle = in.VersionBundle
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AWSConfigSpec.
func (in *AWSConfigSpec) DeepCopy() *AWSConfigSpec {
	if in == nil {
		return nil
	}
	out := new(AWSConfigSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AWSConfigSpecAWS) DeepCopyInto(out *AWSConfigSpecAWS) {
	*out = *in
	out.API = in.API
	out.CredentialSecret = in.CredentialSecret
	out.Etcd = in.Etcd
	out.HostedZones = in.HostedZones
	out.Ingress = in.Ingress
	if in.Masters != nil {
		in, out := &in.Masters, &out.Masters
		*out = make([]AWSConfigSpecAWSNode, len(*in))
		copy(*out, *in)
	}
	in.VPC.DeepCopyInto(&out.VPC)
	if in.Workers != nil {
		in, out := &in.Workers, &out.Workers
		*out = make([]AWSConfigSpecAWSNode, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AWSConfigSpecAWS.
func (in *AWSConfigSpecAWS) DeepCopy() *AWSConfigSpecAWS {
	if in == nil {
		return nil
	}
	out := new(AWSConfigSpecAWS)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AWSConfigSpecAWSAPI) DeepCopyInto(out *AWSConfigSpecAWSAPI) {
	*out = *in
	out.ELB = in.ELB
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AWSConfigSpecAWSAPI.
func (in *AWSConfigSpecAWSAPI) DeepCopy() *AWSConfigSpecAWSAPI {
	if in == nil {
		return nil
	}
	out := new(AWSConfigSpecAWSAPI)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AWSConfigSpecAWSAPIELB) DeepCopyInto(out *AWSConfigSpecAWSAPIELB) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AWSConfigSpecAWSAPIELB.
func (in *AWSConfigSpecAWSAPIELB) DeepCopy() *AWSConfigSpecAWSAPIELB {
	if in == nil {
		return nil
	}
	out := new(AWSConfigSpecAWSAPIELB)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AWSConfigSpecAWSEtcd) DeepCopyInto(out *AWSConfigSpecAWSEtcd) {
	*out = *in
	out.ELB = in.ELB
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AWSConfigSpecAWSEtcd.
func (in *AWSConfigSpecAWSEtcd) DeepCopy() *AWSConfigSpecAWSEtcd {
	if in == nil {
		return nil
	}
	out := new(AWSConfigSpecAWSEtcd)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AWSConfigSpecAWSEtcdELB) DeepCopyInto(out *AWSConfigSpecAWSEtcdELB) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AWSConfigSpecAWSEtcdELB.
func (in *AWSConfigSpecAWSEtcdELB) DeepCopy() *AWSConfigSpecAWSEtcdELB {
	if in == nil {
		return nil
	}
	out := new(AWSConfigSpecAWSEtcdELB)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AWSConfigSpecAWSHostedZones) DeepCopyInto(out *AWSConfigSpecAWSHostedZones) {
	*out = *in
	out.API = in.API
	out.Etcd = in.Etcd
	out.Ingress = in.Ingress
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AWSConfigSpecAWSHostedZones.
func (in *AWSConfigSpecAWSHostedZones) DeepCopy() *AWSConfigSpecAWSHostedZones {
	if in == nil {
		return nil
	}
	out := new(AWSConfigSpecAWSHostedZones)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AWSConfigSpecAWSHostedZonesZone) DeepCopyInto(out *AWSConfigSpecAWSHostedZonesZone) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AWSConfigSpecAWSHostedZonesZone.
func (in *AWSConfigSpecAWSHostedZonesZone) DeepCopy() *AWSConfigSpecAWSHostedZonesZone {
	if in == nil {
		return nil
	}
	out := new(AWSConfigSpecAWSHostedZonesZone)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AWSConfigSpecAWSIngress) DeepCopyInto(out *AWSConfigSpecAWSIngress) {
	*out = *in
	out.ELB = in.ELB
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AWSConfigSpecAWSIngress.
func (in *AWSConfigSpecAWSIngress) DeepCopy() *AWSConfigSpecAWSIngress {
	if in == nil {
		return nil
	}
	out := new(AWSConfigSpecAWSIngress)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AWSConfigSpecAWSIngressELB) DeepCopyInto(out *AWSConfigSpecAWSIngressELB) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AWSConfigSpecAWSIngressELB.
func (in *AWSConfigSpecAWSIngressELB) DeepCopy() *AWSConfigSpecAWSIngressELB {
	if in == nil {
		return nil
	}
	out := new(AWSConfigSpecAWSIngressELB)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AWSConfigSpecAWSNode) DeepCopyInto(out *AWSConfigSpecAWSNode) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AWSConfigSpecAWSNode.
func (in *AWSConfigSpecAWSNode) DeepCopy() *AWSConfigSpecAWSNode {
	if in == nil {
		return nil
	}
	out := new(AWSConfigSpecAWSNode)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AWSConfigSpecAWSVPC) DeepCopyInto(out *AWSConfigSpecAWSVPC) {
	*out = *in
	if in.RouteTableNames != nil {
		in, out := &in.RouteTableNames, &out.RouteTableNames
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AWSConfigSpecAWSVPC.
func (in *AWSConfigSpecAWSVPC) DeepCopy() *AWSConfigSpecAWSVPC {
	if in == nil {
		return nil
	}
	out := new(AWSConfigSpecAWSVPC)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AWSConfigSpecVersionBundle) DeepCopyInto(out *AWSConfigSpecVersionBundle) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AWSConfigSpecVersionBundle.
func (in *AWSConfigSpecVersionBundle) DeepCopy() *AWSConfigSpecVersionBundle {
	if in == nil {
		return nil
	}
	out := new(AWSConfigSpecVersionBundle)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AWSConfigStatus) DeepCopyInto(out *AWSConfigStatus) {
	*out = *in
	in.Cluster.DeepCopyInto(&out.Cluster)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AWSConfigStatus.
func (in *AWSConfigStatus) DeepCopy() *AWSConfigStatus {
	if in == nil {
		return nil
	}
	out := new(AWSConfigStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AzureConfig) DeepCopyInto(out *AzureConfig) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AzureConfig.
func (in *AzureConfig) DeepCopy() *AzureConfig {
	if in == nil {
		return nil
	}
	out := new(AzureConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AzureConfig) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AzureConfigList) DeepCopyInto(out *AzureConfigList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]AzureConfig, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AzureConfigList.
func (in *AzureConfigList) DeepCopy() *AzureConfigList {
	if in == nil {
		return nil
	}
	out := new(AzureConfigList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AzureConfigList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AzureConfigSpec) DeepCopyInto(out *AzureConfigSpec) {
	*out = *in
	in.Cluster.DeepCopyInto(&out.Cluster)
	in.Azure.DeepCopyInto(&out.Azure)
	out.VersionBundle = in.VersionBundle
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AzureConfigSpec.
func (in *AzureConfigSpec) DeepCopy() *AzureConfigSpec {
	if in == nil {
		return nil
	}
	out := new(AzureConfigSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AzureConfigSpecAzure) DeepCopyInto(out *AzureConfigSpecAzure) {
	*out = *in
	out.CredentialSecret = in.CredentialSecret
	out.DNSZones = in.DNSZones
	out.VirtualNetwork = in.VirtualNetwork
	if in.Masters != nil {
		in, out := &in.Masters, &out.Masters
		*out = make([]AzureConfigSpecAzureNode, len(*in))
		copy(*out, *in)
	}
	if in.Workers != nil {
		in, out := &in.Workers, &out.Workers
		*out = make([]AzureConfigSpecAzureNode, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AzureConfigSpecAzure.
func (in *AzureConfigSpecAzure) DeepCopy() *AzureConfigSpecAzure {
	if in == nil {
		return nil
	}
	out := new(AzureConfigSpecAzure)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AzureConfigSpecAzureDNSZones) DeepCopyInto(out *AzureConfigSpecAzureDNSZones) {
	*out = *in
	out.API = in.API
	out.Etcd = in.Etcd
	out.Ingress = in.Ingress
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AzureConfigSpecAzureDNSZones.
func (in *AzureConfigSpecAzureDNSZones) DeepCopy() *AzureConfigSpecAzureDNSZones {
	if in == nil {
		return nil
	}
	out := new(AzureConfigSpecAzureDNSZones)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AzureConfigSpecAzureDNSZonesDNSZone) DeepCopyInto(out *AzureConfigSpecAzureDNSZonesDNSZone) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AzureConfigSpecAzureDNSZonesDNSZone.
func (in *AzureConfigSpecAzureDNSZonesDNSZone) DeepCopy() *AzureConfigSpecAzureDNSZonesDNSZone {
	if in == nil {
		return nil
	}
	out := new(AzureConfigSpecAzureDNSZonesDNSZone)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AzureConfigSpecAzureNode) DeepCopyInto(out *AzureConfigSpecAzureNode) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AzureConfigSpecAzureNode.
func (in *AzureConfigSpecAzureNode) DeepCopy() *AzureConfigSpecAzureNode {
	if in == nil {
		return nil
	}
	out := new(AzureConfigSpecAzureNode)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AzureConfigSpecAzureVirtualNetwork) DeepCopyInto(out *AzureConfigSpecAzureVirtualNetwork) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AzureConfigSpecAzureVirtualNetwork.
func (in *AzureConfigSpecAzureVirtualNetwork) DeepCopy() *AzureConfigSpecAzureVirtualNetwork {
	if in == nil {
		return nil
	}
	out := new(AzureConfigSpecAzureVirtualNetwork)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AzureConfigSpecVersionBundle) DeepCopyInto(out *AzureConfigSpecVersionBundle) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AzureConfigSpecVersionBundle.
func (in *AzureConfigSpecVersionBundle) DeepCopy() *AzureConfigSpecVersionBundle {
	if in == nil {
		return nil
	}
	out := new(AzureConfigSpecVersionBundle)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AzureConfigStatus) DeepCopyInto(out *AzureConfigStatus) {
	*out = *in
	in.Cluster.DeepCopyInto(&out.Cluster)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AzureConfigStatus.
func (in *AzureConfigStatus) DeepCopy() *AzureConfigStatus {
	if in == nil {
		return nil
	}
	out := new(AzureConfigStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Cluster) DeepCopyInto(out *Cluster) {
	*out = *in
	out.Calico = in.Calico
	out.Customer = in.Customer
	out.Docker = in.Docker
	out.Etcd = in.Etcd
	in.Kubernetes.DeepCopyInto(&out.Kubernetes)
	if in.Masters != nil {
		in, out := &in.Masters, &out.Masters
		*out = make([]ClusterNode, len(*in))
		copy(*out, *in)
	}
	out.Vault = in.Vault
	if in.Workers != nil {
		in, out := &in.Workers, &out.Workers
		*out = make([]ClusterNode, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Cluster.
func (in *Cluster) DeepCopy() *Cluster {
	if in == nil {
		return nil
	}
	out := new(Cluster)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterCalico) DeepCopyInto(out *ClusterCalico) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterCalico.
func (in *ClusterCalico) DeepCopy() *ClusterCalico {
	if in == nil {
		return nil
	}
	out := new(ClusterCalico)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterCustomer) DeepCopyInto(out *ClusterCustomer) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterCustomer.
func (in *ClusterCustomer) DeepCopy() *ClusterCustomer {
	if in == nil {
		return nil
	}
	out := new(ClusterCustomer)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterDocker) DeepCopyInto(out *ClusterDocker) {
	*out = *in
	out.Daemon = in.Daemon
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterDocker.
func (in *ClusterDocker) DeepCopy() *ClusterDocker {
	if in == nil {
		return nil
	}
	out := new(ClusterDocker)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterDockerDaemon) DeepCopyInto(out *ClusterDockerDaemon) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterDockerDaemon.
func (in *ClusterDockerDaemon) DeepCopy() *ClusterDockerDaemon {
	if in == nil {
		return nil
	}
	out := new(ClusterDockerDaemon)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterEtcd) DeepCopyInto(out *ClusterEtcd) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterEtcd.
func (in *ClusterEtcd) DeepCopy() *ClusterEtcd {
	if in == nil {
		return nil
	}
	out := new(ClusterEtcd)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterKubernetes) DeepCopyInto(out *ClusterKubernetes) {
	*out = *in
	in.API.DeepCopyInto(&out.API)
	in.DNS.DeepCopyInto(&out.DNS)
	out.Hyperkube = in.Hyperkube
	out.IngressController = in.IngressController
	out.Kubelet = in.Kubelet
	out.NetworkSetup = in.NetworkSetup
	in.SSH.DeepCopyInto(&out.SSH)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterKubernetes.
func (in *ClusterKubernetes) DeepCopy() *ClusterKubernetes {
	if in == nil {
		return nil
	}
	out := new(ClusterKubernetes)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterKubernetesAPI) DeepCopyInto(out *ClusterKubernetesAPI) {
	*out = *in
	if in.IP != nil {
		in, out := &in.IP, &out.IP
		*out = make(net.IP, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterKubernetesAPI.
func (in *ClusterKubernetesAPI) DeepCopy() *ClusterKubernetesAPI {
	if in == nil {
		return nil
	}
	out := new(ClusterKubernetesAPI)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterKubernetesDNS) DeepCopyInto(out *ClusterKubernetesDNS) {
	*out = *in
	if in.IP != nil {
		in, out := &in.IP, &out.IP
		*out = make(net.IP, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterKubernetesDNS.
func (in *ClusterKubernetesDNS) DeepCopy() *ClusterKubernetesDNS {
	if in == nil {
		return nil
	}
	out := new(ClusterKubernetesDNS)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterKubernetesHyperkube) DeepCopyInto(out *ClusterKubernetesHyperkube) {
	*out = *in
	out.Docker = in.Docker
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterKubernetesHyperkube.
func (in *ClusterKubernetesHyperkube) DeepCopy() *ClusterKubernetesHyperkube {
	if in == nil {
		return nil
	}
	out := new(ClusterKubernetesHyperkube)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterKubernetesHyperkubeDocker) DeepCopyInto(out *ClusterKubernetesHyperkubeDocker) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterKubernetesHyperkubeDocker.
func (in *ClusterKubernetesHyperkubeDocker) DeepCopy() *ClusterKubernetesHyperkubeDocker {
	if in == nil {
		return nil
	}
	out := new(ClusterKubernetesHyperkubeDocker)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterKubernetesIngressController) DeepCopyInto(out *ClusterKubernetesIngressController) {
	*out = *in
	out.Docker = in.Docker
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterKubernetesIngressController.
func (in *ClusterKubernetesIngressController) DeepCopy() *ClusterKubernetesIngressController {
	if in == nil {
		return nil
	}
	out := new(ClusterKubernetesIngressController)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterKubernetesIngressControllerDocker) DeepCopyInto(out *ClusterKubernetesIngressControllerDocker) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterKubernetesIngressControllerDocker.
func (in *ClusterKubernetesIngressControllerDocker) DeepCopy() *ClusterKubernetesIngressControllerDocker {
	if in == nil {
		return nil
	}
	out := new(ClusterKubernetesIngressControllerDocker)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterKubernetesKubelet) DeepCopyInto(out *ClusterKubernetesKubelet) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterKubernetesKubelet.
func (in *ClusterKubernetesKubelet) DeepCopy() *ClusterKubernetesKubelet {
	if in == nil {
		return nil
	}
	out := new(ClusterKubernetesKubelet)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterKubernetesNetworkSetup) DeepCopyInto(out *ClusterKubernetesNetworkSetup) {
	*out = *in
	out.Docker = in.Docker
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterKubernetesNetworkSetup.
func (in *ClusterKubernetesNetworkSetup) DeepCopy() *ClusterKubernetesNetworkSetup {
	if in == nil {
		return nil
	}
	out := new(ClusterKubernetesNetworkSetup)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterKubernetesNetworkSetupDocker) DeepCopyInto(out *ClusterKubernetesNetworkSetupDocker) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterKubernetesNetworkSetupDocker.
func (in *ClusterKubernetesNetworkSetupDocker) DeepCopy() *ClusterKubernetesNetworkSetupDocker {
	if in == nil {
		return nil
	}
	out := new(ClusterKubernetesNetworkSetupDocker)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterKubernetesSSH) DeepCopyInto(out *ClusterKubernetesSSH) {
	*out = *in
	if in.UserList != nil {
		in, out := &in.UserList, &out.UserList
		*out = make([]ClusterKubernetesSSHUser, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterKubernetesSSH.
func (in *ClusterKubernetesSSH) DeepCopy() *ClusterKubernetesSSH {
	if in == nil {
		return nil
	}
	out := new(ClusterKubernetesSSH)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterKubernetesSSHUser) DeepCopyInto(out *ClusterKubernetesSSHUser) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterKubernetesSSHUser.
func (in *ClusterKubernetesSSHUser) DeepCopy() *ClusterKubernetesSSHUser {
	if in == nil {
		return nil
	}
	out := new(ClusterKubernetesSSHUser)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterNode) DeepCopyInto(out *ClusterNode) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterNode.
func (in *ClusterNode) DeepCopy() *ClusterNode {
	if in == nil {
		return nil
	}
	out := new(ClusterNode)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterVault) DeepCopyInto(out *ClusterVault) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterVault.
func (in *ClusterVault) DeepCopy() *ClusterVault {
	if in == nil {
		return nil
	}
	out := new(ClusterVault)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CredentialSecret) DeepCopyInto(out *CredentialSecret) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CredentialSecret.
func (in *CredentialSecret) DeepCopy() *CredentialSecret {
	if in == nil {
		return nil
	}
	out := new(CredentialSecret)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KVMConfig) DeepCopyInto(out *KVMConfig) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KVMConfig.
func (in *KVMConfig) DeepCopy() *KVMConfig {
	if in == nil {
		return nil
	}
	out := new(KVMConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *KVMConfig) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KVMConfigList) DeepCopyInto(out *KVMConfigList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]KVMConfig, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KVMConfigList.
func (in *KVMConfigList) DeepCopy() *KVMConfigList {
	if in == nil {
		return nil
	}
	out := new(KVMConfigList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *KVMConfigList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KVMConfigSpec) DeepCopyInto(out *KVMConfigSpec) {
	*out = *in
	in.Cluster.DeepCopyInto(&out.Cluster)
	in.KVM.DeepCopyInto(&out.KVM)
	out.VersionBundle = in.VersionBundle
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KVMConfigSpec.
func (in *KVMConfigSpec) DeepCopy() *KVMConfigSpec {
	if in == nil {
		return nil
	}
	out := new(KVMConfigSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KVMConfigSpecKVM) DeepCopyInto(out *KVMConfigSpecKVM) {
	*out = *in
	out.EndpointUpdater = in.EndpointUpdater
	out.K8sKVM = in.K8sKVM
	if in.Masters != nil {
		in, out := &in.Masters, &out.Masters
		*out = make([]KVMConfigSpecKVMNode, len(*in))
		copy(*out, *in)
	}
	out.Network = in.Network
	out.NodeController = in.NodeController
	if in.Workers != nil {
		in, out := &in.Workers, &out.Workers
		*out = make([]KVMConfigSpecKVMNode, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KVMConfigSpecKVM.
func (in *KVMConfigSpecKVM) DeepCopy() *KVMConfigSpecKVM {
	if in == nil {
		return nil
	}
	out := new(KVMConfigSpecKVM)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KVMConfigSpecKVMEndpointUpdater) DeepCopyInto(out *KVMConfigSpecKVMEndpointUpdater) {
	*out = *in
	out.Docker = in.Docker
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KVMConfigSpecKVMEndpointUpdater.
func (in *KVMConfigSpecKVMEndpointUpdater) DeepCopy() *KVMConfigSpecKVMEndpointUpdater {
	if in == nil {
		return nil
	}
	out := new(KVMConfigSpecKVMEndpointUpdater)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KVMConfigSpecKVMEndpointUpdaterDocker) DeepCopyInto(out *KVMConfigSpecKVMEndpointUpdaterDocker) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KVMConfigSpecKVMEndpointUpdaterDocker.
func (in *KVMConfigSpecKVMEndpointUpdaterDocker) DeepCopy() *KVMConfigSpecKVMEndpointUpdaterDocker {
	if in == nil {
		return nil
	}
	out := new(KVMConfigSpecKVMEndpointUpdaterDocker)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KVMConfigSpecKVMK8sKVM) DeepCopyInto(out *KVMConfigSpecKVMK8sKVM) {
	*out = *in
	out.Docker = in.Docker
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KVMConfigSpecKVMK8sKVM.
func (in *KVMConfigSpecKVMK8sKVM) DeepCopy() *KVMConfigSpecKVMK8sKVM {
	if in == nil {
		return nil
	}
	out := new(KVMConfigSpecKVMK8sKVM)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KVMConfigSpecKVMK8sKVMDocker) DeepCopyInto(out *KVMConfigSpecKVMK8sKVMDocker) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KVMConfigSpecKVMK8sKVMDocker.
func (in *KVMConfigSpecKVMK8sKVMDocker) DeepCopy() *KVMConfigSpecKVMK8sKVMDocker {
	if in == nil {
		return nil
	}
	out := new(KVMConfigSpecKVMK8sKVMDocker)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KVMConfigSpecKVMNetwork) DeepCopyInto(out *KVMConfigSpecKVMNetwork) {
	*out = *in
	out.Flannel = in.Flannel
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KVMConfigSpecKVMNetwork.
func (in *KVMConfigSpecKVMNetwork) DeepCopy() *KVMConfigSpecKVMNetwork {
	if in == nil {
		return nil
	}
	out := new(KVMConfigSpecKVMNetwork)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KVMConfigSpecKVMNetworkFlannel) DeepCopyInto(out *KVMConfigSpecKVMNetworkFlannel) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KVMConfigSpecKVMNetworkFlannel.
func (in *KVMConfigSpecKVMNetworkFlannel) DeepCopy() *KVMConfigSpecKVMNetworkFlannel {
	if in == nil {
		return nil
	}
	out := new(KVMConfigSpecKVMNetworkFlannel)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KVMConfigSpecKVMNode) DeepCopyInto(out *KVMConfigSpecKVMNode) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KVMConfigSpecKVMNode.
func (in *KVMConfigSpecKVMNode) DeepCopy() *KVMConfigSpecKVMNode {
	if in == nil {
		return nil
	}
	out := new(KVMConfigSpecKVMNode)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KVMConfigSpecKVMNodeController) DeepCopyInto(out *KVMConfigSpecKVMNodeController) {
	*out = *in
	out.Docker = in.Docker
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KVMConfigSpecKVMNodeController.
func (in *KVMConfigSpecKVMNodeController) DeepCopy() *KVMConfigSpecKVMNodeController {
	if in == nil {
		return nil
	}
	out := new(KVMConfigSpecKVMNodeController)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KVMConfigSpecKVMNodeControllerDocker) DeepCopyInto(out *KVMConfigSpecKVMNodeControllerDocker) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KVMConfigSpecKVMNodeControllerDocker.
func (in *KVMConfigSpecKVMNodeControllerDocker) DeepCopy() *KVMConfigSpecKVMNodeControllerDocker {
	if in == nil {
		return nil
	}
	out := new(KVMConfigSpecKVMNodeControllerDocker)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KVMConfigSpecVersionBundle) DeepCopyInto(out *KVMConfigSpecVersionBundle) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KVMConfigSpecVersionBundle.
func (in *KVMConfigSpecVersionBundle) DeepCopy() *KVMConfigSpecVersionBundle {
	if in == nil {
		return nil
	}
	out := new(KVMConfigSpecVersionBundle)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KVMConfigStatus) DeepCopyInto(out *KVMConfigStatus) {
	*out = *in
	in.Cluster.DeepCopyInto(&out.Cluster)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KVMConfigStatus.
func (in *KVMConfigStatus) DeepCopy() *KVMConfigStatus {
	if in == nil {
		return nil
	}
	out := new(KVMConfigStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StatusCluster) DeepCopyInto(out *StatusCluster) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]StatusClusterCondition, len(*in))
		copy(*out, *in)
	}
	if in.Versions != nil {
		in, out := &in.Versions, &out.Versions
		*out = make([]StatusClusterVersion, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StatusCluster.
func (in *StatusCluster) DeepCopy() *StatusCluster {
	if in == nil {
		return nil
	}
	out := new(StatusCluster)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StatusClusterCondition) DeepCopyInto(out *StatusClusterCondition) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StatusClusterCondition.
func (in *StatusClusterCondition) DeepCopy() *StatusClusterCondition {
	if in == nil {
		return nil
	}
	out := new(StatusClusterCondition)
	in.DeepCopyInto(out)
	return out
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StatusClusterVersion.
func (in *StatusClusterVersion) DeepCopy() *StatusClusterVersion {
	if in == nil {
		return nil
	}
	out := new(StatusClusterVersion)
	in.DeepCopyInto(out)
	return out
}
