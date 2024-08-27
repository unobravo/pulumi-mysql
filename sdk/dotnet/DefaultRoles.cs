// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Mysql
{
    /// <summary>
    /// The ``mysql.DefaultRoles`` resource creates and manages a user's default roles on a MySQL server.
    /// 
    /// &gt; **Note:** This resource is available on MySQL version 8.0.0 and later.
    /// 
    /// ## Example Usage
    /// 
    /// &lt;!--Start PulumiCodeChooser --&gt;
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using Mysql = Pulumi.Mysql;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var @readonly = new Mysql.Role("readonly");
    /// 
    ///     var jdoeUser = new Mysql.User("jdoeUser", new()
    ///     {
    ///         Username = "jdoe",
    ///         Host = "%",
    ///     });
    /// 
    ///     var jdoeGrant = new Mysql.Grant("jdoeGrant", new()
    ///     {
    ///         User = jdoeUser.Username,
    ///         Host = jdoeUser.Host,
    ///         Database = "",
    ///         Roles = new[]
    ///         {
    ///             @readonly.Name,
    ///         },
    ///     });
    /// 
    ///     var jdoeDefaultRoles = new Mysql.DefaultRoles("jdoeDefaultRoles", new()
    ///     {
    ///         User = jdoeUser.Username,
    ///         Host = jdoeUser.Host,
    ///         Roles = jdoeGrant.Roles,
    ///     });
    /// 
    /// });
    /// ```
    /// &lt;!--End PulumiCodeChooser --&gt;
    /// 
    /// ## Import
    /// 
    /// User default roles can be imported using user and host.
    /// 
    /// ```sh
    /// $ pulumi import mysql:index/defaultRoles:DefaultRoles example user@host
    /// ```
    /// </summary>
    [MysqlResourceType("mysql:index/defaultRoles:DefaultRoles")]
    public partial class DefaultRoles : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The source host of the user. Defaults to "localhost".
        /// </summary>
        [Output("host")]
        public Output<string?> Host { get; private set; } = null!;

        /// <summary>
        /// A list of default roles to assign to the user. By default no roles are assigned.
        /// 
        /// &gt; **Note:** Creating a new default roles resource on an existing user will **overwrite** the user's existing default roles. Likewise, destryoing a default roles resource will **remove** the user's default roles, equivalent to running `ALTER USER ... DEFAULT ROLE NONE`.
        /// </summary>
        [Output("roles")]
        public Output<ImmutableArray<string>> Roles { get; private set; } = null!;

        /// <summary>
        /// The name of the user.
        /// </summary>
        [Output("user")]
        public Output<string> User { get; private set; } = null!;


        /// <summary>
        /// Create a DefaultRoles resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public DefaultRoles(string name, DefaultRolesArgs args, CustomResourceOptions? options = null)
            : base("mysql:index/defaultRoles:DefaultRoles", name, args ?? new DefaultRolesArgs(), MakeResourceOptions(options, ""))
        {
        }

        private DefaultRoles(string name, Input<string> id, DefaultRolesState? state = null, CustomResourceOptions? options = null)
            : base("mysql:index/defaultRoles:DefaultRoles", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                PluginDownloadURL = "github://api.github.com/unobravo/pulumi-mysql",
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing DefaultRoles resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static DefaultRoles Get(string name, Input<string> id, DefaultRolesState? state = null, CustomResourceOptions? options = null)
        {
            return new DefaultRoles(name, id, state, options);
        }
    }

    public sealed class DefaultRolesArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The source host of the user. Defaults to "localhost".
        /// </summary>
        [Input("host")]
        public Input<string>? Host { get; set; }

        [Input("roles", required: true)]
        private InputList<string>? _roles;

        /// <summary>
        /// A list of default roles to assign to the user. By default no roles are assigned.
        /// 
        /// &gt; **Note:** Creating a new default roles resource on an existing user will **overwrite** the user's existing default roles. Likewise, destryoing a default roles resource will **remove** the user's default roles, equivalent to running `ALTER USER ... DEFAULT ROLE NONE`.
        /// </summary>
        public InputList<string> Roles
        {
            get => _roles ?? (_roles = new InputList<string>());
            set => _roles = value;
        }

        /// <summary>
        /// The name of the user.
        /// </summary>
        [Input("user", required: true)]
        public Input<string> User { get; set; } = null!;

        public DefaultRolesArgs()
        {
        }
        public static new DefaultRolesArgs Empty => new DefaultRolesArgs();
    }

    public sealed class DefaultRolesState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The source host of the user. Defaults to "localhost".
        /// </summary>
        [Input("host")]
        public Input<string>? Host { get; set; }

        [Input("roles")]
        private InputList<string>? _roles;

        /// <summary>
        /// A list of default roles to assign to the user. By default no roles are assigned.
        /// 
        /// &gt; **Note:** Creating a new default roles resource on an existing user will **overwrite** the user's existing default roles. Likewise, destryoing a default roles resource will **remove** the user's default roles, equivalent to running `ALTER USER ... DEFAULT ROLE NONE`.
        /// </summary>
        public InputList<string> Roles
        {
            get => _roles ?? (_roles = new InputList<string>());
            set => _roles = value;
        }

        /// <summary>
        /// The name of the user.
        /// </summary>
        [Input("user")]
        public Input<string>? User { get; set; }

        public DefaultRolesState()
        {
        }
        public static new DefaultRolesState Empty => new DefaultRolesState();
    }
}
