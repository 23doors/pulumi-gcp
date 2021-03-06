// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package organizations

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Allows management of the entire IAM policy for an existing Google Cloud Platform Organization.
// 
// !> **Warning:** New organizations have several default policies which will,
//    without extreme caution, be **overwritten** by use of this resource.
//    The safest alternative is to use multiple `organizations.IAMBinding`
//    resources.  It is easy to use this resource to remove your own access to
//    an organization, which will require a call to Google Support to have
//    fixed, and can take multiple days to resolve.  If you do use this resource,
//    the best way to be sure that you are not making dangerous changes is to start
//    by importing your existing policy, and examining the diff very closely.
// 
// > **Note:** This resource __must not__ be used in conjunction with
//    `organizations.IAMMember` or `organizations.IAMBinding`
//    or they will fight over what your policy should be.
// 
// > This content is derived from https://github.com/terraform-providers/terraform-provider-google/blob/master/website/docs/r/organization_iam_policy.html.markdown.
type IAMPolicy struct {
	pulumi.CustomResourceState

	Etag pulumi.StringOutput `pulumi:"etag"`
	// The numeric ID of the organization in which you want to create a custom role.
	OrgId pulumi.StringOutput `pulumi:"orgId"`
	// The `organizations.getIAMPolicy` data source that represents
	// the IAM policy that will be applied to the organization. This policy overrides any existing
	// policy applied to the organization.
	PolicyData pulumi.StringOutput `pulumi:"policyData"`
}

// NewIAMPolicy registers a new resource with the given unique name, arguments, and options.
func NewIAMPolicy(ctx *pulumi.Context,
	name string, args *IAMPolicyArgs, opts ...pulumi.ResourceOption) (*IAMPolicy, error) {
	if args == nil || args.OrgId == nil {
		return nil, errors.New("missing required argument 'OrgId'")
	}
	if args == nil || args.PolicyData == nil {
		return nil, errors.New("missing required argument 'PolicyData'")
	}
	if args == nil {
		args = &IAMPolicyArgs{}
	}
	var resource IAMPolicy
	err := ctx.RegisterResource("gcp:organizations/iAMPolicy:IAMPolicy", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetIAMPolicy gets an existing IAMPolicy resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetIAMPolicy(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *IAMPolicyState, opts ...pulumi.ResourceOption) (*IAMPolicy, error) {
	var resource IAMPolicy
	err := ctx.ReadResource("gcp:organizations/iAMPolicy:IAMPolicy", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering IAMPolicy resources.
type iampolicyState struct {
	Etag *string `pulumi:"etag"`
	// The numeric ID of the organization in which you want to create a custom role.
	OrgId *string `pulumi:"orgId"`
	// The `organizations.getIAMPolicy` data source that represents
	// the IAM policy that will be applied to the organization. This policy overrides any existing
	// policy applied to the organization.
	PolicyData *string `pulumi:"policyData"`
}

type IAMPolicyState struct {
	Etag pulumi.StringPtrInput
	// The numeric ID of the organization in which you want to create a custom role.
	OrgId pulumi.StringPtrInput
	// The `organizations.getIAMPolicy` data source that represents
	// the IAM policy that will be applied to the organization. This policy overrides any existing
	// policy applied to the organization.
	PolicyData pulumi.StringPtrInput
}

func (IAMPolicyState) ElementType() reflect.Type {
	return reflect.TypeOf((*iampolicyState)(nil)).Elem()
}

type iampolicyArgs struct {
	// The numeric ID of the organization in which you want to create a custom role.
	OrgId string `pulumi:"orgId"`
	// The `organizations.getIAMPolicy` data source that represents
	// the IAM policy that will be applied to the organization. This policy overrides any existing
	// policy applied to the organization.
	PolicyData string `pulumi:"policyData"`
}

// The set of arguments for constructing a IAMPolicy resource.
type IAMPolicyArgs struct {
	// The numeric ID of the organization in which you want to create a custom role.
	OrgId pulumi.StringInput
	// The `organizations.getIAMPolicy` data source that represents
	// the IAM policy that will be applied to the organization. This policy overrides any existing
	// policy applied to the organization.
	PolicyData pulumi.StringInput
}

func (IAMPolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*iampolicyArgs)(nil)).Elem()
}

