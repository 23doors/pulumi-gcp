// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * > This content is derived from https://github.com/terraform-providers/terraform-provider-google/blob/master/website/docs/r/compute_url_map.html.markdown.
 */
export class URLMap extends pulumi.CustomResource {
    /**
     * Get an existing URLMap resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: URLMapState, opts?: pulumi.CustomResourceOptions): URLMap {
        return new URLMap(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'gcp:compute/uRLMap:URLMap';

    /**
     * Returns true if the given object is an instance of URLMap.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is URLMap {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === URLMap.__pulumiType;
    }

    public /*out*/ readonly creationTimestamp!: pulumi.Output<string>;
    public readonly defaultService!: pulumi.Output<string>;
    public readonly description!: pulumi.Output<string | undefined>;
    public /*out*/ readonly fingerprint!: pulumi.Output<string>;
    public readonly hostRules!: pulumi.Output<outputs.compute.URLMapHostRule[] | undefined>;
    public /*out*/ readonly mapId!: pulumi.Output<number>;
    public readonly name!: pulumi.Output<string>;
    public readonly pathMatchers!: pulumi.Output<outputs.compute.URLMapPathMatcher[] | undefined>;
    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the provider project is used.
     */
    public readonly project!: pulumi.Output<string>;
    /**
     * The URI of the created resource.
     */
    public /*out*/ readonly selfLink!: pulumi.Output<string>;
    public readonly tests!: pulumi.Output<outputs.compute.URLMapTest[] | undefined>;

    /**
     * Create a URLMap resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: URLMapArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: URLMapArgs | URLMapState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as URLMapState | undefined;
            inputs["creationTimestamp"] = state ? state.creationTimestamp : undefined;
            inputs["defaultService"] = state ? state.defaultService : undefined;
            inputs["description"] = state ? state.description : undefined;
            inputs["fingerprint"] = state ? state.fingerprint : undefined;
            inputs["hostRules"] = state ? state.hostRules : undefined;
            inputs["mapId"] = state ? state.mapId : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["pathMatchers"] = state ? state.pathMatchers : undefined;
            inputs["project"] = state ? state.project : undefined;
            inputs["selfLink"] = state ? state.selfLink : undefined;
            inputs["tests"] = state ? state.tests : undefined;
        } else {
            const args = argsOrState as URLMapArgs | undefined;
            if (!args || args.defaultService === undefined) {
                throw new Error("Missing required property 'defaultService'");
            }
            inputs["defaultService"] = args ? args.defaultService : undefined;
            inputs["description"] = args ? args.description : undefined;
            inputs["hostRules"] = args ? args.hostRules : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["pathMatchers"] = args ? args.pathMatchers : undefined;
            inputs["project"] = args ? args.project : undefined;
            inputs["tests"] = args ? args.tests : undefined;
            inputs["creationTimestamp"] = undefined /*out*/;
            inputs["fingerprint"] = undefined /*out*/;
            inputs["mapId"] = undefined /*out*/;
            inputs["selfLink"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(URLMap.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering URLMap resources.
 */
export interface URLMapState {
    readonly creationTimestamp?: pulumi.Input<string>;
    readonly defaultService?: pulumi.Input<string>;
    readonly description?: pulumi.Input<string>;
    readonly fingerprint?: pulumi.Input<string>;
    readonly hostRules?: pulumi.Input<pulumi.Input<inputs.compute.URLMapHostRule>[]>;
    readonly mapId?: pulumi.Input<number>;
    readonly name?: pulumi.Input<string>;
    readonly pathMatchers?: pulumi.Input<pulumi.Input<inputs.compute.URLMapPathMatcher>[]>;
    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the provider project is used.
     */
    readonly project?: pulumi.Input<string>;
    /**
     * The URI of the created resource.
     */
    readonly selfLink?: pulumi.Input<string>;
    readonly tests?: pulumi.Input<pulumi.Input<inputs.compute.URLMapTest>[]>;
}

/**
 * The set of arguments for constructing a URLMap resource.
 */
export interface URLMapArgs {
    readonly defaultService: pulumi.Input<string>;
    readonly description?: pulumi.Input<string>;
    readonly hostRules?: pulumi.Input<pulumi.Input<inputs.compute.URLMapHostRule>[]>;
    readonly name?: pulumi.Input<string>;
    readonly pathMatchers?: pulumi.Input<pulumi.Input<inputs.compute.URLMapPathMatcher>[]>;
    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the provider project is used.
     */
    readonly project?: pulumi.Input<string>;
    readonly tests?: pulumi.Input<pulumi.Input<inputs.compute.URLMapTest>[]>;
}
