// Code generated by pulumigen DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v1beta1

import (
	"context"
	"reflect"

	"errors"
	metav1 "github.com/pulumi/pulumi-kubernetes/sdk/v4/go/kubernetes/meta/v1"
	"github.com/pulumi/pulumi-kubernetes/sdk/v4/go/kubernetes/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// EventList is a list of Event objects.
type EventList struct {
	pulumi.CustomResourceState

	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion pulumi.StringOutput `pulumi:"apiVersion"`
	// Items is a list of schema objects.
	Items EventTypeArrayOutput `pulumi:"items"`
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind pulumi.StringOutput `pulumi:"kind"`
	// Standard list metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	Metadata metav1.ListMetaOutput `pulumi:"metadata"`
}

// NewEventList registers a new resource with the given unique name, arguments, and options.
func NewEventList(ctx *pulumi.Context,
	name string, args *EventListArgs, opts ...pulumi.ResourceOption) (*EventList, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Items == nil {
		return nil, errors.New("invalid value for required argument 'Items'")
	}
	args.ApiVersion = pulumi.StringPtr("events.k8s.io/v1beta1")
	args.Kind = pulumi.StringPtr("EventList")
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource EventList
	err := ctx.RegisterResource("kubernetes:events.k8s.io/v1beta1:EventList", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetEventList gets an existing EventList resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetEventList(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *EventListState, opts ...pulumi.ResourceOption) (*EventList, error) {
	var resource EventList
	err := ctx.ReadResource("kubernetes:events.k8s.io/v1beta1:EventList", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering EventList resources.
type eventListState struct {
}

type EventListState struct {
}

func (EventListState) ElementType() reflect.Type {
	return reflect.TypeOf((*eventListState)(nil)).Elem()
}

type eventListArgs struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion *string `pulumi:"apiVersion"`
	// Items is a list of schema objects.
	Items []EventType `pulumi:"items"`
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind *string `pulumi:"kind"`
	// Standard list metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	Metadata *metav1.ListMeta `pulumi:"metadata"`
}

// The set of arguments for constructing a EventList resource.
type EventListArgs struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion pulumi.StringPtrInput
	// Items is a list of schema objects.
	Items EventTypeArrayInput
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind pulumi.StringPtrInput
	// Standard list metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	Metadata metav1.ListMetaPtrInput
}

func (EventListArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*eventListArgs)(nil)).Elem()
}

type EventListInput interface {
	pulumi.Input

	ToEventListOutput() EventListOutput
	ToEventListOutputWithContext(ctx context.Context) EventListOutput
}

func (*EventList) ElementType() reflect.Type {
	return reflect.TypeOf((**EventList)(nil)).Elem()
}

func (i *EventList) ToEventListOutput() EventListOutput {
	return i.ToEventListOutputWithContext(context.Background())
}

func (i *EventList) ToEventListOutputWithContext(ctx context.Context) EventListOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EventListOutput)
}

// EventListArrayInput is an input type that accepts EventListArray and EventListArrayOutput values.
// You can construct a concrete instance of `EventListArrayInput` via:
//
//	EventListArray{ EventListArgs{...} }
type EventListArrayInput interface {
	pulumi.Input

	ToEventListArrayOutput() EventListArrayOutput
	ToEventListArrayOutputWithContext(context.Context) EventListArrayOutput
}

type EventListArray []EventListInput

func (EventListArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*EventList)(nil)).Elem()
}

func (i EventListArray) ToEventListArrayOutput() EventListArrayOutput {
	return i.ToEventListArrayOutputWithContext(context.Background())
}

func (i EventListArray) ToEventListArrayOutputWithContext(ctx context.Context) EventListArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EventListArrayOutput)
}

// EventListMapInput is an input type that accepts EventListMap and EventListMapOutput values.
// You can construct a concrete instance of `EventListMapInput` via:
//
//	EventListMap{ "key": EventListArgs{...} }
type EventListMapInput interface {
	pulumi.Input

	ToEventListMapOutput() EventListMapOutput
	ToEventListMapOutputWithContext(context.Context) EventListMapOutput
}

type EventListMap map[string]EventListInput

func (EventListMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*EventList)(nil)).Elem()
}

func (i EventListMap) ToEventListMapOutput() EventListMapOutput {
	return i.ToEventListMapOutputWithContext(context.Background())
}

func (i EventListMap) ToEventListMapOutputWithContext(ctx context.Context) EventListMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EventListMapOutput)
}

type EventListOutput struct{ *pulumi.OutputState }

func (EventListOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**EventList)(nil)).Elem()
}

func (o EventListOutput) ToEventListOutput() EventListOutput {
	return o
}

func (o EventListOutput) ToEventListOutputWithContext(ctx context.Context) EventListOutput {
	return o
}

// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
func (o EventListOutput) ApiVersion() pulumi.StringOutput {
	return o.ApplyT(func(v *EventList) pulumi.StringOutput { return v.ApiVersion }).(pulumi.StringOutput)
}

// Items is a list of schema objects.
func (o EventListOutput) Items() EventTypeArrayOutput {
	return o.ApplyT(func(v *EventList) EventTypeArrayOutput { return v.Items }).(EventTypeArrayOutput)
}

// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
func (o EventListOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v *EventList) pulumi.StringOutput { return v.Kind }).(pulumi.StringOutput)
}

// Standard list metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
func (o EventListOutput) Metadata() metav1.ListMetaOutput {
	return o.ApplyT(func(v *EventList) metav1.ListMetaOutput { return v.Metadata }).(metav1.ListMetaOutput)
}

type EventListArrayOutput struct{ *pulumi.OutputState }

func (EventListArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*EventList)(nil)).Elem()
}

func (o EventListArrayOutput) ToEventListArrayOutput() EventListArrayOutput {
	return o
}

func (o EventListArrayOutput) ToEventListArrayOutputWithContext(ctx context.Context) EventListArrayOutput {
	return o
}

func (o EventListArrayOutput) Index(i pulumi.IntInput) EventListOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *EventList {
		return vs[0].([]*EventList)[vs[1].(int)]
	}).(EventListOutput)
}

type EventListMapOutput struct{ *pulumi.OutputState }

func (EventListMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*EventList)(nil)).Elem()
}

func (o EventListMapOutput) ToEventListMapOutput() EventListMapOutput {
	return o
}

func (o EventListMapOutput) ToEventListMapOutputWithContext(ctx context.Context) EventListMapOutput {
	return o
}

func (o EventListMapOutput) MapIndex(k pulumi.StringInput) EventListOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *EventList {
		return vs[0].(map[string]*EventList)[vs[1].(string)]
	}).(EventListOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*EventListInput)(nil)).Elem(), &EventList{})
	pulumi.RegisterInputType(reflect.TypeOf((*EventListArrayInput)(nil)).Elem(), EventListArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*EventListMapInput)(nil)).Elem(), EventListMap{})
	pulumi.RegisterOutputType(EventListOutput{})
	pulumi.RegisterOutputType(EventListArrayOutput{})
	pulumi.RegisterOutputType(EventListMapOutput{})
}
