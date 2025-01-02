//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
Copyright 2025 The Kubernetes Authors.

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

package removefailedpods

import (
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	api "sigs.k8s.io/descheduler/pkg/api"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RemoveFailedPodsArgs) DeepCopyInto(out *RemoveFailedPodsArgs) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	if in.Namespaces != nil {
		in, out := &in.Namespaces, &out.Namespaces
		*out = new(api.Namespaces)
		(*in).DeepCopyInto(*out)
	}
	if in.LabelSelector != nil {
		in, out := &in.LabelSelector, &out.LabelSelector
		*out = new(v1.LabelSelector)
		(*in).DeepCopyInto(*out)
	}
	if in.ExcludeOwnerKinds != nil {
		in, out := &in.ExcludeOwnerKinds, &out.ExcludeOwnerKinds
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.MinPodLifetimeSeconds != nil {
		in, out := &in.MinPodLifetimeSeconds, &out.MinPodLifetimeSeconds
		*out = new(uint)
		**out = **in
	}
	if in.Reasons != nil {
		in, out := &in.Reasons, &out.Reasons
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.ExitCodes != nil {
		in, out := &in.ExitCodes, &out.ExitCodes
		*out = make([]int32, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RemoveFailedPodsArgs.
func (in *RemoveFailedPodsArgs) DeepCopy() *RemoveFailedPodsArgs {
	if in == nil {
		return nil
	}
	out := new(RemoveFailedPodsArgs)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *RemoveFailedPodsArgs) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}
