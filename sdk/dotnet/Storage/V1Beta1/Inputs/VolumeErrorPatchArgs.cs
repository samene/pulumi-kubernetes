// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Kubernetes.Types.Inputs.Storage.V1Beta1
{

    /// <summary>
    /// VolumeError captures an error encountered during a volume operation.
    /// </summary>
    public class VolumeErrorPatchArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// String detailing the error encountered during Attach or Detach operation. This string may be logged, so it should not contain sensitive information.
        /// </summary>
        [Input("message")]
        public Input<string>? Message { get; set; }

        /// <summary>
        /// Time the error was encountered.
        /// </summary>
        [Input("time")]
        public Input<string>? Time { get; set; }

        public VolumeErrorPatchArgs()
        {
        }
    }
}
