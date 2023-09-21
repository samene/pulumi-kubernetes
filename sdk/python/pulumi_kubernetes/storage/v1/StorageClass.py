# coding=utf-8
# *** WARNING: this file was generated by pulumigen. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Callable, Mapping, Optional, Sequence, Union, overload
from ... import _utilities
from ... import core as _core
from ... import meta as _meta

__all__ = ['StorageClassInitArgs', 'StorageClass']

@pulumi.input_type
class StorageClassInitArgs:
    def __init__(__self__, *,
                 provisioner: pulumi.Input[str],
                 allow_volume_expansion: Optional[pulumi.Input[bool]] = None,
                 allowed_topologies: Optional[pulumi.Input[Sequence[pulumi.Input['_core.v1.TopologySelectorTermArgs']]]] = None,
                 api_version: Optional[pulumi.Input[str]] = None,
                 kind: Optional[pulumi.Input[str]] = None,
                 metadata: Optional[pulumi.Input['_meta.v1.ObjectMetaArgs']] = None,
                 mount_options: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 parameters: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 reclaim_policy: Optional[pulumi.Input[str]] = None,
                 volume_binding_mode: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a StorageClass resource.
        :param pulumi.Input[str] provisioner: provisioner indicates the type of the provisioner.
        :param pulumi.Input[bool] allow_volume_expansion: allowVolumeExpansion shows whether the storage class allow volume expand.
        :param pulumi.Input[Sequence[pulumi.Input['_core.v1.TopologySelectorTermArgs']]] allowed_topologies: allowedTopologies restrict the node topologies where volumes can be dynamically provisioned. Each volume plugin defines its own supported topology specifications. An empty TopologySelectorTerm list means there is no topology restriction. This field is only honored by servers that enable the VolumeScheduling feature.
        :param pulumi.Input[str] api_version: APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
        :param pulumi.Input[str] kind: Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
        :param pulumi.Input['_meta.v1.ObjectMetaArgs'] metadata: Standard object's metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
        :param pulumi.Input[Sequence[pulumi.Input[str]]] mount_options: mountOptions controls the mountOptions for dynamically provisioned PersistentVolumes of this storage class. e.g. ["ro", "soft"]. Not validated - mount of the PVs will simply fail if one is invalid.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] parameters: parameters holds the parameters for the provisioner that should create volumes of this storage class.
        :param pulumi.Input[str] reclaim_policy: reclaimPolicy controls the reclaimPolicy for dynamically provisioned PersistentVolumes of this storage class. Defaults to Delete.
        :param pulumi.Input[str] volume_binding_mode: volumeBindingMode indicates how PersistentVolumeClaims should be provisioned and bound.  When unset, VolumeBindingImmediate is used. This field is only honored by servers that enable the VolumeScheduling feature.
        """
        StorageClassInitArgs._configure(
            lambda key, value: pulumi.set(__self__, key, value),
            provisioner=provisioner,
            allow_volume_expansion=allow_volume_expansion,
            allowed_topologies=allowed_topologies,
            api_version=api_version,
            kind=kind,
            metadata=metadata,
            mount_options=mount_options,
            parameters=parameters,
            reclaim_policy=reclaim_policy,
            volume_binding_mode=volume_binding_mode,
        )
    @staticmethod
    def _configure(
             _setter: Callable[[Any, Any], None],
             provisioner: pulumi.Input[str],
             allow_volume_expansion: Optional[pulumi.Input[bool]] = None,
             allowed_topologies: Optional[pulumi.Input[Sequence[pulumi.Input['_core.v1.TopologySelectorTermArgs']]]] = None,
             api_version: Optional[pulumi.Input[str]] = None,
             kind: Optional[pulumi.Input[str]] = None,
             metadata: Optional[pulumi.Input['_meta.v1.ObjectMetaArgs']] = None,
             mount_options: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
             parameters: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
             reclaim_policy: Optional[pulumi.Input[str]] = None,
             volume_binding_mode: Optional[pulumi.Input[str]] = None,
             opts: Optional[pulumi.ResourceOptions]=None):
        _setter("provisioner", provisioner)
        if allow_volume_expansion is not None:
            _setter("allow_volume_expansion", allow_volume_expansion)
        if allowed_topologies is not None:
            _setter("allowed_topologies", allowed_topologies)
        if api_version is not None:
            _setter("api_version", 'storage.k8s.io/v1')
        if kind is not None:
            _setter("kind", 'StorageClass')
        if metadata is not None:
            _setter("metadata", metadata)
        if mount_options is not None:
            _setter("mount_options", mount_options)
        if parameters is not None:
            _setter("parameters", parameters)
        if reclaim_policy is not None:
            _setter("reclaim_policy", reclaim_policy)
        if volume_binding_mode is not None:
            _setter("volume_binding_mode", volume_binding_mode)

    @property
    @pulumi.getter
    def provisioner(self) -> pulumi.Input[str]:
        """
        provisioner indicates the type of the provisioner.
        """
        return pulumi.get(self, "provisioner")

    @provisioner.setter
    def provisioner(self, value: pulumi.Input[str]):
        pulumi.set(self, "provisioner", value)

    @property
    @pulumi.getter(name="allowVolumeExpansion")
    def allow_volume_expansion(self) -> Optional[pulumi.Input[bool]]:
        """
        allowVolumeExpansion shows whether the storage class allow volume expand.
        """
        return pulumi.get(self, "allow_volume_expansion")

    @allow_volume_expansion.setter
    def allow_volume_expansion(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "allow_volume_expansion", value)

    @property
    @pulumi.getter(name="allowedTopologies")
    def allowed_topologies(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['_core.v1.TopologySelectorTermArgs']]]]:
        """
        allowedTopologies restrict the node topologies where volumes can be dynamically provisioned. Each volume plugin defines its own supported topology specifications. An empty TopologySelectorTerm list means there is no topology restriction. This field is only honored by servers that enable the VolumeScheduling feature.
        """
        return pulumi.get(self, "allowed_topologies")

    @allowed_topologies.setter
    def allowed_topologies(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['_core.v1.TopologySelectorTermArgs']]]]):
        pulumi.set(self, "allowed_topologies", value)

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
    def metadata(self) -> Optional[pulumi.Input['_meta.v1.ObjectMetaArgs']]:
        """
        Standard object's metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
        """
        return pulumi.get(self, "metadata")

    @metadata.setter
    def metadata(self, value: Optional[pulumi.Input['_meta.v1.ObjectMetaArgs']]):
        pulumi.set(self, "metadata", value)

    @property
    @pulumi.getter(name="mountOptions")
    def mount_options(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        mountOptions controls the mountOptions for dynamically provisioned PersistentVolumes of this storage class. e.g. ["ro", "soft"]. Not validated - mount of the PVs will simply fail if one is invalid.
        """
        return pulumi.get(self, "mount_options")

    @mount_options.setter
    def mount_options(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "mount_options", value)

    @property
    @pulumi.getter
    def parameters(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]:
        """
        parameters holds the parameters for the provisioner that should create volumes of this storage class.
        """
        return pulumi.get(self, "parameters")

    @parameters.setter
    def parameters(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]):
        pulumi.set(self, "parameters", value)

    @property
    @pulumi.getter(name="reclaimPolicy")
    def reclaim_policy(self) -> Optional[pulumi.Input[str]]:
        """
        reclaimPolicy controls the reclaimPolicy for dynamically provisioned PersistentVolumes of this storage class. Defaults to Delete.
        """
        return pulumi.get(self, "reclaim_policy")

    @reclaim_policy.setter
    def reclaim_policy(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "reclaim_policy", value)

    @property
    @pulumi.getter(name="volumeBindingMode")
    def volume_binding_mode(self) -> Optional[pulumi.Input[str]]:
        """
        volumeBindingMode indicates how PersistentVolumeClaims should be provisioned and bound.  When unset, VolumeBindingImmediate is used. This field is only honored by servers that enable the VolumeScheduling feature.
        """
        return pulumi.get(self, "volume_binding_mode")

    @volume_binding_mode.setter
    def volume_binding_mode(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "volume_binding_mode", value)


