// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package secretmanager

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

type SecretIamPolicy struct {
	pulumi.CustomResourceState

	Etag pulumi.StringOutput `pulumi:"etag"`
	PolicyData pulumi.StringOutput `pulumi:"policyData"`
	Project pulumi.StringOutput `pulumi:"project"`
	SecretId pulumi.StringOutput `pulumi:"secretId"`
}

// NewSecretIamPolicy registers a new resource with the given unique name, arguments, and options.
func NewSecretIamPolicy(ctx *pulumi.Context,
	name string, args *SecretIamPolicyArgs, opts ...pulumi.ResourceOption) (*SecretIamPolicy, error) {
	if args == nil || args.PolicyData == nil {
		return nil, errors.New("missing required argument 'PolicyData'")
	}
	if args == nil || args.SecretId == nil {
		return nil, errors.New("missing required argument 'SecretId'")
	}
	if args == nil {
		args = &SecretIamPolicyArgs{}
	}
	var resource SecretIamPolicy
	err := ctx.RegisterResource("gcp:secretmanager/secretIamPolicy:SecretIamPolicy", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSecretIamPolicy gets an existing SecretIamPolicy resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSecretIamPolicy(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SecretIamPolicyState, opts ...pulumi.ResourceOption) (*SecretIamPolicy, error) {
	var resource SecretIamPolicy
	err := ctx.ReadResource("gcp:secretmanager/secretIamPolicy:SecretIamPolicy", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SecretIamPolicy resources.
type secretIamPolicyState struct {
	Etag *string `pulumi:"etag"`
	PolicyData *string `pulumi:"policyData"`
	Project *string `pulumi:"project"`
	SecretId *string `pulumi:"secretId"`
}

type SecretIamPolicyState struct {
	Etag pulumi.StringPtrInput
	PolicyData pulumi.StringPtrInput
	Project pulumi.StringPtrInput
	SecretId pulumi.StringPtrInput
}

func (SecretIamPolicyState) ElementType() reflect.Type {
	return reflect.TypeOf((*secretIamPolicyState)(nil)).Elem()
}

type secretIamPolicyArgs struct {
	PolicyData string `pulumi:"policyData"`
	Project *string `pulumi:"project"`
	SecretId string `pulumi:"secretId"`
}

// The set of arguments for constructing a SecretIamPolicy resource.
type SecretIamPolicyArgs struct {
	PolicyData pulumi.StringInput
	Project pulumi.StringPtrInput
	SecretId pulumi.StringInput
}

func (SecretIamPolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*secretIamPolicyArgs)(nil)).Elem()
}

