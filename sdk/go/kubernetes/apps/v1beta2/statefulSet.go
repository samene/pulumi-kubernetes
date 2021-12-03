// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v1beta2

import (
	"context"
	"reflect"

	metav1 "github.com/pulumi/pulumi-kubernetes/sdk/v3/go/kubernetes/meta/v1"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// StatefulSet represents a set of pods with consistent identities. Identities are defined as:
//  - Network: A single stable DNS and hostname.
//  - Storage: As many VolumeClaims as requested.
//    The StatefulSet guarantees that a given network identity will always map to the same storage identity.
//
// This resource waits until its status is ready before registering success
// for create/update, and populating output properties from the current state of the resource.
// The following conditions are used to determine whether the resource creation has
// succeeded or failed:
//
// 1. The value of 'spec.replicas' matches '.status.replicas', '.status.currentReplicas',
//    and '.status.readyReplicas'.
// 2. The value of '.status.updateRevision' matches '.status.currentRevision'.
//
// If the StatefulSet has not reached a Ready state after 10 minutes, it will
// time out and mark the resource update as Failed. You can override the default timeout value
// by setting the 'customTimeouts' option on the resource.
//
// Deprecated: apps/v1beta2/StatefulSet is deprecated by apps/v1/StatefulSet and not supported by Kubernetes v1.16+ clusters.
type StatefulSet struct {
	pulumi.CustomResourceState

	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion pulumi.StringPtrOutput `pulumi:"apiVersion"`
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind     pulumi.StringPtrOutput     `pulumi:"kind"`
	Metadata metav1.ObjectMetaPtrOutput `pulumi:"metadata"`
	// Spec defines the desired identities of pods in this set.
	Spec StatefulSetSpecPtrOutput `pulumi:"spec"`
	// Status is the current status of Pods in this StatefulSet. This data may be out of date by some window of time.
	Status StatefulSetStatusPtrOutput `pulumi:"status"`
}

// NewStatefulSet registers a new resource with the given unique name, arguments, and options.
func NewStatefulSet(ctx *pulumi.Context,
	name string, args *StatefulSetArgs, opts ...pulumi.ResourceOption) (*StatefulSet, error) {
	if args == nil {
		args = &StatefulSetArgs{}
	}

	args.ApiVersion = pulumi.StringPtr("apps/v1beta2")
	args.Kind = pulumi.StringPtr("StatefulSet")
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("kubernetes:apps/v1:StatefulSet"),
		},
		{
			Type: pulumi.String("kubernetes:apps/v1beta1:StatefulSet"),
		},
	})
	opts = append(opts, aliases)
	var resource StatefulSet
	err := ctx.RegisterResource("kubernetes:apps/v1beta2:StatefulSet", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetStatefulSet gets an existing StatefulSet resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetStatefulSet(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *StatefulSetState, opts ...pulumi.ResourceOption) (*StatefulSet, error) {
	var resource StatefulSet
	err := ctx.ReadResource("kubernetes:apps/v1beta2:StatefulSet", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering StatefulSet resources.
type statefulSetState struct {
}

type StatefulSetState struct {
}

func (StatefulSetState) ElementType() reflect.Type {
	return reflect.TypeOf((*statefulSetState)(nil)).Elem()
}

type statefulSetArgs struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion *string `pulumi:"apiVersion"`
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind     *string            `pulumi:"kind"`
	Metadata *metav1.ObjectMeta `pulumi:"metadata"`
	// Spec defines the desired identities of pods in this set.
	Spec *StatefulSetSpec `pulumi:"spec"`
}

// The set of arguments for constructing a StatefulSet resource.
type StatefulSetArgs struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion pulumi.StringPtrInput
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind     pulumi.StringPtrInput
	Metadata metav1.ObjectMetaPtrInput
	// Spec defines the desired identities of pods in this set.
	Spec StatefulSetSpecPtrInput
}

func (StatefulSetArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*statefulSetArgs)(nil)).Elem()
}

type StatefulSetInput interface {
	pulumi.Input

	ToStatefulSetOutput() StatefulSetOutput
	ToStatefulSetOutputWithContext(ctx context.Context) StatefulSetOutput
}

func (*StatefulSet) ElementType() reflect.Type {
	return reflect.TypeOf((**StatefulSet)(nil)).Elem()
}

