// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Mysql.Inputs
{

    public sealed class ProviderAzureConfigArgs : global::Pulumi.ResourceArgs
    {
        [Input("clientId")]
        public Input<string>? ClientId { get; set; }

        [Input("clientSecret")]
        public Input<string>? ClientSecret { get; set; }

        [Input("environment")]
        public Input<string>? Environment { get; set; }

        [Input("tenantId")]
        public Input<string>? TenantId { get; set; }

        public ProviderAzureConfigArgs()
        {
        }
        public static new ProviderAzureConfigArgs Empty => new ProviderAzureConfigArgs();
    }
}
