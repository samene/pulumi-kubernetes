// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Kubernetes.Types.Outputs.Core.V1
{

    /// <summary>
    /// EndpointAddress is a tuple that describes single IP address.
    /// </summary>
    [OutputType]
    public sealed class EndpointAddressPatch
    {
        /// <summary>
        /// The Hostname of this endpoint
        /// </summary>
        public readonly string Hostname;
        /// <summary>
        /// The IP of this endpoint. May not be loopback (127.0.0.0/8), link-local (169.254.0.0/16), or link-local multicast ((224.0.0.0/24). IPv6 is also accepted but not fully supported on all platforms. Also, certain kubernetes components, like kube-proxy, are not IPv6 ready.
        /// </summary>
        public readonly string Ip;
        /// <summary>
        /// Optional: Node hosting this endpoint. This can be used to determine endpoints local to a node.
        /// </summary>
        public readonly string NodeName;
        /// <summary>
        /// Reference to object providing the endpoint.
        /// </summary>
        public readonly Pulumi.Kubernetes.Types.Outputs.Core.V1.ObjectReferencePatch TargetRef;

        [OutputConstructor]
        private EndpointAddressPatch(
            string hostname,

            string ip,

            string nodeName,

            Pulumi.Kubernetes.Types.Outputs.Core.V1.ObjectReferencePatch targetRef)
        {
            Hostname = hostname;
            Ip = ip;
            NodeName = nodeName;
            TargetRef = targetRef;
        }
    }
}
