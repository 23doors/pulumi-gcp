// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";

export class TargetTCPProxy extends pulumi.CustomResource {
    /**
     * Get an existing TargetTCPProxy resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: TargetTCPProxyState): TargetTCPProxy {
        return new TargetTCPProxy(name, <any>state, { id });
    }

    public readonly backendService: pulumi.Output<string>;
    public /*out*/ readonly creationTimestamp: pulumi.Output<string>;
    public readonly description: pulumi.Output<string | undefined>;
    public readonly name: pulumi.Output<string>;
    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the provider project is used.
     */
    public readonly project: pulumi.Output<string>;
    public readonly proxyHeader: pulumi.Output<string | undefined>;
    public /*out*/ readonly proxyId: pulumi.Output<number>;
    /**
     * The URI of the created resource.
     */
    public /*out*/ readonly selfLink: pulumi.Output<string>;

    /**
     * Create a TargetTCPProxy resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: TargetTCPProxyArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: TargetTCPProxyArgs | TargetTCPProxyState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state: TargetTCPProxyState = argsOrState as TargetTCPProxyState | undefined;
            inputs["backendService"] = state ? state.backendService : undefined;
            inputs["creationTimestamp"] = state ? state.creationTimestamp : undefined;
            inputs["description"] = state ? state.description : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["project"] = state ? state.project : undefined;
            inputs["proxyHeader"] = state ? state.proxyHeader : undefined;
            inputs["proxyId"] = state ? state.proxyId : undefined;
            inputs["selfLink"] = state ? state.selfLink : undefined;
        } else {
            const args = argsOrState as TargetTCPProxyArgs | undefined;
            if (!args || args.backendService === undefined) {
                throw new Error("Missing required property 'backendService'");
            }
            inputs["backendService"] = args ? args.backendService : undefined;
            inputs["description"] = args ? args.description : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["project"] = args ? args.project : undefined;
            inputs["proxyHeader"] = args ? args.proxyHeader : undefined;
            inputs["creationTimestamp"] = undefined /*out*/;
            inputs["proxyId"] = undefined /*out*/;
            inputs["selfLink"] = undefined /*out*/;
        }
        super("gcp:compute/targetTCPProxy:TargetTCPProxy", name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering TargetTCPProxy resources.
 */
export interface TargetTCPProxyState {
    readonly backendService?: pulumi.Input<string>;
    readonly creationTimestamp?: pulumi.Input<string>;
    readonly description?: pulumi.Input<string>;
    readonly name?: pulumi.Input<string>;
    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the provider project is used.
     */
    readonly project?: pulumi.Input<string>;
    readonly proxyHeader?: pulumi.Input<string>;
    readonly proxyId?: pulumi.Input<number>;
    /**
     * The URI of the created resource.
     */
    readonly selfLink?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a TargetTCPProxy resource.
 */
export interface TargetTCPProxyArgs {
    readonly backendService: pulumi.Input<string>;
    readonly description?: pulumi.Input<string>;
    readonly name?: pulumi.Input<string>;
    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the provider project is used.
     */
    readonly project?: pulumi.Input<string>;
    readonly proxyHeader?: pulumi.Input<string>;
}
