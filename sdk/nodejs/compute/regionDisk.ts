// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * > This content is derived from https://github.com/terraform-providers/terraform-provider-google/blob/master/website/docs/r/compute_region_disk.html.markdown.
 */
export class RegionDisk extends pulumi.CustomResource {
    /**
     * Get an existing RegionDisk resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: RegionDiskState, opts?: pulumi.CustomResourceOptions): RegionDisk {
        return new RegionDisk(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'gcp:compute/regionDisk:RegionDisk';

    /**
     * Returns true if the given object is an instance of RegionDisk.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is RegionDisk {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === RegionDisk.__pulumiType;
    }

    public /*out*/ readonly creationTimestamp!: pulumi.Output<string>;
    public readonly description!: pulumi.Output<string | undefined>;
    public readonly diskEncryptionKey!: pulumi.Output<outputs.compute.RegionDiskDiskEncryptionKey | undefined>;
    public /*out*/ readonly labelFingerprint!: pulumi.Output<string>;
    public readonly labels!: pulumi.Output<{[key: string]: string} | undefined>;
    public /*out*/ readonly lastAttachTimestamp!: pulumi.Output<string>;
    public /*out*/ readonly lastDetachTimestamp!: pulumi.Output<string>;
    public readonly name!: pulumi.Output<string>;
    public readonly physicalBlockSizeBytes!: pulumi.Output<number>;
    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the provider project is used.
     */
    public readonly project!: pulumi.Output<string>;
    public readonly region!: pulumi.Output<string>;
    public readonly replicaZones!: pulumi.Output<string[]>;
    /**
     * The URI of the created resource.
     */
    public /*out*/ readonly selfLink!: pulumi.Output<string>;
    public readonly size!: pulumi.Output<number>;
    public readonly snapshot!: pulumi.Output<string | undefined>;
    public readonly sourceSnapshotEncryptionKey!: pulumi.Output<outputs.compute.RegionDiskSourceSnapshotEncryptionKey | undefined>;
    public /*out*/ readonly sourceSnapshotId!: pulumi.Output<string>;
    public readonly type!: pulumi.Output<string | undefined>;
    public /*out*/ readonly users!: pulumi.Output<string[]>;

    /**
     * Create a RegionDisk resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: RegionDiskArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: RegionDiskArgs | RegionDiskState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as RegionDiskState | undefined;
            inputs["creationTimestamp"] = state ? state.creationTimestamp : undefined;
            inputs["description"] = state ? state.description : undefined;
            inputs["diskEncryptionKey"] = state ? state.diskEncryptionKey : undefined;
            inputs["labelFingerprint"] = state ? state.labelFingerprint : undefined;
            inputs["labels"] = state ? state.labels : undefined;
            inputs["lastAttachTimestamp"] = state ? state.lastAttachTimestamp : undefined;
            inputs["lastDetachTimestamp"] = state ? state.lastDetachTimestamp : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["physicalBlockSizeBytes"] = state ? state.physicalBlockSizeBytes : undefined;
            inputs["project"] = state ? state.project : undefined;
            inputs["region"] = state ? state.region : undefined;
            inputs["replicaZones"] = state ? state.replicaZones : undefined;
            inputs["selfLink"] = state ? state.selfLink : undefined;
            inputs["size"] = state ? state.size : undefined;
            inputs["snapshot"] = state ? state.snapshot : undefined;
            inputs["sourceSnapshotEncryptionKey"] = state ? state.sourceSnapshotEncryptionKey : undefined;
            inputs["sourceSnapshotId"] = state ? state.sourceSnapshotId : undefined;
            inputs["type"] = state ? state.type : undefined;
            inputs["users"] = state ? state.users : undefined;
        } else {
            const args = argsOrState as RegionDiskArgs | undefined;
            if (!args || args.replicaZones === undefined) {
                throw new Error("Missing required property 'replicaZones'");
            }
            inputs["description"] = args ? args.description : undefined;
            inputs["diskEncryptionKey"] = args ? args.diskEncryptionKey : undefined;
            inputs["labels"] = args ? args.labels : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["physicalBlockSizeBytes"] = args ? args.physicalBlockSizeBytes : undefined;
            inputs["project"] = args ? args.project : undefined;
            inputs["region"] = args ? args.region : undefined;
            inputs["replicaZones"] = args ? args.replicaZones : undefined;
            inputs["size"] = args ? args.size : undefined;
            inputs["snapshot"] = args ? args.snapshot : undefined;
            inputs["sourceSnapshotEncryptionKey"] = args ? args.sourceSnapshotEncryptionKey : undefined;
            inputs["type"] = args ? args.type : undefined;
            inputs["creationTimestamp"] = undefined /*out*/;
            inputs["labelFingerprint"] = undefined /*out*/;
            inputs["lastAttachTimestamp"] = undefined /*out*/;
            inputs["lastDetachTimestamp"] = undefined /*out*/;
            inputs["selfLink"] = undefined /*out*/;
            inputs["sourceSnapshotId"] = undefined /*out*/;
            inputs["users"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(RegionDisk.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering RegionDisk resources.
 */
export interface RegionDiskState {
    readonly creationTimestamp?: pulumi.Input<string>;
    readonly description?: pulumi.Input<string>;
    readonly diskEncryptionKey?: pulumi.Input<inputs.compute.RegionDiskDiskEncryptionKey>;
    readonly labelFingerprint?: pulumi.Input<string>;
    readonly labels?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    readonly lastAttachTimestamp?: pulumi.Input<string>;
    readonly lastDetachTimestamp?: pulumi.Input<string>;
    readonly name?: pulumi.Input<string>;
    readonly physicalBlockSizeBytes?: pulumi.Input<number>;
    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the provider project is used.
     */
    readonly project?: pulumi.Input<string>;
    readonly region?: pulumi.Input<string>;
    readonly replicaZones?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The URI of the created resource.
     */
    readonly selfLink?: pulumi.Input<string>;
    readonly size?: pulumi.Input<number>;
    readonly snapshot?: pulumi.Input<string>;
    readonly sourceSnapshotEncryptionKey?: pulumi.Input<inputs.compute.RegionDiskSourceSnapshotEncryptionKey>;
    readonly sourceSnapshotId?: pulumi.Input<string>;
    readonly type?: pulumi.Input<string>;
    readonly users?: pulumi.Input<pulumi.Input<string>[]>;
}

/**
 * The set of arguments for constructing a RegionDisk resource.
 */
export interface RegionDiskArgs {
    readonly description?: pulumi.Input<string>;
    readonly diskEncryptionKey?: pulumi.Input<inputs.compute.RegionDiskDiskEncryptionKey>;
    readonly labels?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    readonly name?: pulumi.Input<string>;
    readonly physicalBlockSizeBytes?: pulumi.Input<number>;
    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the provider project is used.
     */
    readonly project?: pulumi.Input<string>;
    readonly region?: pulumi.Input<string>;
    readonly replicaZones: pulumi.Input<pulumi.Input<string>[]>;
    readonly size?: pulumi.Input<number>;
    readonly snapshot?: pulumi.Input<string>;
    readonly sourceSnapshotEncryptionKey?: pulumi.Input<inputs.compute.RegionDiskSourceSnapshotEncryptionKey>;
    readonly type?: pulumi.Input<string>;
}
