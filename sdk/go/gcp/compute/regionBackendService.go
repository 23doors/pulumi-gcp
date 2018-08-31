// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package compute

import (
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// A Region Backend Service defines a regionally-scoped group of virtual machines that will serve traffic for load balancing.
// For more information see [the official documentation](https://cloud.google.com/compute/docs/load-balancing/internal/)
// and [API](https://cloud.google.com/compute/docs/reference/latest/regionBackendServices).
// 
// ~> **Note**: Region backend services can only be used when using internal load balancing. For external load balancing, use
//   `google_compute_backend_service` instead.
type RegionBackendService struct {
	s *pulumi.ResourceState
}

// NewRegionBackendService registers a new resource with the given unique name, arguments, and options.
func NewRegionBackendService(ctx *pulumi.Context,
	name string, args *RegionBackendServiceArgs, opts ...pulumi.ResourceOpt) (*RegionBackendService, error) {
	if args == nil || args.HealthChecks == nil {
		return nil, errors.New("missing required argument 'HealthChecks'")
	}
	inputs := make(map[string]interface{})
	if args == nil {
		inputs["backends"] = nil
		inputs["connectionDrainingTimeoutSec"] = nil
		inputs["description"] = nil
		inputs["healthChecks"] = nil
		inputs["name"] = nil
		inputs["project"] = nil
		inputs["protocol"] = nil
		inputs["region"] = nil
		inputs["sessionAffinity"] = nil
		inputs["timeoutSec"] = nil
	} else {
		inputs["backends"] = args.Backends
		inputs["connectionDrainingTimeoutSec"] = args.ConnectionDrainingTimeoutSec
		inputs["description"] = args.Description
		inputs["healthChecks"] = args.HealthChecks
		inputs["name"] = args.Name
		inputs["project"] = args.Project
		inputs["protocol"] = args.Protocol
		inputs["region"] = args.Region
		inputs["sessionAffinity"] = args.SessionAffinity
		inputs["timeoutSec"] = args.TimeoutSec
	}
	inputs["fingerprint"] = nil
	inputs["selfLink"] = nil
	s, err := ctx.RegisterResource("gcp:compute/regionBackendService:RegionBackendService", name, true, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &RegionBackendService{s: s}, nil
}

// GetRegionBackendService gets an existing RegionBackendService resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetRegionBackendService(ctx *pulumi.Context,
	name string, id pulumi.ID, state *RegionBackendServiceState, opts ...pulumi.ResourceOpt) (*RegionBackendService, error) {
	inputs := make(map[string]interface{})
	if state != nil {
		inputs["backends"] = state.Backends
		inputs["connectionDrainingTimeoutSec"] = state.ConnectionDrainingTimeoutSec
		inputs["description"] = state.Description
		inputs["fingerprint"] = state.Fingerprint
		inputs["healthChecks"] = state.HealthChecks
		inputs["name"] = state.Name
		inputs["project"] = state.Project
		inputs["protocol"] = state.Protocol
		inputs["region"] = state.Region
		inputs["selfLink"] = state.SelfLink
		inputs["sessionAffinity"] = state.SessionAffinity
		inputs["timeoutSec"] = state.TimeoutSec
	}
	s, err := ctx.ReadResource("gcp:compute/regionBackendService:RegionBackendService", name, id, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &RegionBackendService{s: s}, nil
}

// URN is this resource's unique name assigned by Pulumi.
func (r *RegionBackendService) URN() *pulumi.URNOutput {
	return r.s.URN
}

// ID is this resource's unique identifier assigned by its provider.
func (r *RegionBackendService) ID() *pulumi.IDOutput {
	return r.s.ID
}

// The list of backends that serve this BackendService.
// Structure is documented below.
func (r *RegionBackendService) Backends() *pulumi.ArrayOutput {
	return (*pulumi.ArrayOutput)(r.s.State["backends"])
}

// Time for which instance will be drained
// (not accept new connections, but still work to finish started ones). Defaults to `0`.
func (r *RegionBackendService) ConnectionDrainingTimeoutSec() *pulumi.IntOutput {
	return (*pulumi.IntOutput)(r.s.State["connectionDrainingTimeoutSec"])
}

// The textual description for the backend service.
func (r *RegionBackendService) Description() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["description"])
}

