// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../../utilities";

// Export members:
export * from "./binding";
export * from "./bindingPatch";
export * from "./configMap";
export * from "./configMapList";
export * from "./configMapPatch";
export * from "./endpoints";
export * from "./endpointsList";
export * from "./endpointsPatch";
export * from "./event";
export * from "./eventList";
export * from "./eventPatch";
export * from "./limitRange";
export * from "./limitRangeList";
export * from "./limitRangePatch";
export * from "./namespace";
export * from "./namespaceList";
export * from "./namespacePatch";
export * from "./node";
export * from "./nodeList";
export * from "./nodePatch";
export * from "./persistentVolume";
export * from "./persistentVolumeClaim";
export * from "./persistentVolumeClaimList";
export * from "./persistentVolumeClaimPatch";
export * from "./persistentVolumeList";
export * from "./persistentVolumePatch";
export * from "./pod";
export * from "./podList";
export * from "./podPatch";
export * from "./podTemplate";
export * from "./podTemplateList";
export * from "./podTemplatePatch";
export * from "./replicationController";
export * from "./replicationControllerList";
export * from "./replicationControllerPatch";
export * from "./resourceQuota";
export * from "./resourceQuotaList";
export * from "./resourceQuotaPatch";
export * from "./secret";
export * from "./secretList";
export * from "./secretPatch";
export * from "./service";
export * from "./serviceAccount";
export * from "./serviceAccountList";
export * from "./serviceAccountPatch";
export * from "./serviceList";
export * from "./servicePatch";

// Export enums:
export * from "../../types/enums/core/v1";

// Import resources to register:
import { Binding } from "./binding";
import { BindingPatch } from "./bindingPatch";
import { ConfigMap } from "./configMap";
import { ConfigMapList } from "./configMapList";
import { ConfigMapPatch } from "./configMapPatch";
import { Endpoints } from "./endpoints";
import { EndpointsList } from "./endpointsList";
import { EndpointsPatch } from "./endpointsPatch";
import { Event } from "./event";
import { EventList } from "./eventList";
import { EventPatch } from "./eventPatch";
import { LimitRange } from "./limitRange";
import { LimitRangeList } from "./limitRangeList";
import { LimitRangePatch } from "./limitRangePatch";
import { Namespace } from "./namespace";
import { NamespaceList } from "./namespaceList";
import { NamespacePatch } from "./namespacePatch";
import { Node } from "./node";
import { NodeList } from "./nodeList";
import { NodePatch } from "./nodePatch";
import { PersistentVolume } from "./persistentVolume";
import { PersistentVolumeClaim } from "./persistentVolumeClaim";
import { PersistentVolumeClaimList } from "./persistentVolumeClaimList";
import { PersistentVolumeClaimPatch } from "./persistentVolumeClaimPatch";
import { PersistentVolumeList } from "./persistentVolumeList";
import { PersistentVolumePatch } from "./persistentVolumePatch";
import { Pod } from "./pod";
import { PodList } from "./podList";
import { PodPatch } from "./podPatch";
import { PodTemplate } from "./podTemplate";
import { PodTemplateList } from "./podTemplateList";
import { PodTemplatePatch } from "./podTemplatePatch";
import { ReplicationController } from "./replicationController";
import { ReplicationControllerList } from "./replicationControllerList";
import { ReplicationControllerPatch } from "./replicationControllerPatch";
import { ResourceQuota } from "./resourceQuota";
import { ResourceQuotaList } from "./resourceQuotaList";
import { ResourceQuotaPatch } from "./resourceQuotaPatch";
import { Secret } from "./secret";
import { SecretList } from "./secretList";
import { SecretPatch } from "./secretPatch";
import { Service } from "./service";
import { ServiceAccount } from "./serviceAccount";
import { ServiceAccountList } from "./serviceAccountList";
import { ServiceAccountPatch } from "./serviceAccountPatch";
import { ServiceList } from "./serviceList";
import { ServicePatch } from "./servicePatch";

