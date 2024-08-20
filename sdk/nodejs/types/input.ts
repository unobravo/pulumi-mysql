// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";

export interface ProviderAwsConfig {
    accessKey?: pulumi.Input<string>;
    profile?: pulumi.Input<string>;
    region?: pulumi.Input<string>;
    secretKey?: pulumi.Input<string>;
}

export interface ProviderAzureConfig {
    clientId?: pulumi.Input<string>;
    clientSecret?: pulumi.Input<string>;
    environment?: pulumi.Input<string>;
    tenantId?: pulumi.Input<string>;
}

export interface ProviderCustomTl {
    caCert: pulumi.Input<string>;
    clientCert: pulumi.Input<string>;
    clientKey: pulumi.Input<string>;
    configKey?: pulumi.Input<string>;
}

export interface UserAadIdentity {
    identity: pulumi.Input<string>;
    type?: pulumi.Input<string>;
}
export namespace config {
}
