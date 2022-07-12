// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Kubernetes.Types.Outputs.FlowControl.V1Alpha1
{

    /// <summary>
    /// ServiceAccountSubject holds detailed information for service-account-kind subject.
    /// </summary>
    [OutputType]
    public sealed class ServiceAccountSubjectPatch
    {
        /// <summary>
        /// `name` is the name of matching ServiceAccount objects, or "*" to match regardless of name. Required.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// `namespace` is the namespace of matching ServiceAccount objects. Required.
        /// </summary>
        public readonly string Namespace;

        [OutputConstructor]
        private ServiceAccountSubjectPatch(
            string name,

            string @namespace)
        {
            Name = name;
            Namespace = @namespace;
        }
    }
}
