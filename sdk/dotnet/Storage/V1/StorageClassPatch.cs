// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Kubernetes.Storage.V1
{
    /// <summary>
    /// Patch resources are used to modify existing Kubernetes resources by using
    /// Server-Side Apply updates. The name of the resource must be specified, but all other properties are optional. More than
    /// one patch may be applied to the same resource, and a random FieldManager name will be used for each Patch resource.
    /// Conflicts will result in an error by default, but can be forced using the "pulumi.com/patchForce" annotation. See the
    /// [Server-Side Apply Docs](https://www.pulumi.com/registry/packages/kubernetes/how-to-guides/managing-resources-with-server-side-apply/) for
    /// additional information about using Server-Side Apply to manage Kubernetes resources with Pulumi.
    /// StorageClass describes the parameters for a class of storage for which PersistentVolumes can be dynamically provisioned.
    /// 
    /// StorageClasses are non-namespaced; the name of the storage class according to etcd is in ObjectMeta.Name.
    /// </summary>
    [KubernetesResourceType("kubernetes:storage.k8s.io/v1:StorageClassPatch")]
    public partial class StorageClassPatch : KubernetesResource
    {
        /// <summary>
        /// allowVolumeExpansion shows whether the storage class allow volume expand.
        /// </summary>
        [Output("allowVolumeExpansion")]
        public Output<bool> AllowVolumeExpansion { get; private set; } = null!;

        /// <summary>
        /// allowedTopologies restrict the node topologies where volumes can be dynamically provisioned. Each volume plugin defines its own supported topology specifications. An empty TopologySelectorTerm list means there is no topology restriction. This field is only honored by servers that enable the VolumeScheduling feature.
        /// </summary>
        [Output("allowedTopologies")]
        public Output<ImmutableArray<Pulumi.Kubernetes.Types.Outputs.Core.V1.TopologySelectorTermPatch>> AllowedTopologies { get; private set; } = null!;

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
        /// mountOptions controls the mountOptions for dynamically provisioned PersistentVolumes of this storage class. e.g. ["ro", "soft"]. Not validated - mount of the PVs will simply fail if one is invalid.
        /// </summary>
        [Output("mountOptions")]
        public Output<ImmutableArray<string>> MountOptions { get; private set; } = null!;

        /// <summary>
        /// parameters holds the parameters for the provisioner that should create volumes of this storage class.
        /// </summary>
        [Output("parameters")]
        public Output<ImmutableDictionary<string, string>> Parameters { get; private set; } = null!;

        /// <summary>
        /// provisioner indicates the type of the provisioner.
        /// </summary>
        [Output("provisioner")]
        public Output<string> Provisioner { get; private set; } = null!;

        /// <summary>
        /// reclaimPolicy controls the reclaimPolicy for dynamically provisioned PersistentVolumes of this storage class. Defaults to Delete.
        /// </summary>
        [Output("reclaimPolicy")]
        public Output<string> ReclaimPolicy { get; private set; } = null!;

        /// <summary>
        /// volumeBindingMode indicates how PersistentVolumeClaims should be provisioned and bound.  When unset, VolumeBindingImmediate is used. This field is only honored by servers that enable the VolumeScheduling feature.
        /// </summary>
        [Output("volumeBindingMode")]
        public Output<string> VolumeBindingMode { get; private set; } = null!;


        /// <summary>
        /// Create a StorageClassPatch resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public StorageClassPatch(string name, Pulumi.Kubernetes.Types.Inputs.Storage.V1.StorageClassPatchArgs? args = null, CustomResourceOptions? options = null)
            : base("kubernetes:storage.k8s.io/v1:StorageClassPatch", name, MakeArgs(args), MakeResourceOptions(options, ""))
        {
        }
        internal StorageClassPatch(string name, ImmutableDictionary<string, object?> dictionary, CustomResourceOptions? options = null)
            : base("kubernetes:storage.k8s.io/v1:StorageClassPatch", name, new DictionaryResourceArgs(dictionary), MakeResourceOptions(options, ""))
        {
        }

        private StorageClassPatch(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("kubernetes:storage.k8s.io/v1:StorageClassPatch", name, null, MakeResourceOptions(options, id))
        {
        }

        private static Pulumi.Kubernetes.Types.Inputs.Storage.V1.StorageClassPatchArgs? MakeArgs(Pulumi.Kubernetes.Types.Inputs.Storage.V1.StorageClassPatchArgs? args)
        {
            args ??= new Pulumi.Kubernetes.Types.Inputs.Storage.V1.StorageClassPatchArgs();
            args.ApiVersion = "storage.k8s.io/v1";
            args.Kind = "StorageClass";
            return args;
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                Aliases =
                {
                    new global::Pulumi.Alias { Type = "kubernetes:storage.k8s.io/v1beta1:StorageClassPatch" },
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing StorageClassPatch resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static StorageClassPatch Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new StorageClassPatch(name, id, options);
        }
    }
}
namespace Pulumi.Kubernetes.Types.Inputs.Storage.V1
{

    public class StorageClassPatchArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// allowVolumeExpansion shows whether the storage class allow volume expand.
        /// </summary>
        [Input("allowVolumeExpansion")]
        public Input<bool>? AllowVolumeExpansion { get; set; }

        [Input("allowedTopologies")]
        private InputList<Pulumi.Kubernetes.Types.Inputs.Core.V1.TopologySelectorTermPatchArgs>? _allowedTopologies;

        /// <summary>
        /// allowedTopologies restrict the node topologies where volumes can be dynamically provisioned. Each volume plugin defines its own supported topology specifications. An empty TopologySelectorTerm list means there is no topology restriction. This field is only honored by servers that enable the VolumeScheduling feature.
        /// </summary>
        public InputList<Pulumi.Kubernetes.Types.Inputs.Core.V1.TopologySelectorTermPatchArgs> AllowedTopologies
        {
            get => _allowedTopologies ?? (_allowedTopologies = new InputList<Pulumi.Kubernetes.Types.Inputs.Core.V1.TopologySelectorTermPatchArgs>());
            set => _allowedTopologies = value;
        }

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

        [Input("mountOptions")]
        private InputList<string>? _mountOptions;

        /// <summary>
        /// mountOptions controls the mountOptions for dynamically provisioned PersistentVolumes of this storage class. e.g. ["ro", "soft"]. Not validated - mount of the PVs will simply fail if one is invalid.
        /// </summary>
        public InputList<string> MountOptions
        {
            get => _mountOptions ?? (_mountOptions = new InputList<string>());
            set => _mountOptions = value;
        }

        [Input("parameters")]
        private InputMap<string>? _parameters;

        /// <summary>
        /// parameters holds the parameters for the provisioner that should create volumes of this storage class.
        /// </summary>
        public InputMap<string> Parameters
        {
            get => _parameters ?? (_parameters = new InputMap<string>());
            set => _parameters = value;
        }

        /// <summary>
        /// provisioner indicates the type of the provisioner.
        /// </summary>
        [Input("provisioner")]
        public Input<string>? Provisioner { get; set; }

        /// <summary>
        /// reclaimPolicy controls the reclaimPolicy for dynamically provisioned PersistentVolumes of this storage class. Defaults to Delete.
        /// </summary>
        [Input("reclaimPolicy")]
        public Input<string>? ReclaimPolicy { get; set; }

        /// <summary>
        /// volumeBindingMode indicates how PersistentVolumeClaims should be provisioned and bound.  When unset, VolumeBindingImmediate is used. This field is only honored by servers that enable the VolumeScheduling feature.
        /// </summary>
        [Input("volumeBindingMode")]
        public Input<string>? VolumeBindingMode { get; set; }

        public StorageClassPatchArgs()
        {
        }
        public static new StorageClassPatchArgs Empty => new StorageClassPatchArgs();
    }
}
