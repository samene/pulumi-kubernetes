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

// CSIStorageCapacity stores the result of one CSI GetCapacity call. For a given StorageClass, this describes the available capacity in a particular topology segment.  This can be used when considering where to instantiate new PersistentVolumes.
//
// For example this can express things like: - StorageClass "standard" has "1234 GiB" available in "topology.kubernetes.io/zone=us-east1" - StorageClass "localssd" has "10 GiB" available in "kubernetes.io/hostname=knode-abc123"
//
// The following three cases all imply that no capacity is available for a certain combination: - no object exists with suitable topology and storage class name - such an object exists, but the capacity is unset - such an object exists, but the capacity is zero
//
// The producer of these objects can decide which approach is more suitable.
//
// They are consumed by the kube-scheduler when a CSI driver opts into capacity-aware scheduling with CSIDriverSpec.StorageCapacity. The scheduler compares the MaximumVolumeSize against the requested size of pending volumes to filter out unsuitable nodes. If MaximumVolumeSize is unset, it falls back to a comparison against the less precise Capacity. If that is also unset, the scheduler assumes that capacity is insufficient and tries some other node.
type CSIStorageCapacity struct {
	pulumi.CustomResourceState

	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion pulumi.StringOutput `pulumi:"apiVersion"`
	// capacity is the value reported by the CSI driver in its GetCapacityResponse for a GetCapacityRequest with topology and parameters that match the previous fields.
	//
	// The semantic is currently (CSI spec 1.2) defined as: The available capacity, in bytes, of the storage that can be used to provision volumes. If not set, that information is currently unavailable.
	Capacity pulumi.StringOutput `pulumi:"capacity"`
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind pulumi.StringOutput `pulumi:"kind"`
	// maximumVolumeSize is the value reported by the CSI driver in its GetCapacityResponse for a GetCapacityRequest with topology and parameters that match the previous fields.
	//
	// This is defined since CSI spec 1.4.0 as the largest size that may be used in a CreateVolumeRequest.capacity_range.required_bytes field to create a volume with the same parameters as those in GetCapacityRequest. The corresponding value in the Kubernetes API is ResourceRequirements.Requests in a volume claim.
	MaximumVolumeSize pulumi.StringOutput `pulumi:"maximumVolumeSize"`
	// Standard object's metadata. The name has no particular meaning. It must be a DNS subdomain (dots allowed, 253 characters). To ensure that there are no conflicts with other CSI drivers on the cluster, the recommendation is to use csisc-<uuid>, a generated name, or a reverse-domain name which ends with the unique CSI driver name.
	//
	// Objects are namespaced.
	//
	// More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	Metadata metav1.ObjectMetaOutput `pulumi:"metadata"`
	// nodeTopology defines which nodes have access to the storage for which capacity was reported. If not set, the storage is not accessible from any node in the cluster. If empty, the storage is accessible from all nodes. This field is immutable.
	NodeTopology metav1.LabelSelectorOutput `pulumi:"nodeTopology"`
	// storageClassName represents the name of the StorageClass that the reported capacity applies to. It must meet the same requirements as the name of a StorageClass object (non-empty, DNS subdomain). If that object no longer exists, the CSIStorageCapacity object is obsolete and should be removed by its creator. This field is immutable.
	StorageClassName pulumi.StringOutput `pulumi:"storageClassName"`
}

