// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package compute

import (
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// The Google Compute Engine Regional Instance Group Manager API creates and manages pools
// of homogeneous Compute Engine virtual machine instances from a common instance
// template. For more information, see [the official documentation](https://cloud.google.com/compute/docs/instance-groups/distributing-instances-with-regional-instance-groups)
// and [API](https://cloud.google.com/compute/docs/reference/latest/regionInstanceGroupManagers)
// 
// > **Note:** Use [google_compute_instance_group_manager](https://www.terraform.io/docs/providers/google/r/compute_instance_group_manager.html) to create a single-zone instance group manager.
type RegionInstanceGroupManager struct {
	s *pulumi.ResourceState
}

// NewRegionInstanceGroupManager registers a new resource with the given unique name, arguments, and options.
func NewRegionInstanceGroupManager(ctx *pulumi.Context,
	name string, args *RegionInstanceGroupManagerArgs, opts ...pulumi.ResourceOpt) (*RegionInstanceGroupManager, error) {
	if args == nil || args.BaseInstanceName == nil {
		return nil, errors.New("missing required argument 'BaseInstanceName'")
	}
	if args == nil || args.Region == nil {
		return nil, errors.New("missing required argument 'Region'")
	}
	inputs := make(map[string]interface{})
	if args == nil {
		inputs["autoHealingPolicies"] = nil
		inputs["baseInstanceName"] = nil
		inputs["description"] = nil
		inputs["distributionPolicyZones"] = nil
		inputs["instanceTemplate"] = nil
		inputs["name"] = nil
		inputs["namedPorts"] = nil
		inputs["project"] = nil
		inputs["region"] = nil
		inputs["rollingUpdatePolicy"] = nil
		inputs["targetPools"] = nil
		inputs["targetSize"] = nil
		inputs["updateStrategy"] = nil
		inputs["versions"] = nil
		inputs["waitForInstances"] = nil
	} else {
		inputs["autoHealingPolicies"] = args.AutoHealingPolicies
		inputs["baseInstanceName"] = args.BaseInstanceName
		inputs["description"] = args.Description
		inputs["distributionPolicyZones"] = args.DistributionPolicyZones
		inputs["instanceTemplate"] = args.InstanceTemplate
		inputs["name"] = args.Name
		inputs["namedPorts"] = args.NamedPorts
		inputs["project"] = args.Project
		inputs["region"] = args.Region
		inputs["rollingUpdatePolicy"] = args.RollingUpdatePolicy
		inputs["targetPools"] = args.TargetPools
		inputs["targetSize"] = args.TargetSize
		inputs["updateStrategy"] = args.UpdateStrategy
		inputs["versions"] = args.Versions
		inputs["waitForInstances"] = args.WaitForInstances
	}
	inputs["fingerprint"] = nil
	inputs["instanceGroup"] = nil
	inputs["selfLink"] = nil
	s, err := ctx.RegisterResource("gcp:compute/regionInstanceGroupManager:RegionInstanceGroupManager", name, true, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &RegionInstanceGroupManager{s: s}, nil
}

// GetRegionInstanceGroupManager gets an existing RegionInstanceGroupManager resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetRegionInstanceGroupManager(ctx *pulumi.Context,
	name string, id pulumi.ID, state *RegionInstanceGroupManagerState, opts ...pulumi.ResourceOpt) (*RegionInstanceGroupManager, error) {
	inputs := make(map[string]interface{})
	if state != nil {
		inputs["autoHealingPolicies"] = state.AutoHealingPolicies
		inputs["baseInstanceName"] = state.BaseInstanceName
		inputs["description"] = state.Description
		inputs["distributionPolicyZones"] = state.DistributionPolicyZones
		inputs["fingerprint"] = state.Fingerprint
		inputs["instanceGroup"] = state.InstanceGroup
		inputs["instanceTemplate"] = state.InstanceTemplate
		inputs["name"] = state.Name
		inputs["namedPorts"] = state.NamedPorts
		inputs["project"] = state.Project
		inputs["region"] = state.Region
		inputs["rollingUpdatePolicy"] = state.RollingUpdatePolicy
		inputs["selfLink"] = state.SelfLink
		inputs["targetPools"] = state.TargetPools
		inputs["targetSize"] = state.TargetSize
		inputs["updateStrategy"] = state.UpdateStrategy
		inputs["versions"] = state.Versions
		inputs["waitForInstances"] = state.WaitForInstances
	}
	s, err := ctx.ReadResource("gcp:compute/regionInstanceGroupManager:RegionInstanceGroupManager", name, id, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &RegionInstanceGroupManager{s: s}, nil
}

// URN is this resource's unique name assigned by Pulumi.
func (r *RegionInstanceGroupManager) URN() *pulumi.URNOutput {
	return r.s.URN()
}

// ID is this resource's unique identifier assigned by its provider.
func (r *RegionInstanceGroupManager) ID() *pulumi.IDOutput {
	return r.s.ID()
}

// The autohealing policies for this managed instance
// group. You can specify only one value. Structure is documented below. For more information, see the [official documentation](https://cloud.google.com/compute/docs/instance-groups/creating-groups-of-managed-instances#monitoring_groups).
// This property is in beta, and should be used with the terraform-provider-google-beta provider.
// See [Provider Versions](https://terraform.io/docs/providers/google/provider_versions.html) for more details on beta fields.
func (r *RegionInstanceGroupManager) AutoHealingPolicies() *pulumi.Output {
	return r.s.State["autoHealingPolicies"]
}

// The base instance name to use for
// instances in this group. The value must be a valid
// [RFC1035](https://www.ietf.org/rfc/rfc1035.txt) name. Supported characters
// are lowercase letters, numbers, and hyphens (-). Instances are named by
// appending a hyphen and a random four-character string to the base instance
// name.
func (r *RegionInstanceGroupManager) BaseInstanceName() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["baseInstanceName"])
}

