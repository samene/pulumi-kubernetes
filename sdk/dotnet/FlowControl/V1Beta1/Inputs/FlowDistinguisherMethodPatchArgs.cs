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
    /// FlowDistinguisherMethod specifies the method of a flow distinguisher.
    /// </summary>
    public class FlowDistinguisherMethodPatchArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// `type` is the type of flow distinguisher method The supported types are "ByUser" and "ByNamespace". Required.
        /// </summary>
        [Input("type")]
        public Input<string>? Type { get; set; }

        public FlowDistinguisherMethodPatchArgs()
        {
        }
    }
}
