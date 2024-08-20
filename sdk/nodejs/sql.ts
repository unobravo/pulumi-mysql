// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

export class Sql extends pulumi.CustomResource {
    /**
     * Get an existing Sql resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: SqlState, opts?: pulumi.CustomResourceOptions): Sql {
        return new Sql(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'mysql:index/sql:Sql';

    /**
     * Returns true if the given object is an instance of Sql.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Sql {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Sql.__pulumiType;
    }

    public readonly createSql!: pulumi.Output<string>;
    public readonly deleteSql!: pulumi.Output<string>;
    public readonly name!: pulumi.Output<string>;

    /**
     * Create a Sql resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: SqlArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: SqlArgs | SqlState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as SqlState | undefined;
            resourceInputs["createSql"] = state ? state.createSql : undefined;
            resourceInputs["deleteSql"] = state ? state.deleteSql : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
        } else {
            const args = argsOrState as SqlArgs | undefined;
            if ((!args || args.createSql === undefined) && !opts.urn) {
                throw new Error("Missing required property 'createSql'");
            }
            if ((!args || args.deleteSql === undefined) && !opts.urn) {
                throw new Error("Missing required property 'deleteSql'");
            }
            resourceInputs["createSql"] = args ? args.createSql : undefined;
            resourceInputs["deleteSql"] = args ? args.deleteSql : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Sql.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Sql resources.
 */
export interface SqlState {
    createSql?: pulumi.Input<string>;
    deleteSql?: pulumi.Input<string>;
    name?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Sql resource.
 */
export interface SqlArgs {
    createSql: pulumi.Input<string>;
    deleteSql: pulumi.Input<string>;
    name?: pulumi.Input<string>;
}
