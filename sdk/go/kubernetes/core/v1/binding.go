// Code generated by pulumigen DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v1

import (
	"context"
	"reflect"

	"errors"
	metav1 "github.com/pulumi/pulumi-kubernetes/sdk/v3/go/kubernetes/meta/v1"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Binding ties one object to another; for example, a pod is bound to a node by a scheduler. Deprecated in 1.7, please use the bindings subresource of pods instead.
type Binding struct {
	pulumi.CustomResourceState

	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion pulumi.StringOutput `pulumi:"apiVersion"`
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind pulumi.StringOutput `pulumi:"kind"`
	// Standard object's metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	Metadata metav1.ObjectMetaOutput `pulumi:"metadata"`
	// The target object that you want to bind to the standard object.
	Target ObjectReferenceOutput `pulumi:"target"`
}

// NewBinding registers a new resource with the given unique name, arguments, and options.
func NewBinding(ctx *pulumi.Context,
	name string, args *BindingArgs, opts ...pulumi.ResourceOption) (*Binding, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Target == nil {
		return nil, errors.New("invalid value for required argument 'Target'")
	}
	args.ApiVersion = pulumi.StringPtr("v1")
	args.Kind = pulumi.StringPtr("Binding")
	var resource Binding
	err := ctx.RegisterResource("kubernetes:core/v1:Binding", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetBinding gets an existing Binding resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetBinding(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *BindingState, opts ...pulumi.ResourceOption) (*Binding, error) {
	var resource Binding
	err := ctx.ReadResource("kubernetes:core/v1:Binding", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Binding resources.
type bindingState struct {
}

type BindingState struct {
}

func (BindingState) ElementType() reflect.Type {
	return reflect.TypeOf((*bindingState)(nil)).Elem()
}

type bindingArgs struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion *string `pulumi:"apiVersion"`
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind *string `pulumi:"kind"`
	// Standard object's metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	Metadata *metav1.ObjectMeta `pulumi:"metadata"`
	// The target object that you want to bind to the standard object.
	Target ObjectReference `pulumi:"target"`
}

// The set of arguments for constructing a Binding resource.
type BindingArgs struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion pulumi.StringPtrInput
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind pulumi.StringPtrInput
	// Standard object's metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	Metadata metav1.ObjectMetaPtrInput
	// The target object that you want to bind to the standard object.
	Target ObjectReferenceInput
}

func (BindingArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*bindingArgs)(nil)).Elem()
}

type BindingInput interface {
	pulumi.Input

	ToBindingOutput() BindingOutput
	ToBindingOutputWithContext(ctx context.Context) BindingOutput
}

func (*Binding) ElementType() reflect.Type {
	return reflect.TypeOf((**Binding)(nil)).Elem()
}

func (i *Binding) ToBindingOutput() BindingOutput {
	return i.ToBindingOutputWithContext(context.Background())
}

func (i *Binding) ToBindingOutputWithContext(ctx context.Context) BindingOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BindingOutput)
}

// BindingArrayInput is an input type that accepts BindingArray and BindingArrayOutput values.
// You can construct a concrete instance of `BindingArrayInput` via:
//
//	BindingArray{ BindingArgs{...} }
type BindingArrayInput interface {
	pulumi.Input

	ToBindingArrayOutput() BindingArrayOutput
	ToBindingArrayOutputWithContext(context.Context) BindingArrayOutput
}

type BindingArray []BindingInput

func (BindingArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Binding)(nil)).Elem()
}

func (i BindingArray) ToBindingArrayOutput() BindingArrayOutput {
	return i.ToBindingArrayOutputWithContext(context.Background())
}

func (i BindingArray) ToBindingArrayOutputWithContext(ctx context.Context) BindingArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BindingArrayOutput)
}

// BindingMapInput is an input type that accepts BindingMap and BindingMapOutput values.
// You can construct a concrete instance of `BindingMapInput` via:
//
//	BindingMap{ "key": BindingArgs{...} }
type BindingMapInput interface {
	pulumi.Input

	ToBindingMapOutput() BindingMapOutput
	ToBindingMapOutputWithContext(context.Context) BindingMapOutput
}

type BindingMap map[string]BindingInput

func (BindingMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Binding)(nil)).Elem()
}

func (i BindingMap) ToBindingMapOutput() BindingMapOutput {
	return i.ToBindingMapOutputWithContext(context.Background())
}

func (i BindingMap) ToBindingMapOutputWithContext(ctx context.Context) BindingMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BindingMapOutput)
}

type BindingOutput struct{ *pulumi.OutputState }

func (BindingOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Binding)(nil)).Elem()
}

func (o BindingOutput) ToBindingOutput() BindingOutput {
	return o
}

func (o BindingOutput) ToBindingOutputWithContext(ctx context.Context) BindingOutput {
	return o
}

// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
func (o BindingOutput) ApiVersion() pulumi.StringOutput {
	return o.ApplyT(func(v *Binding) pulumi.StringOutput { return v.ApiVersion }).(pulumi.StringOutput)
}

// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
func (o BindingOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v *Binding) pulumi.StringOutput { return v.Kind }).(pulumi.StringOutput)
}

// Standard object's metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
func (o BindingOutput) Metadata() metav1.ObjectMetaOutput {
	return o.ApplyT(func(v *Binding) metav1.ObjectMetaOutput { return v.Metadata }).(metav1.ObjectMetaOutput)
}

// The target object that you want to bind to the standard object.
func (o BindingOutput) Target() ObjectReferenceOutput {
	return o.ApplyT(func(v *Binding) ObjectReferenceOutput { return v.Target }).(ObjectReferenceOutput)
}

type BindingArrayOutput struct{ *pulumi.OutputState }

func (BindingArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Binding)(nil)).Elem()
}

func (o BindingArrayOutput) ToBindingArrayOutput() BindingArrayOutput {
	return o
}

func (o BindingArrayOutput) ToBindingArrayOutputWithContext(ctx context.Context) BindingArrayOutput {
	return o
}

func (o BindingArrayOutput) Index(i pulumi.IntInput) BindingOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Binding {
		return vs[0].([]*Binding)[vs[1].(int)]
	}).(BindingOutput)
}

type BindingMapOutput struct{ *pulumi.OutputState }

func (BindingMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Binding)(nil)).Elem()
}

func (o BindingMapOutput) ToBindingMapOutput() BindingMapOutput {
	return o
}

func (o BindingMapOutput) ToBindingMapOutputWithContext(ctx context.Context) BindingMapOutput {
	return o
}

func (o BindingMapOutput) MapIndex(k pulumi.StringInput) BindingOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Binding {
		return vs[0].(map[string]*Binding)[vs[1].(string)]
	}).(BindingOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*BindingInput)(nil)).Elem(), &Binding{})
	pulumi.RegisterInputType(reflect.TypeOf((*BindingArrayInput)(nil)).Elem(), BindingArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*BindingMapInput)(nil)).Elem(), BindingMap{})
	pulumi.RegisterOutputType(BindingOutput{})
	pulumi.RegisterOutputType(BindingArrayOutput{})
	pulumi.RegisterOutputType(BindingMapOutput{})
}
