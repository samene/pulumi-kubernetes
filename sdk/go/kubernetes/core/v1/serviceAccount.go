// Code generated by pulumigen DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v1

import (
	"context"
	"reflect"

	metav1 "github.com/pulumi/pulumi-kubernetes/sdk/v3/go/kubernetes/meta/v1"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// ServiceAccount binds together: * a name, understood by users, and perhaps by peripheral systems, for an identity * a principal that can be authenticated and authorized * a set of secrets
type ServiceAccount struct {
	pulumi.CustomResourceState

	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion pulumi.StringOutput `pulumi:"apiVersion"`
	// AutomountServiceAccountToken indicates whether pods running as this service account should have an API token automatically mounted. Can be overridden at the pod level.
	AutomountServiceAccountToken pulumi.BoolOutput `pulumi:"automountServiceAccountToken"`
	// ImagePullSecrets is a list of references to secrets in the same namespace to use for pulling any images in pods that reference this ServiceAccount. ImagePullSecrets are distinct from Secrets because Secrets can be mounted in the pod, but ImagePullSecrets are only accessed by the kubelet. More info: https://kubernetes.io/docs/concepts/containers/images/#specifying-imagepullsecrets-on-a-pod
	ImagePullSecrets LocalObjectReferenceArrayOutput `pulumi:"imagePullSecrets"`
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind pulumi.StringOutput `pulumi:"kind"`
	// Standard object's metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	Metadata metav1.ObjectMetaOutput `pulumi:"metadata"`
	// Secrets is a list of the secrets in the same namespace that pods running using this ServiceAccount are allowed to use. Pods are only limited to this list if this service account has a "kubernetes.io/enforce-mountable-secrets" annotation set to "true". This field should not be used to find auto-generated service account token secrets for use outside of pods. Instead, tokens can be requested directly using the TokenRequest API, or service account token secrets can be manually created. More info: https://kubernetes.io/docs/concepts/configuration/secret
	Secrets ObjectReferenceArrayOutput `pulumi:"secrets"`
}

// NewServiceAccount registers a new resource with the given unique name, arguments, and options.
func NewServiceAccount(ctx *pulumi.Context,
	name string, args *ServiceAccountArgs, opts ...pulumi.ResourceOption) (*ServiceAccount, error) {
	if args == nil {
		args = &ServiceAccountArgs{}
	}

	args.ApiVersion = pulumi.StringPtr("v1")
	args.Kind = pulumi.StringPtr("ServiceAccount")
	var resource ServiceAccount
	err := ctx.RegisterResource("kubernetes:core/v1:ServiceAccount", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetServiceAccount gets an existing ServiceAccount resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetServiceAccount(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ServiceAccountState, opts ...pulumi.ResourceOption) (*ServiceAccount, error) {
	var resource ServiceAccount
	err := ctx.ReadResource("kubernetes:core/v1:ServiceAccount", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ServiceAccount resources.
type serviceAccountState struct {
}

type ServiceAccountState struct {
}

func (ServiceAccountState) ElementType() reflect.Type {
	return reflect.TypeOf((*serviceAccountState)(nil)).Elem()
}

type serviceAccountArgs struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion *string `pulumi:"apiVersion"`
	// AutomountServiceAccountToken indicates whether pods running as this service account should have an API token automatically mounted. Can be overridden at the pod level.
	AutomountServiceAccountToken *bool `pulumi:"automountServiceAccountToken"`
	// ImagePullSecrets is a list of references to secrets in the same namespace to use for pulling any images in pods that reference this ServiceAccount. ImagePullSecrets are distinct from Secrets because Secrets can be mounted in the pod, but ImagePullSecrets are only accessed by the kubelet. More info: https://kubernetes.io/docs/concepts/containers/images/#specifying-imagepullsecrets-on-a-pod
	ImagePullSecrets []LocalObjectReference `pulumi:"imagePullSecrets"`
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind *string `pulumi:"kind"`
	// Standard object's metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	Metadata *metav1.ObjectMeta `pulumi:"metadata"`
	// Secrets is a list of the secrets in the same namespace that pods running using this ServiceAccount are allowed to use. Pods are only limited to this list if this service account has a "kubernetes.io/enforce-mountable-secrets" annotation set to "true". This field should not be used to find auto-generated service account token secrets for use outside of pods. Instead, tokens can be requested directly using the TokenRequest API, or service account token secrets can be manually created. More info: https://kubernetes.io/docs/concepts/configuration/secret
	Secrets []ObjectReference `pulumi:"secrets"`
}

// The set of arguments for constructing a ServiceAccount resource.
type ServiceAccountArgs struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion pulumi.StringPtrInput
	// AutomountServiceAccountToken indicates whether pods running as this service account should have an API token automatically mounted. Can be overridden at the pod level.
	AutomountServiceAccountToken pulumi.BoolPtrInput
	// ImagePullSecrets is a list of references to secrets in the same namespace to use for pulling any images in pods that reference this ServiceAccount. ImagePullSecrets are distinct from Secrets because Secrets can be mounted in the pod, but ImagePullSecrets are only accessed by the kubelet. More info: https://kubernetes.io/docs/concepts/containers/images/#specifying-imagepullsecrets-on-a-pod
	ImagePullSecrets LocalObjectReferenceArrayInput
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind pulumi.StringPtrInput
	// Standard object's metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	Metadata metav1.ObjectMetaPtrInput
	// Secrets is a list of the secrets in the same namespace that pods running using this ServiceAccount are allowed to use. Pods are only limited to this list if this service account has a "kubernetes.io/enforce-mountable-secrets" annotation set to "true". This field should not be used to find auto-generated service account token secrets for use outside of pods. Instead, tokens can be requested directly using the TokenRequest API, or service account token secrets can be manually created. More info: https://kubernetes.io/docs/concepts/configuration/secret
	Secrets ObjectReferenceArrayInput
}

func (ServiceAccountArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*serviceAccountArgs)(nil)).Elem()
}

type ServiceAccountInput interface {
	pulumi.Input

	ToServiceAccountOutput() ServiceAccountOutput
	ToServiceAccountOutputWithContext(ctx context.Context) ServiceAccountOutput
}

func (*ServiceAccount) ElementType() reflect.Type {
	return reflect.TypeOf((**ServiceAccount)(nil)).Elem()
}

func (i *ServiceAccount) ToServiceAccountOutput() ServiceAccountOutput {
	return i.ToServiceAccountOutputWithContext(context.Background())
}

func (i *ServiceAccount) ToServiceAccountOutputWithContext(ctx context.Context) ServiceAccountOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServiceAccountOutput)
}

