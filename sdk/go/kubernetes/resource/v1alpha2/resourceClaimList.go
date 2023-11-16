// Code generated by pulumigen DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v1alpha2

import (
	"context"
	"reflect"

	"errors"
	metav1 "github.com/pulumi/pulumi-kubernetes/sdk/v4/go/kubernetes/meta/v1"
	"github.com/pulumi/pulumi-kubernetes/sdk/v4/go/kubernetes/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// ResourceClaimList is a collection of claims.
type ResourceClaimList struct {
	pulumi.CustomResourceState

	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion pulumi.StringOutput `pulumi:"apiVersion"`
	// Items is the list of resource claims.
	Items ResourceClaimTypeArrayOutput `pulumi:"items"`
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind pulumi.StringOutput `pulumi:"kind"`
	// Standard list metadata
	Metadata metav1.ListMetaOutput `pulumi:"metadata"`
}

// NewResourceClaimList registers a new resource with the given unique name, arguments, and options.
func NewResourceClaimList(ctx *pulumi.Context,
	name string, args *ResourceClaimListArgs, opts ...pulumi.ResourceOption) (*ResourceClaimList, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Items == nil {
		return nil, errors.New("invalid value for required argument 'Items'")
	}
	args.ApiVersion = pulumi.StringPtr("resource.k8s.io/v1alpha2")
	args.Kind = pulumi.StringPtr("ResourceClaimList")
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource ResourceClaimList
	err := ctx.RegisterResource("kubernetes:resource.k8s.io/v1alpha2:ResourceClaimList", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetResourceClaimList gets an existing ResourceClaimList resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetResourceClaimList(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ResourceClaimListState, opts ...pulumi.ResourceOption) (*ResourceClaimList, error) {
	var resource ResourceClaimList
	err := ctx.ReadResource("kubernetes:resource.k8s.io/v1alpha2:ResourceClaimList", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ResourceClaimList resources.
type resourceClaimListState struct {
}

type ResourceClaimListState struct {
}

func (ResourceClaimListState) ElementType() reflect.Type {
	return reflect.TypeOf((*resourceClaimListState)(nil)).Elem()
}

type resourceClaimListArgs struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion *string `pulumi:"apiVersion"`
	// Items is the list of resource claims.
	Items []ResourceClaimType `pulumi:"items"`
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind *string `pulumi:"kind"`
	// Standard list metadata
	Metadata *metav1.ListMeta `pulumi:"metadata"`
}

// The set of arguments for constructing a ResourceClaimList resource.
type ResourceClaimListArgs struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion pulumi.StringPtrInput
	// Items is the list of resource claims.
	Items ResourceClaimTypeArrayInput
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind pulumi.StringPtrInput
	// Standard list metadata
	Metadata metav1.ListMetaPtrInput
}

func (ResourceClaimListArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*resourceClaimListArgs)(nil)).Elem()
}

type ResourceClaimListInput interface {
	pulumi.Input

	ToResourceClaimListOutput() ResourceClaimListOutput
	ToResourceClaimListOutputWithContext(ctx context.Context) ResourceClaimListOutput
}

func (*ResourceClaimList) ElementType() reflect.Type {
	return reflect.TypeOf((**ResourceClaimList)(nil)).Elem()
}

func (i *ResourceClaimList) ToResourceClaimListOutput() ResourceClaimListOutput {
	return i.ToResourceClaimListOutputWithContext(context.Background())
}

func (i *ResourceClaimList) ToResourceClaimListOutputWithContext(ctx context.Context) ResourceClaimListOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ResourceClaimListOutput)
}

// ResourceClaimListArrayInput is an input type that accepts ResourceClaimListArray and ResourceClaimListArrayOutput values.
// You can construct a concrete instance of `ResourceClaimListArrayInput` via:
//
//	ResourceClaimListArray{ ResourceClaimListArgs{...} }
type ResourceClaimListArrayInput interface {
	pulumi.Input

	ToResourceClaimListArrayOutput() ResourceClaimListArrayOutput
	ToResourceClaimListArrayOutputWithContext(context.Context) ResourceClaimListArrayOutput
}

type ResourceClaimListArray []ResourceClaimListInput

func (ResourceClaimListArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ResourceClaimList)(nil)).Elem()
}

func (i ResourceClaimListArray) ToResourceClaimListArrayOutput() ResourceClaimListArrayOutput {
	return i.ToResourceClaimListArrayOutputWithContext(context.Background())
}

func (i ResourceClaimListArray) ToResourceClaimListArrayOutputWithContext(ctx context.Context) ResourceClaimListArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ResourceClaimListArrayOutput)
}

// ResourceClaimListMapInput is an input type that accepts ResourceClaimListMap and ResourceClaimListMapOutput values.
// You can construct a concrete instance of `ResourceClaimListMapInput` via:
//
//	ResourceClaimListMap{ "key": ResourceClaimListArgs{...} }
type ResourceClaimListMapInput interface {
	pulumi.Input

	ToResourceClaimListMapOutput() ResourceClaimListMapOutput
	ToResourceClaimListMapOutputWithContext(context.Context) ResourceClaimListMapOutput
}

type ResourceClaimListMap map[string]ResourceClaimListInput

func (ResourceClaimListMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ResourceClaimList)(nil)).Elem()
}

func (i ResourceClaimListMap) ToResourceClaimListMapOutput() ResourceClaimListMapOutput {
	return i.ToResourceClaimListMapOutputWithContext(context.Background())
}

func (i ResourceClaimListMap) ToResourceClaimListMapOutputWithContext(ctx context.Context) ResourceClaimListMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ResourceClaimListMapOutput)
}

type ResourceClaimListOutput struct{ *pulumi.OutputState }

func (ResourceClaimListOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ResourceClaimList)(nil)).Elem()
}

func (o ResourceClaimListOutput) ToResourceClaimListOutput() ResourceClaimListOutput {
	return o
}

func (o ResourceClaimListOutput) ToResourceClaimListOutputWithContext(ctx context.Context) ResourceClaimListOutput {
	return o
}

// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
func (o ResourceClaimListOutput) ApiVersion() pulumi.StringOutput {
	return o.ApplyT(func(v *ResourceClaimList) pulumi.StringOutput { return v.ApiVersion }).(pulumi.StringOutput)
}

// Items is the list of resource claims.
func (o ResourceClaimListOutput) Items() ResourceClaimTypeArrayOutput {
	return o.ApplyT(func(v *ResourceClaimList) ResourceClaimTypeArrayOutput { return v.Items }).(ResourceClaimTypeArrayOutput)
}

// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
func (o ResourceClaimListOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v *ResourceClaimList) pulumi.StringOutput { return v.Kind }).(pulumi.StringOutput)
}

// Standard list metadata
func (o ResourceClaimListOutput) Metadata() metav1.ListMetaOutput {
	return o.ApplyT(func(v *ResourceClaimList) metav1.ListMetaOutput { return v.Metadata }).(metav1.ListMetaOutput)
}

type ResourceClaimListArrayOutput struct{ *pulumi.OutputState }

func (ResourceClaimListArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ResourceClaimList)(nil)).Elem()
}

func (o ResourceClaimListArrayOutput) ToResourceClaimListArrayOutput() ResourceClaimListArrayOutput {
	return o
}

func (o ResourceClaimListArrayOutput) ToResourceClaimListArrayOutputWithContext(ctx context.Context) ResourceClaimListArrayOutput {
	return o
}

func (o ResourceClaimListArrayOutput) Index(i pulumi.IntInput) ResourceClaimListOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *ResourceClaimList {
		return vs[0].([]*ResourceClaimList)[vs[1].(int)]
	}).(ResourceClaimListOutput)
}

type ResourceClaimListMapOutput struct{ *pulumi.OutputState }

func (ResourceClaimListMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ResourceClaimList)(nil)).Elem()
}

func (o ResourceClaimListMapOutput) ToResourceClaimListMapOutput() ResourceClaimListMapOutput {
	return o
}

func (o ResourceClaimListMapOutput) ToResourceClaimListMapOutputWithContext(ctx context.Context) ResourceClaimListMapOutput {
	return o
}

func (o ResourceClaimListMapOutput) MapIndex(k pulumi.StringInput) ResourceClaimListOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *ResourceClaimList {
		return vs[0].(map[string]*ResourceClaimList)[vs[1].(string)]
	}).(ResourceClaimListOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ResourceClaimListInput)(nil)).Elem(), &ResourceClaimList{})
	pulumi.RegisterInputType(reflect.TypeOf((*ResourceClaimListArrayInput)(nil)).Elem(), ResourceClaimListArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ResourceClaimListMapInput)(nil)).Elem(), ResourceClaimListMap{})
	pulumi.RegisterOutputType(ResourceClaimListOutput{})
	pulumi.RegisterOutputType(ResourceClaimListArrayOutput{})
	pulumi.RegisterOutputType(ResourceClaimListMapOutput{})
}
