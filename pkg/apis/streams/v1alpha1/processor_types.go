/*
 * Copyright 2019 The original author or authors
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package v1alpha1

import (
	duckv1alpha1 "github.com/knative/pkg/apis/duck/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type Processor struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ProcessorSpec   `json:"spec"`
	Status ProcessorStatus `json:"status"`
}

type ProcessorSpec struct {
	// TODO use arrays for inputs/outputs
	Inputs   string `json:"inputs"`
	Outputs  string `json:"outputs"`
	Function string `json:"function"`
}

const (
	ProcessorConditionReady = duckv1alpha1.ConditionReady
	// TODO add aggregated streams ready status
	ProcessorConditionFunctionReady duckv1alpha1.ConditionType = "FunctionReady"
)

var processorCondSet = duckv1alpha1.NewLivingConditionSet(ProcessorConditionFunctionReady)

type ProcessorStatus struct {
	Conditions         duckv1alpha1.Conditions `json:"conditions,omitempty"`
	ObservedGeneration int64                   `json:"observedGeneration,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type ProcessorList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`

	Items []Processor `json:"items"`
}

func (*Processor) GetGroupVersionKind() schema.GroupVersionKind {
	return SchemeGroupVersion.WithKind("Processor")
}

func (status *ProcessorStatus) IsReady() bool {
	return processorCondSet.Manage(status).IsHappy()
}

func (status *ProcessorStatus) GetCondition(t duckv1alpha1.ConditionType) *duckv1alpha1.Condition {
	return processorCondSet.Manage(status).GetCondition(t)
}

func (status *ProcessorStatus) InitializeConditions() {
	processorCondSet.Manage(status).InitializeConditions()
}
