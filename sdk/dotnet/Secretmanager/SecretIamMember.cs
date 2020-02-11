// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.SecretManager
{
    public partial class SecretIamMember : Pulumi.CustomResource
    {
        [Output("condition")]
        public Output<Outputs.SecretIamMemberCondition?> Condition { get; private set; } = null!;

        [Output("etag")]
        public Output<string> Etag { get; private set; } = null!;

        [Output("member")]
        public Output<string> Member { get; private set; } = null!;

        [Output("project")]
        public Output<string> Project { get; private set; } = null!;

        [Output("role")]
        public Output<string> Role { get; private set; } = null!;

        [Output("secretId")]
        public Output<string> SecretId { get; private set; } = null!;


        /// <summary>
        /// Create a SecretIamMember resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public SecretIamMember(string name, SecretIamMemberArgs args, CustomResourceOptions? options = null)
            : base("gcp:secretmanager/secretIamMember:SecretIamMember", name, args ?? ResourceArgs.Empty, MakeResourceOptions(options, ""))
        {
        }

        private SecretIamMember(string name, Input<string> id, SecretIamMemberState? state = null, CustomResourceOptions? options = null)
            : base("gcp:secretmanager/secretIamMember:SecretIamMember", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing SecretIamMember resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static SecretIamMember Get(string name, Input<string> id, SecretIamMemberState? state = null, CustomResourceOptions? options = null)
        {
            return new SecretIamMember(name, id, state, options);
        }
    }

    public sealed class SecretIamMemberArgs : Pulumi.ResourceArgs
    {
        [Input("condition")]
        public Input<Inputs.SecretIamMemberConditionArgs>? Condition { get; set; }

        [Input("member", required: true)]
        public Input<string> Member { get; set; } = null!;

        [Input("project")]
        public Input<string>? Project { get; set; }

        [Input("role", required: true)]
        public Input<string> Role { get; set; } = null!;

        [Input("secretId", required: true)]
        public Input<string> SecretId { get; set; } = null!;

        public SecretIamMemberArgs()
        {
        }
    }

    public sealed class SecretIamMemberState : Pulumi.ResourceArgs
    {
        [Input("condition")]
        public Input<Inputs.SecretIamMemberConditionGetArgs>? Condition { get; set; }

        [Input("etag")]
        public Input<string>? Etag { get; set; }

        [Input("member")]
        public Input<string>? Member { get; set; }

        [Input("project")]
        public Input<string>? Project { get; set; }

        [Input("role")]
        public Input<string>? Role { get; set; }

        [Input("secretId")]
        public Input<string>? SecretId { get; set; }

        public SecretIamMemberState()
        {
        }
    }

    namespace Inputs
    {

    public sealed class SecretIamMemberConditionArgs : Pulumi.ResourceArgs
    {
        [Input("description")]
        public Input<string>? Description { get; set; }

        [Input("expression", required: true)]
        public Input<string> Expression { get; set; } = null!;

        [Input("title", required: true)]
        public Input<string> Title { get; set; } = null!;

        public SecretIamMemberConditionArgs()
        {
        }
    }

    public sealed class SecretIamMemberConditionGetArgs : Pulumi.ResourceArgs
    {
        [Input("description")]
        public Input<string>? Description { get; set; }

        [Input("expression", required: true)]
        public Input<string> Expression { get; set; } = null!;

        [Input("title", required: true)]
        public Input<string> Title { get; set; } = null!;

        public SecretIamMemberConditionGetArgs()
        {
        }
    }
    }

    namespace Outputs
    {

    [OutputType]
    public sealed class SecretIamMemberCondition
    {
        public readonly string? Description;
        public readonly string Expression;
        public readonly string Title;

        [OutputConstructor]
        private SecretIamMemberCondition(
            string? description,
            string expression,
            string title)
        {
            Description = description;
            Expression = expression;
            Title = title;
        }
    }
    }
}
