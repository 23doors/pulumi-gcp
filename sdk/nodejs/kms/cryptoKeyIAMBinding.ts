// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

export class CryptoKeyIAMBinding extends pulumi.CustomResource {
    /**
     * Get an existing CryptoKeyIAMBinding resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: CryptoKeyIAMBindingState, opts?: pulumi.CustomResourceOptions): CryptoKeyIAMBinding {
        return new CryptoKeyIAMBinding(name, <any>state, { ...opts, id: id });
    }

    public readonly cryptoKeyId: pulumi.Output<string>;
    public /*out*/ readonly etag: pulumi.Output<string>;
    public readonly members: pulumi.Output<string[]>;
    public readonly role: pulumi.Output<string>;

    /**
     * Create a CryptoKeyIAMBinding resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: CryptoKeyIAMBindingArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: CryptoKeyIAMBindingArgs | CryptoKeyIAMBindingState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state: CryptoKeyIAMBindingState = argsOrState as CryptoKeyIAMBindingState | undefined;
            inputs["cryptoKeyId"] = state ? state.cryptoKeyId : undefined;
            inputs["etag"] = state ? state.etag : undefined;
            inputs["members"] = state ? state.members : undefined;
            inputs["role"] = state ? state.role : undefined;
        } else {
            const args = argsOrState as CryptoKeyIAMBindingArgs | undefined;
            if (!args || args.cryptoKeyId === undefined) {
                throw new Error("Missing required property 'cryptoKeyId'");
            }
            if (!args || args.members === undefined) {
                throw new Error("Missing required property 'members'");
            }
            if (!args || args.role === undefined) {
                throw new Error("Missing required property 'role'");
            }
            inputs["cryptoKeyId"] = args ? args.cryptoKeyId : undefined;
            inputs["members"] = args ? args.members : undefined;
            inputs["role"] = args ? args.role : undefined;
            inputs["etag"] = undefined /*out*/;
        }
        super("gcp:kms/cryptoKeyIAMBinding:CryptoKeyIAMBinding", name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering CryptoKeyIAMBinding resources.
 */
export interface CryptoKeyIAMBindingState {
    readonly cryptoKeyId?: pulumi.Input<string>;
    readonly etag?: pulumi.Input<string>;
    readonly members?: pulumi.Input<pulumi.Input<string>[]>;
    readonly role?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a CryptoKeyIAMBinding resource.
 */
export interface CryptoKeyIAMBindingArgs {
    readonly cryptoKeyId: pulumi.Input<string>;
    readonly members: pulumi.Input<pulumi.Input<string>[]>;
    readonly role: pulumi.Input<string>;
}
