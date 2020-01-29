// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package serviceAccount

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Allows management of a [Google Cloud Platform service account](https://cloud.google.com/compute/docs/access/service-accounts)
// 
// > Creation of service accounts is eventually consistent, and that can lead to
// errors when you try to apply ACLs to service accounts immediately after
// creation.
// 
// > This content is derived from https://github.com/terraform-providers/terraform-provider-google/blob/master/website/docs/r/service_account.html.markdown.
type Account struct {
	pulumi.CustomResourceState

	// The account id that is used to generate the service
	// account email address and a stable unique id. It is unique within a project,
	// must be 6-30 characters long, and match the regular expression `a-z`
	// to comply with RFC1035. Changing this forces a new service account to be created.
	AccountId pulumi.StringOutput `pulumi:"accountId"`
	// A text description of the service account.
	// Must be less than or equal to 256 UTF-8 bytes.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// The display name for the service account.
	// Can be updated without creating a new resource.
	DisplayName pulumi.StringPtrOutput `pulumi:"displayName"`
	// The e-mail address of the service account. This value
	// should be referenced from any `organizations.getIAMPolicy` data sources
	// that would grant the service account privileges.
	Email pulumi.StringOutput `pulumi:"email"`
	// The fully-qualified name of the service account.
	Name pulumi.StringOutput `pulumi:"name"`
	// The ID of the project that the service account will be created in.
	// Defaults to the provider project configuration.
	Project pulumi.StringOutput `pulumi:"project"`
	// The unique id of the service account.
	UniqueId pulumi.StringOutput `pulumi:"uniqueId"`
}

// NewAccount registers a new resource with the given unique name, arguments, and options.
func NewAccount(ctx *pulumi.Context,
	name string, args *AccountArgs, opts ...pulumi.ResourceOption) (*Account, error) {
	if args == nil || args.AccountId == nil {
		return nil, errors.New("missing required argument 'AccountId'")
	}
	if args == nil {
		args = &AccountArgs{}
	}
	var resource Account
	err := ctx.RegisterResource("gcp:serviceAccount/account:Account", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAccount gets an existing Account resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAccount(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AccountState, opts ...pulumi.ResourceOption) (*Account, error) {
	var resource Account
	err := ctx.ReadResource("gcp:serviceAccount/account:Account", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Account resources.
type accountState struct {
	// The account id that is used to generate the service
	// account email address and a stable unique id. It is unique within a project,
	// must be 6-30 characters long, and match the regular expression `a-z`
	// to comply with RFC1035. Changing this forces a new service account to be created.
	AccountId *string `pulumi:"accountId"`
	// A text description of the service account.
	// Must be less than or equal to 256 UTF-8 bytes.
	Description *string `pulumi:"description"`
	// The display name for the service account.
	// Can be updated without creating a new resource.
	DisplayName *string `pulumi:"displayName"`
	// The e-mail address of the service account. This value
	// should be referenced from any `organizations.getIAMPolicy` data sources
	// that would grant the service account privileges.
	Email *string `pulumi:"email"`
	// The fully-qualified name of the service account.
	Name *string `pulumi:"name"`
	// The ID of the project that the service account will be created in.
	// Defaults to the provider project configuration.
	Project *string `pulumi:"project"`
	// The unique id of the service account.
	UniqueId *string `pulumi:"uniqueId"`
}

type AccountState struct {
	// The account id that is used to generate the service
	// account email address and a stable unique id. It is unique within a project,
	// must be 6-30 characters long, and match the regular expression `a-z`
	// to comply with RFC1035. Changing this forces a new service account to be created.
	AccountId pulumi.StringPtrInput
	// A text description of the service account.
	// Must be less than or equal to 256 UTF-8 bytes.
	Description pulumi.StringPtrInput
	// The display name for the service account.
	// Can be updated without creating a new resource.
	DisplayName pulumi.StringPtrInput
	// The e-mail address of the service account. This value
	// should be referenced from any `organizations.getIAMPolicy` data sources
	// that would grant the service account privileges.
	Email pulumi.StringPtrInput
	// The fully-qualified name of the service account.
	Name pulumi.StringPtrInput
	// The ID of the project that the service account will be created in.
	// Defaults to the provider project configuration.
	Project pulumi.StringPtrInput
	// The unique id of the service account.
	UniqueId pulumi.StringPtrInput
}

func (AccountState) ElementType() reflect.Type {
	return reflect.TypeOf((*accountState)(nil)).Elem()
}

type accountArgs struct {
	// The account id that is used to generate the service
	// account email address and a stable unique id. It is unique within a project,
	// must be 6-30 characters long, and match the regular expression `a-z`
	// to comply with RFC1035. Changing this forces a new service account to be created.
	AccountId string `pulumi:"accountId"`
	// A text description of the service account.
	// Must be less than or equal to 256 UTF-8 bytes.
	Description *string `pulumi:"description"`
	// The display name for the service account.
	// Can be updated without creating a new resource.
	DisplayName *string `pulumi:"displayName"`
	// The ID of the project that the service account will be created in.
	// Defaults to the provider project configuration.
	Project *string `pulumi:"project"`
}

// The set of arguments for constructing a Account resource.
type AccountArgs struct {
	// The account id that is used to generate the service
	// account email address and a stable unique id. It is unique within a project,
	// must be 6-30 characters long, and match the regular expression `a-z`
	// to comply with RFC1035. Changing this forces a new service account to be created.
	AccountId pulumi.StringInput
	// A text description of the service account.
	// Must be less than or equal to 256 UTF-8 bytes.
	Description pulumi.StringPtrInput
	// The display name for the service account.
	// Can be updated without creating a new resource.
	DisplayName pulumi.StringPtrInput
	// The ID of the project that the service account will be created in.
	// Defaults to the provider project configuration.
	Project pulumi.StringPtrInput
}

func (AccountArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*accountArgs)(nil)).Elem()
}

