// Code generated by pulumigen DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v1beta1

import (
	"context"
	"reflect"

	metav1 "github.com/pulumi/pulumi-kubernetes/sdk/v3/go/kubernetes/meta/v1"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// CSINode holds information about all CSI drivers installed on a node. CSI drivers do not need to create the CSINode object directly. As long as they use the node-driver-registrar sidecar container, the kubelet will automatically populate the CSINode object for the CSI driver as part of kubelet plugin registration. CSINode has the same name as a node. If the object is missing, it means either there are no CSI Drivers available on the node, or the Kubelet version is low enough that it doesn't create this object. CSINode has an OwnerReference that points to the corresponding node object.
//
// Deprecated: storage/v1beta1/CSINode is deprecated by storage.k8s.io/v1/CSINode.
type CSINodePatch struct {
	pulumi.CustomResourceState

	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion pulumi.StringPtrOutput `pulumi:"apiVersion"`
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind pulumi.StringPtrOutput `pulumi:"kind"`
	// metadata.name must be the Kubernetes node name.
	Metadata metav1.ObjectMetaPatchPtrOutput `pulumi:"metadata"`
	// spec is the specification of CSINode
	Spec CSINodeSpecPatchPtrOutput `pulumi:"spec"`
}

// NewCSINodePatch registers a new resource with the given unique name, arguments, and options.
func NewCSINodePatch(ctx *pulumi.Context,
	name string, args *CSINodePatchArgs, opts ...pulumi.ResourceOption) (*CSINodePatch, error) {
	if args == nil {
		args = &CSINodePatchArgs{}
	}

	args.ApiVersion = pulumi.StringPtr("storage.k8s.io/v1beta1")
	args.Kind = pulumi.StringPtr("CSINode")
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("kubernetes:storage.k8s.io/v1:CSINodePatch"),
		},
	})
	opts = append(opts, aliases)
	var resource CSINodePatch
	err := ctx.RegisterResource("kubernetes:storage.k8s.io/v1beta1:CSINodePatch", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetCSINodePatch gets an existing CSINodePatch resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetCSINodePatch(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *CSINodePatchState, opts ...pulumi.ResourceOption) (*CSINodePatch, error) {
	var resource CSINodePatch
	err := ctx.ReadResource("kubernetes:storage.k8s.io/v1beta1:CSINodePatch", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering CSINodePatch resources.
type csinodePatchState struct {
}

type CSINodePatchState struct {
}

func (CSINodePatchState) ElementType() reflect.Type {
	return reflect.TypeOf((*csinodePatchState)(nil)).Elem()
}

type csinodePatchArgs struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion *string `pulumi:"apiVersion"`
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind *string `pulumi:"kind"`
	// metadata.name must be the Kubernetes node name.
	Metadata *metav1.ObjectMetaPatch `pulumi:"metadata"`
	// spec is the specification of CSINode
	Spec *CSINodeSpecPatch `pulumi:"spec"`
}

// The set of arguments for constructing a CSINodePatch resource.
type CSINodePatchArgs struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion pulumi.StringPtrInput
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind pulumi.StringPtrInput
	// metadata.name must be the Kubernetes node name.
	Metadata metav1.ObjectMetaPatchPtrInput
	// spec is the specification of CSINode
	Spec CSINodeSpecPatchPtrInput
}

func (CSINodePatchArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*csinodePatchArgs)(nil)).Elem()
}

type CSINodePatchInput interface {
	pulumi.Input

	ToCSINodePatchOutput() CSINodePatchOutput
	ToCSINodePatchOutputWithContext(ctx context.Context) CSINodePatchOutput
}

func (*CSINodePatch) ElementType() reflect.Type {
	return reflect.TypeOf((**CSINodePatch)(nil)).Elem()
}

func (i *CSINodePatch) ToCSINodePatchOutput() CSINodePatchOutput {
	return i.ToCSINodePatchOutputWithContext(context.Background())
}

func (i *CSINodePatch) ToCSINodePatchOutputWithContext(ctx context.Context) CSINodePatchOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CSINodePatchOutput)
}

// CSINodePatchArrayInput is an input type that accepts CSINodePatchArray and CSINodePatchArrayOutput values.
// You can construct a concrete instance of `CSINodePatchArrayInput` via:
//
//          CSINodePatchArray{ CSINodePatchArgs{...} }
type CSINodePatchArrayInput interface {
	pulumi.Input

	ToCSINodePatchArrayOutput() CSINodePatchArrayOutput
	ToCSINodePatchArrayOutputWithContext(context.Context) CSINodePatchArrayOutput
}

type CSINodePatchArray []CSINodePatchInput

func (CSINodePatchArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*CSINodePatch)(nil)).Elem()
}

