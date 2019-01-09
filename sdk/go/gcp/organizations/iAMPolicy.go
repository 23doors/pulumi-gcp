// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package organizations

import (
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Allows management of the entire IAM policy for an existing Google Cloud Platform Organization.
// 
// > **Warning:** New organizations have several default policies which will,
//    without extreme caution, be **overwritten** by use of this resource.
//    The safest alternative is to use multiple `google_organization_iam_binding`
//    resources.  It is easy to use this resource to remove your own access to
//    an organization, which will require a call to Google Support to have
//    fixed, and can take multiple days to resolve.  If you do use this resource,
//    the best way to be sure that you are not making dangerous changes is to start
//    by importing your existing policy, and examining the diff very closely.
// 
// > **Note:** This resource __must not__ be used in conjunction with
//    `google_organization_iam_member` or `google_organization_iam_binding`
//    or they will fight over what your policy should be.
type IAMPolicy struct {
	s *pulumi.ResourceState
}

// NewIAMPolicy registers a new resource with the given unique name, arguments, and options.
func NewIAMPolicy(ctx *pulumi.Context,
	name string, args *IAMPolicyArgs, opts ...pulumi.ResourceOpt) (*IAMPolicy, error) {
	if args == nil || args.OrgId == nil {
		return nil, errors.New("missing required argument 'OrgId'")
	}
	if args == nil || args.PolicyData == nil {
		return nil, errors.New("missing required argument 'PolicyData'")
	}
	inputs := make(map[string]interface{})
	if args == nil {
		inputs["orgId"] = nil
		inputs["policyData"] = nil
	} else {
		inputs["orgId"] = args.OrgId
		inputs["policyData"] = args.PolicyData
	}
	inputs["etag"] = nil
	s, err := ctx.RegisterResource("gcp:organizations/iAMPolicy:IAMPolicy", name, true, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &IAMPolicy{s: s}, nil
}

// GetIAMPolicy gets an existing IAMPolicy resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetIAMPolicy(ctx *pulumi.Context,
	name string, id pulumi.ID, state *IAMPolicyState, opts ...pulumi.ResourceOpt) (*IAMPolicy, error) {
	inputs := make(map[string]interface{})
	if state != nil {
		inputs["etag"] = state.Etag
		inputs["orgId"] = state.OrgId
		inputs["policyData"] = state.PolicyData
	}
	s, err := ctx.ReadResource("gcp:organizations/iAMPolicy:IAMPolicy", name, id, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &IAMPolicy{s: s}, nil
}

// URN is this resource's unique name assigned by Pulumi.
func (r *IAMPolicy) URN() *pulumi.URNOutput {
	return r.s.URN()
}

// ID is this resource's unique identifier assigned by its provider.
func (r *IAMPolicy) ID() *pulumi.IDOutput {
	return r.s.ID()
}

func (r *IAMPolicy) Etag() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["etag"])
}

// The numeric ID of the organization in which you want to create a custom role.
func (r *IAMPolicy) OrgId() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["orgId"])
}

// The `google_iam_policy` data source that represents
// the IAM policy that will be applied to the organization. This policy overrides any existing
// policy applied to the organization.
func (r *IAMPolicy) PolicyData() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["policyData"])
}

// Input properties used for looking up and filtering IAMPolicy resources.
type IAMPolicyState struct {
	Etag interface{}
	// The numeric ID of the organization in which you want to create a custom role.
	OrgId interface{}
	// The `google_iam_policy` data source that represents
	// the IAM policy that will be applied to the organization. This policy overrides any existing
	// policy applied to the organization.
	PolicyData interface{}
}

// The set of arguments for constructing a IAMPolicy resource.
type IAMPolicyArgs struct {
	// The numeric ID of the organization in which you want to create a custom role.
	OrgId interface{}
	// The `google_iam_policy` data source that represents
	// the IAM policy that will be applied to the organization. This policy overrides any existing
	// policy applied to the organization.
	PolicyData interface{}
}
