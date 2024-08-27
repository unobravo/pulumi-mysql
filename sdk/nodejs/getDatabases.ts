// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * The ``mysql.getDatabases`` gets databases on a MySQL
 * server.
 *
 * ## Example Usage
 *
 * <!--Start PulumiCodeChooser -->
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as mysql from "@pulumi/mysql";
 *
 * const app = mysql.getDatabases({
 *     pattern: "test_%",
 * });
 * ```
 * <!--End PulumiCodeChooser -->
 */
export function getDatabases(args?: GetDatabasesArgs, opts?: pulumi.InvokeOptions): Promise<GetDatabasesResult> {
    args = args || {};

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("mysql:index/getDatabases:getDatabases", {
        "pattern": args.pattern,
    }, opts);
}

/**
 * A collection of arguments for invoking getDatabases.
 */
export interface GetDatabasesArgs {
    /**
     * Patterns for searching databases.
     */
    pattern?: string;
}

/**
 * A collection of values returned by getDatabases.
 */
export interface GetDatabasesResult {
    /**
     * The list of the database names.
     */
    readonly databases: string[];
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly pattern?: string;
}
/**
 * The ``mysql.getDatabases`` gets databases on a MySQL
 * server.
 *
 * ## Example Usage
 *
 * <!--Start PulumiCodeChooser -->
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as mysql from "@pulumi/mysql";
 *
 * const app = mysql.getDatabases({
 *     pattern: "test_%",
 * });
 * ```
 * <!--End PulumiCodeChooser -->
 */
export function getDatabasesOutput(args?: GetDatabasesOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetDatabasesResult> {
    return pulumi.output(args).apply((a: any) => getDatabases(a, opts))
}

/**
 * A collection of arguments for invoking getDatabases.
 */
export interface GetDatabasesOutputArgs {
    /**
     * Patterns for searching databases.
     */
    pattern?: pulumi.Input<string>;
}
