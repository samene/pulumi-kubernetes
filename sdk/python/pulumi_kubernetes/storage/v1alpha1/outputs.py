# coding=utf-8
# *** WARNING: this file was generated by pulumigen. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from ... import _utilities
from . import outputs
from ... import core as _core
from ... import meta as _meta

__all__ = [
    'VolumeAttachment',
    'VolumeAttachmentSource',
    'VolumeAttachmentSourcePatch',
    'VolumeAttachmentSpec',
    'VolumeAttachmentSpecPatch',
    'VolumeAttachmentStatus',
    'VolumeAttachmentStatusPatch',
    'VolumeAttributesClass',
    'VolumeError',
    'VolumeErrorPatch',
]

@pulumi.output_type
class VolumeAttachment(dict):
    """
    VolumeAttachment captures the intent to attach or detach the specified volume to/from the specified node.

    VolumeAttachment objects are non-namespaced.
    """
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "apiVersion":
            suggest = "api_version"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in VolumeAttachment. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        VolumeAttachment.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        VolumeAttachment.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 spec: 'outputs.VolumeAttachmentSpec',
                 api_version: Optional[str] = None,
                 kind: Optional[str] = None,
                 metadata: Optional['_meta.v1.outputs.ObjectMeta'] = None,
                 status: Optional['outputs.VolumeAttachmentStatus'] = None):
        """
        VolumeAttachment captures the intent to attach or detach the specified volume to/from the specified node.

        VolumeAttachment objects are non-namespaced.
        :param 'VolumeAttachmentSpecArgs' spec: Specification of the desired attach/detach volume behavior. Populated by the Kubernetes system.
        :param str api_version: APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
        :param str kind: Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
        :param '_meta.v1.ObjectMetaArgs' metadata: Standard object metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
        :param 'VolumeAttachmentStatusArgs' status: Status of the VolumeAttachment request. Populated by the entity completing the attach or detach operation, i.e. the external-attacher.
        """
        pulumi.set(__self__, "spec", spec)
        if api_version is not None:
            pulumi.set(__self__, "api_version", 'storage.k8s.io/v1alpha1')
        if kind is not None:
            pulumi.set(__self__, "kind", 'VolumeAttachment')
        if metadata is not None:
            pulumi.set(__self__, "metadata", metadata)
        if status is not None:
            pulumi.set(__self__, "status", status)

    @property
    @pulumi.getter
    def spec(self) -> 'outputs.VolumeAttachmentSpec':
        """
        Specification of the desired attach/detach volume behavior. Populated by the Kubernetes system.
        """
        return pulumi.get(self, "spec")

    @property
    @pulumi.getter(name="apiVersion")
    def api_version(self) -> Optional[str]:
        """
        APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
        """
        return pulumi.get(self, "api_version")

    @property
    @pulumi.getter
    def kind(self) -> Optional[str]:
        """
        Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
        """
        return pulumi.get(self, "kind")

    @property
    @pulumi.getter
    def metadata(self) -> Optional['_meta.v1.outputs.ObjectMeta']:
        """
        Standard object metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
        """
        return pulumi.get(self, "metadata")

    @property
    @pulumi.getter
    def status(self) -> Optional['outputs.VolumeAttachmentStatus']:
        """
        Status of the VolumeAttachment request. Populated by the entity completing the attach or detach operation, i.e. the external-attacher.
        """
        return pulumi.get(self, "status")


