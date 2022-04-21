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

// PriorityClassList is a collection of priority classes.
type PriorityClassList struct {
	pulumi.CustomResourceState

	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion pulumi.StringPtrOutput `pulumi:"apiVersion"`
	// items is the list of PriorityClasses
	Items PriorityClassTypeArrayOutput `pulumi:"items"`
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind pulumi.StringPtrOutput `pulumi:"kind"`
	// Standard list metadata More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	Metadata metav1.ListMetaPtrOutput `pulumi:"metadata"`
}

// NewPriorityClassList registers a new resource with the given unique name, arguments, and options.
func NewPriorityClassList(ctx *pulumi.Context,
	name string, args *PriorityClassListArgs, opts ...pulumi.ResourceOption) (*PriorityClassList, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Items == nil {
		return nil, errors.New("invalid value for required argument 'Items'")
	}
	args.ApiVersion = pulumi.StringPtr("scheduling.k8s.io/v1alpha1")
	args.Kind = pulumi.StringPtr("PriorityClassList")
	var resource PriorityClassList
	err := ctx.RegisterResource("kubernetes:scheduling.k8s.io/v1alpha1:PriorityClassList", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetPriorityClassList gets an existing PriorityClassList resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetPriorityClassList(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *PriorityClassListState, opts ...pulumi.ResourceOption) (*PriorityClassList, error) {
	var resource PriorityClassList
	err := ctx.ReadResource("kubernetes:scheduling.k8s.io/v1alpha1:PriorityClassList", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering PriorityClassList resources.
type priorityClassListState struct {
}

type PriorityClassListState struct {
}

func (PriorityClassListState) ElementType() reflect.Type {
	return reflect.TypeOf((*priorityClassListState)(nil)).Elem()
}

type priorityClassListArgs struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion *string `pulumi:"apiVersion"`
	// items is the list of PriorityClasses
	Items []PriorityClassType `pulumi:"items"`
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind *string `pulumi:"kind"`
	// Standard list metadata More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	Metadata *metav1.ListMeta `pulumi:"metadata"`
}

// The set of arguments for constructing a PriorityClassList resource.
type PriorityClassListArgs struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion pulumi.StringPtrInput
	// items is the list of PriorityClasses
	Items PriorityClassTypeArrayInput
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind pulumi.StringPtrInput
	// Standard list metadata More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	Metadata metav1.ListMetaPtrInput
}

func (PriorityClassListArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*priorityClassListArgs)(nil)).Elem()
}

type PriorityClassListInput interface {
	pulumi.Input

	ToPriorityClassListOutput() PriorityClassListOutput
	ToPriorityClassListOutputWithContext(ctx context.Context) PriorityClassListOutput
}

func (*PriorityClassList) ElementType() reflect.Type {
	return reflect.TypeOf((**PriorityClassList)(nil)).Elem()
}

func (i *PriorityClassList) ToPriorityClassListOutput() PriorityClassListOutput {
	return i.ToPriorityClassListOutputWithContext(context.Background())
}

func (i *PriorityClassList) ToPriorityClassListOutputWithContext(ctx context.Context) PriorityClassListOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PriorityClassListOutput)
}

// PriorityClassListArrayInput is an input type that accepts PriorityClassListArray and PriorityClassListArrayOutput values.
// You can construct a concrete instance of `PriorityClassListArrayInput` via:
//
//          PriorityClassListArray{ PriorityClassListArgs{...} }
type PriorityClassListArrayInput interface {
	pulumi.Input

	ToPriorityClassListArrayOutput() PriorityClassListArrayOutput
	ToPriorityClassListArrayOutputWithContext(context.Context) PriorityClassListArrayOutput
}

type PriorityClassListArray []PriorityClassListInput

func (PriorityClassListArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*PriorityClassList)(nil)).Elem()
}

func (i PriorityClassListArray) ToPriorityClassListArrayOutput() PriorityClassListArrayOutput {
	return i.ToPriorityClassListArrayOutputWithContext(context.Background())
}

func (i PriorityClassListArray) ToPriorityClassListArrayOutputWithContext(ctx context.Context) PriorityClassListArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PriorityClassListArrayOutput)
}

// PriorityClassListMapInput is an input type that accepts PriorityClassListMap and PriorityClassListMapOutput values.
// You can construct a concrete instance of `PriorityClassListMapInput` via:
//
//          PriorityClassListMap{ "key": PriorityClassListArgs{...} }
type PriorityClassListMapInput interface {
	pulumi.Input

	ToPriorityClassListMapOutput() PriorityClassListMapOutput
	ToPriorityClassListMapOutputWithContext(context.Context) PriorityClassListMapOutput
}

type PriorityClassListMap map[string]PriorityClassListInput

func (PriorityClassListMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*PriorityClassList)(nil)).Elem()
}

func (i PriorityClassListMap) ToPriorityClassListMapOutput() PriorityClassListMapOutput {
	return i.ToPriorityClassListMapOutputWithContext(context.Background())
}

func (i PriorityClassListMap) ToPriorityClassListMapOutputWithContext(ctx context.Context) PriorityClassListMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PriorityClassListMapOutput)
}

type PriorityClassListOutput struct{ *pulumi.OutputState }

func (PriorityClassListOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PriorityClassList)(nil)).Elem()
}

func (o PriorityClassListOutput) ToPriorityClassListOutput() PriorityClassListOutput {
	return o
}

func (o PriorityClassListOutput) ToPriorityClassListOutputWithContext(ctx context.Context) PriorityClassListOutput {
	return o
}

type PriorityClassListArrayOutput struct{ *pulumi.OutputState }

func (PriorityClassListArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*PriorityClassList)(nil)).Elem()
}

func (o PriorityClassListArrayOutput) ToPriorityClassListArrayOutput() PriorityClassListArrayOutput {
	return o
}

func (o PriorityClassListArrayOutput) ToPriorityClassListArrayOutputWithContext(ctx context.Context) PriorityClassListArrayOutput {
	return o
}

func (o PriorityClassListArrayOutput) Index(i pulumi.IntInput) PriorityClassListOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *PriorityClassList {
		return vs[0].([]*PriorityClassList)[vs[1].(int)]
	}).(PriorityClassListOutput)
}

type PriorityClassListMapOutput struct{ *pulumi.OutputState }

func (PriorityClassListMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*PriorityClassList)(nil)).Elem()
}

func (o PriorityClassListMapOutput) ToPriorityClassListMapOutput() PriorityClassListMapOutput {
	return o
}

func (o PriorityClassListMapOutput) ToPriorityClassListMapOutputWithContext(ctx context.Context) PriorityClassListMapOutput {
	return o
}

func (o PriorityClassListMapOutput) MapIndex(k pulumi.StringInput) PriorityClassListOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *PriorityClassList {
		return vs[0].(map[string]*PriorityClassList)[vs[1].(string)]
	}).(PriorityClassListOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*PriorityClassListInput)(nil)).Elem(), &PriorityClassList{})
	pulumi.RegisterInputType(reflect.TypeOf((*PriorityClassListArrayInput)(nil)).Elem(), PriorityClassListArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*PriorityClassListMapInput)(nil)).Elem(), PriorityClassListMap{})
	pulumi.RegisterOutputType(PriorityClassListOutput{})
	pulumi.RegisterOutputType(PriorityClassListArrayOutput{})
	pulumi.RegisterOutputType(PriorityClassListMapOutput{})
}
