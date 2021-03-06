// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package bigquery

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// > This content is derived from https://github.com/terraform-providers/terraform-provider-google/blob/master/website/docs/r/bigtable_app_profile.html.markdown.
type AppProfile struct {
	pulumi.CustomResourceState

	// The unique name of the app profile in the form '[_a-zA-Z0-9][-_.a-zA-Z0-9]*'.
	AppProfileId pulumi.StringOutput `pulumi:"appProfileId"`
	// Long form description of the use case for this app profile.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// If true, ignore safety checks when deleting/updating the app profile.
	IgnoreWarnings pulumi.BoolPtrOutput `pulumi:"ignoreWarnings"`
	// The name of the instance to create the app profile within.
	Instance pulumi.StringPtrOutput `pulumi:"instance"`
	// If true, read/write requests are routed to the nearest cluster in the instance, and will fail over to the nearest
	// cluster that is available in the event of transient errors or delays. Clusters in a region are considered equidistant.
	// Choosing this option sacrifices read-your-writes consistency to improve availability.
	MultiClusterRoutingUseAny pulumi.BoolPtrOutput `pulumi:"multiClusterRoutingUseAny"`
	// The unique name of the requested app profile. Values are of the form
	// 'projects/<project>/instances/<instance>/appProfiles/<appProfileId>'.
	Name pulumi.StringOutput `pulumi:"name"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringOutput `pulumi:"project"`
	// Use a single-cluster routing policy.
	SingleClusterRouting AppProfileSingleClusterRoutingPtrOutput `pulumi:"singleClusterRouting"`
}

// NewAppProfile registers a new resource with the given unique name, arguments, and options.
func NewAppProfile(ctx *pulumi.Context,
	name string, args *AppProfileArgs, opts ...pulumi.ResourceOption) (*AppProfile, error) {
	if args == nil || args.AppProfileId == nil {
		return nil, errors.New("missing required argument 'AppProfileId'")
	}
	if args == nil {
		args = &AppProfileArgs{}
	}
	var resource AppProfile
	err := ctx.RegisterResource("gcp:bigquery/appProfile:AppProfile", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAppProfile gets an existing AppProfile resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAppProfile(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AppProfileState, opts ...pulumi.ResourceOption) (*AppProfile, error) {
	var resource AppProfile
	err := ctx.ReadResource("gcp:bigquery/appProfile:AppProfile", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering AppProfile resources.
type appProfileState struct {
	// The unique name of the app profile in the form '[_a-zA-Z0-9][-_.a-zA-Z0-9]*'.
	AppProfileId *string `pulumi:"appProfileId"`
	// Long form description of the use case for this app profile.
	Description *string `pulumi:"description"`
	// If true, ignore safety checks when deleting/updating the app profile.
	IgnoreWarnings *bool `pulumi:"ignoreWarnings"`
	// The name of the instance to create the app profile within.
	Instance *string `pulumi:"instance"`
	// If true, read/write requests are routed to the nearest cluster in the instance, and will fail over to the nearest
	// cluster that is available in the event of transient errors or delays. Clusters in a region are considered equidistant.
	// Choosing this option sacrifices read-your-writes consistency to improve availability.
	MultiClusterRoutingUseAny *bool `pulumi:"multiClusterRoutingUseAny"`
	// The unique name of the requested app profile. Values are of the form
	// 'projects/<project>/instances/<instance>/appProfiles/<appProfileId>'.
	Name *string `pulumi:"name"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `pulumi:"project"`
	// Use a single-cluster routing policy.
	SingleClusterRouting *AppProfileSingleClusterRouting `pulumi:"singleClusterRouting"`
}

type AppProfileState struct {
	// The unique name of the app profile in the form '[_a-zA-Z0-9][-_.a-zA-Z0-9]*'.
	AppProfileId pulumi.StringPtrInput
	// Long form description of the use case for this app profile.
	Description pulumi.StringPtrInput
	// If true, ignore safety checks when deleting/updating the app profile.
	IgnoreWarnings pulumi.BoolPtrInput
	// The name of the instance to create the app profile within.
	Instance pulumi.StringPtrInput
	// If true, read/write requests are routed to the nearest cluster in the instance, and will fail over to the nearest
	// cluster that is available in the event of transient errors or delays. Clusters in a region are considered equidistant.
	// Choosing this option sacrifices read-your-writes consistency to improve availability.
	MultiClusterRoutingUseAny pulumi.BoolPtrInput
	// The unique name of the requested app profile. Values are of the form
	// 'projects/<project>/instances/<instance>/appProfiles/<appProfileId>'.
	Name pulumi.StringPtrInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringPtrInput
	// Use a single-cluster routing policy.
	SingleClusterRouting AppProfileSingleClusterRoutingPtrInput
}

func (AppProfileState) ElementType() reflect.Type {
	return reflect.TypeOf((*appProfileState)(nil)).Elem()
}

type appProfileArgs struct {
	// The unique name of the app profile in the form '[_a-zA-Z0-9][-_.a-zA-Z0-9]*'.
	AppProfileId string `pulumi:"appProfileId"`
	// Long form description of the use case for this app profile.
	Description *string `pulumi:"description"`
	// If true, ignore safety checks when deleting/updating the app profile.
	IgnoreWarnings *bool `pulumi:"ignoreWarnings"`
	// The name of the instance to create the app profile within.
	Instance *string `pulumi:"instance"`
	// If true, read/write requests are routed to the nearest cluster in the instance, and will fail over to the nearest
	// cluster that is available in the event of transient errors or delays. Clusters in a region are considered equidistant.
	// Choosing this option sacrifices read-your-writes consistency to improve availability.
	MultiClusterRoutingUseAny *bool `pulumi:"multiClusterRoutingUseAny"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `pulumi:"project"`
	// Use a single-cluster routing policy.
	SingleClusterRouting *AppProfileSingleClusterRouting `pulumi:"singleClusterRouting"`
}

// The set of arguments for constructing a AppProfile resource.
type AppProfileArgs struct {
	// The unique name of the app profile in the form '[_a-zA-Z0-9][-_.a-zA-Z0-9]*'.
	AppProfileId pulumi.StringInput
	// Long form description of the use case for this app profile.
	Description pulumi.StringPtrInput
	// If true, ignore safety checks when deleting/updating the app profile.
	IgnoreWarnings pulumi.BoolPtrInput
	// The name of the instance to create the app profile within.
	Instance pulumi.StringPtrInput
	// If true, read/write requests are routed to the nearest cluster in the instance, and will fail over to the nearest
	// cluster that is available in the event of transient errors or delays. Clusters in a region are considered equidistant.
	// Choosing this option sacrifices read-your-writes consistency to improve availability.
	MultiClusterRoutingUseAny pulumi.BoolPtrInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringPtrInput
	// Use a single-cluster routing policy.
	SingleClusterRouting AppProfileSingleClusterRoutingPtrInput
}

func (AppProfileArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*appProfileArgs)(nil)).Elem()
}

