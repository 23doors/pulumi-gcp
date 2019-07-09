// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package healthcare

import (
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

type DicomStoreIamMember struct {
	s *pulumi.ResourceState
}

// NewDicomStoreIamMember registers a new resource with the given unique name, arguments, and options.
func NewDicomStoreIamMember(ctx *pulumi.Context,
	name string, args *DicomStoreIamMemberArgs, opts ...pulumi.ResourceOpt) (*DicomStoreIamMember, error) {
	if args == nil || args.DicomStoreId == nil {
		return nil, errors.New("missing required argument 'DicomStoreId'")
	}
	if args == nil || args.Member == nil {
		return nil, errors.New("missing required argument 'Member'")
	}
	if args == nil || args.Role == nil {
		return nil, errors.New("missing required argument 'Role'")
	}
	inputs := make(map[string]interface{})
	if args == nil {
		inputs["dicomStoreId"] = nil
		inputs["member"] = nil
		inputs["role"] = nil
	} else {
		inputs["dicomStoreId"] = args.DicomStoreId
		inputs["member"] = args.Member
		inputs["role"] = args.Role
	}
	inputs["etag"] = nil
	s, err := ctx.RegisterResource("gcp:healthcare/dicomStoreIamMember:DicomStoreIamMember", name, true, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &DicomStoreIamMember{s: s}, nil
}

// GetDicomStoreIamMember gets an existing DicomStoreIamMember resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDicomStoreIamMember(ctx *pulumi.Context,
	name string, id pulumi.ID, state *DicomStoreIamMemberState, opts ...pulumi.ResourceOpt) (*DicomStoreIamMember, error) {
	inputs := make(map[string]interface{})
	if state != nil {
		inputs["dicomStoreId"] = state.DicomStoreId
		inputs["etag"] = state.Etag
		inputs["member"] = state.Member
		inputs["role"] = state.Role
	}
	s, err := ctx.ReadResource("gcp:healthcare/dicomStoreIamMember:DicomStoreIamMember", name, id, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &DicomStoreIamMember{s: s}, nil
}

// URN is this resource's unique name assigned by Pulumi.
func (r *DicomStoreIamMember) URN() *pulumi.URNOutput {
	return r.s.URN()
}

// ID is this resource's unique identifier assigned by its provider.
func (r *DicomStoreIamMember) ID() *pulumi.IDOutput {
	return r.s.ID()
}

// The DICOM store ID, in the form
// `{project_id}/{location_name}/{dataset_name}/{dicom_store_name}` or
// `{location_name}/{dataset_name}/{dicom_store_name}`. In the second form, the provider's
// project setting will be used as a fallback.
func (r *DicomStoreIamMember) DicomStoreId() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["dicomStoreId"])
}

// (Computed) The etag of the DICOM store's IAM policy.
func (r *DicomStoreIamMember) Etag() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["etag"])
}

func (r *DicomStoreIamMember) Member() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["member"])
}

// The role that should be applied. Only one
// `google_healthcare_dicom_store_iam_binding` can be used per role. Note that custom roles must be of the format
// `[projects|organizations]/{parent-name}/roles/{role-name}`.
func (r *DicomStoreIamMember) Role() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["role"])
}

// Input properties used for looking up and filtering DicomStoreIamMember resources.
type DicomStoreIamMemberState struct {
	// The DICOM store ID, in the form
	// `{project_id}/{location_name}/{dataset_name}/{dicom_store_name}` or
	// `{location_name}/{dataset_name}/{dicom_store_name}`. In the second form, the provider's
	// project setting will be used as a fallback.
	DicomStoreId interface{}
	// (Computed) The etag of the DICOM store's IAM policy.
	Etag interface{}
	Member interface{}
	// The role that should be applied. Only one
	// `google_healthcare_dicom_store_iam_binding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role interface{}
}

// The set of arguments for constructing a DicomStoreIamMember resource.
type DicomStoreIamMemberArgs struct {
	// The DICOM store ID, in the form
	// `{project_id}/{location_name}/{dataset_name}/{dicom_store_name}` or
	// `{location_name}/{dataset_name}/{dicom_store_name}`. In the second form, the provider's
	// project setting will be used as a fallback.
	DicomStoreId interface{}
	Member interface{}
	// The role that should be applied. Only one
	// `google_healthcare_dicom_store_iam_binding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role interface{}
}