@pulumi.output_type
class VolumeAttachmentSource(dict):
    """
    VolumeAttachmentSource represents a volume that should be attached. Right now only PersistenVolumes can be attached via external attacher, in future we may allow also inline volumes in pods. Exactly one member can be set.
    """
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "inlineVolumeSpec":
            suggest = "inline_volume_spec"
        elif key == "persistentVolumeName":
            suggest = "persistent_volume_name"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in VolumeAttachmentSource. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        VolumeAttachmentSource.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        VolumeAttachmentSource.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 inline_volume_spec: Optional['_core.v1.outputs.PersistentVolumeSpec'] = None,
                 persistent_volume_name: Optional[str] = None):
        """
        VolumeAttachmentSource represents a volume that should be attached. Right now only PersistenVolumes can be attached via external attacher, in future we may allow also inline volumes in pods. Exactly one member can be set.
        :param '_core.v1.PersistentVolumeSpecArgs' inline_volume_spec: inlineVolumeSpec contains all the information necessary to attach a persistent volume defined by a pod's inline VolumeSource. This field is populated only for the CSIMigration feature. It contains translated fields from a pod's inline VolumeSource to a PersistentVolumeSpec. This field is alpha-level and is only honored by servers that enabled the CSIMigration feature.
        :param str persistent_volume_name: Name of the persistent volume to attach.
        """
        if inline_volume_spec is not None:
            pulumi.set(__self__, "inline_volume_spec", inline_volume_spec)
        if persistent_volume_name is not None:
            pulumi.set(__self__, "persistent_volume_name", persistent_volume_name)

    @property
    @pulumi.getter(name="inlineVolumeSpec")
    def inline_volume_spec(self) -> Optional['_core.v1.outputs.PersistentVolumeSpec']:
        """
        inlineVolumeSpec contains all the information necessary to attach a persistent volume defined by a pod's inline VolumeSource. This field is populated only for the CSIMigration feature. It contains translated fields from a pod's inline VolumeSource to a PersistentVolumeSpec. This field is alpha-level and is only honored by servers that enabled the CSIMigration feature.
        """
        return pulumi.get(self, "inline_volume_spec")

    @property
    @pulumi.getter(name="persistentVolumeName")
    def persistent_volume_name(self) -> Optional[str]:
        """
        Name of the persistent volume to attach.
        """
        return pulumi.get(self, "persistent_volume_name")


@pulumi.output_type
class VolumeAttachmentSourcePatch(dict):
    """
    VolumeAttachmentSource represents a volume that should be attached. Right now only PersistenVolumes can be attached via external attacher, in future we may allow also inline volumes in pods. Exactly one member can be set.
    """
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "inlineVolumeSpec":
            suggest = "inline_volume_spec"
        elif key == "persistentVolumeName":
            suggest = "persistent_volume_name"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in VolumeAttachmentSourcePatch. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        VolumeAttachmentSourcePatch.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        VolumeAttachmentSourcePatch.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 inline_volume_spec: Optional['_core.v1.outputs.PersistentVolumeSpecPatch'] = None,
                 persistent_volume_name: Optional[str] = None):
        """
        VolumeAttachmentSource represents a volume that should be attached. Right now only PersistenVolumes can be attached via external attacher, in future we may allow also inline volumes in pods. Exactly one member can be set.
        :param '_core.v1.PersistentVolumeSpecPatchArgs' inline_volume_spec: inlineVolumeSpec contains all the information necessary to attach a persistent volume defined by a pod's inline VolumeSource. This field is populated only for the CSIMigration feature. It contains translated fields from a pod's inline VolumeSource to a PersistentVolumeSpec. This field is alpha-level and is only honored by servers that enabled the CSIMigration feature.
        :param str persistent_volume_name: Name of the persistent volume to attach.
        """
        if inline_volume_spec is not None:
            pulumi.set(__self__, "inline_volume_spec", inline_volume_spec)
        if persistent_volume_name is not None:
            pulumi.set(__self__, "persistent_volume_name", persistent_volume_name)

    @property
    @pulumi.getter(name="inlineVolumeSpec")
    def inline_volume_spec(self) -> Optional['_core.v1.outputs.PersistentVolumeSpecPatch']:
        """
        inlineVolumeSpec contains all the information necessary to attach a persistent volume defined by a pod's inline VolumeSource. This field is populated only for the CSIMigration feature. It contains translated fields from a pod's inline VolumeSource to a PersistentVolumeSpec. This field is alpha-level and is only honored by servers that enabled the CSIMigration feature.
        """
        return pulumi.get(self, "inline_volume_spec")

    @property
    @pulumi.getter(name="persistentVolumeName")
    def persistent_volume_name(self) -> Optional[str]:
        """
        Name of the persistent volume to attach.
        """
        return pulumi.get(self, "persistent_volume_name")