// The fingerprint of the backend service.
func (r *RegionBackendService) Fingerprint() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["fingerprint"])
}

// Specifies a list of health checks
// for checking the health of the backend service. Currently at most
// one health check can be specified, and a health check is required.
func (r *RegionBackendService) HealthChecks() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["healthChecks"])
}

// The name of the backend service.
func (r *RegionBackendService) Name() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["name"])
}

// The ID of the project in which the resource belongs. If it
// is not provided, the provider project is used.
func (r *RegionBackendService) Project() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["project"])
}

// The protocol for incoming requests. Defaults to
// `TCP`.
func (r *RegionBackendService) Protocol() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["protocol"])
}

// The Region in which the created address should reside.
// If it is not provided, the provider region is used.
func (r *RegionBackendService) Region() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["region"])
}

// The URI of the created resource.
func (r *RegionBackendService) SelfLink() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["selfLink"])
}

// How to distribute load. Options are `NONE` (no
// affinity), `CLIENT_IP`, `CLIENT_IP_PROTO`, or `CLIENT_IP_PORT_PROTO`.
// Defaults to `NONE`.
func (r *RegionBackendService) SessionAffinity() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["sessionAffinity"])
}

// The number of secs to wait for a backend to respond
// to a request before considering the request failed. Defaults to `30`.
func (r *RegionBackendService) TimeoutSec() *pulumi.IntOutput {
	return (*pulumi.IntOutput)(r.s.State["timeoutSec"])
}

// Input properties used for looking up and filtering RegionBackendService resources.
type RegionBackendServiceState struct {
	// The list of backends that serve this BackendService.
	// Structure is documented below.
	Backends interface{}
	// Time for which instance will be drained
	// (not accept new connections, but still work to finish started ones). Defaults to `0`.
	ConnectionDrainingTimeoutSec interface{}
	// The textual description for the backend service.
	Description interface{}
	// The fingerprint of the backend service.
	Fingerprint interface{}
	// Specifies a list of health checks
	// for checking the health of the backend service. Currently at most
	// one health check can be specified, and a health check is required.
	HealthChecks interface{}
	// The name of the backend service.
	Name interface{}
	// The ID of the project in which the resource belongs. If it
	// is not provided, the provider project is used.
	Project interface{}
	// The protocol for incoming requests. Defaults to
	// `TCP`.
	Protocol interface{}
	// The Region in which the created address should reside.
	// If it is not provided, the provider region is used.
	Region interface{}
	// The URI of the created resource.
	SelfLink interface{}
	// How to distribute load. Options are `NONE` (no
	// affinity), `CLIENT_IP`, `CLIENT_IP_PROTO`, or `CLIENT_IP_PORT_PROTO`.
	// Defaults to `NONE`.
	SessionAffinity interface{}
	// The number of secs to wait for a backend to respond
	// to a request before considering the request failed. Defaults to `30`.
	TimeoutSec interface{}
}

// The set of arguments for constructing a RegionBackendService resource.
type RegionBackendServiceArgs struct {
	// The list of backends that serve this BackendService.
	// Structure is documented below.
	Backends interface{}
	// Time for which instance will be drained
	// (not accept new connections, but still work to finish started ones). Defaults to `0`.
	ConnectionDrainingTimeoutSec interface{}
	// The textual description for the backend service.
	Description interface{}
	// Specifies a list of health checks
	// for checking the health of the backend service. Currently at most
	// one health check can be specified, and a health check is required.
	HealthChecks interface{}
	// The name of the backend service.
	Name interface{}
	// The ID of the project in which the resource belongs. If it
	// is not provided, the provider project is used.
	Project interface{}
	// The protocol for incoming requests. Defaults to
	// `TCP`.
	Protocol interface{}
	// The Region in which the created address should reside.
	// If it is not provided, the provider region is used.
	Region interface{}
	// How to distribute load. Options are `NONE` (no
	// affinity), `CLIENT_IP`, `CLIENT_IP_PROTO`, or `CLIENT_IP_PORT_PROTO`.
	// Defaults to `NONE`.
	SessionAffinity interface{}
	// The number of secs to wait for a backend to respond
	// to a request before considering the request failed. Defaults to `30`.
	TimeoutSec interface{}
}
