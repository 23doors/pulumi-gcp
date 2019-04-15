# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from .. import utilities, tables

class URLMap(pulumi.CustomResource):
    creation_timestamp: pulumi.Output[str]
    default_service: pulumi.Output[str]
    description: pulumi.Output[str]
    fingerprint: pulumi.Output[str]
    host_rules: pulumi.Output[list]
    map_id: pulumi.Output[float]
    name: pulumi.Output[str]
    path_matchers: pulumi.Output[list]
    project: pulumi.Output[str]
    """
    The ID of the project in which the resource belongs.
    If it is not provided, the provider project is used.
    """
    self_link: pulumi.Output[str]
    """
    The URI of the created resource.
    """
    tests: pulumi.Output[list]
    def __init__(__self__, resource_name, opts=None, default_service=None, description=None, host_rules=None, name=None, path_matchers=None, project=None, tests=None, __name__=None, __opts__=None):
        """
        UrlMaps are used to route requests to a backend service based on rules
        that you define for the host and path of an incoming URL.
        
        
        
        <div class = "oics-button" style="float: right; margin: 0 0 -15px">
          <a href="https://console.cloud.google.com/cloudshell/open?cloudshell_git_repo=https%3A%2F%2Fgithub.com%2Fterraform-google-modules%2Fdocs-examples.git&cloudshell_working_dir=url_map_basic&cloudshell_image=gcr.io%2Fgraphite-cloud-shell-images%2Fterraform%3Alatest&open_in_editor=main.tf&cloudshell_print=.%2Fmotd&cloudshell_tutorial=.%2Ftutorial.md" target="_blank">
            <img alt="Open in Cloud Shell" src="//gstatic.com/cloudssh/images/open-btn.svg" style="max-height: 44px; margin: 32px auto; max-width: 100%;">
          </a>
        </div>
        
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] project: The ID of the project in which the resource belongs.
               If it is not provided, the provider project is used.
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

        if default_service is None:
            raise TypeError("Missing required property 'default_service'")
        __props__['default_service'] = default_service

        __props__['description'] = description

        __props__['host_rules'] = host_rules

        __props__['name'] = name

        __props__['path_matchers'] = path_matchers

        __props__['project'] = project

        __props__['tests'] = tests

        __props__['creation_timestamp'] = None
        __props__['fingerprint'] = None
        __props__['map_id'] = None
        __props__['self_link'] = None

        super(URLMap, __self__).__init__(
            'gcp:compute/uRLMap:URLMap',
            resource_name,
            __props__,
            opts)


    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

