// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Kubernetes.Types.Outputs.Apps.V1
{

    /// <summary>
    /// DeploymentStrategy describes how to replace existing pods with new ones.
    /// </summary>
    [OutputType]
    public sealed class DeploymentStrategyPatch
    {
        /// <summary>
        /// Rolling update config params. Present only if DeploymentStrategyType = RollingUpdate.
        /// </summary>
        public readonly Pulumi.Kubernetes.Types.Outputs.Apps.V1.RollingUpdateDeploymentPatch RollingUpdate;
        /// <summary>
        /// Type of deployment. Can be "Recreate" or "RollingUpdate". Default is RollingUpdate.
        /// </summary>
        public readonly string Type;

        [OutputConstructor]
        private DeploymentStrategyPatch(
            Pulumi.Kubernetes.Types.Outputs.Apps.V1.RollingUpdateDeploymentPatch rollingUpdate,

            string type)
        {
            RollingUpdate = rollingUpdate;
            Type = type;
        }
    }
}
