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
    /// HTTPGetAction describes an action based on HTTP Get requests.
    /// </summary>
    [OutputType]
    public sealed class HTTPGetActionPatch
    {
        /// <summary>
        /// Host name to connect to, defaults to the pod IP. You probably want to set "Host" in httpHeaders instead.
        /// </summary>
        public readonly string Host;
        /// <summary>
        /// Custom headers to set in the request. HTTP allows repeated headers.
        /// </summary>
        public readonly ImmutableArray<Pulumi.Kubernetes.Types.Outputs.Core.V1.HTTPHeaderPatch> HttpHeaders;
        /// <summary>
        /// Path to access on the HTTP server.
        /// </summary>
        public readonly string Path;
        /// <summary>
        /// Name or number of the port to access on the container. Number must be in the range 1 to 65535. Name must be an IANA_SVC_NAME.
        /// </summary>
        public readonly Union<int, string> Port;
        /// <summary>
        /// Scheme to use for connecting to the host. Defaults to HTTP.
        /// </summary>
        public readonly string Scheme;

        [OutputConstructor]
        private HTTPGetActionPatch(
            string host,

            ImmutableArray<Pulumi.Kubernetes.Types.Outputs.Core.V1.HTTPHeaderPatch> httpHeaders,

            string path,

            Union<int, string> port,

            string scheme)
        {
            Host = host;
            HttpHeaders = httpHeaders;
            Path = path;
            Port = port;
            Scheme = scheme;
        }
    }
}
