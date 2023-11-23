// Code generated by pulumigen DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v1beta1

import (
	"context"
	"reflect"

	metav1 "github.com/pulumi/pulumi-kubernetes/sdk/v4/go/kubernetes/meta/v1"
	"github.com/pulumi/pulumi-kubernetes/sdk/v4/go/kubernetes/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Patch resources are used to modify existing Kubernetes resources by using
// Server-Side Apply updates. The name of the resource must be specified, but all other properties are optional. More than
// one patch may be applied to the same resource, and a random FieldManager name will be used for each Patch resource.
// Conflicts will result in an error by default, but can be forced using the "pulumi.com/patchForce" annotation. See the
// [Server-Side Apply Docs](https://www.pulumi.com/registry/packages/kubernetes/how-to-guides/managing-resources-with-server-side-apply/) for
// additional information about using Server-Side Apply to manage Kubernetes resources with Pulumi.
// ValidatingAdmissionPolicyBinding binds the ValidatingAdmissionPolicy with paramerized resources. ValidatingAdmissionPolicyBinding and parameter CRDs together define how cluster administrators configure policies for clusters.
//
// For a given admission request, each binding will cause its policy to be evaluated N times, where N is 1 for policies/bindings that don't use params, otherwise N is the number of parameters selected by the binding.
//
// The CEL expressions of a policy must have a computed CEL cost below the maximum CEL budget. Each evaluation of the policy is given an independent CEL cost budget. Adding/removing policies, bindings, or params can not affect whether a given (policy, binding, param) combination is within its own CEL budget.
type ValidatingAdmissionPolicyBindingPatch struct {
	pulumi.CustomResourceState

	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion pulumi.StringPtrOutput `pulumi:"apiVersion"`
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind pulumi.StringPtrOutput `pulumi:"kind"`
	// Standard object metadata; More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata.
	Metadata metav1.ObjectMetaPatchPtrOutput `pulumi:"metadata"`
	// Specification of the desired behavior of the ValidatingAdmissionPolicyBinding.
	Spec ValidatingAdmissionPolicyBindingSpecPatchPtrOutput `pulumi:"spec"`
}

// NewValidatingAdmissionPolicyBindingPatch registers a new resource with the given unique name, arguments, and options.
func NewValidatingAdmissionPolicyBindingPatch(ctx *pulumi.Context,
	name string, args *ValidatingAdmissionPolicyBindingPatchArgs, opts ...pulumi.ResourceOption) (*ValidatingAdmissionPolicyBindingPatch, error) {
	if args == nil {
		args = &ValidatingAdmissionPolicyBindingPatchArgs{}
	}

	args.ApiVersion = pulumi.StringPtr("admissionregistration.k8s.io/v1beta1")
	args.Kind = pulumi.StringPtr("ValidatingAdmissionPolicyBinding")
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("kubernetes:admissionregistration.k8s.io/v1alpha1:ValidatingAdmissionPolicyBindingPatch"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource ValidatingAdmissionPolicyBindingPatch
	err := ctx.RegisterResource("kubernetes:admissionregistration.k8s.io/v1beta1:ValidatingAdmissionPolicyBindingPatch", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetValidatingAdmissionPolicyBindingPatch gets an existing ValidatingAdmissionPolicyBindingPatch resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetValidatingAdmissionPolicyBindingPatch(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ValidatingAdmissionPolicyBindingPatchState, opts ...pulumi.ResourceOption) (*ValidatingAdmissionPolicyBindingPatch, error) {
	var resource ValidatingAdmissionPolicyBindingPatch
	err := ctx.ReadResource("kubernetes:admissionregistration.k8s.io/v1beta1:ValidatingAdmissionPolicyBindingPatch", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ValidatingAdmissionPolicyBindingPatch resources.
type validatingAdmissionPolicyBindingPatchState struct {
}

type ValidatingAdmissionPolicyBindingPatchState struct {
}

func (ValidatingAdmissionPolicyBindingPatchState) ElementType() reflect.Type {
	return reflect.TypeOf((*validatingAdmissionPolicyBindingPatchState)(nil)).Elem()
}

type validatingAdmissionPolicyBindingPatchArgs struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion *string `pulumi:"apiVersion"`
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind *string `pulumi:"kind"`
	// Standard object metadata; More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata.
	Metadata *metav1.ObjectMetaPatch `pulumi:"metadata"`
	// Specification of the desired behavior of the ValidatingAdmissionPolicyBinding.
	Spec *ValidatingAdmissionPolicyBindingSpecPatch `pulumi:"spec"`
}

// The set of arguments for constructing a ValidatingAdmissionPolicyBindingPatch resource.
type ValidatingAdmissionPolicyBindingPatchArgs struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion pulumi.StringPtrInput
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind pulumi.StringPtrInput
	// Standard object metadata; More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata.
	Metadata metav1.ObjectMetaPatchPtrInput
	// Specification of the desired behavior of the ValidatingAdmissionPolicyBinding.
	Spec ValidatingAdmissionPolicyBindingSpecPatchPtrInput
}

func (ValidatingAdmissionPolicyBindingPatchArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*validatingAdmissionPolicyBindingPatchArgs)(nil)).Elem()
}

type ValidatingAdmissionPolicyBindingPatchInput interface {
	pulumi.Input

	ToValidatingAdmissionPolicyBindingPatchOutput() ValidatingAdmissionPolicyBindingPatchOutput
	ToValidatingAdmissionPolicyBindingPatchOutputWithContext(ctx context.Context) ValidatingAdmissionPolicyBindingPatchOutput
}

func (*ValidatingAdmissionPolicyBindingPatch) ElementType() reflect.Type {
	return reflect.TypeOf((**ValidatingAdmissionPolicyBindingPatch)(nil)).Elem()
}

func (i *ValidatingAdmissionPolicyBindingPatch) ToValidatingAdmissionPolicyBindingPatchOutput() ValidatingAdmissionPolicyBindingPatchOutput {
	return i.ToValidatingAdmissionPolicyBindingPatchOutputWithContext(context.Background())
}

func (i *ValidatingAdmissionPolicyBindingPatch) ToValidatingAdmissionPolicyBindingPatchOutputWithContext(ctx context.Context) ValidatingAdmissionPolicyBindingPatchOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ValidatingAdmissionPolicyBindingPatchOutput)
}

