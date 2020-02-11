// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.Sql
{
    /// <summary>
    /// &gt; This content is derived from https://github.com/terraform-providers/terraform-provider-google/blob/master/website/docs/r/sql_database.html.markdown.
    /// </summary>
    public partial class Database : Pulumi.CustomResource
    {
        /// <summary>
        /// The charset value. See MySQL's [Supported Character Sets and
        /// Collations](https://dev.mysql.com/doc/refman/5.7/en/charset-charsets.html) and Postgres' [Character Set
        /// Support](https://www.postgresql.org/docs/9.6/static/multibyte.html) for more details and supported values.
        /// Postgres databases only support a value of 'UTF8' at creation time.
        /// </summary>
        [Output("charset")]
        public Output<string> Charset { get; private set; } = null!;

        /// <summary>
        /// The collation value. See MySQL's [Supported Character Sets and
        /// Collations](https://dev.mysql.com/doc/refman/5.7/en/charset-charsets.html) and Postgres' [Collation
        /// Support](https://www.postgresql.org/docs/9.6/static/collation.html) for more details and supported values.
        /// Postgres databases only support a value of 'en_US.UTF8' at creation time.
        /// </summary>
        [Output("collation")]
        public Output<string> Collation { get; private set; } = null!;

        /// <summary>
        /// The name of the Cloud SQL instance. This does not include the project ID.
        /// </summary>
        [Output("instance")]
        public Output<string> Instance { get; private set; } = null!;

        /// <summary>
        /// The name of the database in the Cloud SQL instance. This does not include the project ID or instance name.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The ID of the project in which the resource belongs.
        /// If it is not provided, the provider project is used.
        /// </summary>
        [Output("project")]
        public Output<string> Project { get; private set; } = null!;

        /// <summary>
        /// The URI of the created resource.
        /// </summary>
        [Output("selfLink")]
        public Output<string> SelfLink { get; private set; } = null!;


        /// <summary>
        /// Create a Database resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Database(string name, DatabaseArgs args, CustomResourceOptions? options = null)
            : base("gcp:sql/database:Database", name, args ?? ResourceArgs.Empty, MakeResourceOptions(options, ""))
        {
        }

        private Database(string name, Input<string> id, DatabaseState? state = null, CustomResourceOptions? options = null)
            : base("gcp:sql/database:Database", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing Database resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Database Get(string name, Input<string> id, DatabaseState? state = null, CustomResourceOptions? options = null)
        {
            return new Database(name, id, state, options);
        }
    }

    public sealed class DatabaseArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The charset value. See MySQL's [Supported Character Sets and
        /// Collations](https://dev.mysql.com/doc/refman/5.7/en/charset-charsets.html) and Postgres' [Character Set
        /// Support](https://www.postgresql.org/docs/9.6/static/multibyte.html) for more details and supported values.
        /// Postgres databases only support a value of 'UTF8' at creation time.
        /// </summary>
        [Input("charset")]
        public Input<string>? Charset { get; set; }

        /// <summary>
        /// The collation value. See MySQL's [Supported Character Sets and
        /// Collations](https://dev.mysql.com/doc/refman/5.7/en/charset-charsets.html) and Postgres' [Collation
        /// Support](https://www.postgresql.org/docs/9.6/static/collation.html) for more details and supported values.
        /// Postgres databases only support a value of 'en_US.UTF8' at creation time.
        /// </summary>
        [Input("collation")]
        public Input<string>? Collation { get; set; }

        /// <summary>
        /// The name of the Cloud SQL instance. This does not include the project ID.
        /// </summary>
        [Input("instance", required: true)]
        public Input<string> Instance { get; set; } = null!;

        /// <summary>
        /// The name of the database in the Cloud SQL instance. This does not include the project ID or instance name.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The ID of the project in which the resource belongs.
        /// If it is not provided, the provider project is used.
        /// </summary>
        [Input("project")]
        public Input<string>? Project { get; set; }

        public DatabaseArgs()
        {
        }
    }

    public sealed class DatabaseState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The charset value. See MySQL's [Supported Character Sets and
        /// Collations](https://dev.mysql.com/doc/refman/5.7/en/charset-charsets.html) and Postgres' [Character Set
        /// Support](https://www.postgresql.org/docs/9.6/static/multibyte.html) for more details and supported values.
        /// Postgres databases only support a value of 'UTF8' at creation time.
        /// </summary>
        [Input("charset")]
        public Input<string>? Charset { get; set; }

        /// <summary>
        /// The collation value. See MySQL's [Supported Character Sets and
        /// Collations](https://dev.mysql.com/doc/refman/5.7/en/charset-charsets.html) and Postgres' [Collation
        /// Support](https://www.postgresql.org/docs/9.6/static/collation.html) for more details and supported values.
        /// Postgres databases only support a value of 'en_US.UTF8' at creation time.
        /// </summary>
        [Input("collation")]
        public Input<string>? Collation { get; set; }

        /// <summary>
        /// The name of the Cloud SQL instance. This does not include the project ID.
        /// </summary>
        [Input("instance")]
        public Input<string>? Instance { get; set; }

        /// <summary>
        /// The name of the database in the Cloud SQL instance. This does not include the project ID or instance name.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The ID of the project in which the resource belongs.
        /// If it is not provided, the provider project is used.
        /// </summary>
        [Input("project")]
        public Input<string>? Project { get; set; }

        /// <summary>
        /// The URI of the created resource.
        /// </summary>
        [Input("selfLink")]
        public Input<string>? SelfLink { get; set; }

        public DatabaseState()
        {
        }
    }
}
