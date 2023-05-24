// Code generated by pulumigen DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v1beta1

import (
	"context"
	"reflect"

	"errors"
	metav1 "github.com/pulumi/pulumi-kubernetes/sdk/v3/go/kubernetes/meta/v1"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// PodSecurityPolicyList is a list of PodSecurityPolicy objects. Deprecated: use PodSecurityPolicyList from policy API Group instead.
type PodSecurityPolicyList struct {
	pulumi.CustomResourceState

	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion pulumi.StringOutput `pulumi:"apiVersion"`
	// items is a list of schema objects.
	Items PodSecurityPolicyTypeArrayOutput `pulumi:"items"`
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind pulumi.StringOutput `pulumi:"kind"`
	// Standard list metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	Metadata metav1.ListMetaOutput `pulumi:"metadata"`
}

// NewPodSecurityPolicyList registers a new resource with the given unique name, arguments, and options.
func NewPodSecurityPolicyList(ctx *pulumi.Context,
	name string, args *PodSecurityPolicyListArgs, opts ...pulumi.ResourceOption) (*PodSecurityPolicyList, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Items == nil {
		return nil, errors.New("invalid value for required argument 'Items'")
	}
	args.ApiVersion = pulumi.StringPtr("extensions/v1beta1")
	args.Kind = pulumi.StringPtr("PodSecurityPolicyList")
	var resource PodSecurityPolicyList
	err := ctx.RegisterResource("kubernetes:extensions/v1beta1:PodSecurityPolicyList", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetPodSecurityPolicyList gets an existing PodSecurityPolicyList resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetPodSecurityPolicyList(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *PodSecurityPolicyListState, opts ...pulumi.ResourceOption) (*PodSecurityPolicyList, error) {
	var resource PodSecurityPolicyList
	err := ctx.ReadResource("kubernetes:extensions/v1beta1:PodSecurityPolicyList", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering PodSecurityPolicyList resources.
type podSecurityPolicyListState struct {
}

type PodSecurityPolicyListState struct {
}

func (PodSecurityPolicyListState) ElementType() reflect.Type {
	return reflect.TypeOf((*podSecurityPolicyListState)(nil)).Elem()
}

type podSecurityPolicyListArgs struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion *string `pulumi:"apiVersion"`
	// items is a list of schema objects.
	Items []PodSecurityPolicyType `pulumi:"items"`
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind *string `pulumi:"kind"`
	// Standard list metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	Metadata *metav1.ListMeta `pulumi:"metadata"`
}

// The set of arguments for constructing a PodSecurityPolicyList resource.
type PodSecurityPolicyListArgs struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion pulumi.StringPtrInput
	// items is a list of schema objects.
	Items PodSecurityPolicyTypeArrayInput
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind pulumi.StringPtrInput
	// Standard list metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	Metadata metav1.ListMetaPtrInput
}

func (PodSecurityPolicyListArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*podSecurityPolicyListArgs)(nil)).Elem()
}

type PodSecurityPolicyListInput interface {
	pulumi.Input

	ToPodSecurityPolicyListOutput() PodSecurityPolicyListOutput
	ToPodSecurityPolicyListOutputWithContext(ctx context.Context) PodSecurityPolicyListOutput
}

func (*PodSecurityPolicyList) ElementType() reflect.Type {
	return reflect.TypeOf((**PodSecurityPolicyList)(nil)).Elem()
}

func (i *PodSecurityPolicyList) ToPodSecurityPolicyListOutput() PodSecurityPolicyListOutput {
	return i.ToPodSecurityPolicyListOutputWithContext(context.Background())
}

func (i *PodSecurityPolicyList) ToPodSecurityPolicyListOutputWithContext(ctx context.Context) PodSecurityPolicyListOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PodSecurityPolicyListOutput)
}

// PodSecurityPolicyListArrayInput is an input type that accepts PodSecurityPolicyListArray and PodSecurityPolicyListArrayOutput values.
// You can construct a concrete instance of `PodSecurityPolicyListArrayInput` via:
//
//	PodSecurityPolicyListArray{ PodSecurityPolicyListArgs{...} }
type PodSecurityPolicyListArrayInput interface {
	pulumi.Input

	ToPodSecurityPolicyListArrayOutput() PodSecurityPolicyListArrayOutput
	ToPodSecurityPolicyListArrayOutputWithContext(context.Context) PodSecurityPolicyListArrayOutput
}

type PodSecurityPolicyListArray []PodSecurityPolicyListInput

func (PodSecurityPolicyListArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*PodSecurityPolicyList)(nil)).Elem()
}

func (i PodSecurityPolicyListArray) ToPodSecurityPolicyListArrayOutput() PodSecurityPolicyListArrayOutput {
	return i.ToPodSecurityPolicyListArrayOutputWithContext(context.Background())
}

func (i PodSecurityPolicyListArray) ToPodSecurityPolicyListArrayOutputWithContext(ctx context.Context) PodSecurityPolicyListArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PodSecurityPolicyListArrayOutput)
}

