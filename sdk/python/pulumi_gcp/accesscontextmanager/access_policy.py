# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import pulumi
import pulumi.runtime
from .. import utilities, tables

class AccessPolicy(pulumi.CustomResource):
    create_time: pulumi.Output[str]
    name: pulumi.Output[str]
    parent: pulumi.Output[str]
    title: pulumi.Output[str]
    update_time: pulumi.Output[str]
    def __init__(__self__, __name__, __opts__=None, parent=None, title=None):
        """
        Create a AccessPolicy resource with the given unique name, props, and options.
        
        :param str __name__: The name of the resource.
        :param pulumi.ResourceOptions __opts__: Options for the resource.
        :param pulumi.Input[str] parent
        :param pulumi.Input[str] title
        """
        if not __name__:
            raise TypeError('Missing resource name argument (for URN creation)')
        if not isinstance(__name__, str):
            raise TypeError('Expected resource name to be a string')
        if __opts__ and not isinstance(__opts__, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')

        __props__ = dict()

        if not parent:
            raise TypeError('Missing required property parent')
        __props__['parent'] = parent

        if not title:
            raise TypeError('Missing required property title')
        __props__['title'] = title

        __props__['create_time'] = None
        __props__['name'] = None
        __props__['update_time'] = None

        super(AccessPolicy, __self__).__init__(
            'gcp:accesscontextmanager/accessPolicy:AccessPolicy',
            __name__,
            __props__,
            __opts__)


    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop
