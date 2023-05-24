// Code generated by pulumigen DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v1alpha1

import (
	"context"
	"reflect"

	"errors"
	metav1 "github.com/pulumi/pulumi-kubernetes/sdk/v3/go/kubernetes/meta/v1"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// ClusterTrustBundleList is a collection of ClusterTrustBundle objects
type ClusterTrustBundleList struct {
	pulumi.CustomResourceState

	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion pulumi.StringOutput `pulumi:"apiVersion"`
	// items is a collection of ClusterTrustBundle objects
	Items ClusterTrustBundleTypeArrayOutput `pulumi:"items"`
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind pulumi.StringOutput `pulumi:"kind"`
	// metadata contains the list metadata.
	Metadata metav1.ListMetaOutput `pulumi:"metadata"`
}

// NewClusterTrustBundleList registers a new resource with the given unique name, arguments, and options.
func NewClusterTrustBundleList(ctx *pulumi.Context,
	name string, args *ClusterTrustBundleListArgs, opts ...pulumi.ResourceOption) (*ClusterTrustBundleList, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Items == nil {
		return nil, errors.New("invalid value for required argument 'Items'")
	}
	args.ApiVersion = pulumi.StringPtr("certificates.k8s.io/v1alpha1")
	args.Kind = pulumi.StringPtr("ClusterTrustBundleList")
	var resource ClusterTrustBundleList
	err := ctx.RegisterResource("kubernetes:certificates.k8s.io/v1alpha1:ClusterTrustBundleList", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetClusterTrustBundleList gets an existing ClusterTrustBundleList resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetClusterTrustBundleList(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ClusterTrustBundleListState, opts ...pulumi.ResourceOption) (*ClusterTrustBundleList, error) {
	var resource ClusterTrustBundleList
	err := ctx.ReadResource("kubernetes:certificates.k8s.io/v1alpha1:ClusterTrustBundleList", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ClusterTrustBundleList resources.
type clusterTrustBundleListState struct {
}

type ClusterTrustBundleListState struct {
}

func (ClusterTrustBundleListState) ElementType() reflect.Type {
	return reflect.TypeOf((*clusterTrustBundleListState)(nil)).Elem()
}

type clusterTrustBundleListArgs struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion *string `pulumi:"apiVersion"`
	// items is a collection of ClusterTrustBundle objects
	Items []ClusterTrustBundleType `pulumi:"items"`
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind *string `pulumi:"kind"`
	// metadata contains the list metadata.
	Metadata *metav1.ListMeta `pulumi:"metadata"`
}

// The set of arguments for constructing a ClusterTrustBundleList resource.
type ClusterTrustBundleListArgs struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion pulumi.StringPtrInput
	// items is a collection of ClusterTrustBundle objects
	Items ClusterTrustBundleTypeArrayInput
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind pulumi.StringPtrInput
	// metadata contains the list metadata.
	Metadata metav1.ListMetaPtrInput
}

func (ClusterTrustBundleListArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*clusterTrustBundleListArgs)(nil)).Elem()
}

type ClusterTrustBundleListInput interface {
	pulumi.Input

	ToClusterTrustBundleListOutput() ClusterTrustBundleListOutput
	ToClusterTrustBundleListOutputWithContext(ctx context.Context) ClusterTrustBundleListOutput
}

func (*ClusterTrustBundleList) ElementType() reflect.Type {
	return reflect.TypeOf((**ClusterTrustBundleList)(nil)).Elem()
}

func (i *ClusterTrustBundleList) ToClusterTrustBundleListOutput() ClusterTrustBundleListOutput {
	return i.ToClusterTrustBundleListOutputWithContext(context.Background())
}

func (i *ClusterTrustBundleList) ToClusterTrustBundleListOutputWithContext(ctx context.Context) ClusterTrustBundleListOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ClusterTrustBundleListOutput)
}

// ClusterTrustBundleListArrayInput is an input type that accepts ClusterTrustBundleListArray and ClusterTrustBundleListArrayOutput values.
// You can construct a concrete instance of `ClusterTrustBundleListArrayInput` via:
//
//	ClusterTrustBundleListArray{ ClusterTrustBundleListArgs{...} }
type ClusterTrustBundleListArrayInput interface {
	pulumi.Input

	ToClusterTrustBundleListArrayOutput() ClusterTrustBundleListArrayOutput
	ToClusterTrustBundleListArrayOutputWithContext(context.Context) ClusterTrustBundleListArrayOutput
}

type ClusterTrustBundleListArray []ClusterTrustBundleListInput

func (ClusterTrustBundleListArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ClusterTrustBundleList)(nil)).Elem()
}

func (i ClusterTrustBundleListArray) ToClusterTrustBundleListArrayOutput() ClusterTrustBundleListArrayOutput {
	return i.ToClusterTrustBundleListArrayOutputWithContext(context.Background())
}

func (i ClusterTrustBundleListArray) ToClusterTrustBundleListArrayOutputWithContext(ctx context.Context) ClusterTrustBundleListArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ClusterTrustBundleListArrayOutput)
}

