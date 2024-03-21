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

package v1beta1

import (
	v1 "k8s.io/api/authentication/v1"
)

// RequestInfoApplyConfiguration represents an declarative configuration of the RequestInfo type for use
// with apply.
type RequestInfoApplyConfiguration struct {
	Roles             []string     `json:"roles,omitempty"`
	ClusterRoles      []string     `json:"clusterRoles,omitempty"`
	AdmissionUserInfo *v1.UserInfo `json:"userInfo,omitempty"`
}

// RequestInfoApplyConfiguration constructs an declarative configuration of the RequestInfo type for use with
// apply.
func RequestInfo() *RequestInfoApplyConfiguration {
	return &RequestInfoApplyConfiguration{}
}

// WithRoles adds the given value to the Roles field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the Roles field.
func (b *RequestInfoApplyConfiguration) WithRoles(values ...string) *RequestInfoApplyConfiguration {
	for i := range values {
		b.Roles = append(b.Roles, values[i])
	}
	return b
}

// WithClusterRoles adds the given value to the ClusterRoles field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the ClusterRoles field.
func (b *RequestInfoApplyConfiguration) WithClusterRoles(values ...string) *RequestInfoApplyConfiguration {
	for i := range values {
		b.ClusterRoles = append(b.ClusterRoles, values[i])
	}
	return b
}

// WithAdmissionUserInfo sets the AdmissionUserInfo field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the AdmissionUserInfo field is set to the value of the last call.
func (b *RequestInfoApplyConfiguration) WithAdmissionUserInfo(value v1.UserInfo) *RequestInfoApplyConfiguration {
	b.AdmissionUserInfo = &value
	return b
}
