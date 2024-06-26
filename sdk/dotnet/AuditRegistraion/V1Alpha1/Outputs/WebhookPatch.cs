// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Kubernetes.Types.Outputs.AuditRegistraion.V1Alpha1
{

    /// <summary>
    /// Webhook holds the configuration of the webhook
    /// </summary>
    [OutputType]
    public sealed class WebhookPatch
    {
        /// <summary>
        /// ClientConfig holds the connection parameters for the webhook required
        /// </summary>
        public readonly Pulumi.Kubernetes.Types.Outputs.AuditRegistraion.V1Alpha1.WebhookClientConfigPatch ClientConfig;
        /// <summary>
        /// Throttle holds the options for throttling the webhook
        /// </summary>
        public readonly Pulumi.Kubernetes.Types.Outputs.AuditRegistraion.V1Alpha1.WebhookThrottleConfigPatch Throttle;

        [OutputConstructor]
        private WebhookPatch(
            Pulumi.Kubernetes.Types.Outputs.AuditRegistraion.V1Alpha1.WebhookClientConfigPatch clientConfig,

            Pulumi.Kubernetes.Types.Outputs.AuditRegistraion.V1Alpha1.WebhookThrottleConfigPatch throttle)
        {
            ClientConfig = clientConfig;
            Throttle = throttle;
        }
    }
}