// NewCSIStorageCapacity registers a new resource with the given unique name, arguments, and options.
func NewCSIStorageCapacity(ctx *pulumi.Context,
	name string, args *CSIStorageCapacityArgs, opts ...pulumi.ResourceOption) (*CSIStorageCapacity, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.StorageClassName == nil {
		return nil, errors.New("invalid value for required argument 'StorageClassName'")
	}
	args.ApiVersion = pulumi.StringPtr("storage.k8s.io/v1")
	args.Kind = pulumi.StringPtr("CSIStorageCapacity")
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("kubernetes:storage.k8s.io/v1beta1:CSIStorageCapacity"),
		},
		{
			Type: pulumi.String("kubernetes:storage.k8s.io/v1alpha1:CSIStorageCapacity"),
		},
	})
	opts = append(opts, aliases)
	var resource CSIStorageCapacity
	err := ctx.RegisterResource("kubernetes:storage.k8s.io/v1:CSIStorageCapacity", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetCSIStorageCapacity gets an existing CSIStorageCapacity resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetCSIStorageCapacity(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *CSIStorageCapacityState, opts ...pulumi.ResourceOption) (*CSIStorageCapacity, error) {
	var resource CSIStorageCapacity
	err := ctx.ReadResource("kubernetes:storage.k8s.io/v1:CSIStorageCapacity", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering CSIStorageCapacity resources.
type csistorageCapacityState struct {
}

type CSIStorageCapacityState struct {
}

func (CSIStorageCapacityState) ElementType() reflect.Type {
	return reflect.TypeOf((*csistorageCapacityState)(nil)).Elem()
}

type csistorageCapacityArgs struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion *string `pulumi:"apiVersion"`
	// capacity is the value reported by the CSI driver in its GetCapacityResponse for a GetCapacityRequest with topology and parameters that match the previous fields.
	//
	// The semantic is currently (CSI spec 1.2) defined as: The available capacity, in bytes, of the storage that can be used to provision volumes. If not set, that information is currently unavailable.
	Capacity *string `pulumi:"capacity"`
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind *string `pulumi:"kind"`
	// maximumVolumeSize is the value reported by the CSI driver in its GetCapacityResponse for a GetCapacityRequest with topology and parameters that match the previous fields.
	//
	// This is defined since CSI spec 1.4.0 as the largest size that may be used in a CreateVolumeRequest.capacity_range.required_bytes field to create a volume with the same parameters as those in GetCapacityRequest. The corresponding value in the Kubernetes API is ResourceRequirements.Requests in a volume claim.
	MaximumVolumeSize *string `pulumi:"maximumVolumeSize"`
	// Standard object's metadata. The name has no particular meaning. It must be a DNS subdomain (dots allowed, 253 characters). To ensure that there are no conflicts with other CSI drivers on the cluster, the recommendation is to use csisc-<uuid>, a generated name, or a reverse-domain name which ends with the unique CSI driver name.
	//
	// Objects are namespaced.
	//
	// More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	Metadata *metav1.ObjectMeta `pulumi:"metadata"`
	// nodeTopology defines which nodes have access to the storage for which capacity was reported. If not set, the storage is not accessible from any node in the cluster. If empty, the storage is accessible from all nodes. This field is immutable.
	NodeTopology *metav1.LabelSelector `pulumi:"nodeTopology"`
	// storageClassName represents the name of the StorageClass that the reported capacity applies to. It must meet the same requirements as the name of a StorageClass object (non-empty, DNS subdomain). If that object no longer exists, the CSIStorageCapacity object is obsolete and should be removed by its creator. This field is immutable.
	StorageClassName string `pulumi:"storageClassName"`
}

// The set of arguments for constructing a CSIStorageCapacity resource.
type CSIStorageCapacityArgs struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion pulumi.StringPtrInput
	// capacity is the value reported by the CSI driver in its GetCapacityResponse for a GetCapacityRequest with topology and parameters that match the previous fields.
	//
	// The semantic is currently (CSI spec 1.2) defined as: The available capacity, in bytes, of the storage that can be used to provision volumes. If not set, that information is currently unavailable.
	Capacity pulumi.StringPtrInput
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind pulumi.StringPtrInput
	// maximumVolumeSize is the value reported by the CSI driver in its GetCapacityResponse for a GetCapacityRequest with topology and parameters that match the previous fields.
	//
	// This is defined since CSI spec 1.4.0 as the largest size that may be used in a CreateVolumeRequest.capacity_range.required_bytes field to create a volume with the same parameters as those in GetCapacityRequest. The corresponding value in the Kubernetes API is ResourceRequirements.Requests in a volume claim.
	MaximumVolumeSize pulumi.StringPtrInput
	// Standard object's metadata. The name has no particular meaning. It must be a DNS subdomain (dots allowed, 253 characters). To ensure that there are no conflicts with other CSI drivers on the cluster, the recommendation is to use csisc-<uuid>, a generated name, or a reverse-domain name which ends with the unique CSI driver name.
	//
	// Objects are namespaced.
	//
	// More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	Metadata metav1.ObjectMetaPtrInput
	// nodeTopology defines which nodes have access to the storage for which capacity was reported. If not set, the storage is not accessible from any node in the cluster. If empty, the storage is accessible from all nodes. This field is immutable.
	NodeTopology metav1.LabelSelectorPtrInput
	// storageClassName represents the name of the StorageClass that the reported capacity applies to. It must meet the same requirements as the name of a StorageClass object (non-empty, DNS subdomain). If that object no longer exists, the CSIStorageCapacity object is obsolete and should be removed by its creator. This field is immutable.
	StorageClassName pulumi.StringInput
}

func (CSIStorageCapacityArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*csistorageCapacityArgs)(nil)).Elem()
}

