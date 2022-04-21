// Code generated by pulumigen DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v1beta1

import (
	"context"
	"reflect"

	metav1 "github.com/pulumi/pulumi-kubernetes/sdk/v3/go/kubernetes/meta/v1"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Lease defines a lease concept.
type LeaseType struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion *string `pulumi:"apiVersion"`
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind *string `pulumi:"kind"`
	// More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	Metadata *metav1.ObjectMeta `pulumi:"metadata"`
	// Specification of the Lease. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#spec-and-status
	Spec *LeaseSpec `pulumi:"spec"`
}

// LeaseTypeInput is an input type that accepts LeaseTypeArgs and LeaseTypeOutput values.
// You can construct a concrete instance of `LeaseTypeInput` via:
//
//          LeaseTypeArgs{...}
type LeaseTypeInput interface {
	pulumi.Input

	ToLeaseTypeOutput() LeaseTypeOutput
	ToLeaseTypeOutputWithContext(context.Context) LeaseTypeOutput
}

// Lease defines a lease concept.
type LeaseTypeArgs struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion pulumi.StringPtrInput `pulumi:"apiVersion"`
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind pulumi.StringPtrInput `pulumi:"kind"`
	// More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	Metadata metav1.ObjectMetaPtrInput `pulumi:"metadata"`
	// Specification of the Lease. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#spec-and-status
	Spec LeaseSpecPtrInput `pulumi:"spec"`
}

func (LeaseTypeArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LeaseType)(nil)).Elem()
}

func (i LeaseTypeArgs) ToLeaseTypeOutput() LeaseTypeOutput {
	return i.ToLeaseTypeOutputWithContext(context.Background())
}

func (i LeaseTypeArgs) ToLeaseTypeOutputWithContext(ctx context.Context) LeaseTypeOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LeaseTypeOutput)
}

// LeaseTypeArrayInput is an input type that accepts LeaseTypeArray and LeaseTypeArrayOutput values.
// You can construct a concrete instance of `LeaseTypeArrayInput` via:
//
//          LeaseTypeArray{ LeaseTypeArgs{...} }
type LeaseTypeArrayInput interface {
	pulumi.Input

	ToLeaseTypeArrayOutput() LeaseTypeArrayOutput
	ToLeaseTypeArrayOutputWithContext(context.Context) LeaseTypeArrayOutput
}

type LeaseTypeArray []LeaseTypeInput

func (LeaseTypeArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]LeaseType)(nil)).Elem()
}

func (i LeaseTypeArray) ToLeaseTypeArrayOutput() LeaseTypeArrayOutput {
	return i.ToLeaseTypeArrayOutputWithContext(context.Background())
}

func (i LeaseTypeArray) ToLeaseTypeArrayOutputWithContext(ctx context.Context) LeaseTypeArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LeaseTypeArrayOutput)
}

// Lease defines a lease concept.
type LeaseTypeOutput struct{ *pulumi.OutputState }

func (LeaseTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LeaseType)(nil)).Elem()
}

func (o LeaseTypeOutput) ToLeaseTypeOutput() LeaseTypeOutput {
	return o
}

func (o LeaseTypeOutput) ToLeaseTypeOutputWithContext(ctx context.Context) LeaseTypeOutput {
	return o
}

// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
func (o LeaseTypeOutput) ApiVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LeaseType) *string { return v.ApiVersion }).(pulumi.StringPtrOutput)
}

// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
func (o LeaseTypeOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LeaseType) *string { return v.Kind }).(pulumi.StringPtrOutput)
}

// More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
func (o LeaseTypeOutput) Metadata() metav1.ObjectMetaPtrOutput {
	return o.ApplyT(func(v LeaseType) *metav1.ObjectMeta { return v.Metadata }).(metav1.ObjectMetaPtrOutput)
}

// Specification of the Lease. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#spec-and-status
func (o LeaseTypeOutput) Spec() LeaseSpecPtrOutput {
	return o.ApplyT(func(v LeaseType) *LeaseSpec { return v.Spec }).(LeaseSpecPtrOutput)
}

type LeaseTypeArrayOutput struct{ *pulumi.OutputState }

func (LeaseTypeArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]LeaseType)(nil)).Elem()
}

