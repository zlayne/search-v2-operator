//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
Copyright 2022.

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

package v1alpha1

import (
	"k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DeploymentConfig) DeepCopyInto(out *DeploymentConfig) {
	*out = *in
	if in.Resources != nil {
		in, out := &in.Resources, &out.Resources
		*out = new(v1.ResourceRequirements)
		(*in).DeepCopyInto(*out)
	}
	if in.NodeSelector != nil {
		in, out := &in.NodeSelector, &out.NodeSelector
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.ProxyConfig != nil {
		in, out := &in.ProxyConfig, &out.ProxyConfig
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DeploymentConfig.
func (in *DeploymentConfig) DeepCopy() *DeploymentConfig {
	if in == nil {
		return nil
	}
	out := new(DeploymentConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FilterSpec) DeepCopyInto(out *FilterSpec) {
	*out = *in
	if in.AllowedResources != nil {
		in, out := &in.AllowedResources, &out.AllowedResources
		*out = make([]ResourceListSpec, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.DeniedResources != nil {
		in, out := &in.DeniedResources, &out.DeniedResources
		*out = make([]ResourceListSpec, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FilterSpec.
func (in *FilterSpec) DeepCopy() *FilterSpec {
	if in == nil {
		return nil
	}
	out := new(FilterSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ResourceListSpec) DeepCopyInto(out *ResourceListSpec) {
	*out = *in
	if in.APIGroups != nil {
		in, out := &in.APIGroups, &out.APIGroups
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Resources != nil {
		in, out := &in.Resources, &out.Resources
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	in.ClusterLabels.DeepCopyInto(&out.ClusterLabels)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ResourceListSpec.
func (in *ResourceListSpec) DeepCopy() *ResourceListSpec {
	if in == nil {
		return nil
	}
	out := new(ResourceListSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Search) DeepCopyInto(out *Search) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Search.
func (in *Search) DeepCopy() *Search {
	if in == nil {
		return nil
	}
	out := new(Search)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Search) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SearchDeployments) DeepCopyInto(out *SearchDeployments) {
	*out = *in
	in.Database.DeepCopyInto(&out.Database)
	in.Indexer.DeepCopyInto(&out.Indexer)
	in.Collector.DeepCopyInto(&out.Collector)
	in.API.DeepCopyInto(&out.API)
	in.RemoteCollector.DeepCopyInto(&out.RemoteCollector)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SearchDeployments.
func (in *SearchDeployments) DeepCopy() *SearchDeployments {
	if in == nil {
		return nil
	}
	out := new(SearchDeployments)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SearchList) DeepCopyInto(out *SearchList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Search, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SearchList.
func (in *SearchList) DeepCopy() *SearchList {
	if in == nil {
		return nil
	}
	out := new(SearchList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *SearchList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SearchSpec) DeepCopyInto(out *SearchSpec) {
	*out = *in
	in.DBStorage.DeepCopyInto(&out.DBStorage)
	in.Deployments.DeepCopyInto(&out.Deployments)
	in.AllowDenyResources.DeepCopyInto(&out.AllowDenyResources)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SearchSpec.
func (in *SearchSpec) DeepCopy() *SearchSpec {
	if in == nil {
		return nil
	}
	out := new(SearchSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SearchStatus) DeepCopyInto(out *SearchStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]metav1.Condition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SearchStatus.
func (in *SearchStatus) DeepCopy() *SearchStatus {
	if in == nil {
		return nil
	}
	out := new(SearchStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StorageSpec) DeepCopyInto(out *StorageSpec) {
	*out = *in
	if in.Size != nil {
		in, out := &in.Size, &out.Size
		x := (*in).DeepCopy()
		*out = &x
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StorageSpec.
func (in *StorageSpec) DeepCopy() *StorageSpec {
	if in == nil {
		return nil
	}
	out := new(StorageSpec)
	in.DeepCopyInto(out)
	return out
}