// ServiceAccountArrayInput is an input type that accepts ServiceAccountArray and ServiceAccountArrayOutput values.
// You can construct a concrete instance of `ServiceAccountArrayInput` via:
//
//	ServiceAccountArray{ ServiceAccountArgs{...} }
type ServiceAccountArrayInput interface {
	pulumi.Input

	ToServiceAccountArrayOutput() ServiceAccountArrayOutput
	ToServiceAccountArrayOutputWithContext(context.Context) ServiceAccountArrayOutput
}

type ServiceAccountArray []ServiceAccountInput

func (ServiceAccountArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ServiceAccount)(nil)).Elem()
}

func (i ServiceAccountArray) ToServiceAccountArrayOutput() ServiceAccountArrayOutput {
	return i.ToServiceAccountArrayOutputWithContext(context.Background())
}

func (i ServiceAccountArray) ToServiceAccountArrayOutputWithContext(ctx context.Context) ServiceAccountArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServiceAccountArrayOutput)
}

// ServiceAccountMapInput is an input type that accepts ServiceAccountMap and ServiceAccountMapOutput values.
// You can construct a concrete instance of `ServiceAccountMapInput` via:
//
//	ServiceAccountMap{ "key": ServiceAccountArgs{...} }
type ServiceAccountMapInput interface {
	pulumi.Input

	ToServiceAccountMapOutput() ServiceAccountMapOutput
	ToServiceAccountMapOutputWithContext(context.Context) ServiceAccountMapOutput
}

type ServiceAccountMap map[string]ServiceAccountInput

func (ServiceAccountMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ServiceAccount)(nil)).Elem()
}

func (i ServiceAccountMap) ToServiceAccountMapOutput() ServiceAccountMapOutput {
	return i.ToServiceAccountMapOutputWithContext(context.Background())
}

func (i ServiceAccountMap) ToServiceAccountMapOutputWithContext(ctx context.Context) ServiceAccountMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServiceAccountMapOutput)
}

type ServiceAccountOutput struct{ *pulumi.OutputState }

func (ServiceAccountOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ServiceAccount)(nil)).Elem()
}

func (o ServiceAccountOutput) ToServiceAccountOutput() ServiceAccountOutput {
	return o
}

