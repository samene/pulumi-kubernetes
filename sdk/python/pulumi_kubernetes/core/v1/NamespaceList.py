# coding=utf-8
# *** WARNING: this file was generated by pulumigen. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Callable, Mapping, Optional, Sequence, Union, overload
from ... import _utilities
from . import outputs
from ... import meta as _meta
from ._inputs import *

__all__ = ['NamespaceListArgs', 'NamespaceList']

@pulumi.input_type
class NamespaceListArgs:
    def __init__(__self__, *,
                 items: pulumi.Input[Sequence[pulumi.Input['NamespaceArgs']]],
                 api_version: Optional[pulumi.Input[str]] = None,
                 kind: Optional[pulumi.Input[str]] = None,
                 metadata: Optional[pulumi.Input['_meta.v1.ListMetaArgs']] = None):
        """
        The set of arguments for constructing a NamespaceList resource.
        :param pulumi.Input[Sequence[pulumi.Input['NamespaceArgs']]] items: Items is the list of Namespace objects in the list. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/namespaces/
        :param pulumi.Input[str] api_version: APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
        :param pulumi.Input[str] kind: Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
        :param pulumi.Input['_meta.v1.ListMetaArgs'] metadata: Standard list metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
        """
        NamespaceListArgs._configure(
            lambda key, value: pulumi.set(__self__, key, value),
            items=items,
            api_version=api_version,
            kind=kind,
            metadata=metadata,
        )
    @staticmethod
    def _configure(
             _setter: Callable[[Any, Any], None],
             items: pulumi.Input[Sequence[pulumi.Input['NamespaceArgs']]],
             api_version: Optional[pulumi.Input[str]] = None,
             kind: Optional[pulumi.Input[str]] = None,
             metadata: Optional[pulumi.Input['_meta.v1.ListMetaArgs']] = None,
             opts: Optional[pulumi.ResourceOptions]=None):
        _setter("items", items)
        if api_version is not None:
            _setter("api_version", 'v1')
        if kind is not None:
            _setter("kind", 'NamespaceList')
        if metadata is not None:
            _setter("metadata", metadata)

    @property
    @pulumi.getter
    def items(self) -> pulumi.Input[Sequence[pulumi.Input['NamespaceArgs']]]:
        """
        Items is the list of Namespace objects in the list. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/namespaces/
        """
        return pulumi.get(self, "items")

    @items.setter
    def items(self, value: pulumi.Input[Sequence[pulumi.Input['NamespaceArgs']]]):
        pulumi.set(self, "items", value)

    @property
    @pulumi.getter(name="apiVersion")
    def api_version(self) -> Optional[pulumi.Input[str]]:
        """
        APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
        """
        return pulumi.get(self, "api_version")

    @api_version.setter
    def api_version(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "api_version", value)

    @property
    @pulumi.getter
    def kind(self) -> Optional[pulumi.Input[str]]:
        """
        Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
        """
        return pulumi.get(self, "kind")

    @kind.setter
    def kind(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "kind", value)

    @property
    @pulumi.getter
    def metadata(self) -> Optional[pulumi.Input['_meta.v1.ListMetaArgs']]:
        """
        Standard list metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
        """
        return pulumi.get(self, "metadata")

    @metadata.setter
    def metadata(self, value: Optional[pulumi.Input['_meta.v1.ListMetaArgs']]):
        pulumi.set(self, "metadata", value)


class NamespaceList(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 api_version: Optional[pulumi.Input[str]] = None,
                 items: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['NamespaceArgs']]]]] = None,
                 kind: Optional[pulumi.Input[str]] = None,
                 metadata: Optional[pulumi.Input[pulumi.InputType['_meta.v1.ListMetaArgs']]] = None,
                 __props__=None):
        """
        NamespaceList is a list of Namespaces.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] api_version: APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['NamespaceArgs']]]] items: Items is the list of Namespace objects in the list. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/namespaces/
        :param pulumi.Input[str] kind: Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
        :param pulumi.Input[pulumi.InputType['_meta.v1.ListMetaArgs']] metadata: Standard list metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: NamespaceListArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        NamespaceList is a list of Namespaces.

        :param str resource_name: The name of the resource.
        :param NamespaceListArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(NamespaceListArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            kwargs = kwargs or {}
            def _setter(key, value):
                kwargs[key] = value
            NamespaceListArgs._configure(_setter, **kwargs)
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 api_version: Optional[pulumi.Input[str]] = None,
                 items: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['NamespaceArgs']]]]] = None,
                 kind: Optional[pulumi.Input[str]] = None,
                 metadata: Optional[pulumi.Input[pulumi.InputType['_meta.v1.ListMetaArgs']]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = NamespaceListArgs.__new__(NamespaceListArgs)

            __props__.__dict__["api_version"] = 'v1'
            if items is None and not opts.urn:
                raise TypeError("Missing required property 'items'")
            __props__.__dict__["items"] = items
            __props__.__dict__["kind"] = 'NamespaceList'
            if metadata is not None and not isinstance(metadata, _meta.v1.ListMetaArgs):
                metadata = metadata or {}
                def _setter(key, value):
                    metadata[key] = value
                _meta.v1.ListMetaArgs._configure(_setter, **metadata)
            __props__.__dict__["metadata"] = metadata
        super(NamespaceList, __self__).__init__(
            'kubernetes:core/v1:NamespaceList',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'NamespaceList':
        """
        Get an existing NamespaceList resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = NamespaceListArgs.__new__(NamespaceListArgs)

        __props__.__dict__["api_version"] = None
        __props__.__dict__["items"] = None
        __props__.__dict__["kind"] = None
        __props__.__dict__["metadata"] = None
        return NamespaceList(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="apiVersion")
    def api_version(self) -> pulumi.Output[str]:
        """
        APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
        """
        return pulumi.get(self, "api_version")

    @property
    @pulumi.getter
    def items(self) -> pulumi.Output[Sequence['outputs.Namespace']]:
        """
        Items is the list of Namespace objects in the list. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/namespaces/
        """
        return pulumi.get(self, "items")

    @property
    @pulumi.getter
    def kind(self) -> pulumi.Output[str]:
        """
        Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
        """
        return pulumi.get(self, "kind")

    @property
    @pulumi.getter
    def metadata(self) -> pulumi.Output['_meta.v1.outputs.ListMeta']:
        """
        Standard list metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
        """
        return pulumi.get(self, "metadata")

