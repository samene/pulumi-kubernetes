// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Kubernetes.Types.Inputs.Autoscaling.V2Beta2
{

    /// <summary>
    /// HorizontalPodAutoscalerStatus describes the current status of a horizontal pod autoscaler.
    /// </summary>
    public class HorizontalPodAutoscalerStatusPatchArgs : Pulumi.ResourceArgs
    {
        [Input("conditions")]
        private InputList<Pulumi.Kubernetes.Types.Inputs.Autoscaling.V2Beta2.HorizontalPodAutoscalerConditionPatchArgs>? _conditions;

        /// <summary>
        /// conditions is the set of conditions required for this autoscaler to scale its target, and indicates whether or not those conditions are met.
        /// </summary>
        public InputList<Pulumi.Kubernetes.Types.Inputs.Autoscaling.V2Beta2.HorizontalPodAutoscalerConditionPatchArgs> Conditions
        {
            get => _conditions ?? (_conditions = new InputList<Pulumi.Kubernetes.Types.Inputs.Autoscaling.V2Beta2.HorizontalPodAutoscalerConditionPatchArgs>());
            set => _conditions = value;
        }

        [Input("currentMetrics")]
        private InputList<Pulumi.Kubernetes.Types.Inputs.Autoscaling.V2Beta2.MetricStatusPatchArgs>? _currentMetrics;

        /// <summary>
        /// currentMetrics is the last read state of the metrics used by this autoscaler.
        /// </summary>
        public InputList<Pulumi.Kubernetes.Types.Inputs.Autoscaling.V2Beta2.MetricStatusPatchArgs> CurrentMetrics
        {
            get => _currentMetrics ?? (_currentMetrics = new InputList<Pulumi.Kubernetes.Types.Inputs.Autoscaling.V2Beta2.MetricStatusPatchArgs>());
            set => _currentMetrics = value;
        }

        /// <summary>
        /// currentReplicas is current number of replicas of pods managed by this autoscaler, as last seen by the autoscaler.
        /// </summary>
        [Input("currentReplicas")]
        public Input<int>? CurrentReplicas { get; set; }

        /// <summary>
        /// desiredReplicas is the desired number of replicas of pods managed by this autoscaler, as last calculated by the autoscaler.
        /// </summary>
        [Input("desiredReplicas")]
        public Input<int>? DesiredReplicas { get; set; }

        /// <summary>
        /// lastScaleTime is the last time the HorizontalPodAutoscaler scaled the number of pods, used by the autoscaler to control how often the number of pods is changed.
        /// </summary>
        [Input("lastScaleTime")]
        public Input<string>? LastScaleTime { get; set; }

        /// <summary>
        /// observedGeneration is the most recent generation observed by this autoscaler.
        /// </summary>
        [Input("observedGeneration")]
        public Input<int>? ObservedGeneration { get; set; }

        public HorizontalPodAutoscalerStatusPatchArgs()
        {
        }
    }
}
