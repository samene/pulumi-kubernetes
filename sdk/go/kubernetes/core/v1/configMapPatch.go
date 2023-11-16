// Code generated by pulumigen DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v1

import (
	"context"
	"reflect"

	metav1 "github.com/pulumi/pulumi-kubernetes/sdk/v4/go/kubernetes/meta/v1"
	"github.com/pulumi/pulumi-kubernetes/sdk/v4/go/kubernetes/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Patch resources are used to modify existing Kubernetes resources by using
// Server-Side Apply updates. The name of the resource must be specified, but all other properties are optional. More than
// one patch may be applied to the same resource, and a random FieldManager name will be used for each Patch resource.
// Conflicts will result in an error by default, but can be forced using the "pulumi.com/patchForce" annotation. See the
// [Server-Side Apply Docs](https://www.pulumi.com/registry/packages/kubernetes/how-to-guides/managing-resources-with-server-side-apply/) for
// additional information about using Server-Side Apply to manage Kubernetes resources with Pulumi.
// ConfigMap holds configuration data for pods to consume.
type ConfigMapPatch struct {
	pulumi.CustomResourceState

	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion pulumi.StringPtrOutput `pulumi:"apiVersion"`
	// BinaryData contains the binary data. Each key must consist of alphanumeric characters, '-', '_' or '.'. BinaryData can contain byte sequences that are not in the UTF-8 range. The keys stored in BinaryData must not overlap with the ones in the Data field, this is enforced during validation process. Using this field will require 1.10+ apiserver and kubelet.
	BinaryData pulumi.StringMapOutput `pulumi:"binaryData"`
	// Data contains the configuration data. Each key must consist of alphanumeric characters, '-', '_' or '.'. Values with non-UTF-8 byte sequences must use the BinaryData field. The keys stored in Data must not overlap with the keys in the BinaryData field, this is enforced during validation process.
	Data pulumi.StringMapOutput `pulumi:"data"`
	// Immutable, if set to true, ensures that data stored in the ConfigMap cannot be updated (only object metadata can be modified). If not set to true, the field can be modified at any time. Defaulted to nil.
	Immutable pulumi.BoolPtrOutput `pulumi:"immutable"`
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind pulumi.StringPtrOutput `pulumi:"kind"`
	// Standard object's metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	Metadata metav1.ObjectMetaPatchPtrOutput `pulumi:"metadata"`
}

// NewConfigMapPatch registers a new resource with the given unique name, arguments, and options.
func NewConfigMapPatch(ctx *pulumi.Context,
	name string, args *ConfigMapPatchArgs, opts ...pulumi.ResourceOption) (*ConfigMapPatch, error) {
	if args == nil {
		args = &ConfigMapPatchArgs{}
	}

	args.ApiVersion = pulumi.StringPtr("v1")
	args.Kind = pulumi.StringPtr("ConfigMap")
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource ConfigMapPatch
	err := ctx.RegisterResource("kubernetes:core/v1:ConfigMapPatch", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetConfigMapPatch gets an existing ConfigMapPatch resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetConfigMapPatch(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ConfigMapPatchState, opts ...pulumi.ResourceOption) (*ConfigMapPatch, error) {
	var resource ConfigMapPatch
	err := ctx.ReadResource("kubernetes:core/v1:ConfigMapPatch", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ConfigMapPatch resources.
type configMapPatchState struct {
}

type ConfigMapPatchState struct {
}

func (ConfigMapPatchState) ElementType() reflect.Type {
	return reflect.TypeOf((*configMapPatchState)(nil)).Elem()
}

type configMapPatchArgs struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion *string `pulumi:"apiVersion"`
	// BinaryData contains the binary data. Each key must consist of alphanumeric characters, '-', '_' or '.'. BinaryData can contain byte sequences that are not in the UTF-8 range. The keys stored in BinaryData must not overlap with the ones in the Data field, this is enforced during validation process. Using this field will require 1.10+ apiserver and kubelet.
	BinaryData map[string]string `pulumi:"binaryData"`
	// Data contains the configuration data. Each key must consist of alphanumeric characters, '-', '_' or '.'. Values with non-UTF-8 byte sequences must use the BinaryData field. The keys stored in Data must not overlap with the keys in the BinaryData field, this is enforced during validation process.
	Data map[string]string `pulumi:"data"`
	// Immutable, if set to true, ensures that data stored in the ConfigMap cannot be updated (only object metadata can be modified). If not set to true, the field can be modified at any time. Defaulted to nil.
	Immutable *bool `pulumi:"immutable"`
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind *string `pulumi:"kind"`
	// Standard object's metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	Metadata *metav1.ObjectMetaPatch `pulumi:"metadata"`
}

// The set of arguments for constructing a ConfigMapPatch resource.
type ConfigMapPatchArgs struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion pulumi.StringPtrInput
	// BinaryData contains the binary data. Each key must consist of alphanumeric characters, '-', '_' or '.'. BinaryData can contain byte sequences that are not in the UTF-8 range. The keys stored in BinaryData must not overlap with the ones in the Data field, this is enforced during validation process. Using this field will require 1.10+ apiserver and kubelet.
	BinaryData pulumi.StringMapInput
	// Data contains the configuration data. Each key must consist of alphanumeric characters, '-', '_' or '.'. Values with non-UTF-8 byte sequences must use the BinaryData field. The keys stored in Data must not overlap with the keys in the BinaryData field, this is enforced during validation process.
	Data pulumi.StringMapInput
	// Immutable, if set to true, ensures that data stored in the ConfigMap cannot be updated (only object metadata can be modified). If not set to true, the field can be modified at any time. Defaulted to nil.
	Immutable pulumi.BoolPtrInput
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind pulumi.StringPtrInput
	// Standard object's metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	Metadata metav1.ObjectMetaPatchPtrInput
}

func (ConfigMapPatchArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*configMapPatchArgs)(nil)).Elem()
}

type ConfigMapPatchInput interface {
	pulumi.Input

	ToConfigMapPatchOutput() ConfigMapPatchOutput
	ToConfigMapPatchOutputWithContext(ctx context.Context) ConfigMapPatchOutput
}

func (*ConfigMapPatch) ElementType() reflect.Type {
	return reflect.TypeOf((**ConfigMapPatch)(nil)).Elem()
}

func (i *ConfigMapPatch) ToConfigMapPatchOutput() ConfigMapPatchOutput {
	return i.ToConfigMapPatchOutputWithContext(context.Background())
}

func (i *ConfigMapPatch) ToConfigMapPatchOutputWithContext(ctx context.Context) ConfigMapPatchOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConfigMapPatchOutput)
}