// PodSecurityPolicyListMapInput is an input type that accepts PodSecurityPolicyListMap and PodSecurityPolicyListMapOutput values.
// You can construct a concrete instance of `PodSecurityPolicyListMapInput` via:
//
//	PodSecurityPolicyListMap{ "key": PodSecurityPolicyListArgs{...} }
type PodSecurityPolicyListMapInput interface {
	pulumi.Input

	ToPodSecurityPolicyListMapOutput() PodSecurityPolicyListMapOutput
	ToPodSecurityPolicyListMapOutputWithContext(context.Context) PodSecurityPolicyListMapOutput
}

type PodSecurityPolicyListMap map[string]PodSecurityPolicyListInput

func (PodSecurityPolicyListMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*PodSecurityPolicyList)(nil)).Elem()
}

func (i PodSecurityPolicyListMap) ToPodSecurityPolicyListMapOutput() PodSecurityPolicyListMapOutput {
	return i.ToPodSecurityPolicyListMapOutputWithContext(context.Background())
}

func (i PodSecurityPolicyListMap) ToPodSecurityPolicyListMapOutputWithContext(ctx context.Context) PodSecurityPolicyListMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PodSecurityPolicyListMapOutput)
}

type PodSecurityPolicyListOutput struct{ *pulumi.OutputState }

func (PodSecurityPolicyListOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PodSecurityPolicyList)(nil)).Elem()
}

func (o PodSecurityPolicyListOutput) ToPodSecurityPolicyListOutput() PodSecurityPolicyListOutput {
	return o
}

func (o PodSecurityPolicyListOutput) ToPodSecurityPolicyListOutputWithContext(ctx context.Context) PodSecurityPolicyListOutput {
	return o
}

// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
func (o PodSecurityPolicyListOutput) ApiVersion() pulumi.StringOutput {
	return o.ApplyT(func(v *PodSecurityPolicyList) pulumi.StringOutput { return v.ApiVersion }).(pulumi.StringOutput)
}

// items is a list of schema objects.
func (o PodSecurityPolicyListOutput) Items() PodSecurityPolicyTypeArrayOutput {
	return o.ApplyT(func(v *PodSecurityPolicyList) PodSecurityPolicyTypeArrayOutput { return v.Items }).(PodSecurityPolicyTypeArrayOutput)
}

// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
func (o PodSecurityPolicyListOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v *PodSecurityPolicyList) pulumi.StringOutput { return v.Kind }).(pulumi.StringOutput)
}

// Standard list metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
func (o PodSecurityPolicyListOutput) Metadata() metav1.ListMetaOutput {
	return o.ApplyT(func(v *PodSecurityPolicyList) metav1.ListMetaOutput { return v.Metadata }).(metav1.ListMetaOutput)
}

type PodSecurityPolicyListArrayOutput struct{ *pulumi.OutputState }

func (PodSecurityPolicyListArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*PodSecurityPolicyList)(nil)).Elem()
}

func (o PodSecurityPolicyListArrayOutput) ToPodSecurityPolicyListArrayOutput() PodSecurityPolicyListArrayOutput {
	return o
}

func (o PodSecurityPolicyListArrayOutput) ToPodSecurityPolicyListArrayOutputWithContext(ctx context.Context) PodSecurityPolicyListArrayOutput {
	return o
}

func (o PodSecurityPolicyListArrayOutput) Index(i pulumi.IntInput) PodSecurityPolicyListOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *PodSecurityPolicyList {
		return vs[0].([]*PodSecurityPolicyList)[vs[1].(int)]
	}).(PodSecurityPolicyListOutput)
}

type PodSecurityPolicyListMapOutput struct{ *pulumi.OutputState }

func (PodSecurityPolicyListMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*PodSecurityPolicyList)(nil)).Elem()
}

func (o PodSecurityPolicyListMapOutput) ToPodSecurityPolicyListMapOutput() PodSecurityPolicyListMapOutput {
	return o
}

func (o PodSecurityPolicyListMapOutput) ToPodSecurityPolicyListMapOutputWithContext(ctx context.Context) PodSecurityPolicyListMapOutput {
	return o
}

func (o PodSecurityPolicyListMapOutput) MapIndex(k pulumi.StringInput) PodSecurityPolicyListOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *PodSecurityPolicyList {
		return vs[0].(map[string]*PodSecurityPolicyList)[vs[1].(string)]
	}).(PodSecurityPolicyListOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*PodSecurityPolicyListInput)(nil)).Elem(), &PodSecurityPolicyList{})
	pulumi.RegisterInputType(reflect.TypeOf((*PodSecurityPolicyListArrayInput)(nil)).Elem(), PodSecurityPolicyListArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*PodSecurityPolicyListMapInput)(nil)).Elem(), PodSecurityPolicyListMap{})
	pulumi.RegisterOutputType(PodSecurityPolicyListOutput{})
	pulumi.RegisterOutputType(PodSecurityPolicyListArrayOutput{})
	pulumi.RegisterOutputType(PodSecurityPolicyListMapOutput{})
}