func (i *StatefulSet) ToStatefulSetOutput() StatefulSetOutput {
	return i.ToStatefulSetOutputWithContext(context.Background())
}

func (i *StatefulSet) ToStatefulSetOutputWithContext(ctx context.Context) StatefulSetOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StatefulSetOutput)
}

// StatefulSetArrayInput is an input type that accepts StatefulSetArray and StatefulSetArrayOutput values.
// You can construct a concrete instance of `StatefulSetArrayInput` via:
//
//          StatefulSetArray{ StatefulSetArgs{...} }
type StatefulSetArrayInput interface {
	pulumi.Input

	ToStatefulSetArrayOutput() StatefulSetArrayOutput
	ToStatefulSetArrayOutputWithContext(context.Context) StatefulSetArrayOutput
}

type StatefulSetArray []StatefulSetInput

func (StatefulSetArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*StatefulSet)(nil)).Elem()
}

func (i StatefulSetArray) ToStatefulSetArrayOutput() StatefulSetArrayOutput {
	return i.ToStatefulSetArrayOutputWithContext(context.Background())
}

func (i StatefulSetArray) ToStatefulSetArrayOutputWithContext(ctx context.Context) StatefulSetArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StatefulSetArrayOutput)
}

// StatefulSetMapInput is an input type that accepts StatefulSetMap and StatefulSetMapOutput values.
// You can construct a concrete instance of `StatefulSetMapInput` via:
//
//          StatefulSetMap{ "key": StatefulSetArgs{...} }
type StatefulSetMapInput interface {
	pulumi.Input

	ToStatefulSetMapOutput() StatefulSetMapOutput
	ToStatefulSetMapOutputWithContext(context.Context) StatefulSetMapOutput
}

type StatefulSetMap map[string]StatefulSetInput

func (StatefulSetMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*StatefulSet)(nil)).Elem()
}

func (i StatefulSetMap) ToStatefulSetMapOutput() StatefulSetMapOutput {
	return i.ToStatefulSetMapOutputWithContext(context.Background())
}

func (i StatefulSetMap) ToStatefulSetMapOutputWithContext(ctx context.Context) StatefulSetMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StatefulSetMapOutput)
}

type StatefulSetOutput struct{ *pulumi.OutputState }

func (StatefulSetOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**StatefulSet)(nil)).Elem()
}

func (o StatefulSetOutput) ToStatefulSetOutput() StatefulSetOutput {
	return o
}

func (o StatefulSetOutput) ToStatefulSetOutputWithContext(ctx context.Context) StatefulSetOutput {
	return o
}

type StatefulSetArrayOutput struct{ *pulumi.OutputState }

func (StatefulSetArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*StatefulSet)(nil)).Elem()
}

func (o StatefulSetArrayOutput) ToStatefulSetArrayOutput() StatefulSetArrayOutput {
	return o
}

func (o StatefulSetArrayOutput) ToStatefulSetArrayOutputWithContext(ctx context.Context) StatefulSetArrayOutput {
	return o
}

func (o StatefulSetArrayOutput) Index(i pulumi.IntInput) StatefulSetOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *StatefulSet {
		return vs[0].([]*StatefulSet)[vs[1].(int)]
	}).(StatefulSetOutput)
}

type StatefulSetMapOutput struct{ *pulumi.OutputState }

func (StatefulSetMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*StatefulSet)(nil)).Elem()
}

func (o StatefulSetMapOutput) ToStatefulSetMapOutput() StatefulSetMapOutput {
	return o
}

func (o StatefulSetMapOutput) ToStatefulSetMapOutputWithContext(ctx context.Context) StatefulSetMapOutput {
	return o
}

func (o StatefulSetMapOutput) MapIndex(k pulumi.StringInput) StatefulSetOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *StatefulSet {
		return vs[0].(map[string]*StatefulSet)[vs[1].(string)]
	}).(StatefulSetOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*StatefulSetInput)(nil)).Elem(), &StatefulSet{})
	pulumi.RegisterInputType(reflect.TypeOf((*StatefulSetArrayInput)(nil)).Elem(), StatefulSetArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*StatefulSetMapInput)(nil)).Elem(), StatefulSetMap{})
	pulumi.RegisterOutputType(StatefulSetOutput{})
	pulumi.RegisterOutputType(StatefulSetArrayOutput{})
	pulumi.RegisterOutputType(StatefulSetMapOutput{})
}
