// Copyright 2016-2023, Pulumi Corporation.
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

// nolint:goconst
package test

import (
	b64 "encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"testing"
	"time"

	"github.com/pulumi/pulumi-kubernetes/provider/v4/pkg/openapi"
	"github.com/pulumi/pulumi-kubernetes/tests/v4"
	"github.com/pulumi/pulumi/pkg/v3/resource/deploy/providers"
	"github.com/pulumi/pulumi/pkg/v3/testing/integration"
	"github.com/pulumi/pulumi/sdk/v3/go/common/apitype"
	"github.com/pulumi/pulumi/sdk/v3/go/common/resource"
	"github.com/pulumi/pulumi/sdk/v3/go/common/tokens"
	"github.com/stretchr/testify/assert"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
)

var baseOptions = &integration.ProgramTestOptions{
	Verbose: true,
	Dependencies: []string{
		"@pulumi/kubernetes",
	},
	Env: []string{
		"PULUMI_K8S_CLIENT_BURST=200",
		"PULUMI_K8S_CLIENT_QPS=100",
	},
}

func TestAliases(t *testing.T) {
	test := baseOptions.With(integration.ProgramTestOptions{
		Dir:   filepath.Join("aliases", "step1"),
		Quick: true,
		ExtraRuntimeValidation: func(t *testing.T, stackInfo integration.RuntimeValidationStackInfo) {
			assert.NotNil(t, stackInfo.Deployment)
			assert.Equal(t, 4, len(stackInfo.Deployment.Resources))

			tests.SortResourcesByURN(stackInfo)

			deployment := stackInfo.Deployment.Resources[0]
			assert.Equal(t, "alias-test", string(deployment.URN.Name()))
			assert.Equal(t, "kubernetes:apps/v1:Deployment", string(deployment.Type))
			containers, _ := openapi.Pluck(deployment.Outputs, "spec", "template", "spec", "containers")
			containerStatus := containers.([]any)[0].(map[string]any)
			image := containerStatus["image"]
			assert.Equal(t, image.(string), "nginx:1.14")
		},
		EditDirs: []integration.EditDir{
			{
				Dir:      filepath.Join("aliases", "step2"),
				Additive: true,
				ExtraRuntimeValidation: func(t *testing.T, stackInfo integration.RuntimeValidationStackInfo) {
					assert.NotNil(t, stackInfo.Deployment)
					assert.Equal(t, 4, len(stackInfo.Deployment.Resources))

					tests.SortResourcesByURN(stackInfo)

					deployment := stackInfo.Deployment.Resources[0]
					assert.Equal(t, "alias-test", string(deployment.URN.Name()))
					assert.Equal(t, "kubernetes:apps/v1:Deployment", string(deployment.Type))
					containers, _ := openapi.Pluck(deployment.Outputs, "spec", "template", "spec", "containers")
					containerStatus := containers.([]any)[0].(map[string]any)
					image := containerStatus["image"]
					assert.Equal(t, image.(string), "nginx:1.15")
				},
			},
		},
	})
	integration.ProgramTest(t, &test)
}

func TestAutonaming(t *testing.T) {
	var step1Name any
	var step2Name any
	var step3Name any

	test := baseOptions.With(integration.ProgramTestOptions{
		Dir:   filepath.Join("autonaming", "step1"),
		Quick: true,
		ExtraRuntimeValidation: func(t *testing.T, stackInfo integration.RuntimeValidationStackInfo) {
			assert.NotNil(t, stackInfo.Deployment)
			assert.Equal(t, 4, len(stackInfo.Deployment.Resources))

			tests.SortResourcesByURN(stackInfo)

			stackRes := stackInfo.Deployment.Resources[3]
			assert.Equal(t, resource.RootStackType, stackRes.URN.Type())

			provRes := stackInfo.Deployment.Resources[2]
			assert.True(t, providers.IsProviderType(provRes.URN.Type()))

			//
			// Assert Pod is successfully given a unique name by Pulumi.
			//

			pod := stackInfo.Deployment.Resources[1]
			assert.Equal(t, "autonaming-test", string(pod.URN.Name()))
			step1Name, _ = openapi.Pluck(pod.Outputs, "metadata", "name")
			assert.True(t, strings.HasPrefix(step1Name.(string), "autonaming-test-"))

			autonamed, _ := openapi.Pluck(pod.Outputs, "metadata", "annotations", "pulumi.com/autonamed")
			assert.Equal(t, "true", autonamed)
		},
		EditDirs: []integration.EditDir{
			{
				Dir:      filepath.Join("autonaming", "step2"),
				Additive: true,
				ExtraRuntimeValidation: func(t *testing.T, stackInfo integration.RuntimeValidationStackInfo) {
					assert.NotNil(t, stackInfo.Deployment)
					assert.Equal(t, 4, len(stackInfo.Deployment.Resources))

					tests.SortResourcesByURN(stackInfo)

					stackRes := stackInfo.Deployment.Resources[3]
					assert.Equal(t, resource.RootStackType, stackRes.URN.Type())

					provRes := stackInfo.Deployment.Resources[2]
					assert.True(t, providers.IsProviderType(provRes.URN.Type()))

					//
					// Assert Pod was replaced, i.e., destroyed and re-created, with allocating a new name.
					//

					pod := stackInfo.Deployment.Resources[1]
					assert.Equal(t, "autonaming-test", string(pod.URN.Name()))
					step2Name, _ = openapi.Pluck(pod.Outputs, "metadata", "name")
					assert.True(t, strings.HasPrefix(step2Name.(string), "autonaming-test-"))

					autonamed, _ := openapi.Pluck(pod.Outputs, "metadata", "annotations", "pulumi.com/autonamed")
					assert.Equal(t, "true", autonamed)

					assert.NotEqual(t, step1Name, step2Name)
				},
			},
			{
				Dir:      filepath.Join("autonaming", "step3"),
				Additive: true,
				ExtraRuntimeValidation: func(t *testing.T, stackInfo integration.RuntimeValidationStackInfo) {
					assert.NotNil(t, stackInfo.Deployment)
					assert.Equal(t, 4, len(stackInfo.Deployment.Resources))

					tests.SortResourcesByURN(stackInfo)

					stackRes := stackInfo.Deployment.Resources[3]
					assert.Equal(t, resource.RootStackType, stackRes.URN.Type())

					provRes := stackInfo.Deployment.Resources[2]
					assert.True(t, providers.IsProviderType(provRes.URN.Type()))

					//
					// Assert Pod was NOT replaced, and has the same name, previously allocated by Pulumi.
					//

					pod := stackInfo.Deployment.Resources[1]
					assert.Equal(t, "autonaming-test", string(pod.URN.Name()))
					step3Name, _ = openapi.Pluck(pod.Outputs, "metadata", "name")
					assert.True(t, strings.HasPrefix(step3Name.(string), "autonaming-test-"))

					autonamed, _ := openapi.Pluck(pod.Outputs, "metadata", "annotations", "pulumi.com/autonamed")
					assert.Equal(t, "true", autonamed)

					assert.Equal(t, step2Name, step3Name)
				},
			},
			{
				Dir:      filepath.Join("autonaming", "step4"),
				Additive: true,
				ExtraRuntimeValidation: func(t *testing.T, stackInfo integration.RuntimeValidationStackInfo) {
					assert.NotNil(t, stackInfo.Deployment)
					assert.Equal(t, 4, len(stackInfo.Deployment.Resources))

					tests.SortResourcesByURN(stackInfo)

					stackRes := stackInfo.Deployment.Resources[3]
					assert.Equal(t, resource.RootStackType, stackRes.URN.Type())

					provRes := stackInfo.Deployment.Resources[2]
					assert.True(t, providers.IsProviderType(provRes.URN.Type()))

					//
					// User has specified their own name for the Pod, so we replace it, and Pulumi does NOT
					// allocate a name on its own.
					//

					pod := stackInfo.Deployment.Resources[1]
					assert.Equal(t, "autonaming-test", string(pod.URN.Name()))
					name, _ := openapi.Pluck(pod.Outputs, "metadata", "name")
					assert.Equal(t, "autonaming-test", name.(string))

					_, autonamed := openapi.Pluck(pod.Outputs, "metadata", "annotations", "pulumi.com/autonamed")
					assert.False(t, autonamed)
				},
			},
		},
	})
	integration.ProgramTest(t, &test)
}

