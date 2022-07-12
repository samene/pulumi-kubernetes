// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Kubernetes.Types.Inputs.Core.V1
{

    public class PortStatusPatchArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Error is to record the problem with the service port The format of the error shall comply with the following rules: - built-in error values shall be specified in this file and those shall use
        ///   CamelCase names
        /// - cloud provider specific error values must have names that comply with the
        ///   format foo.example.com/CamelCase.
        /// </summary>
        [Input("error")]
        public Input<string>? Error { get; set; }

        /// <summary>
        /// Port is the port number of the service port of which status is recorded here
        /// </summary>
        [Input("port")]
        public Input<int>? Port { get; set; }

        /// <summary>
        /// Protocol is the protocol of the service port of which status is recorded here The supported values are: "TCP", "UDP", "SCTP"
        /// </summary>
        [Input("protocol")]
        public Input<string>? Protocol { get; set; }

        public PortStatusPatchArgs()
        {
        }
    }
}