func (o LeaseTypeArrayOutput) ToLeaseTypeArrayOutput() LeaseTypeArrayOutput {
	return o
}

func (o LeaseTypeArrayOutput) ToLeaseTypeArrayOutputWithContext(ctx context.Context) LeaseTypeArrayOutput {
	return o
}

func (o LeaseTypeArrayOutput) Index(i pulumi.IntInput) LeaseTypeOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) LeaseType {
		return vs[0].([]LeaseType)[vs[1].(int)]
	}).(LeaseTypeOutput)
}

// LeaseList is a list of Lease objects.
type LeaseListType struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion *string `pulumi:"apiVersion"`
	// Items is a list of schema objects.
	Items []LeaseType `pulumi:"items"`
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind *string `pulumi:"kind"`
	// Standard list metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	Metadata *metav1.ListMeta `pulumi:"metadata"`
}

// LeaseListTypeInput is an input type that accepts LeaseListTypeArgs and LeaseListTypeOutput values.
// You can construct a concrete instance of `LeaseListTypeInput` via:
//
//          LeaseListTypeArgs{...}
type LeaseListTypeInput interface {
	pulumi.Input

	ToLeaseListTypeOutput() LeaseListTypeOutput
	ToLeaseListTypeOutputWithContext(context.Context) LeaseListTypeOutput
}

// LeaseList is a list of Lease objects.
type LeaseListTypeArgs struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion pulumi.StringPtrInput `pulumi:"apiVersion"`
	// Items is a list of schema objects.
	Items LeaseTypeArrayInput `pulumi:"items"`
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind pulumi.StringPtrInput `pulumi:"kind"`
	// Standard list metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	Metadata metav1.ListMetaPtrInput `pulumi:"metadata"`
}

func (LeaseListTypeArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LeaseListType)(nil)).Elem()
}

func (i LeaseListTypeArgs) ToLeaseListTypeOutput() LeaseListTypeOutput {
	return i.ToLeaseListTypeOutputWithContext(context.Background())
}

func (i LeaseListTypeArgs) ToLeaseListTypeOutputWithContext(ctx context.Context) LeaseListTypeOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LeaseListTypeOutput)
}

// LeaseList is a list of Lease objects.
type LeaseListTypeOutput struct{ *pulumi.OutputState }

func (LeaseListTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LeaseListType)(nil)).Elem()
}

func (o LeaseListTypeOutput) ToLeaseListTypeOutput() LeaseListTypeOutput {
	return o
}

func (o LeaseListTypeOutput) ToLeaseListTypeOutputWithContext(ctx context.Context) LeaseListTypeOutput {
	return o
}

// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
func (o LeaseListTypeOutput) ApiVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LeaseListType) *string { return v.ApiVersion }).(pulumi.StringPtrOutput)
}

// Items is a list of schema objects.
func (o LeaseListTypeOutput) Items() LeaseTypeArrayOutput {
	return o.ApplyT(func(v LeaseListType) []LeaseType { return v.Items }).(LeaseTypeArrayOutput)
}

// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
func (o LeaseListTypeOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LeaseListType) *string { return v.Kind }).(pulumi.StringPtrOutput)
}

// Standard list metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
func (o LeaseListTypeOutput) Metadata() metav1.ListMetaPtrOutput {
	return o.ApplyT(func(v LeaseListType) *metav1.ListMeta { return v.Metadata }).(metav1.ListMetaPtrOutput)
}

// LeaseSpec is a specification of a Lease.
type LeaseSpec struct {
	// acquireTime is a time when the current lease was acquired.
	AcquireTime *string `pulumi:"acquireTime"`
	// holderIdentity contains the identity of the holder of a current lease.
	HolderIdentity *string `pulumi:"holderIdentity"`
	// leaseDurationSeconds is a duration that candidates for a lease need to wait to force acquire it. This is measure against time of last observed RenewTime.
	LeaseDurationSeconds *int `pulumi:"leaseDurationSeconds"`
	// leaseTransitions is the number of transitions of a lease between holders.
	LeaseTransitions *int `pulumi:"leaseTransitions"`
	// renewTime is a time when the current holder of a lease has last updated the lease.
	RenewTime *string `pulumi:"renewTime"`
}

