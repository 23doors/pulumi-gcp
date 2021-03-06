// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package compute

import (
	"reflect"

	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// > This content is derived from https://github.com/terraform-providers/terraform-provider-google/blob/master/website/docs/r/compute_forwarding_rule.html.markdown.
type ForwardingRule struct {
	pulumi.CustomResourceState

	// For internal TCP/UDP load balancing (i.e. load balancing scheme is INTERNAL and protocol is TCP/UDP), set this to true
	// to allow packets addressed to any ports to be forwarded to the backends configured with this forwarding rule. Used with
	// backend service. Cannot be set if port or portRange are set.
	AllPorts pulumi.BoolPtrOutput `pulumi:"allPorts"`
	// If true, clients can access ILB from all regions. Otherwise only allows from the local region the ILB is located at.
	AllowGlobalAccess pulumi.BoolPtrOutput `pulumi:"allowGlobalAccess"`
	// A BackendService to receive the matched traffic. This is used only for INTERNAL load balancing.
	BackendService pulumi.StringPtrOutput `pulumi:"backendService"`
	// Creation timestamp in RFC3339 text format.
	CreationTimestamp pulumi.StringOutput `pulumi:"creationTimestamp"`
	// An optional description of this resource. Provide this property when you create the resource.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// The IP address that this forwarding rule is serving on behalf of. Addresses are restricted based on the forwarding
	// rule's load balancing scheme (EXTERNAL or INTERNAL) and scope (global or regional). When the load balancing scheme is
	// EXTERNAL, for global forwarding rules, the address must be a global IP, and for regional forwarding rules, the address
	// must live in the same region as the forwarding rule. If this field is empty, an ephemeral IPv4 address from the same
	// scope (global or regional) will be assigned. A regional forwarding rule supports IPv4 only. A global forwarding rule
	// supports either IPv4 or IPv6. When the load balancing scheme is INTERNAL, this can only be an RFC 1918 IP address
	// belonging to the network/subnet configured for the forwarding rule. By default, if this field is empty, an ephemeral
	// internal IP address will be automatically allocated from the IP range of the subnet or network configured for this
	// forwarding rule. An address must be specified by a literal IP address. ~> **NOTE**: While the API allows you to specify
	// various resource paths for an address resource instead, Terraform requires this to specifically be an IP address to
	// avoid needing to fetching the IP address from resource paths on refresh or unnecessary diffs.
	IpAddress pulumi.StringOutput `pulumi:"ipAddress"`
	// The IP protocol to which this rule applies. Valid options are TCP, UDP, ESP, AH, SCTP or ICMP. When the load balancing
	// scheme is INTERNAL, only TCP and UDP are valid.
	IpProtocol pulumi.StringOutput `pulumi:"ipProtocol"`
	// Indicates whether or not this load balancer can be used as a collector for packet mirroring. To prevent mirroring loops,
	// instances behind this load balancer will not have their traffic mirrored even if a PacketMirroring rule applies to them.
	// This can only be set to true for load balancers that have their loadBalancingScheme set to INTERNAL.
	IsMirroringCollector pulumi.BoolPtrOutput `pulumi:"isMirroringCollector"`
	// The fingerprint used for optimistic locking of this resource. Used internally during updates.
	LabelFingerprint pulumi.StringOutput `pulumi:"labelFingerprint"`
	// Labels to apply to this forwarding rule. A list of key->value pairs.
	Labels pulumi.StringMapOutput `pulumi:"labels"`
	// This signifies what the ForwardingRule will be used for and can be EXTERNAL, INTERNAL, or INTERNAL_MANAGED. EXTERNAL is
	// used for Classic Cloud VPN gateways, protocol forwarding to VMs from an external IP address, and HTTP(S), SSL Proxy, TCP
	// Proxy, and Network TCP/UDP load balancers. INTERNAL is used for protocol forwarding to VMs from an internal IP address,
	// and internal TCP/UDP load balancers. INTERNAL_MANAGED is used for internal HTTP(S) load balancers.
	LoadBalancingScheme pulumi.StringPtrOutput `pulumi:"loadBalancingScheme"`
	// Name of the resource; provided by the client when the resource is created. The name must be 1-63 characters long, and
	// comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression
	// '[a-z]([-a-z0-9]*[a-z0-9])?' which means the first character must be a lowercase letter, and all following characters
	// must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
	Name pulumi.StringOutput `pulumi:"name"`
	// For internal load balancing, this field identifies the network that the load balanced IP should belong to for this
	// Forwarding Rule. If this field is not specified, the default network will be used. This field is only used for INTERNAL
	// load balancing.
	Network pulumi.StringOutput `pulumi:"network"`
	// The networking tier used for configuring this address. This field can take the following values: PREMIUM or STANDARD. If
	// this field is not specified, it is assumed to be PREMIUM.
	NetworkTier pulumi.StringOutput `pulumi:"networkTier"`
	// This field is used along with the target field for TargetHttpProxy, TargetHttpsProxy, TargetSslProxy, TargetTcpProxy,
	// TargetVpnGateway, TargetPool, TargetInstance. Applicable only when IPProtocol is TCP, UDP, or SCTP, only packets
	// addressed to ports in the specified range will be forwarded to target. Forwarding rules with the same [IPAddress,
	// IPProtocol] pair must have disjoint port ranges. Some types of forwarding target have constraints on the acceptable
	// ports: * TargetHttpProxy: 80, 8080 * TargetHttpsProxy: 443 * TargetTcpProxy: 25, 43, 110, 143, 195, 443, 465, 587, 700,
	// 993, 995, 1883, 5222 * TargetSslProxy: 25, 43, 110, 143, 195, 443, 465, 587, 700, 993, 995, 1883, 5222 *
	// TargetVpnGateway: 500, 4500
	PortRange pulumi.StringPtrOutput `pulumi:"portRange"`
	// This field is used along with the backend_service field for internal load balancing. When the load balancing scheme is
	// INTERNAL, a single port or a comma separated list of ports can be configured. Only packets addressed to these ports will
	// be forwarded to the backends configured with this forwarding rule. You may specify a maximum of up to 5 ports.
	Ports pulumi.StringArrayOutput `pulumi:"ports"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringOutput `pulumi:"project"`
	// A reference to the region where the regional forwarding rule resides. This field is not applicable to global forwarding
	// rules.
	Region pulumi.StringOutput `pulumi:"region"`
	// The URI of the created resource.
	SelfLink pulumi.StringOutput `pulumi:"selfLink"`
	// An optional prefix to the service name for this Forwarding Rule. If specified, will be the first label of the fully
	// qualified service name. The label must be 1-63 characters long, and comply with RFC1035. Specifically, the label must be
	// 1-63 characters long and match the regular expression '[a-z]([-a-z0-9]*[a-z0-9])?' which means the first character must
	// be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last
	// character, which cannot be a dash. This field is only used for INTERNAL load balancing.
	ServiceLabel pulumi.StringPtrOutput `pulumi:"serviceLabel"`
	// The internal fully qualified service name for this Forwarding Rule. This field is only used for INTERNAL load balancing.
	ServiceName pulumi.StringOutput `pulumi:"serviceName"`
	// The subnetwork that the load balanced IP should belong to for this Forwarding Rule. This field is only used for INTERNAL
	// load balancing. If the network specified is in auto subnet mode, this field is optional. However, if the network is in
	// custom subnet mode, a subnetwork must be specified.
	Subnetwork pulumi.StringOutput `pulumi:"subnetwork"`
	// The URL of the target resource to receive the matched traffic. The target must live in the same region as the forwarding
	// rule. The forwarded traffic must be of a type appropriate to the target object.
	Target pulumi.StringPtrOutput `pulumi:"target"`
}

// NewForwardingRule registers a new resource with the given unique name, arguments, and options.
func NewForwardingRule(ctx *pulumi.Context,
	name string, args *ForwardingRuleArgs, opts ...pulumi.ResourceOption) (*ForwardingRule, error) {
	if args == nil {
		args = &ForwardingRuleArgs{}
	}
	var resource ForwardingRule
	err := ctx.RegisterResource("gcp:compute/forwardingRule:ForwardingRule", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetForwardingRule gets an existing ForwardingRule resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetForwardingRule(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ForwardingRuleState, opts ...pulumi.ResourceOption) (*ForwardingRule, error) {
	var resource ForwardingRule
	err := ctx.ReadResource("gcp:compute/forwardingRule:ForwardingRule", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ForwardingRule resources.
type forwardingRuleState struct {
	// For internal TCP/UDP load balancing (i.e. load balancing scheme is INTERNAL and protocol is TCP/UDP), set this to true
	// to allow packets addressed to any ports to be forwarded to the backends configured with this forwarding rule. Used with
	// backend service. Cannot be set if port or portRange are set.
	AllPorts *bool `pulumi:"allPorts"`
	// If true, clients can access ILB from all regions. Otherwise only allows from the local region the ILB is located at.
	AllowGlobalAccess *bool `pulumi:"allowGlobalAccess"`
	// A BackendService to receive the matched traffic. This is used only for INTERNAL load balancing.
	BackendService *string `pulumi:"backendService"`
	// Creation timestamp in RFC3339 text format.
	CreationTimestamp *string `pulumi:"creationTimestamp"`
	// An optional description of this resource. Provide this property when you create the resource.
	Description *string `pulumi:"description"`
	// The IP address that this forwarding rule is serving on behalf of. Addresses are restricted based on the forwarding
	// rule's load balancing scheme (EXTERNAL or INTERNAL) and scope (global or regional). When the load balancing scheme is
	// EXTERNAL, for global forwarding rules, the address must be a global IP, and for regional forwarding rules, the address
	// must live in the same region as the forwarding rule. If this field is empty, an ephemeral IPv4 address from the same
	// scope (global or regional) will be assigned. A regional forwarding rule supports IPv4 only. A global forwarding rule
	// supports either IPv4 or IPv6. When the load balancing scheme is INTERNAL, this can only be an RFC 1918 IP address
	// belonging to the network/subnet configured for the forwarding rule. By default, if this field is empty, an ephemeral
	// internal IP address will be automatically allocated from the IP range of the subnet or network configured for this
	// forwarding rule. An address must be specified by a literal IP address. ~> **NOTE**: While the API allows you to specify
	// various resource paths for an address resource instead, Terraform requires this to specifically be an IP address to
	// avoid needing to fetching the IP address from resource paths on refresh or unnecessary diffs.
	IpAddress *string `pulumi:"ipAddress"`
	// The IP protocol to which this rule applies. Valid options are TCP, UDP, ESP, AH, SCTP or ICMP. When the load balancing
	// scheme is INTERNAL, only TCP and UDP are valid.
	IpProtocol *string `pulumi:"ipProtocol"`
	// Indicates whether or not this load balancer can be used as a collector for packet mirroring. To prevent mirroring loops,
	// instances behind this load balancer will not have their traffic mirrored even if a PacketMirroring rule applies to them.
	// This can only be set to true for load balancers that have their loadBalancingScheme set to INTERNAL.
	IsMirroringCollector *bool `pulumi:"isMirroringCollector"`
	// The fingerprint used for optimistic locking of this resource. Used internally during updates.
	LabelFingerprint *string `pulumi:"labelFingerprint"`
	// Labels to apply to this forwarding rule. A list of key->value pairs.
	Labels map[string]string `pulumi:"labels"`
	// This signifies what the ForwardingRule will be used for and can be EXTERNAL, INTERNAL, or INTERNAL_MANAGED. EXTERNAL is
	// used for Classic Cloud VPN gateways, protocol forwarding to VMs from an external IP address, and HTTP(S), SSL Proxy, TCP
	// Proxy, and Network TCP/UDP load balancers. INTERNAL is used for protocol forwarding to VMs from an internal IP address,
	// and internal TCP/UDP load balancers. INTERNAL_MANAGED is used for internal HTTP(S) load balancers.
	LoadBalancingScheme *string `pulumi:"loadBalancingScheme"`
	// Name of the resource; provided by the client when the resource is created. The name must be 1-63 characters long, and
	// comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression
	// '[a-z]([-a-z0-9]*[a-z0-9])?' which means the first character must be a lowercase letter, and all following characters
	// must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
	Name *string `pulumi:"name"`
	// For internal load balancing, this field identifies the network that the load balanced IP should belong to for this
	// Forwarding Rule. If this field is not specified, the default network will be used. This field is only used for INTERNAL
	// load balancing.
	Network *string `pulumi:"network"`
	// The networking tier used for configuring this address. This field can take the following values: PREMIUM or STANDARD. If
	// this field is not specified, it is assumed to be PREMIUM.
	NetworkTier *string `pulumi:"networkTier"`
	// This field is used along with the target field for TargetHttpProxy, TargetHttpsProxy, TargetSslProxy, TargetTcpProxy,
	// TargetVpnGateway, TargetPool, TargetInstance. Applicable only when IPProtocol is TCP, UDP, or SCTP, only packets
	// addressed to ports in the specified range will be forwarded to target. Forwarding rules with the same [IPAddress,
	// IPProtocol] pair must have disjoint port ranges. Some types of forwarding target have constraints on the acceptable
	// ports: * TargetHttpProxy: 80, 8080 * TargetHttpsProxy: 443 * TargetTcpProxy: 25, 43, 110, 143, 195, 443, 465, 587, 700,
	// 993, 995, 1883, 5222 * TargetSslProxy: 25, 43, 110, 143, 195, 443, 465, 587, 700, 993, 995, 1883, 5222 *
	// TargetVpnGateway: 500, 4500
	PortRange *string `pulumi:"portRange"`
	// This field is used along with the backend_service field for internal load balancing. When the load balancing scheme is
	// INTERNAL, a single port or a comma separated list of ports can be configured. Only packets addressed to these ports will
	// be forwarded to the backends configured with this forwarding rule. You may specify a maximum of up to 5 ports.
	Ports []string `pulumi:"ports"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `pulumi:"project"`
	// A reference to the region where the regional forwarding rule resides. This field is not applicable to global forwarding
	// rules.
	Region *string `pulumi:"region"`
	// The URI of the created resource.
	SelfLink *string `pulumi:"selfLink"`
	// An optional prefix to the service name for this Forwarding Rule. If specified, will be the first label of the fully
	// qualified service name. The label must be 1-63 characters long, and comply with RFC1035. Specifically, the label must be
	// 1-63 characters long and match the regular expression '[a-z]([-a-z0-9]*[a-z0-9])?' which means the first character must
	// be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last
	// character, which cannot be a dash. This field is only used for INTERNAL load balancing.
	ServiceLabel *string `pulumi:"serviceLabel"`
	// The internal fully qualified service name for this Forwarding Rule. This field is only used for INTERNAL load balancing.
	ServiceName *string `pulumi:"serviceName"`
	// The subnetwork that the load balanced IP should belong to for this Forwarding Rule. This field is only used for INTERNAL
	// load balancing. If the network specified is in auto subnet mode, this field is optional. However, if the network is in
	// custom subnet mode, a subnetwork must be specified.
	Subnetwork *string `pulumi:"subnetwork"`
	// The URL of the target resource to receive the matched traffic. The target must live in the same region as the forwarding
	// rule. The forwarded traffic must be of a type appropriate to the target object.
	Target *string `pulumi:"target"`
}

type ForwardingRuleState struct {
	// For internal TCP/UDP load balancing (i.e. load balancing scheme is INTERNAL and protocol is TCP/UDP), set this to true
	// to allow packets addressed to any ports to be forwarded to the backends configured with this forwarding rule. Used with
	// backend service. Cannot be set if port or portRange are set.
	AllPorts pulumi.BoolPtrInput
	// If true, clients can access ILB from all regions. Otherwise only allows from the local region the ILB is located at.
	AllowGlobalAccess pulumi.BoolPtrInput
	// A BackendService to receive the matched traffic. This is used only for INTERNAL load balancing.
	BackendService pulumi.StringPtrInput
	// Creation timestamp in RFC3339 text format.
	CreationTimestamp pulumi.StringPtrInput
	// An optional description of this resource. Provide this property when you create the resource.
	Description pulumi.StringPtrInput
	// The IP address that this forwarding rule is serving on behalf of. Addresses are restricted based on the forwarding
	// rule's load balancing scheme (EXTERNAL or INTERNAL) and scope (global or regional). When the load balancing scheme is
	// EXTERNAL, for global forwarding rules, the address must be a global IP, and for regional forwarding rules, the address
	// must live in the same region as the forwarding rule. If this field is empty, an ephemeral IPv4 address from the same
	// scope (global or regional) will be assigned. A regional forwarding rule supports IPv4 only. A global forwarding rule
	// supports either IPv4 or IPv6. When the load balancing scheme is INTERNAL, this can only be an RFC 1918 IP address
	// belonging to the network/subnet configured for the forwarding rule. By default, if this field is empty, an ephemeral
	// internal IP address will be automatically allocated from the IP range of the subnet or network configured for this
	// forwarding rule. An address must be specified by a literal IP address. ~> **NOTE**: While the API allows you to specify
	// various resource paths for an address resource instead, Terraform requires this to specifically be an IP address to
	// avoid needing to fetching the IP address from resource paths on refresh or unnecessary diffs.
	IpAddress pulumi.StringPtrInput
	// The IP protocol to which this rule applies. Valid options are TCP, UDP, ESP, AH, SCTP or ICMP. When the load balancing
	// scheme is INTERNAL, only TCP and UDP are valid.
	IpProtocol pulumi.StringPtrInput
	// Indicates whether or not this load balancer can be used as a collector for packet mirroring. To prevent mirroring loops,
	// instances behind this load balancer will not have their traffic mirrored even if a PacketMirroring rule applies to them.
	// This can only be set to true for load balancers that have their loadBalancingScheme set to INTERNAL.
	IsMirroringCollector pulumi.BoolPtrInput
	// The fingerprint used for optimistic locking of this resource. Used internally during updates.
	LabelFingerprint pulumi.StringPtrInput
	// Labels to apply to this forwarding rule. A list of key->value pairs.
	Labels pulumi.StringMapInput
	// This signifies what the ForwardingRule will be used for and can be EXTERNAL, INTERNAL, or INTERNAL_MANAGED. EXTERNAL is
	// used for Classic Cloud VPN gateways, protocol forwarding to VMs from an external IP address, and HTTP(S), SSL Proxy, TCP
	// Proxy, and Network TCP/UDP load balancers. INTERNAL is used for protocol forwarding to VMs from an internal IP address,
	// and internal TCP/UDP load balancers. INTERNAL_MANAGED is used for internal HTTP(S) load balancers.
	LoadBalancingScheme pulumi.StringPtrInput
	// Name of the resource; provided by the client when the resource is created. The name must be 1-63 characters long, and
	// comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression
	// '[a-z]([-a-z0-9]*[a-z0-9])?' which means the first character must be a lowercase letter, and all following characters
	// must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
	Name pulumi.StringPtrInput
	// For internal load balancing, this field identifies the network that the load balanced IP should belong to for this
	// Forwarding Rule. If this field is not specified, the default network will be used. This field is only used for INTERNAL
	// load balancing.
	Network pulumi.StringPtrInput
	// The networking tier used for configuring this address. This field can take the following values: PREMIUM or STANDARD. If
	// this field is not specified, it is assumed to be PREMIUM.
	NetworkTier pulumi.StringPtrInput
	// This field is used along with the target field for TargetHttpProxy, TargetHttpsProxy, TargetSslProxy, TargetTcpProxy,
	// TargetVpnGateway, TargetPool, TargetInstance. Applicable only when IPProtocol is TCP, UDP, or SCTP, only packets
	// addressed to ports in the specified range will be forwarded to target. Forwarding rules with the same [IPAddress,
	// IPProtocol] pair must have disjoint port ranges. Some types of forwarding target have constraints on the acceptable
	// ports: * TargetHttpProxy: 80, 8080 * TargetHttpsProxy: 443 * TargetTcpProxy: 25, 43, 110, 143, 195, 443, 465, 587, 700,
	// 993, 995, 1883, 5222 * TargetSslProxy: 25, 43, 110, 143, 195, 443, 465, 587, 700, 993, 995, 1883, 5222 *
	// TargetVpnGateway: 500, 4500
	PortRange pulumi.StringPtrInput
	// This field is used along with the backend_service field for internal load balancing. When the load balancing scheme is
	// INTERNAL, a single port or a comma separated list of ports can be configured. Only packets addressed to these ports will
	// be forwarded to the backends configured with this forwarding rule. You may specify a maximum of up to 5 ports.
	Ports pulumi.StringArrayInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringPtrInput
	// A reference to the region where the regional forwarding rule resides. This field is not applicable to global forwarding
	// rules.
	Region pulumi.StringPtrInput
	// The URI of the created resource.
	SelfLink pulumi.StringPtrInput
	// An optional prefix to the service name for this Forwarding Rule. If specified, will be the first label of the fully
	// qualified service name. The label must be 1-63 characters long, and comply with RFC1035. Specifically, the label must be
	// 1-63 characters long and match the regular expression '[a-z]([-a-z0-9]*[a-z0-9])?' which means the first character must
	// be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last
	// character, which cannot be a dash. This field is only used for INTERNAL load balancing.
	ServiceLabel pulumi.StringPtrInput
	// The internal fully qualified service name for this Forwarding Rule. This field is only used for INTERNAL load balancing.
	ServiceName pulumi.StringPtrInput
	// The subnetwork that the load balanced IP should belong to for this Forwarding Rule. This field is only used for INTERNAL
	// load balancing. If the network specified is in auto subnet mode, this field is optional. However, if the network is in
	// custom subnet mode, a subnetwork must be specified.
	Subnetwork pulumi.StringPtrInput
	// The URL of the target resource to receive the matched traffic. The target must live in the same region as the forwarding
	// rule. The forwarded traffic must be of a type appropriate to the target object.
	Target pulumi.StringPtrInput
}

func (ForwardingRuleState) ElementType() reflect.Type {
	return reflect.TypeOf((*forwardingRuleState)(nil)).Elem()
}

type forwardingRuleArgs struct {
	// For internal TCP/UDP load balancing (i.e. load balancing scheme is INTERNAL and protocol is TCP/UDP), set this to true
	// to allow packets addressed to any ports to be forwarded to the backends configured with this forwarding rule. Used with
	// backend service. Cannot be set if port or portRange are set.
	AllPorts *bool `pulumi:"allPorts"`
	// If true, clients can access ILB from all regions. Otherwise only allows from the local region the ILB is located at.
	AllowGlobalAccess *bool `pulumi:"allowGlobalAccess"`
	// A BackendService to receive the matched traffic. This is used only for INTERNAL load balancing.
	BackendService *string `pulumi:"backendService"`
	// An optional description of this resource. Provide this property when you create the resource.
	Description *string `pulumi:"description"`
	// The IP address that this forwarding rule is serving on behalf of. Addresses are restricted based on the forwarding
	// rule's load balancing scheme (EXTERNAL or INTERNAL) and scope (global or regional). When the load balancing scheme is
	// EXTERNAL, for global forwarding rules, the address must be a global IP, and for regional forwarding rules, the address
	// must live in the same region as the forwarding rule. If this field is empty, an ephemeral IPv4 address from the same
	// scope (global or regional) will be assigned. A regional forwarding rule supports IPv4 only. A global forwarding rule
	// supports either IPv4 or IPv6. When the load balancing scheme is INTERNAL, this can only be an RFC 1918 IP address
	// belonging to the network/subnet configured for the forwarding rule. By default, if this field is empty, an ephemeral
	// internal IP address will be automatically allocated from the IP range of the subnet or network configured for this
	// forwarding rule. An address must be specified by a literal IP address. ~> **NOTE**: While the API allows you to specify
	// various resource paths for an address resource instead, Terraform requires this to specifically be an IP address to
	// avoid needing to fetching the IP address from resource paths on refresh or unnecessary diffs.
	IpAddress *string `pulumi:"ipAddress"`
	// The IP protocol to which this rule applies. Valid options are TCP, UDP, ESP, AH, SCTP or ICMP. When the load balancing
	// scheme is INTERNAL, only TCP and UDP are valid.
	IpProtocol *string `pulumi:"ipProtocol"`
	// Indicates whether or not this load balancer can be used as a collector for packet mirroring. To prevent mirroring loops,
	// instances behind this load balancer will not have their traffic mirrored even if a PacketMirroring rule applies to them.
	// This can only be set to true for load balancers that have their loadBalancingScheme set to INTERNAL.
	IsMirroringCollector *bool `pulumi:"isMirroringCollector"`
	// Labels to apply to this forwarding rule. A list of key->value pairs.
	Labels map[string]string `pulumi:"labels"`
	// This signifies what the ForwardingRule will be used for and can be EXTERNAL, INTERNAL, or INTERNAL_MANAGED. EXTERNAL is
	// used for Classic Cloud VPN gateways, protocol forwarding to VMs from an external IP address, and HTTP(S), SSL Proxy, TCP
	// Proxy, and Network TCP/UDP load balancers. INTERNAL is used for protocol forwarding to VMs from an internal IP address,
	// and internal TCP/UDP load balancers. INTERNAL_MANAGED is used for internal HTTP(S) load balancers.
	LoadBalancingScheme *string `pulumi:"loadBalancingScheme"`
	// Name of the resource; provided by the client when the resource is created. The name must be 1-63 characters long, and
	// comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression
	// '[a-z]([-a-z0-9]*[a-z0-9])?' which means the first character must be a lowercase letter, and all following characters
	// must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
	Name *string `pulumi:"name"`
	// For internal load balancing, this field identifies the network that the load balanced IP should belong to for this
	// Forwarding Rule. If this field is not specified, the default network will be used. This field is only used for INTERNAL
	// load balancing.
	Network *string `pulumi:"network"`
	// The networking tier used for configuring this address. This field can take the following values: PREMIUM or STANDARD. If
	// this field is not specified, it is assumed to be PREMIUM.
	NetworkTier *string `pulumi:"networkTier"`
	// This field is used along with the target field for TargetHttpProxy, TargetHttpsProxy, TargetSslProxy, TargetTcpProxy,
	// TargetVpnGateway, TargetPool, TargetInstance. Applicable only when IPProtocol is TCP, UDP, or SCTP, only packets
	// addressed to ports in the specified range will be forwarded to target. Forwarding rules with the same [IPAddress,
	// IPProtocol] pair must have disjoint port ranges. Some types of forwarding target have constraints on the acceptable
	// ports: * TargetHttpProxy: 80, 8080 * TargetHttpsProxy: 443 * TargetTcpProxy: 25, 43, 110, 143, 195, 443, 465, 587, 700,
	// 993, 995, 1883, 5222 * TargetSslProxy: 25, 43, 110, 143, 195, 443, 465, 587, 700, 993, 995, 1883, 5222 *
	// TargetVpnGateway: 500, 4500
	PortRange *string `pulumi:"portRange"`
	// This field is used along with the backend_service field for internal load balancing. When the load balancing scheme is
	// INTERNAL, a single port or a comma separated list of ports can be configured. Only packets addressed to these ports will
	// be forwarded to the backends configured with this forwarding rule. You may specify a maximum of up to 5 ports.
	Ports []string `pulumi:"ports"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `pulumi:"project"`
	// A reference to the region where the regional forwarding rule resides. This field is not applicable to global forwarding
	// rules.
	Region *string `pulumi:"region"`
	// An optional prefix to the service name for this Forwarding Rule. If specified, will be the first label of the fully
	// qualified service name. The label must be 1-63 characters long, and comply with RFC1035. Specifically, the label must be
	// 1-63 characters long and match the regular expression '[a-z]([-a-z0-9]*[a-z0-9])?' which means the first character must
	// be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last
	// character, which cannot be a dash. This field is only used for INTERNAL load balancing.
	ServiceLabel *string `pulumi:"serviceLabel"`
	// The subnetwork that the load balanced IP should belong to for this Forwarding Rule. This field is only used for INTERNAL
	// load balancing. If the network specified is in auto subnet mode, this field is optional. However, if the network is in
	// custom subnet mode, a subnetwork must be specified.
	Subnetwork *string `pulumi:"subnetwork"`
	// The URL of the target resource to receive the matched traffic. The target must live in the same region as the forwarding
	// rule. The forwarded traffic must be of a type appropriate to the target object.
	Target *string `pulumi:"target"`
}

// The set of arguments for constructing a ForwardingRule resource.
type ForwardingRuleArgs struct {
	// For internal TCP/UDP load balancing (i.e. load balancing scheme is INTERNAL and protocol is TCP/UDP), set this to true
	// to allow packets addressed to any ports to be forwarded to the backends configured with this forwarding rule. Used with
	// backend service. Cannot be set if port or portRange are set.
	AllPorts pulumi.BoolPtrInput
	// If true, clients can access ILB from all regions. Otherwise only allows from the local region the ILB is located at.
	AllowGlobalAccess pulumi.BoolPtrInput
	// A BackendService to receive the matched traffic. This is used only for INTERNAL load balancing.
	BackendService pulumi.StringPtrInput
	// An optional description of this resource. Provide this property when you create the resource.
	Description pulumi.StringPtrInput
	// The IP address that this forwarding rule is serving on behalf of. Addresses are restricted based on the forwarding
	// rule's load balancing scheme (EXTERNAL or INTERNAL) and scope (global or regional). When the load balancing scheme is
	// EXTERNAL, for global forwarding rules, the address must be a global IP, and for regional forwarding rules, the address
	// must live in the same region as the forwarding rule. If this field is empty, an ephemeral IPv4 address from the same
	// scope (global or regional) will be assigned. A regional forwarding rule supports IPv4 only. A global forwarding rule
	// supports either IPv4 or IPv6. When the load balancing scheme is INTERNAL, this can only be an RFC 1918 IP address
	// belonging to the network/subnet configured for the forwarding rule. By default, if this field is empty, an ephemeral
	// internal IP address will be automatically allocated from the IP range of the subnet or network configured for this
	// forwarding rule. An address must be specified by a literal IP address. ~> **NOTE**: While the API allows you to specify
	// various resource paths for an address resource instead, Terraform requires this to specifically be an IP address to
	// avoid needing to fetching the IP address from resource paths on refresh or unnecessary diffs.
	IpAddress pulumi.StringPtrInput
	// The IP protocol to which this rule applies. Valid options are TCP, UDP, ESP, AH, SCTP or ICMP. When the load balancing
	// scheme is INTERNAL, only TCP and UDP are valid.
	IpProtocol pulumi.StringPtrInput
	// Indicates whether or not this load balancer can be used as a collector for packet mirroring. To prevent mirroring loops,
	// instances behind this load balancer will not have their traffic mirrored even if a PacketMirroring rule applies to them.
	// This can only be set to true for load balancers that have their loadBalancingScheme set to INTERNAL.
	IsMirroringCollector pulumi.BoolPtrInput
	// Labels to apply to this forwarding rule. A list of key->value pairs.
	Labels pulumi.StringMapInput
	// This signifies what the ForwardingRule will be used for and can be EXTERNAL, INTERNAL, or INTERNAL_MANAGED. EXTERNAL is
	// used for Classic Cloud VPN gateways, protocol forwarding to VMs from an external IP address, and HTTP(S), SSL Proxy, TCP
	// Proxy, and Network TCP/UDP load balancers. INTERNAL is used for protocol forwarding to VMs from an internal IP address,
	// and internal TCP/UDP load balancers. INTERNAL_MANAGED is used for internal HTTP(S) load balancers.
	LoadBalancingScheme pulumi.StringPtrInput
	// Name of the resource; provided by the client when the resource is created. The name must be 1-63 characters long, and
	// comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression
	// '[a-z]([-a-z0-9]*[a-z0-9])?' which means the first character must be a lowercase letter, and all following characters
	// must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
	Name pulumi.StringPtrInput
	// For internal load balancing, this field identifies the network that the load balanced IP should belong to for this
	// Forwarding Rule. If this field is not specified, the default network will be used. This field is only used for INTERNAL
	// load balancing.
	Network pulumi.StringPtrInput
	// The networking tier used for configuring this address. This field can take the following values: PREMIUM or STANDARD. If
	// this field is not specified, it is assumed to be PREMIUM.
	NetworkTier pulumi.StringPtrInput
	// This field is used along with the target field for TargetHttpProxy, TargetHttpsProxy, TargetSslProxy, TargetTcpProxy,
	// TargetVpnGateway, TargetPool, TargetInstance. Applicable only when IPProtocol is TCP, UDP, or SCTP, only packets
	// addressed to ports in the specified range will be forwarded to target. Forwarding rules with the same [IPAddress,
	// IPProtocol] pair must have disjoint port ranges. Some types of forwarding target have constraints on the acceptable
	// ports: * TargetHttpProxy: 80, 8080 * TargetHttpsProxy: 443 * TargetTcpProxy: 25, 43, 110, 143, 195, 443, 465, 587, 700,
	// 993, 995, 1883, 5222 * TargetSslProxy: 25, 43, 110, 143, 195, 443, 465, 587, 700, 993, 995, 1883, 5222 *
	// TargetVpnGateway: 500, 4500
	PortRange pulumi.StringPtrInput
	// This field is used along with the backend_service field for internal load balancing. When the load balancing scheme is
	// INTERNAL, a single port or a comma separated list of ports can be configured. Only packets addressed to these ports will
	// be forwarded to the backends configured with this forwarding rule. You may specify a maximum of up to 5 ports.
	Ports pulumi.StringArrayInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringPtrInput
	// A reference to the region where the regional forwarding rule resides. This field is not applicable to global forwarding
	// rules.
	Region pulumi.StringPtrInput
	// An optional prefix to the service name for this Forwarding Rule. If specified, will be the first label of the fully
	// qualified service name. The label must be 1-63 characters long, and comply with RFC1035. Specifically, the label must be
	// 1-63 characters long and match the regular expression '[a-z]([-a-z0-9]*[a-z0-9])?' which means the first character must
	// be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last
	// character, which cannot be a dash. This field is only used for INTERNAL load balancing.
	ServiceLabel pulumi.StringPtrInput
	// The subnetwork that the load balanced IP should belong to for this Forwarding Rule. This field is only used for INTERNAL
	// load balancing. If the network specified is in auto subnet mode, this field is optional. However, if the network is in
	// custom subnet mode, a subnetwork must be specified.
	Subnetwork pulumi.StringPtrInput
	// The URL of the target resource to receive the matched traffic. The target must live in the same region as the forwarding
	// rule. The forwarded traffic must be of a type appropriate to the target object.
	Target pulumi.StringPtrInput
}

func (ForwardingRuleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*forwardingRuleArgs)(nil)).Elem()
}

