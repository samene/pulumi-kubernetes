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
// DEPRECATED - This group version of PriorityClass is deprecated by scheduling.k8s.io/v1/PriorityClass. PriorityClass defines mapping from a priority class name to the priority integer value. The value can be any valid integer.
type PriorityClassPatch struct {
	pulumi.CustomResourceState

	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion pulumi.StringPtrOutput `pulumi:"apiVersion"`
	// description is an arbitrary string that usually provides guidelines on when this priority class should be used.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// globalDefault specifies whether this PriorityClass should be considered as the default priority for pods that do not have any priority class. Only one PriorityClass can be marked as `globalDefault`. However, if more than one PriorityClasses exists with their `globalDefault` field set to true, the smallest value of such global default PriorityClasses will be used as the default priority.
	GlobalDefault pulumi.BoolPtrOutput `pulumi:"globalDefault"`
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind pulumi.StringPtrOutput `pulumi:"kind"`
	// Standard object's metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	Metadata metav1.ObjectMetaPatchPtrOutput `pulumi:"metadata"`
	// PreemptionPolicy is the Policy for preempting pods with lower priority. One of Never, PreemptLowerPriority. Defaults to PreemptLowerPriority if unset. This field is alpha-level and is only honored by servers that enable the NonPreemptingPriority feature.
	PreemptionPolicy pulumi.StringPtrOutput `pulumi:"preemptionPolicy"`
	// The value of this priority class. This is the actual priority that pods receive when they have the name of this class in their pod spec.
	Value pulumi.IntPtrOutput `pulumi:"value"`
}

// NewPriorityClassPatch registers a new resource with the given unique name, arguments, and options.
func NewPriorityClassPatch(ctx *pulumi.Context,
	name string, args *PriorityClassPatchArgs, opts ...pulumi.ResourceOption) (*PriorityClassPatch, error) {
	if args == nil {
		args = &PriorityClassPatchArgs{}
	}

	args.ApiVersion = pulumi.StringPtr("scheduling.k8s.io/v1beta1")
	args.Kind = pulumi.StringPtr("PriorityClass")
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("kubernetes:scheduling.k8s.io/v1:PriorityClassPatch"),
		},
		{
			Type: pulumi.String("kubernetes:scheduling.k8s.io/v1alpha1:PriorityClassPatch"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource PriorityClassPatch
	err := ctx.RegisterResource("kubernetes:scheduling.k8s.io/v1beta1:PriorityClassPatch", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetPriorityClassPatch gets an existing PriorityClassPatch resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetPriorityClassPatch(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *PriorityClassPatchState, opts ...pulumi.ResourceOption) (*PriorityClassPatch, error) {
	var resource PriorityClassPatch
	err := ctx.ReadResource("kubernetes:scheduling.k8s.io/v1beta1:PriorityClassPatch", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering PriorityClassPatch resources.
type priorityClassPatchState struct {
}

type PriorityClassPatchState struct {
}

func (PriorityClassPatchState) ElementType() reflect.Type {
	return reflect.TypeOf((*priorityClassPatchState)(nil)).Elem()
}

type priorityClassPatchArgs struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion *string `pulumi:"apiVersion"`
	// description is an arbitrary string that usually provides guidelines on when this priority class should be used.
	Description *string `pulumi:"description"`
	// globalDefault specifies whether this PriorityClass should be considered as the default priority for pods that do not have any priority class. Only one PriorityClass can be marked as `globalDefault`. However, if more than one PriorityClasses exists with their `globalDefault` field set to true, the smallest value of such global default PriorityClasses will be used as the default priority.
	GlobalDefault *bool `pulumi:"globalDefault"`
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind *string `pulumi:"kind"`
	// Standard object's metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	Metadata *metav1.ObjectMetaPatch `pulumi:"metadata"`
	// PreemptionPolicy is the Policy for preempting pods with lower priority. One of Never, PreemptLowerPriority. Defaults to PreemptLowerPriority if unset. This field is alpha-level and is only honored by servers that enable the NonPreemptingPriority feature.
	PreemptionPolicy *string `pulumi:"preemptionPolicy"`
	// The value of this priority class. This is the actual priority that pods receive when they have the name of this class in their pod spec.
	Value *int `pulumi:"value"`
}

// The set of arguments for constructing a PriorityClassPatch resource.
type PriorityClassPatchArgs struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion pulumi.StringPtrInput
	// description is an arbitrary string that usually provides guidelines on when this priority class should be used.
	Description pulumi.StringPtrInput
	// globalDefault specifies whether this PriorityClass should be considered as the default priority for pods that do not have any priority class. Only one PriorityClass can be marked as `globalDefault`. However, if more than one PriorityClasses exists with their `globalDefault` field set to true, the smallest value of such global default PriorityClasses will be used as the default priority.
	GlobalDefault pulumi.BoolPtrInput
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind pulumi.StringPtrInput
	// Standard object's metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	Metadata metav1.ObjectMetaPatchPtrInput
	// PreemptionPolicy is the Policy for preempting pods with lower priority. One of Never, PreemptLowerPriority. Defaults to PreemptLowerPriority if unset. This field is alpha-level and is only honored by servers that enable the NonPreemptingPriority feature.
	PreemptionPolicy pulumi.StringPtrInput
	// The value of this priority class. This is the actual priority that pods receive when they have the name of this class in their pod spec.
	Value pulumi.IntPtrInput
}

func (PriorityClassPatchArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*priorityClassPatchArgs)(nil)).Elem()
}

type PriorityClassPatchInput interface {
	pulumi.Input

	ToPriorityClassPatchOutput() PriorityClassPatchOutput
	ToPriorityClassPatchOutputWithContext(ctx context.Context) PriorityClassPatchOutput
}

func (*PriorityClassPatch) ElementType() reflect.Type {
	return reflect.TypeOf((**PriorityClassPatch)(nil)).Elem()
}

func (i *PriorityClassPatch) ToPriorityClassPatchOutput() PriorityClassPatchOutput {
	return i.ToPriorityClassPatchOutputWithContext(context.Background())
}

func (i *PriorityClassPatch) ToPriorityClassPatchOutputWithContext(ctx context.Context) PriorityClassPatchOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PriorityClassPatchOutput)
}

