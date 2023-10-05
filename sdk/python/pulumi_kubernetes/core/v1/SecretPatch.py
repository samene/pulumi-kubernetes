# coding=utf-8
# *** WARNING: this file was generated by pulumigen. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Callable, Mapping, Optional, Sequence, Union, overload
from ... import _utilities
from ... import meta as _meta

__all__ = ['SecretPatchArgs', 'SecretPatch']

@pulumi.input_type
class SecretPatchArgs:
    def __init__(__self__, *,
                 api_version: Optional[pulumi.Input[str]] = None,
                 data: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 immutable: Optional[pulumi.Input[bool]] = None,
                 kind: Optional[pulumi.Input[str]] = None,
                 metadata: Optional[pulumi.Input['_meta.v1.ObjectMetaPatchArgs']] = None,
                 string_data: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 type: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a SecretPatch resource.
        :param pulumi.Input[str] api_version: APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] data: Data contains the secret data. Each key must consist of alphanumeric characters, '-', '_' or '.'. The serialized form of the secret data is a base64 encoded string, representing the arbitrary (possibly non-string) data value here. Described in https://tools.ietf.org/html/rfc4648#section-4
        :param pulumi.Input[bool] immutable: Immutable, if set to true, ensures that data stored in the Secret cannot be updated (only object metadata can be modified). If not set to true, the field can be modified at any time. Defaulted to nil.
        :param pulumi.Input[str] kind: Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
        :param pulumi.Input['_meta.v1.ObjectMetaPatchArgs'] metadata: Standard object's metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] string_data: stringData allows specifying non-binary secret data in string form. It is provided as a write-only input field for convenience. All keys and values are merged into the data field on write, overwriting any existing values. The stringData field is never output when reading from the API.
        :param pulumi.Input[str] type: Used to facilitate programmatic handling of secret data. More info: https://kubernetes.io/docs/concepts/configuration/secret/#secret-types
        """
        SecretPatchArgs._configure(
            lambda key, value: pulumi.set(__self__, key, value),
            api_version=api_version,
            data=data,
            immutable=immutable,
            kind=kind,
            metadata=metadata,
            string_data=string_data,
            type=type,
        )
    @staticmethod
    def _configure(
             _setter: Callable[[Any, Any], None],
             api_version: Optional[pulumi.Input[str]] = None,
             data: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
             immutable: Optional[pulumi.Input[bool]] = None,
             kind: Optional[pulumi.Input[str]] = None,
             metadata: Optional[pulumi.Input['_meta.v1.ObjectMetaPatchArgs']] = None,
             string_data: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
             type: Optional[pulumi.Input[str]] = None,
             opts: Optional[pulumi.ResourceOptions]=None):
        if api_version is not None:
            _setter("api_version", 'v1')
        if data is not None:
            _setter("data", data)
        if immutable is not None:
            _setter("immutable", immutable)
        if kind is not None:
            _setter("kind", 'Secret')
        if metadata is not None:
            _setter("metadata", metadata)
        if string_data is not None:
            _setter("string_data", string_data)
        if type is not None:
            _setter("type", type)

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
    def data(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]:
        """
        Data contains the secret data. Each key must consist of alphanumeric characters, '-', '_' or '.'. The serialized form of the secret data is a base64 encoded string, representing the arbitrary (possibly non-string) data value here. Described in https://tools.ietf.org/html/rfc4648#section-4
        """
        return pulumi.get(self, "data")

    @data.setter
    def data(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]):
        pulumi.set(self, "data", value)

    @property
    @pulumi.getter
    def immutable(self) -> Optional[pulumi.Input[bool]]:
        """
        Immutable, if set to true, ensures that data stored in the Secret cannot be updated (only object metadata can be modified). If not set to true, the field can be modified at any time. Defaulted to nil.
        """
        return pulumi.get(self, "immutable")

    @immutable.setter
    def immutable(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "immutable", value)

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
    def metadata(self) -> Optional[pulumi.Input['_meta.v1.ObjectMetaPatchArgs']]:
        """
        Standard object's metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
        """
        return pulumi.get(self, "metadata")

    @metadata.setter
    def metadata(self, value: Optional[pulumi.Input['_meta.v1.ObjectMetaPatchArgs']]):
        pulumi.set(self, "metadata", value)

    @property
    @pulumi.getter(name="stringData")
    def string_data(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]:
        """
        stringData allows specifying non-binary secret data in string form. It is provided as a write-only input field for convenience. All keys and values are merged into the data field on write, overwriting any existing values. The stringData field is never output when reading from the API.
        """
        return pulumi.get(self, "string_data")

    @string_data.setter
    def string_data(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]):
        pulumi.set(self, "string_data", value)

    @property
    @pulumi.getter
    def type(self) -> Optional[pulumi.Input[str]]:
        """
        Used to facilitate programmatic handling of secret data. More info: https://kubernetes.io/docs/concepts/configuration/secret/#secret-types
        """
        return pulumi.get(self, "type")

    @type.setter
    def type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "type", value)


class SecretPatch(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 api_version: Optional[pulumi.Input[str]] = None,
                 data: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 immutable: Optional[pulumi.Input[bool]] = None,
                 kind: Optional[pulumi.Input[str]] = None,
                 metadata: Optional[pulumi.Input[pulumi.InputType['_meta.v1.ObjectMetaPatchArgs']]] = None,
                 string_data: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 type: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Patch resources are used to modify existing Kubernetes resources by using
        Server-Side Apply updates. The name of the resource must be specified, but all other properties are optional. More than
        one patch may be applied to the same resource, and a random FieldManager name will be used for each Patch resource.
        Conflicts will result in an error by default, but can be forced using the "pulumi.com/patchForce" annotation. See the
        [Server-Side Apply Docs](https://www.pulumi.com/registry/packages/kubernetes/how-to-guides/managing-resources-with-server-side-apply/) for
        additional information about using Server-Side Apply to manage Kubernetes resources with Pulumi.
        Secret holds secret data of a certain type. The total bytes of the values in the Data field must be less than MaxSecretSize bytes.

        Note: While Pulumi automatically encrypts the 'data' and 'stringData'
        fields, this encryption only applies to Pulumi's context, including the state file,
        the Service, the CLI, etc. Kubernetes does not encrypt Secret resources by default,
        and the contents are visible to users with access to the Secret in Kubernetes using
        tools like 'kubectl'.

        For more information on securing Kubernetes Secrets, see the following links:
        https://kubernetes.io/docs/concepts/configuration/secret/#security-properties
        https://kubernetes.io/docs/concepts/configuration/secret/#risks

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] api_version: APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] data: Data contains the secret data. Each key must consist of alphanumeric characters, '-', '_' or '.'. The serialized form of the secret data is a base64 encoded string, representing the arbitrary (possibly non-string) data value here. Described in https://tools.ietf.org/html/rfc4648#section-4
        :param pulumi.Input[bool] immutable: Immutable, if set to true, ensures that data stored in the Secret cannot be updated (only object metadata can be modified). If not set to true, the field can be modified at any time. Defaulted to nil.
        :param pulumi.Input[str] kind: Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
        :param pulumi.Input[pulumi.InputType['_meta.v1.ObjectMetaPatchArgs']] metadata: Standard object's metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] string_data: stringData allows specifying non-binary secret data in string form. It is provided as a write-only input field for convenience. All keys and values are merged into the data field on write, overwriting any existing values. The stringData field is never output when reading from the API.
        :param pulumi.Input[str] type: Used to facilitate programmatic handling of secret data. More info: https://kubernetes.io/docs/concepts/configuration/secret/#secret-types
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: Optional[SecretPatchArgs] = None,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Patch resources are used to modify existing Kubernetes resources by using
        Server-Side Apply updates. The name of the resource must be specified, but all other properties are optional. More than
        one patch may be applied to the same resource, and a random FieldManager name will be used for each Patch resource.
        Conflicts will result in an error by default, but can be forced using the "pulumi.com/patchForce" annotation. See the
        [Server-Side Apply Docs](https://www.pulumi.com/registry/packages/kubernetes/how-to-guides/managing-resources-with-server-side-apply/) for
        additional information about using Server-Side Apply to manage Kubernetes resources with Pulumi.
        Secret holds secret data of a certain type. The total bytes of the values in the Data field must be less than MaxSecretSize bytes.

        Note: While Pulumi automatically encrypts the 'data' and 'stringData'
        fields, this encryption only applies to Pulumi's context, including the state file,
        the Service, the CLI, etc. Kubernetes does not encrypt Secret resources by default,
        and the contents are visible to users with access to the Secret in Kubernetes using
        tools like 'kubectl'.

        For more information on securing Kubernetes Secrets, see the following links:
        https://kubernetes.io/docs/concepts/configuration/secret/#security-properties
        https://kubernetes.io/docs/concepts/configuration/secret/#risks

        :param str resource_name: The name of the resource.
        :param SecretPatchArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(SecretPatchArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            kwargs = kwargs or {}
            def _setter(key, value):
                kwargs[key] = value
            SecretPatchArgs._configure(_setter, **kwargs)
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 api_version: Optional[pulumi.Input[str]] = None,
                 data: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 immutable: Optional[pulumi.Input[bool]] = None,
                 kind: Optional[pulumi.Input[str]] = None,
                 metadata: Optional[pulumi.Input[pulumi.InputType['_meta.v1.ObjectMetaPatchArgs']]] = None,
                 string_data: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 type: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = SecretPatchArgs.__new__(SecretPatchArgs)

            __props__.__dict__["api_version"] = 'v1'
            __props__.__dict__["data"] = data
            __props__.__dict__["immutable"] = immutable
            __props__.__dict__["kind"] = 'Secret'
            if metadata is not None and not isinstance(metadata, _meta.v1.ObjectMetaPatchArgs):
                metadata = metadata or {}
                def _setter(key, value):
                    metadata[key] = value
                _meta.v1.ObjectMetaPatchArgs._configure(_setter, **metadata)
            __props__.__dict__["metadata"] = metadata
            __props__.__dict__["string_data"] = string_data
            __props__.__dict__["type"] = type
        super(SecretPatch, __self__).__init__(
            'kubernetes:core/v1:SecretPatch',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'SecretPatch':
        """
        Get an existing SecretPatch resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = SecretPatchArgs.__new__(SecretPatchArgs)

        __props__.__dict__["api_version"] = None
        __props__.__dict__["data"] = None
        __props__.__dict__["immutable"] = None
        __props__.__dict__["kind"] = None
        __props__.__dict__["metadata"] = None
        __props__.__dict__["string_data"] = None
        __props__.__dict__["type"] = None
        return SecretPatch(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="apiVersion")
    def api_version(self) -> pulumi.Output[Optional[str]]:
        """
        APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
        """
        return pulumi.get(self, "api_version")

    @property
    @pulumi.getter
    def data(self) -> pulumi.Output[Optional[Mapping[str, str]]]:
        """
        Data contains the secret data. Each key must consist of alphanumeric characters, '-', '_' or '.'. The serialized form of the secret data is a base64 encoded string, representing the arbitrary (possibly non-string) data value here. Described in https://tools.ietf.org/html/rfc4648#section-4
        """
        return pulumi.get(self, "data")

    @property
    @pulumi.getter
    def immutable(self) -> pulumi.Output[Optional[bool]]:
        """
        Immutable, if set to true, ensures that data stored in the Secret cannot be updated (only object metadata can be modified). If not set to true, the field can be modified at any time. Defaulted to nil.
        """
        return pulumi.get(self, "immutable")

    @property
    @pulumi.getter
    def kind(self) -> pulumi.Output[Optional[str]]:
        """
        Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
        """
        return pulumi.get(self, "kind")

    @property
    @pulumi.getter
    def metadata(self) -> pulumi.Output[Optional['_meta.v1.outputs.ObjectMetaPatch']]:
        """
        Standard object's metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
        """
        return pulumi.get(self, "metadata")

    @property
    @pulumi.getter(name="stringData")
    def string_data(self) -> pulumi.Output[Optional[Mapping[str, str]]]:
        """
        stringData allows specifying non-binary secret data in string form. It is provided as a write-only input field for convenience. All keys and values are merged into the data field on write, overwriting any existing values. The stringData field is never output when reading from the API.
        """
        return pulumi.get(self, "string_data")

    @property
    @pulumi.getter
    def type(self) -> pulumi.Output[Optional[str]]:
        """
        Used to facilitate programmatic handling of secret data. More info: https://kubernetes.io/docs/concepts/configuration/secret/#secret-types
        """
        return pulumi.get(self, "type")

