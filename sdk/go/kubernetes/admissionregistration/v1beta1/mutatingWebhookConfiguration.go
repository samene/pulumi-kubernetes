// Code generated by pulumigen DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v1beta1

import (
	"context"
	"reflect"

	metav1 "github.com/pulumi/pulumi-kubernetes/sdk/v4/go/kubernetes/meta/v1"
	"github.com/pulumi/pulumi-kubernetes/sdk/v4/go/kubernetes/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// MutatingWebhookConfiguration describes the configuration of and admission webhook that accept or reject and may change the object. Deprecated in v1.16, planned for removal in v1.19. Use admissionregistration.k8s.io/v1 MutatingWebhookConfiguration instead.
type MutatingWebhookConfiguration struct {
	pulumi.CustomResourceState

	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion pulumi.StringOutput `pulumi:"apiVersion"`
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind pulumi.StringOutput `pulumi:"kind"`
	// Standard object metadata; More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata.
	Metadata metav1.ObjectMetaOutput `pulumi:"metadata"`
	// Webhooks is a list of webhooks and the affected resources and operations.
	Webhooks MutatingWebhookArrayOutput `pulumi:"webhooks"`
}

// NewMutatingWebhookConfiguration registers a new resource with the given unique name, arguments, and options.
func NewMutatingWebhookConfiguration(ctx *pulumi.Context,
	name string, args *MutatingWebhookConfigurationArgs, opts ...pulumi.ResourceOption) (*MutatingWebhookConfiguration, error) {
	if args == nil {
		args = &MutatingWebhookConfigurationArgs{}
	}

	args.ApiVersion = pulumi.StringPtr("admissionregistration.k8s.io/v1beta1")
	args.Kind = pulumi.StringPtr("MutatingWebhookConfiguration")
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("kubernetes:admissionregistration.k8s.io/v1:MutatingWebhookConfiguration"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource MutatingWebhookConfiguration
	err := ctx.RegisterResource("kubernetes:admissionregistration.k8s.io/v1beta1:MutatingWebhookConfiguration", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetMutatingWebhookConfiguration gets an existing MutatingWebhookConfiguration resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetMutatingWebhookConfiguration(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *MutatingWebhookConfigurationState, opts ...pulumi.ResourceOption) (*MutatingWebhookConfiguration, error) {
	var resource MutatingWebhookConfiguration
	err := ctx.ReadResource("kubernetes:admissionregistration.k8s.io/v1beta1:MutatingWebhookConfiguration", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering MutatingWebhookConfiguration resources.
type mutatingWebhookConfigurationState struct {
}

type MutatingWebhookConfigurationState struct {
}

func (MutatingWebhookConfigurationState) ElementType() reflect.Type {
	return reflect.TypeOf((*mutatingWebhookConfigurationState)(nil)).Elem()
}

type mutatingWebhookConfigurationArgs struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion *string `pulumi:"apiVersion"`
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind *string `pulumi:"kind"`
	// Standard object metadata; More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata.
	Metadata *metav1.ObjectMeta `pulumi:"metadata"`
	// Webhooks is a list of webhooks and the affected resources and operations.
	Webhooks []MutatingWebhook `pulumi:"webhooks"`
}

// The set of arguments for constructing a MutatingWebhookConfiguration resource.
type MutatingWebhookConfigurationArgs struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion pulumi.StringPtrInput
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind pulumi.StringPtrInput
	// Standard object metadata; More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata.
	Metadata metav1.ObjectMetaPtrInput
	// Webhooks is a list of webhooks and the affected resources and operations.
	Webhooks MutatingWebhookArrayInput
}

func (MutatingWebhookConfigurationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*mutatingWebhookConfigurationArgs)(nil)).Elem()
}

type MutatingWebhookConfigurationInput interface {
	pulumi.Input

	ToMutatingWebhookConfigurationOutput() MutatingWebhookConfigurationOutput
	ToMutatingWebhookConfigurationOutputWithContext(ctx context.Context) MutatingWebhookConfigurationOutput
}

func (*MutatingWebhookConfiguration) ElementType() reflect.Type {
	return reflect.TypeOf((**MutatingWebhookConfiguration)(nil)).Elem()
}

func (i *MutatingWebhookConfiguration) ToMutatingWebhookConfigurationOutput() MutatingWebhookConfigurationOutput {
	return i.ToMutatingWebhookConfigurationOutputWithContext(context.Background())
}

func (i *MutatingWebhookConfiguration) ToMutatingWebhookConfigurationOutputWithContext(ctx context.Context) MutatingWebhookConfigurationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MutatingWebhookConfigurationOutput)
}

// MutatingWebhookConfigurationArrayInput is an input type that accepts MutatingWebhookConfigurationArray and MutatingWebhookConfigurationArrayOutput values.
// You can construct a concrete instance of `MutatingWebhookConfigurationArrayInput` via:
//
//	MutatingWebhookConfigurationArray{ MutatingWebhookConfigurationArgs{...} }
type MutatingWebhookConfigurationArrayInput interface {
	pulumi.Input

	ToMutatingWebhookConfigurationArrayOutput() MutatingWebhookConfigurationArrayOutput
	ToMutatingWebhookConfigurationArrayOutputWithContext(context.Context) MutatingWebhookConfigurationArrayOutput
}

type MutatingWebhookConfigurationArray []MutatingWebhookConfigurationInput

func (MutatingWebhookConfigurationArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*MutatingWebhookConfiguration)(nil)).Elem()
}

func (i MutatingWebhookConfigurationArray) ToMutatingWebhookConfigurationArrayOutput() MutatingWebhookConfigurationArrayOutput {
	return i.ToMutatingWebhookConfigurationArrayOutputWithContext(context.Background())
}

