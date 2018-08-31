// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";

/**
 * Creates a table resource in a dataset for Google BigQuery. For more information see
 * [the official documentation](https://cloud.google.com/bigquery/docs/) and
 * [API](https://cloud.google.com/bigquery/docs/reference/rest/v2/tables).
 * 
 */
export class Table extends pulumi.CustomResource {
    /**
     * Get an existing Table resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: TableState): Table {
        return new Table(name, <any>state, { id });
    }

    /**
     * The time when this table was created, in milliseconds since the epoch.
     */
    public /*out*/ readonly creationTime: pulumi.Output<number>;
    /**
     * The dataset ID to create the table in.
     * Changing this forces a new resource to be created.
     */
    public readonly datasetId: pulumi.Output<string>;
    /**
     * The field description.
     */
    public readonly description: pulumi.Output<string | undefined>;
    /**
     * A hash of the resource.
     */
    public /*out*/ readonly etag: pulumi.Output<string>;
    /**
     * The time when this table expires, in
     * milliseconds since the epoch. If not present, the table will persist
     * indefinitely. Expired tables will be deleted and their storage
     * reclaimed.
     */
    public readonly expirationTime: pulumi.Output<number>;
    /**
     * A descriptive name for the table.
     */
    public readonly friendlyName: pulumi.Output<string | undefined>;
    /**
     * A mapping of labels to assign to the resource.
     */
    public readonly labels: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * The time when this table was last modified, in milliseconds since the epoch.
     */
    public /*out*/ readonly lastModifiedTime: pulumi.Output<number>;
    /**
     * The geographic location where the table resides. This value is inherited from the dataset.
     */
    public /*out*/ readonly location: pulumi.Output<string>;
    /**
     * The size of this table in bytes, excluding any data in the streaming buffer.
     */
    public /*out*/ readonly numBytes: pulumi.Output<number>;
    /**
     * The number of bytes in the table that are considered "long-term storage".
     */
    public /*out*/ readonly numLongTermBytes: pulumi.Output<number>;
    /**
     * The number of rows of data in this table, excluding any data in the streaming buffer.
     */
    public /*out*/ readonly numRows: pulumi.Output<number>;
    /**
     * The ID of the project in which the resource belongs. If it
     * is not provided, the provider project is used.
     */
    public readonly project: pulumi.Output<string>;
    /**
     * A JSON schema for the table.
     */
    public readonly schema: pulumi.Output<string>;
    /**
     * The URI of the created resource.
     */
    public /*out*/ readonly selfLink: pulumi.Output<string>;
    /**
     * A unique ID for the resource.
     * Changing this forces a new resource to be created.
     */
    public readonly tableId: pulumi.Output<string>;
    /**
     * If specified, configures time-based
     * partitioning for this table. Structure is documented below.
     */
    public readonly timePartitioning: pulumi.Output<{ expirationMs?: number, field?: string, type: string } | undefined>;
    /**
     * Describes the table type.
     */
    public /*out*/ readonly type: pulumi.Output<string>;
    /**
     * If specified, configures this table as a view.
     * Structure is documented below.
     */
    public readonly view: pulumi.Output<{ query: string, useLegacySql?: boolean } | undefined>;

