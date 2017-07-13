// *** WARNING: this file was generated by the Lumi Terraform Bridge (TFGEN) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as lumi from "@lumi/lumi";

export class Profile extends lumi.NamedResource implements ProfileArgs {
    public readonly dnsConfig: { relativeName: string, ttl: number }[];
    public readonly fqdn?: string;
    public readonly monitorConfig: { path: string, port: number, protocol: string }[];
    public readonly _name: string;
    public readonly profileStatus?: string;
    public readonly resourceGroupName: string;
    public readonly tags?: {[key: string]: any};
    public readonly trafficRoutingMethod: string;

    constructor(name: string, args: ProfileArgs) {
        super(name);
        this.dnsConfig = args.dnsConfig;
        this.fqdn = args.fqdn;
        this.monitorConfig = args.monitorConfig;
        this._name = args._name;
        this.profileStatus = args.profileStatus;
        this.resourceGroupName = args.resourceGroupName;
        this.tags = args.tags;
        this.trafficRoutingMethod = args.trafficRoutingMethod;
    }
}

export interface ProfileArgs {
    readonly dnsConfig: { relativeName: string, ttl: number }[];
    readonly fqdn?: string;
    readonly monitorConfig: { path: string, port: number, protocol: string }[];
    readonly _name: string;
    readonly profileStatus?: string;
    readonly resourceGroupName: string;
    readonly tags?: {[key: string]: any};
    readonly trafficRoutingMethod: string;
}