@pulumi.output_type
class VolumeAttachmentSpec(dict):
    """
    VolumeAttachmentSpec is the specification of a VolumeAttachment request.
    """
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "nodeName":
            suggest = "node_name"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in VolumeAttachmentSpec. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        VolumeAttachmentSpec.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        VolumeAttachmentSpec.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 attacher: str,
                 node_name: str,
                 source: 'outputs.VolumeAttachmentSource'):
        """
        VolumeAttachmentSpec is the specification of a VolumeAttachment request.
        :param str attacher: Attacher indicates the name of the volume driver that MUST handle this request. This is the name returned by GetPluginName().
        :param str node_name: The node that the volume should be attached to.
        :param 'VolumeAttachmentSourceArgs' source: Source represents the volume that should be attached.
        """
        pulumi.set(__self__, "attacher", attacher)
        pulumi.set(__self__, "node_name", node_name)
        pulumi.set(__self__, "source", source)

    @property
    @pulumi.getter
    def attacher(self) -> str:
        """
        Attacher indicates the name of the volume driver that MUST handle this request. This is the name returned by GetPluginName().
        """
        return pulumi.get(self, "attacher")

    @property
    @pulumi.getter(name="nodeName")
    def node_name(self) -> str:
        """
        The node that the volume should be attached to.
        """
        return pulumi.get(self, "node_name")

    @property
    @pulumi.getter
    def source(self) -> 'outputs.VolumeAttachmentSource':
        """
        Source represents the volume that should be attached.
        """
        return pulumi.get(self, "source")


@pulumi.output_type
class VolumeAttachmentSpecPatch(dict):
    """
    VolumeAttachmentSpec is the specification of a VolumeAttachment request.
    """
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "nodeName":
            suggest = "node_name"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in VolumeAttachmentSpecPatch. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        VolumeAttachmentSpecPatch.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        VolumeAttachmentSpecPatch.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 attacher: Optional[str] = None,
                 node_name: Optional[str] = None,
                 source: Optional['outputs.VolumeAttachmentSourcePatch'] = None):
        """
        VolumeAttachmentSpec is the specification of a VolumeAttachment request.
        :param str attacher: Attacher indicates the name of the volume driver that MUST handle this request. This is the name returned by GetPluginName().
        :param str node_name: The node that the volume should be attached to.
        :param 'VolumeAttachmentSourcePatchArgs' source: Source represents the volume that should be attached.
        """
        if attacher is not None:
            pulumi.set(__self__, "attacher", attacher)
        if node_name is not None:
            pulumi.set(__self__, "node_name", node_name)
        if source is not None:
            pulumi.set(__self__, "source", source)

    @property
    @pulumi.getter
    def attacher(self) -> Optional[str]:
        """
        Attacher indicates the name of the volume driver that MUST handle this request. This is the name returned by GetPluginName().
        """
        return pulumi.get(self, "attacher")

    @property
    @pulumi.getter(name="nodeName")
    def node_name(self) -> Optional[str]:
        """
        The node that the volume should be attached to.
        """
        return pulumi.get(self, "node_name")

    @property
    @pulumi.getter
    def source(self) -> Optional['outputs.VolumeAttachmentSourcePatch']:
        """
        Source represents the volume that should be attached.
        """
        return pulumi.get(self, "source")


