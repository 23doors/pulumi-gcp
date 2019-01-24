// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package service

import (
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Get the service account from a project. For more information see
// the official [API](https://cloud.google.com/compute/docs/access/service-accounts) documentation.
func LookupAccount(ctx *pulumi.Context, args *GetAccountArgs) (*GetAccountResult, error) {
	inputs := make(map[string]interface{})
	if args != nil {
		inputs["accountId"] = args.AccountId
		inputs["project"] = args.Project
	}
	outputs, err := ctx.Invoke("gcp:service/getAccount:getAccount", inputs)
	if err != nil {
		return nil, err
	}
	return &GetAccountResult{
		DisplayName: outputs["displayName"],
		Email: outputs["email"],
		Name: outputs["name"],
		UniqueId: outputs["uniqueId"],
		Id: outputs["id"],
	}, nil
}

// A collection of arguments for invoking getAccount.
type GetAccountArgs struct {
	// The Service account id.
	AccountId interface{}
	// The ID of the project that the service account will be created in.
	// Defaults to the provider project configuration.
	Project interface{}
}

// A collection of values returned by getAccount.
type GetAccountResult struct {
	// The display name for the service account.
	DisplayName interface{}
	// The e-mail address of the service account. This value
	// should be referenced from any `google_iam_policy` data sources
	// that would grant the service account privileges.
	Email interface{}
	// The fully-qualified name of the service account.
	Name interface{}
	// The unique id of the service account.
	UniqueId interface{}
	// id is the provider-assigned unique ID for this managed resource.
	Id interface{}
}
