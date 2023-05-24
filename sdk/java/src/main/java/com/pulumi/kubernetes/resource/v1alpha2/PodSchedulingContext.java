// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.kubernetes.resource.v1alpha2;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.kubernetes.Utilities;
import com.pulumi.kubernetes.meta.v1.outputs.ObjectMeta;
import com.pulumi.kubernetes.resource.v1alpha2.PodSchedulingContextArgs;
import com.pulumi.kubernetes.resource.v1alpha2.outputs.PodSchedulingContextSpec;
import com.pulumi.kubernetes.resource.v1alpha2.outputs.PodSchedulingContextStatus;
import java.lang.String;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * PodSchedulingContext objects hold information that is needed to schedule a Pod with ResourceClaims that use &#34;WaitForFirstConsumer&#34; allocation mode.
 * 
 * This is an alpha type and requires enabling the DynamicResourceAllocation feature gate.
 * 
 */
@ResourceType(type="kubernetes:resource.k8s.io/v1alpha2:PodSchedulingContext")
public class PodSchedulingContext extends com.pulumi.resources.CustomResource {
    /**
     * APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
     * 
     */
    @Export(name="apiVersion", refs={String.class}, tree="[0]")
    private Output<String> apiVersion;

    /**
     * @return APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
     * 
     */
    public Output<String> apiVersion() {
        return this.apiVersion;
    }
    /**
     * Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
     * 
     */
    @Export(name="kind", refs={String.class}, tree="[0]")
    private Output<String> kind;

    /**
     * @return Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
     * 
     */
    public Output<String> kind() {
        return this.kind;
    }
    /**
     * Standard object metadata
     * 
     */
    @Export(name="metadata", refs={ObjectMeta.class}, tree="[0]")
    private Output<ObjectMeta> metadata;

    /**
     * @return Standard object metadata
     * 
     */
    public Output<ObjectMeta> metadata() {
        return this.metadata;
    }
    /**
     * Spec describes where resources for the Pod are needed.
     * 
     */
    @Export(name="spec", refs={PodSchedulingContextSpec.class}, tree="[0]")
    private Output<PodSchedulingContextSpec> spec;

    /**
     * @return Spec describes where resources for the Pod are needed.
     * 
     */
    public Output<PodSchedulingContextSpec> spec() {
        return this.spec;
    }
    /**
     * Status describes where resources for the Pod can be allocated.
     * 
     */
    @Export(name="status", refs={PodSchedulingContextStatus.class}, tree="[0]")
    private Output</* @Nullable */ PodSchedulingContextStatus> status;

    /**
     * @return Status describes where resources for the Pod can be allocated.
     * 
     */
    public Output<Optional<PodSchedulingContextStatus>> status() {
        return Codegen.optional(this.status);
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public PodSchedulingContext(String name) {
        this(name, PodSchedulingContextArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public PodSchedulingContext(String name, PodSchedulingContextArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public PodSchedulingContext(String name, PodSchedulingContextArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("kubernetes:resource.k8s.io/v1alpha2:PodSchedulingContext", name, makeArgs(args), makeResourceOptions(options, Codegen.empty()));
    }

    private PodSchedulingContext(String name, Output<String> id, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("kubernetes:resource.k8s.io/v1alpha2:PodSchedulingContext", name, null, makeResourceOptions(options, id));
    }

    private static PodSchedulingContextArgs makeArgs(PodSchedulingContextArgs args) {
        var builder = args == null ? PodSchedulingContextArgs.builder() : PodSchedulingContextArgs.builder(args);
        return builder
            .apiVersion("resource.k8s.io/v1alpha2")
            .kind("PodSchedulingContext")
            .build();
    }

    private static com.pulumi.resources.CustomResourceOptions makeResourceOptions(@Nullable com.pulumi.resources.CustomResourceOptions options, @Nullable Output<String> id) {
        var defaultOptions = com.pulumi.resources.CustomResourceOptions.builder()
            .version(Utilities.getVersion())
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
    public static PodSchedulingContext get(String name, Output<String> id, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new PodSchedulingContext(name, id, options);
    }
}
