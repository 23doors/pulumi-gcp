// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.BinaryAuthorization
{
    /// <summary>
    /// &gt; This content is derived from https://github.com/terraform-providers/terraform-provider-google/blob/master/website/docs/r/binary_authorization_attestor.html.markdown.
    /// </summary>
    public partial class Attestor : Pulumi.CustomResource
    {
        /// <summary>
        /// A Container Analysis ATTESTATION_AUTHORITY Note, created by the user.
        /// </summary>
        [Output("attestationAuthorityNote")]
        public Output<Outputs.AttestorAttestationAuthorityNote> AttestationAuthorityNote { get; private set; } = null!;

        /// <summary>
        /// A descriptive comment. This field may be updated. The field may be displayed in chooser dialogs.
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// The resource name.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        [Output("project")]
        public Output<string> Project { get; private set; } = null!;


        /// <summary>
        /// Create a Attestor resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Attestor(string name, AttestorArgs args, CustomResourceOptions? options = null)
            : base("gcp:binaryauthorization/attestor:Attestor", name, args ?? ResourceArgs.Empty, MakeResourceOptions(options, ""))
        {
        }

        private Attestor(string name, Input<string> id, AttestorState? state = null, CustomResourceOptions? options = null)
            : base("gcp:binaryauthorization/attestor:Attestor", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Attestor resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Attestor Get(string name, Input<string> id, AttestorState? state = null, CustomResourceOptions? options = null)
        {
            return new Attestor(name, id, state, options);
        }
    }

    public sealed class AttestorArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// A Container Analysis ATTESTATION_AUTHORITY Note, created by the user.
        /// </summary>
        [Input("attestationAuthorityNote", required: true)]
        public Input<Inputs.AttestorAttestationAuthorityNoteArgs> AttestationAuthorityNote { get; set; } = null!;

        /// <summary>
        /// A descriptive comment. This field may be updated. The field may be displayed in chooser dialogs.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// The resource name.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("project")]
        public Input<string>? Project { get; set; }

        public AttestorArgs()
        {
        }
    }

    public sealed class AttestorState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// A Container Analysis ATTESTATION_AUTHORITY Note, created by the user.
        /// </summary>
        [Input("attestationAuthorityNote")]
        public Input<Inputs.AttestorAttestationAuthorityNoteGetArgs>? AttestationAuthorityNote { get; set; }

        /// <summary>
        /// A descriptive comment. This field may be updated. The field may be displayed in chooser dialogs.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// The resource name.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("project")]
        public Input<string>? Project { get; set; }

        public AttestorState()
        {
        }
    }

    namespace Inputs
    {

    public sealed class AttestorAttestationAuthorityNoteArgs : Pulumi.ResourceArgs
    {
        [Input("delegationServiceAccountEmail")]
        public Input<string>? DelegationServiceAccountEmail { get; set; }

        [Input("noteReference", required: true)]
        public Input<string> NoteReference { get; set; } = null!;

        [Input("publicKeys")]
        private InputList<AttestorAttestationAuthorityNotePublicKeysArgs>? _publicKeys;
        public InputList<AttestorAttestationAuthorityNotePublicKeysArgs> PublicKeys
        {
            get => _publicKeys ?? (_publicKeys = new InputList<AttestorAttestationAuthorityNotePublicKeysArgs>());
            set => _publicKeys = value;
        }

        public AttestorAttestationAuthorityNoteArgs()
        {
        }
    }

    public sealed class AttestorAttestationAuthorityNoteGetArgs : Pulumi.ResourceArgs
    {
        [Input("delegationServiceAccountEmail")]
        public Input<string>? DelegationServiceAccountEmail { get; set; }

        [Input("noteReference", required: true)]
        public Input<string> NoteReference { get; set; } = null!;

        [Input("publicKeys")]
        private InputList<AttestorAttestationAuthorityNotePublicKeysGetArgs>? _publicKeys;
        public InputList<AttestorAttestationAuthorityNotePublicKeysGetArgs> PublicKeys
        {
            get => _publicKeys ?? (_publicKeys = new InputList<AttestorAttestationAuthorityNotePublicKeysGetArgs>());
            set => _publicKeys = value;
        }

        public AttestorAttestationAuthorityNoteGetArgs()
        {
        }
    }

    public sealed class AttestorAttestationAuthorityNotePublicKeysArgs : Pulumi.ResourceArgs
    {
        [Input("asciiArmoredPgpPublicKey")]
        public Input<string>? AsciiArmoredPgpPublicKey { get; set; }

        [Input("comment")]
        public Input<string>? Comment { get; set; }

        /// <summary>
        /// an identifier for the resource with format `projects/{{project}}/attestors/{{name}}`
        /// </summary>
        [Input("id")]
        public Input<string>? Id { get; set; }

        [Input("pkixPublicKey")]
        public Input<AttestorAttestationAuthorityNotePublicKeysPkixPublicKeyArgs>? PkixPublicKey { get; set; }

        public AttestorAttestationAuthorityNotePublicKeysArgs()
        {
        }
    }

    public sealed class AttestorAttestationAuthorityNotePublicKeysGetArgs : Pulumi.ResourceArgs
    {
        [Input("asciiArmoredPgpPublicKey")]
        public Input<string>? AsciiArmoredPgpPublicKey { get; set; }

        [Input("comment")]
        public Input<string>? Comment { get; set; }

        /// <summary>
        /// an identifier for the resource with format `projects/{{project}}/attestors/{{name}}`
        /// </summary>
        [Input("id")]
        public Input<string>? Id { get; set; }

        [Input("pkixPublicKey")]
        public Input<AttestorAttestationAuthorityNotePublicKeysPkixPublicKeyGetArgs>? PkixPublicKey { get; set; }

        public AttestorAttestationAuthorityNotePublicKeysGetArgs()
        {
        }
    }

    public sealed class AttestorAttestationAuthorityNotePublicKeysPkixPublicKeyArgs : Pulumi.ResourceArgs
    {
        [Input("publicKeyPem")]
        public Input<string>? PublicKeyPem { get; set; }

        [Input("signatureAlgorithm")]
        public Input<string>? SignatureAlgorithm { get; set; }

        public AttestorAttestationAuthorityNotePublicKeysPkixPublicKeyArgs()
        {
        }
    }

    public sealed class AttestorAttestationAuthorityNotePublicKeysPkixPublicKeyGetArgs : Pulumi.ResourceArgs
    {
        [Input("publicKeyPem")]
        public Input<string>? PublicKeyPem { get; set; }

        [Input("signatureAlgorithm")]
        public Input<string>? SignatureAlgorithm { get; set; }

        public AttestorAttestationAuthorityNotePublicKeysPkixPublicKeyGetArgs()
        {
        }
    }
    }

    namespace Outputs
    {

    [OutputType]
    public sealed class AttestorAttestationAuthorityNote
    {
        public readonly string DelegationServiceAccountEmail;
        public readonly string NoteReference;
        public readonly ImmutableArray<AttestorAttestationAuthorityNotePublicKeys> PublicKeys;

        [OutputConstructor]
        private AttestorAttestationAuthorityNote(
            string delegationServiceAccountEmail,
            string noteReference,
            ImmutableArray<AttestorAttestationAuthorityNotePublicKeys> publicKeys)
        {
            DelegationServiceAccountEmail = delegationServiceAccountEmail;
            NoteReference = noteReference;
            PublicKeys = publicKeys;
        }
    }

    [OutputType]
    public sealed class AttestorAttestationAuthorityNotePublicKeys
    {
        public readonly string? AsciiArmoredPgpPublicKey;
        public readonly string? Comment;
        /// <summary>
        /// an identifier for the resource with format `projects/{{project}}/attestors/{{name}}`
        /// </summary>
        public readonly string Id;
        public readonly AttestorAttestationAuthorityNotePublicKeysPkixPublicKey? PkixPublicKey;

        [OutputConstructor]
        private AttestorAttestationAuthorityNotePublicKeys(
            string? asciiArmoredPgpPublicKey,
            string? comment,
            string id,
            AttestorAttestationAuthorityNotePublicKeysPkixPublicKey? pkixPublicKey)
        {
            AsciiArmoredPgpPublicKey = asciiArmoredPgpPublicKey;
            Comment = comment;
            Id = id;
            PkixPublicKey = pkixPublicKey;
        }
    }

    [OutputType]
    public sealed class AttestorAttestationAuthorityNotePublicKeysPkixPublicKey
    {
        public readonly string? PublicKeyPem;
        public readonly string? SignatureAlgorithm;

        [OutputConstructor]
        private AttestorAttestationAuthorityNotePublicKeysPkixPublicKey(
            string? publicKeyPem,
            string? signatureAlgorithm)
        {
            PublicKeyPem = publicKeyPem;
            SignatureAlgorithm = signatureAlgorithm;
        }
    }
    }
}