// ConfigMapPatchArrayInput is an input type that accepts ConfigMapPatchArray and ConfigMapPatchArrayOutput values.
// You can construct a concrete instance of `ConfigMapPatchArrayInput` via:
//
//	ConfigMapPatchArray{ ConfigMapPatchArgs{...} }
type ConfigMapPatchArrayInput interface {
	pulumi.Input

	ToConfigMapPatchArrayOutput() ConfigMapPatchArrayOutput
	ToConfigMapPatchArrayOutputWithContext(context.Context) ConfigMapPatchArrayOutput
}

type ConfigMapPatchArray []ConfigMapPatchInput

func (ConfigMapPatchArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ConfigMapPatch)(nil)).Elem()
}

func (i ConfigMapPatchArray) ToConfigMapPatchArrayOutput() ConfigMapPatchArrayOutput {
	return i.ToConfigMapPatchArrayOutputWithContext(context.Background())
}

func (i ConfigMapPatchArray) ToConfigMapPatchArrayOutputWithContext(ctx context.Context) ConfigMapPatchArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConfigMapPatchArrayOutput)
}

// ConfigMapPatchMapInput is an input type that accepts ConfigMapPatchMap and ConfigMapPatchMapOutput values.
// You can construct a concrete instance of `ConfigMapPatchMapInput` via:
//
//	ConfigMapPatchMap{ "key": ConfigMapPatchArgs{...} }
type ConfigMapPatchMapInput interface {
	pulumi.Input

	ToConfigMapPatchMapOutput() ConfigMapPatchMapOutput
	ToConfigMapPatchMapOutputWithContext(context.Context) ConfigMapPatchMapOutput
}

type ConfigMapPatchMap map[string]ConfigMapPatchInput

func (ConfigMapPatchMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ConfigMapPatch)(nil)).Elem()
}

func (i ConfigMapPatchMap) ToConfigMapPatchMapOutput() ConfigMapPatchMapOutput {
	return i.ToConfigMapPatchMapOutputWithContext(context.Background())
}

func (i ConfigMapPatchMap) ToConfigMapPatchMapOutputWithContext(ctx context.Context) ConfigMapPatchMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConfigMapPatchMapOutput)
}

type ConfigMapPatchOutput struct{ *pulumi.OutputState }

func (ConfigMapPatchOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ConfigMapPatch)(nil)).Elem()
}

