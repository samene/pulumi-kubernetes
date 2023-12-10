// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Kubernetes.Autoscaling.V2
{
    /// <summary>
    /// Patch resources are used to modify existing Kubernetes resources by using
    /// Server-Side Apply updates. The name of the resource must be specified, but all other properties are optional. More than
    /// one patch may be applied to the same resource, and a random FieldManager name will be used for each Patch resource.
    /// Conflicts will result in an error by default, but can be forced using the "pulumi.com/patchForce" annotation. See the
    /// [Server-Side Apply Docs](https://www.pulumi.com/registry/packages/kubernetes/how-to-guides/managing-resources-with-server-side-apply/) for
    /// additional information about using Server-Side Apply to manage Kubernetes resources with Pulumi.
    /// HorizontalPodAutoscaler is the configuration for a horizontal pod autoscaler, which automatically manages the replica count of any resource implementing the scale subresource based on the metrics specified.
    /// </summary>
    [KubernetesResourceType("kubernetes:autoscaling/v2:HorizontalPodAutoscalerPatch")]
    public partial class HorizontalPodAutoscalerPatch : KubernetesResource
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
        /// metadata is the standard object metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
        /// </summary>
        [Output("metadata")]
        public Output<Pulumi.Kubernetes.Types.Outputs.Meta.V1.ObjectMetaPatch> Metadata { get; private set; } = null!;

        /// <summary>
        /// spec is the specification for the behaviour of the autoscaler. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#spec-and-status.
        /// </summary>
        [Output("spec")]
        public Output<Pulumi.Kubernetes.Types.Outputs.Autoscaling.V2.HorizontalPodAutoscalerSpecPatch> Spec { get; private set; } = null!;

        /// <summary>
        /// status is the current information about the autoscaler.
        /// </summary>
        [Output("status")]
        public Output<Pulumi.Kubernetes.Types.Outputs.Autoscaling.V2.HorizontalPodAutoscalerStatusPatch> Status { get; private set; } = null!;


        /// <summary>
        /// Create a HorizontalPodAutoscalerPatch resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public HorizontalPodAutoscalerPatch(string name, Pulumi.Kubernetes.Types.Inputs.Autoscaling.V2.HorizontalPodAutoscalerPatchArgs? args = null, CustomResourceOptions? options = null)
            : base("kubernetes:autoscaling/v2:HorizontalPodAutoscalerPatch", name, MakeArgs(args), MakeResourceOptions(options, ""))
        {
        }
        internal HorizontalPodAutoscalerPatch(string name, ImmutableDictionary<string, object?> dictionary, CustomResourceOptions? options = null)
            : base("kubernetes:autoscaling/v2:HorizontalPodAutoscalerPatch", name, new DictionaryResourceArgs(dictionary), MakeResourceOptions(options, ""))
        {
        }

        private HorizontalPodAutoscalerPatch(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("kubernetes:autoscaling/v2:HorizontalPodAutoscalerPatch", name, null, MakeResourceOptions(options, id))
        {
        }

        private static Pulumi.Kubernetes.Types.Inputs.Autoscaling.V2.HorizontalPodAutoscalerPatchArgs? MakeArgs(Pulumi.Kubernetes.Types.Inputs.Autoscaling.V2.HorizontalPodAutoscalerPatchArgs? args)
        {
            args ??= new Pulumi.Kubernetes.Types.Inputs.Autoscaling.V2.HorizontalPodAutoscalerPatchArgs();
            args.ApiVersion = "autoscaling/v2";
            args.Kind = "HorizontalPodAutoscaler";
            return args;
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                Aliases =
                {
                    new global::Pulumi.Alias { Type = "kubernetes:autoscaling/v1:HorizontalPodAutoscalerPatch" },
                    new global::Pulumi.Alias { Type = "kubernetes:autoscaling/v2beta1:HorizontalPodAutoscalerPatch" },
                    new global::Pulumi.Alias { Type = "kubernetes:autoscaling/v2beta2:HorizontalPodAutoscalerPatch" },
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing HorizontalPodAutoscalerPatch resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static HorizontalPodAutoscalerPatch Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new HorizontalPodAutoscalerPatch(name, id, options);
        }
    }
}
namespace Pulumi.Kubernetes.Types.Inputs.Autoscaling.V2
{

    public class HorizontalPodAutoscalerPatchArgs : global::Pulumi.ResourceArgs
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
        /// metadata is the standard object metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
        /// </summary>
        [Input("metadata")]
        public Input<Pulumi.Kubernetes.Types.Inputs.Meta.V1.ObjectMetaPatchArgs>? Metadata { get; set; }

        /// <summary>
        /// spec is the specification for the behaviour of the autoscaler. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#spec-and-status.
        /// </summary>
        [Input("spec")]
        public Input<Pulumi.Kubernetes.Types.Inputs.Autoscaling.V2.HorizontalPodAutoscalerSpecPatchArgs>? Spec { get; set; }

        public HorizontalPodAutoscalerPatchArgs()
        {
        }
        public static new HorizontalPodAutoscalerPatchArgs Empty => new HorizontalPodAutoscalerPatchArgs();
    }
}