// PriorityClassPatchArrayInput is an input type that accepts PriorityClassPatchArray and PriorityClassPatchArrayOutput values.
// You can construct a concrete instance of `PriorityClassPatchArrayInput` via:
//
//	PriorityClassPatchArray{ PriorityClassPatchArgs{...} }
type PriorityClassPatchArrayInput interface {
	pulumi.Input

	ToPriorityClassPatchArrayOutput() PriorityClassPatchArrayOutput
	ToPriorityClassPatchArrayOutputWithContext(context.Context) PriorityClassPatchArrayOutput
}

type PriorityClassPatchArray []PriorityClassPatchInput

func (PriorityClassPatchArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*PriorityClassPatch)(nil)).Elem()
}

func (i PriorityClassPatchArray) ToPriorityClassPatchArrayOutput() PriorityClassPatchArrayOutput {
	return i.ToPriorityClassPatchArrayOutputWithContext(context.Background())
}

func (i PriorityClassPatchArray) ToPriorityClassPatchArrayOutputWithContext(ctx context.Context) PriorityClassPatchArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PriorityClassPatchArrayOutput)
}

// PriorityClassPatchMapInput is an input type that accepts PriorityClassPatchMap and PriorityClassPatchMapOutput values.
// You can construct a concrete instance of `PriorityClassPatchMapInput` via:
//
//	PriorityClassPatchMap{ "key": PriorityClassPatchArgs{...} }
type PriorityClassPatchMapInput interface {
	pulumi.Input

	ToPriorityClassPatchMapOutput() PriorityClassPatchMapOutput
	ToPriorityClassPatchMapOutputWithContext(context.Context) PriorityClassPatchMapOutput
}

type PriorityClassPatchMap map[string]PriorityClassPatchInput

func (PriorityClassPatchMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*PriorityClassPatch)(nil)).Elem()
}

func (i PriorityClassPatchMap) ToPriorityClassPatchMapOutput() PriorityClassPatchMapOutput {
	return i.ToPriorityClassPatchMapOutputWithContext(context.Background())
}

func (i PriorityClassPatchMap) ToPriorityClassPatchMapOutputWithContext(ctx context.Context) PriorityClassPatchMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PriorityClassPatchMapOutput)
}

type PriorityClassPatchOutput struct{ *pulumi.OutputState }

func (PriorityClassPatchOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PriorityClassPatch)(nil)).Elem()
}

func (o PriorityClassPatchOutput) ToPriorityClassPatchOutput() PriorityClassPatchOutput {
	return o
}

