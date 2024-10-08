// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Mysql
{
    public static class GetTables
    {
        /// <summary>
        /// The ``mysql.getTables`` gets tables on a MySQL
        /// server.
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
        ///     var app = Mysql.GetTables.Invoke(new()
        ///     {
        ///         Database = "my_awesome_app",
        ///     });
        /// 
        /// });
        /// ```
        /// &lt;!--End PulumiCodeChooser --&gt;
        /// </summary>
        public static Task<GetTablesResult> InvokeAsync(GetTablesArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetTablesResult>("mysql:index/getTables:getTables", args ?? new GetTablesArgs(), options.WithDefaults());

        /// <summary>
        /// The ``mysql.getTables`` gets tables on a MySQL
        /// server.
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
        ///     var app = Mysql.GetTables.Invoke(new()
        ///     {
        ///         Database = "my_awesome_app",
        ///     });
        /// 
        /// });
        /// ```
        /// &lt;!--End PulumiCodeChooser --&gt;
        /// </summary>
        public static Output<GetTablesResult> Invoke(GetTablesInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetTablesResult>("mysql:index/getTables:getTables", args ?? new GetTablesInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetTablesArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the database.
        /// </summary>
        [Input("database", required: true)]
        public string Database { get; set; } = null!;

        /// <summary>
        /// Patterns for searching tables.
        /// </summary>
        [Input("pattern")]
        public string? Pattern { get; set; }

        public GetTablesArgs()
        {
        }
        public static new GetTablesArgs Empty => new GetTablesArgs();
    }

    public sealed class GetTablesInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the database.
        /// </summary>
        [Input("database", required: true)]
        public Input<string> Database { get; set; } = null!;

        /// <summary>
        /// Patterns for searching tables.
        /// </summary>
        [Input("pattern")]
        public Input<string>? Pattern { get; set; }

        public GetTablesInvokeArgs()
        {
        }
        public static new GetTablesInvokeArgs Empty => new GetTablesInvokeArgs();
    }


    [OutputType]
    public sealed class GetTablesResult
    {
        public readonly string Database;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly string? Pattern;
        /// <summary>
        /// The list of the table names.
        /// </summary>
        public readonly ImmutableArray<string> Tables;

        [OutputConstructor]
        private GetTablesResult(
            string database,

            string id,

            string? pattern,

            ImmutableArray<string> tables)
        {
            Database = database;
            Id = id;
            Pattern = pattern;
            Tables = tables;
        }
    }
}
