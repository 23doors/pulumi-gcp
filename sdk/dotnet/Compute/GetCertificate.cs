// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.Compute
{
    public static partial class Invokes
    {
        /// <summary>
        /// Get info about a Google Compute SSL Certificate from its name.
        /// 
        /// &gt; This content is derived from https://github.com/terraform-providers/terraform-provider-google/blob/master/website/docs/d/compute_ssl_certificate.html.markdown.
        /// </summary>
        public static Task<GetCertificateResult> GetCertificate(GetCertificateArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetCertificateResult>("gcp:compute/getCertificate:getCertificate", args ?? InvokeArgs.Empty, options.WithVersion());
    }

    public sealed class GetCertificateArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the certificate.
        /// </summary>
        [Input("name", required: true)]
        public string Name { get; set; } = null!;

        /// <summary>
        /// The project in which the resource belongs. If it
        /// is not provided, the provider project is used.
        /// </summary>
        [Input("project")]
        public string? Project { get; set; }

        public GetCertificateArgs()
        {
        }
    }

    [OutputType]
    public sealed class GetCertificateResult
    {
        public readonly string Certificate;
        public readonly int CertificateId;
        public readonly string CreationTimestamp;
        public readonly string Description;
        public readonly string Name;
        public readonly string NamePrefix;
        public readonly string PrivateKey;
        public readonly string? Project;
        public readonly string SelfLink;
        /// <summary>
        /// id is the provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;

        [OutputConstructor]
        private GetCertificateResult(
            string certificate,
            int certificateId,
            string creationTimestamp,
            string description,
            string name,
            string namePrefix,
            string privateKey,
            string? project,
            string selfLink,
            string id)
        {
            Certificate = certificate;
            CertificateId = certificateId;
            CreationTimestamp = creationTimestamp;
            Description = description;
            Name = name;
            NamePrefix = namePrefix;
            PrivateKey = privateKey;
            Project = project;
            SelfLink = selfLink;
            Id = id;
        }
    }
}