func TestCRDs(t *testing.T) {
	test := baseOptions.With(integration.ProgramTestOptions{
		Dir:                  filepath.Join("crds", "step1"),
		Quick:                false,
		ExpectRefreshChanges: true,
		ExtraRuntimeValidation: func(t *testing.T, stackInfo integration.RuntimeValidationStackInfo) {
			assert.NotNil(t, stackInfo.Deployment)
			assert.Equal(t, 5, len(stackInfo.Deployment.Resources))

			tests.SortResourcesByURN(stackInfo)

			crd := stackInfo.Deployment.Resources[0]
			namespace := stackInfo.Deployment.Resources[1]
			ct1 := stackInfo.Deployment.Resources[2]
			provRes := stackInfo.Deployment.Resources[3]
			stackRes := stackInfo.Deployment.Resources[4]

			assert.Equal(t, resource.RootStackType, stackRes.URN.Type())
			assert.True(t, providers.IsProviderType(provRes.URN.Type()))

			assert.Equal(t, tokens.Type("kubernetes:core/v1:Namespace"), namespace.URN.Type())

			//
			// Assert that CRD and CR exist
			//

			assert.Equal(t, "foobar", string(crd.URN.Name()))
			assert.Equal(t, "my-new-foobar-object", string(ct1.URN.Name()))
		},
		EditDirs: []integration.EditDir{
			{
				Dir:      filepath.Join("crds", "step2"),
				Additive: true,
				ExtraRuntimeValidation: func(t *testing.T, stackInfo integration.RuntimeValidationStackInfo) {
					assert.NotNil(t, stackInfo.Deployment)
					assert.Equal(t, 4, len(stackInfo.Deployment.Resources))

					tests.SortResourcesByURN(stackInfo)

					namespace := stackInfo.Deployment.Resources[0]
					provRes := stackInfo.Deployment.Resources[2]
					stackRes := stackInfo.Deployment.Resources[3]

					assert.Equal(t, resource.RootStackType, stackRes.URN.Type())
					assert.True(t, providers.IsProviderType(provRes.URN.Type()))

					assert.Equal(t, tokens.Type("kubernetes:core/v1:Namespace"), namespace.URN.Type())
				},
			},
		},
	})
	integration.ProgramTest(t, &test)
}

func TestPod(t *testing.T) {
	test := baseOptions.With(integration.ProgramTestOptions{
		Dir:   filepath.Join("delete-before-replace", "step1"),
		Quick: true,
		ExtraRuntimeValidation: func(t *testing.T, stackInfo integration.RuntimeValidationStackInfo) {
			assert.NotNil(t, stackInfo.Deployment)
			assert.Equal(t, 4, len(stackInfo.Deployment.Resources))

			tests.SortResourcesByURN(stackInfo)

			stackRes := stackInfo.Deployment.Resources[3]
			assert.Equal(t, resource.RootStackType, stackRes.URN.Type())

			provRes := stackInfo.Deployment.Resources[2]
			assert.True(t, providers.IsProviderType(provRes.URN.Type()))

			//
			// Assert pod is successfully created.
			//

			pod := stackInfo.Deployment.Resources[1]
			name, _ := openapi.Pluck(pod.Outputs, "metadata", "name")
			assert.Equal(t, name.(string), "pod-test")

			// Not autonamed.
			_, autonamed := openapi.Pluck(pod.Outputs, "metadata", "annotations", "pulumi.com/autonamed")
			assert.False(t, autonamed)

			// Status is "Running"
			phase, _ := openapi.Pluck(pod.Outputs, "status", "phase")
			assert.Equal(t, "Running", phase)

			// Status "Ready" is "True".
			conditions, _ := openapi.Pluck(pod.Outputs, "status", "conditions")
			ready := conditions.([]any)[1].(map[string]any)
			readyType := ready["type"]
			assert.Equal(t, "Ready", readyType)
			readyStatus := ready["status"]
			assert.Equal(t, "True", readyStatus)

			// Container is called "nginx" and uses image "docker.io/library/nginx:1.13-alpine".
			containerStatuses, _ := openapi.Pluck(pod.Outputs, "status", "containerStatuses")
			containerStatus := containerStatuses.([]any)[0].(map[string]any)
			containerName := containerStatus["name"]
			assert.Equal(t, "nginx", containerName)
			image := containerStatus["image"]
			assert.Equal(t, "docker.io/library/nginx:1.13-alpine", image)
		},
		EditDirs: []integration.EditDir{
			{
				Dir:      filepath.Join("delete-before-replace", "step2"),
				Additive: true,
				ExtraRuntimeValidation: func(t *testing.T, stackInfo integration.RuntimeValidationStackInfo) {
					assert.NotNil(t, stackInfo.Deployment)
					assert.Equal(t, 4, len(stackInfo.Deployment.Resources))

					tests.SortResourcesByURN(stackInfo)

					stackRes := stackInfo.Deployment.Resources[3]
					assert.Equal(t, resource.RootStackType, stackRes.URN.Type())

					provRes := stackInfo.Deployment.Resources[2]
					assert.True(t, providers.IsProviderType(provRes.URN.Type()))

					//
					// Assert Pod is deleted before being replaced with the new Pod, running
					// nginx:1.15-alpine.
					//
					// Because the Pod name is supplied in the resource definition, we are forced to delete it
					// before replacing it, as otherwise Kubernetes would complain that it can't have two Pods
					// with the same name.
					//

					pod := stackInfo.Deployment.Resources[1]
					name, _ := openapi.Pluck(pod.Outputs, "metadata", "name")
					assert.Equal(t, name.(string), "pod-test")

					// Not autonamed.
					_, autonamed := openapi.Pluck(pod.Outputs, "metadata", "annotations", "pulumi.com/autonamed")
					assert.False(t, autonamed)

					// Status is "Running"
					phase, _ := openapi.Pluck(pod.Outputs, "status", "phase")
					assert.Equal(t, "Running", phase)

					// Status "Ready" is "True".
					conditions, _ := openapi.Pluck(pod.Outputs, "status", "conditions")
					ready := conditions.([]any)[1].(map[string]any)
					readyType := ready["type"]
					assert.Equal(t, "Ready", readyType)
					readyStatus := ready["status"]
					assert.Equal(t, "True", readyStatus)

					// Container is called "nginx" and uses image "docker.io/library/nginx:1.15-alpine".
					containerStatuses, _ := openapi.Pluck(pod.Outputs, "status", "containerStatuses")
					containerStatus := containerStatuses.([]any)[0].(map[string]any)
					containerName := containerStatus["name"]
					assert.Equal(t, "nginx", containerName)
					image := containerStatus["image"]
					assert.Equal(t, "docker.io/library/nginx:1.15-alpine", image)
				},
			},
		},
	})
	integration.ProgramTest(t, &test)
}

func TestDeploymentRollout(t *testing.T) {
	test := baseOptions.With(integration.ProgramTestOptions{
		Dir:   filepath.Join("deployment-rollout", "step1"),
		Quick: true,
		ExtraRuntimeValidation: func(t *testing.T, stackInfo integration.RuntimeValidationStackInfo) {
			assert.NotNil(t, stackInfo.Deployment)
			assert.Equal(t, 4, len(stackInfo.Deployment.Resources))

			tests.SortResourcesByURN(stackInfo)

			appsv1Deploy := stackInfo.Deployment.Resources[0]
			namespace := stackInfo.Deployment.Resources[1]
			provRes := stackInfo.Deployment.Resources[2]
			stackRes := stackInfo.Deployment.Resources[3]

			assert.Equal(t, resource.RootStackType, stackRes.URN.Type())
			assert.True(t, providers.IsProviderType(provRes.URN.Type()))

			assert.Equal(t, tokens.Type("kubernetes:core/v1:Namespace"), namespace.URN.Type())

			//
			// Assert deployment is successfully created.
			//

			name, _ := openapi.Pluck(appsv1Deploy.Outputs, "metadata", "name")
			assert.True(t, strings.Contains(name.(string), "nginx"))
			containers, _ := openapi.Pluck(appsv1Deploy.Outputs, "spec", "template", "spec", "containers")
			containerStatus := containers.([]any)[0].(map[string]any)
			image := containerStatus["image"]
			assert.Equal(t, image.(string), "nginx")

			assert.True(t, strings.Contains(name.(string), "nginx"))
			containers, _ = openapi.Pluck(appsv1Deploy.Outputs, "spec", "template", "spec", "containers")
			containerStatus = containers.([]any)[0].(map[string]any)
			image = containerStatus["image"]
			assert.Equal(t, image.(string), "nginx")
		},
		EditDirs: []integration.EditDir{
			{
				Dir:      filepath.Join("deployment-rollout", "step2"),
				Additive: true,
				ExtraRuntimeValidation: func(t *testing.T, stackInfo integration.RuntimeValidationStackInfo) {
					assert.NotNil(t, stackInfo.Deployment)
					assert.Equal(t, 4, len(stackInfo.Deployment.Resources))

					tests.SortResourcesByURN(stackInfo)

					appsv1Deploy := stackInfo.Deployment.Resources[0]
					namespace := stackInfo.Deployment.Resources[1]
					provRes := stackInfo.Deployment.Resources[2]
					stackRes := stackInfo.Deployment.Resources[3]

					assert.Equal(t, resource.RootStackType, stackRes.URN.Type())
					assert.True(t, providers.IsProviderType(provRes.URN.Type()))

					assert.Equal(t, tokens.Type("kubernetes:core/v1:Namespace"), namespace.URN.Type())

					//
					// Assert deployment is updated successfully.
					//

					name, _ := openapi.Pluck(appsv1Deploy.Outputs, "metadata", "name")
					assert.True(t, strings.Contains(name.(string), "nginx"))
					containers, _ := openapi.Pluck(appsv1Deploy.Outputs, "spec", "template", "spec", "containers")
					containerStatus := containers.([]any)[0].(map[string]any)
					image := containerStatus["image"]
					assert.Equal(t, image.(string), "nginx:stable")
				},
			},
		},
	})
	integration.ProgramTest(t, &test)
}

