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
    /// Represents a host path mapped into a pod. Host path volumes do not support ownership management or SELinux relabeling.
    /// </summary>
    public class HostPathVolumeSourceArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Path of the directory on the host. If the path is a symlink, it will follow the link to the real path. More info: https://kubernetes.io/docs/concepts/storage/volumes#hostpath
        /// </summary>
        [Input("path", required: true)]
        public Input<string> Path { get; set; } = null!;

        /// <summary>
        /// Type for HostPath Volume Defaults to "" More info: https://kubernetes.io/docs/concepts/storage/volumes#hostpath
        /// </summary>
        [Input("type")]
        public Input<string>? Type { get; set; }

        public HostPathVolumeSourceArgs()
        {
        }
    }
}