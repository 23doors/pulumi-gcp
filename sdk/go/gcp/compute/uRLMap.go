// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package compute

import (
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// > This content is derived from https://github.com/terraform-providers/terraform-provider-google/blob/master/website/docs/r/compute_url_map.html.markdown.
type URLMap struct {
	s *pulumi.ResourceState
}

// NewURLMap registers a new resource with the given unique name, arguments, and options.
func NewURLMap(ctx *pulumi.Context,
	name string, args *URLMapArgs, opts ...pulumi.ResourceOpt) (*URLMap, error) {
	if args == nil || args.DefaultService == nil {
		return nil, errors.New("missing required argument 'DefaultService'")
	}
	inputs := make(map[string]interface{})
	if args == nil {
		inputs["defaultService"] = nil
		inputs["description"] = nil
		inputs["hostRules"] = nil
		inputs["name"] = nil
		inputs["pathMatchers"] = nil
		inputs["project"] = nil
		inputs["tests"] = nil
	} else {
		inputs["defaultService"] = args.DefaultService
		inputs["description"] = args.Description
		inputs["hostRules"] = args.HostRules
		inputs["name"] = args.Name
		inputs["pathMatchers"] = args.PathMatchers
		inputs["project"] = args.Project
		inputs["tests"] = args.Tests
	}
	inputs["creationTimestamp"] = nil
	inputs["fingerprint"] = nil
	inputs["mapId"] = nil
	inputs["selfLink"] = nil
	s, err := ctx.RegisterResource("gcp:compute/uRLMap:URLMap", name, true, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &URLMap{s: s}, nil
}

// GetURLMap gets an existing URLMap resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetURLMap(ctx *pulumi.Context,
	name string, id pulumi.ID, state *URLMapState, opts ...pulumi.ResourceOpt) (*URLMap, error) {
	inputs := make(map[string]interface{})
	if state != nil {
		inputs["creationTimestamp"] = state.CreationTimestamp
		inputs["defaultService"] = state.DefaultService
		inputs["description"] = state.Description
		inputs["fingerprint"] = state.Fingerprint
		inputs["hostRules"] = state.HostRules
		inputs["mapId"] = state.MapId
		inputs["name"] = state.Name
		inputs["pathMatchers"] = state.PathMatchers
		inputs["project"] = state.Project
		inputs["selfLink"] = state.SelfLink
		inputs["tests"] = state.Tests
	}
	s, err := ctx.ReadResource("gcp:compute/uRLMap:URLMap", name, id, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &URLMap{s: s}, nil
}

// URN is this resource's unique name assigned by Pulumi.
func (r *URLMap) URN() *pulumi.URNOutput {
	return r.s.URN()
}

// ID is this resource's unique identifier assigned by its provider.
func (r *URLMap) ID() *pulumi.IDOutput {
	return r.s.ID()
}

func (r *URLMap) CreationTimestamp() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["creationTimestamp"])
}

func (r *URLMap) DefaultService() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["defaultService"])
}

func (r *URLMap) Description() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["description"])
}

func (r *URLMap) Fingerprint() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["fingerprint"])
}

func (r *URLMap) HostRules() *pulumi.ArrayOutput {
	return (*pulumi.ArrayOutput)(r.s.State["hostRules"])
}

func (r *URLMap) MapId() *pulumi.IntOutput {
	return (*pulumi.IntOutput)(r.s.State["mapId"])
}

func (r *URLMap) Name() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["name"])
}

func (r *URLMap) PathMatchers() *pulumi.ArrayOutput {
	return (*pulumi.ArrayOutput)(r.s.State["pathMatchers"])
}

// The ID of the project in which the resource belongs.
// If it is not provided, the provider project is used.
func (r *URLMap) Project() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["project"])
}

// The URI of the created resource.
func (r *URLMap) SelfLink() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["selfLink"])
}

func (r *URLMap) Tests() *pulumi.ArrayOutput {
	return (*pulumi.ArrayOutput)(r.s.State["tests"])
}

// Input properties used for looking up and filtering URLMap resources.
type URLMapState struct {
	CreationTimestamp interface{}
	DefaultService interface{}
	Description interface{}
	Fingerprint interface{}
	HostRules interface{}
	MapId interface{}
	Name interface{}
	PathMatchers interface{}
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project interface{}
	// The URI of the created resource.
	SelfLink interface{}
	Tests interface{}
}

// The set of arguments for constructing a URLMap resource.
type URLMapArgs struct {
	DefaultService interface{}
	Description interface{}
	HostRules interface{}
	Name interface{}
	PathMatchers interface{}
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project interface{}
	Tests interface{}
}