// ValidatingAdmissionPolicyBindingPatchArrayInput is an input type that accepts ValidatingAdmissionPolicyBindingPatchArray and ValidatingAdmissionPolicyBindingPatchArrayOutput values.
// You can construct a concrete instance of `ValidatingAdmissionPolicyBindingPatchArrayInput` via:
//
//	ValidatingAdmissionPolicyBindingPatchArray{ ValidatingAdmissionPolicyBindingPatchArgs{...} }
type ValidatingAdmissionPolicyBindingPatchArrayInput interface {
	pulumi.Input

	ToValidatingAdmissionPolicyBindingPatchArrayOutput() ValidatingAdmissionPolicyBindingPatchArrayOutput
	ToValidatingAdmissionPolicyBindingPatchArrayOutputWithContext(context.Context) ValidatingAdmissionPolicyBindingPatchArrayOutput
}

type ValidatingAdmissionPolicyBindingPatchArray []ValidatingAdmissionPolicyBindingPatchInput

func (ValidatingAdmissionPolicyBindingPatchArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ValidatingAdmissionPolicyBindingPatch)(nil)).Elem()
}

func (i ValidatingAdmissionPolicyBindingPatchArray) ToValidatingAdmissionPolicyBindingPatchArrayOutput() ValidatingAdmissionPolicyBindingPatchArrayOutput {
	return i.ToValidatingAdmissionPolicyBindingPatchArrayOutputWithContext(context.Background())
}

func (i ValidatingAdmissionPolicyBindingPatchArray) ToValidatingAdmissionPolicyBindingPatchArrayOutputWithContext(ctx context.Context) ValidatingAdmissionPolicyBindingPatchArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ValidatingAdmissionPolicyBindingPatchArrayOutput)
}

// ValidatingAdmissionPolicyBindingPatchMapInput is an input type that accepts ValidatingAdmissionPolicyBindingPatchMap and ValidatingAdmissionPolicyBindingPatchMapOutput values.
// You can construct a concrete instance of `ValidatingAdmissionPolicyBindingPatchMapInput` via:
//
//	ValidatingAdmissionPolicyBindingPatchMap{ "key": ValidatingAdmissionPolicyBindingPatchArgs{...} }
type ValidatingAdmissionPolicyBindingPatchMapInput interface {
	pulumi.Input

	ToValidatingAdmissionPolicyBindingPatchMapOutput() ValidatingAdmissionPolicyBindingPatchMapOutput
	ToValidatingAdmissionPolicyBindingPatchMapOutputWithContext(context.Context) ValidatingAdmissionPolicyBindingPatchMapOutput
}

type ValidatingAdmissionPolicyBindingPatchMap map[string]ValidatingAdmissionPolicyBindingPatchInput

func (ValidatingAdmissionPolicyBindingPatchMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ValidatingAdmissionPolicyBindingPatch)(nil)).Elem()
}

func (i ValidatingAdmissionPolicyBindingPatchMap) ToValidatingAdmissionPolicyBindingPatchMapOutput() ValidatingAdmissionPolicyBindingPatchMapOutput {
	return i.ToValidatingAdmissionPolicyBindingPatchMapOutputWithContext(context.Background())
}

func (i ValidatingAdmissionPolicyBindingPatchMap) ToValidatingAdmissionPolicyBindingPatchMapOutputWithContext(ctx context.Context) ValidatingAdmissionPolicyBindingPatchMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ValidatingAdmissionPolicyBindingPatchMapOutput)
}

type ValidatingAdmissionPolicyBindingPatchOutput struct{ *pulumi.OutputState }

func (ValidatingAdmissionPolicyBindingPatchOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ValidatingAdmissionPolicyBindingPatch)(nil)).Elem()
}