func TestEmptyArray(t *testing.T) {
	test := baseOptions.With(integration.ProgramTestOptions{
		Dir: filepath.Join("empty-array", "step1"),
		EditDirs: []integration.EditDir{
			{
				Dir:      filepath.Join("empty-array", "step2"),
				Additive: true,
			},
		},
	})
	integration.ProgramTest(t, &test)
}

func TestGet(t *testing.T) {
	test := baseOptions.With(integration.ProgramTestOptions{
		Dir:                  filepath.Join("get", "step1"),
		Quick:                true,
		ExpectRefreshChanges: true,
		ExtraRuntimeValidation: func(t *testing.T, stackInfo integration.RuntimeValidationStackInfo) {
			assert.NotNil(t, stackInfo.Deployment)
			assert.Equal(t, 7, len(stackInfo.Deployment.Resources))

			//
			// Assert we can use .get to retrieve the kube-api Service.
			//

			service := stackInfo.Outputs["svc"].(map[string]any)
			svcURN, _ := openapi.Pluck(service, "urn")
			assert.Containsf(t, svcURN, "kube-api", "urn missing expected name")
			svcName, _ := openapi.Pluck(service, "metadata", "name")
			assert.Equalf(t, "kubernetes", svcName, "unexpected service name")

			//
			// Assert that the uninitialized Service exists
			//

			awaitSvc := stackInfo.Outputs["awaitSvc"].(map[string]any)
			awaitSvcName, _ := openapi.Pluck(awaitSvc, "metadata", "name")
			assert.Equalf(t, "test", awaitSvcName, "unexpected service name")
			awaitSvcAnnotation, ok := openapi.Pluck(awaitSvc, "metadata", "annotations", "pulumi.com/skipAwait")
			assert.Truef(t, ok, "failed to find skipAwait annotation")
			assert.Equalf(t, "true", awaitSvcAnnotation, "expected annotation to be true")

			//
			// Assert that CRD and CR exist
			//

			crd := stackInfo.Outputs["ct"].(map[string]any)
			crdURN, _ := openapi.Pluck(crd, "urn")
			assert.Containsf(t, crdURN, "crontab", "urn missing expected name")

			cr := stackInfo.Outputs["cr"].(map[string]any)
			crURN, _ := openapi.Pluck(cr, "urn")
			assert.Containsf(t, crURN, "my-new-cron-object", "urn missing expected name")
		},
		EditDirs: []integration.EditDir{
			{
				Dir:      filepath.Join("get", "step2"),
				Additive: true,
				ExtraRuntimeValidation: func(t *testing.T, stackInfo integration.RuntimeValidationStackInfo) {
					assert.NotNil(t, stackInfo.Deployment)
					assert.Equal(t, 10, len(stackInfo.Deployment.Resources))

					//
					// Assert we can use .get to retrieve the kube-api Service.
					//

					service := stackInfo.Outputs["svc"].(map[string]any)
					svcURN, _ := openapi.Pluck(service, "urn")
					assert.Containsf(t, svcURN, "kube-api", "urn missing expected name")
					svcName, _ := openapi.Pluck(service, "metadata", "name")
					assert.Equalf(t, "kubernetes", svcName, "unexpected service name")

					//
					// Assert that the uninitialized Service exists
					//

					awaitSvc := stackInfo.Outputs["awaitSvc"].(map[string]any)
					awaitSvcName, _ := openapi.Pluck(awaitSvc, "metadata", "name")
					assert.Equalf(t, "test", awaitSvcName, "unexpected service name")
					awaitSvcAnnotation, ok := openapi.Pluck(awaitSvc, "metadata", "annotations", "pulumi.com/skipAwait")
					assert.Truef(t, ok, "failed to find skipAwait annotation")
					assert.Equalf(t, "true", awaitSvcAnnotation, "expected annotation to be true")

					//
					// Assert we can use .get to retrieve a Service that would fail await logic.
					//

					awaitSvcGet := stackInfo.Outputs["awaitSvcGet"].(map[string]any)
					awaitSvcGetURN, _ := openapi.Pluck(awaitSvcGet, "urn")
					assert.Containsf(t, awaitSvcGetURN, "await", "urn missing expected name")

					//
					// Assert we can use an output from a Service that would fail await logic.
					//

					cm := stackInfo.Outputs["cm"].(map[string]any)
					cmURN, _ := openapi.Pluck(cm, "urn")
					assert.Containsf(t, cmURN, "svc-test", "urn missing expected name")
					clusterIP, _ := openapi.Pluck(cm, "data", "key")
					assert.NotEmptyf(t, clusterIP, "clusterIP should be set")

					//
					// Assert that CRD and CR exist
					//

					crd := stackInfo.Outputs["ct"].(map[string]any)
					crdURN, _ := openapi.Pluck(crd, "urn")
					assert.Containsf(t, crdURN, "crontab", "urn missing expected name")

					cr := stackInfo.Outputs["cr"].(map[string]any)
					crURN, _ := openapi.Pluck(cr, "urn")
					assert.Containsf(t, crURN, "my-new-cron-object", "urn missing expected name")

					//
					// Assert we can use .get to retrieve CRDs.
					//

					crGet := stackInfo.Outputs["crGet"].(map[string]any)
					crGetURN, _ := openapi.Pluck(crGet, "urn")
					assert.Containsf(t, crGetURN, "my-new-cron-object-get", "urn missing expected name")
				},
			},
		},
	})
	integration.ProgramTest(t, &test)
}

func TestIstio(t *testing.T) {
	test := baseOptions.With(integration.ProgramTestOptions{
		Dir:         filepath.Join("istio", "step1"),
		Quick:       true,
		SkipRefresh: true,
		ExtraRuntimeValidation: func(t *testing.T, stackInfo integration.RuntimeValidationStackInfo) {
			frontend := stackInfo.Outputs["frontendIp"].(string)

			// Retry the GET on the Istio gateway repeatedly. Istio doesn't publish `.status` on any
			// of its CRDs, so this is as reliable as we can be right now.
			for i := 1; i < 10; i++ {
				req, err := http.Get(fmt.Sprintf("http://%s", frontend))
				if err != nil {
					fmt.Printf("Request to Istio gateway failed: %v\n", err)
					time.Sleep(time.Second * 10)
				} else if req.StatusCode == 200 {
					return
				}
			}

			assert.Fail(t, "Maximum Istio gateway request retries exceeded")
		},
	})
	integration.ProgramTest(t, &test)
}

func TestKustomize(t *testing.T) {
	test := baseOptions.With(integration.ProgramTestOptions{
		Dir:   filepath.Join("kustomize", "step1"),
		Quick: true,
		OrderedConfig: []integration.ConfigValue{
			{
				Key:   "pulumi:disable-default-providers[0]",
				Value: "kubernetes",
				Path:  true,
			},
		},
		ExtraRuntimeValidation: func(t *testing.T, stackInfo integration.RuntimeValidationStackInfo) {
			assert.NotNil(t, stackInfo.Deployment)

			// Assert that we've retrieved the kustomizations and provisioned them.
			assert.Equal(t, 12, len(stackInfo.Deployment.Resources))
		},
	})
	integration.ProgramTest(t, &test)
}

// TestKustomizeHelmChart verifies the helmChart plugin support for Kustomize. This requires the `helm` binary to be
// on the system PATH to succeed.
// Example based on https://github.com/kubernetes-sigs/kustomize/blob/v3.3.1/examples/chart.md
func TestKustomizeHelmChart(t *testing.T) {
	test := baseOptions.With(integration.ProgramTestOptions{
		Dir: filepath.Join("kustomizeHelmChart", "step1"),
		// Just test that the plugin integration is working with preview
		SkipUpdate:             true,
		SkipRefresh:            true,
		SkipEmptyPreviewUpdate: true,
		SkipExportImport:       true,
		OrderedConfig: []integration.ConfigValue{
			{
				Key:   "pulumi:disable-default-providers[0]",
				Value: "kubernetes",
				Path:  true,
			},
		},
		Env: []string{
			"PULUMI_K8S_KUSTOMIZE_HELM=true", // This experimental feature is currently gated behind a feature flag.
		},
	})
	integration.ProgramTest(t, &test)
}

