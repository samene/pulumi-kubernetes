// Code generated by pulumigen DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v1beta1

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	metav1 "github.com/pulumi/pulumi-kubernetes/sdk/v3/go/kubernetes/meta/v1"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// DeploymentList is a list of Deployments.
type DeploymentList struct {
	pulumi.CustomResourceState

	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion pulumi.StringPtrOutput `pulumi:"apiVersion"`
	// Items is the list of Deployments.
	Items DeploymentTypeArrayOutput `pulumi:"items"`
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind pulumi.StringPtrOutput `pulumi:"kind"`
	// Standard list metadata.
	Metadata metav1.ListMetaPtrOutput `pulumi:"metadata"`
}

// NewDeploymentList registers a new resource with the given unique name, arguments, and options.
func NewDeploymentList(ctx *pulumi.Context,
	name string, args *DeploymentListArgs, opts ...pulumi.ResourceOption) (*DeploymentList, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Items == nil {
		return nil, errors.New("invalid value for required argument 'Items'")
	}
	args.ApiVersion = pulumi.StringPtr("extensions/v1beta1")
	args.Kind = pulumi.StringPtr("DeploymentList")
	var resource DeploymentList
	err := ctx.RegisterResource("kubernetes:extensions/v1beta1:DeploymentList", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDeploymentList gets an existing DeploymentList resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDeploymentList(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DeploymentListState, opts ...pulumi.ResourceOption) (*DeploymentList, error) {
	var resource DeploymentList
	err := ctx.ReadResource("kubernetes:extensions/v1beta1:DeploymentList", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering DeploymentList resources.
type deploymentListState struct {
}

type DeploymentListState struct {
}

func (DeploymentListState) ElementType() reflect.Type {
	return reflect.TypeOf((*deploymentListState)(nil)).Elem()
}

type deploymentListArgs struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion *string `pulumi:"apiVersion"`
	// Items is the list of Deployments.
	Items []DeploymentType `pulumi:"items"`
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind *string `pulumi:"kind"`
	// Standard list metadata.
	Metadata *metav1.ListMeta `pulumi:"metadata"`
}

// The set of arguments for constructing a DeploymentList resource.
type DeploymentListArgs struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion pulumi.StringPtrInput
	// Items is the list of Deployments.
	Items DeploymentTypeArrayInput
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind pulumi.StringPtrInput
	// Standard list metadata.
	Metadata metav1.ListMetaPtrInput
}

func (DeploymentListArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*deploymentListArgs)(nil)).Elem()
}

type DeploymentListInput interface {
	pulumi.Input

	ToDeploymentListOutput() DeploymentListOutput
	ToDeploymentListOutputWithContext(ctx context.Context) DeploymentListOutput
}

func (*DeploymentList) ElementType() reflect.Type {
	return reflect.TypeOf((**DeploymentList)(nil)).Elem()
}

func (i *DeploymentList) ToDeploymentListOutput() DeploymentListOutput {
	return i.ToDeploymentListOutputWithContext(context.Background())
}

func (i *DeploymentList) ToDeploymentListOutputWithContext(ctx context.Context) DeploymentListOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DeploymentListOutput)
}

// DeploymentListArrayInput is an input type that accepts DeploymentListArray and DeploymentListArrayOutput values.
// You can construct a concrete instance of `DeploymentListArrayInput` via:
//
//          DeploymentListArray{ DeploymentListArgs{...} }
type DeploymentListArrayInput interface {
	pulumi.Input

	ToDeploymentListArrayOutput() DeploymentListArrayOutput
	ToDeploymentListArrayOutputWithContext(context.Context) DeploymentListArrayOutput
}

type DeploymentListArray []DeploymentListInput

func (DeploymentListArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*DeploymentList)(nil)).Elem()
}

func (i DeploymentListArray) ToDeploymentListArrayOutput() DeploymentListArrayOutput {
	return i.ToDeploymentListArrayOutputWithContext(context.Background())
}

func (i DeploymentListArray) ToDeploymentListArrayOutputWithContext(ctx context.Context) DeploymentListArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DeploymentListArrayOutput)
}

// DeploymentListMapInput is an input type that accepts DeploymentListMap and DeploymentListMapOutput values.
// You can construct a concrete instance of `DeploymentListMapInput` via:
//
//          DeploymentListMap{ "key": DeploymentListArgs{...} }
type DeploymentListMapInput interface {
	pulumi.Input

	ToDeploymentListMapOutput() DeploymentListMapOutput
	ToDeploymentListMapOutputWithContext(context.Context) DeploymentListMapOutput
}

type DeploymentListMap map[string]DeploymentListInput

func (DeploymentListMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*DeploymentList)(nil)).Elem()
}

func (i DeploymentListMap) ToDeploymentListMapOutput() DeploymentListMapOutput {
	return i.ToDeploymentListMapOutputWithContext(context.Background())
}

func (i DeploymentListMap) ToDeploymentListMapOutputWithContext(ctx context.Context) DeploymentListMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DeploymentListMapOutput)
}

type DeploymentListOutput struct{ *pulumi.OutputState }

func (DeploymentListOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DeploymentList)(nil)).Elem()
}

func (o DeploymentListOutput) ToDeploymentListOutput() DeploymentListOutput {
	return o
}

func (o DeploymentListOutput) ToDeploymentListOutputWithContext(ctx context.Context) DeploymentListOutput {
	return o
}

type DeploymentListArrayOutput struct{ *pulumi.OutputState }

func (DeploymentListArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*DeploymentList)(nil)).Elem()
}

func (o DeploymentListArrayOutput) ToDeploymentListArrayOutput() DeploymentListArrayOutput {
	return o
}

func (o DeploymentListArrayOutput) ToDeploymentListArrayOutputWithContext(ctx context.Context) DeploymentListArrayOutput {
	return o
}

func (o DeploymentListArrayOutput) Index(i pulumi.IntInput) DeploymentListOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *DeploymentList {
		return vs[0].([]*DeploymentList)[vs[1].(int)]
	}).(DeploymentListOutput)
}

type DeploymentListMapOutput struct{ *pulumi.OutputState }

func (DeploymentListMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*DeploymentList)(nil)).Elem()
}

func (o DeploymentListMapOutput) ToDeploymentListMapOutput() DeploymentListMapOutput {
	return o
}

func (o DeploymentListMapOutput) ToDeploymentListMapOutputWithContext(ctx context.Context) DeploymentListMapOutput {
	return o
}

func (o DeploymentListMapOutput) MapIndex(k pulumi.StringInput) DeploymentListOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *DeploymentList {
		return vs[0].(map[string]*DeploymentList)[vs[1].(string)]
	}).(DeploymentListOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*DeploymentListInput)(nil)).Elem(), &DeploymentList{})
	pulumi.RegisterInputType(reflect.TypeOf((*DeploymentListArrayInput)(nil)).Elem(), DeploymentListArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*DeploymentListMapInput)(nil)).Elem(), DeploymentListMap{})
	pulumi.RegisterOutputType(DeploymentListOutput{})
	pulumi.RegisterOutputType(DeploymentListArrayOutput{})
	pulumi.RegisterOutputType(DeploymentListMapOutput{})
}
