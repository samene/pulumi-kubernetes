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
    /// NodeCondition contains condition information for a node.
    /// </summary>
    [OutputType]
    public sealed class NodeConditionPatch
    {
        /// <summary>
        /// Last time we got an update on a given condition.
        /// </summary>
        public readonly string LastHeartbeatTime;
        /// <summary>
        /// Last time the condition transit from one status to another.
        /// </summary>
        public readonly string LastTransitionTime;
        /// <summary>
        /// Human readable message indicating details about last transition.
        /// </summary>
        public readonly string Message;
        /// <summary>
        /// (brief) reason for the condition's last transition.
        /// </summary>
        public readonly string Reason;
        /// <summary>
        /// Status of the condition, one of True, False, Unknown.
        /// </summary>
        public readonly string Status;
        /// <summary>
        /// Type of node condition.
        /// </summary>
        public readonly string Type;

        [OutputConstructor]
        private NodeConditionPatch(
            string lastHeartbeatTime,

            string lastTransitionTime,

            string message,

            string reason,

            string status,

            string type)
        {
            LastHeartbeatTime = lastHeartbeatTime;
            LastTransitionTime = lastTransitionTime;
            Message = message;
            Reason = reason;
            Status = status;
            Type = type;
        }
    }
}
