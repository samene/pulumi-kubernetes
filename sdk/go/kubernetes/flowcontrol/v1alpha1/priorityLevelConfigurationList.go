// Code generated by pulumigen DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v1alpha1

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	metav1 "github.com/pulumi/pulumi-kubernetes/sdk/v3/go/kubernetes/meta/v1"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// PriorityLevelConfigurationList is a list of PriorityLevelConfiguration objects.
type PriorityLevelConfigurationList struct {
	pulumi.CustomResourceState

	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion pulumi.StringPtrOutput `pulumi:"apiVersion"`
	// `items` is a list of request-priorities.
	Items PriorityLevelConfigurationTypeArrayOutput `pulumi:"items"`
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind pulumi.StringPtrOutput `pulumi:"kind"`
	// `metadata` is the standard object's metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	Metadata metav1.ListMetaPtrOutput `pulumi:"metadata"`
}

// NewPriorityLevelConfigurationList registers a new resource with the given unique name, arguments, and options.
func NewPriorityLevelConfigurationList(ctx *pulumi.Context,
	name string, args *PriorityLevelConfigurationListArgs, opts ...pulumi.ResourceOption) (*PriorityLevelConfigurationList, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Items == nil {
		return nil, errors.New("invalid value for required argument 'Items'")
	}
	args.ApiVersion = pulumi.StringPtr("flowcontrol.apiserver.k8s.io/v1alpha1")
	args.Kind = pulumi.StringPtr("PriorityLevelConfigurationList")
	var resource PriorityLevelConfigurationList
	err := ctx.RegisterResource("kubernetes:flowcontrol.apiserver.k8s.io/v1alpha1:PriorityLevelConfigurationList", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetPriorityLevelConfigurationList gets an existing PriorityLevelConfigurationList resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetPriorityLevelConfigurationList(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *PriorityLevelConfigurationListState, opts ...pulumi.ResourceOption) (*PriorityLevelConfigurationList, error) {
	var resource PriorityLevelConfigurationList
	err := ctx.ReadResource("kubernetes:flowcontrol.apiserver.k8s.io/v1alpha1:PriorityLevelConfigurationList", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering PriorityLevelConfigurationList resources.
type priorityLevelConfigurationListState struct {
}

type PriorityLevelConfigurationListState struct {
}

func (PriorityLevelConfigurationListState) ElementType() reflect.Type {
	return reflect.TypeOf((*priorityLevelConfigurationListState)(nil)).Elem()
}

type priorityLevelConfigurationListArgs struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion *string `pulumi:"apiVersion"`
	// `items` is a list of request-priorities.
	Items []PriorityLevelConfigurationType `pulumi:"items"`
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind *string `pulumi:"kind"`
	// `metadata` is the standard object's metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	Metadata *metav1.ListMeta `pulumi:"metadata"`
}

// The set of arguments for constructing a PriorityLevelConfigurationList resource.
type PriorityLevelConfigurationListArgs struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion pulumi.StringPtrInput
	// `items` is a list of request-priorities.
	Items PriorityLevelConfigurationTypeArrayInput
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind pulumi.StringPtrInput
	// `metadata` is the standard object's metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	Metadata metav1.ListMetaPtrInput
}

func (PriorityLevelConfigurationListArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*priorityLevelConfigurationListArgs)(nil)).Elem()
}

type PriorityLevelConfigurationListInput interface {
	pulumi.Input

	ToPriorityLevelConfigurationListOutput() PriorityLevelConfigurationListOutput
	ToPriorityLevelConfigurationListOutputWithContext(ctx context.Context) PriorityLevelConfigurationListOutput
}

func (*PriorityLevelConfigurationList) ElementType() reflect.Type {
	return reflect.TypeOf((**PriorityLevelConfigurationList)(nil)).Elem()
}

func (i *PriorityLevelConfigurationList) ToPriorityLevelConfigurationListOutput() PriorityLevelConfigurationListOutput {
	return i.ToPriorityLevelConfigurationListOutputWithContext(context.Background())
}

func (i *PriorityLevelConfigurationList) ToPriorityLevelConfigurationListOutputWithContext(ctx context.Context) PriorityLevelConfigurationListOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PriorityLevelConfigurationListOutput)
}

// PriorityLevelConfigurationListArrayInput is an input type that accepts PriorityLevelConfigurationListArray and PriorityLevelConfigurationListArrayOutput values.
// You can construct a concrete instance of `PriorityLevelConfigurationListArrayInput` via:
//
//          PriorityLevelConfigurationListArray{ PriorityLevelConfigurationListArgs{...} }
type PriorityLevelConfigurationListArrayInput interface {
	pulumi.Input

	ToPriorityLevelConfigurationListArrayOutput() PriorityLevelConfigurationListArrayOutput
	ToPriorityLevelConfigurationListArrayOutputWithContext(context.Context) PriorityLevelConfigurationListArrayOutput
}

type PriorityLevelConfigurationListArray []PriorityLevelConfigurationListInput

func (PriorityLevelConfigurationListArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*PriorityLevelConfigurationList)(nil)).Elem()
}

func (i PriorityLevelConfigurationListArray) ToPriorityLevelConfigurationListArrayOutput() PriorityLevelConfigurationListArrayOutput {
	return i.ToPriorityLevelConfigurationListArrayOutputWithContext(context.Background())
}

func (i PriorityLevelConfigurationListArray) ToPriorityLevelConfigurationListArrayOutputWithContext(ctx context.Context) PriorityLevelConfigurationListArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PriorityLevelConfigurationListArrayOutput)
}

// PriorityLevelConfigurationListMapInput is an input type that accepts PriorityLevelConfigurationListMap and PriorityLevelConfigurationListMapOutput values.
// You can construct a concrete instance of `PriorityLevelConfigurationListMapInput` via:
//
//          PriorityLevelConfigurationListMap{ "key": PriorityLevelConfigurationListArgs{...} }
type PriorityLevelConfigurationListMapInput interface {
	pulumi.Input

	ToPriorityLevelConfigurationListMapOutput() PriorityLevelConfigurationListMapOutput
	ToPriorityLevelConfigurationListMapOutputWithContext(context.Context) PriorityLevelConfigurationListMapOutput
}

type PriorityLevelConfigurationListMap map[string]PriorityLevelConfigurationListInput

func (PriorityLevelConfigurationListMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*PriorityLevelConfigurationList)(nil)).Elem()
}

func (i PriorityLevelConfigurationListMap) ToPriorityLevelConfigurationListMapOutput() PriorityLevelConfigurationListMapOutput {
	return i.ToPriorityLevelConfigurationListMapOutputWithContext(context.Background())
}

func (i PriorityLevelConfigurationListMap) ToPriorityLevelConfigurationListMapOutputWithContext(ctx context.Context) PriorityLevelConfigurationListMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PriorityLevelConfigurationListMapOutput)
}

type PriorityLevelConfigurationListOutput struct{ *pulumi.OutputState }

func (PriorityLevelConfigurationListOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PriorityLevelConfigurationList)(nil)).Elem()
}

func (o PriorityLevelConfigurationListOutput) ToPriorityLevelConfigurationListOutput() PriorityLevelConfigurationListOutput {
	return o
}

func (o PriorityLevelConfigurationListOutput) ToPriorityLevelConfigurationListOutputWithContext(ctx context.Context) PriorityLevelConfigurationListOutput {
	return o
}

type PriorityLevelConfigurationListArrayOutput struct{ *pulumi.OutputState }

func (PriorityLevelConfigurationListArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*PriorityLevelConfigurationList)(nil)).Elem()
}

func (o PriorityLevelConfigurationListArrayOutput) ToPriorityLevelConfigurationListArrayOutput() PriorityLevelConfigurationListArrayOutput {
	return o
}

func (o PriorityLevelConfigurationListArrayOutput) ToPriorityLevelConfigurationListArrayOutputWithContext(ctx context.Context) PriorityLevelConfigurationListArrayOutput {
	return o
}

func (o PriorityLevelConfigurationListArrayOutput) Index(i pulumi.IntInput) PriorityLevelConfigurationListOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *PriorityLevelConfigurationList {
		return vs[0].([]*PriorityLevelConfigurationList)[vs[1].(int)]
	}).(PriorityLevelConfigurationListOutput)
}

type PriorityLevelConfigurationListMapOutput struct{ *pulumi.OutputState }

func (PriorityLevelConfigurationListMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*PriorityLevelConfigurationList)(nil)).Elem()
}

func (o PriorityLevelConfigurationListMapOutput) ToPriorityLevelConfigurationListMapOutput() PriorityLevelConfigurationListMapOutput {
	return o
}

func (o PriorityLevelConfigurationListMapOutput) ToPriorityLevelConfigurationListMapOutputWithContext(ctx context.Context) PriorityLevelConfigurationListMapOutput {
	return o
}

func (o PriorityLevelConfigurationListMapOutput) MapIndex(k pulumi.StringInput) PriorityLevelConfigurationListOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *PriorityLevelConfigurationList {
		return vs[0].(map[string]*PriorityLevelConfigurationList)[vs[1].(string)]
	}).(PriorityLevelConfigurationListOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*PriorityLevelConfigurationListInput)(nil)).Elem(), &PriorityLevelConfigurationList{})
	pulumi.RegisterInputType(reflect.TypeOf((*PriorityLevelConfigurationListArrayInput)(nil)).Elem(), PriorityLevelConfigurationListArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*PriorityLevelConfigurationListMapInput)(nil)).Elem(), PriorityLevelConfigurationListMap{})
	pulumi.RegisterOutputType(PriorityLevelConfigurationListOutput{})
	pulumi.RegisterOutputType(PriorityLevelConfigurationListArrayOutput{})
	pulumi.RegisterOutputType(PriorityLevelConfigurationListMapOutput{})
}
