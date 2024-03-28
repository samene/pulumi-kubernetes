// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Kubernetes.Types.Outputs.FlowControl.V1
{

    /// <summary>
    /// FlowSchemaSpec describes how the FlowSchema's specification looks like.
    /// </summary>
    [OutputType]
    public sealed class FlowSchemaSpecPatch
    {
        /// <summary>
        /// `distinguisherMethod` defines how to compute the flow distinguisher for requests that match this schema. `nil` specifies that the distinguisher is disabled and thus will always be the empty string.
        /// </summary>
        public readonly Pulumi.Kubernetes.Types.Outputs.FlowControl.V1.FlowDistinguisherMethodPatch DistinguisherMethod;
        /// <summary>
        /// `matchingPrecedence` is used to choose among the FlowSchemas that match a given request. The chosen FlowSchema is among those with the numerically lowest (which we take to be logically highest) MatchingPrecedence.  Each MatchingPrecedence value must be ranged in [1,10000]. Note that if the precedence is not specified, it will be set to 1000 as default.
        /// </summary>
        public readonly int MatchingPrecedence;
        /// <summary>
        /// `priorityLevelConfiguration` should reference a PriorityLevelConfiguration in the cluster. If the reference cannot be resolved, the FlowSchema will be ignored and marked as invalid in its status. Required.
        /// </summary>
        public readonly Pulumi.Kubernetes.Types.Outputs.FlowControl.V1.PriorityLevelConfigurationReferencePatch PriorityLevelConfiguration;
        /// <summary>
        /// `rules` describes which requests will match this flow schema. This FlowSchema matches a request if and only if at least one member of rules matches the request. if it is an empty slice, there will be no requests matching the FlowSchema.
        /// </summary>
        public readonly ImmutableArray<Pulumi.Kubernetes.Types.Outputs.FlowControl.V1.PolicyRulesWithSubjectsPatch> Rules;

        [OutputConstructor]
        private FlowSchemaSpecPatch(
            Pulumi.Kubernetes.Types.Outputs.FlowControl.V1.FlowDistinguisherMethodPatch distinguisherMethod,

            int matchingPrecedence,

            Pulumi.Kubernetes.Types.Outputs.FlowControl.V1.PriorityLevelConfigurationReferencePatch priorityLevelConfiguration,

            ImmutableArray<Pulumi.Kubernetes.Types.Outputs.FlowControl.V1.PolicyRulesWithSubjectsPatch> rules)
        {
            DistinguisherMethod = distinguisherMethod;
            MatchingPrecedence = matchingPrecedence;
            PriorityLevelConfiguration = priorityLevelConfiguration;
            Rules = rules;
        }
    }
}
