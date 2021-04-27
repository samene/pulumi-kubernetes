// Copyright 2016-2021, Pulumi Corporation.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package test

import (
	"path/filepath"
	"testing"

	"github.com/pulumi/pulumi/pkg/v3/testing/integration"
	"github.com/pulumi/pulumi/sdk/v3/go/common/resource"
	"github.com/pulumi/pulumi/sdk/v3/go/common/tokens"
	"github.com/stretchr/testify/assert"
)

var baseOptions = &integration.ProgramTestOptions{
	Verbose: true,
	Dependencies: []string{
		"Pulumi.Kubernetes",
	},
}

func TestDotnet_Basic(t *testing.T) {
	test := baseOptions.With(integration.ProgramTestOptions{
		Dir:                  "basic",
		Quick:                true,
		ExpectRefreshChanges: true, // The CRD sometimes, but not always, has changes during refresh.
	})
	integration.ProgramTest(t, &test)
}

func TestDotnet_Guestbook(t *testing.T) {
	test := baseOptions.With(integration.ProgramTestOptions{
		Dir:   "guestbook",
		Quick: true,
	})
	integration.ProgramTest(t, &test)
}

func TestDotnet_YamlUrl(t *testing.T) {
	test := baseOptions.With(integration.ProgramTestOptions{
		Dir:   "yaml-url",
		Quick: true,
		ExtraRuntimeValidation: func(
			t *testing.T, stackInfo integration.RuntimeValidationStackInfo,
		) {
			assert.NotNil(t, stackInfo.Deployment)
			assert.Equal(t, 18, len(stackInfo.Deployment.Resources))
		},
	})
	integration.ProgramTest(t, &test)
}

func TestDotnet_YamlLocal(t *testing.T) {
	test := baseOptions.With(integration.ProgramTestOptions{
		Dir:   "yaml-local",
		Quick: true,
		ExtraRuntimeValidation: func(
			t *testing.T, stackInfo integration.RuntimeValidationStackInfo,
		) {
			assert.NotNil(t, stackInfo.Deployment)
			assert.Equal(t, 7, len(stackInfo.Deployment.Resources))
		},
	})
	integration.ProgramTest(t, &test)
}

func TestDotnet_Helm(t *testing.T) {
	test := baseOptions.With(integration.ProgramTestOptions{
		Dir:   filepath.Join("helm", "step1"),
		Quick: true,
		ExtraRuntimeValidation: func(t *testing.T, stackInfo integration.RuntimeValidationStackInfo) {
			// Ensure that all `Services` have `status` marked as a `Secret`
			for _, res := range stackInfo.Deployment.Resources {
				if res.Type == tokens.Type("kubernetes:core/v1:Service") {
					spec, has := res.Outputs["status"]
					assert.True(t, has)
					specMap, is := spec.(map[string]interface{})
					assert.True(t, is)
					sigKey, has := specMap[resource.SigKey]
					assert.True(t, has)
					assert.Equal(t, resource.SecretSig, sigKey)
				}
			}
		},
		EditDirs: []integration.EditDir{
			{
				Dir:             filepath.Join("helm", "step2"),
				Additive:        true,
				ExpectNoChanges: true,
			},
		},
	})
	integration.ProgramTest(t, &test)
}

func TestDotnet_HelmLocal(t *testing.T) {
	test := baseOptions.With(integration.ProgramTestOptions{
		Dir:   filepath.Join("helm-local", "step1"),
		Quick: true,
		ExtraRuntimeValidation: func(t *testing.T, stackInfo integration.RuntimeValidationStackInfo) {
			assert.NotNil(t, stackInfo.Deployment)
			assert.Equal(t, 11, len(stackInfo.Deployment.Resources))
		},
		EditDirs: []integration.EditDir{
			{
				Dir:             filepath.Join("helm-local", "step2"),
				Additive:        true,
				ExpectNoChanges: true,
			},
		},
	})
	integration.ProgramTest(t, &test)
}

func TestDotnet_HelmApiVersions(t *testing.T) {
	test := baseOptions.With(integration.ProgramTestOptions{
		Dir:   filepath.Join("helm-api-versions", "step1"),
		Quick: true,
		ExtraRuntimeValidation: func(t *testing.T, stackInfo integration.RuntimeValidationStackInfo) {
			assert.NotNil(t, stackInfo.Deployment)
			assert.Equal(t, 7, len(stackInfo.Deployment.Resources))
		},
		EditDirs: []integration.EditDir{
			{
				Dir:             filepath.Join("helm-api-versions", "step2"),
				Additive:        true,
				ExpectNoChanges: true,
			},
		},
	})
	integration.ProgramTest(t, &test)
}

func TestDotnet_CustomResource(t *testing.T) {
	test := baseOptions.With(integration.ProgramTestOptions{
		Dir:                  "custom-resource",
		Quick:                true,
		ExpectRefreshChanges: true, // The CRD sometimes, but not always, has changes during refresh.
	})
	integration.ProgramTest(t, &test)
}

func TestDotnet_Kustomize(t *testing.T) {
	test := baseOptions.With(integration.ProgramTestOptions{
		Dir:   "kustomize",
		Quick: true,
	})
	integration.ProgramTest(t, &test)
}
