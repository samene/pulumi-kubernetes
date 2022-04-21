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

type CertificateSigningRequestList struct {
	pulumi.CustomResourceState

	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion pulumi.StringPtrOutput                   `pulumi:"apiVersion"`
	Items      CertificateSigningRequestTypeArrayOutput `pulumi:"items"`
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind     pulumi.StringPtrOutput   `pulumi:"kind"`
	Metadata metav1.ListMetaPtrOutput `pulumi:"metadata"`
}

// NewCertificateSigningRequestList registers a new resource with the given unique name, arguments, and options.
func NewCertificateSigningRequestList(ctx *pulumi.Context,
	name string, args *CertificateSigningRequestListArgs, opts ...pulumi.ResourceOption) (*CertificateSigningRequestList, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Items == nil {
		return nil, errors.New("invalid value for required argument 'Items'")
	}
	args.ApiVersion = pulumi.StringPtr("certificates.k8s.io/v1beta1")
	args.Kind = pulumi.StringPtr("CertificateSigningRequestList")
	var resource CertificateSigningRequestList
	err := ctx.RegisterResource("kubernetes:certificates.k8s.io/v1beta1:CertificateSigningRequestList", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetCertificateSigningRequestList gets an existing CertificateSigningRequestList resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetCertificateSigningRequestList(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *CertificateSigningRequestListState, opts ...pulumi.ResourceOption) (*CertificateSigningRequestList, error) {
	var resource CertificateSigningRequestList
	err := ctx.ReadResource("kubernetes:certificates.k8s.io/v1beta1:CertificateSigningRequestList", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering CertificateSigningRequestList resources.
type certificateSigningRequestListState struct {
}

type CertificateSigningRequestListState struct {
}

func (CertificateSigningRequestListState) ElementType() reflect.Type {
	return reflect.TypeOf((*certificateSigningRequestListState)(nil)).Elem()
}

type certificateSigningRequestListArgs struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion *string                         `pulumi:"apiVersion"`
	Items      []CertificateSigningRequestType `pulumi:"items"`
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind     *string          `pulumi:"kind"`
	Metadata *metav1.ListMeta `pulumi:"metadata"`
}

// The set of arguments for constructing a CertificateSigningRequestList resource.
type CertificateSigningRequestListArgs struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion pulumi.StringPtrInput
	Items      CertificateSigningRequestTypeArrayInput
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind     pulumi.StringPtrInput
	Metadata metav1.ListMetaPtrInput
}

func (CertificateSigningRequestListArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*certificateSigningRequestListArgs)(nil)).Elem()
}

type CertificateSigningRequestListInput interface {
	pulumi.Input

	ToCertificateSigningRequestListOutput() CertificateSigningRequestListOutput
	ToCertificateSigningRequestListOutputWithContext(ctx context.Context) CertificateSigningRequestListOutput
}

func (*CertificateSigningRequestList) ElementType() reflect.Type {
	return reflect.TypeOf((**CertificateSigningRequestList)(nil)).Elem()
}

func (i *CertificateSigningRequestList) ToCertificateSigningRequestListOutput() CertificateSigningRequestListOutput {
	return i.ToCertificateSigningRequestListOutputWithContext(context.Background())
}

func (i *CertificateSigningRequestList) ToCertificateSigningRequestListOutputWithContext(ctx context.Context) CertificateSigningRequestListOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CertificateSigningRequestListOutput)
}

// CertificateSigningRequestListArrayInput is an input type that accepts CertificateSigningRequestListArray and CertificateSigningRequestListArrayOutput values.
// You can construct a concrete instance of `CertificateSigningRequestListArrayInput` via:
//
//          CertificateSigningRequestListArray{ CertificateSigningRequestListArgs{...} }
type CertificateSigningRequestListArrayInput interface {
	pulumi.Input

	ToCertificateSigningRequestListArrayOutput() CertificateSigningRequestListArrayOutput
	ToCertificateSigningRequestListArrayOutputWithContext(context.Context) CertificateSigningRequestListArrayOutput
}

type CertificateSigningRequestListArray []CertificateSigningRequestListInput

func (CertificateSigningRequestListArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*CertificateSigningRequestList)(nil)).Elem()
}

func (i CertificateSigningRequestListArray) ToCertificateSigningRequestListArrayOutput() CertificateSigningRequestListArrayOutput {
	return i.ToCertificateSigningRequestListArrayOutputWithContext(context.Background())
}

