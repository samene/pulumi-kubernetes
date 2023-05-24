// Code generated by pulumigen DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v1

import (
	"context"
	"reflect"

	"errors"
	metav1 "github.com/pulumi/pulumi-kubernetes/sdk/v3/go/kubernetes/meta/v1"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// LeaseList is a list of Lease objects.
type LeaseList struct {
	pulumi.CustomResourceState

	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion pulumi.StringOutput `pulumi:"apiVersion"`
	// items is a list of schema objects.
	Items LeaseTypeArrayOutput `pulumi:"items"`
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind pulumi.StringOutput `pulumi:"kind"`
	// Standard list metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	Metadata metav1.ListMetaOutput `pulumi:"metadata"`
}

// NewLeaseList registers a new resource with the given unique name, arguments, and options.
func NewLeaseList(ctx *pulumi.Context,
	name string, args *LeaseListArgs, opts ...pulumi.ResourceOption) (*LeaseList, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Items == nil {
		return nil, errors.New("invalid value for required argument 'Items'")
	}
	args.ApiVersion = pulumi.StringPtr("coordination.k8s.io/v1")
	args.Kind = pulumi.StringPtr("LeaseList")
	var resource LeaseList
	err := ctx.RegisterResource("kubernetes:coordination.k8s.io/v1:LeaseList", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetLeaseList gets an existing LeaseList resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetLeaseList(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *LeaseListState, opts ...pulumi.ResourceOption) (*LeaseList, error) {
	var resource LeaseList
	err := ctx.ReadResource("kubernetes:coordination.k8s.io/v1:LeaseList", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering LeaseList resources.
type leaseListState struct {
}

type LeaseListState struct {
}

func (LeaseListState) ElementType() reflect.Type {
	return reflect.TypeOf((*leaseListState)(nil)).Elem()
}

type leaseListArgs struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion *string `pulumi:"apiVersion"`
	// items is a list of schema objects.
	Items []LeaseType `pulumi:"items"`
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind *string `pulumi:"kind"`
	// Standard list metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	Metadata *metav1.ListMeta `pulumi:"metadata"`
}

// The set of arguments for constructing a LeaseList resource.
type LeaseListArgs struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion pulumi.StringPtrInput
	// items is a list of schema objects.
	Items LeaseTypeArrayInput
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind pulumi.StringPtrInput
	// Standard list metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	Metadata metav1.ListMetaPtrInput
}

func (LeaseListArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*leaseListArgs)(nil)).Elem()
}

type LeaseListInput interface {
	pulumi.Input

	ToLeaseListOutput() LeaseListOutput
	ToLeaseListOutputWithContext(ctx context.Context) LeaseListOutput
}

func (*LeaseList) ElementType() reflect.Type {
	return reflect.TypeOf((**LeaseList)(nil)).Elem()
}

func (i *LeaseList) ToLeaseListOutput() LeaseListOutput {
	return i.ToLeaseListOutputWithContext(context.Background())
}

func (i *LeaseList) ToLeaseListOutputWithContext(ctx context.Context) LeaseListOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LeaseListOutput)
}

// LeaseListArrayInput is an input type that accepts LeaseListArray and LeaseListArrayOutput values.
// You can construct a concrete instance of `LeaseListArrayInput` via:
//
//	LeaseListArray{ LeaseListArgs{...} }
type LeaseListArrayInput interface {
	pulumi.Input

	ToLeaseListArrayOutput() LeaseListArrayOutput
	ToLeaseListArrayOutputWithContext(context.Context) LeaseListArrayOutput
}

type LeaseListArray []LeaseListInput

func (LeaseListArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*LeaseList)(nil)).Elem()
}

func (i LeaseListArray) ToLeaseListArrayOutput() LeaseListArrayOutput {
	return i.ToLeaseListArrayOutputWithContext(context.Background())
}

