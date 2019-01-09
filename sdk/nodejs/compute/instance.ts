// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

export class Instance extends pulumi.CustomResource {
    /**
     * Get an existing Instance resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: InstanceState, opts?: pulumi.CustomResourceOptions): Instance {
        return new Instance(name, <any>state, { ...opts, id: id });
    }

    public readonly allowStoppingForUpdate: pulumi.Output<boolean | undefined>;
    public readonly attachedDisks: pulumi.Output<{ deviceName: string, diskEncryptionKeyRaw?: string, diskEncryptionKeySha256: string, mode?: string, source: string }[] | undefined>;
    public readonly bootDisk: pulumi.Output<{ autoDelete?: boolean, deviceName: string, diskEncryptionKeyRaw?: string, diskEncryptionKeySha256: string, initializeParams: { image: string, size: number, type: string }, source: string }>;
    public readonly canIpForward: pulumi.Output<boolean | undefined>;
    public /*out*/ readonly cpuPlatform: pulumi.Output<string>;
    public readonly createTimeout: pulumi.Output<number | undefined>;
    public readonly deletionProtection: pulumi.Output<boolean | undefined>;
    public readonly description: pulumi.Output<string | undefined>;
    public readonly guestAccelerators: pulumi.Output<{ count: number, type: string }[]>;
    public /*out*/ readonly instanceId: pulumi.Output<string>;
    public /*out*/ readonly labelFingerprint: pulumi.Output<string>;
    public readonly labels: pulumi.Output<{[key: string]: string} | undefined>;
    public readonly machineType: pulumi.Output<string>;
    public readonly metadata: pulumi.Output<{[key: string]: string} | undefined>;
    public /*out*/ readonly metadataFingerprint: pulumi.Output<string>;
    public readonly metadataStartupScript: pulumi.Output<string | undefined>;
    public readonly minCpuPlatform: pulumi.Output<string | undefined>;
    public readonly name: pulumi.Output<string>;
    public readonly networkInterfaces: pulumi.Output<{ accessConfigs?: { assignedNatIp: string, natIp: string, networkTier: string, publicPtrDomainName?: string }[], address: string, aliasIpRanges?: { ipCidrRange: string, subnetworkRangeName?: string }[], name: string, network: string, networkIp: string, subnetwork: string, subnetworkProject: string }[]>;
    public readonly project: pulumi.Output<string>;
    public readonly scheduling: pulumi.Output<{ automaticRestart?: boolean, onHostMaintenance: string, preemptible?: boolean }>;
    public readonly scratchDisks: pulumi.Output<{ interface?: string }[] | undefined>;
    public /*out*/ readonly selfLink: pulumi.Output<string>;
    public readonly serviceAccount: pulumi.Output<{ email: string, scopes: string[] } | undefined>;
    public readonly tags: pulumi.Output<string[] | undefined>;
    public /*out*/ readonly tagsFingerprint: pulumi.Output<string>;
    public readonly zone: pulumi.Output<string>;