func (o ServiceAccountOutput) ToServiceAccountOutputWithContext(ctx context.Context) ServiceAccountOutput {
	return o
}

// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
func (o ServiceAccountOutput) ApiVersion() pulumi.StringOutput {
	return o.ApplyT(func(v *ServiceAccount) pulumi.StringOutput { return v.ApiVersion }).(pulumi.StringOutput)
}

// AutomountServiceAccountToken indicates whether pods running as this service account should have an API token automatically mounted. Can be overridden at the pod level.
func (o ServiceAccountOutput) AutomountServiceAccountToken() pulumi.BoolOutput {
	return o.ApplyT(func(v *ServiceAccount) pulumi.BoolOutput { return v.AutomountServiceAccountToken }).(pulumi.BoolOutput)
}

// ImagePullSecrets is a list of references to secrets in the same namespace to use for pulling any images in pods that reference this ServiceAccount. ImagePullSecrets are distinct from Secrets because Secrets can be mounted in the pod, but ImagePullSecrets are only accessed by the kubelet. More info: https://kubernetes.io/docs/concepts/containers/images/#specifying-imagepullsecrets-on-a-pod
func (o ServiceAccountOutput) ImagePullSecrets() LocalObjectReferenceArrayOutput {
	return o.ApplyT(func(v *ServiceAccount) LocalObjectReferenceArrayOutput { return v.ImagePullSecrets }).(LocalObjectReferenceArrayOutput)
}

// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
func (o ServiceAccountOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v *ServiceAccount) pulumi.StringOutput { return v.Kind }).(pulumi.StringOutput)
}

// Standard object's metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
func (o ServiceAccountOutput) Metadata() metav1.ObjectMetaOutput {
	return o.ApplyT(func(v *ServiceAccount) metav1.ObjectMetaOutput { return v.Metadata }).(metav1.ObjectMetaOutput)
}

// Secrets is a list of the secrets in the same namespace that pods running using this ServiceAccount are allowed to use. Pods are only limited to this list if this service account has a "kubernetes.io/enforce-mountable-secrets" annotation set to "true". This field should not be used to find auto-generated service account token secrets for use outside of pods. Instead, tokens can be requested directly using the TokenRequest API, or service account token secrets can be manually created. More info: https://kubernetes.io/docs/concepts/configuration/secret
func (o ServiceAccountOutput) Secrets() ObjectReferenceArrayOutput {
	return o.ApplyT(func(v *ServiceAccount) ObjectReferenceArrayOutput { return v.Secrets }).(ObjectReferenceArrayOutput)
}

type ServiceAccountArrayOutput struct{ *pulumi.OutputState }

func (ServiceAccountArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ServiceAccount)(nil)).Elem()
}

func (o ServiceAccountArrayOutput) ToServiceAccountArrayOutput() ServiceAccountArrayOutput {
	return o
}

func (o ServiceAccountArrayOutput) ToServiceAccountArrayOutputWithContext(ctx context.Context) ServiceAccountArrayOutput {
	return o
}

func (o ServiceAccountArrayOutput) Index(i pulumi.IntInput) ServiceAccountOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *ServiceAccount {
		return vs[0].([]*ServiceAccount)[vs[1].(int)]
	}).(ServiceAccountOutput)
}

type ServiceAccountMapOutput struct{ *pulumi.OutputState }

func (ServiceAccountMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ServiceAccount)(nil)).Elem()
}

func (o ServiceAccountMapOutput) ToServiceAccountMapOutput() ServiceAccountMapOutput {
	return o
}

func (o ServiceAccountMapOutput) ToServiceAccountMapOutputWithContext(ctx context.Context) ServiceAccountMapOutput {
	return o
}

func (o ServiceAccountMapOutput) MapIndex(k pulumi.StringInput) ServiceAccountOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *ServiceAccount {
		return vs[0].(map[string]*ServiceAccount)[vs[1].(string)]
	}).(ServiceAccountOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ServiceAccountInput)(nil)).Elem(), &ServiceAccount{})
	pulumi.RegisterInputType(reflect.TypeOf((*ServiceAccountArrayInput)(nil)).Elem(), ServiceAccountArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ServiceAccountMapInput)(nil)).Elem(), ServiceAccountMap{})
	pulumi.RegisterOutputType(ServiceAccountOutput{})
	pulumi.RegisterOutputType(ServiceAccountArrayOutput{})
	pulumi.RegisterOutputType(ServiceAccountMapOutput{})
}