type CSIStorageCapacityInput interface {
	pulumi.Input

	ToCSIStorageCapacityOutput() CSIStorageCapacityOutput
	ToCSIStorageCapacityOutputWithContext(ctx context.Context) CSIStorageCapacityOutput
}

func (*CSIStorageCapacity) ElementType() reflect.Type {
	return reflect.TypeOf((**CSIStorageCapacity)(nil)).Elem()
}

func (i *CSIStorageCapacity) ToCSIStorageCapacityOutput() CSIStorageCapacityOutput {
	return i.ToCSIStorageCapacityOutputWithContext(context.Background())
}

func (i *CSIStorageCapacity) ToCSIStorageCapacityOutputWithContext(ctx context.Context) CSIStorageCapacityOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CSIStorageCapacityOutput)
}

// CSIStorageCapacityArrayInput is an input type that accepts CSIStorageCapacityArray and CSIStorageCapacityArrayOutput values.
// You can construct a concrete instance of `CSIStorageCapacityArrayInput` via:
//
//	CSIStorageCapacityArray{ CSIStorageCapacityArgs{...} }
type CSIStorageCapacityArrayInput interface {
	pulumi.Input

	ToCSIStorageCapacityArrayOutput() CSIStorageCapacityArrayOutput
	ToCSIStorageCapacityArrayOutputWithContext(context.Context) CSIStorageCapacityArrayOutput
}

type CSIStorageCapacityArray []CSIStorageCapacityInput

func (CSIStorageCapacityArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*CSIStorageCapacity)(nil)).Elem()
}

func (i CSIStorageCapacityArray) ToCSIStorageCapacityArrayOutput() CSIStorageCapacityArrayOutput {
	return i.ToCSIStorageCapacityArrayOutputWithContext(context.Background())
}

func (i CSIStorageCapacityArray) ToCSIStorageCapacityArrayOutputWithContext(ctx context.Context) CSIStorageCapacityArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CSIStorageCapacityArrayOutput)
}

// CSIStorageCapacityMapInput is an input type that accepts CSIStorageCapacityMap and CSIStorageCapacityMapOutput values.
// You can construct a concrete instance of `CSIStorageCapacityMapInput` via:
//
//	CSIStorageCapacityMap{ "key": CSIStorageCapacityArgs{...} }
type CSIStorageCapacityMapInput interface {
	pulumi.Input

	ToCSIStorageCapacityMapOutput() CSIStorageCapacityMapOutput
	ToCSIStorageCapacityMapOutputWithContext(context.Context) CSIStorageCapacityMapOutput
}

type CSIStorageCapacityMap map[string]CSIStorageCapacityInput

func (CSIStorageCapacityMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*CSIStorageCapacity)(nil)).Elem()
}

func (i CSIStorageCapacityMap) ToCSIStorageCapacityMapOutput() CSIStorageCapacityMapOutput {
	return i.ToCSIStorageCapacityMapOutputWithContext(context.Background())
}

func (i CSIStorageCapacityMap) ToCSIStorageCapacityMapOutputWithContext(ctx context.Context) CSIStorageCapacityMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CSIStorageCapacityMapOutput)
}

type CSIStorageCapacityOutput struct{ *pulumi.OutputState }

func (CSIStorageCapacityOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CSIStorageCapacity)(nil)).Elem()
}

func (o CSIStorageCapacityOutput) ToCSIStorageCapacityOutput() CSIStorageCapacityOutput {
	return o
}

func (o CSIStorageCapacityOutput) ToCSIStorageCapacityOutputWithContext(ctx context.Context) CSIStorageCapacityOutput {
	return o
}

// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
func (o CSIStorageCapacityOutput) ApiVersion() pulumi.StringOutput {
	return o.ApplyT(func(v *CSIStorageCapacity) pulumi.StringOutput { return v.ApiVersion }).(pulumi.StringOutput)
}

// capacity is the value reported by the CSI driver in its GetCapacityResponse for a GetCapacityRequest with topology and parameters that match the previous fields.
//
// The semantic is currently (CSI spec 1.2) defined as: The available capacity, in bytes, of the storage that can be used to provision volumes. If not set, that information is currently unavailable.
func (o CSIStorageCapacityOutput) Capacity() pulumi.StringOutput {
	return o.ApplyT(func(v *CSIStorageCapacity) pulumi.StringOutput { return v.Capacity }).(pulumi.StringOutput)
}

// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
func (o CSIStorageCapacityOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v *CSIStorageCapacity) pulumi.StringOutput { return v.Kind }).(pulumi.StringOutput)
}

