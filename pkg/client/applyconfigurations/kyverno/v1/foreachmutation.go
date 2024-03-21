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
	v1 "github.com/kyverno/kyverno/api/kyverno/v1"
	apiextensionsv1 "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"
)

// ForEachMutationApplyConfiguration represents an declarative configuration of the ForEachMutation type for use
// with apply.
type ForEachMutationApplyConfiguration struct {
	List                   *string                             `json:"list,omitempty"`
	Order                  *v1.ForeachOrder                    `json:"order,omitempty"`
	Context                []ContextEntryApplyConfiguration    `json:"context,omitempty"`
	AnyAllConditions       *AnyAllConditionsApplyConfiguration `json:"preconditions,omitempty"`
	RawPatchStrategicMerge *apiextensionsv1.JSON               `json:"patchStrategicMerge,omitempty"`
	PatchesJSON6902        *string                             `json:"patchesJson6902,omitempty"`
	ForEachMutation        *apiextensionsv1.JSON               `json:"foreach,omitempty"`
}

// ForEachMutationApplyConfiguration constructs an declarative configuration of the ForEachMutation type for use with
// apply.
func ForEachMutation() *ForEachMutationApplyConfiguration {
	return &ForEachMutationApplyConfiguration{}
}

// WithList sets the List field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the List field is set to the value of the last call.
func (b *ForEachMutationApplyConfiguration) WithList(value string) *ForEachMutationApplyConfiguration {
	b.List = &value
	return b
}

// WithOrder sets the Order field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Order field is set to the value of the last call.
func (b *ForEachMutationApplyConfiguration) WithOrder(value v1.ForeachOrder) *ForEachMutationApplyConfiguration {
	b.Order = &value
	return b
}

// WithContext adds the given value to the Context field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the Context field.
func (b *ForEachMutationApplyConfiguration) WithContext(values ...*ContextEntryApplyConfiguration) *ForEachMutationApplyConfiguration {
	for i := range values {
		if values[i] == nil {
			panic("nil value passed to WithContext")
		}
		b.Context = append(b.Context, *values[i])
	}
	return b
}

// WithAnyAllConditions sets the AnyAllConditions field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the AnyAllConditions field is set to the value of the last call.
func (b *ForEachMutationApplyConfiguration) WithAnyAllConditions(value *AnyAllConditionsApplyConfiguration) *ForEachMutationApplyConfiguration {
	b.AnyAllConditions = value
	return b
}

// WithRawPatchStrategicMerge sets the RawPatchStrategicMerge field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the RawPatchStrategicMerge field is set to the value of the last call.
func (b *ForEachMutationApplyConfiguration) WithRawPatchStrategicMerge(value apiextensionsv1.JSON) *ForEachMutationApplyConfiguration {
	b.RawPatchStrategicMerge = &value
	return b
}

// WithPatchesJSON6902 sets the PatchesJSON6902 field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the PatchesJSON6902 field is set to the value of the last call.
func (b *ForEachMutationApplyConfiguration) WithPatchesJSON6902(value string) *ForEachMutationApplyConfiguration {
	b.PatchesJSON6902 = &value
	return b
}

// WithForEachMutation sets the ForEachMutation field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the ForEachMutation field is set to the value of the last call.
func (b *ForEachMutationApplyConfiguration) WithForEachMutation(value apiextensionsv1.JSON) *ForEachMutationApplyConfiguration {
	b.ForEachMutation = &value
	return b
}