// ClusterTrustBundleListMapInput is an input type that accepts ClusterTrustBundleListMap and ClusterTrustBundleListMapOutput values.
// You can construct a concrete instance of `ClusterTrustBundleListMapInput` via:
//
//	ClusterTrustBundleListMap{ "key": ClusterTrustBundleListArgs{...} }
type ClusterTrustBundleListMapInput interface {
	pulumi.Input

	ToClusterTrustBundleListMapOutput() ClusterTrustBundleListMapOutput
	ToClusterTrustBundleListMapOutputWithContext(context.Context) ClusterTrustBundleListMapOutput
}

type ClusterTrustBundleListMap map[string]ClusterTrustBundleListInput

func (ClusterTrustBundleListMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ClusterTrustBundleList)(nil)).Elem()
}

func (i ClusterTrustBundleListMap) ToClusterTrustBundleListMapOutput() ClusterTrustBundleListMapOutput {
	return i.ToClusterTrustBundleListMapOutputWithContext(context.Background())
}

func (i ClusterTrustBundleListMap) ToClusterTrustBundleListMapOutputWithContext(ctx context.Context) ClusterTrustBundleListMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ClusterTrustBundleListMapOutput)
}

type ClusterTrustBundleListOutput struct{ *pulumi.OutputState }

func (ClusterTrustBundleListOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ClusterTrustBundleList)(nil)).Elem()
}

func (o ClusterTrustBundleListOutput) ToClusterTrustBundleListOutput() ClusterTrustBundleListOutput {
	return o
}

func (o ClusterTrustBundleListOutput) ToClusterTrustBundleListOutputWithContext(ctx context.Context) ClusterTrustBundleListOutput {
	return o
}

// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
func (o ClusterTrustBundleListOutput) ApiVersion() pulumi.StringOutput {
	return o.ApplyT(func(v *ClusterTrustBundleList) pulumi.StringOutput { return v.ApiVersion }).(pulumi.StringOutput)
}

// items is a collection of ClusterTrustBundle objects
func (o ClusterTrustBundleListOutput) Items() ClusterTrustBundleTypeArrayOutput {
	return o.ApplyT(func(v *ClusterTrustBundleList) ClusterTrustBundleTypeArrayOutput { return v.Items }).(ClusterTrustBundleTypeArrayOutput)
}

// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
func (o ClusterTrustBundleListOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v *ClusterTrustBundleList) pulumi.StringOutput { return v.Kind }).(pulumi.StringOutput)
}

// metadata contains the list metadata.
func (o ClusterTrustBundleListOutput) Metadata() metav1.ListMetaOutput {
	return o.ApplyT(func(v *ClusterTrustBundleList) metav1.ListMetaOutput { return v.Metadata }).(metav1.ListMetaOutput)
}

type ClusterTrustBundleListArrayOutput struct{ *pulumi.OutputState }

func (ClusterTrustBundleListArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ClusterTrustBundleList)(nil)).Elem()
}

func (o ClusterTrustBundleListArrayOutput) ToClusterTrustBundleListArrayOutput() ClusterTrustBundleListArrayOutput {
	return o
}

func (o ClusterTrustBundleListArrayOutput) ToClusterTrustBundleListArrayOutputWithContext(ctx context.Context) ClusterTrustBundleListArrayOutput {
	return o
}

func (o ClusterTrustBundleListArrayOutput) Index(i pulumi.IntInput) ClusterTrustBundleListOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *ClusterTrustBundleList {
		return vs[0].([]*ClusterTrustBundleList)[vs[1].(int)]
	}).(ClusterTrustBundleListOutput)
}

type ClusterTrustBundleListMapOutput struct{ *pulumi.OutputState }

func (ClusterTrustBundleListMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ClusterTrustBundleList)(nil)).Elem()
}

func (o ClusterTrustBundleListMapOutput) ToClusterTrustBundleListMapOutput() ClusterTrustBundleListMapOutput {
	return o
}

func (o ClusterTrustBundleListMapOutput) ToClusterTrustBundleListMapOutputWithContext(ctx context.Context) ClusterTrustBundleListMapOutput {
	return o
}

func (o ClusterTrustBundleListMapOutput) MapIndex(k pulumi.StringInput) ClusterTrustBundleListOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *ClusterTrustBundleList {
		return vs[0].(map[string]*ClusterTrustBundleList)[vs[1].(string)]
	}).(ClusterTrustBundleListOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ClusterTrustBundleListInput)(nil)).Elem(), &ClusterTrustBundleList{})
	pulumi.RegisterInputType(reflect.TypeOf((*ClusterTrustBundleListArrayInput)(nil)).Elem(), ClusterTrustBundleListArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ClusterTrustBundleListMapInput)(nil)).Elem(), ClusterTrustBundleListMap{})
	pulumi.RegisterOutputType(ClusterTrustBundleListOutput{})
	pulumi.RegisterOutputType(ClusterTrustBundleListArrayOutput{})
	pulumi.RegisterOutputType(ClusterTrustBundleListMapOutput{})
}
