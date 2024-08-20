// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * The ``mysql.Database`` resource creates and manages a database on a MySQL
 * server.
 *
 * > **Caution:** The ``mysql.Database`` resource can completely delete your
 * database just as easily as it can create it. To avoid costly accidents,
 * consider setting
 * [``preventDestroy``](https://www.terraform.io/docs/configuration/resources.html#prevent_destroy)
 * on your database resources as an extra safety measure.
 *
 * ## Example Usage
 *
 * <!--Start PulumiCodeChooser -->
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as mysql from "@pulumi/mysql";
 *
 * const app = new mysql.Database("app", {});
 * ```
 * <!--End PulumiCodeChooser -->
 *
 * ## Import
 *
 * Databases can be imported using their name, e.g.
 *
 * ```sh
 * $ pulumi import mysql:index/database:Database example my-example-database
 * ```
 */
export class Database extends pulumi.CustomResource {
    /**
     * Get an existing Database resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: DatabaseState, opts?: pulumi.CustomResourceOptions): Database {
        return new Database(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'mysql:index/database:Database';

    /**
     * Returns true if the given object is an instance of Database.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Database {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Database.__pulumiType;
    }

    /**
     * The default character set to use when
     * a table is created without specifying an explicit character set. Defaults
     * to ``utf8mb4``.
     */
    public readonly defaultCharacterSet!: pulumi.Output<string | undefined>;
    public readonly defaultCollation!: pulumi.Output<string | undefined>;
    /**
     * The name of the database. This must be unique within
     * a given MySQL server and may or may not be case-sensitive depending on
     * the operating system on which the MySQL server is running.
     */
    public readonly name!: pulumi.Output<string>;

    /**
     * Create a Database resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: DatabaseArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: DatabaseArgs | DatabaseState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as DatabaseState | undefined;
            resourceInputs["defaultCharacterSet"] = state ? state.defaultCharacterSet : undefined;
            resourceInputs["defaultCollation"] = state ? state.defaultCollation : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
        } else {
            const args = argsOrState as DatabaseArgs | undefined;
            resourceInputs["defaultCharacterSet"] = args ? args.defaultCharacterSet : undefined;
            resourceInputs["defaultCollation"] = args ? args.defaultCollation : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Database.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Database resources.
 */
export interface DatabaseState {
    /**
     * The default character set to use when
     * a table is created without specifying an explicit character set. Defaults
     * to ``utf8mb4``.
     */
    defaultCharacterSet?: pulumi.Input<string>;
    defaultCollation?: pulumi.Input<string>;
    /**
     * The name of the database. This must be unique within
     * a given MySQL server and may or may not be case-sensitive depending on
     * the operating system on which the MySQL server is running.
     */
    name?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Database resource.
 */
export interface DatabaseArgs {
    /**
     * The default character set to use when
     * a table is created without specifying an explicit character set. Defaults
     * to ``utf8mb4``.
     */
    defaultCharacterSet?: pulumi.Input<string>;
    defaultCollation?: pulumi.Input<string>;
    /**
     * The name of the database. This must be unique within
     * a given MySQL server and may or may not be case-sensitive depending on
     * the operating system on which the MySQL server is running.
     */
    name?: pulumi.Input<string>;
}