    /**
     * Create a Table resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: TableArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: TableArgs | TableState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state: TableState = argsOrState as TableState | undefined;
            inputs["creationTime"] = state ? state.creationTime : undefined;
            inputs["datasetId"] = state ? state.datasetId : undefined;
            inputs["description"] = state ? state.description : undefined;
            inputs["etag"] = state ? state.etag : undefined;
            inputs["expirationTime"] = state ? state.expirationTime : undefined;
            inputs["friendlyName"] = state ? state.friendlyName : undefined;
            inputs["labels"] = state ? state.labels : undefined;
            inputs["lastModifiedTime"] = state ? state.lastModifiedTime : undefined;
            inputs["location"] = state ? state.location : undefined;
            inputs["numBytes"] = state ? state.numBytes : undefined;
            inputs["numLongTermBytes"] = state ? state.numLongTermBytes : undefined;
            inputs["numRows"] = state ? state.numRows : undefined;
            inputs["project"] = state ? state.project : undefined;
            inputs["schema"] = state ? state.schema : undefined;
            inputs["selfLink"] = state ? state.selfLink : undefined;
            inputs["tableId"] = state ? state.tableId : undefined;
            inputs["timePartitioning"] = state ? state.timePartitioning : undefined;
            inputs["type"] = state ? state.type : undefined;
            inputs["view"] = state ? state.view : undefined;
        } else {
            const args = argsOrState as TableArgs | undefined;
            if (!args || args.datasetId === undefined) {
                throw new Error("Missing required property 'datasetId'");
            }
            if (!args || args.tableId === undefined) {
                throw new Error("Missing required property 'tableId'");
            }
            inputs["datasetId"] = args ? args.datasetId : undefined;
            inputs["description"] = args ? args.description : undefined;
            inputs["expirationTime"] = args ? args.expirationTime : undefined;
            inputs["friendlyName"] = args ? args.friendlyName : undefined;
            inputs["labels"] = args ? args.labels : undefined;
            inputs["project"] = args ? args.project : undefined;
            inputs["schema"] = args ? args.schema : undefined;
            inputs["tableId"] = args ? args.tableId : undefined;
            inputs["timePartitioning"] = args ? args.timePartitioning : undefined;
            inputs["view"] = args ? args.view : undefined;
            inputs["creationTime"] = undefined /*out*/;
            inputs["etag"] = undefined /*out*/;
            inputs["lastModifiedTime"] = undefined /*out*/;
            inputs["location"] = undefined /*out*/;
            inputs["numBytes"] = undefined /*out*/;
            inputs["numLongTermBytes"] = undefined /*out*/;
            inputs["numRows"] = undefined /*out*/;
            inputs["selfLink"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
        }
        super("gcp:bigquery/table:Table", name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Table resources.
 */
export interface TableState {
    /**
     * The time when this table was created, in milliseconds since the epoch.
     */
    readonly creationTime?: pulumi.Input<number>;
    /**
     * The dataset ID to create the table in.
     * Changing this forces a new resource to be created.
     */
    readonly datasetId?: pulumi.Input<string>;
    /**
     * The field description.
     */
    readonly description?: pulumi.Input<string>;
    /**
     * A hash of the resource.
     */
    readonly etag?: pulumi.Input<string>;
    /**
     * The time when this table expires, in
     * milliseconds since the epoch. If not present, the table will persist
     * indefinitely. Expired tables will be deleted and their storage
     * reclaimed.
     */
    readonly expirationTime?: pulumi.Input<number>;
    /**
     * A descriptive name for the table.
     */
    readonly friendlyName?: pulumi.Input<string>;
    /**
     * A mapping of labels to assign to the resource.
     */
    readonly labels?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * The time when this table was last modified, in milliseconds since the epoch.
     */
    readonly lastModifiedTime?: pulumi.Input<number>;
    /**
     * The geographic location where the table resides. This value is inherited from the dataset.
     */
    readonly location?: pulumi.Input<string>;
    /**
     * The size of this table in bytes, excluding any data in the streaming buffer.
     */
    readonly numBytes?: pulumi.Input<number>;
    /**
     * The number of bytes in the table that are considered "long-term storage".
     */
    readonly numLongTermBytes?: pulumi.Input<number>;
    /**
     * The number of rows of data in this table, excluding any data in the streaming buffer.
     */
    readonly numRows?: pulumi.Input<number>;
    /**
     * The ID of the project in which the resource belongs. If it
     * is not provided, the provider project is used.
     */
    readonly project?: pulumi.Input<string>;
    /**
     * A JSON schema for the table.
     */
    readonly schema?: pulumi.Input<string>;
    /**
     * The URI of the created resource.
     */
    readonly selfLink?: pulumi.Input<string>;
    /**
     * A unique ID for the resource.
     * Changing this forces a new resource to be created.
     */
    readonly tableId?: pulumi.Input<string>;
    /**
     * If specified, configures time-based
     * partitioning for this table. Structure is documented below.
     */
    readonly timePartitioning?: pulumi.Input<{ expirationMs?: pulumi.Input<number>, field?: pulumi.Input<string>, type: pulumi.Input<string> }>;
    /**
     * Describes the table type.
     */
    readonly type?: pulumi.Input<string>;
    /**
     * If specified, configures this table as a view.
     * Structure is documented below.
     */
    readonly view?: pulumi.Input<{ query: pulumi.Input<string>, useLegacySql?: pulumi.Input<boolean> }>;
}

/**
 * The set of arguments for constructing a Table resource.
 */
export interface TableArgs {
    /**
     * The dataset ID to create the table in.
     * Changing this forces a new resource to be created.
     */
    readonly datasetId: pulumi.Input<string>;
    /**
     * The field description.
     */
    readonly description?: pulumi.Input<string>;
    /**
     * The time when this table expires, in
     * milliseconds since the epoch. If not present, the table will persist
     * indefinitely. Expired tables will be deleted and their storage
     * reclaimed.
     */
    readonly expirationTime?: pulumi.Input<number>;
    /**
     * A descriptive name for the table.
     */
    readonly friendlyName?: pulumi.Input<string>;
    /**
     * A mapping of labels to assign to the resource.
     */
    readonly labels?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * The ID of the project in which the resource belongs. If it
     * is not provided, the provider project is used.
     */
    readonly project?: pulumi.Input<string>;
    /**
     * A JSON schema for the table.
     */
    readonly schema?: pulumi.Input<string>;
    /**
     * A unique ID for the resource.
     * Changing this forces a new resource to be created.
     */
    readonly tableId: pulumi.Input<string>;
    /**
     * If specified, configures time-based
     * partitioning for this table. Structure is documented below.
     */
    readonly timePartitioning?: pulumi.Input<{ expirationMs?: pulumi.Input<number>, field?: pulumi.Input<string>, type: pulumi.Input<string> }>;
    /**
     * If specified, configures this table as a view.
     * Structure is documented below.
     */
    readonly view?: pulumi.Input<{ query: pulumi.Input<string>, useLegacySql?: pulumi.Input<boolean> }>;
}
