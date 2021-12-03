// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v1

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	metav1 "github.com/pulumi/pulumi-kubernetes/sdk/v3/go/kubernetes/meta/v1"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// MutatingWebhookConfigurationList is a list of MutatingWebhookConfiguration.
type MutatingWebhookConfigurationList struct {
	pulumi.CustomResourceState

	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion pulumi.StringPtrOutput `pulumi:"apiVersion"`
	// List of MutatingWebhookConfiguration.
	Items MutatingWebhookConfigurationTypeArrayOutput `pulumi:"items"`
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind pulumi.StringPtrOutput `pulumi:"kind"`
	// Standard list metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Metadata metav1.ListMetaPtrOutput `pulumi:"metadata"`
}

// NewMutatingWebhookConfigurationList registers a new resource with the given unique name, arguments, and options.
func NewMutatingWebhookConfigurationList(ctx *pulumi.Context,
	name string, args *MutatingWebhookConfigurationListArgs, opts ...pulumi.ResourceOption) (*MutatingWebhookConfigurationList, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Items == nil {
		return nil, errors.New("invalid value for required argument 'Items'")
	}
	args.ApiVersion = pulumi.StringPtr("admissionregistration.k8s.io/v1")
	args.Kind = pulumi.StringPtr("MutatingWebhookConfigurationList")
	var resource MutatingWebhookConfigurationList
	err := ctx.RegisterResource("kubernetes:admissionregistration.k8s.io/v1:MutatingWebhookConfigurationList", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetMutatingWebhookConfigurationList gets an existing MutatingWebhookConfigurationList resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetMutatingWebhookConfigurationList(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *MutatingWebhookConfigurationListState, opts ...pulumi.ResourceOption) (*MutatingWebhookConfigurationList, error) {
	var resource MutatingWebhookConfigurationList
	err := ctx.ReadResource("kubernetes:admissionregistration.k8s.io/v1:MutatingWebhookConfigurationList", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering MutatingWebhookConfigurationList resources.
type mutatingWebhookConfigurationListState struct {
}

type MutatingWebhookConfigurationListState struct {
}

func (MutatingWebhookConfigurationListState) ElementType() reflect.Type {
	return reflect.TypeOf((*mutatingWebhookConfigurationListState)(nil)).Elem()
}

type mutatingWebhookConfigurationListArgs struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion *string `pulumi:"apiVersion"`
	// List of MutatingWebhookConfiguration.
	Items []MutatingWebhookConfigurationType `pulumi:"items"`
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind *string `pulumi:"kind"`
	// Standard list metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Metadata *metav1.ListMeta `pulumi:"metadata"`
}

// The set of arguments for constructing a MutatingWebhookConfigurationList resource.
type MutatingWebhookConfigurationListArgs struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion pulumi.StringPtrInput
	// List of MutatingWebhookConfiguration.
	Items MutatingWebhookConfigurationTypeArrayInput
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind pulumi.StringPtrInput
	// Standard list metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Metadata metav1.ListMetaPtrInput
}

func (MutatingWebhookConfigurationListArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*mutatingWebhookConfigurationListArgs)(nil)).Elem()
}

type MutatingWebhookConfigurationListInput interface {
	pulumi.Input

	ToMutatingWebhookConfigurationListOutput() MutatingWebhookConfigurationListOutput
	ToMutatingWebhookConfigurationListOutputWithContext(ctx context.Context) MutatingWebhookConfigurationListOutput
}

func (*MutatingWebhookConfigurationList) ElementType() reflect.Type {
	return reflect.TypeOf((**MutatingWebhookConfigurationList)(nil)).Elem()
}

func (i *MutatingWebhookConfigurationList) ToMutatingWebhookConfigurationListOutput() MutatingWebhookConfigurationListOutput {
	return i.ToMutatingWebhookConfigurationListOutputWithContext(context.Background())
}

func (i *MutatingWebhookConfigurationList) ToMutatingWebhookConfigurationListOutputWithContext(ctx context.Context) MutatingWebhookConfigurationListOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MutatingWebhookConfigurationListOutput)
}

// MutatingWebhookConfigurationListArrayInput is an input type that accepts MutatingWebhookConfigurationListArray and MutatingWebhookConfigurationListArrayOutput values.
// You can construct a concrete instance of `MutatingWebhookConfigurationListArrayInput` via:
//
//          MutatingWebhookConfigurationListArray{ MutatingWebhookConfigurationListArgs{...} }
type MutatingWebhookConfigurationListArrayInput interface {
	pulumi.Input

	ToMutatingWebhookConfigurationListArrayOutput() MutatingWebhookConfigurationListArrayOutput
	ToMutatingWebhookConfigurationListArrayOutputWithContext(context.Context) MutatingWebhookConfigurationListArrayOutput
}

type MutatingWebhookConfigurationListArray []MutatingWebhookConfigurationListInput

func (MutatingWebhookConfigurationListArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*MutatingWebhookConfigurationList)(nil)).Elem()
}

