// Code generated by pulumigen DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v1

import (
	"context"
	"reflect"

	"errors"
	metav1 "github.com/pulumi/pulumi-kubernetes/sdk/v4/go/kubernetes/meta/v1"
	"github.com/pulumi/pulumi-kubernetes/sdk/v4/go/kubernetes/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// ResourceQuotaList is a list of ResourceQuota items.
type ResourceQuotaList struct {
	pulumi.CustomResourceState

	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion pulumi.StringOutput `pulumi:"apiVersion"`
	// Items is a list of ResourceQuota objects. More info: https://kubernetes.io/docs/concepts/policy/resource-quotas/
	Items ResourceQuotaTypeArrayOutput `pulumi:"items"`
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind pulumi.StringOutput `pulumi:"kind"`
	// Standard list metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Metadata metav1.ListMetaOutput `pulumi:"metadata"`
}

// NewResourceQuotaList registers a new resource with the given unique name, arguments, and options.
func NewResourceQuotaList(ctx *pulumi.Context,
	name string, args *ResourceQuotaListArgs, opts ...pulumi.ResourceOption) (*ResourceQuotaList, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Items == nil {
		return nil, errors.New("invalid value for required argument 'Items'")
	}
	args.ApiVersion = pulumi.StringPtr("v1")
	args.Kind = pulumi.StringPtr("ResourceQuotaList")
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource ResourceQuotaList
	err := ctx.RegisterResource("kubernetes:core/v1:ResourceQuotaList", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetResourceQuotaList gets an existing ResourceQuotaList resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetResourceQuotaList(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ResourceQuotaListState, opts ...pulumi.ResourceOption) (*ResourceQuotaList, error) {
	var resource ResourceQuotaList
	err := ctx.ReadResource("kubernetes:core/v1:ResourceQuotaList", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ResourceQuotaList resources.
type resourceQuotaListState struct {
}

type ResourceQuotaListState struct {
}

func (ResourceQuotaListState) ElementType() reflect.Type {
	return reflect.TypeOf((*resourceQuotaListState)(nil)).Elem()
}

type resourceQuotaListArgs struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion *string `pulumi:"apiVersion"`
	// Items is a list of ResourceQuota objects. More info: https://kubernetes.io/docs/concepts/policy/resource-quotas/
	Items []ResourceQuotaType `pulumi:"items"`
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind *string `pulumi:"kind"`
	// Standard list metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Metadata *metav1.ListMeta `pulumi:"metadata"`
}

// The set of arguments for constructing a ResourceQuotaList resource.
type ResourceQuotaListArgs struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion pulumi.StringPtrInput
	// Items is a list of ResourceQuota objects. More info: https://kubernetes.io/docs/concepts/policy/resource-quotas/
	Items ResourceQuotaTypeArrayInput
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind pulumi.StringPtrInput
	// Standard list metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Metadata metav1.ListMetaPtrInput
}

func (ResourceQuotaListArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*resourceQuotaListArgs)(nil)).Elem()
}

type ResourceQuotaListInput interface {
	pulumi.Input

	ToResourceQuotaListOutput() ResourceQuotaListOutput
	ToResourceQuotaListOutputWithContext(ctx context.Context) ResourceQuotaListOutput
}

func (*ResourceQuotaList) ElementType() reflect.Type {
	return reflect.TypeOf((**ResourceQuotaList)(nil)).Elem()
}

func (i *ResourceQuotaList) ToResourceQuotaListOutput() ResourceQuotaListOutput {
	return i.ToResourceQuotaListOutputWithContext(context.Background())
}

func (i *ResourceQuotaList) ToResourceQuotaListOutputWithContext(ctx context.Context) ResourceQuotaListOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ResourceQuotaListOutput)
}

// ResourceQuotaListArrayInput is an input type that accepts ResourceQuotaListArray and ResourceQuotaListArrayOutput values.
// You can construct a concrete instance of `ResourceQuotaListArrayInput` via:
//
//	ResourceQuotaListArray{ ResourceQuotaListArgs{...} }
type ResourceQuotaListArrayInput interface {
	pulumi.Input

	ToResourceQuotaListArrayOutput() ResourceQuotaListArrayOutput
	ToResourceQuotaListArrayOutputWithContext(context.Context) ResourceQuotaListArrayOutput
}

type ResourceQuotaListArray []ResourceQuotaListInput

func (ResourceQuotaListArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ResourceQuotaList)(nil)).Elem()
}

func (i ResourceQuotaListArray) ToResourceQuotaListArrayOutput() ResourceQuotaListArrayOutput {
	return i.ToResourceQuotaListArrayOutputWithContext(context.Background())
}

func (i ResourceQuotaListArray) ToResourceQuotaListArrayOutputWithContext(ctx context.Context) ResourceQuotaListArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ResourceQuotaListArrayOutput)
}

