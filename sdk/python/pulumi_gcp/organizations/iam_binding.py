# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import pulumi
import pulumi.runtime
from .. import utilities, tables

class IAMBinding(pulumi.CustomResource):
    """
    Allows creation and management of a single binding within IAM policy for
    an existing Google Cloud Platform Organization.
    
    > **Note:** This resource __must not__ be used in conjunction with
       `google_organization_iam_member` for the __same role__ or they will fight over
       what your policy should be.
    """
    def __init__(__self__, __name__, __opts__=None, members=None, org_id=None, role=None):
        """Create a IAMBinding resource with the given unique name, props, and options."""
        if not __name__:
            raise TypeError('Missing resource name argument (for URN creation)')
        if not isinstance(__name__, str):
            raise TypeError('Expected resource name to be a string')
        if __opts__ and not isinstance(__opts__, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')

        __props__ = dict()

        if not members:
            raise TypeError('Missing required property members')
        __props__['members'] = members

        if not org_id:
            raise TypeError('Missing required property org_id')
        __props__['org_id'] = org_id

        if not role:
            raise TypeError('Missing required property role')
        __props__['role'] = role

        __props__['etag'] = None

        super(IAMBinding, __self__).__init__(
            'gcp:organizations/iAMBinding:IAMBinding',
            __name__,
            __props__,
            __opts__)


    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

