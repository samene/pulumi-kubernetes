// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Kubernetes.Core.V1
{
    /// <summary>
    /// Service is a named abstraction of software service (for example, mysql) consisting of local port (for example 3306) that the proxy listens on, and the selector that determines which pods will answer requests sent through the proxy.
    /// 
    /// This resource waits until its status is ready before registering success
    /// for create/update, and populating output properties from the current state of the resource.
    /// The following conditions are used to determine whether the resource creation has
    /// succeeded or failed:
    /// 
    /// 1. Service object exists.
    /// 2. Related Endpoint objects are created. Each time we get an update, wait 10 seconds
    ///    for any stragglers.
    /// 3. The endpoints objects target some number of living objects (unless the Service is
    ///    an "empty headless" Service [1] or a Service with '.spec.type: ExternalName').
    /// 4. External IP address is allocated (if Service has '.spec.type: LoadBalancer').
    /// 
    /// Known limitations:
    /// Services targeting ReplicaSets (and, by extension, Deployments,
    /// StatefulSets, etc.) with '.spec.replicas' set to 0 are not handled, and will time
    /// out. To work around this limitation, set 'pulumi.com/skipAwait: "true"' on
    /// '.metadata.annotations' for the Service. Work to handle this case is in progress [2].
    /// 
    /// [1] https://kubernetes.io/docs/concepts/services-networking/service/#headless-services
    /// [2] https://github.com/pulumi/pulumi-kubernetes/pull/703
    /// 
    /// If the Service has not reached a Ready state after 10 minutes, it will
    /// time out and mark the resource update as Failed. You can override the default timeout value
    /// by setting the 'customTimeouts' option on the resource.
    /// </summary>
    [KubernetesResourceType("kubernetes:core/v1:ServicePatch")]
    public partial class ServicePatch : KubernetesResource
    {
        /// <summary>
        /// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
        /// </summary>
        [Output("apiVersion")]
        public Output<string> ApiVersion { get; private set; } = null!;

        /// <summary>
        /// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
        /// </summary>
        [Output("kind")]
        public Output<string> Kind { get; private set; } = null!;

        /// <summary>
        /// Standard object's metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
        /// </summary>
        [Output("metadata")]
        public Output<Pulumi.Kubernetes.Types.Outputs.Meta.V1.ObjectMetaPatch> Metadata { get; private set; } = null!;

        /// <summary>
        /// Spec defines the behavior of a service. https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#spec-and-status
        /// </summary>
        [Output("spec")]
        public Output<Pulumi.Kubernetes.Types.Outputs.Core.V1.ServiceSpecPatch> Spec { get; private set; } = null!;

        /// <summary>
        /// Most recently observed status of the service. Populated by the system. Read-only. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#spec-and-status
        /// </summary>
        [Output("status")]
        public Output<Pulumi.Kubernetes.Types.Outputs.Core.V1.ServiceStatusPatch> Status { get; private set; } = null!;


        /// <summary>
        /// Create a ServicePatch resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public ServicePatch(string name, Pulumi.Kubernetes.Types.Inputs.Core.V1.ServicePatchArgs? args = null, CustomResourceOptions? options = null)
            : base("kubernetes:core/v1:ServicePatch", name, MakeArgs(args), MakeResourceOptions(options, ""))
        {
        }
        internal ServicePatch(string name, ImmutableDictionary<string, object?> dictionary, CustomResourceOptions? options = null)
            : base("kubernetes:core/v1:ServicePatch", name, new DictionaryResourceArgs(dictionary), MakeResourceOptions(options, ""))
        {
        }

        private ServicePatch(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("kubernetes:core/v1:ServicePatch", name, null, MakeResourceOptions(options, id))
        {
        }

        private static Pulumi.Kubernetes.Types.Inputs.Core.V1.ServicePatchArgs? MakeArgs(Pulumi.Kubernetes.Types.Inputs.Core.V1.ServicePatchArgs? args)
        {
            args ??= new Pulumi.Kubernetes.Types.Inputs.Core.V1.ServicePatchArgs();
            args.ApiVersion = "v1";
            args.Kind = "Service";
            return args;
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing ServicePatch resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static ServicePatch Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new ServicePatch(name, id, options);
        }
    }
}
namespace Pulumi.Kubernetes.Types.Inputs.Core.V1
{

    public class ServicePatchArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
        /// </summary>
        [Input("apiVersion")]
        public Input<string>? ApiVersion { get; set; }

        /// <summary>
        /// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
        /// </summary>
        [Input("kind")]
        public Input<string>? Kind { get; set; }

        /// <summary>
        /// Standard object's metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
        /// </summary>
        [Input("metadata")]
        public Input<Pulumi.Kubernetes.Types.Inputs.Meta.V1.ObjectMetaPatchArgs>? Metadata { get; set; }

        /// <summary>
        /// Spec defines the behavior of a service. https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#spec-and-status
        /// </summary>
        [Input("spec")]
        public Input<Pulumi.Kubernetes.Types.Inputs.Core.V1.ServiceSpecPatchArgs>? Spec { get; set; }

        public ServicePatchArgs()
        {
        }
    }
}
