# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import utilities, tables

class ServicePerimeter(pulumi.CustomResource):
    create_time: pulumi.Output[str]
    description: pulumi.Output[str]
    name: pulumi.Output[str]
    parent: pulumi.Output[str]
    perimeter_type: pulumi.Output[str]
    status: pulumi.Output[dict]
    title: pulumi.Output[str]
    update_time: pulumi.Output[str]
    def __init__(__self__, resource_name, opts=None, description=None, name=None, parent=None, perimeter_type=None, status=None, title=None, __props__=None, __name__=None, __opts__=None):
        """
        ServicePerimeter describes a set of GCP resources which can freely import
        and export data amongst themselves, but not export outside of the
        ServicePerimeter. If a request with a source within this ServicePerimeter
        has a target outside of the ServicePerimeter, the request will be blocked.
        Otherwise the request is allowed. There are two types of Service Perimeter
        - Regular and Bridge. Regular Service Perimeters cannot overlap, a single
        GCP project can only belong to a single regular Service Perimeter. Service
        Perimeter Bridges can contain only GCP projects as members, a single GCP
        project may belong to multiple Service Perimeter Bridges.
        
        
        To get more information about ServicePerimeter, see:
        
        * [API documentation](https://cloud.google.com/access-context-manager/docs/reference/rest/v1/accessPolicies.servicePerimeters)
        * How-to Guides
            * [Service Perimeter Quickstart](https://cloud.google.com/vpc-service-controls/docs/quickstart)
        
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        
        The **status** object supports the following:
        
          * `accessLevels` (`pulumi.Input[list]`)
          * `resources` (`pulumi.Input[list]`)
          * `restrictedServices` (`pulumi.Input[list]`)

        > This content is derived from https://github.com/terraform-providers/terraform-provider-google/blob/master/website/docs/r/access_context_manager_service_perimeter.html.markdown.
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

            __props__['description'] = description
            __props__['name'] = name
            if parent is None:
                raise TypeError("Missing required property 'parent'")
            __props__['parent'] = parent
            __props__['perimeter_type'] = perimeter_type
            __props__['status'] = status
            if title is None:
                raise TypeError("Missing required property 'title'")
            __props__['title'] = title
            __props__['create_time'] = None
            __props__['update_time'] = None
        super(ServicePerimeter, __self__).__init__(
            'gcp:accesscontextmanager/servicePerimeter:ServicePerimeter',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None, create_time=None, description=None, name=None, parent=None, perimeter_type=None, status=None, title=None, update_time=None):
        """
        Get an existing ServicePerimeter resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.
        
        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        
        The **status** object supports the following:
        
          * `accessLevels` (`pulumi.Input[list]`)
          * `resources` (`pulumi.Input[list]`)
          * `restrictedServices` (`pulumi.Input[list]`)

        > This content is derived from https://github.com/terraform-providers/terraform-provider-google/blob/master/website/docs/r/access_context_manager_service_perimeter.html.markdown.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()
        __props__["create_time"] = create_time
        __props__["description"] = description
        __props__["name"] = name
        __props__["parent"] = parent
        __props__["perimeter_type"] = perimeter_type
        __props__["status"] = status
        __props__["title"] = title
        __props__["update_time"] = update_time
        return ServicePerimeter(resource_name, opts=opts, __props__=__props__)
    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

