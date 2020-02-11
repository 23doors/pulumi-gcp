// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package compute

import (
	"reflect"

	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// > This content is derived from https://github.com/terraform-providers/terraform-provider-google/blob/master/website/docs/r/compute_external_vpn_gateway.html.markdown.
type ExternalVpnGateway struct {
	pulumi.CustomResourceState

	// An optional description of this resource.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// A list of interfaces on this external VPN gateway.
	Interfaces ExternalVpnGatewayInterfaceArrayOutput `pulumi:"interfaces"`
	// Name of the resource. Provided by the client when the resource is created. The name must be 1-63 characters long, and
	// comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression
	// '[a-z]([-a-z0-9]*[a-z0-9])?' which means the first character must be a lowercase letter, and all following characters
	// must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
	Name pulumi.StringOutput `pulumi:"name"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringOutput `pulumi:"project"`
	// Indicates the redundancy type of this external VPN gateway
	RedundancyType pulumi.StringPtrOutput `pulumi:"redundancyType"`
	// The URI of the created resource.
	SelfLink pulumi.StringOutput `pulumi:"selfLink"`
}

// NewExternalVpnGateway registers a new resource with the given unique name, arguments, and options.
func NewExternalVpnGateway(ctx *pulumi.Context,
	name string, args *ExternalVpnGatewayArgs, opts ...pulumi.ResourceOption) (*ExternalVpnGateway, error) {
	if args == nil {
		args = &ExternalVpnGatewayArgs{}
	}
	var resource ExternalVpnGateway
	err := ctx.RegisterResource("gcp:compute/externalVpnGateway:ExternalVpnGateway", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetExternalVpnGateway gets an existing ExternalVpnGateway resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetExternalVpnGateway(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ExternalVpnGatewayState, opts ...pulumi.ResourceOption) (*ExternalVpnGateway, error) {
	var resource ExternalVpnGateway
	err := ctx.ReadResource("gcp:compute/externalVpnGateway:ExternalVpnGateway", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ExternalVpnGateway resources.
type externalVpnGatewayState struct {
	// An optional description of this resource.
	Description *string `pulumi:"description"`
	// A list of interfaces on this external VPN gateway.
	Interfaces []ExternalVpnGatewayInterface `pulumi:"interfaces"`
	// Name of the resource. Provided by the client when the resource is created. The name must be 1-63 characters long, and
	// comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression
	// '[a-z]([-a-z0-9]*[a-z0-9])?' which means the first character must be a lowercase letter, and all following characters
	// must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
	Name *string `pulumi:"name"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `pulumi:"project"`
	// Indicates the redundancy type of this external VPN gateway
	RedundancyType *string `pulumi:"redundancyType"`
	// The URI of the created resource.
	SelfLink *string `pulumi:"selfLink"`
}

type ExternalVpnGatewayState struct {
	// An optional description of this resource.
	Description pulumi.StringPtrInput
	// A list of interfaces on this external VPN gateway.
	Interfaces ExternalVpnGatewayInterfaceArrayInput
	// Name of the resource. Provided by the client when the resource is created. The name must be 1-63 characters long, and
	// comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression
	// '[a-z]([-a-z0-9]*[a-z0-9])?' which means the first character must be a lowercase letter, and all following characters
	// must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
	Name pulumi.StringPtrInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringPtrInput
	// Indicates the redundancy type of this external VPN gateway
	RedundancyType pulumi.StringPtrInput
	// The URI of the created resource.
	SelfLink pulumi.StringPtrInput
}

func (ExternalVpnGatewayState) ElementType() reflect.Type {
	return reflect.TypeOf((*externalVpnGatewayState)(nil)).Elem()
}

type externalVpnGatewayArgs struct {
	// An optional description of this resource.
	Description *string `pulumi:"description"`
	// A list of interfaces on this external VPN gateway.
	Interfaces []ExternalVpnGatewayInterface `pulumi:"interfaces"`
	// Name of the resource. Provided by the client when the resource is created. The name must be 1-63 characters long, and
	// comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression
	// '[a-z]([-a-z0-9]*[a-z0-9])?' which means the first character must be a lowercase letter, and all following characters
	// must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
	Name *string `pulumi:"name"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `pulumi:"project"`
	// Indicates the redundancy type of this external VPN gateway
	RedundancyType *string `pulumi:"redundancyType"`
}

// The set of arguments for constructing a ExternalVpnGateway resource.
type ExternalVpnGatewayArgs struct {
	// An optional description of this resource.
	Description pulumi.StringPtrInput
	// A list of interfaces on this external VPN gateway.
	Interfaces ExternalVpnGatewayInterfaceArrayInput
	// Name of the resource. Provided by the client when the resource is created. The name must be 1-63 characters long, and
	// comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression
	// '[a-z]([-a-z0-9]*[a-z0-9])?' which means the first character must be a lowercase letter, and all following characters
	// must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
	Name pulumi.StringPtrInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringPtrInput
	// Indicates the redundancy type of this external VPN gateway
	RedundancyType pulumi.StringPtrInput
}

func (ExternalVpnGatewayArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*externalVpnGatewayArgs)(nil)).Elem()
}

