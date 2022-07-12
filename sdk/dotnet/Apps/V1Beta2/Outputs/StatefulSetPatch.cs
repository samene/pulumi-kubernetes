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
    /// StatefulSet represents a set of pods with consistent identities. Identities are defined as:
    ///  - Network: A single stable DNS and hostname.
    ///  - Storage: As many VolumeClaims as requested.
    /// The StatefulSet guarantees that a given network identity will always map to the same storage identity.
    /// 
    /// This resource waits until its status is ready before registering success
    /// for create/update, and populating output properties from the current state of the resource.
    /// The following conditions are used to determine whether the resource creation has
    /// succeeded or failed:
    /// 
    /// 1. The value of 'spec.replicas' matches '.status.replicas', '.status.currentReplicas',
    ///    and '.status.readyReplicas'.
    /// 2. The value of '.status.updateRevision' matches '.status.currentRevision'.
    /// 
    /// If the StatefulSet has not reached a Ready state after 10 minutes, it will
    /// time out and mark the resource update as Failed. You can override the default timeout value
    /// by setting the 'customTimeouts' option on the resource.
    /// </summary>
    [OutputType]
    public sealed class StatefulSetPatch
    {
        /// <summary>
        /// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
        /// </summary>
        public readonly string ApiVersion;
        /// <summary>
        /// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
        /// </summary>
        public readonly string Kind;
        public readonly Pulumi.Kubernetes.Types.Outputs.Meta.V1.ObjectMetaPatch Metadata;
        /// <summary>
        /// Spec defines the desired identities of pods in this set.
        /// </summary>
        public readonly Pulumi.Kubernetes.Types.Outputs.Apps.V1Beta2.StatefulSetSpecPatch Spec;
        /// <summary>
        /// Status is the current status of Pods in this StatefulSet. This data may be out of date by some window of time.
        /// </summary>
        public readonly Pulumi.Kubernetes.Types.Outputs.Apps.V1Beta2.StatefulSetStatusPatch Status;

        [OutputConstructor]
        private StatefulSetPatch(
            string apiVersion,

            string kind,

            Pulumi.Kubernetes.Types.Outputs.Meta.V1.ObjectMetaPatch metadata,

            Pulumi.Kubernetes.Types.Outputs.Apps.V1Beta2.StatefulSetSpecPatch spec,

            Pulumi.Kubernetes.Types.Outputs.Apps.V1Beta2.StatefulSetStatusPatch status)
        {
            ApiVersion = apiVersion;
            Kind = kind;
            Metadata = metadata;
            Spec = spec;
            Status = status;
        }
    }
}
