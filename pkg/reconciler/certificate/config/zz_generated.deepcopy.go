// +build !ignore_autogenerated

/*
Copyright 2019 The Knative Authors

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

package config

import (
	v1alpha1 "github.com/jetstack/cert-manager/pkg/apis/certmanager/v1alpha1"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CertManagerConfig) DeepCopyInto(out *CertManagerConfig) {
	*out = *in
	if in.SolverConfig != nil {
		in, out := &in.SolverConfig, &out.SolverConfig
		*out = new(v1alpha1.SolverConfig)
		(*in).DeepCopyInto(*out)
	}
	if in.IssuerRef != nil {
		in, out := &in.IssuerRef, &out.IssuerRef
		*out = new(v1alpha1.ObjectReference)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CertManagerConfig.
func (in *CertManagerConfig) DeepCopy() *CertManagerConfig {
	if in == nil {
		return nil
	}
	out := new(CertManagerConfig)
	in.DeepCopyInto(out)
	return out
}