func TestNamespace(t *testing.T) {
	var nmPodName, defaultPodName string
	test := baseOptions.With(integration.ProgramTestOptions{
		Dir:   filepath.Join("namespace", "step1"),
		Quick: true,
		ExtraRuntimeValidation: func(t *testing.T, stackInfo integration.RuntimeValidationStackInfo) {
			assert.NotNil(t, stackInfo.Deployment)
			assert.Equal(t, 5, len(stackInfo.Deployment.Resources))

			tests.SortResourcesByURN(stackInfo)

			stackRes := stackInfo.Deployment.Resources[4]
			assert.Equal(t, resource.RootStackType, stackRes.URN.Type())

			provRes := stackInfo.Deployment.Resources[3]
			assert.True(t, providers.IsProviderType(provRes.URN.Type()))

			// Assert the Namespace was created
			namespace := stackInfo.Deployment.Resources[0]
			assert.Equal(t, tokens.Type("kubernetes:core/v1:Namespace"), namespace.URN.Type())

			// Assert the "no metadata" Pod was created in the "default" namespace.
			nmPod := stackInfo.Deployment.Resources[2]
			assert.Equal(t, tokens.Type("kubernetes:core/v1:Pod"), nmPod.URN.Type())
			nmPodNamespace, _ := openapi.Pluck(nmPod.Outputs, "metadata", "namespace")
			assert.Equal(t, nmPodNamespace.(string), "default")
			nmPodNameRaw, _ := openapi.Pluck(nmPod.Outputs, "metadata", "name")
			nmPodName = nmPodNameRaw.(string)
			assert.True(t, strings.HasPrefix(nmPodName, "no-metadata-pod"))

			// Assert the "explicit default namespace" Pod was created in the "default" namespace.
			defaultPod := stackInfo.Deployment.Resources[1]
			assert.Equal(t, tokens.Type("kubernetes:core/v1:Pod"), defaultPod.URN.Type())
			defaultPodNamespace, _ := openapi.Pluck(defaultPod.Outputs, "metadata", "namespace")
			assert.Equal(t, defaultPodNamespace.(string), "default")
			defaultPodNameRaw, _ := openapi.Pluck(defaultPod.Outputs, "metadata", "name")
			defaultPodName = defaultPodNameRaw.(string)
			assert.True(t, strings.HasPrefix(defaultPodName, "default-ns-pod"))
		},
		EditDirs: []integration.EditDir{
			{
				Dir:      filepath.Join("namespace", "step2"),
				Additive: true,
				ExtraRuntimeValidation: func(t *testing.T, stackInfo integration.RuntimeValidationStackInfo) {
					assert.NotNil(t, stackInfo.Deployment)
					assert.Equal(t, 5, len(stackInfo.Deployment.Resources))

					tests.SortResourcesByURN(stackInfo)

					stackRes := stackInfo.Deployment.Resources[4]
					assert.Equal(t, resource.RootStackType, stackRes.URN.Type())

					provRes := stackInfo.Deployment.Resources[3]
					assert.True(t, providers.IsProviderType(provRes.URN.Type()))

					// Assert that the Namespace was updated with the expected label.
					namespace := stackInfo.Deployment.Resources[0]
					assert.Equal(t, tokens.Type("kubernetes:core/v1:Namespace"), namespace.URN.Type())
					namespaceLabels, _ := openapi.Pluck(namespace.Outputs, "metadata", "labels")
					assert.True(t, namespaceLabels.(map[string]any)["hello"] == "world")
				},
			},
			{
				Dir:      filepath.Join("namespace", "step3"),
				Additive: true,
				ExtraRuntimeValidation: func(t *testing.T, stackInfo integration.RuntimeValidationStackInfo) {
					assert.NotNil(t, stackInfo.Deployment)
					assert.Equal(t, 5, len(stackInfo.Deployment.Resources))

					tests.SortResourcesByURN(stackInfo)

					stackRes := stackInfo.Deployment.Resources[4]
					assert.Equal(t, resource.RootStackType, stackRes.URN.Type())

					provRes := stackInfo.Deployment.Resources[3]
					assert.True(t, providers.IsProviderType(provRes.URN.Type()))

					namespace := stackInfo.Deployment.Resources[0]
					assert.Equal(t, tokens.Type("kubernetes:core/v1:Namespace"), namespace.URN.Type())

					// Assert the "no metadata" Pod has not changed.
					nmPod := stackInfo.Deployment.Resources[2]
					assert.Equal(t, tokens.Type("kubernetes:core/v1:Pod"), nmPod.URN.Type())
					nmPodNamespace, _ := openapi.Pluck(nmPod.Outputs, "metadata", "namespace")
					assert.Equal(t, nmPodNamespace.(string), "default")
					nmPodNameRaw, _ := openapi.Pluck(nmPod.Outputs, "metadata", "name")
					assert.True(t, strings.HasPrefix(nmPodNameRaw.(string), "no-metadata-pod"))
					assert.Equal(t, nmPodNameRaw.(string), nmPodName)

					// Assert the "explicit default namespace" has not changed.
					defaultPod := stackInfo.Deployment.Resources[1]
					assert.Equal(t, tokens.Type("kubernetes:core/v1:Pod"), defaultPod.URN.Type())
					defaultPodNamespace, _ := openapi.Pluck(defaultPod.Outputs, "metadata", "namespace")
					assert.Equal(t, defaultPodNamespace.(string), "default")
					defaultPodNameRaw, _ := openapi.Pluck(defaultPod.Outputs, "metadata", "name")
					assert.True(t, strings.HasPrefix(defaultPodNameRaw.(string), "default-ns-pod"))
					assert.Equal(t, defaultPodNameRaw.(string), defaultPodName)
				},
			},
		},
	})
	integration.ProgramTest(t, &test)
}

// FIXME(levi): Figure out why this test is flaky, and re-enable it in CI.
//  https://github.com/pulumi/pulumi-kubernetes/issues/1016
//func TestPerformance(t *testing.T) {
//	test := baseOptions.With(integration.ProgramTestOptions{
//		Dir:                  filepath.Join("performance", "step1"),
//		ExpectRefreshChanges: true, // The Mutating and Validating webhooks update on refresh.
//	})
//	integration.ProgramTest(t, &test)
//}

func TestProvider(t *testing.T) {
	test := baseOptions.With(integration.ProgramTestOptions{
		Dir:   filepath.Join("provider", "step1"),
		Quick: true,
		ExtraRuntimeValidation: func(t *testing.T, stackInfo integration.RuntimeValidationStackInfo) {
			assert.NotNil(t, stackInfo.Deployment)
			assert.Equal(t, 10, len(stackInfo.Deployment.Resources))

			tests.SortResourcesByURN(stackInfo)

			stackRes := stackInfo.Deployment.Resources[9]
			assert.Equal(t, resource.RootStackType, stackRes.URN.Type())

			k8sPathProvider := stackInfo.Deployment.Resources[8]
			assert.True(t, providers.IsProviderType(k8sPathProvider.URN.Type()))

			k8sContentsProvider := stackInfo.Deployment.Resources[7]
			assert.True(t, providers.IsProviderType(k8sContentsProvider.URN.Type()))

			defaultProvider := stackInfo.Deployment.Resources[6]
			assert.True(t, providers.IsProviderType(defaultProvider.URN.Type()))

			// Assert the provider default Namespace (ns1) was created
			ns1 := stackInfo.Deployment.Resources[0]
			assert.Equal(t, tokens.Type("kubernetes:core/v1:Namespace"), ns1.URN.Type())
			providerNsName, _ := openapi.Pluck(ns1.Outputs, "metadata", "name")

			// Assert the ns2 Namespace was created
			ns2 := stackInfo.Deployment.Resources[1]
			assert.Equal(t, tokens.Type("kubernetes:core/v1:Namespace"), ns2.URN.Type())
			ns2Name, _ := openapi.Pluck(ns2.Outputs, "metadata", "name")

			// Assert the other Namespace was created and doesn't use the provider default.
			otherNamespace := stackInfo.Deployment.Resources[2]
			assert.Equal(t, tokens.Type("kubernetes:core/v1:Namespace"), otherNamespace.URN.Type())
			nsName, _ := openapi.Pluck(otherNamespace.Outputs, "metadata", "name")
			assert.NotEqual(t, nsName.(string), providerNsName.(string))

			// Assert the first Pod was created in the provider default namespace.
			pod1 := stackInfo.Deployment.Resources[4]
			assert.Equal(t, "nginx1", string(pod1.URN.Name()))
			podNamespace1, _ := openapi.Pluck(pod1.Outputs, "metadata", "namespace")
			assert.Equal(t, providerNsName.(string), podNamespace1.(string))

			// Assert the second Pod was created in the provider default namespace.
			pod2 := stackInfo.Deployment.Resources[5]
			assert.Equal(t, "nginx2", string(pod2.URN.Name()))
			podNamespace2, _ := openapi.Pluck(pod2.Outputs, "metadata", "namespace")
			assert.Equal(t, providerNsName.(string), podNamespace2.(string))

			// Assert the Pod was created in the specified namespace rather than the provider default namespace.
			namespacedPod := stackInfo.Deployment.Resources[3]
			assert.Equal(t, "namespaced-nginx", string(namespacedPod.URN.Name()))
			namespacedPodNamespace, _ := openapi.Pluck(namespacedPod.Outputs, "metadata", "namespace")
			assert.NotEqual(t, providerNsName.(string), namespacedPodNamespace.(string))
			assert.Equal(t, ns2Name.(string), namespacedPodNamespace.(string))
		},
	})
	integration.ProgramTest(t, &test)
}

