# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from .. import utilities, tables

class AccessLevel(pulumi.CustomResource):
    basic: pulumi.Output[dict]
    description: pulumi.Output[str]
    name: pulumi.Output[str]
    parent: pulumi.Output[str]
    title: pulumi.Output[str]
    def __init__(__self__, resource_name, opts=None, basic=None, description=None, name=None, parent=None, title=None, __name__=None, __opts__=None):
        """
        An AccessLevel is a label that can be applied to requests to GCP services,
        along with a list of requirements necessary for the label to be applied.
        
        > **Warning:** This resource is in beta, and should be used with the terraform-provider-google-beta provider.
        See [Provider Versions](https://terraform.io/docs/providers/google/provider_versions.html) for more details on beta resources.
        
        To get more information about AccessLevel, see:
        
        * [API documentation](https://cloud.google.com/access-context-manager/docs/reference/rest/v1beta/accessPolicies.accessLevels)
        * How-to Guides
            * [Access Policy Quickstart](https://cloud.google.com/access-context-manager/docs/quickstart)
        
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        if __name__ is not None:
            warnings.warn("explicit use of __name__ is deprecated", DeprecationWarning)
            resource_name = __name__
        if __opts__ is not None:
            warnings.warn("explicit use of __opts__ is deprecated, use 'opts' instead", DeprecationWarning)
            opts = __opts__
        if not resource_name:
            raise TypeError('Missing resource name argument (for URN creation)')
        if not isinstance(resource_name, str):
            raise TypeError('Expected resource name to be a string')
        if opts and not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')

        __props__ = dict()

        __props__['basic'] = basic

        __props__['description'] = description

        __props__['name'] = name

        if parent is None:
            raise TypeError("Missing required property 'parent'")
        __props__['parent'] = parent

        if title is None:
            raise TypeError("Missing required property 'title'")
        __props__['title'] = title

        super(AccessLevel, __self__).__init__(
            'gcp:accesscontextmanager/accessLevel:AccessLevel',
            resource_name,
            __props__,
            opts)


    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

