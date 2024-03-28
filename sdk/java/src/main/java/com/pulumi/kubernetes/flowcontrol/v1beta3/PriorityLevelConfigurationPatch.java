// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.kubernetes.flowcontrol.v1beta3;

import com.pulumi.core.Alias;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.kubernetes.Utilities;
import com.pulumi.kubernetes.flowcontrol.v1beta3.PriorityLevelConfigurationPatchArgs;
import com.pulumi.kubernetes.flowcontrol.v1beta3.outputs.PriorityLevelConfigurationSpecPatch;
import com.pulumi.kubernetes.flowcontrol.v1beta3.outputs.PriorityLevelConfigurationStatusPatch;
import com.pulumi.kubernetes.meta.v1.outputs.ObjectMetaPatch;
import java.lang.String;
import java.util.List;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Patch resources are used to modify existing Kubernetes resources by using
 * Server-Side Apply updates. The name of the resource must be specified, but all other properties are optional. More than
 * one patch may be applied to the same resource, and a random FieldManager name will be used for each Patch resource.
 * Conflicts will result in an error by default, but can be forced using the &#34;pulumi.com/patchForce&#34; annotation. See the
 * [Server-Side Apply Docs](https://www.pulumi.com/registry/packages/kubernetes/how-to-guides/managing-resources-with-server-side-apply/) for
 * additional information about using Server-Side Apply to manage Kubernetes resources with Pulumi.
 * PriorityLevelConfiguration represents the configuration of a priority level.
 * 
 */
@ResourceType(type="kubernetes:flowcontrol.apiserver.k8s.io/v1beta3:PriorityLevelConfigurationPatch")
public class PriorityLevelConfigurationPatch extends com.pulumi.resources.CustomResource {
    /**
     * APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
     * 
     */
    @Export(name="apiVersion", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> apiVersion;

    /**
     * @return APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
     * 
     */
    public Output<Optional<String>> apiVersion() {
        return Codegen.optional(this.apiVersion);
    }
    /**
     * Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
     * 
     */
    @Export(name="kind", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> kind;

    /**
     * @return Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
     * 
     */
    public Output<Optional<String>> kind() {
        return Codegen.optional(this.kind);
    }
    /**
     * `metadata` is the standard object&#39;s metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
     * 
     */
    @Export(name="metadata", refs={ObjectMetaPatch.class}, tree="[0]")
    private Output</* @Nullable */ ObjectMetaPatch> metadata;

    /**
     * @return `metadata` is the standard object&#39;s metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
     * 
     */
    public Output<Optional<ObjectMetaPatch>> metadata() {
        return Codegen.optional(this.metadata);
    }
    /**
     * `spec` is the specification of the desired behavior of a &#34;request-priority&#34;. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#spec-and-status
     * 
     */
    @Export(name="spec", refs={PriorityLevelConfigurationSpecPatch.class}, tree="[0]")
    private Output</* @Nullable */ PriorityLevelConfigurationSpecPatch> spec;

    /**
     * @return `spec` is the specification of the desired behavior of a &#34;request-priority&#34;. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#spec-and-status
     * 
     */
    public Output<Optional<PriorityLevelConfigurationSpecPatch>> spec() {
        return Codegen.optional(this.spec);
    }
    /**
     * `status` is the current status of a &#34;request-priority&#34;. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#spec-and-status
     * 
     */
    @Export(name="status", refs={PriorityLevelConfigurationStatusPatch.class}, tree="[0]")
    private Output</* @Nullable */ PriorityLevelConfigurationStatusPatch> status;

    /**
     * @return `status` is the current status of a &#34;request-priority&#34;. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#spec-and-status
     * 
     */
    public Output<Optional<PriorityLevelConfigurationStatusPatch>> status() {
        return Codegen.optional(this.status);
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public PriorityLevelConfigurationPatch(String name) {
        this(name, PriorityLevelConfigurationPatchArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public PriorityLevelConfigurationPatch(String name, @Nullable PriorityLevelConfigurationPatchArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public PriorityLevelConfigurationPatch(String name, @Nullable PriorityLevelConfigurationPatchArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("kubernetes:flowcontrol.apiserver.k8s.io/v1beta3:PriorityLevelConfigurationPatch", name, makeArgs(args), makeResourceOptions(options, Codegen.empty()));
    }

    private PriorityLevelConfigurationPatch(String name, Output<String> id, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("kubernetes:flowcontrol.apiserver.k8s.io/v1beta3:PriorityLevelConfigurationPatch", name, null, makeResourceOptions(options, id));
    }

    private static PriorityLevelConfigurationPatchArgs makeArgs(@Nullable PriorityLevelConfigurationPatchArgs args) {
        var builder = args == null ? PriorityLevelConfigurationPatchArgs.builder() : PriorityLevelConfigurationPatchArgs.builder(args);
        return builder
            .apiVersion("flowcontrol.apiserver.k8s.io/v1beta3")
            .kind("PriorityLevelConfiguration")
            .build();
    }

    private static com.pulumi.resources.CustomResourceOptions makeResourceOptions(@Nullable com.pulumi.resources.CustomResourceOptions options, @Nullable Output<String> id) {
        var defaultOptions = com.pulumi.resources.CustomResourceOptions.builder()
            .version(Utilities.getVersion())
            .aliases(List.of(
                Output.of(Alias.builder().type("kubernetes:flowcontrol.apiserver.k8s.io/v1:PriorityLevelConfigurationPatch").build()),
                Output.of(Alias.builder().type("kubernetes:flowcontrol.apiserver.k8s.io/v1alpha1:PriorityLevelConfigurationPatch").build()),
                Output.of(Alias.builder().type("kubernetes:flowcontrol.apiserver.k8s.io/v1beta1:PriorityLevelConfigurationPatch").build()),
                Output.of(Alias.builder().type("kubernetes:flowcontrol.apiserver.k8s.io/v1beta2:PriorityLevelConfigurationPatch").build())
            ))
            .build();
        return com.pulumi.resources.CustomResourceOptions.merge(defaultOptions, options, id);
    }

    /**
     * Get an existing Host resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param options Optional settings to control the behavior of the CustomResource.
     */
    public static PriorityLevelConfigurationPatch get(String name, Output<String> id, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new PriorityLevelConfigurationPatch(name, id, options);
    }
}