func TestQuery(t *testing.T) {
	test := baseOptions.With(integration.ProgramTestOptions{
		Dir:       filepath.Join("query", "step1"),
		Quick:     true,
		StackName: "query-test-c186bcc3-1572-44d8-b7d5-1028853682c3", // Chosen by fair dice roll. Guaranteed to be random.
		CloudURL:  "file://~",                                        // Required; we hard-code the stack name
		ExtraRuntimeValidation: func(t *testing.T, stackInfo integration.RuntimeValidationStackInfo) {
			assert.NotNil(t, stackInfo.Deployment)
			assert.Equal(t, 4, len(stackInfo.Deployment.Resources))

			tests.SortResourcesByURN(stackInfo)

			stackRes := stackInfo.Deployment.Resources[3]
			assert.Equal(t, resource.RootStackType, stackRes.URN.Type())

			provRes := stackInfo.Deployment.Resources[2]
			assert.True(t, providers.IsProviderType(provRes.URN.Type()))

			//
			// Assert Pod is successfully given a unique name by Pulumi.
			//

			pod := stackInfo.Deployment.Resources[1]
			assert.Equal(t, "query-test", string(pod.URN.Name()))
		},
		EditDirs: []integration.EditDir{
			{
				Dir:       filepath.Join("query", "step2"),
				Additive:  true,
				QueryMode: true,
				ExtraRuntimeValidation: func(t *testing.T, stackInfo integration.RuntimeValidationStackInfo) {
					//
					// Verify no resources were deleted.
					//
					assert.NotNil(t, stackInfo.Deployment)
					assert.Equal(t, 4, len(stackInfo.Deployment.Resources))

					tests.SortResourcesByURN(stackInfo)

					stackRes := stackInfo.Deployment.Resources[3]
					assert.Equal(t, resource.RootStackType, stackRes.URN.Type())

					provRes := stackInfo.Deployment.Resources[2]
					assert.True(t, providers.IsProviderType(provRes.URN.Type()))

					//
					// If we pass this point, the query did NOT throw an error, and is therefore
					// successful.
					//
				},
			},
		},
	})
	integration.ProgramTest(t, &test)
}

func TestRenderYAML(t *testing.T) {
	// Create a temporary directory to hold rendered YAML manifests.
	dir, err := ioutil.TempDir("", "")
	assert.NoError(t, err)
	defer os.RemoveAll(dir)

	test := baseOptions.With(integration.ProgramTestOptions{
		Config: map[string]string{
			"renderDir": dir,
		},
		Dir:   filepath.Join("render-yaml", "step1"),
		Quick: true,
		ExtraRuntimeValidation: func(t *testing.T, stackInfo integration.RuntimeValidationStackInfo) {
			assert.NotNil(t, stackInfo.Deployment)
			assert.Equal(t, 4, len(stackInfo.Deployment.Resources))

			// Verify that YAML directory was created.
			files, err := ioutil.ReadDir(dir)
			assert.NoError(t, err)
			assert.Equal(t, len(files), 2)

			// Verify that CRD manifest directory was created.
			files, err = ioutil.ReadDir(filepath.Join(dir, "0-crd"))
			assert.NoError(t, err)
			assert.Equal(t, len(files), 0)

			// Verify that manifest directory was created.
			files, err = ioutil.ReadDir(filepath.Join(dir, "1-manifest"))
			assert.NoError(t, err)
			assert.Equal(t, len(files), 2)
		},
	})
	integration.ProgramTest(t, &test)
}

func TestReplaceUnready(t *testing.T) {
	test := baseOptions.With(integration.ProgramTestOptions{
		Dir:                  filepath.Join("replace-unready", "step1"),
		Quick:                true,
		ExpectFailure:        true, // The Job is intended to fail.
		ExpectRefreshChanges: true,
		ExtraRuntimeValidation: func(t *testing.T, stackInfo integration.RuntimeValidationStackInfo) {
			assert.NotNil(t, stackInfo.Deployment)
			assert.Equal(t, 3, len(stackInfo.Deployment.Resources))

			tests.SortResourcesByURN(stackInfo)

			job := stackInfo.Deployment.Resources[0]
			provRes := stackInfo.Deployment.Resources[1]
			stackRes := stackInfo.Deployment.Resources[2]

			assert.Equal(t, resource.RootStackType, stackRes.URN.Type())
			assert.True(t, providers.IsProviderType(provRes.URN.Type()))

			assert.Equal(t, tokens.Type("kubernetes:batch/v1:Job"), job.URN.Type())
		},
		EditDirs: []integration.EditDir{
			{
				Dir:           filepath.Join("replace-unready", "step2"),
				Additive:      true,
				ExpectFailure: true,
				ExtraRuntimeValidation: func(t *testing.T, stackInfo integration.RuntimeValidationStackInfo) {
					assert.NotNil(t, stackInfo.Deployment)
					assert.Equal(t, 3, len(stackInfo.Deployment.Resources))

					tests.SortResourcesByURN(stackInfo)

					job := stackInfo.Deployment.Resources[0]
					provRes := stackInfo.Deployment.Resources[1]
					stackRes := stackInfo.Deployment.Resources[2]

					assert.Equal(t, resource.RootStackType, stackRes.URN.Type())
					assert.True(t, providers.IsProviderType(provRes.URN.Type()))

					assert.Equal(t, tokens.Type("kubernetes:batch/v1:Job"), job.URN.Type())

					// Check the event stream for a preview showing that the Job will be updated.
					for _, e := range stackInfo.Events {
						if e.ResourcePreEvent != nil && e.ResourcePreEvent.Metadata.Type == "kubernetes:batch/v1:Job" {
							assert.Equal(t, e.ResourcePreEvent.Metadata.Op, apitype.OpUpdate)
						}
					}
				},
			},
			{
				Dir:           filepath.Join("replace-unready", "step3"),
				Additive:      true,
				ExpectFailure: true,
				ExtraRuntimeValidation: func(t *testing.T, stackInfo integration.RuntimeValidationStackInfo) {
					// The stack has an extra Job now from the failed update in step2.
					assert.NotNil(t, stackInfo.Deployment)
					assert.Equal(t, 4, len(stackInfo.Deployment.Resources))

					tests.SortResourcesByURN(stackInfo)

					job := stackInfo.Deployment.Resources[0]
					jobOld := stackInfo.Deployment.Resources[1]
					provRes := stackInfo.Deployment.Resources[2]
					stackRes := stackInfo.Deployment.Resources[3]

					assert.Equal(t, resource.RootStackType, stackRes.URN.Type())
					assert.True(t, providers.IsProviderType(provRes.URN.Type()))

					assert.Equal(t, tokens.Type("kubernetes:batch/v1:Job"), job.URN.Type())
					assert.Equal(t, tokens.Type("kubernetes:batch/v1:Job"), jobOld.URN.Type())

					// Check the event stream for a preview showing that the Job will be replaced.
					for _, e := range stackInfo.Events {
						if e.ResourcePreEvent != nil && e.ResourcePreEvent.Metadata.Type == "kubernetes:batch/v1:Job" {
							assert.Equal(t, e.ResourcePreEvent.Metadata.Op, apitype.OpCreateReplacement)
						}
					}
				},
			},
		},
	})
	integration.ProgramTest(t, &test)
}

func TestRetry(t *testing.T) {
	test := baseOptions.With(integration.ProgramTestOptions{
		Dir:   filepath.Join("retry", "step1"),
		Quick: true,
		ExtraRuntimeValidation: func(t *testing.T, stackInfo integration.RuntimeValidationStackInfo) {
			assert.NotNil(t, stackInfo.Deployment)
			assert.Equal(t, 6, len(stackInfo.Deployment.Resources))

			tests.SortResourcesByURN(stackInfo)

			stackRes := stackInfo.Deployment.Resources[4]
			assert.Equal(t, resource.RootStackType, stackRes.URN.Type())

			provRes := stackInfo.Deployment.Resources[2]
			assert.True(t, providers.IsProviderType(provRes.URN.Type()))

			// Assert the Namespace was created
			namespace := stackInfo.Deployment.Resources[0]
			assert.Equal(t, tokens.Type("kubernetes:core/v1:Namespace"), namespace.URN.Type())

			// Assert the Pod was created
			pod := stackInfo.Deployment.Resources[1]
			assert.Equal(t, "nginx", string(pod.URN.Name()))
			step1Name, _ := openapi.Pluck(pod.Outputs, "metadata", "name")
			assert.Equal(t, "nginx", step1Name.(string))
			step1PodNamespace, _ := openapi.Pluck(pod.Outputs, "metadata", "namespace")
			assert.Equal(t, namespace.ID.String(), step1PodNamespace.(string))
		},
	})
	integration.ProgramTest(t, &test)
}

