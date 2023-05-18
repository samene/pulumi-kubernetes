// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as enums from "./types/enums";
import * as utilities from "./utilities";

/**
 * The provider type for the kubernetes package.
 */
export class Provider extends pulumi.ProviderResource {
    /** @internal */
    public static readonly __pulumiType = 'kubernetes';

    /**
     * Returns true if the given object is an instance of Provider.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Provider {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Provider.__pulumiType;
    }


    /**
     * Create a Provider resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: ProviderArgs, opts?: pulumi.ResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        {
            resourceInputs["cluster"] = args ? args.cluster : undefined;
            resourceInputs["context"] = args ? args.context : undefined;
            resourceInputs["deleteUnreachable"] = pulumi.output(args ? args.deleteUnreachable : undefined).apply(JSON.stringify);
            resourceInputs["enableConfigMapMutable"] = pulumi.output((args ? args.enableConfigMapMutable : undefined) ?? utilities.getEnvBoolean("PULUMI_K8S_ENABLE_CONFIGMAP_MUTABLE")).apply(JSON.stringify);
            resourceInputs["enableServerSideApply"] = pulumi.output((args ? args.enableServerSideApply : undefined) ?? utilities.getEnvBoolean("PULUMI_K8S_ENABLE_SERVER_SIDE_APPLY")).apply(JSON.stringify);
            resourceInputs["helmReleaseSettings"] = pulumi.output(args ? (args.helmReleaseSettings ? pulumi.output(args.helmReleaseSettings).apply(inputs.helmReleaseSettingsProvideDefaults) : undefined) : undefined).apply(JSON.stringify);
            resourceInputs["kubeClientSettings"] = pulumi.output(args ? (args.kubeClientSettings ? pulumi.output(args.kubeClientSettings).apply(inputs.kubeClientSettingsProvideDefaults) : undefined) : undefined).apply(JSON.stringify);
            resourceInputs["kubeconfig"] = (args ? args.kubeconfig : undefined) ?? utilities.getEnv("KUBECONFIG");
            resourceInputs["namespace"] = args ? args.namespace : undefined;
            resourceInputs["renderYamlToDirectory"] = args ? args.renderYamlToDirectory : undefined;
            resourceInputs["suppressDeprecationWarnings"] = pulumi.output((args ? args.suppressDeprecationWarnings : undefined) ?? utilities.getEnvBoolean("PULUMI_K8S_SUPPRESS_DEPRECATION_WARNINGS")).apply(JSON.stringify);
            resourceInputs["suppressHelmHookWarnings"] = pulumi.output((args ? args.suppressHelmHookWarnings : undefined) ?? utilities.getEnvBoolean("PULUMI_K8S_SUPPRESS_HELM_HOOK_WARNINGS")).apply(JSON.stringify);
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Provider.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a Provider resource.
 */
export interface ProviderArgs {
    /**
     * If present, the name of the kubeconfig cluster to use.
     */
    cluster?: pulumi.Input<string>;
    /**
     * If present, the name of the kubeconfig context to use.
     */
    context?: pulumi.Input<string>;
    /**
     * If present and set to true, the provider will delete resources associated with an unreachable Kubernetes cluster from Pulumi state
     */
    deleteUnreachable?: pulumi.Input<boolean>;
    /**
     * BETA FEATURE - If present and set to true, allow ConfigMaps to be mutated.
     * This feature is in developer preview, and is disabled by default.
     *
     * This config can be specified in the following ways using this precedence:
     * 1. This `enableConfigMapMutable` parameter.
     * 2. The `PULUMI_K8S_ENABLE_CONFIGMAP_MUTABLE` environment variable.
     */
    enableConfigMapMutable?: pulumi.Input<boolean>;
    /**
     * If present and set to false, disable Server-Side Apply mode.
     * See https://github.com/pulumi/pulumi-kubernetes/issues/2011 for additional details.
     */
    enableServerSideApply?: pulumi.Input<boolean>;
    /**
     * Options to configure the Helm Release resource.
     */
    helmReleaseSettings?: pulumi.Input<inputs.HelmReleaseSettings>;
    /**
     * Options for tuning the Kubernetes client used by a Provider.
     */
    kubeClientSettings?: pulumi.Input<inputs.KubeClientSettings>;
    /**
     * The contents of a kubeconfig file or the path to a kubeconfig file.
     */
    kubeconfig?: pulumi.Input<string>;
    /**
     * If present, the default namespace to use. This flag is ignored for cluster-scoped resources.
     *
     * A namespace can be specified in multiple places, and the precedence is as follows:
     * 1. `.metadata.namespace` set on the resource.
     * 2. This `namespace` parameter.
     * 3. `namespace` set for the active context in the kubeconfig.
     */
    namespace?: pulumi.Input<string>;
    /**
     * BETA FEATURE - If present, render resource manifests to this directory. In this mode, resources will not
     * be created on a Kubernetes cluster, but the rendered manifests will be kept in sync with changes
     * to the Pulumi program. This feature is in developer preview, and is disabled by default.
     *
     * Note that some computed Outputs such as status fields will not be populated
     * since the resources are not created on a Kubernetes cluster. These Output values will remain undefined,
     * and may result in an error if they are referenced by other resources. Also note that any secret values
     * used in these resources will be rendered in plaintext to the resulting YAML.
     */
    renderYamlToDirectory?: pulumi.Input<string>;
    /**
     * If present and set to true, suppress apiVersion deprecation warnings from the CLI.
     */
    suppressDeprecationWarnings?: pulumi.Input<boolean>;
    /**
     * If present and set to true, suppress unsupported Helm hook warnings from the CLI.
     */
    suppressHelmHookWarnings?: pulumi.Input<boolean>;
}
