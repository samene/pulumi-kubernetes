// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../../utilities";

// Export members:
export * from "./ingress";
export * from "./ingressClass";
export * from "./ingressClassList";
export * from "./ingressClassPatch";
export * from "./ingressList";
export * from "./ingressPatch";

// Import resources to register:
import { Ingress } from "./ingress";
import { IngressClass } from "./ingressClass";
import { IngressClassList } from "./ingressClassList";
import { IngressClassPatch } from "./ingressClassPatch";
import { IngressList } from "./ingressList";
import { IngressPatch } from "./ingressPatch";

const _module = {
    version: utilities.getVersion(),
    construct: (name: string, type: string, urn: string): pulumi.Resource => {
        switch (type) {
            case "kubernetes:networking.k8s.io/v1beta1:Ingress":
                return new Ingress(name, <any>undefined, { urn })
            case "kubernetes:networking.k8s.io/v1beta1:IngressClass":
                return new IngressClass(name, <any>undefined, { urn })
            case "kubernetes:networking.k8s.io/v1beta1:IngressClassList":
                return new IngressClassList(name, <any>undefined, { urn })
            case "kubernetes:networking.k8s.io/v1beta1:IngressClassPatch":
                return new IngressClassPatch(name, <any>undefined, { urn })
            case "kubernetes:networking.k8s.io/v1beta1:IngressList":
                return new IngressList(name, <any>undefined, { urn })
            case "kubernetes:networking.k8s.io/v1beta1:IngressPatch":
                return new IngressPatch(name, <any>undefined, { urn })
            default:
                throw new Error(`unknown resource type ${type}`);
        }
    },
};
pulumi.runtime.registerResourceModule("kubernetes", "networking.k8s.io/v1beta1", _module)
