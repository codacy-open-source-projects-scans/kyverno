/*
Copyright The Kubernetes Authors.

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

// Code generated by applyconfiguration-gen. DO NOT EDIT.

package v1

import (
	api "k8s.io/pod-security-admission/api"
)

// PodSecurityApplyConfiguration represents an declarative configuration of the PodSecurity type for use
// with apply.
type PodSecurityApplyConfiguration struct {
	Level   *api.Level                              `json:"level,omitempty"`
	Version *string                                 `json:"version,omitempty"`
	Exclude []PodSecurityStandardApplyConfiguration `json:"exclude,omitempty"`
}

// PodSecurityApplyConfiguration constructs an declarative configuration of the PodSecurity type for use with
// apply.
func PodSecurity() *PodSecurityApplyConfiguration {
	return &PodSecurityApplyConfiguration{}
}

// WithLevel sets the Level field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Level field is set to the value of the last call.
func (b *PodSecurityApplyConfiguration) WithLevel(value api.Level) *PodSecurityApplyConfiguration {
	b.Level = &value
	return b
}

// WithVersion sets the Version field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Version field is set to the value of the last call.
func (b *PodSecurityApplyConfiguration) WithVersion(value string) *PodSecurityApplyConfiguration {
	b.Version = &value
	return b
}

// WithExclude adds the given value to the Exclude field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the Exclude field.
func (b *PodSecurityApplyConfiguration) WithExclude(values ...*PodSecurityStandardApplyConfiguration) *PodSecurityApplyConfiguration {
	for i := range values {
		if values[i] == nil {
			panic("nil value passed to WithExclude")
		}
		b.Exclude = append(b.Exclude, *values[i])
	}
	return b
}
