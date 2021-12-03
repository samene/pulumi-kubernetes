// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package helm

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// A Release is an instance of a chart running in a Kubernetes cluster.
// A Chart is a Helm package. It contains all the resource definitions necessary to run an application, tool, or service inside a Kubernetes cluster.
// Note - Helm Release is currently in BETA and may change. Use in production environment is discouraged.
type Release struct {
	pulumi.CustomResourceState

	// If set, installation process purges chart on fail. `skipAwait` will be disabled automatically if atomic is used.
	Atomic pulumi.BoolPtrOutput `pulumi:"atomic"`
	// Chart name to be installed. A path may be used.
	Chart pulumi.StringOutput `pulumi:"chart"`
	// Allow deletion of new resources created in this upgrade when upgrade fails.
	CleanupOnFail pulumi.BoolPtrOutput `pulumi:"cleanupOnFail"`
	// Create the namespace if it does not exist.
	CreateNamespace pulumi.BoolPtrOutput `pulumi:"createNamespace"`
	// Run helm dependency update before installing the chart.
	DependencyUpdate pulumi.BoolPtrOutput `pulumi:"dependencyUpdate"`
	// Add a custom description
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// Use chart development versions, too. Equivalent to version '>0.0.0-0'. If `version` is set, this is ignored.
	Devel pulumi.BoolPtrOutput `pulumi:"devel"`
	// Prevent CRD hooks from, running, but run other hooks.  See helm install --no-crd-hook
	DisableCRDHooks pulumi.BoolPtrOutput `pulumi:"disableCRDHooks"`
	// If set, the installation process will not validate rendered templates against the Kubernetes OpenAPI Schema
	DisableOpenapiValidation pulumi.BoolPtrOutput `pulumi:"disableOpenapiValidation"`
	// Prevent hooks from running.
	DisableWebhooks pulumi.BoolPtrOutput `pulumi:"disableWebhooks"`
	// Force resource update through delete/recreate if needed.
	ForceUpdate pulumi.BoolPtrOutput `pulumi:"forceUpdate"`
	// Location of public keys used for verification. Used only if `verify` is true
	Keyring pulumi.StringPtrOutput `pulumi:"keyring"`
	// Run helm lint when planning.
	Lint pulumi.BoolPtrOutput `pulumi:"lint"`
	// The rendered manifests as JSON. Not yet supported.
	Manifest pulumi.MapOutput `pulumi:"manifest"`
	// Limit the maximum number of revisions saved per release. Use 0 for no limit.
	MaxHistory pulumi.IntPtrOutput `pulumi:"maxHistory"`
	// Release name.
	Name pulumi.StringPtrOutput `pulumi:"name"`
	// Namespace to install the release into.
	Namespace pulumi.StringPtrOutput `pulumi:"namespace"`
	// Postrender command to run.
	Postrender pulumi.StringPtrOutput `pulumi:"postrender"`
	// Perform pods restart during upgrade/rollback.
	RecreatePods pulumi.BoolPtrOutput `pulumi:"recreatePods"`
	// If set, render subchart notes along with the parent.
	RenderSubchartNotes pulumi.BoolPtrOutput `pulumi:"renderSubchartNotes"`
	// Re-use the given name, even if that name is already used. This is unsafe in production
	Replace pulumi.BoolPtrOutput `pulumi:"replace"`
	// Specification defining the Helm chart repository to use.
	RepositoryOpts RepositoryOptsPtrOutput `pulumi:"repositoryOpts"`
	// When upgrading, reset the values to the ones built into the chart.
	ResetValues pulumi.BoolPtrOutput `pulumi:"resetValues"`
	// Names of resources created by the release grouped by "kind/version".
	ResourceNames pulumi.StringArrayMapOutput `pulumi:"resourceNames"`
	// When upgrading, reuse the last release's values and merge in any overrides. If 'resetValues' is specified, this is ignored
	ReuseValues pulumi.BoolPtrOutput `pulumi:"reuseValues"`
	// By default, the provider waits until all resources are in a ready state before marking the release as successful. Setting this to true will skip such await logic.
	SkipAwait pulumi.BoolPtrOutput `pulumi:"skipAwait"`
	// If set, no CRDs will be installed. By default, CRDs are installed if not already present.
	SkipCrds pulumi.BoolPtrOutput `pulumi:"skipCrds"`
	// Status of the deployed release.
	Status ReleaseStatusOutput `pulumi:"status"`
	// Time in seconds to wait for any individual kubernetes operation.
	Timeout pulumi.IntPtrOutput `pulumi:"timeout"`
	// List of assets (raw yaml files). Content is read and merged with values. Not yet supported.
	ValueYamlFiles pulumi.AssetOrArchiveArrayOutput `pulumi:"valueYamlFiles"`
	// Custom values set for the release.
	Values pulumi.MapOutput `pulumi:"values"`
	// Verify the package before installing it.
	Verify pulumi.BoolPtrOutput `pulumi:"verify"`
	// Specify the exact chart version to install. If this is not specified, the latest version is installed.
	Version pulumi.StringPtrOutput `pulumi:"version"`
	// Will wait until all Jobs have been completed before marking the release as successful. This is ignored if `skipAwait` is enabled.
	WaitForJobs pulumi.BoolPtrOutput `pulumi:"waitForJobs"`
}

