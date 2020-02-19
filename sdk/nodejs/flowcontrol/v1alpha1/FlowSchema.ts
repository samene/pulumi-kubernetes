// *** WARNING: this file was generated by the Pulumi Kubernetes codegen tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { core } from "../..";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import { getVersion } from "../../version";

    /**
     * FlowSchema defines the schema of a group of flows. Note that a flow is made up of a set of
     * inbound API requests with similar attributes and is identified by a pair of strings: the name
     * of the FlowSchema and a "flow distinguisher".
     */
    export class FlowSchema extends pulumi.CustomResource {
      /**
       * APIVersion defines the versioned schema of this representation of an object. Servers should
       * convert recognized schemas to the latest internal value, and may reject unrecognized
       * values. More info:
       * https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
       */
      public readonly apiVersion: pulumi.Output<"flowcontrol.apiserver.k8s.io/v1alpha1">;

      /**
       * Kind is a string value representing the REST resource this object represents. Servers may
       * infer this from the endpoint the client submits requests to. Cannot be updated. In
       * CamelCase. More info:
       * https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
       */
      public readonly kind: pulumi.Output<"FlowSchema">;

      /**
       * `metadata` is the standard object's metadata. More info:
       * https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
       */
      public readonly metadata: pulumi.Output<outputs.meta.v1.ObjectMeta>;

      /**
       * `spec` is the specification of the desired behavior of a FlowSchema. More info:
       * https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#spec-and-status
       */
      public readonly spec: pulumi.Output<outputs.flowcontrol.v1alpha1.FlowSchemaSpec>;

      /**
       * `status` is the current status of a FlowSchema. More info:
       * https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#spec-and-status
       */
      public readonly status: pulumi.Output<outputs.flowcontrol.v1alpha1.FlowSchemaStatus>;

      /**
       * Get the state of an existing `FlowSchema` resource, as identified by `id`.
       * The ID is of the form `[namespace]/<name>`; if `namespace` is omitted, then (per
       * Kubernetes convention) the ID becomes `default/<name>`.
       *
       * Pulumi will keep track of this resource using `name` as the Pulumi ID.
       *
       * @param name _Unique_ name used to register this resource with Pulumi.
       * @param id An ID for the Kubernetes resource to retrieve. Takes the form `[namespace]/<name>`.
       * @param opts Uniquely specifies a CustomResource to select.
       */
      public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): FlowSchema {
          return new FlowSchema(name, undefined, { ...opts, id: id });
      }

      /** @internal */
      private static readonly __pulumiType = "kubernetes:flowcontrol.apiserver.k8s.io/v1alpha1:FlowSchema";

      /**
       * Returns true if the given object is an instance of FlowSchema.  This is designed to work even
       * when multiple copies of the Pulumi SDK have been loaded into the same process.
       */
      public static isInstance(obj: any): obj is FlowSchema {
          if (obj === undefined || obj === null) {
              return false;
          }

          return obj["__pulumiType"] === FlowSchema.__pulumiType;
      }

      /**
       * Create a flowcontrol.v1alpha1.FlowSchema resource with the given unique name, arguments, and options.
       *
       * @param name The _unique_ name of the resource.
       * @param args The arguments to use to populate this resource's properties.
       * @param opts A bag of options that control this resource's behavior.
       */
      constructor(name: string, args?: inputs.flowcontrol.v1alpha1.FlowSchema, opts?: pulumi.CustomResourceOptions) {
          const props: pulumi.Inputs = {};

          props["apiVersion"] = "flowcontrol.apiserver.k8s.io/v1alpha1";
          props["kind"] = "FlowSchema";
          props["metadata"] = args?.metadata;
          props["spec"] = args?.spec;

          props["status"] = undefined;

          if (!opts) {
              opts = {};
          }

          if (!opts.version) {
              opts.version = getVersion();
          }

          super(FlowSchema.__pulumiType, name, props, opts);
      }
    }