func (i MutatingWebhookConfigurationListArray) ToMutatingWebhookConfigurationListArrayOutput() MutatingWebhookConfigurationListArrayOutput {
	return i.ToMutatingWebhookConfigurationListArrayOutputWithContext(context.Background())
}

func (i MutatingWebhookConfigurationListArray) ToMutatingWebhookConfigurationListArrayOutputWithContext(ctx context.Context) MutatingWebhookConfigurationListArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MutatingWebhookConfigurationListArrayOutput)
}

// MutatingWebhookConfigurationListMapInput is an input type that accepts MutatingWebhookConfigurationListMap and MutatingWebhookConfigurationListMapOutput values.
// You can construct a concrete instance of `MutatingWebhookConfigurationListMapInput` via:
//
//          MutatingWebhookConfigurationListMap{ "key": MutatingWebhookConfigurationListArgs{...} }
type MutatingWebhookConfigurationListMapInput interface {
	pulumi.Input

	ToMutatingWebhookConfigurationListMapOutput() MutatingWebhookConfigurationListMapOutput
	ToMutatingWebhookConfigurationListMapOutputWithContext(context.Context) MutatingWebhookConfigurationListMapOutput
}

type MutatingWebhookConfigurationListMap map[string]MutatingWebhookConfigurationListInput

func (MutatingWebhookConfigurationListMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*MutatingWebhookConfigurationList)(nil)).Elem()
}

func (i MutatingWebhookConfigurationListMap) ToMutatingWebhookConfigurationListMapOutput() MutatingWebhookConfigurationListMapOutput {
	return i.ToMutatingWebhookConfigurationListMapOutputWithContext(context.Background())
}

func (i MutatingWebhookConfigurationListMap) ToMutatingWebhookConfigurationListMapOutputWithContext(ctx context.Context) MutatingWebhookConfigurationListMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MutatingWebhookConfigurationListMapOutput)
}

type MutatingWebhookConfigurationListOutput struct{ *pulumi.OutputState }

func (MutatingWebhookConfigurationListOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**MutatingWebhookConfigurationList)(nil)).Elem()
}

func (o MutatingWebhookConfigurationListOutput) ToMutatingWebhookConfigurationListOutput() MutatingWebhookConfigurationListOutput {
	return o
}

func (o MutatingWebhookConfigurationListOutput) ToMutatingWebhookConfigurationListOutputWithContext(ctx context.Context) MutatingWebhookConfigurationListOutput {
	return o
}

type MutatingWebhookConfigurationListArrayOutput struct{ *pulumi.OutputState }

func (MutatingWebhookConfigurationListArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*MutatingWebhookConfigurationList)(nil)).Elem()
}

func (o MutatingWebhookConfigurationListArrayOutput) ToMutatingWebhookConfigurationListArrayOutput() MutatingWebhookConfigurationListArrayOutput {
	return o
}

func (o MutatingWebhookConfigurationListArrayOutput) ToMutatingWebhookConfigurationListArrayOutputWithContext(ctx context.Context) MutatingWebhookConfigurationListArrayOutput {
	return o
}

func (o MutatingWebhookConfigurationListArrayOutput) Index(i pulumi.IntInput) MutatingWebhookConfigurationListOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *MutatingWebhookConfigurationList {
		return vs[0].([]*MutatingWebhookConfigurationList)[vs[1].(int)]
	}).(MutatingWebhookConfigurationListOutput)
}

type MutatingWebhookConfigurationListMapOutput struct{ *pulumi.OutputState }

func (MutatingWebhookConfigurationListMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*MutatingWebhookConfigurationList)(nil)).Elem()
}

func (o MutatingWebhookConfigurationListMapOutput) ToMutatingWebhookConfigurationListMapOutput() MutatingWebhookConfigurationListMapOutput {
	return o
}

func (o MutatingWebhookConfigurationListMapOutput) ToMutatingWebhookConfigurationListMapOutputWithContext(ctx context.Context) MutatingWebhookConfigurationListMapOutput {
	return o
}

func (o MutatingWebhookConfigurationListMapOutput) MapIndex(k pulumi.StringInput) MutatingWebhookConfigurationListOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *MutatingWebhookConfigurationList {
		return vs[0].(map[string]*MutatingWebhookConfigurationList)[vs[1].(string)]
	}).(MutatingWebhookConfigurationListOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*MutatingWebhookConfigurationListInput)(nil)).Elem(), &MutatingWebhookConfigurationList{})
	pulumi.RegisterInputType(reflect.TypeOf((*MutatingWebhookConfigurationListArrayInput)(nil)).Elem(), MutatingWebhookConfigurationListArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*MutatingWebhookConfigurationListMapInput)(nil)).Elem(), MutatingWebhookConfigurationListMap{})
	pulumi.RegisterOutputType(MutatingWebhookConfigurationListOutput{})
	pulumi.RegisterOutputType(MutatingWebhookConfigurationListArrayOutput{})
	pulumi.RegisterOutputType(MutatingWebhookConfigurationListMapOutput{})
}
