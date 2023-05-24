// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.kubernetes.storage.v1alpha1;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.kubernetes.Utilities;
import com.pulumi.kubernetes.meta.v1.outputs.ListMeta;
import com.pulumi.kubernetes.storage.v1alpha1.VolumeAttachmentListArgs;
import com.pulumi.kubernetes.storage.v1alpha1.outputs.VolumeAttachment;
import java.lang.String;
import java.util.List;
import javax.annotation.Nullable;

/**
 * VolumeAttachmentList is a collection of VolumeAttachment objects.
 * 
 */
@ResourceType(type="kubernetes:storage.k8s.io/v1alpha1:VolumeAttachmentList")
public class VolumeAttachmentList extends com.pulumi.resources.CustomResource {
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
     * Items is the list of VolumeAttachments
     * 
     */
    @Export(name="items", refs={List.class,VolumeAttachment.class}, tree="[0,1]")
    private Output<List<VolumeAttachment>> items;

    /**
     * @return Items is the list of VolumeAttachments
     * 
     */
    public Output<List<VolumeAttachment>> items() {
        return this.items;
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
     * Standard list metadata More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
     * 
     */
    @Export(name="metadata", refs={ListMeta.class}, tree="[0]")
    private Output<ListMeta> metadata;

    /**
     * @return Standard list metadata More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
     * 
     */
    public Output<ListMeta> metadata() {
        return this.metadata;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public VolumeAttachmentList(String name) {
        this(name, VolumeAttachmentListArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public VolumeAttachmentList(String name, VolumeAttachmentListArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public VolumeAttachmentList(String name, VolumeAttachmentListArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("kubernetes:storage.k8s.io/v1alpha1:VolumeAttachmentList", name, makeArgs(args), makeResourceOptions(options, Codegen.empty()));
    }

    private VolumeAttachmentList(String name, Output<String> id, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("kubernetes:storage.k8s.io/v1alpha1:VolumeAttachmentList", name, null, makeResourceOptions(options, id));
    }

    private static VolumeAttachmentListArgs makeArgs(VolumeAttachmentListArgs args) {
        var builder = args == null ? VolumeAttachmentListArgs.builder() : VolumeAttachmentListArgs.builder(args);
        return builder
            .apiVersion("storage.k8s.io/v1alpha1")
            .kind("VolumeAttachmentList")
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
    public static VolumeAttachmentList get(String name, Output<String> id, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new VolumeAttachmentList(name, id, options);
    }
}
