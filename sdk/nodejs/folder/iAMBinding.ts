// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";

/**
 * Allows creation and management of a single binding within IAM policy for
 * an existing Google Cloud Platform folder.
 * 
 * ~> **Note:** This resource _must not_ be used in conjunction with
 *    `google_folder_iam_policy` or they will fight over what your policy
 *    should be.
 */
export class IAMBinding extends pulumi.CustomResource {
    /**
     * Get an existing IAMBinding resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: IAMBindingState): IAMBinding {
        return new IAMBinding(name, <any>state, { id });
    }

    /**
     * (Computed) The etag of the folder's IAM policy.
     */
    public /*out*/ readonly etag: pulumi.Output<string>;
    /**
     * The resource name of the folder the policy is attached to. Its format is folders/{folder_id}.
     */
    public readonly folder: pulumi.Output<string>;
    /**
     * An array of identites that will be granted the privilege in the `role`.
     * Each entry can have one of the following values:
     * * **user:{emailid}**: An email address that represents a specific Google account. For example, alice@gmail.com or joe@example.com.
     * * **serviceAccount:{emailid}**: An email address that represents a service account. For example, my-other-app@appspot.gserviceaccount.com.
     * * **group:{emailid}**: An email address that represents a Google group. For example, admins@example.com.
     * * **domain:{domain}**: A Google Apps domain name that represents all the users of that domain. For example, google.com or example.com.
     */
    public readonly members: pulumi.Output<string[]>;
    /**
     * The role that should be applied. Only one
     * `google_folder_iam_binding` can be used per role. Note that custom roles must be of the format
     * `[projects|organizations]/{parent-name}/roles/{role-name}`.
     */
    public readonly role: pulumi.Output<string>;

    /**
     * Create a IAMBinding resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: IAMBindingArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: IAMBindingArgs | IAMBindingState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state: IAMBindingState = argsOrState as IAMBindingState | undefined;
            inputs["etag"] = state ? state.etag : undefined;
            inputs["folder"] = state ? state.folder : undefined;
            inputs["members"] = state ? state.members : undefined;
            inputs["role"] = state ? state.role : undefined;
        } else {
            const args = argsOrState as IAMBindingArgs | undefined;
            if (!args || args.folder === undefined) {
                throw new Error("Missing required property 'folder'");
            }
            if (!args || args.members === undefined) {
                throw new Error("Missing required property 'members'");
            }
            if (!args || args.role === undefined) {
                throw new Error("Missing required property 'role'");
            }
            inputs["folder"] = args ? args.folder : undefined;
            inputs["members"] = args ? args.members : undefined;
            inputs["role"] = args ? args.role : undefined;
            inputs["etag"] = undefined /*out*/;
        }
        super("gcp:folder/iAMBinding:IAMBinding", name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering IAMBinding resources.
 */
export interface IAMBindingState {
    /**
     * (Computed) The etag of the folder's IAM policy.
     */
    readonly etag?: pulumi.Input<string>;
    /**
     * The resource name of the folder the policy is attached to. Its format is folders/{folder_id}.
     */
    readonly folder?: pulumi.Input<string>;
    /**
     * An array of identites that will be granted the privilege in the `role`.
     * Each entry can have one of the following values:
     * * **user:{emailid}**: An email address that represents a specific Google account. For example, alice@gmail.com or joe@example.com.
     * * **serviceAccount:{emailid}**: An email address that represents a service account. For example, my-other-app@appspot.gserviceaccount.com.
     * * **group:{emailid}**: An email address that represents a Google group. For example, admins@example.com.
     * * **domain:{domain}**: A Google Apps domain name that represents all the users of that domain. For example, google.com or example.com.
     */
    readonly members?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The role that should be applied. Only one
     * `google_folder_iam_binding` can be used per role. Note that custom roles must be of the format
     * `[projects|organizations]/{parent-name}/roles/{role-name}`.
     */
    readonly role?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a IAMBinding resource.
 */
export interface IAMBindingArgs {
    /**
     * The resource name of the folder the policy is attached to. Its format is folders/{folder_id}.
     */
    readonly folder: pulumi.Input<string>;
    /**
     * An array of identites that will be granted the privilege in the `role`.
     * Each entry can have one of the following values:
     * * **user:{emailid}**: An email address that represents a specific Google account. For example, alice@gmail.com or joe@example.com.
     * * **serviceAccount:{emailid}**: An email address that represents a service account. For example, my-other-app@appspot.gserviceaccount.com.
     * * **group:{emailid}**: An email address that represents a Google group. For example, admins@example.com.
     * * **domain:{domain}**: A Google Apps domain name that represents all the users of that domain. For example, google.com or example.com.
     */
    readonly members: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The role that should be applied. Only one
     * `google_folder_iam_binding` can be used per role. Note that custom roles must be of the format
     * `[projects|organizations]/{parent-name}/roles/{role-name}`.
     */
    readonly role: pulumi.Input<string>;
}
