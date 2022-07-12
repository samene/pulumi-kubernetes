// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Kubernetes.Types.Inputs.FlowControl.V1Beta1
{

    /// <summary>
    /// LimitResponse defines how to handle requests that can not be executed right now.
    /// </summary>
    public class LimitResponsePatchArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// `queuing` holds the configuration parameters for queuing. This field may be non-empty only if `type` is `"Queue"`.
        /// </summary>
        [Input("queuing")]
        public Input<Pulumi.Kubernetes.Types.Inputs.FlowControl.V1Beta1.QueuingConfigurationPatchArgs>? Queuing { get; set; }

        /// <summary>
        /// `type` is "Queue" or "Reject". "Queue" means that requests that can not be executed upon arrival are held in a queue until they can be executed or a queuing limit is reached. "Reject" means that requests that can not be executed upon arrival are rejected. Required.
        /// </summary>
        [Input("type")]
        public Input<string>? Type { get; set; }

        public LimitResponsePatchArgs()
        {
        }
    }
}