func TestSecrets(t *testing.T) {
	secretMessage := "secret message for testing"

	test := baseOptions.With(integration.ProgramTestOptions{
		Dir: filepath.Join("secrets", "step1"),
		Config: map[string]string{
			"message": secretMessage,
		},
		OrderedConfig: []integration.ConfigValue{
			{
				Key:   "pulumi:disable-default-providers[0]",
				Value: "kubernetes",
				Path:  true,
			},
		},
		Quick: true,
		ExtraRuntimeValidation: func(t *testing.T, stackInfo integration.RuntimeValidationStackInfo) {
			assert.NotNil(t, stackInfo.Deployment)
			state, err := json.Marshal(stackInfo.Deployment)
			assert.NoError(t, err)

			ssStringDataData, ok := stackInfo.Outputs["ssStringDataData"]
			assert.Truef(t, ok, "missing expected output \"ssStringDataData\"")

			ssStringDataStringData, ok := stackInfo.Outputs["ssStringDataStringData"]
			assert.Truef(t, ok, "missing expected output \"ssStringDataStringData\"")

			assert.NotEmptyf(t, ssStringDataData, "data field is empty")
			assert.NotEmptyf(t, ssStringDataStringData, "stringData field is empty")

			assert.NotContains(t, string(state), secretMessage)

			// The program converts the secret message to base64, to make a ConfigMap from it, so the state
			// should also not contain the base64 encoding of secret message.
			assert.NotContains(t, string(state), b64.StdEncoding.EncodeToString([]byte(secretMessage)))
		},
		EditDirs: []integration.EditDir{
			{
				Dir:      filepath.Join("secrets", "step2"),
				Additive: true,
				ExtraRuntimeValidation: func(t *testing.T, stackInfo integration.RuntimeValidationStackInfo) {
					secretMessage += "updated"

					assert.NotNil(t, stackInfo.Deployment)
					state, err := json.Marshal(stackInfo.Deployment)
					assert.NoError(t, err)

					ssStringDataData, ok := stackInfo.Outputs["ssStringDataData"]
					assert.Truef(t, ok, "missing expected output \"ssStringDataData\"")

					ssStringDataStringData, ok := stackInfo.Outputs["ssStringDataStringData"]
					assert.Truef(t, ok, "missing expected output \"ssStringDataStringData\"")

					assert.NotEmptyf(t, ssStringDataData, "data field is empty")
					assert.NotEmptyf(t, ssStringDataStringData, "stringData field is empty")

					assert.NotContains(t, string(state), secretMessage)

					// The program converts the secret message to base64, to make a ConfigMap from it, so the state
					// should also not contain the base64 encoding of secret message.
					assert.NotContains(t, string(state), b64.StdEncoding.EncodeToString([]byte(secretMessage)))
				},
			},
		},
	})
	integration.ProgramTest(t, &test)
}

func TestServerSideApply(t *testing.T) {
	test := baseOptions.With(integration.ProgramTestOptions{
		Dir:                  filepath.Join("server-side-apply", "step1"),
		ExpectRefreshChanges: true,
		OrderedConfig: []integration.ConfigValue{
			{
				Key:   "pulumi:disable-default-providers[0]",
				Value: "kubernetes",
				Path:  true,
			},
		},
		ExtraRuntimeValidation: func(t *testing.T, stackInfo integration.RuntimeValidationStackInfo) {
			// Validate patched Namespace
			nsPatched := stackInfo.Outputs["nsPatched"].(map[string]any)
			fooV, ok, err := unstructured.NestedString(nsPatched, "metadata", "labels", "foo")
			assert.True(t, ok)
			assert.NoError(t, err)
			assert.Equal(t, "foo", fooV)

			// Validate patched CustomResource
			crPatched := stackInfo.Outputs["crPatched"].(map[string]any)
			fooV, ok, err = unstructured.NestedString(crPatched, "metadata", "labels", "foo")
			assert.True(t, ok)
			assert.NoError(t, err)
			assert.Equal(t, "foo", fooV)

			for _, res := range stackInfo.Deployment.Resources {
				// Validate that the last-applied-configuration annotation is not present on SSA resources.
				annotations, ok, err := unstructured.NestedStringMap(res.Outputs, "metadata", "labels")
				assert.NoError(t, err)
				if ok {
					assert.NotContains(t, annotations, "kubectl.kubernetes.io/last-applied-configuration")
				}

				// Validate that the managed-by label is not present on SSA resources.
				labels, ok, err := unstructured.NestedStringMap(res.Outputs, "metadata", "labels")
				assert.NoError(t, err)
				if ok {
					assert.NotContains(t, labels, "app.kubernetes.io/managed-by")
				}

				if res.Type == "kubernetes:core/v1:ConfigMap" {
					dataV, ok, err := unstructured.NestedString(res.Outputs, "data", "foo")
					assert.True(t, ok)
					assert.NoError(t, err)
					assert.Equal(t, "bar", dataV)
				}
			}
		},
		EditDirs: []integration.EditDir{
			{
				Dir:      filepath.Join("server-side-apply", "step2"),
				Additive: true,
				ExtraRuntimeValidation: func(t *testing.T, stackInfo integration.RuntimeValidationStackInfo) {
					// Validate patched Deployment
					deploymentPatched := stackInfo.Outputs["deploymentPatched"].(map[string]any)
					containersV, ok, err := unstructured.NestedSlice(
						deploymentPatched, "spec", "template", "spec", "containers")
					assert.True(t, ok)
					assert.NoError(t, err)
					assert.Len(t, containersV, 1)
					limitsV, ok, err := unstructured.NestedMap(
						containersV[0].(map[string]any), "resources", "limits")
					assert.True(t, ok)
					assert.NoError(t, err)
					assert.Contains(t, limitsV, "memory")

					for _, res := range stackInfo.Deployment.Resources {
						if res.Type == "kubernetes:core/v1:ConfigMap" {
							dataV, ok, err := unstructured.NestedString(res.Outputs, "data", "foo")
							assert.True(t, ok)
							assert.NoError(t, err)
							assert.Equal(t, "baz", dataV)
						}
					}
				},
			},
			{
				Dir:      filepath.Join("server-side-apply", "step3"),
				Additive: true,
				ExtraRuntimeValidation: func(t *testing.T, stackInfo integration.RuntimeValidationStackInfo) {
					for _, res := range stackInfo.Deployment.Resources {
						if res.Type == "kubernetes:core/v1:ConfigMap" {
							dataV, ok, err := unstructured.NestedString(res.Outputs, "data", "foo")
							assert.True(t, ok)
							assert.NoError(t, err)
							assert.Equal(t, "baz", dataV) // Data should be unchanged from step2.
						}
					}
				},
			},
		},
	})
	integration.ProgramTest(t, &test)
}

// TestServerSideApplyEmptyMaps tests that we correctly handle merging structs containing empty maps when diffing live and wanted
// states. This is a regression test for issue #2332 to ensure Pulumi can handle updating a resource which has a
// map field that is empty in the live state but non-empty in the wanted state.
func TestServerSideApplyEmptyMaps(t *testing.T) {
	var ns, cmName string

	applyStep := baseOptions.With(integration.ProgramTestOptions{
		Dir:                  filepath.Join("server-side-apply-empty-maps", "configmap"),
		ExpectRefreshChanges: true,
		// Enable destroy-on-cleanup so we can shell out to kubectl to make external changes to the resource and reuse the same stack.
		DestroyOnCleanup: true,
		Quick:            true,
		OrderedConfig: []integration.ConfigValue{
			{
				Key:   "pulumi:disable-default-providers[0]",
				Value: "kubernetes",
				Path:  true,
			},
		},
		ExtraRuntimeValidation: func(t *testing.T, stackInfo integration.RuntimeValidationStackInfo) {
			cm := stackInfo.Outputs["cm"].(map[string]any)
			// Save the name and namespace for later use with kubectl. We check that the vars are empty,
			// in case pulumi up creates a new ConfigMap/Namespace instead of updating the existing one on
			// subsequent runs.
			if ns == "" && cmName == "" {
				ns = cm["metadata"].(map[string]any)["namespace"].(string)
				cmName = cm["metadata"].(map[string]any)["name"].(string)
			}

			// Validate we applied ConfigMap with wanted labels.
			fooV, ok, err := unstructured.NestedString(cm, "metadata", "labels", "foo")
			assert.True(t, ok)
			assert.NoError(t, err)
			assert.Equal(t, "bar", fooV)
		},
	})

	// Use manual lifecycle management since we need to run external commands in between pulumi up steps, while referencing
	// the same stack.
	pt := integration.ProgramTestManualLifeCycle(t, &applyStep)
	err := pt.TestLifeCycleInitAndDestroy()
	assert.NoError(t, err)

	// Sanity check with kubectl to verify that the ConfigMap was created with the wanted label.
	out, err := exec.Command("kubectl", "get", "configmap", "-o", "yaml", "-n", ns, cmName).CombinedOutput()
	assert.NoError(t, err)
	assert.Contains(t, string(out), "bar") // ConfigMap should have been created with label foo=bar.

	// Update the ConfigMap and remove label using kubectl.
	out, err = exec.Command("kubectl", "label", "configmap", "-n", ns, cmName, "foo-").CombinedOutput()
	assert.NoError(t, err)
	assert.Contains(t, string(out), "configmap/"+cmName+" unlabeled") // Ensure CM was unlabeled.

	// Use kubectl to verify that the ConfigMap was updated and no longer has the label.
	out, err = exec.Command("kubectl", "get", "configmap", "-o", "yaml", "-n", ns, cmName).CombinedOutput()
	assert.NoError(t, err)
	assert.NotContains(t, string(out), "bar") // ConfigMap should no longer have label foo=bar.

	// Run `pulumi up` + `pulumi refresh` to refresh the state and detect the missing label.
	// (The program tester runs these as separate steps, so the `pulumi up` doesn't detect a change until after the
	// subsequent refresh is performed.)
	err = pt.TestPreviewUpdateAndEdits()
	assert.NoError(t, err)

	// Re-run `pulumi up` to update the ConfigMap and re-add the label.
	err = pt.TestPreviewUpdateAndEdits()
	assert.NoError(t, err)

	// Use kubectl to verify that the ConfigMap was updated and has the label again.
	out, err = exec.Command("kubectl", "get", "configmap", "-o", "yaml", "-n", ns, cmName).CombinedOutput()
	assert.NoError(t, err)
	assert.Contains(t, string(out), "bar") // ConfigMap should have been updated with label foo=bar.
}

