// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.kubernetes.autoscaling.v2;

import com.pulumi.core.Alias;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.kubernetes.Utilities;
import com.pulumi.kubernetes.autoscaling.v2.HorizontalPodAutoscalerArgs;
import com.pulumi.kubernetes.autoscaling.v2.outputs.HorizontalPodAutoscalerSpec;
import com.pulumi.kubernetes.autoscaling.v2.outputs.HorizontalPodAutoscalerStatus;
import com.pulumi.kubernetes.meta.v1.outputs.ObjectMeta;
import java.lang.String;
import java.util.List;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * HorizontalPodAutoscaler is the configuration for a horizontal pod autoscaler, which automatically manages the replica count of any resource implementing the scale subresource based on the metrics specified.
 * 
 */
@ResourceType(type="kubernetes:autoscaling/v2:HorizontalPodAutoscaler")
public class HorizontalPodAutoscaler extends com.pulumi.resources.CustomResource {
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
     * metadata is the standard object metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
     * 
     */
    @Export(name="metadata", refs={ObjectMeta.class}, tree="[0]")
    private Output<ObjectMeta> metadata;

    /**
     * @return metadata is the standard object metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
     * 
     */
    public Output<ObjectMeta> metadata() {
        return this.metadata;
    }
    /**
     * spec is the specification for the behaviour of the autoscaler. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#spec-and-status.
     * 
     */
    @Export(name="spec", refs={HorizontalPodAutoscalerSpec.class}, tree="[0]")
    private Output<HorizontalPodAutoscalerSpec> spec;

    /**
     * @return spec is the specification for the behaviour of the autoscaler. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#spec-and-status.
     * 
     */
    public Output<HorizontalPodAutoscalerSpec> spec() {
        return this.spec;
    }
    /**
     * status is the current information about the autoscaler.
     * 
     */
    @Export(name="status", refs={HorizontalPodAutoscalerStatus.class}, tree="[0]")
    private Output</* @Nullable */ HorizontalPodAutoscalerStatus> status;

    /**
     * @return status is the current information about the autoscaler.
     * 
     */
    public Output<Optional<HorizontalPodAutoscalerStatus>> status() {
        return Codegen.optional(this.status);
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public HorizontalPodAutoscaler(String name) {
        this(name, HorizontalPodAutoscalerArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public HorizontalPodAutoscaler(String name, @Nullable HorizontalPodAutoscalerArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public HorizontalPodAutoscaler(String name, @Nullable HorizontalPodAutoscalerArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("kubernetes:autoscaling/v2:HorizontalPodAutoscaler", name, makeArgs(args), makeResourceOptions(options, Codegen.empty()));
    }

    private HorizontalPodAutoscaler(String name, Output<String> id, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("kubernetes:autoscaling/v2:HorizontalPodAutoscaler", name, null, makeResourceOptions(options, id));
    }

    private static HorizontalPodAutoscalerArgs makeArgs(@Nullable HorizontalPodAutoscalerArgs args) {
        var builder = args == null ? HorizontalPodAutoscalerArgs.builder() : HorizontalPodAutoscalerArgs.builder(args);
        return builder
            .apiVersion("autoscaling/v2")
            .kind("HorizontalPodAutoscaler")
            .build();
    }

    private static com.pulumi.resources.CustomResourceOptions makeResourceOptions(@Nullable com.pulumi.resources.CustomResourceOptions options, @Nullable Output<String> id) {
        var defaultOptions = com.pulumi.resources.CustomResourceOptions.builder()
            .version(Utilities.getVersion())
            .aliases(List.of(
                Output.of(Alias.builder().type("kubernetes:autoscaling/v1:HorizontalPodAutoscaler").build()),
                Output.of(Alias.builder().type("kubernetes:autoscaling/v2beta1:HorizontalPodAutoscaler").build()),
                Output.of(Alias.builder().type("kubernetes:autoscaling/v2beta2:HorizontalPodAutoscaler").build())
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
    public static HorizontalPodAutoscaler get(String name, Output<String> id, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new HorizontalPodAutoscaler(name, id, options);
    }
}
