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
    /// PodReadinessGate contains the reference to a pod condition
    /// </summary>
    public class PodReadinessGatePatchArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// ConditionType refers to a condition in the pod's condition list with matching type.
        /// </summary>
        [Input("conditionType")]
        public Input<string>? ConditionType { get; set; }

        public PodReadinessGatePatchArgs()
        {
        }
        public static new PodReadinessGatePatchArgs Empty => new PodReadinessGatePatchArgs();
    }
}