func (o ConfigMapPatchOutput) ToConfigMapPatchOutput() ConfigMapPatchOutput {
	return o
}

func (o ConfigMapPatchOutput) ToConfigMapPatchOutputWithContext(ctx context.Context) ConfigMapPatchOutput {
	return o
}

// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
func (o ConfigMapPatchOutput) ApiVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ConfigMapPatch) pulumi.StringPtrOutput { return v.ApiVersion }).(pulumi.StringPtrOutput)
}

// BinaryData contains the binary data. Each key must consist of alphanumeric characters, '-', '_' or '.'. BinaryData can contain byte sequences that are not in the UTF-8 range. The keys stored in BinaryData must not overlap with the ones in the Data field, this is enforced during validation process. Using this field will require 1.10+ apiserver and kubelet.
func (o ConfigMapPatchOutput) BinaryData() pulumi.StringMapOutput {
	return o.ApplyT(func(v *ConfigMapPatch) pulumi.StringMapOutput { return v.BinaryData }).(pulumi.StringMapOutput)
}

// Data contains the configuration data. Each key must consist of alphanumeric characters, '-', '_' or '.'. Values with non-UTF-8 byte sequences must use the BinaryData field. The keys stored in Data must not overlap with the keys in the BinaryData field, this is enforced during validation process.
func (o ConfigMapPatchOutput) Data() pulumi.StringMapOutput {
	return o.ApplyT(func(v *ConfigMapPatch) pulumi.StringMapOutput { return v.Data }).(pulumi.StringMapOutput)
}

// Immutable, if set to true, ensures that data stored in the ConfigMap cannot be updated (only object metadata can be modified). If not set to true, the field can be modified at any time. Defaulted to nil.
func (o ConfigMapPatchOutput) Immutable() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *ConfigMapPatch) pulumi.BoolPtrOutput { return v.Immutable }).(pulumi.BoolPtrOutput)
}

// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
func (o ConfigMapPatchOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ConfigMapPatch) pulumi.StringPtrOutput { return v.Kind }).(pulumi.StringPtrOutput)
}

// Standard object's metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
func (o ConfigMapPatchOutput) Metadata() metav1.ObjectMetaPatchPtrOutput {
	return o.ApplyT(func(v *ConfigMapPatch) metav1.ObjectMetaPatchPtrOutput { return v.Metadata }).(metav1.ObjectMetaPatchPtrOutput)
}

type ConfigMapPatchArrayOutput struct{ *pulumi.OutputState }

func (ConfigMapPatchArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ConfigMapPatch)(nil)).Elem()
}

func (o ConfigMapPatchArrayOutput) ToConfigMapPatchArrayOutput() ConfigMapPatchArrayOutput {
	return o
}

func (o ConfigMapPatchArrayOutput) ToConfigMapPatchArrayOutputWithContext(ctx context.Context) ConfigMapPatchArrayOutput {
	return o
}

func (o ConfigMapPatchArrayOutput) Index(i pulumi.IntInput) ConfigMapPatchOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *ConfigMapPatch {
		return vs[0].([]*ConfigMapPatch)[vs[1].(int)]
	}).(ConfigMapPatchOutput)
}

type ConfigMapPatchMapOutput struct{ *pulumi.OutputState }

func (ConfigMapPatchMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ConfigMapPatch)(nil)).Elem()
}

func (o ConfigMapPatchMapOutput) ToConfigMapPatchMapOutput() ConfigMapPatchMapOutput {
	return o
}

func (o ConfigMapPatchMapOutput) ToConfigMapPatchMapOutputWithContext(ctx context.Context) ConfigMapPatchMapOutput {
	return o
}

func (o ConfigMapPatchMapOutput) MapIndex(k pulumi.StringInput) ConfigMapPatchOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *ConfigMapPatch {
		return vs[0].(map[string]*ConfigMapPatch)[vs[1].(string)]
	}).(ConfigMapPatchOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ConfigMapPatchInput)(nil)).Elem(), &ConfigMapPatch{})
	pulumi.RegisterInputType(reflect.TypeOf((*ConfigMapPatchArrayInput)(nil)).Elem(), ConfigMapPatchArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ConfigMapPatchMapInput)(nil)).Elem(), ConfigMapPatchMap{})
	pulumi.RegisterOutputType(ConfigMapPatchOutput{})
	pulumi.RegisterOutputType(ConfigMapPatchArrayOutput{})
	pulumi.RegisterOutputType(ConfigMapPatchMapOutput{})
}
