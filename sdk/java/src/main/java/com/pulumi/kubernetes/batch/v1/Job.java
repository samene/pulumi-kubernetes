// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.kubernetes.batch.v1;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.kubernetes.Utilities;
import com.pulumi.kubernetes.batch.v1.JobArgs;
import com.pulumi.kubernetes.batch.v1.outputs.JobSpec;
import com.pulumi.kubernetes.batch.v1.outputs.JobStatus;
import com.pulumi.kubernetes.meta.v1.outputs.ObjectMeta;
import java.lang.String;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Job represents the configuration of a single job.
 * 
 * This resource waits until its status is ready before registering success
 * for create/update, and populating output properties from the current state of the resource.
 * The following conditions are used to determine whether the resource creation has
 * succeeded or failed:
 * 
 * 1. The Job&#39;s &#39;.status.startTime&#39; is set, which indicates that the Job has started running.
 * 2. The Job&#39;s &#39;.status.conditions&#39; has a status of type &#39;Complete&#39;, and a &#39;status&#39; set
 *    to &#39;True&#39;.
 * 3. The Job&#39;s &#39;.status.conditions&#39; do not have a status of type &#39;Failed&#39;, with a
 *     &#39;status&#39; set to &#39;True&#39;. If this condition is set, we should fail the Job immediately.
 * 
 * If the Job has not reached a Ready state after 10 minutes, it will
 * time out and mark the resource update as Failed. You can override the default timeout value
 * by setting the &#39;customTimeouts&#39; option on the resource.
 * 
 * By default, if a resource failed to become ready in a previous update,
 * Pulumi will continue to wait for readiness on the next update. If you would prefer
 * to schedule a replacement for an unready resource on the next update, you can add the
 * &#34;pulumi.com/replaceUnready&#34;: &#34;true&#34; annotation to the resource definition.
 * 
 * ## Example Usage
 * ### Create a Job with auto-naming
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.kubernetes.batch_v1.Job;
 * import com.pulumi.kubernetes.batch_v1.JobArgs;
 * import com.pulumi.kubernetes.batch_v1.inputs.JobSpecArgs;
 * import com.pulumi.kubernetes.core_v1.inputs.PodTemplateSpecArgs;
 * import com.pulumi.kubernetes.core_v1.inputs.PodSpecArgs;
 * import java.util.List;
 * import java.util.ArrayList;
 * import java.util.Map;
 * import java.io.File;
 * import java.nio.file.Files;
 * import java.nio.file.Paths;
 * 
 * public class App {
 *     public static void main(String[] args) {
 *         Pulumi.run(App::stack);
 *     }
 * 
 *     public static void stack(Context ctx) {
 *         var job = new Job(&#34;job&#34;, JobArgs.builder()        
 *             .metadata(null)
 *             .spec(JobSpecArgs.builder()
 *                 .backoffLimit(4)
 *                 .template(PodTemplateSpecArgs.builder()
 *                     .spec(PodSpecArgs.builder()
 *                         .containers(ContainerArgs.builder()
 *                             .command(                            
 *                                 &#34;perl&#34;,
 *                                 &#34;-Mbignum=bpi&#34;,
 *                                 &#34;-wle&#34;,
 *                                 &#34;print bpi(2000)&#34;)
 *                             .image(&#34;perl&#34;)
 *                             .name(&#34;pi&#34;)
 *                             .build())
 *                         .restartPolicy(&#34;Never&#34;)
 *                         .build())
 *                     .build())
 *                 .build())
 *             .build());
 * 
 *     }
 * }
 * ```
 * ### Create a Job with a user-specified name
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.kubernetes.batch_v1.Job;
 * import com.pulumi.kubernetes.batch_v1.JobArgs;
 * import com.pulumi.kubernetes.meta_v1.inputs.ObjectMetaArgs;
 * import com.pulumi.kubernetes.batch_v1.inputs.JobSpecArgs;
 * import com.pulumi.kubernetes.core_v1.inputs.PodTemplateSpecArgs;
 * import com.pulumi.kubernetes.core_v1.inputs.PodSpecArgs;
 * import java.util.List;
 * import java.util.ArrayList;
 * import java.util.Map;
 * import java.io.File;
 * import java.nio.file.Files;
 * import java.nio.file.Paths;
 * 
 * public class App {
 *     public static void main(String[] args) {
 *         Pulumi.run(App::stack);
 *     }
 * 
 *     public static void stack(Context ctx) {
 *         var job = new Job(&#34;job&#34;, JobArgs.builder()        
 *             .metadata(ObjectMetaArgs.builder()
 *                 .name(&#34;pi&#34;)
 *                 .build())
 *             .spec(JobSpecArgs.builder()
 *                 .backoffLimit(4)
 *                 .template(PodTemplateSpecArgs.builder()
 *                     .spec(PodSpecArgs.builder()
 *                         .containers(ContainerArgs.builder()
 *                             .command(                            
 *                                 &#34;perl&#34;,
 *                                 &#34;-Mbignum=bpi&#34;,
 *                                 &#34;-wle&#34;,
 *                                 &#34;print bpi(2000)&#34;)
 *                             .image(&#34;perl&#34;)
 *                             .name(&#34;pi&#34;)
 *                             .build())
 *                         .restartPolicy(&#34;Never&#34;)
 *                         .build())
 *                     .build())
 *                 .build())
 *             .build());
 * 
 *     }
 * }
 * ```
 * 
 */
@ResourceType(type="kubernetes:batch/v1:Job")
public class Job extends com.pulumi.resources.CustomResource {
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
     * Standard object&#39;s metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
     * 
     */
    @Export(name="metadata", refs={ObjectMeta.class}, tree="[0]")
    private Output</* @Nullable */ ObjectMeta> metadata;

    /**
     * @return Standard object&#39;s metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
     * 
     */
    public Output<Optional<ObjectMeta>> metadata() {
        return Codegen.optional(this.metadata);
    }
    /**
     * Specification of the desired behavior of a job. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#spec-and-status
     * 
     */
    @Export(name="spec", refs={JobSpec.class}, tree="[0]")
    private Output</* @Nullable */ JobSpec> spec;

    /**
     * @return Specification of the desired behavior of a job. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#spec-and-status
     * 
     */
    public Output<Optional<JobSpec>> spec() {
        return Codegen.optional(this.spec);
    }
    /**
     * Current status of a job. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#spec-and-status
     * 
     */
    @Export(name="status", refs={JobStatus.class}, tree="[0]")
    private Output</* @Nullable */ JobStatus> status;

    /**
     * @return Current status of a job. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#spec-and-status
     * 
     */
    public Output<Optional<JobStatus>> status() {
        return Codegen.optional(this.status);
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public Job(String name) {
        this(name, JobArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public Job(String name, @Nullable JobArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public Job(String name, @Nullable JobArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("kubernetes:batch/v1:Job", name, makeArgs(args), makeResourceOptions(options, Codegen.empty()));
    }

    private Job(String name, Output<String> id, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("kubernetes:batch/v1:Job", name, null, makeResourceOptions(options, id));
    }

    private static JobArgs makeArgs(@Nullable JobArgs args) {
        var builder = args == null ? JobArgs.builder() : JobArgs.builder(args);
        return builder
            .apiVersion("batch/v1")
            .kind("Job")
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
    public static Job get(String name, Output<String> id, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new Job(name, id, options);
    }
}