// NewRelease registers a new resource with the given unique name, arguments, and options.
func NewRelease(ctx *pulumi.Context,
	name string, args *ReleaseArgs, opts ...pulumi.ResourceOption) (*Release, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Chart == nil {
		return nil, errors.New("invalid value for required argument 'Chart'")
	}
	args.Compat = pulumi.StringPtr("true")
	var resource Release
	err := ctx.RegisterResource("kubernetes:helm.sh/v3:Release", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetRelease gets an existing Release resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetRelease(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ReleaseState, opts ...pulumi.ResourceOption) (*Release, error) {
	var resource Release
	err := ctx.ReadResource("kubernetes:helm.sh/v3:Release", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Release resources.
type releaseState struct {
}

type ReleaseState struct {
}

func (ReleaseState) ElementType() reflect.Type {
	return reflect.TypeOf((*releaseState)(nil)).Elem()
}

type releaseArgs struct {
	// If set, installation process purges chart on fail. `skipAwait` will be disabled automatically if atomic is used.
	Atomic *bool `pulumi:"atomic"`
	// Chart name to be installed. A path may be used.
	Chart string `pulumi:"chart"`
	// Allow deletion of new resources created in this upgrade when upgrade fails.
	CleanupOnFail *bool   `pulumi:"cleanupOnFail"`
	Compat        *string `pulumi:"compat"`
	// Create the namespace if it does not exist.
	CreateNamespace *bool `pulumi:"createNamespace"`
	// Run helm dependency update before installing the chart.
	DependencyUpdate *bool `pulumi:"dependencyUpdate"`
	// Add a custom description
	Description *string `pulumi:"description"`
	// Use chart development versions, too. Equivalent to version '>0.0.0-0'. If `version` is set, this is ignored.
	Devel *bool `pulumi:"devel"`
	// Prevent CRD hooks from, running, but run other hooks.  See helm install --no-crd-hook
	DisableCRDHooks *bool `pulumi:"disableCRDHooks"`
	// If set, the installation process will not validate rendered templates against the Kubernetes OpenAPI Schema
	DisableOpenapiValidation *bool `pulumi:"disableOpenapiValidation"`
	// Prevent hooks from running.
	DisableWebhooks *bool `pulumi:"disableWebhooks"`
	// Force resource update through delete/recreate if needed.
	ForceUpdate *bool `pulumi:"forceUpdate"`
	// Location of public keys used for verification. Used only if `verify` is true
	Keyring *string `pulumi:"keyring"`
	// Run helm lint when planning.
	Lint *bool `pulumi:"lint"`
	// The rendered manifests as JSON. Not yet supported.
	Manifest map[string]interface{} `pulumi:"manifest"`
	// Limit the maximum number of revisions saved per release. Use 0 for no limit.
	MaxHistory *int `pulumi:"maxHistory"`
	// Release name.
	Name *string `pulumi:"name"`
	// Namespace to install the release into.
	Namespace *string `pulumi:"namespace"`
	// Postrender command to run.
	Postrender *string `pulumi:"postrender"`
	// Perform pods restart during upgrade/rollback.
	RecreatePods *bool `pulumi:"recreatePods"`
	// If set, render subchart notes along with the parent.
	RenderSubchartNotes *bool `pulumi:"renderSubchartNotes"`
	// Re-use the given name, even if that name is already used. This is unsafe in production
	Replace *bool `pulumi:"replace"`
	// Specification defining the Helm chart repository to use.
	RepositoryOpts *RepositoryOpts `pulumi:"repositoryOpts"`
	// When upgrading, reset the values to the ones built into the chart.
	ResetValues *bool `pulumi:"resetValues"`
	// Names of resources created by the release grouped by "kind/version".
	ResourceNames map[string][]string `pulumi:"resourceNames"`
	// When upgrading, reuse the last release's values and merge in any overrides. If 'resetValues' is specified, this is ignored
	ReuseValues *bool `pulumi:"reuseValues"`
	// By default, the provider waits until all resources are in a ready state before marking the release as successful. Setting this to true will skip such await logic.
	SkipAwait *bool `pulumi:"skipAwait"`
	// If set, no CRDs will be installed. By default, CRDs are installed if not already present.
	SkipCrds *bool `pulumi:"skipCrds"`
	// Time in seconds to wait for any individual kubernetes operation.
	Timeout *int `pulumi:"timeout"`
	// List of assets (raw yaml files). Content is read and merged with values. Not yet supported.
	ValueYamlFiles []pulumi.AssetOrArchive `pulumi:"valueYamlFiles"`
	// Custom values set for the release.
	Values map[string]interface{} `pulumi:"values"`
	// Verify the package before installing it.
	Verify *bool `pulumi:"verify"`
	// Specify the exact chart version to install. If this is not specified, the latest version is installed.
	Version *string `pulumi:"version"`
	// Will wait until all Jobs have been completed before marking the release as successful. This is ignored if `skipAwait` is enabled.
	WaitForJobs *bool `pulumi:"waitForJobs"`
}

// The set of arguments for constructing a Release resource.
type ReleaseArgs struct {
	// If set, installation process purges chart on fail. `skipAwait` will be disabled automatically if atomic is used.
	Atomic pulumi.BoolPtrInput
	// Chart name to be installed. A path may be used.
	Chart pulumi.StringInput
	// Allow deletion of new resources created in this upgrade when upgrade fails.
	CleanupOnFail pulumi.BoolPtrInput
	Compat        pulumi.StringPtrInput
	// Create the namespace if it does not exist.
	CreateNamespace pulumi.BoolPtrInput
	// Run helm dependency update before installing the chart.
	DependencyUpdate pulumi.BoolPtrInput
	// Add a custom description
	Description pulumi.StringPtrInput
	// Use chart development versions, too. Equivalent to version '>0.0.0-0'. If `version` is set, this is ignored.
	Devel pulumi.BoolPtrInput
	// Prevent CRD hooks from, running, but run other hooks.  See helm install --no-crd-hook
	DisableCRDHooks pulumi.BoolPtrInput
	// If set, the installation process will not validate rendered templates against the Kubernetes OpenAPI Schema
	DisableOpenapiValidation pulumi.BoolPtrInput
	// Prevent hooks from running.
	DisableWebhooks pulumi.BoolPtrInput
	// Force resource update through delete/recreate if needed.
	ForceUpdate pulumi.BoolPtrInput
	// Location of public keys used for verification. Used only if `verify` is true
	Keyring pulumi.StringPtrInput
	// Run helm lint when planning.
	Lint pulumi.BoolPtrInput
	// The rendered manifests as JSON. Not yet supported.
	Manifest pulumi.MapInput
	// Limit the maximum number of revisions saved per release. Use 0 for no limit.
	MaxHistory pulumi.IntPtrInput
	// Release name.
	Name pulumi.StringPtrInput
	// Namespace to install the release into.
	Namespace pulumi.StringPtrInput
	// Postrender command to run.
	Postrender pulumi.StringPtrInput
	// Perform pods restart during upgrade/rollback.
	RecreatePods pulumi.BoolPtrInput
	// If set, render subchart notes along with the parent.
	RenderSubchartNotes pulumi.BoolPtrInput
	// Re-use the given name, even if that name is already used. This is unsafe in production
	Replace pulumi.BoolPtrInput
	// Specification defining the Helm chart repository to use.
	RepositoryOpts RepositoryOptsPtrInput
	// When upgrading, reset the values to the ones built into the chart.
	ResetValues pulumi.BoolPtrInput
	// Names of resources created by the release grouped by "kind/version".
	ResourceNames pulumi.StringArrayMapInput
	// When upgrading, reuse the last release's values and merge in any overrides. If 'resetValues' is specified, this is ignored
	ReuseValues pulumi.BoolPtrInput
	// By default, the provider waits until all resources are in a ready state before marking the release as successful. Setting this to true will skip such await logic.
	SkipAwait pulumi.BoolPtrInput
	// If set, no CRDs will be installed. By default, CRDs are installed if not already present.
	SkipCrds pulumi.BoolPtrInput
	// Time in seconds to wait for any individual kubernetes operation.
	Timeout pulumi.IntPtrInput
	// List of assets (raw yaml files). Content is read and merged with values. Not yet supported.
	ValueYamlFiles pulumi.AssetOrArchiveArrayInput
	// Custom values set for the release.
	Values pulumi.MapInput
	// Verify the package before installing it.
	Verify pulumi.BoolPtrInput
	// Specify the exact chart version to install. If this is not specified, the latest version is installed.
	Version pulumi.StringPtrInput
	// Will wait until all Jobs have been completed before marking the release as successful. This is ignored if `skipAwait` is enabled.
	WaitForJobs pulumi.BoolPtrInput
}

func (ReleaseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*releaseArgs)(nil)).Elem()
}

type ReleaseInput interface {
	pulumi.Input

	ToReleaseOutput() ReleaseOutput
	ToReleaseOutputWithContext(ctx context.Context) ReleaseOutput
}

func (*Release) ElementType() reflect.Type {
	return reflect.TypeOf((**Release)(nil)).Elem()
}

func (i *Release) ToReleaseOutput() ReleaseOutput {
	return i.ToReleaseOutputWithContext(context.Background())
}

func (i *Release) ToReleaseOutputWithContext(ctx context.Context) ReleaseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ReleaseOutput)
}

