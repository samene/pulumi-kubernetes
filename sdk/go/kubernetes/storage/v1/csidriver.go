// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v1

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	metav1 "github.com/pulumi/pulumi-kubernetes/sdk/v3/go/kubernetes/meta/v1"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// CSIDriver captures information about a Container Storage Interface (CSI) volume driver deployed on the cluster. Kubernetes attach detach controller uses this object to determine whether attach is required. Kubelet uses this object to determine whether pod information needs to be passed on mount. CSIDriver objects are non-namespaced.
type CSIDriver struct {
	pulumi.CustomResourceState

	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion pulumi.StringPtrOutput `pulumi:"apiVersion"`
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind pulumi.StringPtrOutput `pulumi:"kind"`
	// Standard object metadata. metadata.Name indicates the name of the CSI driver that this object refers to; it MUST be the same name returned by the CSI GetPluginName() call for that driver. The driver name must be 63 characters or less, beginning and ending with an alphanumeric character ([a-z0-9A-Z]) with dashes (-), dots (.), and alphanumerics between. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	Metadata metav1.ObjectMetaPtrOutput `pulumi:"metadata"`
	// Specification of the CSI Driver.
	Spec CSIDriverSpecOutput `pulumi:"spec"`
}

// NewCSIDriver registers a new resource with the given unique name, arguments, and options.
func NewCSIDriver(ctx *pulumi.Context,
	name string, args *CSIDriverArgs, opts ...pulumi.ResourceOption) (*CSIDriver, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Spec == nil {
		return nil, errors.New("invalid value for required argument 'Spec'")
	}
	args.ApiVersion = pulumi.StringPtr("storage.k8s.io/v1")
	args.Kind = pulumi.StringPtr("CSIDriver")
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("kubernetes:storage.k8s.io/v1beta1:CSIDriver"),
		},
	})
	opts = append(opts, aliases)
	var resource CSIDriver
	err := ctx.RegisterResource("kubernetes:storage.k8s.io/v1:CSIDriver", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetCSIDriver gets an existing CSIDriver resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetCSIDriver(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *CSIDriverState, opts ...pulumi.ResourceOption) (*CSIDriver, error) {
	var resource CSIDriver
	err := ctx.ReadResource("kubernetes:storage.k8s.io/v1:CSIDriver", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering CSIDriver resources.
type csidriverState struct {
}

type CSIDriverState struct {
}

func (CSIDriverState) ElementType() reflect.Type {
	return reflect.TypeOf((*csidriverState)(nil)).Elem()
}

type csidriverArgs struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion *string `pulumi:"apiVersion"`
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind *string `pulumi:"kind"`
	// Standard object metadata. metadata.Name indicates the name of the CSI driver that this object refers to; it MUST be the same name returned by the CSI GetPluginName() call for that driver. The driver name must be 63 characters or less, beginning and ending with an alphanumeric character ([a-z0-9A-Z]) with dashes (-), dots (.), and alphanumerics between. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	Metadata *metav1.ObjectMeta `pulumi:"metadata"`
	// Specification of the CSI Driver.
	Spec CSIDriverSpec `pulumi:"spec"`
}

// The set of arguments for constructing a CSIDriver resource.
type CSIDriverArgs struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion pulumi.StringPtrInput
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind pulumi.StringPtrInput
	// Standard object metadata. metadata.Name indicates the name of the CSI driver that this object refers to; it MUST be the same name returned by the CSI GetPluginName() call for that driver. The driver name must be 63 characters or less, beginning and ending with an alphanumeric character ([a-z0-9A-Z]) with dashes (-), dots (.), and alphanumerics between. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	Metadata metav1.ObjectMetaPtrInput
	// Specification of the CSI Driver.
	Spec CSIDriverSpecInput
}

func (CSIDriverArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*csidriverArgs)(nil)).Elem()
}

type CSIDriverInput interface {
	pulumi.Input

	ToCSIDriverOutput() CSIDriverOutput
	ToCSIDriverOutputWithContext(ctx context.Context) CSIDriverOutput
}

func (*CSIDriver) ElementType() reflect.Type {
	return reflect.TypeOf((**CSIDriver)(nil)).Elem()
}

func (i *CSIDriver) ToCSIDriverOutput() CSIDriverOutput {
	return i.ToCSIDriverOutputWithContext(context.Background())
}