func (o ValidatingAdmissionPolicyBindingPatchOutput) ToValidatingAdmissionPolicyBindingPatchOutput() ValidatingAdmissionPolicyBindingPatchOutput {
	return o
}

func (o ValidatingAdmissionPolicyBindingPatchOutput) ToValidatingAdmissionPolicyBindingPatchOutputWithContext(ctx context.Context) ValidatingAdmissionPolicyBindingPatchOutput {
	return o
}

// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
func (o ValidatingAdmissionPolicyBindingPatchOutput) ApiVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ValidatingAdmissionPolicyBindingPatch) pulumi.StringPtrOutput { return v.ApiVersion }).(pulumi.StringPtrOutput)
}

// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
func (o ValidatingAdmissionPolicyBindingPatchOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ValidatingAdmissionPolicyBindingPatch) pulumi.StringPtrOutput { return v.Kind }).(pulumi.StringPtrOutput)
}

// Standard object metadata; More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata.
func (o ValidatingAdmissionPolicyBindingPatchOutput) Metadata() metav1.ObjectMetaPatchPtrOutput {
	return o.ApplyT(func(v *ValidatingAdmissionPolicyBindingPatch) metav1.ObjectMetaPatchPtrOutput { return v.Metadata }).(metav1.ObjectMetaPatchPtrOutput)
}

// Specification of the desired behavior of the ValidatingAdmissionPolicyBinding.
func (o ValidatingAdmissionPolicyBindingPatchOutput) Spec() ValidatingAdmissionPolicyBindingSpecPatchPtrOutput {
	return o.ApplyT(func(v *ValidatingAdmissionPolicyBindingPatch) ValidatingAdmissionPolicyBindingSpecPatchPtrOutput {
		return v.Spec
	}).(ValidatingAdmissionPolicyBindingSpecPatchPtrOutput)
}

type ValidatingAdmissionPolicyBindingPatchArrayOutput struct{ *pulumi.OutputState }

func (ValidatingAdmissionPolicyBindingPatchArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ValidatingAdmissionPolicyBindingPatch)(nil)).Elem()
}

func (o ValidatingAdmissionPolicyBindingPatchArrayOutput) ToValidatingAdmissionPolicyBindingPatchArrayOutput() ValidatingAdmissionPolicyBindingPatchArrayOutput {
	return o
}

func (o ValidatingAdmissionPolicyBindingPatchArrayOutput) ToValidatingAdmissionPolicyBindingPatchArrayOutputWithContext(ctx context.Context) ValidatingAdmissionPolicyBindingPatchArrayOutput {
	return o
}

func (o ValidatingAdmissionPolicyBindingPatchArrayOutput) Index(i pulumi.IntInput) ValidatingAdmissionPolicyBindingPatchOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *ValidatingAdmissionPolicyBindingPatch {
		return vs[0].([]*ValidatingAdmissionPolicyBindingPatch)[vs[1].(int)]
	}).(ValidatingAdmissionPolicyBindingPatchOutput)
}

type ValidatingAdmissionPolicyBindingPatchMapOutput struct{ *pulumi.OutputState }

func (ValidatingAdmissionPolicyBindingPatchMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ValidatingAdmissionPolicyBindingPatch)(nil)).Elem()
}

func (o ValidatingAdmissionPolicyBindingPatchMapOutput) ToValidatingAdmissionPolicyBindingPatchMapOutput() ValidatingAdmissionPolicyBindingPatchMapOutput {
	return o
}

func (o ValidatingAdmissionPolicyBindingPatchMapOutput) ToValidatingAdmissionPolicyBindingPatchMapOutputWithContext(ctx context.Context) ValidatingAdmissionPolicyBindingPatchMapOutput {
	return o
}

func (o ValidatingAdmissionPolicyBindingPatchMapOutput) MapIndex(k pulumi.StringInput) ValidatingAdmissionPolicyBindingPatchOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *ValidatingAdmissionPolicyBindingPatch {
		return vs[0].(map[string]*ValidatingAdmissionPolicyBindingPatch)[vs[1].(string)]
	}).(ValidatingAdmissionPolicyBindingPatchOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ValidatingAdmissionPolicyBindingPatchInput)(nil)).Elem(), &ValidatingAdmissionPolicyBindingPatch{})
	pulumi.RegisterInputType(reflect.TypeOf((*ValidatingAdmissionPolicyBindingPatchArrayInput)(nil)).Elem(), ValidatingAdmissionPolicyBindingPatchArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ValidatingAdmissionPolicyBindingPatchMapInput)(nil)).Elem(), ValidatingAdmissionPolicyBindingPatchMap{})
	pulumi.RegisterOutputType(ValidatingAdmissionPolicyBindingPatchOutput{})
	pulumi.RegisterOutputType(ValidatingAdmissionPolicyBindingPatchArrayOutput{})
	pulumi.RegisterOutputType(ValidatingAdmissionPolicyBindingPatchMapOutput{})
}
