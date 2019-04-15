// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package compute

import (
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Get the IP address from a static address. For more information see
// the official [API](https://cloud.google.com/compute/docs/reference/latest/addresses/get) documentation.
func LookupAddress(ctx *pulumi.Context, args *GetAddressArgs) (*GetAddressResult, error) {
	inputs := make(map[string]interface{})
	if args != nil {
		inputs["name"] = args.Name
		inputs["project"] = args.Project
		inputs["region"] = args.Region
	}
	outputs, err := ctx.Invoke("gcp:compute/getAddress:getAddress", inputs)
	if err != nil {
		return nil, err
	}
	return &GetAddressResult{
		Address: outputs["address"],
		Name: outputs["name"],
		Project: outputs["project"],
		Region: outputs["region"],
		SelfLink: outputs["selfLink"],
		Status: outputs["status"],
		Id: outputs["id"],
	}, nil
}

// A collection of arguments for invoking getAddress.
type GetAddressArgs struct {
	// A unique name for the resource, required by GCE.
	Name interface{}
	// The project in which the resource belongs. If it
	// is not provided, the provider project is used.
	Project interface{}
	// The Region in which the created address reside.
	// If it is not provided, the provider region is used.
	Region interface{}
}

// A collection of values returned by getAddress.
type GetAddressResult struct {
	// The IP of the created resource.
	Address interface{}
	Name interface{}
	Project interface{}
	Region interface{}
	// The URI of the created resource.
	SelfLink interface{}
	// Indicates if the address is used. Possible values are: RESERVED or IN_USE.
	Status interface{}
	// id is the provider-assigned unique ID for this managed resource.
	Id interface{}
}