func (i MutatingWebhookConfigurationArray) ToMutatingWebhookConfigurationArrayOutputWithContext(ctx context.Context) MutatingWebhookConfigurationArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MutatingWebhookConfigurationArrayOutput)
}

// MutatingWebhookConfigurationMapInput is an input type that accepts MutatingWebhookConfigurationMap and MutatingWebhookConfigurationMapOutput values.
// You can construct a concrete instance of `MutatingWebhookConfigurationMapInput` via:
//
//	MutatingWebhookConfigurationMap{ "key": MutatingWebhookConfigurationArgs{...} }
type MutatingWebhookConfigurationMapInput interface {
	pulumi.Input

	ToMutatingWebhookConfigurationMapOutput() MutatingWebhookConfigurationMapOutput
	ToMutatingWebhookConfigurationMapOutputWithContext(context.Context) MutatingWebhookConfigurationMapOutput
}

type MutatingWebhookConfigurationMap map[string]MutatingWebhookConfigurationInput

func (MutatingWebhookConfigurationMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*MutatingWebhookConfiguration)(nil)).Elem()
}

func (i MutatingWebhookConfigurationMap) ToMutatingWebhookConfigurationMapOutput() MutatingWebhookConfigurationMapOutput {
	return i.ToMutatingWebhookConfigurationMapOutputWithContext(context.Background())
}

func (i MutatingWebhookConfigurationMap) ToMutatingWebhookConfigurationMapOutputWithContext(ctx context.Context) MutatingWebhookConfigurationMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MutatingWebhookConfigurationMapOutput)
}

type MutatingWebhookConfigurationOutput struct{ *pulumi.OutputState }

func (MutatingWebhookConfigurationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**MutatingWebhookConfiguration)(nil)).Elem()
}

func (o MutatingWebhookConfigurationOutput) ToMutatingWebhookConfigurationOutput() MutatingWebhookConfigurationOutput {
	return o
}

func (o MutatingWebhookConfigurationOutput) ToMutatingWebhookConfigurationOutputWithContext(ctx context.Context) MutatingWebhookConfigurationOutput {
	return o
}

// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
func (o MutatingWebhookConfigurationOutput) ApiVersion() pulumi.StringOutput {
	return o.ApplyT(func(v *MutatingWebhookConfiguration) pulumi.StringOutput { return v.ApiVersion }).(pulumi.StringOutput)
}

// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
func (o MutatingWebhookConfigurationOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v *MutatingWebhookConfiguration) pulumi.StringOutput { return v.Kind }).(pulumi.StringOutput)
}

// Standard object metadata; More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata.
func (o MutatingWebhookConfigurationOutput) Metadata() metav1.ObjectMetaOutput {
	return o.ApplyT(func(v *MutatingWebhookConfiguration) metav1.ObjectMetaOutput { return v.Metadata }).(metav1.ObjectMetaOutput)
}

// Webhooks is a list of webhooks and the affected resources and operations.
func (o MutatingWebhookConfigurationOutput) Webhooks() MutatingWebhookArrayOutput {
	return o.ApplyT(func(v *MutatingWebhookConfiguration) MutatingWebhookArrayOutput { return v.Webhooks }).(MutatingWebhookArrayOutput)
}

type MutatingWebhookConfigurationArrayOutput struct{ *pulumi.OutputState }

func (MutatingWebhookConfigurationArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*MutatingWebhookConfiguration)(nil)).Elem()
}

func (o MutatingWebhookConfigurationArrayOutput) ToMutatingWebhookConfigurationArrayOutput() MutatingWebhookConfigurationArrayOutput {
	return o
}

func (o MutatingWebhookConfigurationArrayOutput) ToMutatingWebhookConfigurationArrayOutputWithContext(ctx context.Context) MutatingWebhookConfigurationArrayOutput {
	return o
}

func (o MutatingWebhookConfigurationArrayOutput) Index(i pulumi.IntInput) MutatingWebhookConfigurationOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *MutatingWebhookConfiguration {
		return vs[0].([]*MutatingWebhookConfiguration)[vs[1].(int)]
	}).(MutatingWebhookConfigurationOutput)
}

type MutatingWebhookConfigurationMapOutput struct{ *pulumi.OutputState }

func (MutatingWebhookConfigurationMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*MutatingWebhookConfiguration)(nil)).Elem()
}

func (o MutatingWebhookConfigurationMapOutput) ToMutatingWebhookConfigurationMapOutput() MutatingWebhookConfigurationMapOutput {
	return o
}

func (o MutatingWebhookConfigurationMapOutput) ToMutatingWebhookConfigurationMapOutputWithContext(ctx context.Context) MutatingWebhookConfigurationMapOutput {
	return o
}

func (o MutatingWebhookConfigurationMapOutput) MapIndex(k pulumi.StringInput) MutatingWebhookConfigurationOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *MutatingWebhookConfiguration {
		return vs[0].(map[string]*MutatingWebhookConfiguration)[vs[1].(string)]
	}).(MutatingWebhookConfigurationOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*MutatingWebhookConfigurationInput)(nil)).Elem(), &MutatingWebhookConfiguration{})
	pulumi.RegisterInputType(reflect.TypeOf((*MutatingWebhookConfigurationArrayInput)(nil)).Elem(), MutatingWebhookConfigurationArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*MutatingWebhookConfigurationMapInput)(nil)).Elem(), MutatingWebhookConfigurationMap{})
	pulumi.RegisterOutputType(MutatingWebhookConfigurationOutput{})
	pulumi.RegisterOutputType(MutatingWebhookConfigurationArrayOutput{})
	pulumi.RegisterOutputType(MutatingWebhookConfigurationMapOutput{})
}
