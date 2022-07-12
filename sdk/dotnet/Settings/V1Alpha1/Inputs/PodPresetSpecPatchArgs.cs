// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Kubernetes.Types.Inputs.Settings.V1Alpha1
{

    /// <summary>
    /// PodPresetSpec is a description of a pod preset.
    /// </summary>
    public class PodPresetSpecPatchArgs : Pulumi.ResourceArgs
    {
        [Input("env")]
        private InputList<Pulumi.Kubernetes.Types.Inputs.Core.V1.EnvVarPatchArgs>? _env;

        /// <summary>
        /// Env defines the collection of EnvVar to inject into containers.
        /// </summary>
        public InputList<Pulumi.Kubernetes.Types.Inputs.Core.V1.EnvVarPatchArgs> Env
        {
            get => _env ?? (_env = new InputList<Pulumi.Kubernetes.Types.Inputs.Core.V1.EnvVarPatchArgs>());
            set => _env = value;
        }

        [Input("envFrom")]
        private InputList<Pulumi.Kubernetes.Types.Inputs.Core.V1.EnvFromSourcePatchArgs>? _envFrom;

        /// <summary>
        /// EnvFrom defines the collection of EnvFromSource to inject into containers.
        /// </summary>
        public InputList<Pulumi.Kubernetes.Types.Inputs.Core.V1.EnvFromSourcePatchArgs> EnvFrom
        {
            get => _envFrom ?? (_envFrom = new InputList<Pulumi.Kubernetes.Types.Inputs.Core.V1.EnvFromSourcePatchArgs>());
            set => _envFrom = value;
        }

        /// <summary>
        /// Selector is a label query over a set of resources, in this case pods. Required.
        /// </summary>
        [Input("selector")]
        public Input<Pulumi.Kubernetes.Types.Inputs.Meta.V1.LabelSelectorPatchArgs>? Selector { get; set; }

        [Input("volumeMounts")]
        private InputList<Pulumi.Kubernetes.Types.Inputs.Core.V1.VolumeMountPatchArgs>? _volumeMounts;

        /// <summary>
        /// VolumeMounts defines the collection of VolumeMount to inject into containers.
        /// </summary>
        public InputList<Pulumi.Kubernetes.Types.Inputs.Core.V1.VolumeMountPatchArgs> VolumeMounts
        {
            get => _volumeMounts ?? (_volumeMounts = new InputList<Pulumi.Kubernetes.Types.Inputs.Core.V1.VolumeMountPatchArgs>());
            set => _volumeMounts = value;
        }

        [Input("volumes")]
        private InputList<Pulumi.Kubernetes.Types.Inputs.Core.V1.VolumePatchArgs>? _volumes;

        /// <summary>
        /// Volumes defines the collection of Volume to inject into the pod.
        /// </summary>
        public InputList<Pulumi.Kubernetes.Types.Inputs.Core.V1.VolumePatchArgs> Volumes
        {
            get => _volumes ?? (_volumes = new InputList<Pulumi.Kubernetes.Types.Inputs.Core.V1.VolumePatchArgs>());
            set => _volumes = value;
        }

        public PodPresetSpecPatchArgs()
        {
        }
    }
}
