// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Mysql.Outputs
{

    [OutputType]
    public sealed class UserAadIdentity
    {
        public readonly string Identity;
        public readonly string? Type;

        [OutputConstructor]
        private UserAadIdentity(
            string identity,

            string? type)
        {
            Identity = identity;
            Type = type;
        }
    }
}
