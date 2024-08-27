// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

// Export members:
export { DatabaseArgs, DatabaseState } from "./database";
export type Database = import("./database").Database;
export const Database: typeof import("./database").Database = null as any;
utilities.lazyLoad(exports, ["Database"], () => require("./database"));

export { DefaultRolesArgs, DefaultRolesState } from "./defaultRoles";
export type DefaultRoles = import("./defaultRoles").DefaultRoles;
export const DefaultRoles: typeof import("./defaultRoles").DefaultRoles = null as any;
utilities.lazyLoad(exports, ["DefaultRoles"], () => require("./defaultRoles"));

export { GetDatabasesArgs, GetDatabasesResult, GetDatabasesOutputArgs } from "./getDatabases";
export const getDatabases: typeof import("./getDatabases").getDatabases = null as any;
export const getDatabasesOutput: typeof import("./getDatabases").getDatabasesOutput = null as any;
utilities.lazyLoad(exports, ["getDatabases","getDatabasesOutput"], () => require("./getDatabases"));

export { GetTablesArgs, GetTablesResult, GetTablesOutputArgs } from "./getTables";
export const getTables: typeof import("./getTables").getTables = null as any;
export const getTablesOutput: typeof import("./getTables").getTablesOutput = null as any;
utilities.lazyLoad(exports, ["getTables","getTablesOutput"], () => require("./getTables"));

export { GlobalVariableArgs, GlobalVariableState } from "./globalVariable";
export type GlobalVariable = import("./globalVariable").GlobalVariable;
export const GlobalVariable: typeof import("./globalVariable").GlobalVariable = null as any;
utilities.lazyLoad(exports, ["GlobalVariable"], () => require("./globalVariable"));

export { GrantArgs, GrantState } from "./grant";
export type Grant = import("./grant").Grant;
export const Grant: typeof import("./grant").Grant = null as any;
utilities.lazyLoad(exports, ["Grant"], () => require("./grant"));

export { ProviderArgs } from "./provider";
export type Provider = import("./provider").Provider;
export const Provider: typeof import("./provider").Provider = null as any;
utilities.lazyLoad(exports, ["Provider"], () => require("./provider"));

export { RdsConfigArgs, RdsConfigState } from "./rdsConfig";
export type RdsConfig = import("./rdsConfig").RdsConfig;
export const RdsConfig: typeof import("./rdsConfig").RdsConfig = null as any;
utilities.lazyLoad(exports, ["RdsConfig"], () => require("./rdsConfig"));

export { RoleArgs, RoleState } from "./role";
export type Role = import("./role").Role;
export const Role: typeof import("./role").Role = null as any;
utilities.lazyLoad(exports, ["Role"], () => require("./role"));

export { SqlArgs, SqlState } from "./sql";
export type Sql = import("./sql").Sql;
export const Sql: typeof import("./sql").Sql = null as any;
utilities.lazyLoad(exports, ["Sql"], () => require("./sql"));

export { TiConfigArgs, TiConfigState } from "./tiConfig";
export type TiConfig = import("./tiConfig").TiConfig;
export const TiConfig: typeof import("./tiConfig").TiConfig = null as any;
utilities.lazyLoad(exports, ["TiConfig"], () => require("./tiConfig"));

export { UserArgs, UserState } from "./user";
export type User = import("./user").User;
export const User: typeof import("./user").User = null as any;
utilities.lazyLoad(exports, ["User"], () => require("./user"));

export { UserPasswordArgs, UserPasswordState } from "./userPassword";
export type UserPassword = import("./userPassword").UserPassword;
export const UserPassword: typeof import("./userPassword").UserPassword = null as any;
utilities.lazyLoad(exports, ["UserPassword"], () => require("./userPassword"));


// Export sub-modules:
import * as config from "./config";
import * as types from "./types";

export {
    config,
    types,
};

const _module = {
    version: utilities.getVersion(),
    construct: (name: string, type: string, urn: string): pulumi.Resource => {
        switch (type) {
            case "mysql:index/database:Database":
                return new Database(name, <any>undefined, { urn })
            case "mysql:index/defaultRoles:DefaultRoles":
                return new DefaultRoles(name, <any>undefined, { urn })
            case "mysql:index/globalVariable:GlobalVariable":
                return new GlobalVariable(name, <any>undefined, { urn })
            case "mysql:index/grant:Grant":
                return new Grant(name, <any>undefined, { urn })
            case "mysql:index/rdsConfig:RdsConfig":
                return new RdsConfig(name, <any>undefined, { urn })
            case "mysql:index/role:Role":
                return new Role(name, <any>undefined, { urn })
            case "mysql:index/sql:Sql":
                return new Sql(name, <any>undefined, { urn })
            case "mysql:index/tiConfig:TiConfig":
                return new TiConfig(name, <any>undefined, { urn })
            case "mysql:index/user:User":
                return new User(name, <any>undefined, { urn })
            case "mysql:index/userPassword:UserPassword":
                return new UserPassword(name, <any>undefined, { urn })
            default:
                throw new Error(`unknown resource type ${type}`);
        }
    },
};
pulumi.runtime.registerResourceModule("mysql", "index/database", _module)
pulumi.runtime.registerResourceModule("mysql", "index/defaultRoles", _module)
pulumi.runtime.registerResourceModule("mysql", "index/globalVariable", _module)
pulumi.runtime.registerResourceModule("mysql", "index/grant", _module)
pulumi.runtime.registerResourceModule("mysql", "index/rdsConfig", _module)
pulumi.runtime.registerResourceModule("mysql", "index/role", _module)
pulumi.runtime.registerResourceModule("mysql", "index/sql", _module)
pulumi.runtime.registerResourceModule("mysql", "index/tiConfig", _module)
pulumi.runtime.registerResourceModule("mysql", "index/user", _module)
pulumi.runtime.registerResourceModule("mysql", "index/userPassword", _module)
pulumi.runtime.registerResourcePackage("mysql", {
    version: utilities.getVersion(),
    constructProvider: (name: string, type: string, urn: string): pulumi.ProviderResource => {
        if (type !== "pulumi:providers:mysql") {
            throw new Error(`unknown provider type ${type}`);
        }
        return new Provider(name, <any>undefined, { urn });
    },
});
