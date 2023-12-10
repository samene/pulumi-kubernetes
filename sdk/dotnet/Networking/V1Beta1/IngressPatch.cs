// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Kubernetes.Networking.V1Beta1
{
    /// <summary>
    /// Patch resources are used to modify existing Kubernetes resources by using
    /// Server-Side Apply updates. The name of the resource must be specified, but all other properties are optional. More than
    /// one patch may be applied to the same resource, and a random FieldManager name will be used for each Patch resource.
    /// Conflicts will result in an error by default, but can be forced using the "pulumi.com/patchForce" annotation. See the
    /// [Server-Side Apply Docs](https://www.pulumi.com/registry/packages/kubernetes/how-to-guides/managing-resources-with-server-side-apply/) for
    /// additional information about using Server-Side Apply to manage Kubernetes resources with Pulumi.
    /// Ingress is a collection of rules that allow inbound connections to reach the endpoints defined by a backend. An Ingress can be configured to give services externally-reachable urls, load balance traffic, terminate SSL, offer name based virtual hosting etc.
    /// 
    /// This resource waits until its status is ready before registering success
    /// for create/update, and populating output properties from the current state of the resource.
    /// The following conditions are used to determine whether the resource creation has
    /// succeeded or failed:
    /// 
    /// 1.  Ingress object exists.
    /// 2.  Endpoint objects exist with matching names for each Ingress path (except when Service
    ///     type is ExternalName).
    /// 3.  Ingress entry exists for '.status.loadBalancer.ingress'.
    /// 
    /// If the Ingress has not reached a Ready state after 10 minutes, it will
    /// time out and mark the resource update as Failed. You can override the default timeout value
    /// by setting the 'customTimeouts' option on the resource.
    /// </summary>
    [KubernetesResourceType("kubernetes:networking.k8s.io/v1beta1:IngressPatch")]
    public partial class IngressPatch : KubernetesResource
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
        /// Spec is the desired state of the Ingress. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#spec-and-status
        /// </summary>
        [Output("spec")]
        public Output<Pulumi.Kubernetes.Types.Outputs.Networking.V1Beta1.IngressSpecPatch> Spec { get; private set; } = null!;

        /// <summary>
        /// Status is the current state of the Ingress. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#spec-and-status
        /// </summary>
        [Output("status")]
        public Output<Pulumi.Kubernetes.Types.Outputs.Networking.V1Beta1.IngressStatusPatch> Status { get; private set; } = null!;


        /// <summary>
        /// Create a IngressPatch resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public IngressPatch(string name, Pulumi.Kubernetes.Types.Inputs.Networking.V1Beta1.IngressPatchArgs? args = null, CustomResourceOptions? options = null)
            : base("kubernetes:networking.k8s.io/v1beta1:IngressPatch", name, MakeArgs(args), MakeResourceOptions(options, ""))
        {
        }
        internal IngressPatch(string name, ImmutableDictionary<string, object?> dictionary, CustomResourceOptions? options = null)
            : base("kubernetes:networking.k8s.io/v1beta1:IngressPatch", name, new DictionaryResourceArgs(dictionary), MakeResourceOptions(options, ""))
        {
        }

        private IngressPatch(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("kubernetes:networking.k8s.io/v1beta1:IngressPatch", name, null, MakeResourceOptions(options, id))
        {
        }

        private static Pulumi.Kubernetes.Types.Inputs.Networking.V1Beta1.IngressPatchArgs? MakeArgs(Pulumi.Kubernetes.Types.Inputs.Networking.V1Beta1.IngressPatchArgs? args)
        {
            args ??= new Pulumi.Kubernetes.Types.Inputs.Networking.V1Beta1.IngressPatchArgs();
            args.ApiVersion = "networking.k8s.io/v1beta1";
            args.Kind = "Ingress";
            return args;
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                Aliases =
                {
                    new global::Pulumi.Alias { Type = "kubernetes:extensions/v1beta1:IngressPatch" },
                    new global::Pulumi.Alias { Type = "kubernetes:networking.k8s.io/v1:IngressPatch" },
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing IngressPatch resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static IngressPatch Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new IngressPatch(name, id, options);
        }
    }
}
namespace Pulumi.Kubernetes.Types.Inputs.Networking.V1Beta1
{

    public class IngressPatchArgs : global::Pulumi.ResourceArgs
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
        /// Spec is the desired state of the Ingress. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#spec-and-status
        /// </summary>
        [Input("spec")]
        public Input<Pulumi.Kubernetes.Types.Inputs.Networking.V1Beta1.IngressSpecPatchArgs>? Spec { get; set; }

        public IngressPatchArgs()
        {
        }
        public static new IngressPatchArgs Empty => new IngressPatchArgs();
    }
}