// LeaseSpecInput is an input type that accepts LeaseSpecArgs and LeaseSpecOutput values.
// You can construct a concrete instance of `LeaseSpecInput` via:
//
//          LeaseSpecArgs{...}
type LeaseSpecInput interface {
	pulumi.Input

	ToLeaseSpecOutput() LeaseSpecOutput
	ToLeaseSpecOutputWithContext(context.Context) LeaseSpecOutput
}

// LeaseSpec is a specification of a Lease.
type LeaseSpecArgs struct {
	// acquireTime is a time when the current lease was acquired.
	AcquireTime pulumi.StringPtrInput `pulumi:"acquireTime"`
	// holderIdentity contains the identity of the holder of a current lease.
	HolderIdentity pulumi.StringPtrInput `pulumi:"holderIdentity"`
	// leaseDurationSeconds is a duration that candidates for a lease need to wait to force acquire it. This is measure against time of last observed RenewTime.
	LeaseDurationSeconds pulumi.IntPtrInput `pulumi:"leaseDurationSeconds"`
	// leaseTransitions is the number of transitions of a lease between holders.
	LeaseTransitions pulumi.IntPtrInput `pulumi:"leaseTransitions"`
	// renewTime is a time when the current holder of a lease has last updated the lease.
	RenewTime pulumi.StringPtrInput `pulumi:"renewTime"`
}

func (LeaseSpecArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LeaseSpec)(nil)).Elem()
}

func (i LeaseSpecArgs) ToLeaseSpecOutput() LeaseSpecOutput {
	return i.ToLeaseSpecOutputWithContext(context.Background())
}

func (i LeaseSpecArgs) ToLeaseSpecOutputWithContext(ctx context.Context) LeaseSpecOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LeaseSpecOutput)
}

func (i LeaseSpecArgs) ToLeaseSpecPtrOutput() LeaseSpecPtrOutput {
	return i.ToLeaseSpecPtrOutputWithContext(context.Background())
}

func (i LeaseSpecArgs) ToLeaseSpecPtrOutputWithContext(ctx context.Context) LeaseSpecPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LeaseSpecOutput).ToLeaseSpecPtrOutputWithContext(ctx)
}

// LeaseSpecPtrInput is an input type that accepts LeaseSpecArgs, LeaseSpecPtr and LeaseSpecPtrOutput values.
// You can construct a concrete instance of `LeaseSpecPtrInput` via:
//
//          LeaseSpecArgs{...}
//
//  or:
//
//          nil
type LeaseSpecPtrInput interface {
	pulumi.Input

	ToLeaseSpecPtrOutput() LeaseSpecPtrOutput
	ToLeaseSpecPtrOutputWithContext(context.Context) LeaseSpecPtrOutput
}

type leaseSpecPtrType LeaseSpecArgs

func LeaseSpecPtr(v *LeaseSpecArgs) LeaseSpecPtrInput {
	return (*leaseSpecPtrType)(v)
}

func (*leaseSpecPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**LeaseSpec)(nil)).Elem()
}

func (i *leaseSpecPtrType) ToLeaseSpecPtrOutput() LeaseSpecPtrOutput {
	return i.ToLeaseSpecPtrOutputWithContext(context.Background())
}

func (i *leaseSpecPtrType) ToLeaseSpecPtrOutputWithContext(ctx context.Context) LeaseSpecPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LeaseSpecPtrOutput)
}

// LeaseSpec is a specification of a Lease.
type LeaseSpecOutput struct{ *pulumi.OutputState }

func (LeaseSpecOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LeaseSpec)(nil)).Elem()
}

func (o LeaseSpecOutput) ToLeaseSpecOutput() LeaseSpecOutput {
	return o
}

func (o LeaseSpecOutput) ToLeaseSpecOutputWithContext(ctx context.Context) LeaseSpecOutput {
	return o
}

func (o LeaseSpecOutput) ToLeaseSpecPtrOutput() LeaseSpecPtrOutput {
	return o.ToLeaseSpecPtrOutputWithContext(context.Background())
}

func (o LeaseSpecOutput) ToLeaseSpecPtrOutputWithContext(ctx context.Context) LeaseSpecPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v LeaseSpec) *LeaseSpec {
		return &v
	}).(LeaseSpecPtrOutput)
}