// ReleaseArrayInput is an input type that accepts ReleaseArray and ReleaseArrayOutput values.
// You can construct a concrete instance of `ReleaseArrayInput` via:
//
//          ReleaseArray{ ReleaseArgs{...} }
type ReleaseArrayInput interface {
	pulumi.Input

	ToReleaseArrayOutput() ReleaseArrayOutput
	ToReleaseArrayOutputWithContext(context.Context) ReleaseArrayOutput
}

type ReleaseArray []ReleaseInput

func (ReleaseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Release)(nil)).Elem()
}

func (i ReleaseArray) ToReleaseArrayOutput() ReleaseArrayOutput {
	return i.ToReleaseArrayOutputWithContext(context.Background())
}

func (i ReleaseArray) ToReleaseArrayOutputWithContext(ctx context.Context) ReleaseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ReleaseArrayOutput)
}

// ReleaseMapInput is an input type that accepts ReleaseMap and ReleaseMapOutput values.
// You can construct a concrete instance of `ReleaseMapInput` via:
//
//          ReleaseMap{ "key": ReleaseArgs{...} }
type ReleaseMapInput interface {
	pulumi.Input

	ToReleaseMapOutput() ReleaseMapOutput
	ToReleaseMapOutputWithContext(context.Context) ReleaseMapOutput
}

