// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Kubernetes.Types.Outputs.Core.V1
{

    /// <summary>
    /// Lifecycle describes actions that the management system should take in response to container lifecycle events. For the PostStart and PreStop lifecycle handlers, management of the container blocks until the action is complete, unless the container process fails, in which case the handler is aborted.
    /// </summary>
    [OutputType]
    public sealed class LifecyclePatch
    {
        /// <summary>
        /// PostStart is called immediately after a container is created. If the handler fails, the container is terminated and restarted according to its restart policy. Other management of the container blocks until the hook completes. More info: https://kubernetes.io/docs/concepts/containers/container-lifecycle-hooks/#container-hooks
        /// </summary>
        public readonly Pulumi.Kubernetes.Types.Outputs.Core.V1.LifecycleHandlerPatch PostStart;
        /// <summary>
        /// PreStop is called immediately before a container is terminated due to an API request or management event such as liveness/startup probe failure, preemption, resource contention, etc. The handler is not called if the container crashes or exits. The Pod's termination grace period countdown begins before the PreStop hook is executed. Regardless of the outcome of the handler, the container will eventually terminate within the Pod's termination grace period (unless delayed by finalizers). Other management of the container blocks until the hook completes or until the termination grace period is reached. More info: https://kubernetes.io/docs/concepts/containers/container-lifecycle-hooks/#container-hooks
        /// </summary>
        public readonly Pulumi.Kubernetes.Types.Outputs.Core.V1.LifecycleHandlerPatch PreStop;

        [OutputConstructor]
        private LifecyclePatch(
            Pulumi.Kubernetes.Types.Outputs.Core.V1.LifecycleHandlerPatch postStart,

            Pulumi.Kubernetes.Types.Outputs.Core.V1.LifecycleHandlerPatch preStop)
        {
            PostStart = postStart;
            PreStop = preStop;
        }
    }
}
