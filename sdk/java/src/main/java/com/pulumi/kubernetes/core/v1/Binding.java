// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.kubernetes.core.v1;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.kubernetes.Utilities;
import com.pulumi.kubernetes.core.v1.BindingArgs;
import com.pulumi.kubernetes.core.v1.outputs.ObjectReference;
import com.pulumi.kubernetes.meta.v1.outputs.ObjectMeta;
import java.lang.String;
import javax.annotation.Nullable;

/**
 * Binding ties one object to another; for example, a pod is bound to a node by a scheduler. Deprecated in 1.7, please use the bindings subresource of pods instead.
 * 
 */
@ResourceType(type="kubernetes:core/v1:Binding")
public class Binding extends com.pulumi.resources.CustomResource {
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
     * Standard object&#39;s metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
     * 
     */
    @Export(name="metadata", refs={ObjectMeta.class}, tree="[0]")
    private Output<ObjectMeta> metadata;

    /**
     * @return Standard object&#39;s metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
     * 
     */
    public Output<ObjectMeta> metadata() {
        return this.metadata;
    }
    /**
     * The target object that you want to bind to the standard object.
     * 
     */
    @Export(name="target", refs={ObjectReference.class}, tree="[0]")
    private Output<ObjectReference> target;

    /**
     * @return The target object that you want to bind to the standard object.
     * 
     */
    public Output<ObjectReference> target() {
        return this.target;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public Binding(String name) {
        this(name, BindingArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public Binding(String name, BindingArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public Binding(String name, BindingArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("kubernetes:core/v1:Binding", name, makeArgs(args), makeResourceOptions(options, Codegen.empty()));
    }

    private Binding(String name, Output<String> id, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("kubernetes:core/v1:Binding", name, null, makeResourceOptions(options, id));
    }

    private static BindingArgs makeArgs(BindingArgs args) {
        var builder = args == null ? BindingArgs.builder() : BindingArgs.builder(args);
        return builder
            .apiVersion("v1")
            .kind("Binding")
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
    public static Binding get(String name, Output<String> id, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new Binding(name, id, options);
    }
}
