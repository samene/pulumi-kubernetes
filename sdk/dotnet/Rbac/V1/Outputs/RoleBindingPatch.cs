// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Kubernetes.Types.Outputs.Rbac.V1
{

    /// <summary>
    /// RoleBinding references a role, but does not contain it.  It can reference a Role in the same namespace or a ClusterRole in the global namespace. It adds who information via Subjects and namespace information by which namespace it exists in.  RoleBindings in a given namespace only have effect in that namespace.
    /// </summary>
    [OutputType]
    public sealed class RoleBindingPatch
    {
        /// <summary>
        /// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
        /// </summary>
        public readonly string ApiVersion;
        /// <summary>
        /// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
        /// </summary>
        public readonly string Kind;
        /// <summary>
        /// Standard object's metadata.
        /// </summary>
        public readonly Pulumi.Kubernetes.Types.Outputs.Meta.V1.ObjectMetaPatch Metadata;
        /// <summary>
        /// RoleRef can reference a Role in the current namespace or a ClusterRole in the global namespace. If the RoleRef cannot be resolved, the Authorizer must return an error.
        /// </summary>
        public readonly Pulumi.Kubernetes.Types.Outputs.Rbac.V1.RoleRefPatch RoleRef;
        /// <summary>
        /// Subjects holds references to the objects the role applies to.
        /// </summary>
        public readonly ImmutableArray<Pulumi.Kubernetes.Types.Outputs.Rbac.V1.SubjectPatch> Subjects;

        [OutputConstructor]
        private RoleBindingPatch(
            string apiVersion,

            string kind,

            Pulumi.Kubernetes.Types.Outputs.Meta.V1.ObjectMetaPatch metadata,

            Pulumi.Kubernetes.Types.Outputs.Rbac.V1.RoleRefPatch roleRef,

            ImmutableArray<Pulumi.Kubernetes.Types.Outputs.Rbac.V1.SubjectPatch> subjects)
        {
            ApiVersion = apiVersion;
            Kind = kind;
            Metadata = metadata;
            RoleRef = roleRef;
            Subjects = subjects;
        }
    }
}