@pulumi.output_type
class VolumeAttachmentStatus(dict):
    """
    VolumeAttachmentStatus is the status of a VolumeAttachment request.
    """
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "attachError":
            suggest = "attach_error"
        elif key == "attachmentMetadata":
            suggest = "attachment_metadata"
        elif key == "detachError":
            suggest = "detach_error"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in VolumeAttachmentStatus. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        VolumeAttachmentStatus.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        VolumeAttachmentStatus.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 attached: bool,
                 attach_error: Optional['outputs.VolumeError'] = None,
                 attachment_metadata: Optional[Mapping[str, str]] = None,
                 detach_error: Optional['outputs.VolumeError'] = None):
        """
        VolumeAttachmentStatus is the status of a VolumeAttachment request.
        :param bool attached: Indicates the volume is successfully attached. This field must only be set by the entity completing the attach operation, i.e. the external-attacher.
        :param 'VolumeErrorArgs' attach_error: The last error encountered during attach operation, if any. This field must only be set by the entity completing the attach operation, i.e. the external-attacher.
        :param Mapping[str, str] attachment_metadata: Upon successful attach, this field is populated with any information returned by the attach operation that must be passed into subsequent WaitForAttach or Mount calls. This field must only be set by the entity completing the attach operation, i.e. the external-attacher.
        :param 'VolumeErrorArgs' detach_error: The last error encountered during detach operation, if any. This field must only be set by the entity completing the detach operation, i.e. the external-attacher.
        """
        pulumi.set(__self__, "attached", attached)
        if attach_error is not None:
            pulumi.set(__self__, "attach_error", attach_error)
        if attachment_metadata is not None:
            pulumi.set(__self__, "attachment_metadata", attachment_metadata)
        if detach_error is not None:
            pulumi.set(__self__, "detach_error", detach_error)

    @property
    @pulumi.getter
    def attached(self) -> bool:
        """
        Indicates the volume is successfully attached. This field must only be set by the entity completing the attach operation, i.e. the external-attacher.
        """
        return pulumi.get(self, "attached")

    @property
    @pulumi.getter(name="attachError")
    def attach_error(self) -> Optional['outputs.VolumeError']:
        """
        The last error encountered during attach operation, if any. This field must only be set by the entity completing the attach operation, i.e. the external-attacher.
        """
        return pulumi.get(self, "attach_error")

    @property
    @pulumi.getter(name="attachmentMetadata")
    def attachment_metadata(self) -> Optional[Mapping[str, str]]:
        """
        Upon successful attach, this field is populated with any information returned by the attach operation that must be passed into subsequent WaitForAttach or Mount calls. This field must only be set by the entity completing the attach operation, i.e. the external-attacher.
        """
        return pulumi.get(self, "attachment_metadata")

    @property
    @pulumi.getter(name="detachError")
    def detach_error(self) -> Optional['outputs.VolumeError']:
        """
        The last error encountered during detach operation, if any. This field must only be set by the entity completing the detach operation, i.e. the external-attacher.
        """
        return pulumi.get(self, "detach_error")