// An optional textual description of the instance
// group manager.
func (r *RegionInstanceGroupManager) Description() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["description"])
}

// The distribution policy for this managed instance
// group. You can specify one or more values. For more information, see the [official documentation](https://cloud.google.com/compute/docs/instance-groups/distributing-instances-with-regional-instance-groups#selectingzones).
// - - -
func (r *RegionInstanceGroupManager) DistributionPolicyZones() *pulumi.ArrayOutput {
	return (*pulumi.ArrayOutput)(r.s.State["distributionPolicyZones"])
}

// The fingerprint of the instance group manager.
func (r *RegionInstanceGroupManager) Fingerprint() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["fingerprint"])
}

// The full URL of the instance group created by the manager.
func (r *RegionInstanceGroupManager) InstanceGroup() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["instanceGroup"])
}

// - The full URL to an instance template from which all new instances of this version will be created.
func (r *RegionInstanceGroupManager) InstanceTemplate() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["instanceTemplate"])
}

// - Version name.
func (r *RegionInstanceGroupManager) Name() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["name"])
}

// The named port configuration. See the section below
// for details on configuration.
func (r *RegionInstanceGroupManager) NamedPorts() *pulumi.ArrayOutput {
	return (*pulumi.ArrayOutput)(r.s.State["namedPorts"])
}

// The ID of the project in which the resource belongs. If it
// is not provided, the provider project is used.
func (r *RegionInstanceGroupManager) Project() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["project"])
}

// The region where the managed instance group resides.
func (r *RegionInstanceGroupManager) Region() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["region"])
}

// The update policy for this managed instance group. Structure is documented below. For more information, see the [official documentation](https://cloud.google.com/compute/docs/instance-groups/updating-managed-instance-groups) and [API](https://cloud.google.com/compute/docs/reference/rest/beta/regionInstanceGroupManagers/patch)
// This property is in beta, and should be used with the terraform-provider-google-beta provider.
// See [Provider Versions](https://terraform.io/docs/providers/google/provider_versions.html) for more details on beta fields.
func (r *RegionInstanceGroupManager) RollingUpdatePolicy() *pulumi.Output {
	return r.s.State["rollingUpdatePolicy"]
}

// The URL of the created resource.
func (r *RegionInstanceGroupManager) SelfLink() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["selfLink"])
}

// The full URL of all target pools to which new
// instances in the group are added. Updating the target pools attribute does
// not affect existing instances.
func (r *RegionInstanceGroupManager) TargetPools() *pulumi.ArrayOutput {
	return (*pulumi.ArrayOutput)(r.s.State["targetPools"])
}

// - The number of instances calculated as a fixed number or a percentage depending on the settings. Structure is documented below.
func (r *RegionInstanceGroupManager) TargetSize() *pulumi.IntOutput {
	return (*pulumi.IntOutput)(r.s.State["targetSize"])
}