// acquireTime is a time when the current lease was acquired.
func (o LeaseSpecOutput) AcquireTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LeaseSpec) *string { return v.AcquireTime }).(pulumi.StringPtrOutput)
}

// holderIdentity contains the identity of the holder of a current lease.
func (o LeaseSpecOutput) HolderIdentity() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LeaseSpec) *string { return v.HolderIdentity }).(pulumi.StringPtrOutput)
}

// leaseDurationSeconds is a duration that candidates for a lease need to wait to force acquire it. This is measure against time of last observed RenewTime.
func (o LeaseSpecOutput) LeaseDurationSeconds() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LeaseSpec) *int { return v.LeaseDurationSeconds }).(pulumi.IntPtrOutput)
}

// leaseTransitions is the number of transitions of a lease between holders.
func (o LeaseSpecOutput) LeaseTransitions() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LeaseSpec) *int { return v.LeaseTransitions }).(pulumi.IntPtrOutput)
}

// renewTime is a time when the current holder of a lease has last updated the lease.
func (o LeaseSpecOutput) RenewTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LeaseSpec) *string { return v.RenewTime }).(pulumi.StringPtrOutput)
}

type LeaseSpecPtrOutput struct{ *pulumi.OutputState }

func (LeaseSpecPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**LeaseSpec)(nil)).Elem()
}

func (o LeaseSpecPtrOutput) ToLeaseSpecPtrOutput() LeaseSpecPtrOutput {
	return o
}

func (o LeaseSpecPtrOutput) ToLeaseSpecPtrOutputWithContext(ctx context.Context) LeaseSpecPtrOutput {
	return o
}

func (o LeaseSpecPtrOutput) Elem() LeaseSpecOutput {
	return o.ApplyT(func(v *LeaseSpec) LeaseSpec {
		if v != nil {
			return *v
		}
		var ret LeaseSpec
		return ret
	}).(LeaseSpecOutput)
}

// acquireTime is a time when the current lease was acquired.
func (o LeaseSpecPtrOutput) AcquireTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LeaseSpec) *string {
		if v == nil {
			return nil
		}
		return v.AcquireTime
	}).(pulumi.StringPtrOutput)
}

// holderIdentity contains the identity of the holder of a current lease.
func (o LeaseSpecPtrOutput) HolderIdentity() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LeaseSpec) *string {
		if v == nil {
			return nil
		}
		return v.HolderIdentity
	}).(pulumi.StringPtrOutput)
}

// leaseDurationSeconds is a duration that candidates for a lease need to wait to force acquire it. This is measure against time of last observed RenewTime.
func (o LeaseSpecPtrOutput) LeaseDurationSeconds() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *LeaseSpec) *int {
		if v == nil {
			return nil
		}
		return v.LeaseDurationSeconds
	}).(pulumi.IntPtrOutput)
}

// leaseTransitions is the number of transitions of a lease between holders.
func (o LeaseSpecPtrOutput) LeaseTransitions() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *LeaseSpec) *int {
		if v == nil {
			return nil
		}
		return v.LeaseTransitions
	}).(pulumi.IntPtrOutput)
}

// renewTime is a time when the current holder of a lease has last updated the lease.
func (o LeaseSpecPtrOutput) RenewTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LeaseSpec) *string {
		if v == nil {
			return nil
		}
		return v.RenewTime
	}).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*LeaseTypeInput)(nil)).Elem(), LeaseTypeArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*LeaseTypeArrayInput)(nil)).Elem(), LeaseTypeArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*LeaseListTypeInput)(nil)).Elem(), LeaseListTypeArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*LeaseSpecInput)(nil)).Elem(), LeaseSpecArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*LeaseSpecPtrInput)(nil)).Elem(), LeaseSpecArgs{})
	pulumi.RegisterOutputType(LeaseTypeOutput{})
	pulumi.RegisterOutputType(LeaseTypeArrayOutput{})
	pulumi.RegisterOutputType(LeaseListTypeOutput{})
	pulumi.RegisterOutputType(LeaseSpecOutput{})
	pulumi.RegisterOutputType(LeaseSpecPtrOutput{})
}
