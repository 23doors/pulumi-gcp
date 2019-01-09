// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package compute

import (
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Each network has its own firewall controlling access to and from the
// instances.
// 
// All traffic to instances, even from other instances, is blocked by the
// firewall unless firewall rules are created to allow it.
// 
// The default network has automatically created firewall rules that are
// shown in default firewall rules. No manually created network has
// automatically created firewall rules except for a default "allow" rule for
// outgoing traffic and a default "deny" for incoming traffic. For all
// networks except the default network, you must create any firewall rules
// you need.
// 
// 
// To get more information about Firewall, see:
// 
// * [API documentation](https://cloud.google.com/compute/docs/reference/latest/firewalls)
// * How-to Guides
//     * [Official Documentation](https://cloud.google.com/vpc/docs/firewalls)
// 
// <div class = "oics-button" style="float: right; margin: 0 0 -15px">
//   <a href="https://console.cloud.google.com/cloudshell/open?cloudshell_git_repo=https%3A%2F%2Fgithub.com%2Fterraform-google-modules%2Fdocs-examples.git&cloudshell_working_dir=firewall_basic&cloudshell_image=gcr.io%2Fgraphite-cloud-shell-images%2Fterraform%3Alatest&open_in_editor=main.tf&cloudshell_print=.%2Fmotd&cloudshell_tutorial=.%2Ftutorial.md" target="_blank">
//     <img alt="Open in Cloud Shell" src="//gstatic.com/cloudssh/images/open-btn.svg" style="max-height: 44px; margin: 32px auto; max-width: 100%;">
//   </a>
// </div>
type Firewall struct {
	s *pulumi.ResourceState
}

// NewFirewall registers a new resource with the given unique name, arguments, and options.
func NewFirewall(ctx *pulumi.Context,
	name string, args *FirewallArgs, opts ...pulumi.ResourceOpt) (*Firewall, error) {
	if args == nil || args.Network == nil {
		return nil, errors.New("missing required argument 'Network'")
	}
	inputs := make(map[string]interface{})
	if args == nil {
		inputs["allows"] = nil
		inputs["denies"] = nil
		inputs["description"] = nil
		inputs["destinationRanges"] = nil
		inputs["direction"] = nil
		inputs["disabled"] = nil
		inputs["enableLogging"] = nil
		inputs["name"] = nil
		inputs["network"] = nil
		inputs["priority"] = nil
		inputs["project"] = nil
		inputs["sourceRanges"] = nil
		inputs["sourceServiceAccounts"] = nil
		inputs["sourceTags"] = nil
		inputs["targetServiceAccounts"] = nil
		inputs["targetTags"] = nil
	} else {
		inputs["allows"] = args.Allows
		inputs["denies"] = args.Denies
		inputs["description"] = args.Description
		inputs["destinationRanges"] = args.DestinationRanges
		inputs["direction"] = args.Direction
		inputs["disabled"] = args.Disabled
		inputs["enableLogging"] = args.EnableLogging
		inputs["name"] = args.Name
		inputs["network"] = args.Network
		inputs["priority"] = args.Priority
		inputs["project"] = args.Project
		inputs["sourceRanges"] = args.SourceRanges
		inputs["sourceServiceAccounts"] = args.SourceServiceAccounts
		inputs["sourceTags"] = args.SourceTags
		inputs["targetServiceAccounts"] = args.TargetServiceAccounts
		inputs["targetTags"] = args.TargetTags
	}
	inputs["creationTimestamp"] = nil
	inputs["selfLink"] = nil
	s, err := ctx.RegisterResource("gcp:compute/firewall:Firewall", name, true, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &Firewall{s: s}, nil
}

// GetFirewall gets an existing Firewall resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetFirewall(ctx *pulumi.Context,
	name string, id pulumi.ID, state *FirewallState, opts ...pulumi.ResourceOpt) (*Firewall, error) {
	inputs := make(map[string]interface{})
	if state != nil {
		inputs["allows"] = state.Allows
		inputs["creationTimestamp"] = state.CreationTimestamp
		inputs["denies"] = state.Denies
		inputs["description"] = state.Description
		inputs["destinationRanges"] = state.DestinationRanges
		inputs["direction"] = state.Direction
		inputs["disabled"] = state.Disabled
		inputs["enableLogging"] = state.EnableLogging
		inputs["name"] = state.Name
		inputs["network"] = state.Network
		inputs["priority"] = state.Priority
		inputs["project"] = state.Project
		inputs["selfLink"] = state.SelfLink
		inputs["sourceRanges"] = state.SourceRanges
		inputs["sourceServiceAccounts"] = state.SourceServiceAccounts
		inputs["sourceTags"] = state.SourceTags
		inputs["targetServiceAccounts"] = state.TargetServiceAccounts
		inputs["targetTags"] = state.TargetTags
	}
	s, err := ctx.ReadResource("gcp:compute/firewall:Firewall", name, id, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &Firewall{s: s}, nil
}

// URN is this resource's unique name assigned by Pulumi.
func (r *Firewall) URN() *pulumi.URNOutput {
	return r.s.URN()
}

// ID is this resource's unique identifier assigned by its provider.
func (r *Firewall) ID() *pulumi.IDOutput {
	return r.s.ID()
}

func (r *Firewall) Allows() *pulumi.ArrayOutput {
	return (*pulumi.ArrayOutput)(r.s.State["allows"])
}

func (r *Firewall) CreationTimestamp() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["creationTimestamp"])
}

