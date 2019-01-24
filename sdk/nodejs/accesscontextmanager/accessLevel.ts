// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

export class AccessLevel extends pulumi.CustomResource {
    /**
     * Get an existing AccessLevel resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: AccessLevelState, opts?: pulumi.CustomResourceOptions): AccessLevel {
        return new AccessLevel(name, <any>state, { ...opts, id: id });
    }

    public readonly basic: pulumi.Output<{ combiningFunction?: string, conditions: { devicePolicy?: { allowedDeviceManagementLevels?: string[], allowedEncryptionStatuses?: string[], osConstraints?: { minimumVersion?: string, osType?: string }[], requireScreenLock?: boolean }, ipSubnetworks?: string[], members?: string[], negate?: boolean, requiredAccessLevels?: string[] }[] } | undefined>;
    public readonly description: pulumi.Output<string | undefined>;
    public readonly name: pulumi.Output<string>;
    public readonly parent: pulumi.Output<string>;
    public readonly title: pulumi.Output<string>;

    /**
     * Create a AccessLevel resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: AccessLevelArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: AccessLevelArgs | AccessLevelState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state: AccessLevelState = argsOrState as AccessLevelState | undefined;
            inputs["basic"] = state ? state.basic : undefined;
            inputs["description"] = state ? state.description : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["parent"] = state ? state.parent : undefined;
            inputs["title"] = state ? state.title : undefined;
        } else {
            const args = argsOrState as AccessLevelArgs | undefined;
            if (!args || args.parent === undefined) {
                throw new Error("Missing required property 'parent'");
            }
            if (!args || args.title === undefined) {
                throw new Error("Missing required property 'title'");
            }
            inputs["basic"] = args ? args.basic : undefined;
            inputs["description"] = args ? args.description : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["parent"] = args ? args.parent : undefined;
            inputs["title"] = args ? args.title : undefined;
        }
        super("gcp:accesscontextmanager/accessLevel:AccessLevel", name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering AccessLevel resources.
 */
export interface AccessLevelState {
    readonly basic?: pulumi.Input<{ combiningFunction?: pulumi.Input<string>, conditions: pulumi.Input<pulumi.Input<{ devicePolicy?: pulumi.Input<{ allowedDeviceManagementLevels?: pulumi.Input<pulumi.Input<string>[]>, allowedEncryptionStatuses?: pulumi.Input<pulumi.Input<string>[]>, osConstraints?: pulumi.Input<pulumi.Input<{ minimumVersion?: pulumi.Input<string>, osType?: pulumi.Input<string> }>[]>, requireScreenLock?: pulumi.Input<boolean> }>, ipSubnetworks?: pulumi.Input<pulumi.Input<string>[]>, members?: pulumi.Input<pulumi.Input<string>[]>, negate?: pulumi.Input<boolean>, requiredAccessLevels?: pulumi.Input<pulumi.Input<string>[]> }>[]> }>;
    readonly description?: pulumi.Input<string>;
    readonly name?: pulumi.Input<string>;
    readonly parent?: pulumi.Input<string>;
    readonly title?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a AccessLevel resource.
 */
export interface AccessLevelArgs {
    readonly basic?: pulumi.Input<{ combiningFunction?: pulumi.Input<string>, conditions: pulumi.Input<pulumi.Input<{ devicePolicy?: pulumi.Input<{ allowedDeviceManagementLevels?: pulumi.Input<pulumi.Input<string>[]>, allowedEncryptionStatuses?: pulumi.Input<pulumi.Input<string>[]>, osConstraints?: pulumi.Input<pulumi.Input<{ minimumVersion?: pulumi.Input<string>, osType?: pulumi.Input<string> }>[]>, requireScreenLock?: pulumi.Input<boolean> }>, ipSubnetworks?: pulumi.Input<pulumi.Input<string>[]>, members?: pulumi.Input<pulumi.Input<string>[]>, negate?: pulumi.Input<boolean>, requiredAccessLevels?: pulumi.Input<pulumi.Input<string>[]> }>[]> }>;
    readonly description?: pulumi.Input<string>;
    readonly name?: pulumi.Input<string>;
    readonly parent: pulumi.Input<string>;
    readonly title: pulumi.Input<string>;
}