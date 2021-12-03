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

// APIServiceList is a list of APIService objects.
type APIServiceList struct {
	pulumi.CustomResourceState

	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion pulumi.StringPtrOutput    `pulumi:"apiVersion"`
	Items      APIServiceTypeArrayOutput `pulumi:"items"`
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind     pulumi.StringPtrOutput   `pulumi:"kind"`
	Metadata metav1.ListMetaPtrOutput `pulumi:"metadata"`
}

// NewAPIServiceList registers a new resource with the given unique name, arguments, and options.
func NewAPIServiceList(ctx *pulumi.Context,
	name string, args *APIServiceListArgs, opts ...pulumi.ResourceOption) (*APIServiceList, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Items == nil {
		return nil, errors.New("invalid value for required argument 'Items'")
	}
	args.ApiVersion = pulumi.StringPtr("apiregistration.k8s.io/v1beta1")
	args.Kind = pulumi.StringPtr("APIServiceList")
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("kubernetes:apiregistration/v1beta1:APIServiceList"),
		},
	})
	opts = append(opts, aliases)
	var resource APIServiceList
	err := ctx.RegisterResource("kubernetes:apiregistration.k8s.io/v1beta1:APIServiceList", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAPIServiceList gets an existing APIServiceList resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAPIServiceList(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *APIServiceListState, opts ...pulumi.ResourceOption) (*APIServiceList, error) {
	var resource APIServiceList
	err := ctx.ReadResource("kubernetes:apiregistration.k8s.io/v1beta1:APIServiceList", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering APIServiceList resources.
type apiserviceListState struct {
}

type APIServiceListState struct {
}

func (APIServiceListState) ElementType() reflect.Type {
	return reflect.TypeOf((*apiserviceListState)(nil)).Elem()
}

type apiserviceListArgs struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion *string          `pulumi:"apiVersion"`
	Items      []APIServiceType `pulumi:"items"`
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind     *string          `pulumi:"kind"`
	Metadata *metav1.ListMeta `pulumi:"metadata"`
}

// The set of arguments for constructing a APIServiceList resource.
type APIServiceListArgs struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion pulumi.StringPtrInput
	Items      APIServiceTypeArrayInput
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind     pulumi.StringPtrInput
	Metadata metav1.ListMetaPtrInput
}

func (APIServiceListArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*apiserviceListArgs)(nil)).Elem()
}

type APIServiceListInput interface {
	pulumi.Input

	ToAPIServiceListOutput() APIServiceListOutput
	ToAPIServiceListOutputWithContext(ctx context.Context) APIServiceListOutput
}

func (*APIServiceList) ElementType() reflect.Type {
	return reflect.TypeOf((**APIServiceList)(nil)).Elem()
}

func (i *APIServiceList) ToAPIServiceListOutput() APIServiceListOutput {
	return i.ToAPIServiceListOutputWithContext(context.Background())
}

func (i *APIServiceList) ToAPIServiceListOutputWithContext(ctx context.Context) APIServiceListOutput {
	return pulumi.ToOutputWithContext(ctx, i).(APIServiceListOutput)
}

// APIServiceListArrayInput is an input type that accepts APIServiceListArray and APIServiceListArrayOutput values.
// You can construct a concrete instance of `APIServiceListArrayInput` via:
//
//          APIServiceListArray{ APIServiceListArgs{...} }
type APIServiceListArrayInput interface {
	pulumi.Input

	ToAPIServiceListArrayOutput() APIServiceListArrayOutput
	ToAPIServiceListArrayOutputWithContext(context.Context) APIServiceListArrayOutput
}

type APIServiceListArray []APIServiceListInput

func (APIServiceListArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*APIServiceList)(nil)).Elem()
}

func (i APIServiceListArray) ToAPIServiceListArrayOutput() APIServiceListArrayOutput {
	return i.ToAPIServiceListArrayOutputWithContext(context.Background())
}

func (i APIServiceListArray) ToAPIServiceListArrayOutputWithContext(ctx context.Context) APIServiceListArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(APIServiceListArrayOutput)
}

// APIServiceListMapInput is an input type that accepts APIServiceListMap and APIServiceListMapOutput values.
// You can construct a concrete instance of `APIServiceListMapInput` via:
//
//          APIServiceListMap{ "key": APIServiceListArgs{...} }
type APIServiceListMapInput interface {
	pulumi.Input

	ToAPIServiceListMapOutput() APIServiceListMapOutput
	ToAPIServiceListMapOutputWithContext(context.Context) APIServiceListMapOutput
}

type APIServiceListMap map[string]APIServiceListInput

func (APIServiceListMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*APIServiceList)(nil)).Elem()
}

func (i APIServiceListMap) ToAPIServiceListMapOutput() APIServiceListMapOutput {
	return i.ToAPIServiceListMapOutputWithContext(context.Background())
}

func (i APIServiceListMap) ToAPIServiceListMapOutputWithContext(ctx context.Context) APIServiceListMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(APIServiceListMapOutput)
}

type APIServiceListOutput struct{ *pulumi.OutputState }

func (APIServiceListOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**APIServiceList)(nil)).Elem()
}

func (o APIServiceListOutput) ToAPIServiceListOutput() APIServiceListOutput {
	return o
}

func (o APIServiceListOutput) ToAPIServiceListOutputWithContext(ctx context.Context) APIServiceListOutput {
	return o
}

type APIServiceListArrayOutput struct{ *pulumi.OutputState }

func (APIServiceListArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*APIServiceList)(nil)).Elem()
}

func (o APIServiceListArrayOutput) ToAPIServiceListArrayOutput() APIServiceListArrayOutput {
	return o
}

func (o APIServiceListArrayOutput) ToAPIServiceListArrayOutputWithContext(ctx context.Context) APIServiceListArrayOutput {
	return o
}

func (o APIServiceListArrayOutput) Index(i pulumi.IntInput) APIServiceListOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *APIServiceList {
		return vs[0].([]*APIServiceList)[vs[1].(int)]
	}).(APIServiceListOutput)
}

type APIServiceListMapOutput struct{ *pulumi.OutputState }

func (APIServiceListMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*APIServiceList)(nil)).Elem()
}

func (o APIServiceListMapOutput) ToAPIServiceListMapOutput() APIServiceListMapOutput {
	return o
}

func (o APIServiceListMapOutput) ToAPIServiceListMapOutputWithContext(ctx context.Context) APIServiceListMapOutput {
	return o
}

func (o APIServiceListMapOutput) MapIndex(k pulumi.StringInput) APIServiceListOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *APIServiceList {
		return vs[0].(map[string]*APIServiceList)[vs[1].(string)]
	}).(APIServiceListOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*APIServiceListInput)(nil)).Elem(), &APIServiceList{})
	pulumi.RegisterInputType(reflect.TypeOf((*APIServiceListArrayInput)(nil)).Elem(), APIServiceListArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*APIServiceListMapInput)(nil)).Elem(), APIServiceListMap{})
	pulumi.RegisterOutputType(APIServiceListOutput{})
	pulumi.RegisterOutputType(APIServiceListArrayOutput{})
	pulumi.RegisterOutputType(APIServiceListMapOutput{})
}