// If the `instance_template`
// resource is modified, a value of `"NONE"` will prevent any of the managed
// instances from being restarted by Terraform. A value of `"ROLLING_UPDATE"`
// is supported as a beta feature. A value of `"ROLLING_UPDATE"` requires
// `rolling_update_policy` block to be set. This field is deprecated as in
// `2.0.0` it has no functionality anymore. It will be removed then. This field
// is only present in the `google` provider.
func (r *RegionInstanceGroupManager) UpdateStrategy() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["updateStrategy"])
}

// Application versions managed by this instance group. Each
// version deals with a specific instance template, allowing canary release scenarios.
// Conflicts with `instance_template`. Structure is documented below. Beware that
// exactly one version must not specify a target size. It means that versions with
// a target size will respect the setting, and the one without target size will
// be applied to all remaining Instances (top level target_size - each version target_size).
// This property is in beta, and should be used with the terraform-provider-google-beta provider.
// See [Provider Versions](https://terraform.io/docs/providers/google/provider_versions.html) for more details on beta fields.
func (r *RegionInstanceGroupManager) Versions() *pulumi.ArrayOutput {
	return (*pulumi.ArrayOutput)(r.s.State["versions"])
}

// Whether to wait for all instances to be created/updated before
// returning. Note that if this is set to true and the operation does not succeed, Terraform will
// continue trying until it times out.
func (r *RegionInstanceGroupManager) WaitForInstances() *pulumi.BoolOutput {
	return (*pulumi.BoolOutput)(r.s.State["waitForInstances"])
}

// Input properties used for looking up and filtering RegionInstanceGroupManager resources.
type RegionInstanceGroupManagerState struct {
	// The autohealing policies for this managed instance
	// group. You can specify only one value. Structure is documented below. For more information, see the [official documentation](https://cloud.google.com/compute/docs/instance-groups/creating-groups-of-managed-instances#monitoring_groups).
	// This property is in beta, and should be used with the terraform-provider-google-beta provider.
	// See [Provider Versions](https://terraform.io/docs/providers/google/provider_versions.html) for more details on beta fields.
	AutoHealingPolicies interface{}
	// The base instance name to use for
	// instances in this group. The value must be a valid
	// [RFC1035](https://www.ietf.org/rfc/rfc1035.txt) name. Supported characters
	// are lowercase letters, numbers, and hyphens (-). Instances are named by
	// appending a hyphen and a random four-character string to the base instance
	// name.
	BaseInstanceName interface{}
	// An optional textual description of the instance
	// group manager.
	Description interface{}
	// The distribution policy for this managed instance
	// group. You can specify one or more values. For more information, see the [official documentation](https://cloud.google.com/compute/docs/instance-groups/distributing-instances-with-regional-instance-groups#selectingzones).
	// - - -
	DistributionPolicyZones interface{}
	// The fingerprint of the instance group manager.
	Fingerprint interface{}
	// The full URL of the instance group created by the manager.
	InstanceGroup interface{}
	// - The full URL to an instance template from which all new instances of this version will be created.
	InstanceTemplate interface{}
	// - Version name.
	Name interface{}
	// The named port configuration. See the section below
	// for details on configuration.
	NamedPorts interface{}
	// The ID of the project in which the resource belongs. If it
	// is not provided, the provider project is used.
	Project interface{}
	// The region where the managed instance group resides.
	Region interface{}
	// The update policy for this managed instance group. Structure is documented below. For more information, see the [official documentation](https://cloud.google.com/compute/docs/instance-groups/updating-managed-instance-groups) and [API](https://cloud.google.com/compute/docs/reference/rest/beta/regionInstanceGroupManagers/patch)
	// This property is in beta, and should be used with the terraform-provider-google-beta provider.
	// See [Provider Versions](https://terraform.io/docs/providers/google/provider_versions.html) for more details on beta fields.
	RollingUpdatePolicy interface{}
	// The URL of the created resource.
	SelfLink interface{}
	// The full URL of all target pools to which new
	// instances in the group are added. Updating the target pools attribute does
	// not affect existing instances.
	TargetPools interface{}
	// - The number of instances calculated as a fixed number or a percentage depending on the settings. Structure is documented below.
	TargetSize interface{}
	// If the `instance_template`
	// resource is modified, a value of `"NONE"` will prevent any of the managed
	// instances from being restarted by Terraform. A value of `"ROLLING_UPDATE"`
	// is supported as a beta feature. A value of `"ROLLING_UPDATE"` requires
	// `rolling_update_policy` block to be set. This field is deprecated as in
	// `2.0.0` it has no functionality anymore. It will be removed then. This field
	// is only present in the `google` provider.
	UpdateStrategy interface{}
	// Application versions managed by this instance group. Each
	// version deals with a specific instance template, allowing canary release scenarios.
	// Conflicts with `instance_template`. Structure is documented below. Beware that
	// exactly one version must not specify a target size. It means that versions with
	// a target size will respect the setting, and the one without target size will
	// be applied to all remaining Instances (top level target_size - each version target_size).
	// This property is in beta, and should be used with the terraform-provider-google-beta provider.
	// See [Provider Versions](https://terraform.io/docs/providers/google/provider_versions.html) for more details on beta fields.
	Versions interface{}
	// Whether to wait for all instances to be created/updated before
	// returning. Note that if this is set to true and the operation does not succeed, Terraform will
	// continue trying until it times out.
	WaitForInstances interface{}
}