func (r *Firewall) Denies() *pulumi.ArrayOutput {
	return (*pulumi.ArrayOutput)(r.s.State["denies"])
}

func (r *Firewall) Description() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["description"])
}

func (r *Firewall) DestinationRanges() *pulumi.ArrayOutput {
	return (*pulumi.ArrayOutput)(r.s.State["destinationRanges"])
}

func (r *Firewall) Direction() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["direction"])
}

func (r *Firewall) Disabled() *pulumi.BoolOutput {
	return (*pulumi.BoolOutput)(r.s.State["disabled"])
}

func (r *Firewall) EnableLogging() *pulumi.BoolOutput {
	return (*pulumi.BoolOutput)(r.s.State["enableLogging"])
}

func (r *Firewall) Name() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["name"])
}

func (r *Firewall) Network() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["network"])
}

func (r *Firewall) Priority() *pulumi.IntOutput {
	return (*pulumi.IntOutput)(r.s.State["priority"])
}

// The ID of the project in which the resource belongs.
// If it is not provided, the provider project is used.
func (r *Firewall) Project() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["project"])
}

// The URI of the created resource.
func (r *Firewall) SelfLink() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["selfLink"])
}

func (r *Firewall) SourceRanges() *pulumi.ArrayOutput {
	return (*pulumi.ArrayOutput)(r.s.State["sourceRanges"])
}

func (r *Firewall) SourceServiceAccounts() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["sourceServiceAccounts"])
}

func (r *Firewall) SourceTags() *pulumi.ArrayOutput {
	return (*pulumi.ArrayOutput)(r.s.State["sourceTags"])
}

func (r *Firewall) TargetServiceAccounts() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["targetServiceAccounts"])
}

func (r *Firewall) TargetTags() *pulumi.ArrayOutput {
	return (*pulumi.ArrayOutput)(r.s.State["targetTags"])
}

// Input properties used for looking up and filtering Firewall resources.
type FirewallState struct {
	Allows interface{}
	CreationTimestamp interface{}
	Denies interface{}
	Description interface{}
	DestinationRanges interface{}
	Direction interface{}
	Disabled interface{}
	EnableLogging interface{}
	Name interface{}
	Network interface{}
	Priority interface{}
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project interface{}
	// The URI of the created resource.
	SelfLink interface{}
	SourceRanges interface{}
	SourceServiceAccounts interface{}
	SourceTags interface{}
	TargetServiceAccounts interface{}
	TargetTags interface{}
}

// The set of arguments for constructing a Firewall resource.
type FirewallArgs struct {
	Allows interface{}
	Denies interface{}
	Description interface{}
	DestinationRanges interface{}
	Direction interface{}
	Disabled interface{}
	EnableLogging interface{}
	Name interface{}
	Network interface{}
	Priority interface{}
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project interface{}
	SourceRanges interface{}
	SourceServiceAccounts interface{}
	SourceTags interface{}
	TargetServiceAccounts interface{}
	TargetTags interface{}
}
