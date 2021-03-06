// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

export function getTensorflowVersions(args?: GetTensorflowVersionsArgs, opts?: pulumi.InvokeOptions): Promise<GetTensorflowVersionsResult> & GetTensorflowVersionsResult {
    args = args || {};
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    const promise: Promise<GetTensorflowVersionsResult> = pulumi.runtime.invoke("gcp:tpu/getTensorflowVersions:getTensorflowVersions", {
        "project": args.project,
        "zone": args.zone,
    }, opts);

    return pulumi.utils.liftProperties(promise, opts);
}

/**
 * A collection of arguments for invoking getTensorflowVersions.
 */
export interface GetTensorflowVersionsArgs {
    readonly project?: string;
    readonly zone?: string;
}

/**
 * A collection of values returned by getTensorflowVersions.
 */
export interface GetTensorflowVersionsResult {
    readonly project: string;
    readonly versions: string[];
    readonly zone: string;
    /**
     * id is the provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
}
