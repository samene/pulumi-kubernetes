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

// ReplicaSetList is a collection of ReplicaSets.
type ReplicaSetList struct {
	pulumi.CustomResourceState

	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion pulumi.StringOutput `pulumi:"apiVersion"`
	// List of ReplicaSets. More info: https://kubernetes.io/docs/concepts/workloads/controllers/replicationcontroller
	Items ReplicaSetTypeArrayOutput `pulumi:"items"`
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind pulumi.StringOutput `pulumi:"kind"`
	// Standard list metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Metadata metav1.ListMetaOutput `pulumi:"metadata"`
}

// NewReplicaSetList registers a new resource with the given unique name, arguments, and options.
func NewReplicaSetList(ctx *pulumi.Context,
	name string, args *ReplicaSetListArgs, opts ...pulumi.ResourceOption) (*ReplicaSetList, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Items == nil {
		return nil, errors.New("invalid value for required argument 'Items'")
	}
	args.ApiVersion = pulumi.StringPtr("extensions/v1beta1")
	args.Kind = pulumi.StringPtr("ReplicaSetList")
	var resource ReplicaSetList
	err := ctx.RegisterResource("kubernetes:extensions/v1beta1:ReplicaSetList", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetReplicaSetList gets an existing ReplicaSetList resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetReplicaSetList(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ReplicaSetListState, opts ...pulumi.ResourceOption) (*ReplicaSetList, error) {
	var resource ReplicaSetList
	err := ctx.ReadResource("kubernetes:extensions/v1beta1:ReplicaSetList", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ReplicaSetList resources.
type replicaSetListState struct {
}

type ReplicaSetListState struct {
}

func (ReplicaSetListState) ElementType() reflect.Type {
	return reflect.TypeOf((*replicaSetListState)(nil)).Elem()
}

type replicaSetListArgs struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion *string `pulumi:"apiVersion"`
	// List of ReplicaSets. More info: https://kubernetes.io/docs/concepts/workloads/controllers/replicationcontroller
	Items []ReplicaSetType `pulumi:"items"`
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind *string `pulumi:"kind"`
	// Standard list metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Metadata *metav1.ListMeta `pulumi:"metadata"`
}

// The set of arguments for constructing a ReplicaSetList resource.
type ReplicaSetListArgs struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion pulumi.StringPtrInput
	// List of ReplicaSets. More info: https://kubernetes.io/docs/concepts/workloads/controllers/replicationcontroller
	Items ReplicaSetTypeArrayInput
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind pulumi.StringPtrInput
	// Standard list metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Metadata metav1.ListMetaPtrInput
}

func (ReplicaSetListArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*replicaSetListArgs)(nil)).Elem()
}

type ReplicaSetListInput interface {
	pulumi.Input

	ToReplicaSetListOutput() ReplicaSetListOutput
	ToReplicaSetListOutputWithContext(ctx context.Context) ReplicaSetListOutput
}

func (*ReplicaSetList) ElementType() reflect.Type {
	return reflect.TypeOf((**ReplicaSetList)(nil)).Elem()
}

func (i *ReplicaSetList) ToReplicaSetListOutput() ReplicaSetListOutput {
	return i.ToReplicaSetListOutputWithContext(context.Background())
}

func (i *ReplicaSetList) ToReplicaSetListOutputWithContext(ctx context.Context) ReplicaSetListOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ReplicaSetListOutput)
}

// ReplicaSetListArrayInput is an input type that accepts ReplicaSetListArray and ReplicaSetListArrayOutput values.
// You can construct a concrete instance of `ReplicaSetListArrayInput` via:
//
//	ReplicaSetListArray{ ReplicaSetListArgs{...} }
type ReplicaSetListArrayInput interface {
	pulumi.Input

	ToReplicaSetListArrayOutput() ReplicaSetListArrayOutput
	ToReplicaSetListArrayOutputWithContext(context.Context) ReplicaSetListArrayOutput
}

type ReplicaSetListArray []ReplicaSetListInput

func (ReplicaSetListArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ReplicaSetList)(nil)).Elem()
}

func (i ReplicaSetListArray) ToReplicaSetListArrayOutput() ReplicaSetListArrayOutput {
	return i.ToReplicaSetListArrayOutputWithContext(context.Background())
}

func (i ReplicaSetListArray) ToReplicaSetListArrayOutputWithContext(ctx context.Context) ReplicaSetListArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ReplicaSetListArrayOutput)
}

