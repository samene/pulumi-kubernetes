// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Kubernetes.Resource.V1Alpha2
{
    /// <summary>
    /// ResourceClass is used by administrators to influence how resources are allocated.
    /// 
    /// This is an alpha type and requires enabling the DynamicResourceAllocation feature gate.
    /// </summary>
    [KubernetesResourceType("kubernetes:resource.k8s.io/v1alpha2:ResourceClass")]
    public partial class ResourceClass : KubernetesResource
    {
        /// <summary>
        /// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
        /// </summary>
        [Output("apiVersion")]
        public Output<string> ApiVersion { get; private set; } = null!;

        /// <summary>
        /// DriverName defines the name of the dynamic resource driver that is used for allocation of a ResourceClaim that uses this class.
        /// 
        /// Resource drivers have a unique name in forward domain order (acme.example.com).
        /// </summary>
        [Output("driverName")]
        public Output<string> DriverName { get; private set; } = null!;

        /// <summary>
        /// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
        /// </summary>
        [Output("kind")]
        public Output<string> Kind { get; private set; } = null!;

        /// <summary>
        /// Standard object metadata
        /// </summary>
        [Output("metadata")]
        public Output<Pulumi.Kubernetes.Types.Outputs.Meta.V1.ObjectMeta> Metadata { get; private set; } = null!;

        /// <summary>
        /// ParametersRef references an arbitrary separate object that may hold parameters that will be used by the driver when allocating a resource that uses this class. A dynamic resource driver can distinguish between parameters stored here and and those stored in ResourceClaimSpec.
        /// </summary>
        [Output("parametersRef")]
        public Output<Pulumi.Kubernetes.Types.Outputs.Resource.V1Alpha2.ResourceClassParametersReference> ParametersRef { get; private set; } = null!;

        /// <summary>
        /// Only nodes matching the selector will be considered by the scheduler when trying to find a Node that fits a Pod when that Pod uses a ResourceClaim that has not been allocated yet.
        /// 
        /// Setting this field is optional. If null, all nodes are candidates.
        /// </summary>
        [Output("suitableNodes")]
        public Output<Pulumi.Kubernetes.Types.Outputs.Core.V1.NodeSelector> SuitableNodes { get; private set; } = null!;


        /// <summary>
        /// Create a ResourceClass resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public ResourceClass(string name, Pulumi.Kubernetes.Types.Inputs.Resource.V1Alpha2.ResourceClassArgs? args = null, CustomResourceOptions? options = null)
            : base("kubernetes:resource.k8s.io/v1alpha2:ResourceClass", name, MakeArgs(args), MakeResourceOptions(options, ""))
        {
        }
        internal ResourceClass(string name, ImmutableDictionary<string, object?> dictionary, CustomResourceOptions? options = null)
            : base("kubernetes:resource.k8s.io/v1alpha2:ResourceClass", name, new DictionaryResourceArgs(dictionary), MakeResourceOptions(options, ""))
        {
        }

        private ResourceClass(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("kubernetes:resource.k8s.io/v1alpha2:ResourceClass", name, null, MakeResourceOptions(options, id))
        {
        }

        private static Pulumi.Kubernetes.Types.Inputs.Resource.V1Alpha2.ResourceClassArgs? MakeArgs(Pulumi.Kubernetes.Types.Inputs.Resource.V1Alpha2.ResourceClassArgs? args)
        {
            args ??= new Pulumi.Kubernetes.Types.Inputs.Resource.V1Alpha2.ResourceClassArgs();
            args.ApiVersion = "resource.k8s.io/v1alpha2";
            args.Kind = "ResourceClass";
            return args;
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                Aliases =
                {
                    new global::Pulumi.Alias { Type = "kubernetes:resource.k8s.io/v1alpha1:ResourceClass" },
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing ResourceClass resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static ResourceClass Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new ResourceClass(name, id, options);
        }
    }
}
namespace Pulumi.Kubernetes.Types.Inputs.Resource.V1Alpha2
{

    public class ResourceClassArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
        /// </summary>
        [Input("apiVersion")]
        public Input<string>? ApiVersion { get; set; }

        /// <summary>
        /// DriverName defines the name of the dynamic resource driver that is used for allocation of a ResourceClaim that uses this class.
        /// 
        /// Resource drivers have a unique name in forward domain order (acme.example.com).
        /// </summary>
        [Input("driverName", required: true)]
        public Input<string> DriverName { get; set; } = null!;

        /// <summary>
        /// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
        /// </summary>
        [Input("kind")]
        public Input<string>? Kind { get; set; }

        /// <summary>
        /// Standard object metadata
        /// </summary>
        [Input("metadata")]
        public Input<Pulumi.Kubernetes.Types.Inputs.Meta.V1.ObjectMetaArgs>? Metadata { get; set; }

        /// <summary>
        /// ParametersRef references an arbitrary separate object that may hold parameters that will be used by the driver when allocating a resource that uses this class. A dynamic resource driver can distinguish between parameters stored here and and those stored in ResourceClaimSpec.
        /// </summary>
        [Input("parametersRef")]
        public Input<Pulumi.Kubernetes.Types.Inputs.Resource.V1Alpha2.ResourceClassParametersReferenceArgs>? ParametersRef { get; set; }

        /// <summary>
        /// Only nodes matching the selector will be considered by the scheduler when trying to find a Node that fits a Pod when that Pod uses a ResourceClaim that has not been allocated yet.
        /// 
        /// Setting this field is optional. If null, all nodes are candidates.
        /// </summary>
        [Input("suitableNodes")]
        public Input<Pulumi.Kubernetes.Types.Inputs.Core.V1.NodeSelectorArgs>? SuitableNodes { get; set; }

        public ResourceClassArgs()
        {
        }
        public static new ResourceClassArgs Empty => new ResourceClassArgs();
    }
}