func (i LeaseListArray) ToLeaseListArrayOutputWithContext(ctx context.Context) LeaseListArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LeaseListArrayOutput)
}

// LeaseListMapInput is an input type that accepts LeaseListMap and LeaseListMapOutput values.
// You can construct a concrete instance of `LeaseListMapInput` via:
//
//	LeaseListMap{ "key": LeaseListArgs{...} }
type LeaseListMapInput interface {
	pulumi.Input

	ToLeaseListMapOutput() LeaseListMapOutput
	ToLeaseListMapOutputWithContext(context.Context) LeaseListMapOutput
}

type LeaseListMap map[string]LeaseListInput

func (LeaseListMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*LeaseList)(nil)).Elem()
}

func (i LeaseListMap) ToLeaseListMapOutput() LeaseListMapOutput {
	return i.ToLeaseListMapOutputWithContext(context.Background())
}

func (i LeaseListMap) ToLeaseListMapOutputWithContext(ctx context.Context) LeaseListMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LeaseListMapOutput)
}

type LeaseListOutput struct{ *pulumi.OutputState }

func (LeaseListOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**LeaseList)(nil)).Elem()
}

func (o LeaseListOutput) ToLeaseListOutput() LeaseListOutput {
	return o
}

func (o LeaseListOutput) ToLeaseListOutputWithContext(ctx context.Context) LeaseListOutput {
	return o
}

// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
func (o LeaseListOutput) ApiVersion() pulumi.StringOutput {
	return o.ApplyT(func(v *LeaseList) pulumi.StringOutput { return v.ApiVersion }).(pulumi.StringOutput)
}

// items is a list of schema objects.
func (o LeaseListOutput) Items() LeaseTypeArrayOutput {
	return o.ApplyT(func(v *LeaseList) LeaseTypeArrayOutput { return v.Items }).(LeaseTypeArrayOutput)
}

// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
func (o LeaseListOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v *LeaseList) pulumi.StringOutput { return v.Kind }).(pulumi.StringOutput)
}

// Standard list metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
func (o LeaseListOutput) Metadata() metav1.ListMetaOutput {
	return o.ApplyT(func(v *LeaseList) metav1.ListMetaOutput { return v.Metadata }).(metav1.ListMetaOutput)
}

type LeaseListArrayOutput struct{ *pulumi.OutputState }

func (LeaseListArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*LeaseList)(nil)).Elem()
}

func (o LeaseListArrayOutput) ToLeaseListArrayOutput() LeaseListArrayOutput {
	return o
}

func (o LeaseListArrayOutput) ToLeaseListArrayOutputWithContext(ctx context.Context) LeaseListArrayOutput {
	return o
}

func (o LeaseListArrayOutput) Index(i pulumi.IntInput) LeaseListOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *LeaseList {
		return vs[0].([]*LeaseList)[vs[1].(int)]
	}).(LeaseListOutput)
}

type LeaseListMapOutput struct{ *pulumi.OutputState }

func (LeaseListMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*LeaseList)(nil)).Elem()
}

func (o LeaseListMapOutput) ToLeaseListMapOutput() LeaseListMapOutput {
	return o
}

func (o LeaseListMapOutput) ToLeaseListMapOutputWithContext(ctx context.Context) LeaseListMapOutput {
	return o
}

func (o LeaseListMapOutput) MapIndex(k pulumi.StringInput) LeaseListOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *LeaseList {
		return vs[0].(map[string]*LeaseList)[vs[1].(string)]
	}).(LeaseListOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*LeaseListInput)(nil)).Elem(), &LeaseList{})
	pulumi.RegisterInputType(reflect.TypeOf((*LeaseListArrayInput)(nil)).Elem(), LeaseListArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*LeaseListMapInput)(nil)).Elem(), LeaseListMap{})
	pulumi.RegisterOutputType(LeaseListOutput{})
	pulumi.RegisterOutputType(LeaseListArrayOutput{})
	pulumi.RegisterOutputType(LeaseListMapOutput{})
}
