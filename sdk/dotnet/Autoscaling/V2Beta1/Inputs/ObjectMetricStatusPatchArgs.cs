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
    /// ObjectMetricStatus indicates the current value of a metric describing a kubernetes object (for example, hits-per-second on an Ingress object).
    /// </summary>
    public class ObjectMetricStatusPatchArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// averageValue is the current value of the average of the metric across all relevant pods (as a quantity)
        /// </summary>
        [Input("averageValue")]
        public Input<string>? AverageValue { get; set; }

        /// <summary>
        /// currentValue is the current value of the metric (as a quantity).
        /// </summary>
        [Input("currentValue")]
        public Input<string>? CurrentValue { get; set; }

        /// <summary>
        /// metricName is the name of the metric in question.
        /// </summary>
        [Input("metricName")]
        public Input<string>? MetricName { get; set; }

        /// <summary>
        /// selector is the string-encoded form of a standard kubernetes label selector for the given metric When set in the ObjectMetricSource, it is passed as an additional parameter to the metrics server for more specific metrics scoping. When unset, just the metricName will be used to gather metrics.
        /// </summary>
        [Input("selector")]
        public Input<Pulumi.Kubernetes.Types.Inputs.Meta.V1.LabelSelectorPatchArgs>? Selector { get; set; }

        /// <summary>
        /// target is the described Kubernetes object.
        /// </summary>
        [Input("target")]
        public Input<Pulumi.Kubernetes.Types.Inputs.Autoscaling.V2Beta1.CrossVersionObjectReferencePatchArgs>? Target { get; set; }

        public ObjectMetricStatusPatchArgs()
        {
        }
    }
}
