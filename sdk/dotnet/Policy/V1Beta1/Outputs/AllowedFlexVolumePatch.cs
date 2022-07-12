// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Kubernetes.Types.Outputs.Policy.V1Beta1
{

    /// <summary>
    /// AllowedFlexVolume represents a single Flexvolume that is allowed to be used.
    /// </summary>
    [OutputType]
    public sealed class AllowedFlexVolumePatch
    {
        /// <summary>
        /// driver is the name of the Flexvolume driver.
        /// </summary>
        public readonly string Driver;

        [OutputConstructor]
        private AllowedFlexVolumePatch(string driver)
        {
            Driver = driver;
        }
    }
}
