// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

export class Function extends pulumi.CustomResource {
    /**
     * Get an existing Function resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: FunctionState, opts?: pulumi.CustomResourceOptions): Function {
        return new Function(name, <any>state, { ...opts, id: id });
    }

    public readonly availableMemoryMb: pulumi.Output<number | undefined>;
    public readonly description: pulumi.Output<string | undefined>;
    public readonly entryPoint: pulumi.Output<string | undefined>;
    public readonly environmentVariables: pulumi.Output<{[key: string]: any} | undefined>;
    public readonly eventTrigger: pulumi.Output<{ eventType: string, failurePolicy: { retry: boolean }, resource: string }>;
    public readonly httpsTriggerUrl: pulumi.Output<string>;
    public readonly labels: pulumi.Output<{[key: string]: any} | undefined>;
    public readonly name: pulumi.Output<string>;
    public readonly project: pulumi.Output<string>;
    public readonly region: pulumi.Output<string>;
    public readonly retryOnFailure: pulumi.Output<boolean>;
    public readonly runtime: pulumi.Output<string>;
    public readonly sourceArchiveBucket: pulumi.Output<string>;
    public readonly sourceArchiveObject: pulumi.Output<string>;
    public readonly timeout: pulumi.Output<number | undefined>;
    public readonly triggerBucket: pulumi.Output<string>;
    public readonly triggerHttp: pulumi.Output<boolean | undefined>;
    public readonly triggerTopic: pulumi.Output<string>;

    /**
     * Create a Function resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: FunctionArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: FunctionArgs | FunctionState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state: FunctionState = argsOrState as FunctionState | undefined;
            inputs["availableMemoryMb"] = state ? state.availableMemoryMb : undefined;
            inputs["description"] = state ? state.description : undefined;
            inputs["entryPoint"] = state ? state.entryPoint : undefined;
            inputs["environmentVariables"] = state ? state.environmentVariables : undefined;
            inputs["eventTrigger"] = state ? state.eventTrigger : undefined;
            inputs["httpsTriggerUrl"] = state ? state.httpsTriggerUrl : undefined;
            inputs["labels"] = state ? state.labels : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["project"] = state ? state.project : undefined;
            inputs["region"] = state ? state.region : undefined;
            inputs["retryOnFailure"] = state ? state.retryOnFailure : undefined;
            inputs["runtime"] = state ? state.runtime : undefined;
            inputs["sourceArchiveBucket"] = state ? state.sourceArchiveBucket : undefined;
            inputs["sourceArchiveObject"] = state ? state.sourceArchiveObject : undefined;
            inputs["timeout"] = state ? state.timeout : undefined;
            inputs["triggerBucket"] = state ? state.triggerBucket : undefined;
            inputs["triggerHttp"] = state ? state.triggerHttp : undefined;
            inputs["triggerTopic"] = state ? state.triggerTopic : undefined;
        } else {
            const args = argsOrState as FunctionArgs | undefined;
            if (!args || args.sourceArchiveBucket === undefined) {
                throw new Error("Missing required property 'sourceArchiveBucket'");
            }
            if (!args || args.sourceArchiveObject === undefined) {
                throw new Error("Missing required property 'sourceArchiveObject'");
            }
            inputs["availableMemoryMb"] = args ? args.availableMemoryMb : undefined;
            inputs["description"] = args ? args.description : undefined;
            inputs["entryPoint"] = args ? args.entryPoint : undefined;
            inputs["environmentVariables"] = args ? args.environmentVariables : undefined;
            inputs["eventTrigger"] = args ? args.eventTrigger : undefined;
            inputs["httpsTriggerUrl"] = args ? args.httpsTriggerUrl : undefined;
            inputs["labels"] = args ? args.labels : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["project"] = args ? args.project : undefined;
            inputs["region"] = args ? args.region : undefined;
            inputs["retryOnFailure"] = args ? args.retryOnFailure : undefined;
            inputs["runtime"] = args ? args.runtime : undefined;
            inputs["sourceArchiveBucket"] = args ? args.sourceArchiveBucket : undefined;
            inputs["sourceArchiveObject"] = args ? args.sourceArchiveObject : undefined;
            inputs["timeout"] = args ? args.timeout : undefined;
            inputs["triggerBucket"] = args ? args.triggerBucket : undefined;
            inputs["triggerHttp"] = args ? args.triggerHttp : undefined;
            inputs["triggerTopic"] = args ? args.triggerTopic : undefined;
        }
        super("gcp:cloudfunctions/function:Function", name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Function resources.
 */
export interface FunctionState {
    readonly availableMemoryMb?: pulumi.Input<number>;
    readonly description?: pulumi.Input<string>;
    readonly entryPoint?: pulumi.Input<string>;
    readonly environmentVariables?: pulumi.Input<{[key: string]: any}>;
    readonly eventTrigger?: pulumi.Input<{ eventType: pulumi.Input<string>, failurePolicy?: pulumi.Input<{ retry: pulumi.Input<boolean> }>, resource: pulumi.Input<string> }>;
    readonly httpsTriggerUrl?: pulumi.Input<string>;
    readonly labels?: pulumi.Input<{[key: string]: any}>;
    readonly name?: pulumi.Input<string>;
    readonly project?: pulumi.Input<string>;
    readonly region?: pulumi.Input<string>;
    readonly retryOnFailure?: pulumi.Input<boolean>;
    readonly runtime?: pulumi.Input<string>;
    readonly sourceArchiveBucket?: pulumi.Input<string>;
    readonly sourceArchiveObject?: pulumi.Input<string>;
    readonly timeout?: pulumi.Input<number>;
    readonly triggerBucket?: pulumi.Input<string>;
    readonly triggerHttp?: pulumi.Input<boolean>;
    readonly triggerTopic?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Function resource.
 */
export interface FunctionArgs {
    readonly availableMemoryMb?: pulumi.Input<number>;
    readonly description?: pulumi.Input<string>;
    readonly entryPoint?: pulumi.Input<string>;
    readonly environmentVariables?: pulumi.Input<{[key: string]: any}>;
    readonly eventTrigger?: pulumi.Input<{ eventType: pulumi.Input<string>, failurePolicy?: pulumi.Input<{ retry: pulumi.Input<boolean> }>, resource: pulumi.Input<string> }>;
    readonly httpsTriggerUrl?: pulumi.Input<string>;
    readonly labels?: pulumi.Input<{[key: string]: any}>;
    readonly name?: pulumi.Input<string>;
    readonly project?: pulumi.Input<string>;
    readonly region?: pulumi.Input<string>;
    readonly retryOnFailure?: pulumi.Input<boolean>;
    readonly runtime?: pulumi.Input<string>;
    readonly sourceArchiveBucket: pulumi.Input<string>;
    readonly sourceArchiveObject: pulumi.Input<string>;
    readonly timeout?: pulumi.Input<number>;
    readonly triggerBucket?: pulumi.Input<string>;
    readonly triggerHttp?: pulumi.Input<boolean>;
    readonly triggerTopic?: pulumi.Input<string>;
}
