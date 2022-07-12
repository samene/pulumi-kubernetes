// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Kubernetes.Types.Outputs.Certificates.V1Beta1
{

    [OutputType]
    public sealed class CertificateSigningRequestStatusPatch
    {
        /// <summary>
        /// If request was approved, the controller will place the issued certificate here.
        /// </summary>
        public readonly string Certificate;
        /// <summary>
        /// Conditions applied to the request, such as approval or denial.
        /// </summary>
        public readonly ImmutableArray<Pulumi.Kubernetes.Types.Outputs.Certificates.V1Beta1.CertificateSigningRequestConditionPatch> Conditions;

        [OutputConstructor]
        private CertificateSigningRequestStatusPatch(
            string certificate,

            ImmutableArray<Pulumi.Kubernetes.Types.Outputs.Certificates.V1Beta1.CertificateSigningRequestConditionPatch> conditions)
        {
            Certificate = certificate;
            Conditions = conditions;
        }
    }
}
