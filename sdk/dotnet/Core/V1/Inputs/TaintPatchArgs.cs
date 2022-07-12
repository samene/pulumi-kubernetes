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
    /// The node this Taint is attached to has the "effect" on any pod that does not tolerate the Taint.
    /// </summary>
    public class TaintPatchArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Required. The effect of the taint on pods that do not tolerate the taint. Valid effects are NoSchedule, PreferNoSchedule and NoExecute.
        /// </summary>
        [Input("effect")]
        public Input<string>? Effect { get; set; }

        /// <summary>
        /// Required. The taint key to be applied to a node.
        /// </summary>
        [Input("key")]
        public Input<string>? Key { get; set; }

        /// <summary>
        /// TimeAdded represents the time at which the taint was added. It is only written for NoExecute taints.
        /// </summary>
        [Input("timeAdded")]
        public Input<string>? TimeAdded { get; set; }

        /// <summary>
        /// The taint value corresponding to the taint key.
        /// </summary>
        [Input("value")]
        public Input<string>? Value { get; set; }

        public TaintPatchArgs()
        {
        }
    }
}
