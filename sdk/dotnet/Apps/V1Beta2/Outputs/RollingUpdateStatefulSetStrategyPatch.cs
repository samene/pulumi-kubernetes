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
    /// RollingUpdateStatefulSetStrategy is used to communicate parameter for RollingUpdateStatefulSetStrategyType.
    /// </summary>
    [OutputType]
    public sealed class RollingUpdateStatefulSetStrategyPatch
    {
        /// <summary>
        /// Partition indicates the ordinal at which the StatefulSet should be partitioned. Default value is 0.
        /// </summary>
        public readonly int Partition;

        [OutputConstructor]
        private RollingUpdateStatefulSetStrategyPatch(int partition)
        {
            Partition = partition;
        }
    }
}
