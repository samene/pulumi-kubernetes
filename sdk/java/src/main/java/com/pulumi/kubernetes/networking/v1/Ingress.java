// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.kubernetes.networking.v1;

import com.pulumi.core.Alias;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.kubernetes.Utilities;
import com.pulumi.kubernetes.meta.v1.outputs.ObjectMeta;
import com.pulumi.kubernetes.networking.v1.IngressArgs;
import com.pulumi.kubernetes.networking.v1.outputs.IngressSpec;
import com.pulumi.kubernetes.networking.v1.outputs.IngressStatus;
import java.lang.String;
import java.util.List;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Ingress is a collection of rules that allow inbound connections to reach the endpoints defined by a backend. An Ingress can be configured to give services externally-reachable urls, load balance traffic, terminate SSL, offer name based virtual hosting etc.
 * 
 * This resource waits until its status is ready before registering success
 * for create/update, and populating output properties from the current state of the resource.
 * The following conditions are used to determine whether the resource creation has
 * succeeded or failed:
 * 
 * 1.  Ingress object exists.
 * 2.  Endpoint objects exist with matching names for each Ingress path (except when Service
 *     type is ExternalName).
 * 3.  Ingress entry exists for &#39;.status.loadBalancer.ingress&#39;.
 * 
 * If the Ingress has not reached a Ready state after 10 minutes, it will
 * time out and mark the resource update as Failed. You can override the default timeout value
 * by setting the &#39;customTimeouts&#39; option on the resource.
 * 
 * ## Example Usage
 * ### Create an Ingress with auto-naming
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.kubernetes.networking.k8s.io_v1.Ingress;
 * import com.pulumi.kubernetes.networking.k8s.io_v1.IngressArgs;
 * import com.pulumi.kubernetes.meta_v1.inputs.ObjectMetaArgs;
 * import com.pulumi.kubernetes.networking.k8s.io_v1.inputs.IngressSpecArgs;
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
 *         var ingress = new Ingress(&#34;ingress&#34;, IngressArgs.builder()        
 *             .metadata(ObjectMetaArgs.builder()
 *                 .annotations(Map.of(&#34;nginx.ingress.kubernetes.io/rewrite-target&#34;, &#34;/&#34;))
 *                 .build())
 *             .spec(IngressSpecArgs.builder()
 *                 .rules(IngressRuleArgs.builder()
 *                     .http(HTTPIngressRuleValueArgs.builder()
 *                         .paths(HTTPIngressPathArgs.builder()
 *                             .backend(IngressBackendArgs.builder()
 *                                 .service(IngressServiceBackendArgs.builder()
 *                                     .name(&#34;test&#34;)
 *                                     .port(ServiceBackendPortArgs.builder()
 *                                         .number(80)
 *                                         .build())
 *                                     .build())
 *                                 .build())
 *                             .path(&#34;/testpath&#34;)
 *                             .pathType(&#34;Prefix&#34;)
 *                             .build())
 *                         .build())
 *                     .build())
 *                 .build())
 *             .build());
 * 
 *     }
 * }
 * ```
 * ### Create an Ingress with a user-specified name
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.kubernetes.networking.k8s.io_v1.Ingress;
 * import com.pulumi.kubernetes.networking.k8s.io_v1.IngressArgs;
 * import com.pulumi.kubernetes.meta_v1.inputs.ObjectMetaArgs;
 * import com.pulumi.kubernetes.networking.k8s.io_v1.inputs.IngressSpecArgs;
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
 *         var ingress = new Ingress(&#34;ingress&#34;, IngressArgs.builder()        
 *             .metadata(ObjectMetaArgs.builder()
 *                 .annotations(Map.of(&#34;nginx.ingress.kubernetes.io/rewrite-target&#34;, &#34;/&#34;))
 *                 .name(&#34;minimal-ingress&#34;)
 *                 .build())
 *             .spec(IngressSpecArgs.builder()
 *                 .rules(IngressRuleArgs.builder()
 *                     .http(HTTPIngressRuleValueArgs.builder()
 *                         .paths(HTTPIngressPathArgs.builder()
 *                             .backend(IngressBackendArgs.builder()
 *                                 .service(IngressServiceBackendArgs.builder()
 *                                     .name(&#34;test&#34;)
 *                                     .port(ServiceBackendPortArgs.builder()
 *                                         .number(80)
 *                                         .build())
 *                                     .build())
 *                                 .build())
 *                             .path(&#34;/testpath&#34;)
 *                             .pathType(&#34;Prefix&#34;)
 *                             .build())
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
@ResourceType(type="kubernetes:networking.k8s.io/v1:Ingress")
public class Ingress extends com.pulumi.resources.CustomResource {
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
     * spec is the desired state of the Ingress. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#spec-and-status
     * 
     */
    @Export(name="spec", refs={IngressSpec.class}, tree="[0]")
    private Output<IngressSpec> spec;

    /**
     * @return spec is the desired state of the Ingress. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#spec-and-status
     * 
     */
    public Output<IngressSpec> spec() {
        return this.spec;
    }
    /**
     * status is the current state of the Ingress. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#spec-and-status
     * 
     */
    @Export(name="status", refs={IngressStatus.class}, tree="[0]")
    private Output</* @Nullable */ IngressStatus> status;

    /**
     * @return status is the current state of the Ingress. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#spec-and-status
     * 
     */
    public Output<Optional<IngressStatus>> status() {
        return Codegen.optional(this.status);
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public Ingress(String name) {
        this(name, IngressArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public Ingress(String name, @Nullable IngressArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public Ingress(String name, @Nullable IngressArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("kubernetes:networking.k8s.io/v1:Ingress", name, makeArgs(args), makeResourceOptions(options, Codegen.empty()));
    }

    private Ingress(String name, Output<String> id, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("kubernetes:networking.k8s.io/v1:Ingress", name, null, makeResourceOptions(options, id));
    }

    private static IngressArgs makeArgs(@Nullable IngressArgs args) {
        var builder = args == null ? IngressArgs.builder() : IngressArgs.builder(args);
        return builder
            .apiVersion("networking.k8s.io/v1")
            .kind("Ingress")
            .build();
    }

    private static com.pulumi.resources.CustomResourceOptions makeResourceOptions(@Nullable com.pulumi.resources.CustomResourceOptions options, @Nullable Output<String> id) {
        var defaultOptions = com.pulumi.resources.CustomResourceOptions.builder()
            .version(Utilities.getVersion())
            .aliases(List.of(
                Output.of(Alias.builder().type("kubernetes:extensions/v1beta1:Ingress").build()),
                Output.of(Alias.builder().type("kubernetes:networking.k8s.io/v1beta1:Ingress").build())
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
    public static Ingress get(String name, Output<String> id, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new Ingress(name, id, options);
    }
}
