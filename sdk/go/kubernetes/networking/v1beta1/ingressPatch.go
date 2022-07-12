// Code generated by pulumigen DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v1beta1

import (
	"context"
	"reflect"

	metav1 "github.com/pulumi/pulumi-kubernetes/sdk/v3/go/kubernetes/meta/v1"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Ingress is a collection of rules that allow inbound connections to reach the endpoints defined by a backend. An Ingress can be configured to give services externally-reachable urls, load balance traffic, terminate SSL, offer name based virtual hosting etc.
//
// This resource waits until its status is ready before registering success
// for create/update, and populating output properties from the current state of the resource.
// The following conditions are used to determine whether the resource creation has
// succeeded or failed:
//
// 1.  Ingress object exists.
// 2.  Endpoint objects exist with matching names for each Ingress path (except when Service
//     type is ExternalName).
// 3.  Ingress entry exists for '.status.loadBalancer.ingress'.
//
// If the Ingress has not reached a Ready state after 10 minutes, it will
// time out and mark the resource update as Failed. You can override the default timeout value
// by setting the 'customTimeouts' option on the resource.
type IngressPatch struct {
	pulumi.CustomResourceState

	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion pulumi.StringPtrOutput `pulumi:"apiVersion"`
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind pulumi.StringPtrOutput `pulumi:"kind"`
	// Standard object's metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	Metadata metav1.ObjectMetaPatchPtrOutput `pulumi:"metadata"`
	// Spec is the desired state of the Ingress. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#spec-and-status
	Spec IngressSpecPatchPtrOutput `pulumi:"spec"`
	// Status is the current state of the Ingress. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#spec-and-status
	Status IngressStatusPatchPtrOutput `pulumi:"status"`
}

// NewIngressPatch registers a new resource with the given unique name, arguments, and options.
func NewIngressPatch(ctx *pulumi.Context,
	name string, args *IngressPatchArgs, opts ...pulumi.ResourceOption) (*IngressPatch, error) {
	if args == nil {
		args = &IngressPatchArgs{}
	}

	args.ApiVersion = pulumi.StringPtr("networking.k8s.io/v1beta1")
	args.Kind = pulumi.StringPtr("Ingress")
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("kubernetes:extensions/v1beta1:IngressPatch"),
		},
		{
			Type: pulumi.String("kubernetes:networking.k8s.io/v1:IngressPatch"),
		},
	})
	opts = append(opts, aliases)
	var resource IngressPatch
	err := ctx.RegisterResource("kubernetes:networking.k8s.io/v1beta1:IngressPatch", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetIngressPatch gets an existing IngressPatch resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetIngressPatch(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *IngressPatchState, opts ...pulumi.ResourceOption) (*IngressPatch, error) {
	var resource IngressPatch
	err := ctx.ReadResource("kubernetes:networking.k8s.io/v1beta1:IngressPatch", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering IngressPatch resources.
type ingressPatchState struct {
}

type IngressPatchState struct {
}

func (IngressPatchState) ElementType() reflect.Type {
	return reflect.TypeOf((*ingressPatchState)(nil)).Elem()
}

type ingressPatchArgs struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion *string `pulumi:"apiVersion"`
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind *string `pulumi:"kind"`
	// Standard object's metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	Metadata *metav1.ObjectMetaPatch `pulumi:"metadata"`
	// Spec is the desired state of the Ingress. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#spec-and-status
	Spec *IngressSpecPatch `pulumi:"spec"`
}

// The set of arguments for constructing a IngressPatch resource.
type IngressPatchArgs struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion pulumi.StringPtrInput
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind pulumi.StringPtrInput
	// Standard object's metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	Metadata metav1.ObjectMetaPatchPtrInput
	// Spec is the desired state of the Ingress. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#spec-and-status
	Spec IngressSpecPatchPtrInput
}

func (IngressPatchArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ingressPatchArgs)(nil)).Elem()
}

type IngressPatchInput interface {
	pulumi.Input

	ToIngressPatchOutput() IngressPatchOutput
	ToIngressPatchOutputWithContext(ctx context.Context) IngressPatchOutput
}

func (*IngressPatch) ElementType() reflect.Type {
	return reflect.TypeOf((**IngressPatch)(nil)).Elem()
}

func (i *IngressPatch) ToIngressPatchOutput() IngressPatchOutput {
	return i.ToIngressPatchOutputWithContext(context.Background())
}

func (i *IngressPatch) ToIngressPatchOutputWithContext(ctx context.Context) IngressPatchOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IngressPatchOutput)
}

// IngressPatchArrayInput is an input type that accepts IngressPatchArray and IngressPatchArrayOutput values.
// You can construct a concrete instance of `IngressPatchArrayInput` via:
//
//          IngressPatchArray{ IngressPatchArgs{...} }
type IngressPatchArrayInput interface {
	pulumi.Input

	ToIngressPatchArrayOutput() IngressPatchArrayOutput
	ToIngressPatchArrayOutputWithContext(context.Context) IngressPatchArrayOutput
}

