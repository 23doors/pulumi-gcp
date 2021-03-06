// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.Monitoring
{
    /// <summary>
    /// &gt; This content is derived from https://github.com/terraform-providers/terraform-provider-google/blob/master/website/docs/r/monitoring_alert_policy.html.markdown.
    /// </summary>
    public partial class AlertPolicy : Pulumi.CustomResource
    {
        /// <summary>
        /// How to combine the results of multiple conditions to determine if an incident should be opened.
        /// </summary>
        [Output("combiner")]
        public Output<string> Combiner { get; private set; } = null!;

        /// <summary>
        /// A list of conditions for the policy. The conditions are combined by AND or OR according to the combiner
        /// field. If the combined conditions evaluate to true, then an incident is created. A policy can have from one
        /// to six conditions.
        /// </summary>
        [Output("conditions")]
        public Output<ImmutableArray<Outputs.AlertPolicyConditions>> Conditions { get; private set; } = null!;

        /// <summary>
        /// A read-only record of the creation of the alerting policy. If provided in a call to create or update, this
        /// field will be ignored.
        /// </summary>
        [Output("creationRecord")]
        public Output<Outputs.AlertPolicyCreationRecord> CreationRecord { get; private set; } = null!;

        /// <summary>
        /// A short name or phrase used to identify the policy in dashboards, notifications, and incidents. To avoid
        /// confusion, don't use the same display name for multiple policies in the same project. The name is limited to
        /// 512 Unicode characters.
        /// </summary>
        [Output("displayName")]
        public Output<string> DisplayName { get; private set; } = null!;

        /// <summary>
        /// A short name or phrase used to identify the policy in dashboards, notifications, and incidents. To avoid
        /// confusion, don't use the same display name for multiple policies in the same project. The name is limited to
        /// 512 Unicode characters.
        /// </summary>
        [Output("documentation")]
        public Output<Outputs.AlertPolicyDocumentation?> Documentation { get; private set; } = null!;

        /// <summary>
        /// Whether or not the policy is enabled. The default is true.
        /// </summary>
        [Output("enabled")]
        public Output<bool?> Enabled { get; private set; } = null!;

        /// <summary>
        /// The unique resource name for this policy. Its syntax is:
        /// projects/[PROJECT_ID]/alertPolicies/[ALERT_POLICY_ID]
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Identifies the notification channels to which notifications should be sent when incidents are opened or
        /// closed or when new violations occur on an already opened incident. Each element of this array corresponds to
        /// the name field in each of the NotificationChannel objects that are returned from the
        /// notificationChannels.list method. The syntax of the entries in this field is
        /// 'projects/[PROJECT_ID]/notificationChannels/[CHANNEL_ID]'
        /// </summary>
        [Output("notificationChannels")]
        public Output<ImmutableArray<string>> NotificationChannels { get; private set; } = null!;

        [Output("project")]
        public Output<string> Project { get; private set; } = null!;

        /// <summary>
        /// This field is intended to be used for organizing and identifying the AlertPolicy objects.The field can
        /// contain up to 64 entries. Each key and value is limited to 63 Unicode characters or 128 bytes, whichever is
        /// smaller. Labels and values can contain only lowercase letters, numerals, underscores, and dashes. Keys must
        /// begin with a letter.
        /// </summary>
        [Output("userLabels")]
        public Output<ImmutableDictionary<string, string>?> UserLabels { get; private set; } = null!;


        /// <summary>
        /// Create a AlertPolicy resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public AlertPolicy(string name, AlertPolicyArgs args, CustomResourceOptions? options = null)
            : base("gcp:monitoring/alertPolicy:AlertPolicy", name, args ?? ResourceArgs.Empty, MakeResourceOptions(options, ""))
        {
        }

        private AlertPolicy(string name, Input<string> id, AlertPolicyState? state = null, CustomResourceOptions? options = null)
            : base("gcp:monitoring/alertPolicy:AlertPolicy", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing AlertPolicy resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static AlertPolicy Get(string name, Input<string> id, AlertPolicyState? state = null, CustomResourceOptions? options = null)
        {
            return new AlertPolicy(name, id, state, options);
        }
    }

    public sealed class AlertPolicyArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// How to combine the results of multiple conditions to determine if an incident should be opened.
        /// </summary>
        [Input("combiner", required: true)]
        public Input<string> Combiner { get; set; } = null!;

        [Input("conditions", required: true)]
        private InputList<Inputs.AlertPolicyConditionsArgs>? _conditions;

        /// <summary>
        /// A list of conditions for the policy. The conditions are combined by AND or OR according to the combiner
        /// field. If the combined conditions evaluate to true, then an incident is created. A policy can have from one
        /// to six conditions.
        /// </summary>
        public InputList<Inputs.AlertPolicyConditionsArgs> Conditions
        {
            get => _conditions ?? (_conditions = new InputList<Inputs.AlertPolicyConditionsArgs>());
            set => _conditions = value;
        }

        /// <summary>
        /// A short name or phrase used to identify the policy in dashboards, notifications, and incidents. To avoid
        /// confusion, don't use the same display name for multiple policies in the same project. The name is limited to
        /// 512 Unicode characters.
        /// </summary>
        [Input("displayName", required: true)]
        public Input<string> DisplayName { get; set; } = null!;

        /// <summary>
        /// A short name or phrase used to identify the policy in dashboards, notifications, and incidents. To avoid
        /// confusion, don't use the same display name for multiple policies in the same project. The name is limited to
        /// 512 Unicode characters.
        /// </summary>
        [Input("documentation")]
        public Input<Inputs.AlertPolicyDocumentationArgs>? Documentation { get; set; }

        /// <summary>
        /// Whether or not the policy is enabled. The default is true.
        /// </summary>
        [Input("enabled")]
        public Input<bool>? Enabled { get; set; }

        [Input("notificationChannels")]
        private InputList<string>? _notificationChannels;

        /// <summary>
        /// Identifies the notification channels to which notifications should be sent when incidents are opened or
        /// closed or when new violations occur on an already opened incident. Each element of this array corresponds to
        /// the name field in each of the NotificationChannel objects that are returned from the
        /// notificationChannels.list method. The syntax of the entries in this field is
        /// 'projects/[PROJECT_ID]/notificationChannels/[CHANNEL_ID]'
        /// </summary>
        public InputList<string> NotificationChannels
        {
            get => _notificationChannels ?? (_notificationChannels = new InputList<string>());
            set => _notificationChannels = value;
        }

        [Input("project")]
        public Input<string>? Project { get; set; }

        [Input("userLabels")]
        private InputMap<string>? _userLabels;

        /// <summary>
        /// This field is intended to be used for organizing and identifying the AlertPolicy objects.The field can
        /// contain up to 64 entries. Each key and value is limited to 63 Unicode characters or 128 bytes, whichever is
        /// smaller. Labels and values can contain only lowercase letters, numerals, underscores, and dashes. Keys must
        /// begin with a letter.
        /// </summary>
        public InputMap<string> UserLabels
        {
            get => _userLabels ?? (_userLabels = new InputMap<string>());
            set => _userLabels = value;
        }

        public AlertPolicyArgs()
        {
        }
    }

    public sealed class AlertPolicyState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// How to combine the results of multiple conditions to determine if an incident should be opened.
        /// </summary>
        [Input("combiner")]
        public Input<string>? Combiner { get; set; }

        [Input("conditions")]
        private InputList<Inputs.AlertPolicyConditionsGetArgs>? _conditions;

        /// <summary>
        /// A list of conditions for the policy. The conditions are combined by AND or OR according to the combiner
        /// field. If the combined conditions evaluate to true, then an incident is created. A policy can have from one
        /// to six conditions.
        /// </summary>
        public InputList<Inputs.AlertPolicyConditionsGetArgs> Conditions
        {
            get => _conditions ?? (_conditions = new InputList<Inputs.AlertPolicyConditionsGetArgs>());
            set => _conditions = value;
        }

        /// <summary>
        /// A read-only record of the creation of the alerting policy. If provided in a call to create or update, this
        /// field will be ignored.
        /// </summary>
        [Input("creationRecord")]
        public Input<Inputs.AlertPolicyCreationRecordGetArgs>? CreationRecord { get; set; }

        /// <summary>
        /// A short name or phrase used to identify the policy in dashboards, notifications, and incidents. To avoid
        /// confusion, don't use the same display name for multiple policies in the same project. The name is limited to
        /// 512 Unicode characters.
        /// </summary>
        [Input("displayName")]
        public Input<string>? DisplayName { get; set; }

        /// <summary>
        /// A short name or phrase used to identify the policy in dashboards, notifications, and incidents. To avoid
        /// confusion, don't use the same display name for multiple policies in the same project. The name is limited to
        /// 512 Unicode characters.
        /// </summary>
        [Input("documentation")]
        public Input<Inputs.AlertPolicyDocumentationGetArgs>? Documentation { get; set; }

        /// <summary>
        /// Whether or not the policy is enabled. The default is true.
        /// </summary>
        [Input("enabled")]
        public Input<bool>? Enabled { get; set; }

        /// <summary>
        /// The unique resource name for this policy. Its syntax is:
        /// projects/[PROJECT_ID]/alertPolicies/[ALERT_POLICY_ID]
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("notificationChannels")]
        private InputList<string>? _notificationChannels;

        /// <summary>
        /// Identifies the notification channels to which notifications should be sent when incidents are opened or
        /// closed or when new violations occur on an already opened incident. Each element of this array corresponds to
        /// the name field in each of the NotificationChannel objects that are returned from the
        /// notificationChannels.list method. The syntax of the entries in this field is
        /// 'projects/[PROJECT_ID]/notificationChannels/[CHANNEL_ID]'
        /// </summary>
        public InputList<string> NotificationChannels
        {
            get => _notificationChannels ?? (_notificationChannels = new InputList<string>());
            set => _notificationChannels = value;
        }

        [Input("project")]
        public Input<string>? Project { get; set; }

        [Input("userLabels")]
        private InputMap<string>? _userLabels;

        /// <summary>
        /// This field is intended to be used for organizing and identifying the AlertPolicy objects.The field can
        /// contain up to 64 entries. Each key and value is limited to 63 Unicode characters or 128 bytes, whichever is
        /// smaller. Labels and values can contain only lowercase letters, numerals, underscores, and dashes. Keys must
        /// begin with a letter.
        /// </summary>
        public InputMap<string> UserLabels
        {
            get => _userLabels ?? (_userLabels = new InputMap<string>());
            set => _userLabels = value;
        }

        public AlertPolicyState()
        {
        }
    }

    namespace Inputs
    {

    public sealed class AlertPolicyConditionsArgs : Pulumi.ResourceArgs
    {
        [Input("conditionAbsent")]
        public Input<AlertPolicyConditionsConditionAbsentArgs>? ConditionAbsent { get; set; }

        [Input("conditionThreshold")]
        public Input<AlertPolicyConditionsConditionThresholdArgs>? ConditionThreshold { get; set; }

        [Input("displayName", required: true)]
        public Input<string> DisplayName { get; set; } = null!;

        [Input("name")]
        public Input<string>? Name { get; set; }

        public AlertPolicyConditionsArgs()
        {
        }
    }

    public sealed class AlertPolicyConditionsConditionAbsentAggregationsArgs : Pulumi.ResourceArgs
    {
        [Input("alignmentPeriod")]
        public Input<string>? AlignmentPeriod { get; set; }

        [Input("crossSeriesReducer")]
        public Input<string>? CrossSeriesReducer { get; set; }

        [Input("groupByFields")]
        private InputList<string>? _groupByFields;
        public InputList<string> GroupByFields
        {
            get => _groupByFields ?? (_groupByFields = new InputList<string>());
            set => _groupByFields = value;
        }

        [Input("perSeriesAligner")]
        public Input<string>? PerSeriesAligner { get; set; }

        public AlertPolicyConditionsConditionAbsentAggregationsArgs()
        {
        }
    }

    public sealed class AlertPolicyConditionsConditionAbsentAggregationsGetArgs : Pulumi.ResourceArgs
    {
        [Input("alignmentPeriod")]
        public Input<string>? AlignmentPeriod { get; set; }

        [Input("crossSeriesReducer")]
        public Input<string>? CrossSeriesReducer { get; set; }

        [Input("groupByFields")]
        private InputList<string>? _groupByFields;
        public InputList<string> GroupByFields
        {
            get => _groupByFields ?? (_groupByFields = new InputList<string>());
            set => _groupByFields = value;
        }

        [Input("perSeriesAligner")]
        public Input<string>? PerSeriesAligner { get; set; }

        public AlertPolicyConditionsConditionAbsentAggregationsGetArgs()
        {
        }
    }

    public sealed class AlertPolicyConditionsConditionAbsentArgs : Pulumi.ResourceArgs
    {
        [Input("aggregations")]
        private InputList<AlertPolicyConditionsConditionAbsentAggregationsArgs>? _aggregations;
        public InputList<AlertPolicyConditionsConditionAbsentAggregationsArgs> Aggregations
        {
            get => _aggregations ?? (_aggregations = new InputList<AlertPolicyConditionsConditionAbsentAggregationsArgs>());
            set => _aggregations = value;
        }

        [Input("duration", required: true)]
        public Input<string> Duration { get; set; } = null!;

        [Input("filter")]
        public Input<string>? Filter { get; set; }

        [Input("trigger")]
        public Input<AlertPolicyConditionsConditionAbsentTriggerArgs>? Trigger { get; set; }

        public AlertPolicyConditionsConditionAbsentArgs()
        {
        }
    }

    public sealed class AlertPolicyConditionsConditionAbsentGetArgs : Pulumi.ResourceArgs
    {
        [Input("aggregations")]
        private InputList<AlertPolicyConditionsConditionAbsentAggregationsGetArgs>? _aggregations;
        public InputList<AlertPolicyConditionsConditionAbsentAggregationsGetArgs> Aggregations
        {
            get => _aggregations ?? (_aggregations = new InputList<AlertPolicyConditionsConditionAbsentAggregationsGetArgs>());
            set => _aggregations = value;
        }

        [Input("duration", required: true)]
        public Input<string> Duration { get; set; } = null!;

        [Input("filter")]
        public Input<string>? Filter { get; set; }

        [Input("trigger")]
        public Input<AlertPolicyConditionsConditionAbsentTriggerGetArgs>? Trigger { get; set; }

        public AlertPolicyConditionsConditionAbsentGetArgs()
        {
        }
    }

    public sealed class AlertPolicyConditionsConditionAbsentTriggerArgs : Pulumi.ResourceArgs
    {
        [Input("count")]
        public Input<int>? Count { get; set; }

        [Input("percent")]
        public Input<double>? Percent { get; set; }

        public AlertPolicyConditionsConditionAbsentTriggerArgs()
        {
        }
    }

    public sealed class AlertPolicyConditionsConditionAbsentTriggerGetArgs : Pulumi.ResourceArgs
    {
        [Input("count")]
        public Input<int>? Count { get; set; }

        [Input("percent")]
        public Input<double>? Percent { get; set; }

        public AlertPolicyConditionsConditionAbsentTriggerGetArgs()
        {
        }
    }

    public sealed class AlertPolicyConditionsConditionThresholdAggregationsArgs : Pulumi.ResourceArgs
    {
        [Input("alignmentPeriod")]
        public Input<string>? AlignmentPeriod { get; set; }

        [Input("crossSeriesReducer")]
        public Input<string>? CrossSeriesReducer { get; set; }

        [Input("groupByFields")]
        private InputList<string>? _groupByFields;
        public InputList<string> GroupByFields
        {
            get => _groupByFields ?? (_groupByFields = new InputList<string>());
            set => _groupByFields = value;
        }

        [Input("perSeriesAligner")]
        public Input<string>? PerSeriesAligner { get; set; }

        public AlertPolicyConditionsConditionThresholdAggregationsArgs()
        {
        }
    }

    public sealed class AlertPolicyConditionsConditionThresholdAggregationsGetArgs : Pulumi.ResourceArgs
    {
        [Input("alignmentPeriod")]
        public Input<string>? AlignmentPeriod { get; set; }

        [Input("crossSeriesReducer")]
        public Input<string>? CrossSeriesReducer { get; set; }

        [Input("groupByFields")]
        private InputList<string>? _groupByFields;
        public InputList<string> GroupByFields
        {
            get => _groupByFields ?? (_groupByFields = new InputList<string>());
            set => _groupByFields = value;
        }

        [Input("perSeriesAligner")]
        public Input<string>? PerSeriesAligner { get; set; }

        public AlertPolicyConditionsConditionThresholdAggregationsGetArgs()
        {
        }
    }

    public sealed class AlertPolicyConditionsConditionThresholdArgs : Pulumi.ResourceArgs
    {
        [Input("aggregations")]
        private InputList<AlertPolicyConditionsConditionThresholdAggregationsArgs>? _aggregations;
        public InputList<AlertPolicyConditionsConditionThresholdAggregationsArgs> Aggregations
        {
            get => _aggregations ?? (_aggregations = new InputList<AlertPolicyConditionsConditionThresholdAggregationsArgs>());
            set => _aggregations = value;
        }

        [Input("comparison", required: true)]
        public Input<string> Comparison { get; set; } = null!;

        [Input("denominatorAggregations")]
        private InputList<AlertPolicyConditionsConditionThresholdDenominatorAggregationsArgs>? _denominatorAggregations;
        public InputList<AlertPolicyConditionsConditionThresholdDenominatorAggregationsArgs> DenominatorAggregations
        {
            get => _denominatorAggregations ?? (_denominatorAggregations = new InputList<AlertPolicyConditionsConditionThresholdDenominatorAggregationsArgs>());
            set => _denominatorAggregations = value;
        }

        [Input("denominatorFilter")]
        public Input<string>? DenominatorFilter { get; set; }

        [Input("duration", required: true)]
        public Input<string> Duration { get; set; } = null!;

        [Input("filter")]
        public Input<string>? Filter { get; set; }

        [Input("thresholdValue")]
        public Input<double>? ThresholdValue { get; set; }

        [Input("trigger")]
        public Input<AlertPolicyConditionsConditionThresholdTriggerArgs>? Trigger { get; set; }

        public AlertPolicyConditionsConditionThresholdArgs()
        {
        }
    }

    public sealed class AlertPolicyConditionsConditionThresholdDenominatorAggregationsArgs : Pulumi.ResourceArgs
    {
        [Input("alignmentPeriod")]
        public Input<string>? AlignmentPeriod { get; set; }

        [Input("crossSeriesReducer")]
        public Input<string>? CrossSeriesReducer { get; set; }

        [Input("groupByFields")]
        private InputList<string>? _groupByFields;
        public InputList<string> GroupByFields
        {
            get => _groupByFields ?? (_groupByFields = new InputList<string>());
            set => _groupByFields = value;
        }

        [Input("perSeriesAligner")]
        public Input<string>? PerSeriesAligner { get; set; }

        public AlertPolicyConditionsConditionThresholdDenominatorAggregationsArgs()
        {
        }
    }

    public sealed class AlertPolicyConditionsConditionThresholdDenominatorAggregationsGetArgs : Pulumi.ResourceArgs
    {
        [Input("alignmentPeriod")]
        public Input<string>? AlignmentPeriod { get; set; }

        [Input("crossSeriesReducer")]
        public Input<string>? CrossSeriesReducer { get; set; }

        [Input("groupByFields")]
        private InputList<string>? _groupByFields;
        public InputList<string> GroupByFields
        {
            get => _groupByFields ?? (_groupByFields = new InputList<string>());
            set => _groupByFields = value;
        }

        [Input("perSeriesAligner")]
        public Input<string>? PerSeriesAligner { get; set; }

        public AlertPolicyConditionsConditionThresholdDenominatorAggregationsGetArgs()
        {
        }
    }

    public sealed class AlertPolicyConditionsConditionThresholdGetArgs : Pulumi.ResourceArgs
    {
        [Input("aggregations")]
        private InputList<AlertPolicyConditionsConditionThresholdAggregationsGetArgs>? _aggregations;
        public InputList<AlertPolicyConditionsConditionThresholdAggregationsGetArgs> Aggregations
        {
            get => _aggregations ?? (_aggregations = new InputList<AlertPolicyConditionsConditionThresholdAggregationsGetArgs>());
            set => _aggregations = value;
        }

        [Input("comparison", required: true)]
        public Input<string> Comparison { get; set; } = null!;

        [Input("denominatorAggregations")]
        private InputList<AlertPolicyConditionsConditionThresholdDenominatorAggregationsGetArgs>? _denominatorAggregations;
        public InputList<AlertPolicyConditionsConditionThresholdDenominatorAggregationsGetArgs> DenominatorAggregations
        {
            get => _denominatorAggregations ?? (_denominatorAggregations = new InputList<AlertPolicyConditionsConditionThresholdDenominatorAggregationsGetArgs>());
            set => _denominatorAggregations = value;
        }

        [Input("denominatorFilter")]
        public Input<string>? DenominatorFilter { get; set; }

        [Input("duration", required: true)]
        public Input<string> Duration { get; set; } = null!;

        [Input("filter")]
        public Input<string>? Filter { get; set; }

        [Input("thresholdValue")]
        public Input<double>? ThresholdValue { get; set; }

        [Input("trigger")]
        public Input<AlertPolicyConditionsConditionThresholdTriggerGetArgs>? Trigger { get; set; }

        public AlertPolicyConditionsConditionThresholdGetArgs()
        {
        }
    }

    public sealed class AlertPolicyConditionsConditionThresholdTriggerArgs : Pulumi.ResourceArgs
    {
        [Input("count")]
        public Input<int>? Count { get; set; }

        [Input("percent")]
        public Input<double>? Percent { get; set; }

        public AlertPolicyConditionsConditionThresholdTriggerArgs()
        {
        }
    }

    public sealed class AlertPolicyConditionsConditionThresholdTriggerGetArgs : Pulumi.ResourceArgs
    {
        [Input("count")]
        public Input<int>? Count { get; set; }

        [Input("percent")]
        public Input<double>? Percent { get; set; }

        public AlertPolicyConditionsConditionThresholdTriggerGetArgs()
        {
        }
    }

    public sealed class AlertPolicyConditionsGetArgs : Pulumi.ResourceArgs
    {
        [Input("conditionAbsent")]
        public Input<AlertPolicyConditionsConditionAbsentGetArgs>? ConditionAbsent { get; set; }

        [Input("conditionThreshold")]
        public Input<AlertPolicyConditionsConditionThresholdGetArgs>? ConditionThreshold { get; set; }

        [Input("displayName", required: true)]
        public Input<string> DisplayName { get; set; } = null!;

        [Input("name")]
        public Input<string>? Name { get; set; }

        public AlertPolicyConditionsGetArgs()
        {
        }
    }

    public sealed class AlertPolicyCreationRecordGetArgs : Pulumi.ResourceArgs
    {
        [Input("mutateTime")]
        public Input<string>? MutateTime { get; set; }

        [Input("mutatedBy")]
        public Input<string>? MutatedBy { get; set; }

        public AlertPolicyCreationRecordGetArgs()
        {
        }
    }

    public sealed class AlertPolicyDocumentationArgs : Pulumi.ResourceArgs
    {
        [Input("content")]
        public Input<string>? Content { get; set; }

        [Input("mimeType")]
        public Input<string>? MimeType { get; set; }

        public AlertPolicyDocumentationArgs()
        {
        }
    }

    public sealed class AlertPolicyDocumentationGetArgs : Pulumi.ResourceArgs
    {
        [Input("content")]
        public Input<string>? Content { get; set; }

        [Input("mimeType")]
        public Input<string>? MimeType { get; set; }

        public AlertPolicyDocumentationGetArgs()
        {
        }
    }
    }

    namespace Outputs
    {

    [OutputType]
    public sealed class AlertPolicyConditions
    {
        public readonly AlertPolicyConditionsConditionAbsent? ConditionAbsent;
        public readonly AlertPolicyConditionsConditionThreshold? ConditionThreshold;
        public readonly string DisplayName;
        public readonly string Name;

        [OutputConstructor]
        private AlertPolicyConditions(
            AlertPolicyConditionsConditionAbsent? conditionAbsent,
            AlertPolicyConditionsConditionThreshold? conditionThreshold,
            string displayName,
            string name)
        {
            ConditionAbsent = conditionAbsent;
            ConditionThreshold = conditionThreshold;
            DisplayName = displayName;
            Name = name;
        }
    }

    [OutputType]
    public sealed class AlertPolicyConditionsConditionAbsent
    {
        public readonly ImmutableArray<AlertPolicyConditionsConditionAbsentAggregations> Aggregations;
        public readonly string Duration;
        public readonly string? Filter;
        public readonly AlertPolicyConditionsConditionAbsentTrigger? Trigger;

        [OutputConstructor]
        private AlertPolicyConditionsConditionAbsent(
            ImmutableArray<AlertPolicyConditionsConditionAbsentAggregations> aggregations,
            string duration,
            string? filter,
            AlertPolicyConditionsConditionAbsentTrigger? trigger)
        {
            Aggregations = aggregations;
            Duration = duration;
            Filter = filter;
            Trigger = trigger;
        }
    }

    [OutputType]
    public sealed class AlertPolicyConditionsConditionAbsentAggregations
    {
        public readonly string? AlignmentPeriod;
        public readonly string? CrossSeriesReducer;
        public readonly ImmutableArray<string> GroupByFields;
        public readonly string? PerSeriesAligner;

        [OutputConstructor]
        private AlertPolicyConditionsConditionAbsentAggregations(
            string? alignmentPeriod,
            string? crossSeriesReducer,
            ImmutableArray<string> groupByFields,
            string? perSeriesAligner)
        {
            AlignmentPeriod = alignmentPeriod;
            CrossSeriesReducer = crossSeriesReducer;
            GroupByFields = groupByFields;
            PerSeriesAligner = perSeriesAligner;
        }
    }

    [OutputType]
    public sealed class AlertPolicyConditionsConditionAbsentTrigger
    {
        public readonly int? Count;
        public readonly double? Percent;

        [OutputConstructor]
        private AlertPolicyConditionsConditionAbsentTrigger(
            int? count,
            double? percent)
        {
            Count = count;
            Percent = percent;
        }
    }

    [OutputType]
    public sealed class AlertPolicyConditionsConditionThreshold
    {
        public readonly ImmutableArray<AlertPolicyConditionsConditionThresholdAggregations> Aggregations;
        public readonly string Comparison;
        public readonly ImmutableArray<AlertPolicyConditionsConditionThresholdDenominatorAggregations> DenominatorAggregations;
        public readonly string? DenominatorFilter;
        public readonly string Duration;
        public readonly string? Filter;
        public readonly double? ThresholdValue;
        public readonly AlertPolicyConditionsConditionThresholdTrigger? Trigger;

        [OutputConstructor]
        private AlertPolicyConditionsConditionThreshold(
            ImmutableArray<AlertPolicyConditionsConditionThresholdAggregations> aggregations,
            string comparison,
            ImmutableArray<AlertPolicyConditionsConditionThresholdDenominatorAggregations> denominatorAggregations,
            string? denominatorFilter,
            string duration,
            string? filter,
            double? thresholdValue,
            AlertPolicyConditionsConditionThresholdTrigger? trigger)
        {
            Aggregations = aggregations;
            Comparison = comparison;
            DenominatorAggregations = denominatorAggregations;
            DenominatorFilter = denominatorFilter;
            Duration = duration;
            Filter = filter;
            ThresholdValue = thresholdValue;
            Trigger = trigger;
        }
    }

    [OutputType]
    public sealed class AlertPolicyConditionsConditionThresholdAggregations
    {
        public readonly string? AlignmentPeriod;
        public readonly string? CrossSeriesReducer;
        public readonly ImmutableArray<string> GroupByFields;
        public readonly string? PerSeriesAligner;

        [OutputConstructor]
        private AlertPolicyConditionsConditionThresholdAggregations(
            string? alignmentPeriod,
            string? crossSeriesReducer,
            ImmutableArray<string> groupByFields,
            string? perSeriesAligner)
        {
            AlignmentPeriod = alignmentPeriod;
            CrossSeriesReducer = crossSeriesReducer;
            GroupByFields = groupByFields;
            PerSeriesAligner = perSeriesAligner;
        }
    }

    [OutputType]
    public sealed class AlertPolicyConditionsConditionThresholdDenominatorAggregations
    {
        public readonly string? AlignmentPeriod;
        public readonly string? CrossSeriesReducer;
        public readonly ImmutableArray<string> GroupByFields;
        public readonly string? PerSeriesAligner;

        [OutputConstructor]
        private AlertPolicyConditionsConditionThresholdDenominatorAggregations(
            string? alignmentPeriod,
            string? crossSeriesReducer,
            ImmutableArray<string> groupByFields,
            string? perSeriesAligner)
        {
            AlignmentPeriod = alignmentPeriod;
            CrossSeriesReducer = crossSeriesReducer;
            GroupByFields = groupByFields;
            PerSeriesAligner = perSeriesAligner;
        }
    }

    [OutputType]
    public sealed class AlertPolicyConditionsConditionThresholdTrigger
    {
        public readonly int? Count;
        public readonly double? Percent;

        [OutputConstructor]
        private AlertPolicyConditionsConditionThresholdTrigger(
            int? count,
            double? percent)
        {
            Count = count;
            Percent = percent;
        }
    }

    [OutputType]
    public sealed class AlertPolicyCreationRecord
    {
        public readonly string MutateTime;
        public readonly string MutatedBy;

        [OutputConstructor]
        private AlertPolicyCreationRecord(
            string mutateTime,
            string mutatedBy)
        {
            MutateTime = mutateTime;
            MutatedBy = mutatedBy;
        }
    }

    [OutputType]
    public sealed class AlertPolicyDocumentation
    {
        public readonly string? Content;
        public readonly string? MimeType;

        [OutputConstructor]
        private AlertPolicyDocumentation(
            string? content,
            string? mimeType)
        {
            Content = content;
            MimeType = mimeType;
        }
    }
    }
}
