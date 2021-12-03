// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../../types";
import * as utilities from "../../utilities";

/**
 * StatefulSet represents a set of pods with consistent identities. Identities are defined as:
 *  - Network: A single stable DNS and hostname.
 *  - Storage: As many VolumeClaims as requested.
 *    The StatefulSet guarantees that a given network identity will always map to the same storage identity.
 *
 * This resource waits until its status is ready before registering success
 * for create/update, and populating output properties from the current state of the resource.
 * The following conditions are used to determine whether the resource creation has
 * succeeded or failed:
 *
 * 1. The value of 'spec.replicas' matches '.status.replicas', '.status.currentReplicas',
 *    and '.status.readyReplicas'.
 * 2. The value of '.status.updateRevision' matches '.status.currentRevision'.
 *
 * If the StatefulSet has not reached a Ready state after 10 minutes, it will
 * time out and mark the resource update as Failed. You can override the default timeout value
 * by setting the 'customTimeouts' option on the resource.
 *
 * @deprecated apps/v1beta2/StatefulSet is deprecated by apps/v1/StatefulSet and not supported by Kubernetes v1.16+ clusters.
 */
export class StatefulSet extends pulumi.CustomResource {
    /**
     * Get an existing StatefulSet resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): StatefulSet {
        return new StatefulSet(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'kubernetes:apps/v1beta2:StatefulSet';

    /**
     * Returns true if the given object is an instance of StatefulSet.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is StatefulSet {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === StatefulSet.__pulumiType;
    }

    /**
     * APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
     */
    public readonly apiVersion!: pulumi.Output<"apps/v1beta2">;
    /**
     * Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
     */
    public readonly kind!: pulumi.Output<"StatefulSet">;
    public readonly metadata!: pulumi.Output<outputs.meta.v1.ObjectMeta>;
    /**
     * Spec defines the desired identities of pods in this set.
     */
    public readonly spec!: pulumi.Output<outputs.apps.v1beta2.StatefulSetSpec>;
    /**
     * Status is the current status of Pods in this StatefulSet. This data may be out of date by some window of time.
     */
    public /*out*/ readonly status!: pulumi.Output<outputs.apps.v1beta2.StatefulSetStatus>;

    /**
     * Create a StatefulSet resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    /** @deprecated apps/v1beta2/StatefulSet is deprecated by apps/v1/StatefulSet and not supported by Kubernetes v1.16+ clusters. */
    constructor(name: string, args?: StatefulSetArgs, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            resourceInputs["apiVersion"] = "apps/v1beta2";
            resourceInputs["kind"] = "StatefulSet";
            resourceInputs["metadata"] = args ? args.metadata : undefined;
            resourceInputs["spec"] = args ? args.spec : undefined;
            resourceInputs["status"] = undefined /*out*/;
        } else {
            resourceInputs["apiVersion"] = undefined /*out*/;
            resourceInputs["kind"] = undefined /*out*/;
            resourceInputs["metadata"] = undefined /*out*/;
            resourceInputs["spec"] = undefined /*out*/;
            resourceInputs["status"] = undefined /*out*/;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        const aliasOpts = { aliases: [{ type: "kubernetes:apps/v1:StatefulSet" }, { type: "kubernetes:apps/v1beta1:StatefulSet" }] };
        opts = pulumi.mergeOptions(opts, aliasOpts);
        super(StatefulSet.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a StatefulSet resource.
 */
export interface StatefulSetArgs {
    /**
     * APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
     */
    apiVersion?: pulumi.Input<"apps/v1beta2">;
    /**
     * Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
     */
    kind?: pulumi.Input<"StatefulSet">;
    metadata?: pulumi.Input<inputs.meta.v1.ObjectMeta>;
    /**
     * Spec defines the desired identities of pods in this set.
     */
    spec?: pulumi.Input<inputs.apps.v1beta2.StatefulSetSpec>;
}
