// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.kubernetes.batch.v1beta1;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.kubernetes.Utilities;
import com.pulumi.kubernetes.batch.v1beta1.CronJobListArgs;
import com.pulumi.kubernetes.batch.v1beta1.outputs.CronJob;
import com.pulumi.kubernetes.meta.v1.outputs.ListMeta;
import java.lang.String;
import java.util.List;
import javax.annotation.Nullable;

/**
 * CronJobList is a collection of cron jobs.
 * 
 */
@ResourceType(type="kubernetes:batch/v1beta1:CronJobList")
public class CronJobList extends com.pulumi.resources.CustomResource {
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
     * items is the list of CronJobs.
     * 
     */
    @Export(name="items", refs={List.class,CronJob.class}, tree="[0,1]")
    private Output<List<CronJob>> items;

    /**
     * @return items is the list of CronJobs.
     * 
     */
    public Output<List<CronJob>> items() {
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
     * Standard list metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
     * 
     */
    @Export(name="metadata", refs={ListMeta.class}, tree="[0]")
    private Output<ListMeta> metadata;

    /**
     * @return Standard list metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
     * 
     */
    public Output<ListMeta> metadata() {
        return this.metadata;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public CronJobList(String name) {
        this(name, CronJobListArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public CronJobList(String name, CronJobListArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public CronJobList(String name, CronJobListArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("kubernetes:batch/v1beta1:CronJobList", name, makeArgs(args), makeResourceOptions(options, Codegen.empty()));
    }

    private CronJobList(String name, Output<String> id, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("kubernetes:batch/v1beta1:CronJobList", name, null, makeResourceOptions(options, id));
    }

    private static CronJobListArgs makeArgs(CronJobListArgs args) {
        var builder = args == null ? CronJobListArgs.builder() : CronJobListArgs.builder(args);
        return builder
            .apiVersion("batch/v1beta1")
            .kind("CronJobList")
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
    public static CronJobList get(String name, Output<String> id, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new CronJobList(name, id, options);
    }
}
