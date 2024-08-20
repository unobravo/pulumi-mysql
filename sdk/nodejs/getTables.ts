// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * The ``mysql.getTables`` gets tables on a MySQL
 * server.
 *
 * ## Example Usage
 *
 * <!--Start PulumiCodeChooser -->
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as mysql from "@pulumi/mysql";
 *
 * const app = mysql.getTables({
 *     database: "my_awesome_app",
 * });
 * ```
 * <!--End PulumiCodeChooser -->
 */
export function getTables(args: GetTablesArgs, opts?: pulumi.InvokeOptions): Promise<GetTablesResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("mysql:index/getTables:getTables", {
        "database": args.database,
        "pattern": args.pattern,
    }, opts);
}

/**
 * A collection of arguments for invoking getTables.
 */
export interface GetTablesArgs {
    /**
     * The name of the database.
     */
    database: string;
    /**
     * Patterns for searching tables.
     */
    pattern?: string;
}

/**
 * A collection of values returned by getTables.
 */
export interface GetTablesResult {
    readonly database: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly pattern?: string;
    /**
     * The list of the table names.
     */
    readonly tables: string[];
}
/**
 * The ``mysql.getTables`` gets tables on a MySQL
 * server.
 *
 * ## Example Usage
 *
 * <!--Start PulumiCodeChooser -->
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as mysql from "@pulumi/mysql";
 *
 * const app = mysql.getTables({
 *     database: "my_awesome_app",
 * });
 * ```
 * <!--End PulumiCodeChooser -->
 */
export function getTablesOutput(args: GetTablesOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetTablesResult> {
    return pulumi.output(args).apply((a: any) => getTables(a, opts))
}

/**
 * A collection of arguments for invoking getTables.
 */
export interface GetTablesOutputArgs {
    /**
     * The name of the database.
     */
    database: pulumi.Input<string>;
    /**
     * Patterns for searching tables.
     */
    pattern?: pulumi.Input<string>;
}
