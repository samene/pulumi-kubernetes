// Code generated by pulumigen DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v1

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	metav1 "github.com/pulumi/pulumi-kubernetes/sdk/v3/go/kubernetes/meta/v1"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// NodeList is the whole list of all Nodes which have been registered with master.
type NodeList struct {
	pulumi.CustomResourceState

	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion pulumi.StringPtrOutput `pulumi:"apiVersion"`
	// List of nodes
	Items NodeTypeArrayOutput `pulumi:"items"`
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind pulumi.StringPtrOutput `pulumi:"kind"`
	// Standard list metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Metadata metav1.ListMetaPtrOutput `pulumi:"metadata"`
}

// NewNodeList registers a new resource with the given unique name, arguments, and options.
func NewNodeList(ctx *pulumi.Context,
	name string, args *NodeListArgs, opts ...pulumi.ResourceOption) (*NodeList, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Items == nil {
		return nil, errors.New("invalid value for required argument 'Items'")
	}
	args.ApiVersion = pulumi.StringPtr("v1")
	args.Kind = pulumi.StringPtr("NodeList")
	var resource NodeList
	err := ctx.RegisterResource("kubernetes:core/v1:NodeList", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetNodeList gets an existing NodeList resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetNodeList(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *NodeListState, opts ...pulumi.ResourceOption) (*NodeList, error) {
	var resource NodeList
	err := ctx.ReadResource("kubernetes:core/v1:NodeList", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering NodeList resources.
type nodeListState struct {
}

type NodeListState struct {
}

func (NodeListState) ElementType() reflect.Type {
	return reflect.TypeOf((*nodeListState)(nil)).Elem()
}

type nodeListArgs struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion *string `pulumi:"apiVersion"`
	// List of nodes
	Items []NodeType `pulumi:"items"`
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind *string `pulumi:"kind"`
	// Standard list metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Metadata *metav1.ListMeta `pulumi:"metadata"`
}

// The set of arguments for constructing a NodeList resource.
type NodeListArgs struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion pulumi.StringPtrInput
	// List of nodes
	Items NodeTypeArrayInput
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind pulumi.StringPtrInput
	// Standard list metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Metadata metav1.ListMetaPtrInput
}

func (NodeListArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*nodeListArgs)(nil)).Elem()
}

type NodeListInput interface {
	pulumi.Input

	ToNodeListOutput() NodeListOutput
	ToNodeListOutputWithContext(ctx context.Context) NodeListOutput
}

func (*NodeList) ElementType() reflect.Type {
	return reflect.TypeOf((**NodeList)(nil)).Elem()
}

func (i *NodeList) ToNodeListOutput() NodeListOutput {
	return i.ToNodeListOutputWithContext(context.Background())
}

func (i *NodeList) ToNodeListOutputWithContext(ctx context.Context) NodeListOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NodeListOutput)
}

// NodeListArrayInput is an input type that accepts NodeListArray and NodeListArrayOutput values.
// You can construct a concrete instance of `NodeListArrayInput` via:
//
//          NodeListArray{ NodeListArgs{...} }
type NodeListArrayInput interface {
	pulumi.Input

	ToNodeListArrayOutput() NodeListArrayOutput
	ToNodeListArrayOutputWithContext(context.Context) NodeListArrayOutput
}

type NodeListArray []NodeListInput

func (NodeListArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*NodeList)(nil)).Elem()
}

func (i NodeListArray) ToNodeListArrayOutput() NodeListArrayOutput {
	return i.ToNodeListArrayOutputWithContext(context.Background())
}

func (i NodeListArray) ToNodeListArrayOutputWithContext(ctx context.Context) NodeListArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NodeListArrayOutput)
}

// NodeListMapInput is an input type that accepts NodeListMap and NodeListMapOutput values.
// You can construct a concrete instance of `NodeListMapInput` via:
//
//          NodeListMap{ "key": NodeListArgs{...} }
type NodeListMapInput interface {
	pulumi.Input

	ToNodeListMapOutput() NodeListMapOutput
	ToNodeListMapOutputWithContext(context.Context) NodeListMapOutput
}

type NodeListMap map[string]NodeListInput

func (NodeListMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*NodeList)(nil)).Elem()
}

func (i NodeListMap) ToNodeListMapOutput() NodeListMapOutput {
	return i.ToNodeListMapOutputWithContext(context.Background())
}

func (i NodeListMap) ToNodeListMapOutputWithContext(ctx context.Context) NodeListMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NodeListMapOutput)
}

type NodeListOutput struct{ *pulumi.OutputState }

func (NodeListOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**NodeList)(nil)).Elem()
}

func (o NodeListOutput) ToNodeListOutput() NodeListOutput {
	return o
}

func (o NodeListOutput) ToNodeListOutputWithContext(ctx context.Context) NodeListOutput {
	return o
}

type NodeListArrayOutput struct{ *pulumi.OutputState }

func (NodeListArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*NodeList)(nil)).Elem()
}

func (o NodeListArrayOutput) ToNodeListArrayOutput() NodeListArrayOutput {
	return o
}

func (o NodeListArrayOutput) ToNodeListArrayOutputWithContext(ctx context.Context) NodeListArrayOutput {
	return o
}

func (o NodeListArrayOutput) Index(i pulumi.IntInput) NodeListOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *NodeList {
		return vs[0].([]*NodeList)[vs[1].(int)]
	}).(NodeListOutput)
}

type NodeListMapOutput struct{ *pulumi.OutputState }

func (NodeListMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*NodeList)(nil)).Elem()
}

func (o NodeListMapOutput) ToNodeListMapOutput() NodeListMapOutput {
	return o
}

func (o NodeListMapOutput) ToNodeListMapOutputWithContext(ctx context.Context) NodeListMapOutput {
	return o
}

func (o NodeListMapOutput) MapIndex(k pulumi.StringInput) NodeListOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *NodeList {
		return vs[0].(map[string]*NodeList)[vs[1].(string)]
	}).(NodeListOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*NodeListInput)(nil)).Elem(), &NodeList{})
	pulumi.RegisterInputType(reflect.TypeOf((*NodeListArrayInput)(nil)).Elem(), NodeListArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*NodeListMapInput)(nil)).Elem(), NodeListMap{})
	pulumi.RegisterOutputType(NodeListOutput{})
	pulumi.RegisterOutputType(NodeListArrayOutput{})
	pulumi.RegisterOutputType(NodeListMapOutput{})
}
