// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v1beta1

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	metav1 "github.com/pulumi/pulumi-kubernetes/sdk/v3/go/kubernetes/meta/v1"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// StatefulSetList is a collection of StatefulSets.
type StatefulSetList struct {
	pulumi.CustomResourceState

	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion pulumi.StringPtrOutput     `pulumi:"apiVersion"`
	Items      StatefulSetTypeArrayOutput `pulumi:"items"`
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind     pulumi.StringPtrOutput   `pulumi:"kind"`
	Metadata metav1.ListMetaPtrOutput `pulumi:"metadata"`
}

// NewStatefulSetList registers a new resource with the given unique name, arguments, and options.
func NewStatefulSetList(ctx *pulumi.Context,
	name string, args *StatefulSetListArgs, opts ...pulumi.ResourceOption) (*StatefulSetList, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Items == nil {
		return nil, errors.New("invalid value for required argument 'Items'")
	}
	args.ApiVersion = pulumi.StringPtr("apps/v1beta1")
	args.Kind = pulumi.StringPtr("StatefulSetList")
	var resource StatefulSetList
	err := ctx.RegisterResource("kubernetes:apps/v1beta1:StatefulSetList", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetStatefulSetList gets an existing StatefulSetList resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetStatefulSetList(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *StatefulSetListState, opts ...pulumi.ResourceOption) (*StatefulSetList, error) {
	var resource StatefulSetList
	err := ctx.ReadResource("kubernetes:apps/v1beta1:StatefulSetList", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering StatefulSetList resources.
type statefulSetListState struct {
}

type StatefulSetListState struct {
}

func (StatefulSetListState) ElementType() reflect.Type {
	return reflect.TypeOf((*statefulSetListState)(nil)).Elem()
}

type statefulSetListArgs struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion *string           `pulumi:"apiVersion"`
	Items      []StatefulSetType `pulumi:"items"`
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind     *string          `pulumi:"kind"`
	Metadata *metav1.ListMeta `pulumi:"metadata"`
}

// The set of arguments for constructing a StatefulSetList resource.
type StatefulSetListArgs struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion pulumi.StringPtrInput
	Items      StatefulSetTypeArrayInput
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind     pulumi.StringPtrInput
	Metadata metav1.ListMetaPtrInput
}

func (StatefulSetListArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*statefulSetListArgs)(nil)).Elem()
}

type StatefulSetListInput interface {
	pulumi.Input

	ToStatefulSetListOutput() StatefulSetListOutput
	ToStatefulSetListOutputWithContext(ctx context.Context) StatefulSetListOutput
}

func (*StatefulSetList) ElementType() reflect.Type {
	return reflect.TypeOf((**StatefulSetList)(nil)).Elem()
}

func (i *StatefulSetList) ToStatefulSetListOutput() StatefulSetListOutput {
	return i.ToStatefulSetListOutputWithContext(context.Background())
}

func (i *StatefulSetList) ToStatefulSetListOutputWithContext(ctx context.Context) StatefulSetListOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StatefulSetListOutput)
}

// StatefulSetListArrayInput is an input type that accepts StatefulSetListArray and StatefulSetListArrayOutput values.
// You can construct a concrete instance of `StatefulSetListArrayInput` via:
//
//          StatefulSetListArray{ StatefulSetListArgs{...} }
type StatefulSetListArrayInput interface {
	pulumi.Input

	ToStatefulSetListArrayOutput() StatefulSetListArrayOutput
	ToStatefulSetListArrayOutputWithContext(context.Context) StatefulSetListArrayOutput
}

type StatefulSetListArray []StatefulSetListInput

func (StatefulSetListArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*StatefulSetList)(nil)).Elem()
}

func (i StatefulSetListArray) ToStatefulSetListArrayOutput() StatefulSetListArrayOutput {
	return i.ToStatefulSetListArrayOutputWithContext(context.Background())
}

func (i StatefulSetListArray) ToStatefulSetListArrayOutputWithContext(ctx context.Context) StatefulSetListArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StatefulSetListArrayOutput)
}

// StatefulSetListMapInput is an input type that accepts StatefulSetListMap and StatefulSetListMapOutput values.
// You can construct a concrete instance of `StatefulSetListMapInput` via:
//
//          StatefulSetListMap{ "key": StatefulSetListArgs{...} }
type StatefulSetListMapInput interface {
	pulumi.Input

	ToStatefulSetListMapOutput() StatefulSetListMapOutput
	ToStatefulSetListMapOutputWithContext(context.Context) StatefulSetListMapOutput
}

type StatefulSetListMap map[string]StatefulSetListInput

func (StatefulSetListMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*StatefulSetList)(nil)).Elem()
}

func (i StatefulSetListMap) ToStatefulSetListMapOutput() StatefulSetListMapOutput {
	return i.ToStatefulSetListMapOutputWithContext(context.Background())
}

func (i StatefulSetListMap) ToStatefulSetListMapOutputWithContext(ctx context.Context) StatefulSetListMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StatefulSetListMapOutput)
}

type StatefulSetListOutput struct{ *pulumi.OutputState }

func (StatefulSetListOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**StatefulSetList)(nil)).Elem()
}

func (o StatefulSetListOutput) ToStatefulSetListOutput() StatefulSetListOutput {
	return o
}

func (o StatefulSetListOutput) ToStatefulSetListOutputWithContext(ctx context.Context) StatefulSetListOutput {
	return o
}

type StatefulSetListArrayOutput struct{ *pulumi.OutputState }

func (StatefulSetListArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*StatefulSetList)(nil)).Elem()
}

func (o StatefulSetListArrayOutput) ToStatefulSetListArrayOutput() StatefulSetListArrayOutput {
	return o
}

func (o StatefulSetListArrayOutput) ToStatefulSetListArrayOutputWithContext(ctx context.Context) StatefulSetListArrayOutput {
	return o
}

func (o StatefulSetListArrayOutput) Index(i pulumi.IntInput) StatefulSetListOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *StatefulSetList {
		return vs[0].([]*StatefulSetList)[vs[1].(int)]
	}).(StatefulSetListOutput)
}

type StatefulSetListMapOutput struct{ *pulumi.OutputState }

func (StatefulSetListMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*StatefulSetList)(nil)).Elem()
}

func (o StatefulSetListMapOutput) ToStatefulSetListMapOutput() StatefulSetListMapOutput {
	return o
}

func (o StatefulSetListMapOutput) ToStatefulSetListMapOutputWithContext(ctx context.Context) StatefulSetListMapOutput {
	return o
}

func (o StatefulSetListMapOutput) MapIndex(k pulumi.StringInput) StatefulSetListOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *StatefulSetList {
		return vs[0].(map[string]*StatefulSetList)[vs[1].(string)]
	}).(StatefulSetListOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*StatefulSetListInput)(nil)).Elem(), &StatefulSetList{})
	pulumi.RegisterInputType(reflect.TypeOf((*StatefulSetListArrayInput)(nil)).Elem(), StatefulSetListArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*StatefulSetListMapInput)(nil)).Elem(), StatefulSetListMap{})
	pulumi.RegisterOutputType(StatefulSetListOutput{})
	pulumi.RegisterOutputType(StatefulSetListArrayOutput{})
	pulumi.RegisterOutputType(StatefulSetListMapOutput{})
}
