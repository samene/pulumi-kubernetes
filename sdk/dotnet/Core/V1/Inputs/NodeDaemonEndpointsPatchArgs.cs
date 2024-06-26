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
    /// NodeDaemonEndpoints lists ports opened by daemons running on the Node.
    /// </summary>
    public class NodeDaemonEndpointsPatchArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Endpoint on which Kubelet is listening.
        /// </summary>
        [Input("kubeletEndpoint")]
        public Input<Pulumi.Kubernetes.Types.Inputs.Core.V1.DaemonEndpointPatchArgs>? KubeletEndpoint { get; set; }

        public NodeDaemonEndpointsPatchArgs()
        {
        }
    }
}
