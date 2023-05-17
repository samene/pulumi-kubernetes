// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.kubernetes;

import com.pulumi.core.internal.Codegen;
import java.lang.Boolean;
import java.lang.String;
import java.util.Optional;

public final class Config {

    private static final com.pulumi.Config config = com.pulumi.Config.of("kubernetes");
/**
 * If present, the name of the kubeconfig cluster to use.
 * 
 */
    public Optional<String> cluster() {
        return Codegen.stringProp("cluster").config(config).get();
    }
/**
 * If present, the name of the kubeconfig context to use.
 * 
 */
    public Optional<String> context() {
        return Codegen.stringProp("context").config(config).get();
    }
/**
 * If present and set to true, the provider will delete resources associated with an unreachable Kubernetes cluster from Pulumi state
 * 
 */
    public Optional<Boolean> deleteUnreachable() {
        return Codegen.booleanProp("deleteUnreachable").config(config).get();
    }
/**
 * BETA FEATURE - If present and set to true, allow ConfigMaps to be mutated.
 * This feature is in developer preview, and is disabled by default.
 * 
 * This config can be specified in the following ways using this precedence:
 * 1. This `enableConfigMapMutable` parameter.
 * 2. The `PULUMI_K8S_ENABLE_CONFIGMAP_MUTABLE` environment variable.
 * 
 */
    public Optional<Boolean> enableConfigMapMutable() {
        return Codegen.booleanProp("enableConfigMapMutable").config(config).get();
    }
/**
 * Obsolete. This option has no effect.
 * 
 */
    public Optional<Boolean> enableReplaceCRD() {
        return Codegen.booleanProp("enableReplaceCRD").config(config).get();
    }
/**
 * BETA FEATURE - If present and set to true, enable Server-Side Apply mode.
 * See https://github.com/pulumi/pulumi-kubernetes/issues/2011 for additional details.
 * This feature is in developer preview, and is disabled by default.
 * 
 */
    public Optional<Boolean> enableServerSideApply() {
        return Codegen.booleanProp("enableServerSideApply").config(config).get();
    }
/**
 * The contents of a kubeconfig file or the path to a kubeconfig file. If this is set, this config will be used instead of $KUBECONFIG.
 * 
 */
    public Optional<String> kubeconfig() {
        return Codegen.stringProp("kubeconfig").config(config).get();
    }
/**
 * If present, the default namespace to use. This flag is ignored for cluster-scoped resources.
 * 
 * A namespace can be specified in multiple places, and the precedence is as follows:
 * 1. `.metadata.namespace` set on the resource.
 * 2. This `namespace` parameter.
 * 3. `namespace` set for the active context in the kubeconfig.
 * 
 */
    public Optional<String> namespace() {
        return Codegen.stringProp("namespace").config(config).get();
    }
/**
 * BETA FEATURE - If present, render resource manifests to this directory. In this mode, resources will not
 * be created on a Kubernetes cluster, but the rendered manifests will be kept in sync with changes
 * to the Pulumi program. This feature is in developer preview, and is disabled by default.
 * 
 * Note that some computed Outputs such as status fields will not be populated
 * since the resources are not created on a Kubernetes cluster. These Output values will remain undefined,
 * and may result in an error if they are referenced by other resources. Also note that any secret values
 * used in these resources will be rendered in plaintext to the resulting YAML.
 * 
 */
    public Optional<String> renderYamlToDirectory() {
        return Codegen.stringProp("renderYamlToDirectory").config(config).get();
    }
/**
 * If present and set to true, the provider will use strict configuration mode. Recommended for production stacks. In this mode, the default Kubernetes provider is disabled, and the `kubeconfig` and `context` settings are required for Provider configuration. These settings unambiguously ensure that every Kubernetes resource is associated with a particular cluster.
 * 
 */
    public Optional<Boolean> strictMode() {
        return Codegen.booleanProp("strictMode").config(config).get();
    }
/**
 * If present and set to true, suppress apiVersion deprecation warnings from the CLI.
 * 
 * This config can be specified in the following ways, using this precedence:
 * 1. This `suppressDeprecationWarnings` parameter.
 * 2. The `PULUMI_K8S_SUPPRESS_DEPRECATION_WARNINGS` environment variable.
 * 
 */
    public Optional<Boolean> suppressDeprecationWarnings() {
        return Codegen.booleanProp("suppressDeprecationWarnings").config(config).get();
    }
/**
 * If present and set to true, suppress unsupported Helm hook warnings from the CLI.
 * 
 * This config can be specified in the following ways, using this precedence:
 * 1. This `suppressHelmHookWarnings` parameter.
 * 2. The `PULUMI_K8S_SUPPRESS_HELM_HOOK_WARNINGS` environment variable.
 * 
 */
    public Optional<Boolean> suppressHelmHookWarnings() {
        return Codegen.booleanProp("suppressHelmHookWarnings").config(config).get();
    }
}
