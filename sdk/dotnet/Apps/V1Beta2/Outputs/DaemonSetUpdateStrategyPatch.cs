// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Kubernetes.Types.Outputs.Apps.V1Beta2
{

    /// <summary>
    /// DaemonSetUpdateStrategy is a struct used to control the update strategy for a DaemonSet.
    /// </summary>
    [OutputType]
    public sealed class DaemonSetUpdateStrategyPatch
    {
        /// <summary>
        /// Rolling update config params. Present only if type = "RollingUpdate".
        /// </summary>
        public readonly Pulumi.Kubernetes.Types.Outputs.Apps.V1Beta2.RollingUpdateDaemonSetPatch RollingUpdate;
        /// <summary>
        /// Type of daemon set update. Can be "RollingUpdate" or "OnDelete". Default is RollingUpdate.
        /// </summary>
        public readonly string Type;

        [OutputConstructor]
        private DaemonSetUpdateStrategyPatch(
            Pulumi.Kubernetes.Types.Outputs.Apps.V1Beta2.RollingUpdateDaemonSetPatch rollingUpdate,

            string type)
        {
            RollingUpdate = rollingUpdate;
            Type = type;
        }
    }
}
