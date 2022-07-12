// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Kubernetes.Types.Inputs.Autoscaling.V2Beta1
{

    /// <summary>
    /// ContainerResourceMetricSource indicates how to scale on a resource metric known to Kubernetes, as specified in requests and limits, describing each pod in the current scale target (e.g. CPU or memory).  The values will be averaged together before being compared to the target.  Such metrics are built in to Kubernetes, and have special scaling options on top of those available to normal per-pod metrics using the "pods" source.  Only one "target" type should be set.
    /// </summary>
    public class ContainerResourceMetricSourcePatchArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// container is the name of the container in the pods of the scaling target
        /// </summary>
        [Input("container")]
        public Input<string>? Container { get; set; }

        /// <summary>
        /// name is the name of the resource in question.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// targetAverageUtilization is the target value of the average of the resource metric across all relevant pods, represented as a percentage of the requested value of the resource for the pods.
        /// </summary>
        [Input("targetAverageUtilization")]
        public Input<int>? TargetAverageUtilization { get; set; }

        /// <summary>
        /// targetAverageValue is the target value of the average of the resource metric across all relevant pods, as a raw value (instead of as a percentage of the request), similar to the "pods" metric source type.
        /// </summary>
        [Input("targetAverageValue")]
        public Input<string>? TargetAverageValue { get; set; }

        public ContainerResourceMetricSourcePatchArgs()
        {
        }
    }
}