func (i CSINodePatchArray) ToCSINodePatchArrayOutput() CSINodePatchArrayOutput {
	return i.ToCSINodePatchArrayOutputWithContext(context.Background())
}

func (i CSINodePatchArray) ToCSINodePatchArrayOutputWithContext(ctx context.Context) CSINodePatchArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CSINodePatchArrayOutput)
}

// CSINodePatchMapInput is an input type that accepts CSINodePatchMap and CSINodePatchMapOutput values.
// You can construct a concrete instance of `CSINodePatchMapInput` via:
//
//          CSINodePatchMap{ "key": CSINodePatchArgs{...} }
type CSINodePatchMapInput interface {
	pulumi.Input

	ToCSINodePatchMapOutput() CSINodePatchMapOutput
	ToCSINodePatchMapOutputWithContext(context.Context) CSINodePatchMapOutput
}

type CSINodePatchMap map[string]CSINodePatchInput

func (CSINodePatchMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*CSINodePatch)(nil)).Elem()
}

func (i CSINodePatchMap) ToCSINodePatchMapOutput() CSINodePatchMapOutput {
	return i.ToCSINodePatchMapOutputWithContext(context.Background())
}

func (i CSINodePatchMap) ToCSINodePatchMapOutputWithContext(ctx context.Context) CSINodePatchMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CSINodePatchMapOutput)
}

type CSINodePatchOutput struct{ *pulumi.OutputState }

func (CSINodePatchOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CSINodePatch)(nil)).Elem()
}

func (o CSINodePatchOutput) ToCSINodePatchOutput() CSINodePatchOutput {
	return o
}

func (o CSINodePatchOutput) ToCSINodePatchOutputWithContext(ctx context.Context) CSINodePatchOutput {
	return o
}

// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
func (o CSINodePatchOutput) ApiVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CSINodePatch) pulumi.StringPtrOutput { return v.ApiVersion }).(pulumi.StringPtrOutput)
}

// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
func (o CSINodePatchOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CSINodePatch) pulumi.StringPtrOutput { return v.Kind }).(pulumi.StringPtrOutput)
}

// metadata.name must be the Kubernetes node name.
func (o CSINodePatchOutput) Metadata() metav1.ObjectMetaPatchPtrOutput {
	return o.ApplyT(func(v *CSINodePatch) metav1.ObjectMetaPatchPtrOutput { return v.Metadata }).(metav1.ObjectMetaPatchPtrOutput)
}

// spec is the specification of CSINode
func (o CSINodePatchOutput) Spec() CSINodeSpecPatchPtrOutput {
	return o.ApplyT(func(v *CSINodePatch) CSINodeSpecPatchPtrOutput { return v.Spec }).(CSINodeSpecPatchPtrOutput)
}

type CSINodePatchArrayOutput struct{ *pulumi.OutputState }

func (CSINodePatchArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*CSINodePatch)(nil)).Elem()
}

func (o CSINodePatchArrayOutput) ToCSINodePatchArrayOutput() CSINodePatchArrayOutput {
	return o
}

func (o CSINodePatchArrayOutput) ToCSINodePatchArrayOutputWithContext(ctx context.Context) CSINodePatchArrayOutput {
	return o
}

func (o CSINodePatchArrayOutput) Index(i pulumi.IntInput) CSINodePatchOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *CSINodePatch {
		return vs[0].([]*CSINodePatch)[vs[1].(int)]
	}).(CSINodePatchOutput)
}

type CSINodePatchMapOutput struct{ *pulumi.OutputState }

func (CSINodePatchMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*CSINodePatch)(nil)).Elem()
}

func (o CSINodePatchMapOutput) ToCSINodePatchMapOutput() CSINodePatchMapOutput {
	return o
}

func (o CSINodePatchMapOutput) ToCSINodePatchMapOutputWithContext(ctx context.Context) CSINodePatchMapOutput {
	return o
}

func (o CSINodePatchMapOutput) MapIndex(k pulumi.StringInput) CSINodePatchOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *CSINodePatch {
		return vs[0].(map[string]*CSINodePatch)[vs[1].(string)]
	}).(CSINodePatchOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*CSINodePatchInput)(nil)).Elem(), &CSINodePatch{})
	pulumi.RegisterInputType(reflect.TypeOf((*CSINodePatchArrayInput)(nil)).Elem(), CSINodePatchArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*CSINodePatchMapInput)(nil)).Elem(), CSINodePatchMap{})
	pulumi.RegisterOutputType(CSINodePatchOutput{})
	pulumi.RegisterOutputType(CSINodePatchArrayOutput{})
	pulumi.RegisterOutputType(CSINodePatchMapOutput{})
}
