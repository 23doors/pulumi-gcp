// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * An HttpHealthCheck resource. This resource defines a template for how
 * individual VMs should be checked for health, via HTTP.
 * 
 * 
 * > **Note:** google_compute_http_health_check is a legacy health check.
 * The newer [google_compute_health_check](https://www.terraform.io/docs/providers/google/r/compute_health_check.html)
 * should be preferred for all uses except
 * [Network Load Balancers](https://cloud.google.com/compute/docs/load-balancing/network/)
 * which still require the legacy version.
 * 
 * 
 * To get more information about HttpHealthCheck, see:
 * 
 * * [API documentation](https://cloud.google.com/compute/docs/reference/v1/httpHealthChecks)
 * * How-to Guides
 *     * [Adding Health Checks](https://cloud.google.com/compute/docs/load-balancing/health-checks#legacy_health_checks)
 * 
 * <div class = "oics-button" style="float: right; margin: 0 0 -15px">
 *   <a href="https://console.cloud.google.com/cloudshell/open?cloudshell_git_repo=https%3A%2F%2Fgithub.com%2Fterraform-google-modules%2Fdocs-examples.git&cloudshell_working_dir=http_health_check_basic&cloudshell_image=gcr.io%2Fgraphite-cloud-shell-images%2Fterraform%3Alatest&open_in_editor=main.tf&cloudshell_print=.%2Fmotd&cloudshell_tutorial=.%2Ftutorial.md" target="_blank">
 *     <img alt="Open in Cloud Shell" src="//gstatic.com/cloudssh/images/open-btn.svg" style="max-height: 44px; margin: 32px auto; max-width: 100%;">
 *   </a>
 * </div>
 * ## Example Usage - Http Health Check Basic
 * 
 * 
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 * 
 * const defaultHttpHealthCheck = new gcp.compute.HttpHealthCheck("default", {
 *     checkIntervalSec: 1,
 *     requestPath: "/health_check",
 *     timeoutSec: 1,
 * });
 * ```
 */
export class HttpHealthCheck extends pulumi.CustomResource {
    /**
     * Get an existing HttpHealthCheck resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: HttpHealthCheckState, opts?: pulumi.CustomResourceOptions): HttpHealthCheck {
        return new HttpHealthCheck(name, <any>state, { ...opts, id: id });
    }

    public readonly checkIntervalSec: pulumi.Output<number | undefined>;
    public /*out*/ readonly creationTimestamp: pulumi.Output<string>;
    public readonly description: pulumi.Output<string | undefined>;
    public readonly healthyThreshold: pulumi.Output<number | undefined>;
    public readonly host: pulumi.Output<string | undefined>;
    public readonly name: pulumi.Output<string>;
    public readonly port: pulumi.Output<number | undefined>;
    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the provider project is used.
     */
    public readonly project: pulumi.Output<string>;
    public readonly requestPath: pulumi.Output<string | undefined>;
    /**
     * The URI of the created resource.
     */
    public /*out*/ readonly selfLink: pulumi.Output<string>;
    public readonly timeoutSec: pulumi.Output<number | undefined>;
    public readonly unhealthyThreshold: pulumi.Output<number | undefined>;

    /**
     * Create a HttpHealthCheck resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: HttpHealthCheckArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: HttpHealthCheckArgs | HttpHealthCheckState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state: HttpHealthCheckState = argsOrState as HttpHealthCheckState | undefined;
            inputs["checkIntervalSec"] = state ? state.checkIntervalSec : undefined;
            inputs["creationTimestamp"] = state ? state.creationTimestamp : undefined;
            inputs["description"] = state ? state.description : undefined;
            inputs["healthyThreshold"] = state ? state.healthyThreshold : undefined;
            inputs["host"] = state ? state.host : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["port"] = state ? state.port : undefined;
            inputs["project"] = state ? state.project : undefined;
            inputs["requestPath"] = state ? state.requestPath : undefined;
            inputs["selfLink"] = state ? state.selfLink : undefined;
            inputs["timeoutSec"] = state ? state.timeoutSec : undefined;
            inputs["unhealthyThreshold"] = state ? state.unhealthyThreshold : undefined;
        } else {
            const args = argsOrState as HttpHealthCheckArgs | undefined;
            inputs["checkIntervalSec"] = args ? args.checkIntervalSec : undefined;
            inputs["description"] = args ? args.description : undefined;
            inputs["healthyThreshold"] = args ? args.healthyThreshold : undefined;
            inputs["host"] = args ? args.host : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["port"] = args ? args.port : undefined;
            inputs["project"] = args ? args.project : undefined;
            inputs["requestPath"] = args ? args.requestPath : undefined;
            inputs["timeoutSec"] = args ? args.timeoutSec : undefined;
            inputs["unhealthyThreshold"] = args ? args.unhealthyThreshold : undefined;
            inputs["creationTimestamp"] = undefined /*out*/;
            inputs["selfLink"] = undefined /*out*/;
        }
        super("gcp:compute/httpHealthCheck:HttpHealthCheck", name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering HttpHealthCheck resources.
 */
export interface HttpHealthCheckState {
    readonly checkIntervalSec?: pulumi.Input<number>;
    readonly creationTimestamp?: pulumi.Input<string>;
    readonly description?: pulumi.Input<string>;
    readonly healthyThreshold?: pulumi.Input<number>;
    readonly host?: pulumi.Input<string>;
    readonly name?: pulumi.Input<string>;
    readonly port?: pulumi.Input<number>;
    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the provider project is used.
     */
    readonly project?: pulumi.Input<string>;
    readonly requestPath?: pulumi.Input<string>;
    /**
     * The URI of the created resource.
     */
    readonly selfLink?: pulumi.Input<string>;
    readonly timeoutSec?: pulumi.Input<number>;
    readonly unhealthyThreshold?: pulumi.Input<number>;
}

/**
 * The set of arguments for constructing a HttpHealthCheck resource.
 */
export interface HttpHealthCheckArgs {
    readonly checkIntervalSec?: pulumi.Input<number>;
    readonly description?: pulumi.Input<string>;
    readonly healthyThreshold?: pulumi.Input<number>;
    readonly host?: pulumi.Input<string>;
    readonly name?: pulumi.Input<string>;
    readonly port?: pulumi.Input<number>;
    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the provider project is used.
     */
    readonly project?: pulumi.Input<string>;
    readonly requestPath?: pulumi.Input<string>;
    readonly timeoutSec?: pulumi.Input<number>;
    readonly unhealthyThreshold?: pulumi.Input<number>;
}
