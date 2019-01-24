// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package storage

import (
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

func LookupBucketObject(ctx *pulumi.Context, args *GetBucketObjectArgs) (*GetBucketObjectResult, error) {
	inputs := make(map[string]interface{})
	if args != nil {
		inputs["bucket"] = args.Bucket
		inputs["name"] = args.Name
	}
	outputs, err := ctx.Invoke("gcp:storage/getBucketObject:getBucketObject", inputs)
	if err != nil {
		return nil, err
	}
	return &GetBucketObjectResult{
		CacheControl: outputs["cacheControl"],
		Content: outputs["content"],
		ContentDisposition: outputs["contentDisposition"],
		ContentEncoding: outputs["contentEncoding"],
		ContentLanguage: outputs["contentLanguage"],
		ContentType: outputs["contentType"],
		Crc32c: outputs["crc32c"],
		DetectMd5hash: outputs["detectMd5hash"],
		Md5hash: outputs["md5hash"],
		OutputName: outputs["outputName"],
		PredefinedAcl: outputs["predefinedAcl"],
		SelfLink: outputs["selfLink"],
		Source: outputs["source"],
		StorageClass: outputs["storageClass"],
		Id: outputs["id"],
	}, nil
}

// A collection of arguments for invoking getBucketObject.
type GetBucketObjectArgs struct {
	Bucket interface{}
	Name interface{}
}

// A collection of values returned by getBucketObject.
type GetBucketObjectResult struct {
	CacheControl interface{}
	Content interface{}
	ContentDisposition interface{}
	ContentEncoding interface{}
	ContentLanguage interface{}
	ContentType interface{}
	Crc32c interface{}
	DetectMd5hash interface{}
	Md5hash interface{}
	OutputName interface{}
	PredefinedAcl interface{}
	SelfLink interface{}
	Source interface{}
	StorageClass interface{}
	// id is the provider-assigned unique ID for this managed resource.
	Id interface{}
}