type IngressPatchArray []IngressPatchInput

func (IngressPatchArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*IngressPatch)(nil)).Elem()
}

func (i IngressPatchArray) ToIngressPatchArrayOutput() IngressPatchArrayOutput {
	return i.ToIngressPatchArrayOutputWithContext(context.Background())
}

func (i IngressPatchArray) ToIngressPatchArrayOutputWithContext(ctx context.Context) IngressPatchArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IngressPatchArrayOutput)
}

// IngressPatchMapInput is an input type that accepts IngressPatchMap and IngressPatchMapOutput values.
// You can construct a concrete instance of `IngressPatchMapInput` via:
//
//          IngressPatchMap{ "key": IngressPatchArgs{...} }
type IngressPatchMapInput interface {
	pulumi.Input

	ToIngressPatchMapOutput() IngressPatchMapOutput
	ToIngressPatchMapOutputWithContext(context.Context) IngressPatchMapOutput
}

type IngressPatchMap map[string]IngressPatchInput

func (IngressPatchMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*IngressPatch)(nil)).Elem()
}

func (i IngressPatchMap) ToIngressPatchMapOutput() IngressPatchMapOutput {
	return i.ToIngressPatchMapOutputWithContext(context.Background())
}

func (i IngressPatchMap) ToIngressPatchMapOutputWithContext(ctx context.Context) IngressPatchMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IngressPatchMapOutput)
}

type IngressPatchOutput struct{ *pulumi.OutputState }

func (IngressPatchOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**IngressPatch)(nil)).Elem()
}

func (o IngressPatchOutput) ToIngressPatchOutput() IngressPatchOutput {
	return o
}

func (o IngressPatchOutput) ToIngressPatchOutputWithContext(ctx context.Context) IngressPatchOutput {
	return o
}

// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
func (o IngressPatchOutput) ApiVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *IngressPatch) pulumi.StringPtrOutput { return v.ApiVersion }).(pulumi.StringPtrOutput)
}

// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
func (o IngressPatchOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *IngressPatch) pulumi.StringPtrOutput { return v.Kind }).(pulumi.StringPtrOutput)
}

// Standard object's metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
func (o IngressPatchOutput) Metadata() metav1.ObjectMetaPatchPtrOutput {
	return o.ApplyT(func(v *IngressPatch) metav1.ObjectMetaPatchPtrOutput { return v.Metadata }).(metav1.ObjectMetaPatchPtrOutput)
}

// Spec is the desired state of the Ingress. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#spec-and-status
func (o IngressPatchOutput) Spec() IngressSpecPatchPtrOutput {
	return o.ApplyT(func(v *IngressPatch) IngressSpecPatchPtrOutput { return v.Spec }).(IngressSpecPatchPtrOutput)
}

// Status is the current state of the Ingress. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#spec-and-status
func (o IngressPatchOutput) Status() IngressStatusPatchPtrOutput {
	return o.ApplyT(func(v *IngressPatch) IngressStatusPatchPtrOutput { return v.Status }).(IngressStatusPatchPtrOutput)
}

type IngressPatchArrayOutput struct{ *pulumi.OutputState }

func (IngressPatchArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*IngressPatch)(nil)).Elem()
}

func (o IngressPatchArrayOutput) ToIngressPatchArrayOutput() IngressPatchArrayOutput {
	return o
}

func (o IngressPatchArrayOutput) ToIngressPatchArrayOutputWithContext(ctx context.Context) IngressPatchArrayOutput {
	return o
}

func (o IngressPatchArrayOutput) Index(i pulumi.IntInput) IngressPatchOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *IngressPatch {
		return vs[0].([]*IngressPatch)[vs[1].(int)]
	}).(IngressPatchOutput)
}

type IngressPatchMapOutput struct{ *pulumi.OutputState }

func (IngressPatchMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*IngressPatch)(nil)).Elem()
}

func (o IngressPatchMapOutput) ToIngressPatchMapOutput() IngressPatchMapOutput {
	return o
}

func (o IngressPatchMapOutput) ToIngressPatchMapOutputWithContext(ctx context.Context) IngressPatchMapOutput {
	return o
}

func (o IngressPatchMapOutput) MapIndex(k pulumi.StringInput) IngressPatchOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *IngressPatch {
		return vs[0].(map[string]*IngressPatch)[vs[1].(string)]
	}).(IngressPatchOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*IngressPatchInput)(nil)).Elem(), &IngressPatch{})
	pulumi.RegisterInputType(reflect.TypeOf((*IngressPatchArrayInput)(nil)).Elem(), IngressPatchArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*IngressPatchMapInput)(nil)).Elem(), IngressPatchMap{})
	pulumi.RegisterOutputType(IngressPatchOutput{})
	pulumi.RegisterOutputType(IngressPatchArrayOutput{})
	pulumi.RegisterOutputType(IngressPatchMapOutput{})
}
