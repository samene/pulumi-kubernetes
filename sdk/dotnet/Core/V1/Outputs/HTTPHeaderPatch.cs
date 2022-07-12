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
    /// HTTPHeader describes a custom header to be used in HTTP probes
    /// </summary>
    [OutputType]
    public sealed class HTTPHeaderPatch
    {
        /// <summary>
        /// The header field name
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// The header field value
        /// </summary>
        public readonly string Value;

        [OutputConstructor]
        private HTTPHeaderPatch(
            string name,

            string value)
        {
            Name = name;
            Value = value;
        }
    }
}