func TestServerSideApplyUpgrade(t *testing.T) {
	test := baseOptions.With(integration.ProgramTestOptions{
		Dir:                  filepath.Join("server-side-apply-upgrade", "step1"),
		ExpectRefreshChanges: true,
		OrderedConfig: []integration.ConfigValue{
			{
				Key:   "pulumi:disable-default-providers[0]",
				Value: "kubernetes",
				Path:  true,
			},
		},
		ExtraRuntimeValidation: func(t *testing.T, stackInfo integration.RuntimeValidationStackInfo) {
			// Validate Provider config
			provider := stackInfo.Outputs["provider"].(map[string]any)
			enableSSA, ok, err := unstructured.NestedString(provider, "enableServerSideApply")
			assert.True(t, ok)
			assert.NoError(t, err)
			assert.Equalf(t, "false", enableSSA, "SSA should be disabled")
			for _, res := range stackInfo.Deployment.Resources {
				if res.Type == "kubernetes:apps/v1:Deployment" {
					live := &unstructured.Unstructured{Object: res.Outputs}
					foundExpectedManager := false
					for _, managedField := range live.GetManagedFields() {
						if managedField.Manager == "pulumi-kubernetes" && managedField.Operation == "Update" {
							foundExpectedManager = true
							break
						}
					}
					assert.Truef(t, foundExpectedManager, "missing expected pulumi-kubernetes field manager with operation type Update")
				}
			}
		},
		EditDirs: []integration.EditDir{
			{
				Dir:      filepath.Join("server-side-apply-upgrade", "step2"),
				Additive: true,
				ExtraRuntimeValidation: func(t *testing.T, stackInfo integration.RuntimeValidationStackInfo) {
					// Validate Provider config
					provider := stackInfo.Outputs["provider"].(map[string]any)
					enableSSA, ok, err := unstructured.NestedString(provider, "enableServerSideApply")
					assert.True(t, ok)
					assert.NoError(t, err)
					assert.Equalf(t, "true", enableSSA, "SSA should be enabled")
					for _, res := range stackInfo.Deployment.Resources {
						if res.Type == "kubernetes:apps/v1:Deployment" {
							live := &unstructured.Unstructured{Object: res.Outputs}
							foundExpectedManager := false
							for _, managedField := range live.GetManagedFields() {
								if managedField.Manager == "pulumi-kubernetes" && managedField.Operation == "Update" {
									foundExpectedManager = true
									break
								}
							}
							assert.Truef(t, foundExpectedManager, "missing expected pulumi-kubernetes field manager with operation type Update")
						}
					}
				},
			},
			{
				Dir:      filepath.Join("server-side-apply-upgrade", "step3"),
				Additive: true,
				ExtraRuntimeValidation: func(t *testing.T, stackInfo integration.RuntimeValidationStackInfo) {
					for _, res := range stackInfo.Deployment.Resources {
						if res.Type == "kubernetes:apps/v1:Deployment" {
							live := &unstructured.Unstructured{Object: res.Outputs}
							foundExpectedManager := false
							for _, managedField := range live.GetManagedFields() {
								// Operation type should change to Apply on first update to the Deployment.
								if managedField.Manager == "pulumi-kubernetes" && managedField.Operation == "Apply" {
									foundExpectedManager = true
									break
								}
							}
							assert.Truef(t, foundExpectedManager, "missing expected pulumi-kubernetes field manager with operation type Apply")
							containers, ok := openapi.Pluck(res.Outputs, "spec", "template", "spec", "containers")
							assert.Truef(t, ok, "failed to get containers")
							containerStatus := containers.([]any)[0].(map[string]any)
							image := containerStatus["image"]
							assert.Equalf(t, image.(string), "nginx:1.17", "image should be updated")
							portsRaw := containerStatus["ports"]
							portsArray := portsRaw.([]any)
							assert.Equalf(t, len(portsArray), 1, "only one port should be set")
							portMap := portsArray[0].(map[string]any)
							assert.Equalf(t, portMap["containerPort"], float64(81), "port should be updated to 81")
						}
					}
				},
			},
		},
	})
	integration.ProgramTest(t, &test)
}

func TestYAMLURL(t *testing.T) {
	test := baseOptions.With(integration.ProgramTestOptions{
		Dir:   filepath.Join("yaml-url", "step1"),
		Quick: true,
		OrderedConfig: []integration.ConfigValue{
			{
				Key:   "pulumi:disable-default-providers[0]",
				Value: "kubernetes",
				Path:  true,
			},
		},
		ExtraRuntimeValidation: func(t *testing.T, stackInfo integration.RuntimeValidationStackInfo) {
			assert.NotNil(t, stackInfo.Deployment)

			// Assert that we've retrieved the YAML from the URL and provisioned them.
			assert.Equal(t, 18, len(stackInfo.Deployment.Resources))
		},
	})
	integration.ProgramTest(t, &test)
}

func TestReplaceDaemonSet(t *testing.T) {
	daemonSetName := ""
	test := baseOptions.With(integration.ProgramTestOptions{
		Dir:           filepath.Join("replace-daemonset", "step1"),
		Quick:         true,
		ExpectFailure: false,
		SkipRefresh:   true,
		ExtraRuntimeValidation: func(t *testing.T, stackInfo integration.RuntimeValidationStackInfo) {
			assert.NotNil(t, stackInfo.Deployment)
			assert.Equal(t, 3, len(stackInfo.Deployment.Resources))

			// Save the DaemonSet name to compare it in the step2
			daemonSetName = stackInfo.Outputs["name"].(string)

			// Assert that the DaemonSet was created
			assert.True(t, strings.HasPrefix(stackInfo.Outputs["name"].(string), "test-replacement-"))
		},
		EditDirs: []integration.EditDir{
			{
				Dir:           filepath.Join("replace-daemonset", "step2"),
				Additive:      true,
				ExpectFailure: false,
				ExtraRuntimeValidation: func(t *testing.T, stackInfo integration.RuntimeValidationStackInfo) {
					assert.NotNil(t, stackInfo.Deployment)
					assert.Equal(t, 3, len(stackInfo.Deployment.Resources))

					newDaemonSetName := stackInfo.Outputs["name"].(string)

					// Assert that the DaemonSet still exists
					assert.True(t, strings.HasPrefix(newDaemonSetName, "test-replacement-"))

					// DaemonSet should have a different name as it was replaced
					assert.True(t, daemonSetName != newDaemonSetName)
				},
			},
		},
	})
	integration.ProgramTest(t, &test)
}

func TestServiceAccountTokenSecret(t *testing.T) {
	test := baseOptions.With(integration.ProgramTestOptions{
		Dir:           filepath.Join("service-account-token-secret", "step1"),
		Quick:         true,
		ExpectFailure: false,
		SkipRefresh:   true,
		ExtraRuntimeValidation: func(t *testing.T, stackInfo integration.RuntimeValidationStackInfo) {
			assert.NotNil(t, stackInfo.Deployment)
			_, err := json.Marshal(stackInfo.Deployment)
			assert.NoError(t, err)

			secretData := stackInfo.Outputs["data"].(map[string]any)

			assert.Contains(t, secretData, "ca.crt")
			assert.Contains(t, secretData, "token")
		},
	})
	integration.ProgramTest(t, &test)
}

func TestStrictMode(t *testing.T) {
	test := baseOptions.With(integration.ProgramTestOptions{
		Dir:           filepath.Join("strict-mode", "step1"),
		Quick:         true,
		ExpectFailure: true,
		SkipRefresh:   true,
		OrderedConfig: []integration.ConfigValue{
			{
				Key:   "kubernetes:strictMode",
				Value: "true",
			},
		},
		ExtraRuntimeValidation: func(t *testing.T, stackInfo integration.RuntimeValidationStackInfo) {
			// Check the event stream for a diagnostic event showing that a default provider is prohibited.
			foundMessage := false
			msg := "strict mode prohibits default provider"
			for _, e := range stackInfo.Events {
				if e.DiagnosticEvent != nil && strings.Contains(e.DiagnosticEvent.Message, msg) {
					foundMessage = true
					break
				}
			}
			assert.Truef(t, foundMessage, "did not find expected failure message: %q", msg)
		},
		EditDirs: []integration.EditDir{
			{
				Dir:           filepath.Join("strict-mode", "step2"),
				Additive:      true,
				ExpectFailure: true,
				ExtraRuntimeValidation: func(t *testing.T, stackInfo integration.RuntimeValidationStackInfo) {
					// Check the event stream for a diagnostic event showing that a Provider requires a "context".
					foundMessage := false
					msg := `strict mode requires Provider "context" argument`
					for _, e := range stackInfo.Events {
						if e.DiagnosticEvent != nil && strings.Contains(e.DiagnosticEvent.Message, msg) {
							foundMessage = true
							break
						}
					}
					assert.Truef(t, foundMessage, "did not find expected failure message: %q", msg)
					_, ok := stackInfo.Outputs["cm"]
					assert.Falsef(t, ok, "ConfigMap should not be present since Provider is invalid")
				},
			},
			{
				Dir:           filepath.Join("strict-mode", "step3"),
				Additive:      true,
				ExpectFailure: true,
				ExtraRuntimeValidation: func(t *testing.T, stackInfo integration.RuntimeValidationStackInfo) {
					// Check the event stream for a diagnostic event showing that a Provider requires a "kubeconfig".
					foundMessage := false
					msg := `strict mode requires Provider "kubeconfig" argument`
					for _, e := range stackInfo.Events {
						if e.DiagnosticEvent != nil && strings.Contains(e.DiagnosticEvent.Message, msg) {
							foundMessage = true
							break
						}
					}
					assert.Truef(t, foundMessage, "did not find expected failure message: %q", msg)
					_, ok := stackInfo.Outputs["cm"]
					assert.Falsef(t, ok, "ConfigMap should not be present since Provider is invalid")
				},
			},
		},
	})
	integration.ProgramTest(t, &test)
}

