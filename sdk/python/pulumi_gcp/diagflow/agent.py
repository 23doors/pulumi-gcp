# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import utilities, tables

class Agent(pulumi.CustomResource):
    api_version: pulumi.Output[str]
    avatar_uri: pulumi.Output[str]
    avatar_uri_backend: pulumi.Output[str]
    classification_threshold: pulumi.Output[float]
    default_language_code: pulumi.Output[str]
    description: pulumi.Output[str]
    display_name: pulumi.Output[str]
    enable_logging: pulumi.Output[bool]
    match_mode: pulumi.Output[str]
    project: pulumi.Output[str]
    """
    The ID of the project in which the resource belongs.
    If it is not provided, the provider project is used.
    """
    supported_language_codes: pulumi.Output[list]
    tier: pulumi.Output[str]
    time_zone: pulumi.Output[str]
    def __init__(__self__, resource_name, opts=None, api_version=None, avatar_uri=None, classification_threshold=None, default_language_code=None, description=None, display_name=None, enable_logging=None, match_mode=None, project=None, supported_language_codes=None, tier=None, time_zone=None, __props__=None, __name__=None, __opts__=None):
        """
        Create a Agent resource with the given unique name, props, and options.
        
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] project: The ID of the project in which the resource belongs.
               If it is not provided, the provider project is used.

        > This content is derived from https://github.com/terraform-providers/terraform-provider-google/blob/master/website/docs/r/dialogflow_agent.html.markdown.
        """
        if __name__ is not None:
            warnings.warn("explicit use of __name__ is deprecated", DeprecationWarning)
            resource_name = __name__
        if __opts__ is not None:
            warnings.warn("explicit use of __opts__ is deprecated, use 'opts' instead", DeprecationWarning)
            opts = __opts__
        if opts is None:
            opts = pulumi.ResourceOptions()
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.version is None:
            opts.version = utilities.get_version()
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = dict()

            __props__['api_version'] = api_version
            __props__['avatar_uri'] = avatar_uri
            __props__['classification_threshold'] = classification_threshold
            if default_language_code is None:
                raise TypeError("Missing required property 'default_language_code'")
            __props__['default_language_code'] = default_language_code
            __props__['description'] = description
            if display_name is None:
                raise TypeError("Missing required property 'display_name'")
            __props__['display_name'] = display_name
            __props__['enable_logging'] = enable_logging
            __props__['match_mode'] = match_mode
            __props__['project'] = project
            __props__['supported_language_codes'] = supported_language_codes
            __props__['tier'] = tier
            if time_zone is None:
                raise TypeError("Missing required property 'time_zone'")
            __props__['time_zone'] = time_zone
            __props__['avatar_uri_backend'] = None
        super(Agent, __self__).__init__(
            'gcp:diagflow/agent:Agent',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None, api_version=None, avatar_uri=None, avatar_uri_backend=None, classification_threshold=None, default_language_code=None, description=None, display_name=None, enable_logging=None, match_mode=None, project=None, supported_language_codes=None, tier=None, time_zone=None):
        """
        Get an existing Agent resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.
        
        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] project: The ID of the project in which the resource belongs.
               If it is not provided, the provider project is used.

        > This content is derived from https://github.com/terraform-providers/terraform-provider-google/blob/master/website/docs/r/dialogflow_agent.html.markdown.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()
        __props__["api_version"] = api_version
        __props__["avatar_uri"] = avatar_uri
        __props__["avatar_uri_backend"] = avatar_uri_backend
        __props__["classification_threshold"] = classification_threshold
        __props__["default_language_code"] = default_language_code
        __props__["description"] = description
        __props__["display_name"] = display_name
        __props__["enable_logging"] = enable_logging
        __props__["match_mode"] = match_mode
        __props__["project"] = project
        __props__["supported_language_codes"] = supported_language_codes
        __props__["tier"] = tier
        __props__["time_zone"] = time_zone
        return Agent(resource_name, opts=opts, __props__=__props__)
    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