    /**
     * Create a Instance resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: InstanceArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: InstanceArgs | InstanceState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state: InstanceState = argsOrState as InstanceState | undefined;
            inputs["allowStoppingForUpdate"] = state ? state.allowStoppingForUpdate : undefined;
            inputs["attachedDisks"] = state ? state.attachedDisks : undefined;
            inputs["bootDisk"] = state ? state.bootDisk : undefined;
            inputs["canIpForward"] = state ? state.canIpForward : undefined;
            inputs["cpuPlatform"] = state ? state.cpuPlatform : undefined;
            inputs["createTimeout"] = state ? state.createTimeout : undefined;
            inputs["deletionProtection"] = state ? state.deletionProtection : undefined;
            inputs["description"] = state ? state.description : undefined;
            inputs["guestAccelerators"] = state ? state.guestAccelerators : undefined;
            inputs["instanceId"] = state ? state.instanceId : undefined;
            inputs["labelFingerprint"] = state ? state.labelFingerprint : undefined;
            inputs["labels"] = state ? state.labels : undefined;
            inputs["machineType"] = state ? state.machineType : undefined;
            inputs["metadata"] = state ? state.metadata : undefined;
            inputs["metadataFingerprint"] = state ? state.metadataFingerprint : undefined;
            inputs["metadataStartupScript"] = state ? state.metadataStartupScript : undefined;
            inputs["minCpuPlatform"] = state ? state.minCpuPlatform : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["networkInterfaces"] = state ? state.networkInterfaces : undefined;
            inputs["project"] = state ? state.project : undefined;
            inputs["scheduling"] = state ? state.scheduling : undefined;
            inputs["scratchDisks"] = state ? state.scratchDisks : undefined;
            inputs["selfLink"] = state ? state.selfLink : undefined;
            inputs["serviceAccount"] = state ? state.serviceAccount : undefined;
            inputs["tags"] = state ? state.tags : undefined;
            inputs["tagsFingerprint"] = state ? state.tagsFingerprint : undefined;
            inputs["zone"] = state ? state.zone : undefined;
        } else {
            const args = argsOrState as InstanceArgs | undefined;
            if (!args || args.bootDisk === undefined) {
                throw new Error("Missing required property 'bootDisk'");
            }
            if (!args || args.machineType === undefined) {
                throw new Error("Missing required property 'machineType'");
            }
            if (!args || args.networkInterfaces === undefined) {
                throw new Error("Missing required property 'networkInterfaces'");
            }
            inputs["allowStoppingForUpdate"] = args ? args.allowStoppingForUpdate : undefined;
            inputs["attachedDisks"] = args ? args.attachedDisks : undefined;
            inputs["bootDisk"] = args ? args.bootDisk : undefined;
            inputs["canIpForward"] = args ? args.canIpForward : undefined;
            inputs["createTimeout"] = args ? args.createTimeout : undefined;
            inputs["deletionProtection"] = args ? args.deletionProtection : undefined;
            inputs["description"] = args ? args.description : undefined;
            inputs["guestAccelerators"] = args ? args.guestAccelerators : undefined;
            inputs["labels"] = args ? args.labels : undefined;
            inputs["machineType"] = args ? args.machineType : undefined;
            inputs["metadata"] = args ? args.metadata : undefined;
            inputs["metadataStartupScript"] = args ? args.metadataStartupScript : undefined;
            inputs["minCpuPlatform"] = args ? args.minCpuPlatform : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["networkInterfaces"] = args ? args.networkInterfaces : undefined;
            inputs["project"] = args ? args.project : undefined;
            inputs["scheduling"] = args ? args.scheduling : undefined;
            inputs["scratchDisks"] = args ? args.scratchDisks : undefined;
            inputs["serviceAccount"] = args ? args.serviceAccount : undefined;
            inputs["tags"] = args ? args.tags : undefined;
            inputs["zone"] = args ? args.zone : undefined;
            inputs["cpuPlatform"] = undefined /*out*/;
            inputs["instanceId"] = undefined /*out*/;
            inputs["labelFingerprint"] = undefined /*out*/;
            inputs["metadataFingerprint"] = undefined /*out*/;
            inputs["selfLink"] = undefined /*out*/;
            inputs["tagsFingerprint"] = undefined /*out*/;
        }
        super("gcp:compute/instance:Instance", name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Instance resources.
 */
export interface InstanceState {
    readonly allowStoppingForUpdate?: pulumi.Input<boolean>;
    readonly attachedDisks?: pulumi.Input<pulumi.Input<{ deviceName?: pulumi.Input<string>, diskEncryptionKeyRaw?: pulumi.Input<string>, diskEncryptionKeySha256?: pulumi.Input<string>, mode?: pulumi.Input<string>, source: pulumi.Input<string> }>[]>;
    readonly bootDisk?: pulumi.Input<{ autoDelete?: pulumi.Input<boolean>, deviceName?: pulumi.Input<string>, diskEncryptionKeyRaw?: pulumi.Input<string>, diskEncryptionKeySha256?: pulumi.Input<string>, initializeParams?: pulumi.Input<{ image?: pulumi.Input<string>, size?: pulumi.Input<number>, type?: pulumi.Input<string> }>, source?: pulumi.Input<string> }>;
    readonly canIpForward?: pulumi.Input<boolean>;
    readonly cpuPlatform?: pulumi.Input<string>;
    readonly createTimeout?: pulumi.Input<number>;
    readonly deletionProtection?: pulumi.Input<boolean>;
    readonly description?: pulumi.Input<string>;
    readonly guestAccelerators?: pulumi.Input<pulumi.Input<{ count: pulumi.Input<number>, type: pulumi.Input<string> }>[]>;
    readonly instanceId?: pulumi.Input<string>;
    readonly labelFingerprint?: pulumi.Input<string>;
    readonly labels?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    readonly machineType?: pulumi.Input<string>;
    readonly metadata?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    readonly metadataFingerprint?: pulumi.Input<string>;
    readonly metadataStartupScript?: pulumi.Input<string>;
    readonly minCpuPlatform?: pulumi.Input<string>;
    readonly name?: pulumi.Input<string>;
    readonly networkInterfaces?: pulumi.Input<pulumi.Input<{ accessConfigs?: pulumi.Input<pulumi.Input<{ assignedNatIp?: pulumi.Input<string>, natIp?: pulumi.Input<string>, networkTier?: pulumi.Input<string>, publicPtrDomainName?: pulumi.Input<string> }>[]>, address?: pulumi.Input<string>, aliasIpRanges?: pulumi.Input<pulumi.Input<{ ipCidrRange: pulumi.Input<string>, subnetworkRangeName?: pulumi.Input<string> }>[]>, name?: pulumi.Input<string>, network?: pulumi.Input<string>, networkIp?: pulumi.Input<string>, subnetwork?: pulumi.Input<string>, subnetworkProject?: pulumi.Input<string> }>[]>;
    readonly project?: pulumi.Input<string>;
    readonly scheduling?: pulumi.Input<{ automaticRestart?: pulumi.Input<boolean>, onHostMaintenance?: pulumi.Input<string>, preemptible?: pulumi.Input<boolean> }>;
    readonly scratchDisks?: pulumi.Input<pulumi.Input<{ interface?: pulumi.Input<string> }>[]>;
    readonly selfLink?: pulumi.Input<string>;
    readonly serviceAccount?: pulumi.Input<{ email?: pulumi.Input<string>, scopes: pulumi.Input<pulumi.Input<string>[]> }>;
    readonly tags?: pulumi.Input<pulumi.Input<string>[]>;
    readonly tagsFingerprint?: pulumi.Input<string>;
    readonly zone?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Instance resource.
 */
export interface InstanceArgs {
    readonly allowStoppingForUpdate?: pulumi.Input<boolean>;
    readonly attachedDisks?: pulumi.Input<pulumi.Input<{ deviceName?: pulumi.Input<string>, diskEncryptionKeyRaw?: pulumi.Input<string>, diskEncryptionKeySha256?: pulumi.Input<string>, mode?: pulumi.Input<string>, source: pulumi.Input<string> }>[]>;
    readonly bootDisk: pulumi.Input<{ autoDelete?: pulumi.Input<boolean>, deviceName?: pulumi.Input<string>, diskEncryptionKeyRaw?: pulumi.Input<string>, diskEncryptionKeySha256?: pulumi.Input<string>, initializeParams?: pulumi.Input<{ image?: pulumi.Input<string>, size?: pulumi.Input<number>, type?: pulumi.Input<string> }>, source?: pulumi.Input<string> }>;
    readonly canIpForward?: pulumi.Input<boolean>;
    readonly createTimeout?: pulumi.Input<number>;
    readonly deletionProtection?: pulumi.Input<boolean>;
    readonly description?: pulumi.Input<string>;
    readonly guestAccelerators?: pulumi.Input<pulumi.Input<{ count: pulumi.Input<number>, type: pulumi.Input<string> }>[]>;
    readonly labels?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    readonly machineType: pulumi.Input<string>;
    readonly metadata?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    readonly metadataStartupScript?: pulumi.Input<string>;
    readonly minCpuPlatform?: pulumi.Input<string>;
    readonly name?: pulumi.Input<string>;
    readonly networkInterfaces: pulumi.Input<pulumi.Input<{ accessConfigs?: pulumi.Input<pulumi.Input<{ assignedNatIp?: pulumi.Input<string>, natIp?: pulumi.Input<string>, networkTier?: pulumi.Input<string>, publicPtrDomainName?: pulumi.Input<string> }>[]>, address?: pulumi.Input<string>, aliasIpRanges?: pulumi.Input<pulumi.Input<{ ipCidrRange: pulumi.Input<string>, subnetworkRangeName?: pulumi.Input<string> }>[]>, name?: pulumi.Input<string>, network?: pulumi.Input<string>, networkIp?: pulumi.Input<string>, subnetwork?: pulumi.Input<string>, subnetworkProject?: pulumi.Input<string> }>[]>;
    readonly project?: pulumi.Input<string>;
    readonly scheduling?: pulumi.Input<{ automaticRestart?: pulumi.Input<boolean>, onHostMaintenance?: pulumi.Input<string>, preemptible?: pulumi.Input<boolean> }>;
    readonly scratchDisks?: pulumi.Input<pulumi.Input<{ interface?: pulumi.Input<string> }>[]>;
    readonly serviceAccount?: pulumi.Input<{ email?: pulumi.Input<string>, scopes: pulumi.Input<pulumi.Input<string>[]> }>;
    readonly tags?: pulumi.Input<pulumi.Input<string>[]>;
    readonly zone?: pulumi.Input<string>;
}
