# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from .. import utilities, tables

class Policy(pulumi.CustomResource):
    alternative_name_server_config: pulumi.Output[dict]
    description: pulumi.Output[str]
    enable_inbound_forwarding: pulumi.Output[bool]
    enable_logging: pulumi.Output[bool]
    name: pulumi.Output[str]
    networks: pulumi.Output[list]
    project: pulumi.Output[str]
    """
    The ID of the project in which the resource belongs.
    If it is not provided, the provider project is used.
    """
    def __init__(__self__, resource_name, opts=None, alternative_name_server_config=None, description=None, enable_inbound_forwarding=None, enable_logging=None, name=None, networks=None, project=None, __name__=None, __opts__=None):
        """
        A policy is a collection of DNS rules applied to one or more Virtual
        Private Cloud resources.
        
        To get more information about Policy, see:
        
        * [API documentation](https://cloud.google.com/dns/docs/reference/v1beta2/policies)
        * How-to Guides
            * [Using DNS server policies](https://cloud.google.com/dns/zones/#using-dns-server-policies)
        
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] project: The ID of the project in which the resource belongs.
               If it is not provided, the provider project is used.

        > This content is derived from https://github.com/terraform-providers/terraform-provider-google/blob/master/website/docs/r/dns_policy.html.markdown.
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

        __props__['alternative_name_server_config'] = alternative_name_server_config

        __props__['description'] = description

        __props__['enable_inbound_forwarding'] = enable_inbound_forwarding

        __props__['enable_logging'] = enable_logging

        __props__['name'] = name

        __props__['networks'] = networks

        __props__['project'] = project

        if opts is None:
            opts = pulumi.ResourceOptions()
        if opts.version is None:
            opts.version = utilities.get_version()
        super(Policy, __self__).__init__(
            'gcp:dns/policy:Policy',
            resource_name,
            __props__,
            opts)


    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