// TestClientSideDriftCorrectCSA tests that we can successfully reapply a resource that has been
// modified outside of Pulumi using the client-side apply strategy.
func TestClientSideDriftCorrectCSA(t *testing.T) {
	var ns, cmName string

	applyStep := baseOptions.With(integration.ProgramTestOptions{
		Dir:                  filepath.Join("drift-correct", "configmap-csa"),
		ExpectRefreshChanges: true,
		// Enable destroy-on-cleanup so we can shell out to kubectl to make external changes to the resource and reuse the same stack.
		DestroyOnCleanup: true,
		OrderedConfig: []integration.ConfigValue{
			{
				Key:   "pulumi:disable-default-providers[0]",
				Value: "kubernetes",
				Path:  true,
			},
		},
		ExtraRuntimeValidation: func(t *testing.T, stackInfo integration.RuntimeValidationStackInfo) {
			cm, ok := stackInfo.Outputs["cm"].(map[string]interface{})
			assert.True(t, ok)

			// Save the name and namespace for later use with kubectl. We check that the vars are empty,
			// in case pulumi up creates a new ConfigMap/Namespace instead of updating the existing one on
			// subsequent runs.
			if ns == "" && cmName == "" {
				ns = cm["metadata"].(map[string]interface{})["namespace"].(string)
				cmName = cm["metadata"].(map[string]interface{})["name"].(string)
			}

			// Validate we applied ConfigMap with data.
			fooV, ok, err := unstructured.NestedString(cm, "data", "foo")
			assert.True(t, ok)
			assert.NoError(t, err)
			assert.Equal(t, "bar", fooV)
		},
	})

	// Use manual lifecycle management since we need to run external commands in between pulumi up steps, while referencing
	// the same stack.
	pt := integration.ProgramTestManualLifeCycle(t, &applyStep)
	err := pt.TestLifeCycleInitAndDestroy()
	assert.NoError(t, err)

	// Sanity check with kubectl to verify that the ConfigMap was created with the wanted label.
	out, err := exec.Command("kubectl", "get", "configmap", "-o", "yaml", "-n", ns, cmName).CombinedOutput()
	assert.NoError(t, err)
	assert.Contains(t, string(out), "bar") // ConfigMap should have been created with data foo: bar.

	// Update the ConfigMap and change the data foo: bar to foo: baz.
	out, err = exec.Command("kubectl", "patch", "configmap", "-n", ns, cmName, "-p", `{"data":{"foo":"baz"}}`).CombinedOutput()
	assert.NoError(t, err)
	assert.Contains(t, string(out), "configmap/"+cmName+" patched") // Ensure CM was patched.

	// Use kubectl to verify that the ConfigMap was updated and now has data foo: baz.
	out, err = exec.Command("kubectl", "get", "configmap", "-o", "yaml", "-n", ns, cmName).CombinedOutput()
	assert.NoError(t, err)
	assert.NotContains(t, string(out), "foo: bar") // ConfigMap should no longer have data foo: bar.
	assert.Contains(t, string(out), "foo: baz")    // ConfigMap should have data foo: baz.

	// Run `pulumi up` + `pulumi refresh` to refresh the state and detect the missing label.
	// (The program tester runs these as separate steps, so the `pulumi up` doesn't detect a change until after the
	// subsequent refresh is performed.)
	err = pt.TestPreviewUpdateAndEdits()
	assert.NoError(t, err)

	// Re-run `pulumi up` to update the ConfigMap and re-add the label.
	err = pt.TestPreviewUpdateAndEdits()
	assert.NoError(t, err)

	// Use kubectl to verify that the ConfigMap was updated and has the label again.
	out, err = exec.Command("kubectl", "get", "configmap", "-o", "yaml", "-n", ns, cmName).CombinedOutput()
	assert.NoError(t, err)

	assert.Contains(t, string(out), "foo: bar")    // ConfigMap should have been updated with data foo: bar.
	assert.NotContains(t, string(out), "foo: baz") // ConfigMap should no longer have data foo: baz.
}

// TestClientSideDriftCorrectSSA tests that we can successfully reapply a resource that has been
// modified outside of Pulumi, with SSA enabled.
func TestClientSideDriftCorrectSSA(t *testing.T) {
	var ns, cmName string

	applyStep := baseOptions.With(integration.ProgramTestOptions{
		Dir:                  filepath.Join("drift-correct", "configmap-ssa"),
		ExpectRefreshChanges: true,
		// Enable destroy-on-cleanup so we can shell out to kubectl to make external changes to the resource and reuse the same stack.
		DestroyOnCleanup: true,
		OrderedConfig: []integration.ConfigValue{
			{
				Key:   "pulumi:disable-default-providers[0]",
				Value: "kubernetes",
				Path:  true,
			},
		},
		ExtraRuntimeValidation: func(t *testing.T, stackInfo integration.RuntimeValidationStackInfo) {
			cm, ok := stackInfo.Outputs["cm"].(map[string]interface{})
			assert.True(t, ok)

			// Save the name and namespace for later use with kubectl. We check that the vars are empty,
			// in case pulumi up creates a new ConfigMap/Namespace instead of updating the existing one on
			// subsequent runs.
			if ns == "" && cmName == "" {
				ns = cm["metadata"].(map[string]interface{})["namespace"].(string)
				cmName = cm["metadata"].(map[string]interface{})["name"].(string)
			}

			// Validate we applied ConfigMap with data.
			fooV, ok, err := unstructured.NestedString(cm, "data", "foo")
			assert.True(t, ok)
			assert.NoError(t, err)
			assert.Equal(t, "bar", fooV)
		},
	})

	// Use manual lifecycle management since we need to run external commands in between pulumi up steps, while referencing
	// the same stack.
	pt := integration.ProgramTestManualLifeCycle(t, &applyStep)
	err := pt.TestLifeCycleInitAndDestroy()
	assert.NoError(t, err)

	// Sanity check with kubectl to verify that the ConfigMap was created with the wanted label.
	out, err := exec.Command("kubectl", "get", "configmap", "-o", "yaml", "-n", ns, cmName).CombinedOutput()
	assert.NoError(t, err)
	assert.Contains(t, string(out), "bar") // ConfigMap should have been created with data foo: bar.

	// Update the ConfigMap and change the data foo: bar to foo: baz.
	out, err = exec.Command("kubectl", "patch", "configmap", "-n", ns, cmName, "-p", `{"data":{"foo":"baz"}}`).CombinedOutput()
	assert.NoError(t, err)
	assert.Contains(t, string(out), "configmap/"+cmName+" patched") // Ensure CM was patched.

	// Use kubectl to verify that the ConfigMap was updated and now has data foo: baz.
	out, err = exec.Command("kubectl", "get", "configmap", "-o", "yaml", "-n", ns, cmName).CombinedOutput()
	assert.NoError(t, err)
	assert.NotContains(t, string(out), "foo: bar") // ConfigMap should no longer have data foo: bar.
	assert.Contains(t, string(out), "foo: baz")    // ConfigMap should have data foo: baz.

	// Run `pulumi up` + `pulumi refresh` to refresh the state and detect the missing label.
	// (The program tester runs these as separate steps, so the `pulumi up` doesn't detect a change until after the
	// subsequent refresh is performed.)
	err = pt.TestPreviewUpdateAndEdits()
	assert.NoError(t, err)

	// Re-run `pulumi up` to update the ConfigMap and re-add the label.
	err = pt.TestPreviewUpdateAndEdits()
	assert.NoError(t, err)

	// Use kubectl to verify that the ConfigMap was updated and has the label again.
	out, err = exec.Command("kubectl", "get", "configmap", "-o", "yaml", "-n", ns, cmName).CombinedOutput()
	assert.NoError(t, err)

	assert.Contains(t, string(out), "foo: bar")    // ConfigMap should have been updated with data foo: bar.
	assert.NotContains(t, string(out), "foo: baz") // ConfigMap should no longer have data foo: baz.
}