func (o PriorityClassPatchOutput) ToPriorityClassPatchOutputWithContext(ctx context.Context) PriorityClassPatchOutput {
	return o
}

// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
func (o PriorityClassPatchOutput) ApiVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PriorityClassPatch) pulumi.StringPtrOutput { return v.ApiVersion }).(pulumi.StringPtrOutput)
}

// description is an arbitrary string that usually provides guidelines on when this priority class should be used.
func (o PriorityClassPatchOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PriorityClassPatch) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// globalDefault specifies whether this PriorityClass should be considered as the default priority for pods that do not have any priority class. Only one PriorityClass can be marked as `globalDefault`. However, if more than one PriorityClasses exists with their `globalDefault` field set to true, the smallest value of such global default PriorityClasses will be used as the default priority.
func (o PriorityClassPatchOutput) GlobalDefault() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *PriorityClassPatch) pulumi.BoolPtrOutput { return v.GlobalDefault }).(pulumi.BoolPtrOutput)
}

// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
func (o PriorityClassPatchOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PriorityClassPatch) pulumi.StringPtrOutput { return v.Kind }).(pulumi.StringPtrOutput)
}

// Standard object's metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
func (o PriorityClassPatchOutput) Metadata() metav1.ObjectMetaPatchPtrOutput {
	return o.ApplyT(func(v *PriorityClassPatch) metav1.ObjectMetaPatchPtrOutput { return v.Metadata }).(metav1.ObjectMetaPatchPtrOutput)
}

// PreemptionPolicy is the Policy for preempting pods with lower priority. One of Never, PreemptLowerPriority. Defaults to PreemptLowerPriority if unset. This field is alpha-level and is only honored by servers that enable the NonPreemptingPriority feature.
func (o PriorityClassPatchOutput) PreemptionPolicy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PriorityClassPatch) pulumi.StringPtrOutput { return v.PreemptionPolicy }).(pulumi.StringPtrOutput)
}

// The value of this priority class. This is the actual priority that pods receive when they have the name of this class in their pod spec.
func (o PriorityClassPatchOutput) Value() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *PriorityClassPatch) pulumi.IntPtrOutput { return v.Value }).(pulumi.IntPtrOutput)
}

type PriorityClassPatchArrayOutput struct{ *pulumi.OutputState }

func (PriorityClassPatchArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*PriorityClassPatch)(nil)).Elem()
}

func (o PriorityClassPatchArrayOutput) ToPriorityClassPatchArrayOutput() PriorityClassPatchArrayOutput {
	return o
}

func (o PriorityClassPatchArrayOutput) ToPriorityClassPatchArrayOutputWithContext(ctx context.Context) PriorityClassPatchArrayOutput {
	return o
}

func (o PriorityClassPatchArrayOutput) Index(i pulumi.IntInput) PriorityClassPatchOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *PriorityClassPatch {
		return vs[0].([]*PriorityClassPatch)[vs[1].(int)]
	}).(PriorityClassPatchOutput)
}

type PriorityClassPatchMapOutput struct{ *pulumi.OutputState }

func (PriorityClassPatchMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*PriorityClassPatch)(nil)).Elem()
}

func (o PriorityClassPatchMapOutput) ToPriorityClassPatchMapOutput() PriorityClassPatchMapOutput {
	return o
}

func (o PriorityClassPatchMapOutput) ToPriorityClassPatchMapOutputWithContext(ctx context.Context) PriorityClassPatchMapOutput {
	return o
}

func (o PriorityClassPatchMapOutput) MapIndex(k pulumi.StringInput) PriorityClassPatchOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *PriorityClassPatch {
		return vs[0].(map[string]*PriorityClassPatch)[vs[1].(string)]
	}).(PriorityClassPatchOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*PriorityClassPatchInput)(nil)).Elem(), &PriorityClassPatch{})
	pulumi.RegisterInputType(reflect.TypeOf((*PriorityClassPatchArrayInput)(nil)).Elem(), PriorityClassPatchArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*PriorityClassPatchMapInput)(nil)).Elem(), PriorityClassPatchMap{})
	pulumi.RegisterOutputType(PriorityClassPatchOutput{})
	pulumi.RegisterOutputType(PriorityClassPatchArrayOutput{})
	pulumi.RegisterOutputType(PriorityClassPatchMapOutput{})
}
