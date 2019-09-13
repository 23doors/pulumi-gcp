# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import utilities, tables

class Router(pulumi.CustomResource):
    bgp: pulumi.Output[dict]
    creation_timestamp: pulumi.Output[str]
    description: pulumi.Output[str]
    name: pulumi.Output[str]
    network: pulumi.Output[str]
    project: pulumi.Output[str]
    """
    The ID of the project in which the resource belongs.
    If it is not provided, the provider project is used.
    """
    region: pulumi.Output[str]
    self_link: pulumi.Output[str]
    """
    The URI of the created resource.
    """
    def __init__(__self__, resource_name, opts=None, bgp=None, description=None, name=None, network=None, project=None, region=None, __props__=None, __name__=None, __opts__=None):
        """
        Create a Router resource with the given unique name, props, and options.
        
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] project: The ID of the project in which the resource belongs.
               If it is not provided, the provider project is used.
        
        The **bgp** object supports the following:
        
          * `advertiseMode` (`pulumi.Input[str]`)
          * `advertisedGroups` (`pulumi.Input[list]`)
          * `advertisedIpRanges` (`pulumi.Input[list]`)
        
            * `description` (`pulumi.Input[str]`)
            * `range` (`pulumi.Input[str]`)
        
          * `asn` (`pulumi.Input[float]`)

        > This content is derived from https://github.com/terraform-providers/terraform-provider-google/blob/master/website/docs/r/compute_router.html.markdown.
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

            __props__['bgp'] = bgp
            __props__['description'] = description
            __props__['name'] = name
            if network is None:
                raise TypeError("Missing required property 'network'")
            __props__['network'] = network
            __props__['project'] = project
            __props__['region'] = region
            __props__['creation_timestamp'] = None
            __props__['self_link'] = None
        super(Router, __self__).__init__(
            'gcp:compute/router:Router',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None, bgp=None, creation_timestamp=None, description=None, name=None, network=None, project=None, region=None, self_link=None):
        """
        Get an existing Router resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.
        
        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] project: The ID of the project in which the resource belongs.
               If it is not provided, the provider project is used.
        :param pulumi.Input[str] self_link: The URI of the created resource.
        
        The **bgp** object supports the following:
        
          * `advertiseMode` (`pulumi.Input[str]`)
          * `advertisedGroups` (`pulumi.Input[list]`)
          * `advertisedIpRanges` (`pulumi.Input[list]`)
        
            * `description` (`pulumi.Input[str]`)
            * `range` (`pulumi.Input[str]`)
        
          * `asn` (`pulumi.Input[float]`)

        > This content is derived from https://github.com/terraform-providers/terraform-provider-google/blob/master/website/docs/r/compute_router.html.markdown.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()
        __props__["bgp"] = bgp
        __props__["creation_timestamp"] = creation_timestamp
        __props__["description"] = description
        __props__["name"] = name
        __props__["network"] = network
        __props__["project"] = project
        __props__["region"] = region
        __props__["self_link"] = self_link
        return Router(resource_name, opts=opts, __props__=__props__)
    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

