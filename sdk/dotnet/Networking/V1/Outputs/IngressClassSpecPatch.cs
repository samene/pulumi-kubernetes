// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Kubernetes.Types.Outputs.Networking.V1
{

    /// <summary>
    /// IngressClassSpec provides information about the class of an Ingress.
    /// </summary>
    [OutputType]
    public sealed class IngressClassSpecPatch
    {
        /// <summary>
        /// Controller refers to the name of the controller that should handle this class. This allows for different "flavors" that are controlled by the same controller. For example, you may have different Parameters for the same implementing controller. This should be specified as a domain-prefixed path no more than 250 characters in length, e.g. "acme.io/ingress-controller". This field is immutable.
        /// </summary>
        public readonly string Controller;
        /// <summary>
        /// Parameters is a link to a custom resource containing additional configuration for the controller. This is optional if the controller does not require extra parameters.
        /// </summary>
        public readonly Pulumi.Kubernetes.Types.Outputs.Networking.V1.IngressClassParametersReferencePatch Parameters;

        [OutputConstructor]
        private IngressClassSpecPatch(
            string controller,

            Pulumi.Kubernetes.Types.Outputs.Networking.V1.IngressClassParametersReferencePatch parameters)
        {
            Controller = controller;
            Parameters = parameters;
        }
    }
}
