/*
Copyright 2018 The Knative Authors

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

package resources

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	runv1alpha1 "github.com/projectriff/system/pkg/apis/run/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

const (
	testRequestProcessorName      = "test-requestprocessor"
	testRequestProcessorNamespace = "test-requestprocessor-namespace"
	testConfigurationName         = "test-requestprocessor-0"
	testRouteName                 = "test-requestprocessor"
	testLabelKey                  = "test-label-key"
	testLabelValue                = "test-label-value"
	testServiceAccount            = "test-sa"
	testImage                     = "test-image"
	testTag                       = "test-tag"
)

func expectOwnerReferencesSetCorrectly(t *testing.T, ownerRefs []metav1.OwnerReference) {
	if got, want := len(ownerRefs), 1; got != want {
		t.Errorf("expected %d owner refs got %d", want, got)
		return
	}

	expectedRefs := []metav1.OwnerReference{{
		APIVersion: "run.projectriff.io/v1alpha1",
		Kind:       "RequestProcessor",
		Name:       testRequestProcessorName,
	}}
	if diff := cmp.Diff(expectedRefs, ownerRefs, cmpopts.IgnoreFields(expectedRefs[0], "Controller", "BlockOwnerDeletion")); diff != "" {
		t.Errorf("Unexpected requestprocessor owner refs diff (-want +got): %v", diff)
	}
}

func createRequestProcessorMeta() *runv1alpha1.RequestProcessor {
	return &runv1alpha1.RequestProcessor{
		ObjectMeta: metav1.ObjectMeta{
			Name:      testRequestProcessorName,
			Namespace: testRequestProcessorNamespace,
		},
	}
}