func (i *CSIDriver) ToCSIDriverOutputWithContext(ctx context.Context) CSIDriverOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CSIDriverOutput)
}

// CSIDriverArrayInput is an input type that accepts CSIDriverArray and CSIDriverArrayOutput values.
// You can construct a concrete instance of `CSIDriverArrayInput` via:
//
//          CSIDriverArray{ CSIDriverArgs{...} }
type CSIDriverArrayInput interface {
	pulumi.Input

	ToCSIDriverArrayOutput() CSIDriverArrayOutput
	ToCSIDriverArrayOutputWithContext(context.Context) CSIDriverArrayOutput
}

type CSIDriverArray []CSIDriverInput

func (CSIDriverArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*CSIDriver)(nil)).Elem()
}

func (i CSIDriverArray) ToCSIDriverArrayOutput() CSIDriverArrayOutput {
	return i.ToCSIDriverArrayOutputWithContext(context.Background())
}

func (i CSIDriverArray) ToCSIDriverArrayOutputWithContext(ctx context.Context) CSIDriverArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CSIDriverArrayOutput)
}

// CSIDriverMapInput is an input type that accepts CSIDriverMap and CSIDriverMapOutput values.
// You can construct a concrete instance of `CSIDriverMapInput` via:
//
//          CSIDriverMap{ "key": CSIDriverArgs{...} }
type CSIDriverMapInput interface {
	pulumi.Input

	ToCSIDriverMapOutput() CSIDriverMapOutput
	ToCSIDriverMapOutputWithContext(context.Context) CSIDriverMapOutput
}

type CSIDriverMap map[string]CSIDriverInput

func (CSIDriverMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*CSIDriver)(nil)).Elem()
}

func (i CSIDriverMap) ToCSIDriverMapOutput() CSIDriverMapOutput {
	return i.ToCSIDriverMapOutputWithContext(context.Background())
}

func (i CSIDriverMap) ToCSIDriverMapOutputWithContext(ctx context.Context) CSIDriverMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CSIDriverMapOutput)
}

type CSIDriverOutput struct{ *pulumi.OutputState }

func (CSIDriverOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CSIDriver)(nil)).Elem()
}

func (o CSIDriverOutput) ToCSIDriverOutput() CSIDriverOutput {
	return o
}

func (o CSIDriverOutput) ToCSIDriverOutputWithContext(ctx context.Context) CSIDriverOutput {
	return o
}

type CSIDriverArrayOutput struct{ *pulumi.OutputState }

func (CSIDriverArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*CSIDriver)(nil)).Elem()
}

func (o CSIDriverArrayOutput) ToCSIDriverArrayOutput() CSIDriverArrayOutput {
	return o
}

func (o CSIDriverArrayOutput) ToCSIDriverArrayOutputWithContext(ctx context.Context) CSIDriverArrayOutput {
	return o
}

func (o CSIDriverArrayOutput) Index(i pulumi.IntInput) CSIDriverOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *CSIDriver {
		return vs[0].([]*CSIDriver)[vs[1].(int)]
	}).(CSIDriverOutput)
}

type CSIDriverMapOutput struct{ *pulumi.OutputState }

func (CSIDriverMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*CSIDriver)(nil)).Elem()
}

func (o CSIDriverMapOutput) ToCSIDriverMapOutput() CSIDriverMapOutput {
	return o
}

func (o CSIDriverMapOutput) ToCSIDriverMapOutputWithContext(ctx context.Context) CSIDriverMapOutput {
	return o
}

func (o CSIDriverMapOutput) MapIndex(k pulumi.StringInput) CSIDriverOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *CSIDriver {
		return vs[0].(map[string]*CSIDriver)[vs[1].(string)]
	}).(CSIDriverOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*CSIDriverInput)(nil)).Elem(), &CSIDriver{})
	pulumi.RegisterInputType(reflect.TypeOf((*CSIDriverArrayInput)(nil)).Elem(), CSIDriverArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*CSIDriverMapInput)(nil)).Elem(), CSIDriverMap{})
	pulumi.RegisterOutputType(CSIDriverOutput{})
	pulumi.RegisterOutputType(CSIDriverArrayOutput{})
	pulumi.RegisterOutputType(CSIDriverMapOutput{})
}
