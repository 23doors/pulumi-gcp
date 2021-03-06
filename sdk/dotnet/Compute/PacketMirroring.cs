// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.Compute
{
    /// <summary>
    /// &gt; This content is derived from https://github.com/terraform-providers/terraform-provider-google/blob/master/website/docs/r/compute_packet_mirroring.html.markdown.
    /// </summary>
    public partial class PacketMirroring : Pulumi.CustomResource
    {
        /// <summary>
        /// The Forwarding Rule resource (of type load_balancing_scheme=INTERNAL) that will be used as collector for
        /// mirrored traffic. The specified forwarding rule must have is_mirroring_collector set to true.
        /// </summary>
        [Output("collectorIlb")]
        public Output<Outputs.PacketMirroringCollectorIlb> CollectorIlb { get; private set; } = null!;

        /// <summary>
        /// A human-readable description of the rule.
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// A filter for mirrored traffic. If unset, all traffic is mirrored.
        /// </summary>
        [Output("filter")]
        public Output<Outputs.PacketMirroringFilter?> Filter { get; private set; } = null!;

        /// <summary>
        /// A means of specifying which resources to mirror.
        /// </summary>
        [Output("mirroredResources")]
        public Output<Outputs.PacketMirroringMirroredResources> MirroredResources { get; private set; } = null!;

        /// <summary>
        /// The name of the packet mirroring rule
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Specifies the mirrored VPC network. Only packets in this network will be mirrored. All mirrored VMs should
        /// have a NIC in the given network. All mirrored subnetworks should belong to the given network.
        /// </summary>
        [Output("network")]
        public Output<Outputs.PacketMirroringNetwork> Network { get; private set; } = null!;

        /// <summary>
        /// Since only one rule can be active at a time, priority is used to break ties in the case of two rules that
        /// apply to the same instances.
        /// </summary>
        [Output("priority")]
        public Output<int> Priority { get; private set; } = null!;

        [Output("project")]
        public Output<string> Project { get; private set; } = null!;

        /// <summary>
        /// The Region in which the created address should reside. If it is not provided, the provider region is used.
        /// </summary>
        [Output("region")]
        public Output<string> Region { get; private set; } = null!;


        /// <summary>
        /// Create a PacketMirroring resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public PacketMirroring(string name, PacketMirroringArgs args, CustomResourceOptions? options = null)
            : base("gcp:compute/packetMirroring:PacketMirroring", name, args ?? ResourceArgs.Empty, MakeResourceOptions(options, ""))
        {
        }

        private PacketMirroring(string name, Input<string> id, PacketMirroringState? state = null, CustomResourceOptions? options = null)
            : base("gcp:compute/packetMirroring:PacketMirroring", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing PacketMirroring resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static PacketMirroring Get(string name, Input<string> id, PacketMirroringState? state = null, CustomResourceOptions? options = null)
        {
            return new PacketMirroring(name, id, state, options);
        }
    }

    public sealed class PacketMirroringArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The Forwarding Rule resource (of type load_balancing_scheme=INTERNAL) that will be used as collector for
        /// mirrored traffic. The specified forwarding rule must have is_mirroring_collector set to true.
        /// </summary>
        [Input("collectorIlb", required: true)]
        public Input<Inputs.PacketMirroringCollectorIlbArgs> CollectorIlb { get; set; } = null!;

        /// <summary>
        /// A human-readable description of the rule.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// A filter for mirrored traffic. If unset, all traffic is mirrored.
        /// </summary>
        [Input("filter")]
        public Input<Inputs.PacketMirroringFilterArgs>? Filter { get; set; }

        /// <summary>
        /// A means of specifying which resources to mirror.
        /// </summary>
        [Input("mirroredResources", required: true)]
        public Input<Inputs.PacketMirroringMirroredResourcesArgs> MirroredResources { get; set; } = null!;

        /// <summary>
        /// The name of the packet mirroring rule
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Specifies the mirrored VPC network. Only packets in this network will be mirrored. All mirrored VMs should
        /// have a NIC in the given network. All mirrored subnetworks should belong to the given network.
        /// </summary>
        [Input("network", required: true)]
        public Input<Inputs.PacketMirroringNetworkArgs> Network { get; set; } = null!;

        /// <summary>
        /// Since only one rule can be active at a time, priority is used to break ties in the case of two rules that
        /// apply to the same instances.
        /// </summary>
        [Input("priority")]
        public Input<int>? Priority { get; set; }

        [Input("project")]
        public Input<string>? Project { get; set; }

        /// <summary>
        /// The Region in which the created address should reside. If it is not provided, the provider region is used.
        /// </summary>
        [Input("region")]
        public Input<string>? Region { get; set; }

        public PacketMirroringArgs()
        {
        }
    }

    public sealed class PacketMirroringState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The Forwarding Rule resource (of type load_balancing_scheme=INTERNAL) that will be used as collector for
        /// mirrored traffic. The specified forwarding rule must have is_mirroring_collector set to true.
        /// </summary>
        [Input("collectorIlb")]
        public Input<Inputs.PacketMirroringCollectorIlbGetArgs>? CollectorIlb { get; set; }

        /// <summary>
        /// A human-readable description of the rule.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// A filter for mirrored traffic. If unset, all traffic is mirrored.
        /// </summary>
        [Input("filter")]
        public Input<Inputs.PacketMirroringFilterGetArgs>? Filter { get; set; }

        /// <summary>
        /// A means of specifying which resources to mirror.
        /// </summary>
        [Input("mirroredResources")]
        public Input<Inputs.PacketMirroringMirroredResourcesGetArgs>? MirroredResources { get; set; }

        /// <summary>
        /// The name of the packet mirroring rule
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Specifies the mirrored VPC network. Only packets in this network will be mirrored. All mirrored VMs should
        /// have a NIC in the given network. All mirrored subnetworks should belong to the given network.
        /// </summary>
        [Input("network")]
        public Input<Inputs.PacketMirroringNetworkGetArgs>? Network { get; set; }

        /// <summary>
        /// Since only one rule can be active at a time, priority is used to break ties in the case of two rules that
        /// apply to the same instances.
        /// </summary>
        [Input("priority")]
        public Input<int>? Priority { get; set; }

        [Input("project")]
        public Input<string>? Project { get; set; }

        /// <summary>
        /// The Region in which the created address should reside. If it is not provided, the provider region is used.
        /// </summary>
        [Input("region")]
        public Input<string>? Region { get; set; }

        public PacketMirroringState()
        {
        }
    }

    namespace Inputs
    {

    public sealed class PacketMirroringCollectorIlbArgs : Pulumi.ResourceArgs
    {
        [Input("url", required: true)]
        public Input<string> Url { get; set; } = null!;

        public PacketMirroringCollectorIlbArgs()
        {
        }
    }

    public sealed class PacketMirroringCollectorIlbGetArgs : Pulumi.ResourceArgs
    {
        [Input("url", required: true)]
        public Input<string> Url { get; set; } = null!;

        public PacketMirroringCollectorIlbGetArgs()
        {
        }
    }

    public sealed class PacketMirroringFilterArgs : Pulumi.ResourceArgs
    {
        [Input("cidrRanges")]
        private InputList<string>? _cidrRanges;
        public InputList<string> CidrRanges
        {
            get => _cidrRanges ?? (_cidrRanges = new InputList<string>());
            set => _cidrRanges = value;
        }

        [Input("ipProtocols")]
        private InputList<string>? _ipProtocols;
        public InputList<string> IpProtocols
        {
            get => _ipProtocols ?? (_ipProtocols = new InputList<string>());
            set => _ipProtocols = value;
        }

        public PacketMirroringFilterArgs()
        {
        }
    }

    public sealed class PacketMirroringFilterGetArgs : Pulumi.ResourceArgs
    {
        [Input("cidrRanges")]
        private InputList<string>? _cidrRanges;
        public InputList<string> CidrRanges
        {
            get => _cidrRanges ?? (_cidrRanges = new InputList<string>());
            set => _cidrRanges = value;
        }

        [Input("ipProtocols")]
        private InputList<string>? _ipProtocols;
        public InputList<string> IpProtocols
        {
            get => _ipProtocols ?? (_ipProtocols = new InputList<string>());
            set => _ipProtocols = value;
        }

        public PacketMirroringFilterGetArgs()
        {
        }
    }

    public sealed class PacketMirroringMirroredResourcesArgs : Pulumi.ResourceArgs
    {
        [Input("instances")]
        private InputList<PacketMirroringMirroredResourcesInstancesArgs>? _instances;
        public InputList<PacketMirroringMirroredResourcesInstancesArgs> Instances
        {
            get => _instances ?? (_instances = new InputList<PacketMirroringMirroredResourcesInstancesArgs>());
            set => _instances = value;
        }

        [Input("subnetworks")]
        private InputList<PacketMirroringMirroredResourcesSubnetworksArgs>? _subnetworks;
        public InputList<PacketMirroringMirroredResourcesSubnetworksArgs> Subnetworks
        {
            get => _subnetworks ?? (_subnetworks = new InputList<PacketMirroringMirroredResourcesSubnetworksArgs>());
            set => _subnetworks = value;
        }

        [Input("tags")]
        private InputList<string>? _tags;
        public InputList<string> Tags
        {
            get => _tags ?? (_tags = new InputList<string>());
            set => _tags = value;
        }

        public PacketMirroringMirroredResourcesArgs()
        {
        }
    }

    public sealed class PacketMirroringMirroredResourcesGetArgs : Pulumi.ResourceArgs
    {
        [Input("instances")]
        private InputList<PacketMirroringMirroredResourcesInstancesGetArgs>? _instances;
        public InputList<PacketMirroringMirroredResourcesInstancesGetArgs> Instances
        {
            get => _instances ?? (_instances = new InputList<PacketMirroringMirroredResourcesInstancesGetArgs>());
            set => _instances = value;
        }

        [Input("subnetworks")]
        private InputList<PacketMirroringMirroredResourcesSubnetworksGetArgs>? _subnetworks;
        public InputList<PacketMirroringMirroredResourcesSubnetworksGetArgs> Subnetworks
        {
            get => _subnetworks ?? (_subnetworks = new InputList<PacketMirroringMirroredResourcesSubnetworksGetArgs>());
            set => _subnetworks = value;
        }

        [Input("tags")]
        private InputList<string>? _tags;
        public InputList<string> Tags
        {
            get => _tags ?? (_tags = new InputList<string>());
            set => _tags = value;
        }

        public PacketMirroringMirroredResourcesGetArgs()
        {
        }
    }

    public sealed class PacketMirroringMirroredResourcesInstancesArgs : Pulumi.ResourceArgs
    {
        [Input("url", required: true)]
        public Input<string> Url { get; set; } = null!;

        public PacketMirroringMirroredResourcesInstancesArgs()
        {
        }
    }

    public sealed class PacketMirroringMirroredResourcesInstancesGetArgs : Pulumi.ResourceArgs
    {
        [Input("url", required: true)]
        public Input<string> Url { get; set; } = null!;

        public PacketMirroringMirroredResourcesInstancesGetArgs()
        {
        }
    }

    public sealed class PacketMirroringMirroredResourcesSubnetworksArgs : Pulumi.ResourceArgs
    {
        [Input("url", required: true)]
        public Input<string> Url { get; set; } = null!;

        public PacketMirroringMirroredResourcesSubnetworksArgs()
        {
        }
    }

    public sealed class PacketMirroringMirroredResourcesSubnetworksGetArgs : Pulumi.ResourceArgs
    {
        [Input("url", required: true)]
        public Input<string> Url { get; set; } = null!;

        public PacketMirroringMirroredResourcesSubnetworksGetArgs()
        {
        }
    }

    public sealed class PacketMirroringNetworkArgs : Pulumi.ResourceArgs
    {
        [Input("url", required: true)]
        public Input<string> Url { get; set; } = null!;

        public PacketMirroringNetworkArgs()
        {
        }
    }

    public sealed class PacketMirroringNetworkGetArgs : Pulumi.ResourceArgs
    {
        [Input("url", required: true)]
        public Input<string> Url { get; set; } = null!;

        public PacketMirroringNetworkGetArgs()
        {
        }
    }
    }

    namespace Outputs
    {

    [OutputType]
    public sealed class PacketMirroringCollectorIlb
    {
        public readonly string Url;

        [OutputConstructor]
        private PacketMirroringCollectorIlb(string url)
        {
            Url = url;
        }
    }

    [OutputType]
    public sealed class PacketMirroringFilter
    {
        public readonly ImmutableArray<string> CidrRanges;
        public readonly ImmutableArray<string> IpProtocols;

        [OutputConstructor]
        private PacketMirroringFilter(
            ImmutableArray<string> cidrRanges,
            ImmutableArray<string> ipProtocols)
        {
            CidrRanges = cidrRanges;
            IpProtocols = ipProtocols;
        }
    }

    [OutputType]
    public sealed class PacketMirroringMirroredResources
    {
        public readonly ImmutableArray<PacketMirroringMirroredResourcesInstances> Instances;
        public readonly ImmutableArray<PacketMirroringMirroredResourcesSubnetworks> Subnetworks;
        public readonly ImmutableArray<string> Tags;

        [OutputConstructor]
        private PacketMirroringMirroredResources(
            ImmutableArray<PacketMirroringMirroredResourcesInstances> instances,
            ImmutableArray<PacketMirroringMirroredResourcesSubnetworks> subnetworks,
            ImmutableArray<string> tags)
        {
            Instances = instances;
            Subnetworks = subnetworks;
            Tags = tags;
        }
    }

    [OutputType]
    public sealed class PacketMirroringMirroredResourcesInstances
    {
        public readonly string Url;

        [OutputConstructor]
        private PacketMirroringMirroredResourcesInstances(string url)
        {
            Url = url;
        }
    }

    [OutputType]
    public sealed class PacketMirroringMirroredResourcesSubnetworks
    {
        public readonly string Url;

        [OutputConstructor]
        private PacketMirroringMirroredResourcesSubnetworks(string url)
        {
            Url = url;
        }
    }

    [OutputType]
    public sealed class PacketMirroringNetwork
    {
        public readonly string Url;

        [OutputConstructor]
        private PacketMirroringNetwork(string url)
        {
            Url = url;
        }
    }
    }
}
