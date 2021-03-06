// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.GameServices
{
    /// <summary>
    /// &gt; This content is derived from https://github.com/terraform-providers/terraform-provider-google/blob/master/website/docs/r/game_services_game_server_cluster.html.markdown.
    /// </summary>
    public partial class GameServerCluster : Pulumi.CustomResource
    {
        /// <summary>
        /// Required. The resource name of the game server cluster
        /// </summary>
        [Output("clusterId")]
        public Output<string> ClusterId { get; private set; } = null!;

        /// <summary>
        /// Game server cluster connection information. This information is used to manage game server clusters.
        /// </summary>
        [Output("connectionInfo")]
        public Output<Outputs.GameServerClusterConnectionInfo> ConnectionInfo { get; private set; } = null!;

        /// <summary>
        /// Human readable description of the cluster.
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// The labels associated with this game server cluster. Each label is a key-value pair.
        /// </summary>
        [Output("labels")]
        public Output<ImmutableDictionary<string, string>?> Labels { get; private set; } = null!;

        /// <summary>
        /// Location of the Cluster.
        /// </summary>
        [Output("location")]
        public Output<string?> Location { get; private set; } = null!;

        /// <summary>
        /// The resource id of the game server cluster, eg:
        /// 'projects/{project_id}/locations/{location}/realms/{realm_id}/gameServerClusters/{cluster_id}'. For example,
        /// 'projects/my-project/locations/{location}/realms/zanzibar/gameServerClusters/my-onprem-cluster'.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        [Output("project")]
        public Output<string> Project { get; private set; } = null!;

        /// <summary>
        /// The realm id of the game server realm.
        /// </summary>
        [Output("realmId")]
        public Output<string> RealmId { get; private set; } = null!;


        /// <summary>
        /// Create a GameServerCluster resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public GameServerCluster(string name, GameServerClusterArgs args, CustomResourceOptions? options = null)
            : base("gcp:gameservices/gameServerCluster:GameServerCluster", name, args ?? ResourceArgs.Empty, MakeResourceOptions(options, ""))
        {
        }

        private GameServerCluster(string name, Input<string> id, GameServerClusterState? state = null, CustomResourceOptions? options = null)
            : base("gcp:gameservices/gameServerCluster:GameServerCluster", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing GameServerCluster resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static GameServerCluster Get(string name, Input<string> id, GameServerClusterState? state = null, CustomResourceOptions? options = null)
        {
            return new GameServerCluster(name, id, state, options);
        }
    }

    public sealed class GameServerClusterArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Required. The resource name of the game server cluster
        /// </summary>
        [Input("clusterId", required: true)]
        public Input<string> ClusterId { get; set; } = null!;

        /// <summary>
        /// Game server cluster connection information. This information is used to manage game server clusters.
        /// </summary>
        [Input("connectionInfo", required: true)]
        public Input<Inputs.GameServerClusterConnectionInfoArgs> ConnectionInfo { get; set; } = null!;

        /// <summary>
        /// Human readable description of the cluster.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        [Input("labels")]
        private InputMap<string>? _labels;

        /// <summary>
        /// The labels associated with this game server cluster. Each label is a key-value pair.
        /// </summary>
        public InputMap<string> Labels
        {
            get => _labels ?? (_labels = new InputMap<string>());
            set => _labels = value;
        }

        /// <summary>
        /// Location of the Cluster.
        /// </summary>
        [Input("location")]
        public Input<string>? Location { get; set; }

        [Input("project")]
        public Input<string>? Project { get; set; }

        /// <summary>
        /// The realm id of the game server realm.
        /// </summary>
        [Input("realmId", required: true)]
        public Input<string> RealmId { get; set; } = null!;

        public GameServerClusterArgs()
        {
        }
    }

    public sealed class GameServerClusterState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Required. The resource name of the game server cluster
        /// </summary>
        [Input("clusterId")]
        public Input<string>? ClusterId { get; set; }

        /// <summary>
        /// Game server cluster connection information. This information is used to manage game server clusters.
        /// </summary>
        [Input("connectionInfo")]
        public Input<Inputs.GameServerClusterConnectionInfoGetArgs>? ConnectionInfo { get; set; }

        /// <summary>
        /// Human readable description of the cluster.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        [Input("labels")]
        private InputMap<string>? _labels;

        /// <summary>
        /// The labels associated with this game server cluster. Each label is a key-value pair.
        /// </summary>
        public InputMap<string> Labels
        {
            get => _labels ?? (_labels = new InputMap<string>());
            set => _labels = value;
        }

        /// <summary>
        /// Location of the Cluster.
        /// </summary>
        [Input("location")]
        public Input<string>? Location { get; set; }

        /// <summary>
        /// The resource id of the game server cluster, eg:
        /// 'projects/{project_id}/locations/{location}/realms/{realm_id}/gameServerClusters/{cluster_id}'. For example,
        /// 'projects/my-project/locations/{location}/realms/zanzibar/gameServerClusters/my-onprem-cluster'.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("project")]
        public Input<string>? Project { get; set; }

        /// <summary>
        /// The realm id of the game server realm.
        /// </summary>
        [Input("realmId")]
        public Input<string>? RealmId { get; set; }

        public GameServerClusterState()
        {
        }
    }

    namespace Inputs
    {

    public sealed class GameServerClusterConnectionInfoArgs : Pulumi.ResourceArgs
    {
        [Input("gkeClusterReference", required: true)]
        public Input<GameServerClusterConnectionInfoGkeClusterReferenceArgs> GkeClusterReference { get; set; } = null!;

        [Input("namespace", required: true)]
        public Input<string> Namespace { get; set; } = null!;

        public GameServerClusterConnectionInfoArgs()
        {
        }
    }

    public sealed class GameServerClusterConnectionInfoGetArgs : Pulumi.ResourceArgs
    {
        [Input("gkeClusterReference", required: true)]
        public Input<GameServerClusterConnectionInfoGkeClusterReferenceGetArgs> GkeClusterReference { get; set; } = null!;

        [Input("namespace", required: true)]
        public Input<string> Namespace { get; set; } = null!;

        public GameServerClusterConnectionInfoGetArgs()
        {
        }
    }

    public sealed class GameServerClusterConnectionInfoGkeClusterReferenceArgs : Pulumi.ResourceArgs
    {
        [Input("cluster", required: true)]
        public Input<string> Cluster { get; set; } = null!;

        public GameServerClusterConnectionInfoGkeClusterReferenceArgs()
        {
        }
    }

    public sealed class GameServerClusterConnectionInfoGkeClusterReferenceGetArgs : Pulumi.ResourceArgs
    {
        [Input("cluster", required: true)]
        public Input<string> Cluster { get; set; } = null!;

        public GameServerClusterConnectionInfoGkeClusterReferenceGetArgs()
        {
        }
    }
    }

    namespace Outputs
    {

    [OutputType]
    public sealed class GameServerClusterConnectionInfo
    {
        public readonly GameServerClusterConnectionInfoGkeClusterReference GkeClusterReference;
        public readonly string Namespace;

        [OutputConstructor]
        private GameServerClusterConnectionInfo(
            GameServerClusterConnectionInfoGkeClusterReference gkeClusterReference,
            string @namespace)
        {
            GkeClusterReference = gkeClusterReference;
            Namespace = @namespace;
        }
    }

    [OutputType]
    public sealed class GameServerClusterConnectionInfoGkeClusterReference
    {
        public readonly string Cluster;

        [OutputConstructor]
        private GameServerClusterConnectionInfoGkeClusterReference(string cluster)
        {
            Cluster = cluster;
        }
    }
    }
}
