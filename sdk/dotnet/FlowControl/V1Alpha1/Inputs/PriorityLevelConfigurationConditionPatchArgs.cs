// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Kubernetes.Types.Inputs.FlowControl.V1Alpha1
{

    /// <summary>
    /// PriorityLevelConfigurationCondition defines the condition of priority level.
    /// </summary>
    public class PriorityLevelConfigurationConditionPatchArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// `lastTransitionTime` is the last time the condition transitioned from one status to another.
        /// </summary>
        [Input("lastTransitionTime")]
        public Input<string>? LastTransitionTime { get; set; }

        /// <summary>
        /// `message` is a human-readable message indicating details about last transition.
        /// </summary>
        [Input("message")]
        public Input<string>? Message { get; set; }

        /// <summary>
        /// `reason` is a unique, one-word, CamelCase reason for the condition's last transition.
        /// </summary>
        [Input("reason")]
        public Input<string>? Reason { get; set; }

        /// <summary>
        /// `status` is the status of the condition. Can be True, False, Unknown. Required.
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        /// <summary>
        /// `type` is the type of the condition. Required.
        /// </summary>
        [Input("type")]
        public Input<string>? Type { get; set; }

        public PriorityLevelConfigurationConditionPatchArgs()
        {
        }
    }
}
