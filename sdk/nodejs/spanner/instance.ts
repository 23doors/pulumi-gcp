// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * > This content is derived from https://github.com/terraform-providers/terraform-provider-google/blob/master/website/docs/r/spanner_instance.html.markdown.
 */
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

    /** @internal */
    public static readonly __pulumiType = 'gcp:spanner/instance:Instance';

    /**
     * Returns true if the given object is an instance of Instance.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Instance {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Instance.__pulumiType;
    }

    /**
     * The name of the instance's configuration (similar but not quite the same as a region) which defines defines the
     * geographic placement and replication of your databases in this instance. It determines where your data is stored.
     * Values are typically of the form 'regional-europe-west1' , 'us-central' etc. In order to obtain a valid list please
     * consult the [Configuration section of the docs](https://cloud.google.com/spanner/docs/instances).
     */
    public readonly config!: pulumi.Output<string>;
    /**
     * The descriptive name for this instance as it appears in UIs. Must be unique per project and between 4 and 30
     * characters in length.
     */
    public readonly displayName!: pulumi.Output<string>;
    /**
     * An object containing a list of "key": value pairs. Example: { "name": "wrench", "mass": "1.3kg", "count": "3" }.
     */
    public readonly labels!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * A unique identifier for the instance, which cannot be changed after the instance is created. The name must be
     * between 6 and 30 characters in length. If not provided, a random string starting with 'tf-' will be selected.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The number of nodes allocated to this instance.
     */
    public readonly numNodes!: pulumi.Output<number | undefined>;
    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the provider project is used.
     */
    public readonly project!: pulumi.Output<string>;
    /**
     * Instance status: 'CREATING' or 'READY'.
     */
    public /*out*/ readonly state!: pulumi.Output<string>;

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
            const state = argsOrState as InstanceState | undefined;
            inputs["config"] = state ? state.config : undefined;
            inputs["displayName"] = state ? state.displayName : undefined;
            inputs["labels"] = state ? state.labels : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["numNodes"] = state ? state.numNodes : undefined;
            inputs["project"] = state ? state.project : undefined;
            inputs["state"] = state ? state.state : undefined;
        } else {
            const args = argsOrState as InstanceArgs | undefined;
            if (!args || args.config === undefined) {
                throw new Error("Missing required property 'config'");
            }
            if (!args || args.displayName === undefined) {
                throw new Error("Missing required property 'displayName'");
            }
            inputs["config"] = args ? args.config : undefined;
            inputs["displayName"] = args ? args.displayName : undefined;
            inputs["labels"] = args ? args.labels : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["numNodes"] = args ? args.numNodes : undefined;
            inputs["project"] = args ? args.project : undefined;
            inputs["state"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(Instance.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Instance resources.
 */
export interface InstanceState {
    /**
     * The name of the instance's configuration (similar but not quite the same as a region) which defines defines the
     * geographic placement and replication of your databases in this instance. It determines where your data is stored.
     * Values are typically of the form 'regional-europe-west1' , 'us-central' etc. In order to obtain a valid list please
     * consult the [Configuration section of the docs](https://cloud.google.com/spanner/docs/instances).
     */
    readonly config?: pulumi.Input<string>;
    /**
     * The descriptive name for this instance as it appears in UIs. Must be unique per project and between 4 and 30
     * characters in length.
     */
    readonly displayName?: pulumi.Input<string>;
    /**
     * An object containing a list of "key": value pairs. Example: { "name": "wrench", "mass": "1.3kg", "count": "3" }.
     */
    readonly labels?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * A unique identifier for the instance, which cannot be changed after the instance is created. The name must be
     * between 6 and 30 characters in length. If not provided, a random string starting with 'tf-' will be selected.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The number of nodes allocated to this instance.
     */
    readonly numNodes?: pulumi.Input<number>;
    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the provider project is used.
     */
    readonly project?: pulumi.Input<string>;
    /**
     * Instance status: 'CREATING' or 'READY'.
     */
    readonly state?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Instance resource.
 */
export interface InstanceArgs {
    /**
     * The name of the instance's configuration (similar but not quite the same as a region) which defines defines the
     * geographic placement and replication of your databases in this instance. It determines where your data is stored.
     * Values are typically of the form 'regional-europe-west1' , 'us-central' etc. In order to obtain a valid list please
     * consult the [Configuration section of the docs](https://cloud.google.com/spanner/docs/instances).
     */
    readonly config: pulumi.Input<string>;
    /**
     * The descriptive name for this instance as it appears in UIs. Must be unique per project and between 4 and 30
     * characters in length.
     */
    readonly displayName: pulumi.Input<string>;
    /**
     * An object containing a list of "key": value pairs. Example: { "name": "wrench", "mass": "1.3kg", "count": "3" }.
     */
    readonly labels?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * A unique identifier for the instance, which cannot be changed after the instance is created. The name must be
     * between 6 and 30 characters in length. If not provided, a random string starting with 'tf-' will be selected.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The number of nodes allocated to this instance.
     */
    readonly numNodes?: pulumi.Input<number>;
    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the provider project is used.
     */
    readonly project?: pulumi.Input<string>;
}
