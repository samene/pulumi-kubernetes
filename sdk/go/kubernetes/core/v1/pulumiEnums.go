// Code generated by pulumigen DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v1

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

type ServiceSpecType string

const (
	ServiceSpecTypeExternalName = ServiceSpecType("ExternalName")
	ServiceSpecTypeClusterIP    = ServiceSpecType("ClusterIP")
	ServiceSpecTypeNodePort     = ServiceSpecType("NodePort")
	ServiceSpecTypeLoadBalancer = ServiceSpecType("LoadBalancer")
)

func (ServiceSpecType) ElementType() reflect.Type {
	return reflect.TypeOf((*ServiceSpecType)(nil)).Elem()
}

func (e ServiceSpecType) ToServiceSpecTypeOutput() ServiceSpecTypeOutput {
	return pulumi.ToOutput(e).(ServiceSpecTypeOutput)
}

func (e ServiceSpecType) ToServiceSpecTypeOutputWithContext(ctx context.Context) ServiceSpecTypeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(ServiceSpecTypeOutput)
}

func (e ServiceSpecType) ToServiceSpecTypePtrOutput() ServiceSpecTypePtrOutput {
	return e.ToServiceSpecTypePtrOutputWithContext(context.Background())
}

func (e ServiceSpecType) ToServiceSpecTypePtrOutputWithContext(ctx context.Context) ServiceSpecTypePtrOutput {
	return ServiceSpecType(e).ToServiceSpecTypeOutputWithContext(ctx).ToServiceSpecTypePtrOutputWithContext(ctx)
}

func (e ServiceSpecType) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e ServiceSpecType) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e ServiceSpecType) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e ServiceSpecType) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type ServiceSpecTypeOutput struct{ *pulumi.OutputState }

func (ServiceSpecTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ServiceSpecType)(nil)).Elem()
}

func (o ServiceSpecTypeOutput) ToServiceSpecTypeOutput() ServiceSpecTypeOutput {
	return o
}

func (o ServiceSpecTypeOutput) ToServiceSpecTypeOutputWithContext(ctx context.Context) ServiceSpecTypeOutput {
	return o
}

func (o ServiceSpecTypeOutput) ToServiceSpecTypePtrOutput() ServiceSpecTypePtrOutput {
	return o.ToServiceSpecTypePtrOutputWithContext(context.Background())
}

func (o ServiceSpecTypeOutput) ToServiceSpecTypePtrOutputWithContext(ctx context.Context) ServiceSpecTypePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ServiceSpecType) *ServiceSpecType {
		return &v
	}).(ServiceSpecTypePtrOutput)
}

func (o ServiceSpecTypeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o ServiceSpecTypeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ServiceSpecType) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o ServiceSpecTypeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ServiceSpecTypeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ServiceSpecType) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type ServiceSpecTypePtrOutput struct{ *pulumi.OutputState }

func (ServiceSpecTypePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ServiceSpecType)(nil)).Elem()
}

func (o ServiceSpecTypePtrOutput) ToServiceSpecTypePtrOutput() ServiceSpecTypePtrOutput {
	return o
}

func (o ServiceSpecTypePtrOutput) ToServiceSpecTypePtrOutputWithContext(ctx context.Context) ServiceSpecTypePtrOutput {
	return o
}

func (o ServiceSpecTypePtrOutput) Elem() ServiceSpecTypeOutput {
	return o.ApplyT(func(v *ServiceSpecType) ServiceSpecType {
		if v != nil {
			return *v
		}
		var ret ServiceSpecType
		return ret
	}).(ServiceSpecTypeOutput)
}

func (o ServiceSpecTypePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ServiceSpecTypePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *ServiceSpecType) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}

// ServiceSpecTypeInput is an input type that accepts values of the ServiceSpecType enum
// A concrete instance of `ServiceSpecTypeInput` can be one of the following:
//
//	ServiceSpecTypeExternalName
//	ServiceSpecTypeClusterIP
//	ServiceSpecTypeNodePort
//	ServiceSpecTypeLoadBalancer
type ServiceSpecTypeInput interface {
	pulumi.Input

	ToServiceSpecTypeOutput() ServiceSpecTypeOutput
	ToServiceSpecTypeOutputWithContext(context.Context) ServiceSpecTypeOutput
}

var serviceSpecTypePtrType = reflect.TypeOf((**ServiceSpecType)(nil)).Elem()

type ServiceSpecTypePtrInput interface {
	pulumi.Input

	ToServiceSpecTypePtrOutput() ServiceSpecTypePtrOutput
	ToServiceSpecTypePtrOutputWithContext(context.Context) ServiceSpecTypePtrOutput
}

type serviceSpecTypePtr string

func ServiceSpecTypePtr(v string) ServiceSpecTypePtrInput {
	return (*serviceSpecTypePtr)(&v)
}

func (*serviceSpecTypePtr) ElementType() reflect.Type {
	return serviceSpecTypePtrType
}

func (in *serviceSpecTypePtr) ToServiceSpecTypePtrOutput() ServiceSpecTypePtrOutput {
	return pulumi.ToOutput(in).(ServiceSpecTypePtrOutput)
}

func (in *serviceSpecTypePtr) ToServiceSpecTypePtrOutputWithContext(ctx context.Context) ServiceSpecTypePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(ServiceSpecTypePtrOutput)
}

func (in *serviceSpecTypePtr) ToOutput(ctx context.Context) pulumix.Output[*ServiceSpecType] {
	return pulumix.Output[*ServiceSpecType]{
		OutputState: in.ToServiceSpecTypePtrOutputWithContext(ctx).OutputState,
	}
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ServiceSpecTypeInput)(nil)).Elem(), ServiceSpecType("ExternalName"))
	pulumi.RegisterInputType(reflect.TypeOf((*ServiceSpecTypePtrInput)(nil)).Elem(), ServiceSpecType("ExternalName"))
	pulumi.RegisterOutputType(ServiceSpecTypeOutput{})
	pulumi.RegisterOutputType(ServiceSpecTypePtrOutput{})
}