@pulumi.output_type
class VolumeAttachmentStatusPatch(dict):
    """
    VolumeAttachmentStatus is the status of a VolumeAttachment request.
    """
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "attachError":
            suggest = "attach_error"
        elif key == "attachmentMetadata":
            suggest = "attachment_metadata"
        elif key == "detachError":
            suggest = "detach_error"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in VolumeAttachmentStatusPatch. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        VolumeAttachmentStatusPatch.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        VolumeAttachmentStatusPatch.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 attach_error: Optional['outputs.VolumeErrorPatch'] = None,
                 attached: Optional[bool] = None,
                 attachment_metadata: Optional[Mapping[str, str]] = None,
                 detach_error: Optional['outputs.VolumeErrorPatch'] = None):
        """
        VolumeAttachmentStatus is the status of a VolumeAttachment request.
        :param 'VolumeErrorPatchArgs' attach_error: The last error encountered during attach operation, if any. This field must only be set by the entity completing the attach operation, i.e. the external-attacher.
        :param bool attached: Indicates the volume is successfully attached. This field must only be set by the entity completing the attach operation, i.e. the external-attacher.
        :param Mapping[str, str] attachment_metadata: Upon successful attach, this field is populated with any information returned by the attach operation that must be passed into subsequent WaitForAttach or Mount calls. This field must only be set by the entity completing the attach operation, i.e. the external-attacher.
        :param 'VolumeErrorPatchArgs' detach_error: The last error encountered during detach operation, if any. This field must only be set by the entity completing the detach operation, i.e. the external-attacher.
        """
        if attach_error is not None:
            pulumi.set(__self__, "attach_error", attach_error)
        if attached is not None:
            pulumi.set(__self__, "attached", attached)
        if attachment_metadata is not None:
            pulumi.set(__self__, "attachment_metadata", attachment_metadata)
        if detach_error is not None:
            pulumi.set(__self__, "detach_error", detach_error)

    @property
    @pulumi.getter(name="attachError")
    def attach_error(self) -> Optional['outputs.VolumeErrorPatch']:
        """
        The last error encountered during attach operation, if any. This field must only be set by the entity completing the attach operation, i.e. the external-attacher.
        """
        return pulumi.get(self, "attach_error")

    @property
    @pulumi.getter
    def attached(self) -> Optional[bool]:
        """
        Indicates the volume is successfully attached. This field must only be set by the entity completing the attach operation, i.e. the external-attacher.
        """
        return pulumi.get(self, "attached")

    @property
    @pulumi.getter(name="attachmentMetadata")
    def attachment_metadata(self) -> Optional[Mapping[str, str]]:
        """
        Upon successful attach, this field is populated with any information returned by the attach operation that must be passed into subsequent WaitForAttach or Mount calls. This field must only be set by the entity completing the attach operation, i.e. the external-attacher.
        """
        return pulumi.get(self, "attachment_metadata")

    @property
    @pulumi.getter(name="detachError")
    def detach_error(self) -> Optional['outputs.VolumeErrorPatch']:
        """
        The last error encountered during detach operation, if any. This field must only be set by the entity completing the detach operation, i.e. the external-attacher.
        """
        return pulumi.get(self, "detach_error")


