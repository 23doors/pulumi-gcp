// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";

/**
 * Enables the Google Compute Engine
 * [Shared VPC](https://cloud.google.com/compute/docs/shared-vpc)
 * feature for a project, assigning it as a Shared VPC service project associated
 * with a given host project.
 * 
 * For more information, see,
 * [the Project API documentation](https://cloud.google.com/compute/docs/reference/latest/projects),
 * where the Shared VPC feature is referred to by its former name "XPN".
 */
export class SharedVPCServiceProject extends pulumi.CustomResource {
    /**
     * Get an existing SharedVPCServiceProject resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: SharedVPCServiceProjectState): SharedVPCServiceProject {
        return new SharedVPCServiceProject(name, <any>state, { id });
    }

    /**
     * The ID of a host project to associate.
     */
    public readonly hostProject: pulumi.Output<string>;
    /**
     * The ID of the project that will serve as a Shared VPC service project.
     */
    public readonly serviceProject: pulumi.Output<string>;

    /**
     * Create a SharedVPCServiceProject resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: SharedVPCServiceProjectArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: SharedVPCServiceProjectArgs | SharedVPCServiceProjectState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state: SharedVPCServiceProjectState = argsOrState as SharedVPCServiceProjectState | undefined;
            inputs["hostProject"] = state ? state.hostProject : undefined;
            inputs["serviceProject"] = state ? state.serviceProject : undefined;
        } else {
            const args = argsOrState as SharedVPCServiceProjectArgs | undefined;
            if (!args || args.hostProject === undefined) {
                throw new Error("Missing required property 'hostProject'");
            }
            if (!args || args.serviceProject === undefined) {
                throw new Error("Missing required property 'serviceProject'");
            }
            inputs["hostProject"] = args ? args.hostProject : undefined;
            inputs["serviceProject"] = args ? args.serviceProject : undefined;
        }
        super("gcp:compute/sharedVPCServiceProject:SharedVPCServiceProject", name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering SharedVPCServiceProject resources.
 */
export interface SharedVPCServiceProjectState {
    /**
     * The ID of a host project to associate.
     */
    readonly hostProject?: pulumi.Input<string>;
    /**
     * The ID of the project that will serve as a Shared VPC service project.
     */
    readonly serviceProject?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a SharedVPCServiceProject resource.
 */
export interface SharedVPCServiceProjectArgs {
    /**
     * The ID of a host project to associate.
     */
    readonly hostProject: pulumi.Input<string>;
    /**
     * The ID of the project that will serve as a Shared VPC service project.
     */
    readonly serviceProject: pulumi.Input<string>;
}