type ReleaseMap map[string]ReleaseInput

func (ReleaseMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Release)(nil)).Elem()
}

func (i ReleaseMap) ToReleaseMapOutput() ReleaseMapOutput {
	return i.ToReleaseMapOutputWithContext(context.Background())
}

func (i ReleaseMap) ToReleaseMapOutputWithContext(ctx context.Context) ReleaseMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ReleaseMapOutput)
}

type ReleaseOutput struct{ *pulumi.OutputState }

func (ReleaseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Release)(nil)).Elem()
}

func (o ReleaseOutput) ToReleaseOutput() ReleaseOutput {
	return o
}

func (o ReleaseOutput) ToReleaseOutputWithContext(ctx context.Context) ReleaseOutput {
	return o
}

type ReleaseArrayOutput struct{ *pulumi.OutputState }

func (ReleaseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Release)(nil)).Elem()
}

func (o ReleaseArrayOutput) ToReleaseArrayOutput() ReleaseArrayOutput {
	return o
}

func (o ReleaseArrayOutput) ToReleaseArrayOutputWithContext(ctx context.Context) ReleaseArrayOutput {
	return o
}

func (o ReleaseArrayOutput) Index(i pulumi.IntInput) ReleaseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Release {
		return vs[0].([]*Release)[vs[1].(int)]
	}).(ReleaseOutput)
}

type ReleaseMapOutput struct{ *pulumi.OutputState }

func (ReleaseMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Release)(nil)).Elem()
}

func (o ReleaseMapOutput) ToReleaseMapOutput() ReleaseMapOutput {
	return o
}

func (o ReleaseMapOutput) ToReleaseMapOutputWithContext(ctx context.Context) ReleaseMapOutput {
	return o
}

func (o ReleaseMapOutput) MapIndex(k pulumi.StringInput) ReleaseOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Release {
		return vs[0].(map[string]*Release)[vs[1].(string)]
	}).(ReleaseOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ReleaseInput)(nil)).Elem(), &Release{})
	pulumi.RegisterInputType(reflect.TypeOf((*ReleaseArrayInput)(nil)).Elem(), ReleaseArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ReleaseMapInput)(nil)).Elem(), ReleaseMap{})
	pulumi.RegisterOutputType(ReleaseOutput{})
	pulumi.RegisterOutputType(ReleaseArrayOutput{})
	pulumi.RegisterOutputType(ReleaseMapOutput{})
}