const _module = {
    version: utilities.getVersion(),
    construct: (name: string, type: string, urn: string): pulumi.Resource => {
        switch (type) {
            case "kubernetes:core/v1:Binding":
                return new Binding(name, <any>undefined, { urn })
            case "kubernetes:core/v1:BindingPatch":
                return new BindingPatch(name, <any>undefined, { urn })
            case "kubernetes:core/v1:ConfigMap":
                return new ConfigMap(name, <any>undefined, { urn })
            case "kubernetes:core/v1:ConfigMapList":
                return new ConfigMapList(name, <any>undefined, { urn })
            case "kubernetes:core/v1:ConfigMapPatch":
                return new ConfigMapPatch(name, <any>undefined, { urn })
            case "kubernetes:core/v1:Endpoints":
                return new Endpoints(name, <any>undefined, { urn })
            case "kubernetes:core/v1:EndpointsList":
                return new EndpointsList(name, <any>undefined, { urn })
            case "kubernetes:core/v1:EndpointsPatch":
                return new EndpointsPatch(name, <any>undefined, { urn })
            case "kubernetes:core/v1:Event":
                return new Event(name, <any>undefined, { urn })
            case "kubernetes:core/v1:EventList":
                return new EventList(name, <any>undefined, { urn })
            case "kubernetes:core/v1:EventPatch":
                return new EventPatch(name, <any>undefined, { urn })
            case "kubernetes:core/v1:LimitRange":
                return new LimitRange(name, <any>undefined, { urn })
            case "kubernetes:core/v1:LimitRangeList":
                return new LimitRangeList(name, <any>undefined, { urn })
            case "kubernetes:core/v1:LimitRangePatch":
                return new LimitRangePatch(name, <any>undefined, { urn })
            case "kubernetes:core/v1:Namespace":
                return new Namespace(name, <any>undefined, { urn })
            case "kubernetes:core/v1:NamespaceList":
                return new NamespaceList(name, <any>undefined, { urn })
            case "kubernetes:core/v1:NamespacePatch":
                return new NamespacePatch(name, <any>undefined, { urn })
            case "kubernetes:core/v1:Node":
                return new Node(name, <any>undefined, { urn })
            case "kubernetes:core/v1:NodeList":
                return new NodeList(name, <any>undefined, { urn })
            case "kubernetes:core/v1:NodePatch":
                return new NodePatch(name, <any>undefined, { urn })
            case "kubernetes:core/v1:PersistentVolume":
                return new PersistentVolume(name, <any>undefined, { urn })
            case "kubernetes:core/v1:PersistentVolumeClaim":
                return new PersistentVolumeClaim(name, <any>undefined, { urn })
            case "kubernetes:core/v1:PersistentVolumeClaimList":
                return new PersistentVolumeClaimList(name, <any>undefined, { urn })
            case "kubernetes:core/v1:PersistentVolumeClaimPatch":
                return new PersistentVolumeClaimPatch(name, <any>undefined, { urn })
            case "kubernetes:core/v1:PersistentVolumeList":
                return new PersistentVolumeList(name, <any>undefined, { urn })
            case "kubernetes:core/v1:PersistentVolumePatch":
                return new PersistentVolumePatch(name, <any>undefined, { urn })
            case "kubernetes:core/v1:Pod":
                return new Pod(name, <any>undefined, { urn })
            case "kubernetes:core/v1:PodList":
                return new PodList(name, <any>undefined, { urn })
            case "kubernetes:core/v1:PodPatch":
                return new PodPatch(name, <any>undefined, { urn })
            case "kubernetes:core/v1:PodTemplate":
                return new PodTemplate(name, <any>undefined, { urn })
            case "kubernetes:core/v1:PodTemplateList":
                return new PodTemplateList(name, <any>undefined, { urn })
            case "kubernetes:core/v1:PodTemplatePatch":
                return new PodTemplatePatch(name, <any>undefined, { urn })
            case "kubernetes:core/v1:ReplicationController":
                return new ReplicationController(name, <any>undefined, { urn })
            case "kubernetes:core/v1:ReplicationControllerList":
                return new ReplicationControllerList(name, <any>undefined, { urn })
            case "kubernetes:core/v1:ReplicationControllerPatch":
                return new ReplicationControllerPatch(name, <any>undefined, { urn })
            case "kubernetes:core/v1:ResourceQuota":
                return new ResourceQuota(name, <any>undefined, { urn })
            case "kubernetes:core/v1:ResourceQuotaList":
                return new ResourceQuotaList(name, <any>undefined, { urn })
            case "kubernetes:core/v1:ResourceQuotaPatch":
                return new ResourceQuotaPatch(name, <any>undefined, { urn })
            case "kubernetes:core/v1:Secret":
                return new Secret(name, <any>undefined, { urn })
            case "kubernetes:core/v1:SecretList":
                return new SecretList(name, <any>undefined, { urn })
            case "kubernetes:core/v1:SecretPatch":
                return new SecretPatch(name, <any>undefined, { urn })
            case "kubernetes:core/v1:Service":
                return new Service(name, <any>undefined, { urn })
            case "kubernetes:core/v1:ServiceAccount":
                return new ServiceAccount(name, <any>undefined, { urn })
            case "kubernetes:core/v1:ServiceAccountList":
                return new ServiceAccountList(name, <any>undefined, { urn })
            case "kubernetes:core/v1:ServiceAccountPatch":
                return new ServiceAccountPatch(name, <any>undefined, { urn })
            case "kubernetes:core/v1:ServiceList":
                return new ServiceList(name, <any>undefined, { urn })
            case "kubernetes:core/v1:ServicePatch":
                return new ServicePatch(name, <any>undefined, { urn })
            default:
                throw new Error(`unknown resource type ${type}`);
        }
    },
};
pulumi.runtime.registerResourceModule("kubernetes", "core/v1", _module)