// ResourceQuotaListMapInput is an input type that accepts ResourceQuotaListMap and ResourceQuotaListMapOutput values.
// You can construct a concrete instance of `ResourceQuotaListMapInput` via:
//
//	ResourceQuotaListMap{ "key": ResourceQuotaListArgs{...} }
type ResourceQuotaListMapInput interface {
	pulumi.Input

	ToResourceQuotaListMapOutput() ResourceQuotaListMapOutput
	ToResourceQuotaListMapOutputWithContext(context.Context) ResourceQuotaListMapOutput
}

type ResourceQuotaListMap map[string]ResourceQuotaListInput

func (ResourceQuotaListMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ResourceQuotaList)(nil)).Elem()
}

func (i ResourceQuotaListMap) ToResourceQuotaListMapOutput() ResourceQuotaListMapOutput {
	return i.ToResourceQuotaListMapOutputWithContext(context.Background())
}

func (i ResourceQuotaListMap) ToResourceQuotaListMapOutputWithContext(ctx context.Context) ResourceQuotaListMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ResourceQuotaListMapOutput)
}

type ResourceQuotaListOutput struct{ *pulumi.OutputState }

func (ResourceQuotaListOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ResourceQuotaList)(nil)).Elem()
}

func (o ResourceQuotaListOutput) ToResourceQuotaListOutput() ResourceQuotaListOutput {
	return o
}

func (o ResourceQuotaListOutput) ToResourceQuotaListOutputWithContext(ctx context.Context) ResourceQuotaListOutput {
	return o
}

// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
func (o ResourceQuotaListOutput) ApiVersion() pulumi.StringOutput {
	return o.ApplyT(func(v *ResourceQuotaList) pulumi.StringOutput { return v.ApiVersion }).(pulumi.StringOutput)
}

// Items is a list of ResourceQuota objects. More info: https://kubernetes.io/docs/concepts/policy/resource-quotas/
func (o ResourceQuotaListOutput) Items() ResourceQuotaTypeArrayOutput {
	return o.ApplyT(func(v *ResourceQuotaList) ResourceQuotaTypeArrayOutput { return v.Items }).(ResourceQuotaTypeArrayOutput)
}

// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
func (o ResourceQuotaListOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v *ResourceQuotaList) pulumi.StringOutput { return v.Kind }).(pulumi.StringOutput)
}

// Standard list metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
func (o ResourceQuotaListOutput) Metadata() metav1.ListMetaOutput {
	return o.ApplyT(func(v *ResourceQuotaList) metav1.ListMetaOutput { return v.Metadata }).(metav1.ListMetaOutput)
}

type ResourceQuotaListArrayOutput struct{ *pulumi.OutputState }

func (ResourceQuotaListArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ResourceQuotaList)(nil)).Elem()
}

func (o ResourceQuotaListArrayOutput) ToResourceQuotaListArrayOutput() ResourceQuotaListArrayOutput {
	return o
}

func (o ResourceQuotaListArrayOutput) ToResourceQuotaListArrayOutputWithContext(ctx context.Context) ResourceQuotaListArrayOutput {
	return o
}

func (o ResourceQuotaListArrayOutput) Index(i pulumi.IntInput) ResourceQuotaListOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *ResourceQuotaList {
		return vs[0].([]*ResourceQuotaList)[vs[1].(int)]
	}).(ResourceQuotaListOutput)
}

type ResourceQuotaListMapOutput struct{ *pulumi.OutputState }

func (ResourceQuotaListMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ResourceQuotaList)(nil)).Elem()
}

func (o ResourceQuotaListMapOutput) ToResourceQuotaListMapOutput() ResourceQuotaListMapOutput {
	return o
}

func (o ResourceQuotaListMapOutput) ToResourceQuotaListMapOutputWithContext(ctx context.Context) ResourceQuotaListMapOutput {
	return o
}

func (o ResourceQuotaListMapOutput) MapIndex(k pulumi.StringInput) ResourceQuotaListOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *ResourceQuotaList {
		return vs[0].(map[string]*ResourceQuotaList)[vs[1].(string)]
	}).(ResourceQuotaListOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ResourceQuotaListInput)(nil)).Elem(), &ResourceQuotaList{})
	pulumi.RegisterInputType(reflect.TypeOf((*ResourceQuotaListArrayInput)(nil)).Elem(), ResourceQuotaListArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ResourceQuotaListMapInput)(nil)).Elem(), ResourceQuotaListMap{})
	pulumi.RegisterOutputType(ResourceQuotaListOutput{})
	pulumi.RegisterOutputType(ResourceQuotaListArrayOutput{})
	pulumi.RegisterOutputType(ResourceQuotaListMapOutput{})
}
