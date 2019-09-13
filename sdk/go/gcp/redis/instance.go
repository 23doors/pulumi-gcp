// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package redis

import (
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// > This content is derived from https://github.com/terraform-providers/terraform-provider-google/blob/master/website/docs/r/redis_instance.html.markdown.
type Instance struct {
	s *pulumi.ResourceState
}

// NewInstance registers a new resource with the given unique name, arguments, and options.
func NewInstance(ctx *pulumi.Context,
	name string, args *InstanceArgs, opts ...pulumi.ResourceOpt) (*Instance, error) {
	if args == nil || args.MemorySizeGb == nil {
		return nil, errors.New("missing required argument 'MemorySizeGb'")
	}
	inputs := make(map[string]interface{})
	if args == nil {
		inputs["alternativeLocationId"] = nil
		inputs["authorizedNetwork"] = nil
		inputs["displayName"] = nil
		inputs["labels"] = nil
		inputs["locationId"] = nil
		inputs["memorySizeGb"] = nil
		inputs["name"] = nil
		inputs["project"] = nil
		inputs["redisConfigs"] = nil
		inputs["redisVersion"] = nil
		inputs["region"] = nil
		inputs["reservedIpRange"] = nil
		inputs["tier"] = nil
	} else {
		inputs["alternativeLocationId"] = args.AlternativeLocationId
		inputs["authorizedNetwork"] = args.AuthorizedNetwork
		inputs["displayName"] = args.DisplayName
		inputs["labels"] = args.Labels
		inputs["locationId"] = args.LocationId
		inputs["memorySizeGb"] = args.MemorySizeGb
		inputs["name"] = args.Name
		inputs["project"] = args.Project
		inputs["redisConfigs"] = args.RedisConfigs
		inputs["redisVersion"] = args.RedisVersion
		inputs["region"] = args.Region
		inputs["reservedIpRange"] = args.ReservedIpRange
		inputs["tier"] = args.Tier
	}
	inputs["createTime"] = nil
	inputs["currentLocationId"] = nil
	inputs["host"] = nil
	inputs["port"] = nil
	s, err := ctx.RegisterResource("gcp:redis/instance:Instance", name, true, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &Instance{s: s}, nil
}

// GetInstance gets an existing Instance resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetInstance(ctx *pulumi.Context,
	name string, id pulumi.ID, state *InstanceState, opts ...pulumi.ResourceOpt) (*Instance, error) {
	inputs := make(map[string]interface{})
	if state != nil {
		inputs["alternativeLocationId"] = state.AlternativeLocationId
		inputs["authorizedNetwork"] = state.AuthorizedNetwork
		inputs["createTime"] = state.CreateTime
		inputs["currentLocationId"] = state.CurrentLocationId
		inputs["displayName"] = state.DisplayName
		inputs["host"] = state.Host
		inputs["labels"] = state.Labels
		inputs["locationId"] = state.LocationId
		inputs["memorySizeGb"] = state.MemorySizeGb
		inputs["name"] = state.Name
		inputs["port"] = state.Port
		inputs["project"] = state.Project
		inputs["redisConfigs"] = state.RedisConfigs
		inputs["redisVersion"] = state.RedisVersion
		inputs["region"] = state.Region
		inputs["reservedIpRange"] = state.ReservedIpRange
		inputs["tier"] = state.Tier
	}
	s, err := ctx.ReadResource("gcp:redis/instance:Instance", name, id, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &Instance{s: s}, nil
}

// URN is this resource's unique name assigned by Pulumi.
func (r *Instance) URN() *pulumi.URNOutput {
	return r.s.URN()
}

// ID is this resource's unique identifier assigned by its provider.
func (r *Instance) ID() *pulumi.IDOutput {
	return r.s.ID()
}

func (r *Instance) AlternativeLocationId() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["alternativeLocationId"])
}

func (r *Instance) AuthorizedNetwork() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["authorizedNetwork"])
}

func (r *Instance) CreateTime() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["createTime"])
}

func (r *Instance) CurrentLocationId() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["currentLocationId"])
}

func (r *Instance) DisplayName() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["displayName"])
}

func (r *Instance) Host() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["host"])
}

func (r *Instance) Labels() *pulumi.MapOutput {
	return (*pulumi.MapOutput)(r.s.State["labels"])
}

func (r *Instance) LocationId() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["locationId"])
}

func (r *Instance) MemorySizeGb() *pulumi.IntOutput {
	return (*pulumi.IntOutput)(r.s.State["memorySizeGb"])
}

func (r *Instance) Name() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["name"])
}

func (r *Instance) Port() *pulumi.IntOutput {
	return (*pulumi.IntOutput)(r.s.State["port"])
}

// The ID of the project in which the resource belongs.
// If it is not provided, the provider project is used.
func (r *Instance) Project() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["project"])
}

func (r *Instance) RedisConfigs() *pulumi.MapOutput {
	return (*pulumi.MapOutput)(r.s.State["redisConfigs"])
}

func (r *Instance) RedisVersion() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["redisVersion"])
}

func (r *Instance) Region() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["region"])
}

func (r *Instance) ReservedIpRange() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["reservedIpRange"])
}

func (r *Instance) Tier() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["tier"])
}

// Input properties used for looking up and filtering Instance resources.
type InstanceState struct {
	AlternativeLocationId interface{}
	AuthorizedNetwork interface{}
	CreateTime interface{}
	CurrentLocationId interface{}
	DisplayName interface{}
	Host interface{}
	Labels interface{}
	LocationId interface{}
	MemorySizeGb interface{}
	Name interface{}
	Port interface{}
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project interface{}
	RedisConfigs interface{}
	RedisVersion interface{}
	Region interface{}
	ReservedIpRange interface{}
	Tier interface{}
}

// The set of arguments for constructing a Instance resource.
type InstanceArgs struct {
	AlternativeLocationId interface{}
	AuthorizedNetwork interface{}
	DisplayName interface{}
	Labels interface{}
	LocationId interface{}
	MemorySizeGb interface{}
	Name interface{}
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project interface{}
	RedisConfigs interface{}
	RedisVersion interface{}
	Region interface{}
	ReservedIpRange interface{}
	Tier interface{}
}