func (i CertificateSigningRequestListArray) ToCertificateSigningRequestListArrayOutputWithContext(ctx context.Context) CertificateSigningRequestListArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CertificateSigningRequestListArrayOutput)
}

// CertificateSigningRequestListMapInput is an input type that accepts CertificateSigningRequestListMap and CertificateSigningRequestListMapOutput values.
// You can construct a concrete instance of `CertificateSigningRequestListMapInput` via:
//
//          CertificateSigningRequestListMap{ "key": CertificateSigningRequestListArgs{...} }
type CertificateSigningRequestListMapInput interface {
	pulumi.Input

	ToCertificateSigningRequestListMapOutput() CertificateSigningRequestListMapOutput
	ToCertificateSigningRequestListMapOutputWithContext(context.Context) CertificateSigningRequestListMapOutput
}

type CertificateSigningRequestListMap map[string]CertificateSigningRequestListInput

func (CertificateSigningRequestListMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*CertificateSigningRequestList)(nil)).Elem()
}

func (i CertificateSigningRequestListMap) ToCertificateSigningRequestListMapOutput() CertificateSigningRequestListMapOutput {
	return i.ToCertificateSigningRequestListMapOutputWithContext(context.Background())
}

func (i CertificateSigningRequestListMap) ToCertificateSigningRequestListMapOutputWithContext(ctx context.Context) CertificateSigningRequestListMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CertificateSigningRequestListMapOutput)
}

type CertificateSigningRequestListOutput struct{ *pulumi.OutputState }

func (CertificateSigningRequestListOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CertificateSigningRequestList)(nil)).Elem()
}

func (o CertificateSigningRequestListOutput) ToCertificateSigningRequestListOutput() CertificateSigningRequestListOutput {
	return o
}

func (o CertificateSigningRequestListOutput) ToCertificateSigningRequestListOutputWithContext(ctx context.Context) CertificateSigningRequestListOutput {
	return o
}

type CertificateSigningRequestListArrayOutput struct{ *pulumi.OutputState }

func (CertificateSigningRequestListArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*CertificateSigningRequestList)(nil)).Elem()
}

func (o CertificateSigningRequestListArrayOutput) ToCertificateSigningRequestListArrayOutput() CertificateSigningRequestListArrayOutput {
	return o
}

func (o CertificateSigningRequestListArrayOutput) ToCertificateSigningRequestListArrayOutputWithContext(ctx context.Context) CertificateSigningRequestListArrayOutput {
	return o
}

func (o CertificateSigningRequestListArrayOutput) Index(i pulumi.IntInput) CertificateSigningRequestListOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *CertificateSigningRequestList {
		return vs[0].([]*CertificateSigningRequestList)[vs[1].(int)]
	}).(CertificateSigningRequestListOutput)
}

type CertificateSigningRequestListMapOutput struct{ *pulumi.OutputState }

func (CertificateSigningRequestListMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*CertificateSigningRequestList)(nil)).Elem()
}

func (o CertificateSigningRequestListMapOutput) ToCertificateSigningRequestListMapOutput() CertificateSigningRequestListMapOutput {
	return o
}

func (o CertificateSigningRequestListMapOutput) ToCertificateSigningRequestListMapOutputWithContext(ctx context.Context) CertificateSigningRequestListMapOutput {
	return o
}

func (o CertificateSigningRequestListMapOutput) MapIndex(k pulumi.StringInput) CertificateSigningRequestListOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *CertificateSigningRequestList {
		return vs[0].(map[string]*CertificateSigningRequestList)[vs[1].(string)]
	}).(CertificateSigningRequestListOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*CertificateSigningRequestListInput)(nil)).Elem(), &CertificateSigningRequestList{})
	pulumi.RegisterInputType(reflect.TypeOf((*CertificateSigningRequestListArrayInput)(nil)).Elem(), CertificateSigningRequestListArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*CertificateSigningRequestListMapInput)(nil)).Elem(), CertificateSigningRequestListMap{})
	pulumi.RegisterOutputType(CertificateSigningRequestListOutput{})
	pulumi.RegisterOutputType(CertificateSigningRequestListArrayOutput{})
	pulumi.RegisterOutputType(CertificateSigningRequestListMapOutput{})
}
