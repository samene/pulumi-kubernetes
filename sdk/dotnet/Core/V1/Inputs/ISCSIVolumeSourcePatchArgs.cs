// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Kubernetes.Types.Inputs.Core.V1
{

    /// <summary>
    /// Represents an ISCSI disk. ISCSI volumes can only be mounted as read/write once. ISCSI volumes support ownership management and SELinux relabeling.
    /// </summary>
    public class ISCSIVolumeSourcePatchArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// chapAuthDiscovery defines whether support iSCSI Discovery CHAP authentication
        /// </summary>
        [Input("chapAuthDiscovery")]
        public Input<bool>? ChapAuthDiscovery { get; set; }

        /// <summary>
        /// chapAuthSession defines whether support iSCSI Session CHAP authentication
        /// </summary>
        [Input("chapAuthSession")]
        public Input<bool>? ChapAuthSession { get; set; }

        /// <summary>
        /// fsType is the filesystem type of the volume that you want to mount. Tip: Ensure that the filesystem type is supported by the host operating system. Examples: "ext4", "xfs", "ntfs". Implicitly inferred to be "ext4" if unspecified. More info: https://kubernetes.io/docs/concepts/storage/volumes#iscsi
        /// </summary>
        [Input("fsType")]
        public Input<string>? FsType { get; set; }

        /// <summary>
        /// initiatorName is the custom iSCSI Initiator Name. If initiatorName is specified with iscsiInterface simultaneously, new iSCSI interface &lt;target portal&gt;:&lt;volume name&gt; will be created for the connection.
        /// </summary>
        [Input("initiatorName")]
        public Input<string>? InitiatorName { get; set; }

        /// <summary>
        /// iqn is the target iSCSI Qualified Name.
        /// </summary>
        [Input("iqn")]
        public Input<string>? Iqn { get; set; }

        /// <summary>
        /// iscsiInterface is the interface Name that uses an iSCSI transport. Defaults to 'default' (tcp).
        /// </summary>
        [Input("iscsiInterface")]
        public Input<string>? IscsiInterface { get; set; }

        /// <summary>
        /// lun represents iSCSI Target Lun number.
        /// </summary>
        [Input("lun")]
        public Input<int>? Lun { get; set; }

        [Input("portals")]
        private InputList<string>? _portals;

        /// <summary>
        /// portals is the iSCSI Target Portal List. The portal is either an IP or ip_addr:port if the port is other than default (typically TCP ports 860 and 3260).
        /// </summary>
        public InputList<string> Portals
        {
            get => _portals ?? (_portals = new InputList<string>());
            set => _portals = value;
        }

        /// <summary>
        /// readOnly here will force the ReadOnly setting in VolumeMounts. Defaults to false.
        /// </summary>
        [Input("readOnly")]
        public Input<bool>? ReadOnly { get; set; }

        /// <summary>
        /// secretRef is the CHAP Secret for iSCSI target and initiator authentication
        /// </summary>
        [Input("secretRef")]
        public Input<Pulumi.Kubernetes.Types.Inputs.Core.V1.LocalObjectReferencePatchArgs>? SecretRef { get; set; }

        /// <summary>
        /// targetPortal is iSCSI Target Portal. The Portal is either an IP or ip_addr:port if the port is other than default (typically TCP ports 860 and 3260).
        /// </summary>
        [Input("targetPortal")]
        public Input<string>? TargetPortal { get; set; }

        public ISCSIVolumeSourcePatchArgs()
        {
        }
    }
}
