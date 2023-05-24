// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.kubernetes.admissionregistration.v1alpha1;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.kubernetes.Utilities;
import com.pulumi.kubernetes.admissionregistration.v1alpha1.ValidatingAdmissionPolicyListArgs;
import com.pulumi.kubernetes.admissionregistration.v1alpha1.outputs.ValidatingAdmissionPolicy;
import com.pulumi.kubernetes.meta.v1.outputs.ListMeta;
import java.lang.String;
import java.util.List;
import javax.annotation.Nullable;

/**
 * ValidatingAdmissionPolicyList is a list of ValidatingAdmissionPolicy.
 * 
 */
@ResourceType(type="kubernetes:admissionregistration.k8s.io/v1alpha1:ValidatingAdmissionPolicyList")
public class ValidatingAdmissionPolicyList extends com.pulumi.resources.CustomResource {
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
     * List of ValidatingAdmissionPolicy.
     * 
     */
    @Export(name="items", refs={List.class,ValidatingAdmissionPolicy.class}, tree="[0,1]")
    private Output<List<ValidatingAdmissionPolicy>> items;

    /**
     * @return List of ValidatingAdmissionPolicy.
     * 
     */
    public Output<List<ValidatingAdmissionPolicy>> items() {
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
     * Standard list metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
     * 
     */
    @Export(name="metadata", refs={ListMeta.class}, tree="[0]")
    private Output<ListMeta> metadata;

    /**
     * @return Standard list metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
     * 
     */
    public Output<ListMeta> metadata() {
        return this.metadata;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public ValidatingAdmissionPolicyList(String name) {
        this(name, ValidatingAdmissionPolicyListArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public ValidatingAdmissionPolicyList(String name, @Nullable ValidatingAdmissionPolicyListArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public ValidatingAdmissionPolicyList(String name, @Nullable ValidatingAdmissionPolicyListArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("kubernetes:admissionregistration.k8s.io/v1alpha1:ValidatingAdmissionPolicyList", name, makeArgs(args), makeResourceOptions(options, Codegen.empty()));
    }

    private ValidatingAdmissionPolicyList(String name, Output<String> id, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("kubernetes:admissionregistration.k8s.io/v1alpha1:ValidatingAdmissionPolicyList", name, null, makeResourceOptions(options, id));
    }

    private static ValidatingAdmissionPolicyListArgs makeArgs(@Nullable ValidatingAdmissionPolicyListArgs args) {
        var builder = args == null ? ValidatingAdmissionPolicyListArgs.builder() : ValidatingAdmissionPolicyListArgs.builder(args);
        return builder
            .apiVersion("admissionregistration.k8s.io/v1alpha1")
            .kind("ValidatingAdmissionPolicyList")
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
    public static ValidatingAdmissionPolicyList get(String name, Output<String> id, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new ValidatingAdmissionPolicyList(name, id, options);
    }
}
