// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Kubernetes.Types.Inputs.Autoscaling.V2
{

    /// <summary>
    /// MetricTarget defines the target value, average value, or average utilization of a specific metric
    /// </summary>
    public class MetricTargetPatchArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// averageUtilization is the target value of the average of the resource metric across all relevant pods, represented as a percentage of the requested value of the resource for the pods. Currently only valid for Resource metric source type
        /// </summary>
        [Input("averageUtilization")]
        public Input<int>? AverageUtilization { get; set; }

        /// <summary>
        /// averageValue is the target value of the average of the metric across all relevant pods (as a quantity)
        /// </summary>
        [Input("averageValue")]
        public Input<string>? AverageValue { get; set; }

        /// <summary>
        /// type represents whether the metric type is Utilization, Value, or AverageValue
        /// </summary>
        [Input("type")]
        public Input<string>? Type { get; set; }

        /// <summary>
        /// value is the target value of the metric (as a quantity).
        /// </summary>
        [Input("value")]
        public Input<string>? Value { get; set; }

        public MetricTargetPatchArgs()
        {
        }
    }
}