@pulumi.output_type
class VolumeAttributesClass(dict):
    """
    VolumeAttributesClass represents a specification of mutable volume attributes defined by the CSI driver. The class can be specified during dynamic provisioning of PersistentVolumeClaims, and changed in the PersistentVolumeClaim spec after provisioning.
    """
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "driverName":
            suggest = "driver_name"
        elif key == "apiVersion":
            suggest = "api_version"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in VolumeAttributesClass. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        VolumeAttributesClass.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        VolumeAttributesClass.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 driver_name: str,
                 api_version: Optional[str] = None,
                 kind: Optional[str] = None,
                 metadata: Optional['_meta.v1.outputs.ObjectMeta'] = None,
                 parameters: Optional[Mapping[str, str]] = None):
        """
        VolumeAttributesClass represents a specification of mutable volume attributes defined by the CSI driver. The class can be specified during dynamic provisioning of PersistentVolumeClaims, and changed in the PersistentVolumeClaim spec after provisioning.
        :param str driver_name: Name of the CSI driver This field is immutable.
        :param str api_version: APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
        :param str kind: Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
        :param '_meta.v1.ObjectMetaArgs' metadata: Standard object's metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
        :param Mapping[str, str] parameters: parameters hold volume attributes defined by the CSI driver. These values are opaque to the Kubernetes and are passed directly to the CSI driver. The underlying storage provider supports changing these attributes on an existing volume, however the parameters field itself is immutable. To invoke a volume update, a new VolumeAttributesClass should be created with new parameters, and the PersistentVolumeClaim should be updated to reference the new VolumeAttributesClass.
               
               This field is required and must contain at least one key/value pair. The keys cannot be empty, and the maximum number of parameters is 512, with a cumulative max size of 256K. If the CSI driver rejects invalid parameters, the target PersistentVolumeClaim will be set to an "Infeasible" state in the modifyVolumeStatus field.
        """
        pulumi.set(__self__, "driver_name", driver_name)
        if api_version is not None:
            pulumi.set(__self__, "api_version", 'storage.k8s.io/v1alpha1')
        if kind is not None:
            pulumi.set(__self__, "kind", 'VolumeAttributesClass')
        if metadata is not None:
            pulumi.set(__self__, "metadata", metadata)
        if parameters is not None:
            pulumi.set(__self__, "parameters", parameters)

    @property
    @pulumi.getter(name="driverName")
    def driver_name(self) -> str:
        """
        Name of the CSI driver This field is immutable.
        """
        return pulumi.get(self, "driver_name")

    @property
    @pulumi.getter(name="apiVersion")
    def api_version(self) -> Optional[str]:
        """
        APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
        """
        return pulumi.get(self, "api_version")

    @property
    @pulumi.getter
    def kind(self) -> Optional[str]:
        """
        Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
        """
        return pulumi.get(self, "kind")

    @property
    @pulumi.getter
    def metadata(self) -> Optional['_meta.v1.outputs.ObjectMeta']:
        """
        Standard object's metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
        """
        return pulumi.get(self, "metadata")

    @property
    @pulumi.getter
    def parameters(self) -> Optional[Mapping[str, str]]:
        """
        parameters hold volume attributes defined by the CSI driver. These values are opaque to the Kubernetes and are passed directly to the CSI driver. The underlying storage provider supports changing these attributes on an existing volume, however the parameters field itself is immutable. To invoke a volume update, a new VolumeAttributesClass should be created with new parameters, and the PersistentVolumeClaim should be updated to reference the new VolumeAttributesClass.

        This field is required and must contain at least one key/value pair. The keys cannot be empty, and the maximum number of parameters is 512, with a cumulative max size of 256K. If the CSI driver rejects invalid parameters, the target PersistentVolumeClaim will be set to an "Infeasible" state in the modifyVolumeStatus field.
        """
        return pulumi.get(self, "parameters")


@pulumi.output_type
class VolumeError(dict):
    """
    VolumeError captures an error encountered during a volume operation.
    """
    def __init__(__self__, *,
                 message: Optional[str] = None,
                 time: Optional[str] = None):
        """
        VolumeError captures an error encountered during a volume operation.
        :param str message: String detailing the error encountered during Attach or Detach operation. This string maybe logged, so it should not contain sensitive information.
        :param str time: Time the error was encountered.
        """
        if message is not None:
            pulumi.set(__self__, "message", message)
        if time is not None:
            pulumi.set(__self__, "time", time)

    @property
    @pulumi.getter
    def message(self) -> Optional[str]:
        """
        String detailing the error encountered during Attach or Detach operation. This string maybe logged, so it should not contain sensitive information.
        """
        return pulumi.get(self, "message")

    @property
    @pulumi.getter
    def time(self) -> Optional[str]:
        """
        Time the error was encountered.
        """
        return pulumi.get(self, "time")


@pulumi.output_type
class VolumeErrorPatch(dict):
    """
    VolumeError captures an error encountered during a volume operation.
    """
    def __init__(__self__, *,
                 message: Optional[str] = None,
                 time: Optional[str] = None):
        """
        VolumeError captures an error encountered during a volume operation.
        :param str message: String detailing the error encountered during Attach or Detach operation. This string maybe logged, so it should not contain sensitive information.
        :param str time: Time the error was encountered.
        """
        if message is not None:
            pulumi.set(__self__, "message", message)
        if time is not None:
            pulumi.set(__self__, "time", time)

    @property
    @pulumi.getter
    def message(self) -> Optional[str]:
        """
        String detailing the error encountered during Attach or Detach operation. This string maybe logged, so it should not contain sensitive information.
        """
        return pulumi.get(self, "message")

    @property
    @pulumi.getter
    def time(self) -> Optional[str]:
        """
        Time the error was encountered.
        """
        return pulumi.get(self, "time")