// ReplicaSetListMapInput is an input type that accepts ReplicaSetListMap and ReplicaSetListMapOutput values.
// You can construct a concrete instance of `ReplicaSetListMapInput` via:
//
//	ReplicaSetListMap{ "key": ReplicaSetListArgs{...} }
type ReplicaSetListMapInput interface {
	pulumi.Input

	ToReplicaSetListMapOutput() ReplicaSetListMapOutput
	ToReplicaSetListMapOutputWithContext(context.Context) ReplicaSetListMapOutput
}

type ReplicaSetListMap map[string]ReplicaSetListInput

func (ReplicaSetListMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ReplicaSetList)(nil)).Elem()
}

func (i ReplicaSetListMap) ToReplicaSetListMapOutput() ReplicaSetListMapOutput {
	return i.ToReplicaSetListMapOutputWithContext(context.Background())
}

func (i ReplicaSetListMap) ToReplicaSetListMapOutputWithContext(ctx context.Context) ReplicaSetListMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ReplicaSetListMapOutput)
}

type ReplicaSetListOutput struct{ *pulumi.OutputState }

func (ReplicaSetListOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ReplicaSetList)(nil)).Elem()
}

func (o ReplicaSetListOutput) ToReplicaSetListOutput() ReplicaSetListOutput {
	return o
}

func (o ReplicaSetListOutput) ToReplicaSetListOutputWithContext(ctx context.Context) ReplicaSetListOutput {
	return o
}

// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
func (o ReplicaSetListOutput) ApiVersion() pulumi.StringOutput {
	return o.ApplyT(func(v *ReplicaSetList) pulumi.StringOutput { return v.ApiVersion }).(pulumi.StringOutput)
}

// List of ReplicaSets. More info: https://kubernetes.io/docs/concepts/workloads/controllers/replicationcontroller
func (o ReplicaSetListOutput) Items() ReplicaSetTypeArrayOutput {
	return o.ApplyT(func(v *ReplicaSetList) ReplicaSetTypeArrayOutput { return v.Items }).(ReplicaSetTypeArrayOutput)
}

// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
func (o ReplicaSetListOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v *ReplicaSetList) pulumi.StringOutput { return v.Kind }).(pulumi.StringOutput)
}

// Standard list metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
func (o ReplicaSetListOutput) Metadata() metav1.ListMetaOutput {
	return o.ApplyT(func(v *ReplicaSetList) metav1.ListMetaOutput { return v.Metadata }).(metav1.ListMetaOutput)
}

type ReplicaSetListArrayOutput struct{ *pulumi.OutputState }

func (ReplicaSetListArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ReplicaSetList)(nil)).Elem()
}

func (o ReplicaSetListArrayOutput) ToReplicaSetListArrayOutput() ReplicaSetListArrayOutput {
	return o
}

func (o ReplicaSetListArrayOutput) ToReplicaSetListArrayOutputWithContext(ctx context.Context) ReplicaSetListArrayOutput {
	return o
}

func (o ReplicaSetListArrayOutput) Index(i pulumi.IntInput) ReplicaSetListOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *ReplicaSetList {
		return vs[0].([]*ReplicaSetList)[vs[1].(int)]
	}).(ReplicaSetListOutput)
}

type ReplicaSetListMapOutput struct{ *pulumi.OutputState }

func (ReplicaSetListMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ReplicaSetList)(nil)).Elem()
}

func (o ReplicaSetListMapOutput) ToReplicaSetListMapOutput() ReplicaSetListMapOutput {
	return o
}

func (o ReplicaSetListMapOutput) ToReplicaSetListMapOutputWithContext(ctx context.Context) ReplicaSetListMapOutput {
	return o
}

func (o ReplicaSetListMapOutput) MapIndex(k pulumi.StringInput) ReplicaSetListOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *ReplicaSetList {
		return vs[0].(map[string]*ReplicaSetList)[vs[1].(string)]
	}).(ReplicaSetListOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ReplicaSetListInput)(nil)).Elem(), &ReplicaSetList{})
	pulumi.RegisterInputType(reflect.TypeOf((*ReplicaSetListArrayInput)(nil)).Elem(), ReplicaSetListArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ReplicaSetListMapInput)(nil)).Elem(), ReplicaSetListMap{})
	pulumi.RegisterOutputType(ReplicaSetListOutput{})
	pulumi.RegisterOutputType(ReplicaSetListArrayOutput{})
	pulumi.RegisterOutputType(ReplicaSetListMapOutput{})
}
