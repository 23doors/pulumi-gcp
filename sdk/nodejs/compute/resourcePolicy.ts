// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * > This content is derived from https://github.com/terraform-providers/terraform-provider-google/blob/master/website/docs/r/compute_resource_policy.html.markdown.
 */
export class ResourcePolicy extends pulumi.CustomResource {
    /**
     * Get an existing ResourcePolicy resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ResourcePolicyState, opts?: pulumi.CustomResourceOptions): ResourcePolicy {
        return new ResourcePolicy(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'gcp:compute/resourcePolicy:ResourcePolicy';

    /**
     * Returns true if the given object is an instance of ResourcePolicy.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ResourcePolicy {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ResourcePolicy.__pulumiType;
    }

    /**
     * The name of the resource, provided by the client when initially creating the resource. The resource name must be
     * 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the
     * regular expression '[a-z]([-a-z0-9]*[a-z0-9])'? which means the first character must be a lowercase letter, and all
     * following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the provider project is used.
     */
    public readonly project!: pulumi.Output<string>;
    /**
     * Region where resource policy resides.
     */
    public readonly region!: pulumi.Output<string>;
    /**
     * The URI of the created resource.
     */
    public /*out*/ readonly selfLink!: pulumi.Output<string>;
    /**
     * Policy for creating snapshots of persistent disks.
     */
    public readonly snapshotSchedulePolicy!: pulumi.Output<outputs.compute.ResourcePolicySnapshotSchedulePolicy | undefined>;

    /**
     * Create a ResourcePolicy resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: ResourcePolicyArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ResourcePolicyArgs | ResourcePolicyState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as ResourcePolicyState | undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["project"] = state ? state.project : undefined;
            inputs["region"] = state ? state.region : undefined;
            inputs["selfLink"] = state ? state.selfLink : undefined;
            inputs["snapshotSchedulePolicy"] = state ? state.snapshotSchedulePolicy : undefined;
        } else {
            const args = argsOrState as ResourcePolicyArgs | undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["project"] = args ? args.project : undefined;
            inputs["region"] = args ? args.region : undefined;
            inputs["snapshotSchedulePolicy"] = args ? args.snapshotSchedulePolicy : undefined;
            inputs["selfLink"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(ResourcePolicy.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering ResourcePolicy resources.
 */
export interface ResourcePolicyState {
    /**
     * The name of the resource, provided by the client when initially creating the resource. The resource name must be
     * 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the
     * regular expression '[a-z]([-a-z0-9]*[a-z0-9])'? which means the first character must be a lowercase letter, and all
     * following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the provider project is used.
     */
    readonly project?: pulumi.Input<string>;
    /**
     * Region where resource policy resides.
     */
    readonly region?: pulumi.Input<string>;
    /**
     * The URI of the created resource.
     */
    readonly selfLink?: pulumi.Input<string>;
    /**
     * Policy for creating snapshots of persistent disks.
     */
    readonly snapshotSchedulePolicy?: pulumi.Input<inputs.compute.ResourcePolicySnapshotSchedulePolicy>;
}

/**
 * The set of arguments for constructing a ResourcePolicy resource.
 */
export interface ResourcePolicyArgs {
    /**
     * The name of the resource, provided by the client when initially creating the resource. The resource name must be
     * 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the
     * regular expression '[a-z]([-a-z0-9]*[a-z0-9])'? which means the first character must be a lowercase letter, and all
     * following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the provider project is used.
     */
    readonly project?: pulumi.Input<string>;
    /**
     * Region where resource policy resides.
     */
    readonly region?: pulumi.Input<string>;
    /**
     * Policy for creating snapshots of persistent disks.
     */
    readonly snapshotSchedulePolicy?: pulumi.Input<inputs.compute.ResourcePolicySnapshotSchedulePolicy>;
}