// maximumVolumeSize is the value reported by the CSI driver in its GetCapacityResponse for a GetCapacityRequest with topology and parameters that match the previous fields.
//
// This is defined since CSI spec 1.4.0 as the largest size that may be used in a CreateVolumeRequest.capacity_range.required_bytes field to create a volume with the same parameters as those in GetCapacityRequest. The corresponding value in the Kubernetes API is ResourceRequirements.Requests in a volume claim.
func (o CSIStorageCapacityOutput) MaximumVolumeSize() pulumi.StringOutput {
	return o.ApplyT(func(v *CSIStorageCapacity) pulumi.StringOutput { return v.MaximumVolumeSize }).(pulumi.StringOutput)
}

// Standard object's metadata. The name has no particular meaning. It must be a DNS subdomain (dots allowed, 253 characters). To ensure that there are no conflicts with other CSI drivers on the cluster, the recommendation is to use csisc-<uuid>, a generated name, or a reverse-domain name which ends with the unique CSI driver name.
//
// Objects are namespaced.
//
// More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
func (o CSIStorageCapacityOutput) Metadata() metav1.ObjectMetaOutput {
	return o.ApplyT(func(v *CSIStorageCapacity) metav1.ObjectMetaOutput { return v.Metadata }).(metav1.ObjectMetaOutput)
}

// nodeTopology defines which nodes have access to the storage for which capacity was reported. If not set, the storage is not accessible from any node in the cluster. If empty, the storage is accessible from all nodes. This field is immutable.
func (o CSIStorageCapacityOutput) NodeTopology() metav1.LabelSelectorOutput {
	return o.ApplyT(func(v *CSIStorageCapacity) metav1.LabelSelectorOutput { return v.NodeTopology }).(metav1.LabelSelectorOutput)
}

// storageClassName represents the name of the StorageClass that the reported capacity applies to. It must meet the same requirements as the name of a StorageClass object (non-empty, DNS subdomain). If that object no longer exists, the CSIStorageCapacity object is obsolete and should be removed by its creator. This field is immutable.
func (o CSIStorageCapacityOutput) StorageClassName() pulumi.StringOutput {
	return o.ApplyT(func(v *CSIStorageCapacity) pulumi.StringOutput { return v.StorageClassName }).(pulumi.StringOutput)
}

type CSIStorageCapacityArrayOutput struct{ *pulumi.OutputState }

func (CSIStorageCapacityArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*CSIStorageCapacity)(nil)).Elem()
}

func (o CSIStorageCapacityArrayOutput) ToCSIStorageCapacityArrayOutput() CSIStorageCapacityArrayOutput {
	return o
}

func (o CSIStorageCapacityArrayOutput) ToCSIStorageCapacityArrayOutputWithContext(ctx context.Context) CSIStorageCapacityArrayOutput {
	return o
}

func (o CSIStorageCapacityArrayOutput) Index(i pulumi.IntInput) CSIStorageCapacityOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *CSIStorageCapacity {
		return vs[0].([]*CSIStorageCapacity)[vs[1].(int)]
	}).(CSIStorageCapacityOutput)
}

type CSIStorageCapacityMapOutput struct{ *pulumi.OutputState }

func (CSIStorageCapacityMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*CSIStorageCapacity)(nil)).Elem()
}

func (o CSIStorageCapacityMapOutput) ToCSIStorageCapacityMapOutput() CSIStorageCapacityMapOutput {
	return o
}

func (o CSIStorageCapacityMapOutput) ToCSIStorageCapacityMapOutputWithContext(ctx context.Context) CSIStorageCapacityMapOutput {
	return o
}

func (o CSIStorageCapacityMapOutput) MapIndex(k pulumi.StringInput) CSIStorageCapacityOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *CSIStorageCapacity {
		return vs[0].(map[string]*CSIStorageCapacity)[vs[1].(string)]
	}).(CSIStorageCapacityOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*CSIStorageCapacityInput)(nil)).Elem(), &CSIStorageCapacity{})
	pulumi.RegisterInputType(reflect.TypeOf((*CSIStorageCapacityArrayInput)(nil)).Elem(), CSIStorageCapacityArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*CSIStorageCapacityMapInput)(nil)).Elem(), CSIStorageCapacityMap{})
	pulumi.RegisterOutputType(CSIStorageCapacityOutput{})
	pulumi.RegisterOutputType(CSIStorageCapacityArrayOutput{})
	pulumi.RegisterOutputType(CSIStorageCapacityMapOutput{})
}
