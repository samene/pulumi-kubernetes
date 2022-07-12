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
    /// PersistentVolumeClaimVolumeSource references the user's PVC in the same namespace. This volume finds the bound PV and mounts that volume for the pod. A PersistentVolumeClaimVolumeSource is, essentially, a wrapper around another type of volume that is owned by someone else (the system).
    /// </summary>
    [OutputType]
    public sealed class PersistentVolumeClaimVolumeSourcePatch
    {
        /// <summary>
        /// claimName is the name of a PersistentVolumeClaim in the same namespace as the pod using this volume. More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#persistentvolumeclaims
        /// </summary>
        public readonly string ClaimName;
        /// <summary>
        /// readOnly Will force the ReadOnly setting in VolumeMounts. Default false.
        /// </summary>
        public readonly bool ReadOnly;

        [OutputConstructor]
        private PersistentVolumeClaimVolumeSourcePatch(
            string claimName,

            bool readOnly)
        {
            ClaimName = claimName;
            ReadOnly = readOnly;
        }
    }
}