class StorageClass(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 allow_volume_expansion: Optional[pulumi.Input[bool]] = None,
                 allowed_topologies: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['_core.v1.TopologySelectorTermArgs']]]]] = None,
                 api_version: Optional[pulumi.Input[str]] = None,
                 kind: Optional[pulumi.Input[str]] = None,
                 metadata: Optional[pulumi.Input[pulumi.InputType['_meta.v1.ObjectMetaArgs']]] = None,
                 mount_options: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 parameters: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 provisioner: Optional[pulumi.Input[str]] = None,
                 reclaim_policy: Optional[pulumi.Input[str]] = None,
                 volume_binding_mode: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        StorageClass describes the parameters for a class of storage for which PersistentVolumes can be dynamically provisioned.

        StorageClasses are non-namespaced; the name of the storage class according to etcd is in ObjectMeta.Name.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[bool] allow_volume_expansion: allowVolumeExpansion shows whether the storage class allow volume expand.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['_core.v1.TopologySelectorTermArgs']]]] allowed_topologies: allowedTopologies restrict the node topologies where volumes can be dynamically provisioned. Each volume plugin defines its own supported topology specifications. An empty TopologySelectorTerm list means there is no topology restriction. This field is only honored by servers that enable the VolumeScheduling feature.
        :param pulumi.Input[str] api_version: APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
        :param pulumi.Input[str] kind: Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
        :param pulumi.Input[pulumi.InputType['_meta.v1.ObjectMetaArgs']] metadata: Standard object's metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
        :param pulumi.Input[Sequence[pulumi.Input[str]]] mount_options: mountOptions controls the mountOptions for dynamically provisioned PersistentVolumes of this storage class. e.g. ["ro", "soft"]. Not validated - mount of the PVs will simply fail if one is invalid.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] parameters: parameters holds the parameters for the provisioner that should create volumes of this storage class.
        :param pulumi.Input[str] provisioner: provisioner indicates the type of the provisioner.
        :param pulumi.Input[str] reclaim_policy: reclaimPolicy controls the reclaimPolicy for dynamically provisioned PersistentVolumes of this storage class. Defaults to Delete.
        :param pulumi.Input[str] volume_binding_mode: volumeBindingMode indicates how PersistentVolumeClaims should be provisioned and bound.  When unset, VolumeBindingImmediate is used. This field is only honored by servers that enable the VolumeScheduling feature.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: StorageClassInitArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        StorageClass describes the parameters for a class of storage for which PersistentVolumes can be dynamically provisioned.

        StorageClasses are non-namespaced; the name of the storage class according to etcd is in ObjectMeta.Name.

        :param str resource_name: The name of the resource.
        :param StorageClassInitArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(StorageClassInitArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            kwargs = kwargs or {}
            def _setter(key, value):
                kwargs[key] = value
            StorageClassInitArgs._configure(_setter, **kwargs)
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 allow_volume_expansion: Optional[pulumi.Input[bool]] = None,
                 allowed_topologies: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['_core.v1.TopologySelectorTermArgs']]]]] = None,
                 api_version: Optional[pulumi.Input[str]] = None,
                 kind: Optional[pulumi.Input[str]] = None,
                 metadata: Optional[pulumi.Input[pulumi.InputType['_meta.v1.ObjectMetaArgs']]] = None,
                 mount_options: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 parameters: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 provisioner: Optional[pulumi.Input[str]] = None,
                 reclaim_policy: Optional[pulumi.Input[str]] = None,
                 volume_binding_mode: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = StorageClassInitArgs.__new__(StorageClassInitArgs)

            __props__.__dict__["allow_volume_expansion"] = allow_volume_expansion
            __props__.__dict__["allowed_topologies"] = allowed_topologies
            __props__.__dict__["api_version"] = 'storage.k8s.io/v1'
            __props__.__dict__["kind"] = 'StorageClass'
            if not isinstance(metadata, _meta.v1.ObjectMetaArgs):
                metadata = metadata or {}
                def _setter(key, value):
                    metadata[key] = value
                _meta.v1.ObjectMetaArgs._configure(_setter, **metadata)
            __props__.__dict__["metadata"] = metadata
            __props__.__dict__["mount_options"] = mount_options
            __props__.__dict__["parameters"] = parameters
            if provisioner is None and not opts.urn:
                raise TypeError("Missing required property 'provisioner'")
            __props__.__dict__["provisioner"] = provisioner
            __props__.__dict__["reclaim_policy"] = reclaim_policy
            __props__.__dict__["volume_binding_mode"] = volume_binding_mode
        alias_opts = pulumi.ResourceOptions(aliases=[pulumi.Alias(type_="kubernetes:storage.k8s.io/v1beta1:StorageClass")])
        opts = pulumi.ResourceOptions.merge(opts, alias_opts)
        super(StorageClass, __self__).__init__(
            'kubernetes:storage.k8s.io/v1:StorageClass',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'StorageClass':
        """
        Get an existing StorageClass resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = StorageClassInitArgs.__new__(StorageClassInitArgs)

        __props__.__dict__["allow_volume_expansion"] = None
        __props__.__dict__["allowed_topologies"] = None
        __props__.__dict__["api_version"] = None
        __props__.__dict__["kind"] = None
        __props__.__dict__["metadata"] = None
        __props__.__dict__["mount_options"] = None
        __props__.__dict__["parameters"] = None
        __props__.__dict__["provisioner"] = None
        __props__.__dict__["reclaim_policy"] = None
        __props__.__dict__["volume_binding_mode"] = None
        return StorageClass(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="allowVolumeExpansion")
    def allow_volume_expansion(self) -> pulumi.Output[bool]:
        """
        allowVolumeExpansion shows whether the storage class allow volume expand.
        """
        return pulumi.get(self, "allow_volume_expansion")

    @property
    @pulumi.getter(name="allowedTopologies")
    def allowed_topologies(self) -> pulumi.Output[Sequence['_core.v1.outputs.TopologySelectorTerm']]:
        """
        allowedTopologies restrict the node topologies where volumes can be dynamically provisioned. Each volume plugin defines its own supported topology specifications. An empty TopologySelectorTerm list means there is no topology restriction. This field is only honored by servers that enable the VolumeScheduling feature.
        """
        return pulumi.get(self, "allowed_topologies")

    @property
    @pulumi.getter(name="apiVersion")
    def api_version(self) -> pulumi.Output[str]:
        """
        APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
        """
        return pulumi.get(self, "api_version")

    @property
    @pulumi.getter
    def kind(self) -> pulumi.Output[str]:
        """
        Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
        """
        return pulumi.get(self, "kind")

    @property
    @pulumi.getter
    def metadata(self) -> pulumi.Output['_meta.v1.outputs.ObjectMeta']:
        """
        Standard object's metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
        """
        return pulumi.get(self, "metadata")

    @property
    @pulumi.getter(name="mountOptions")
    def mount_options(self) -> pulumi.Output[Sequence[str]]:
        """
        mountOptions controls the mountOptions for dynamically provisioned PersistentVolumes of this storage class. e.g. ["ro", "soft"]. Not validated - mount of the PVs will simply fail if one is invalid.
        """
        return pulumi.get(self, "mount_options")

    @property
    @pulumi.getter
    def parameters(self) -> pulumi.Output[Mapping[str, str]]:
        """
        parameters holds the parameters for the provisioner that should create volumes of this storage class.
        """
        return pulumi.get(self, "parameters")

    @property
    @pulumi.getter
    def provisioner(self) -> pulumi.Output[str]:
        """
        provisioner indicates the type of the provisioner.
        """
        return pulumi.get(self, "provisioner")

    @property
    @pulumi.getter(name="reclaimPolicy")
    def reclaim_policy(self) -> pulumi.Output[str]:
        """
        reclaimPolicy controls the reclaimPolicy for dynamically provisioned PersistentVolumes of this storage class. Defaults to Delete.
        """
        return pulumi.get(self, "reclaim_policy")

    @property
    @pulumi.getter(name="volumeBindingMode")
    def volume_binding_mode(self) -> pulumi.Output[str]:
        """
        volumeBindingMode indicates how PersistentVolumeClaims should be provisioned and bound.  When unset, VolumeBindingImmediate is used. This field is only honored by servers that enable the VolumeScheduling feature.
        """
        return pulumi.get(self, "volume_binding_mode")