// The set of arguments for constructing a RegionInstanceGroupManager resource.
type RegionInstanceGroupManagerArgs struct {
	// The autohealing policies for this managed instance
	// group. You can specify only one value. Structure is documented below. For more information, see the [official documentation](https://cloud.google.com/compute/docs/instance-groups/creating-groups-of-managed-instances#monitoring_groups).
	// This property is in beta, and should be used with the terraform-provider-google-beta provider.
	// See [Provider Versions](https://terraform.io/docs/providers/google/provider_versions.html) for more details on beta fields.
	AutoHealingPolicies interface{}
	// The base instance name to use for
	// instances in this group. The value must be a valid
	// [RFC1035](https://www.ietf.org/rfc/rfc1035.txt) name. Supported characters
	// are lowercase letters, numbers, and hyphens (-). Instances are named by
	// appending a hyphen and a random four-character string to the base instance
	// name.
	BaseInstanceName interface{}
	// An optional textual description of the instance
	// group manager.
	Description interface{}
	// The distribution policy for this managed instance
	// group. You can specify one or more values. For more information, see the [official documentation](https://cloud.google.com/compute/docs/instance-groups/distributing-instances-with-regional-instance-groups#selectingzones).
	// - - -
	DistributionPolicyZones interface{}
	// - The full URL to an instance template from which all new instances of this version will be created.
	InstanceTemplate interface{}
	// - Version name.
	Name interface{}
	// The named port configuration. See the section below
	// for details on configuration.
	NamedPorts interface{}
	// The ID of the project in which the resource belongs. If it
	// is not provided, the provider project is used.
	Project interface{}
	// The region where the managed instance group resides.
	Region interface{}
	// The update policy for this managed instance group. Structure is documented below. For more information, see the [official documentation](https://cloud.google.com/compute/docs/instance-groups/updating-managed-instance-groups) and [API](https://cloud.google.com/compute/docs/reference/rest/beta/regionInstanceGroupManagers/patch)
	// This property is in beta, and should be used with the terraform-provider-google-beta provider.
	// See [Provider Versions](https://terraform.io/docs/providers/google/provider_versions.html) for more details on beta fields.
	RollingUpdatePolicy interface{}
	// The full URL of all target pools to which new
	// instances in the group are added. Updating the target pools attribute does
	// not affect existing instances.
	TargetPools interface{}
	// - The number of instances calculated as a fixed number or a percentage depending on the settings. Structure is documented below.
	TargetSize interface{}
	// If the `instance_template`
	// resource is modified, a value of `"NONE"` will prevent any of the managed
	// instances from being restarted by Terraform. A value of `"ROLLING_UPDATE"`
	// is supported as a beta feature. A value of `"ROLLING_UPDATE"` requires
	// `rolling_update_policy` block to be set. This field is deprecated as in
	// `2.0.0` it has no functionality anymore. It will be removed then. This field
	// is only present in the `google` provider.
	UpdateStrategy interface{}
	// Application versions managed by this instance group. Each
	// version deals with a specific instance template, allowing canary release scenarios.
	// Conflicts with `instance_template`. Structure is documented below. Beware that
	// exactly one version must not specify a target size. It means that versions with
	// a target size will respect the setting, and the one without target size will
	// be applied to all remaining Instances (top level target_size - each version target_size).
	// This property is in beta, and should be used with the terraform-provider-google-beta provider.
	// See [Provider Versions](https://terraform.io/docs/providers/google/provider_versions.html) for more details on beta fields.
	Versions interface{}
	// Whether to wait for all instances to be created/updated before
	// returning. Note that if this is set to true and the operation does not succeed, Terraform will
	// continue trying until it times out.
	WaitForInstances interface{}
}
