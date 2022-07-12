// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Kubernetes.Types.Inputs.Autoscaling.V1
{

    /// <summary>
    /// CrossVersionObjectReference contains enough information to let you identify the referred resource.
    /// </summary>
    public class CrossVersionObjectReferencePatchArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// API version of the referent
        /// </summary>
        [Input("apiVersion")]
        public Input<string>? ApiVersion { get; set; }

        /// <summary>
        /// Kind of the referent; More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds"
        /// </summary>
        [Input("kind")]
        public Input<string>? Kind { get; set; }

        /// <summary>
        /// Name of the referent; More info: http://kubernetes.io/docs/user-guide/identifiers#names
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        public CrossVersionObjectReferencePatchArgs()
        {
        }
    }
}
