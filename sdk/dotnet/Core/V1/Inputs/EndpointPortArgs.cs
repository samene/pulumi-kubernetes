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
    /// EndpointPort is a tuple that describes a single port.
    /// </summary>
    public class EndpointPortArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The application protocol for this port. This is used as a hint for implementations to offer richer behavior for protocols that they understand. This field follows standard Kubernetes label syntax. Valid values are either:
        /// 
        /// * Un-prefixed protocol names - reserved for IANA standard service names (as per RFC-6335 and https://www.iana.org/assignments/service-names).
        /// 
        /// * Kubernetes-defined prefixed names:
        ///   * 'kubernetes.io/h2c' - HTTP/2 prior knowledge over cleartext as described in https://www.rfc-editor.org/rfc/rfc9113.html#name-starting-http-2-with-prior-
        ///   * 'kubernetes.io/ws'  - WebSocket over cleartext as described in https://www.rfc-editor.org/rfc/rfc6455
        ///   * 'kubernetes.io/wss' - WebSocket over TLS as described in https://www.rfc-editor.org/rfc/rfc6455
        /// 
        /// * Other protocols should use implementation-defined prefixed names such as mycompany.com/my-custom-protocol.
        /// </summary>
        [Input("appProtocol")]
        public Input<string>? AppProtocol { get; set; }

        /// <summary>
        /// The name of this port.  This must match the 'name' field in the corresponding ServicePort. Must be a DNS_LABEL. Optional only if one port is defined.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The port number of the endpoint.
        /// </summary>
        [Input("port", required: true)]
        public Input<int> Port { get; set; } = null!;

        /// <summary>
        /// The IP protocol for this port. Must be UDP, TCP, or SCTP. Default is TCP.
        /// </summary>
        [Input("protocol")]
        public Input<string>? Protocol { get; set; }

        public EndpointPortArgs()
        {
        }
        public static new EndpointPortArgs Empty => new EndpointPortArgs();
    }
}
