// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * > This content is derived from https://github.com/terraform-providers/terraform-provider-google/blob/master/website/docs/r/binary_authorization_policy.html.markdown.
 */
export class Policy extends pulumi.CustomResource {
    /**
     * Get an existing Policy resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: PolicyState, opts?: pulumi.CustomResourceOptions): Policy {
        return new Policy(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'gcp:binaryauthorization/policy:Policy';

    /**
     * Returns true if the given object is an instance of Policy.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Policy {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Policy.__pulumiType;
    }

    public readonly admissionWhitelistPatterns!: pulumi.Output<outputs.binaryauthorization.PolicyAdmissionWhitelistPattern[] | undefined>;
    public readonly clusterAdmissionRules!: pulumi.Output<outputs.binaryauthorization.PolicyClusterAdmissionRule[] | undefined>;
    public readonly defaultAdmissionRule!: pulumi.Output<outputs.binaryauthorization.PolicyDefaultAdmissionRule>;
    public readonly description!: pulumi.Output<string | undefined>;
    public readonly globalPolicyEvaluationMode!: pulumi.Output<string>;
    public readonly project!: pulumi.Output<string>;

    /**
     * Create a Policy resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: PolicyArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: PolicyArgs | PolicyState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as PolicyState | undefined;
            inputs["admissionWhitelistPatterns"] = state ? state.admissionWhitelistPatterns : undefined;
            inputs["clusterAdmissionRules"] = state ? state.clusterAdmissionRules : undefined;
            inputs["defaultAdmissionRule"] = state ? state.defaultAdmissionRule : undefined;
            inputs["description"] = state ? state.description : undefined;
            inputs["globalPolicyEvaluationMode"] = state ? state.globalPolicyEvaluationMode : undefined;
            inputs["project"] = state ? state.project : undefined;
        } else {
            const args = argsOrState as PolicyArgs | undefined;
            if (!args || args.defaultAdmissionRule === undefined) {
                throw new Error("Missing required property 'defaultAdmissionRule'");
            }
            inputs["admissionWhitelistPatterns"] = args ? args.admissionWhitelistPatterns : undefined;
            inputs["clusterAdmissionRules"] = args ? args.clusterAdmissionRules : undefined;
            inputs["defaultAdmissionRule"] = args ? args.defaultAdmissionRule : undefined;
            inputs["description"] = args ? args.description : undefined;
            inputs["globalPolicyEvaluationMode"] = args ? args.globalPolicyEvaluationMode : undefined;
            inputs["project"] = args ? args.project : undefined;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(Policy.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Policy resources.
 */
export interface PolicyState {
    readonly admissionWhitelistPatterns?: pulumi.Input<pulumi.Input<inputs.binaryauthorization.PolicyAdmissionWhitelistPattern>[]>;
    readonly clusterAdmissionRules?: pulumi.Input<pulumi.Input<inputs.binaryauthorization.PolicyClusterAdmissionRule>[]>;
    readonly defaultAdmissionRule?: pulumi.Input<inputs.binaryauthorization.PolicyDefaultAdmissionRule>;
    readonly description?: pulumi.Input<string>;
    readonly globalPolicyEvaluationMode?: pulumi.Input<string>;
    readonly project?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Policy resource.
 */
export interface PolicyArgs {
    readonly admissionWhitelistPatterns?: pulumi.Input<pulumi.Input<inputs.binaryauthorization.PolicyAdmissionWhitelistPattern>[]>;
    readonly clusterAdmissionRules?: pulumi.Input<pulumi.Input<inputs.binaryauthorization.PolicyClusterAdmissionRule>[]>;
    readonly defaultAdmissionRule: pulumi.Input<inputs.binaryauthorization.PolicyDefaultAdmissionRule>;
    readonly description?: pulumi.Input<string>;
    readonly globalPolicyEvaluationMode?: pulumi.Input<string>;
    readonly project?: pulumi.Input<string>;
}
