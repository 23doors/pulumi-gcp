# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import utilities, tables

class StandardAppVersion(pulumi.CustomResource):
    delete_service_on_destroy: pulumi.Output[bool]
    """
    If set to `true`, the service will be deleted if it is the last version.    
    """
    deployment: pulumi.Output[dict]
    entrypoint: pulumi.Output[dict]
    env_variables: pulumi.Output[dict]
    handlers: pulumi.Output[list]
    instance_class: pulumi.Output[str]
    libraries: pulumi.Output[list]
    name: pulumi.Output[str]
    noop_on_destroy: pulumi.Output[bool]
    """
    If set to `true`, the application version will not be deleted.
    """
    project: pulumi.Output[str]
    """
    The ID of the project in which the resource belongs.
    If it is not provided, the provider project is used.
    """
    runtime: pulumi.Output[str]
    runtime_api_version: pulumi.Output[str]
    service: pulumi.Output[str]
    threadsafe: pulumi.Output[bool]
    version_id: pulumi.Output[str]
    def __init__(__self__, resource_name, opts=None, delete_service_on_destroy=None, deployment=None, entrypoint=None, env_variables=None, handlers=None, instance_class=None, libraries=None, noop_on_destroy=None, project=None, runtime=None, runtime_api_version=None, service=None, threadsafe=None, version_id=None, __props__=None, __name__=None, __opts__=None):
        """
        Create a StandardAppVersion resource with the given unique name, props, and options.
        
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[bool] delete_service_on_destroy: If set to `true`, the service will be deleted if it is the last version.    
        :param pulumi.Input[bool] noop_on_destroy: If set to `true`, the application version will not be deleted.
        :param pulumi.Input[str] project: The ID of the project in which the resource belongs.
               If it is not provided, the provider project is used.
        
        The **deployment** object supports the following:
        
          * `files` (`pulumi.Input[list]`)
        
            * `name` (`pulumi.Input[str]`)
            * `sha1Sum` (`pulumi.Input[str]`)
            * `sourceUrl` (`pulumi.Input[str]`)
        
          * `zip` (`pulumi.Input[dict]`)
        
            * `filesCount` (`pulumi.Input[float]`)
            * `sourceUrl` (`pulumi.Input[str]`)
        
        The **entrypoint** object supports the following:
        
          * `shell` (`pulumi.Input[str]`)
        
        The **handlers** object supports the following:
        
          * `authFailAction` (`pulumi.Input[str]`)
          * `login` (`pulumi.Input[str]`)
          * `redirectHttpResponseCode` (`pulumi.Input[str]`)
          * `script` (`pulumi.Input[dict]`)
        
            * `scriptPath` (`pulumi.Input[str]`)
        
          * `securityLevel` (`pulumi.Input[str]`)
          * `staticFiles` (`pulumi.Input[dict]`)
        
            * `applicationReadable` (`pulumi.Input[bool]`)
            * `expiration` (`pulumi.Input[str]`)
            * `httpHeaders` (`pulumi.Input[dict]`)
            * `mimeType` (`pulumi.Input[str]`)
            * `path` (`pulumi.Input[str]`)
            * `requireMatchingFile` (`pulumi.Input[bool]`)
            * `uploadPathRegex` (`pulumi.Input[str]`)
        
          * `urlRegex` (`pulumi.Input[str]`)
        
        The **libraries** object supports the following:
        
          * `name` (`pulumi.Input[str]`)
          * `version` (`pulumi.Input[str]`)

        > This content is derived from https://github.com/terraform-providers/terraform-provider-google/blob/master/website/docs/r/app_engine_standard_app_version.html.markdown.
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

            __props__['delete_service_on_destroy'] = delete_service_on_destroy
            __props__['deployment'] = deployment
            __props__['entrypoint'] = entrypoint
            __props__['env_variables'] = env_variables
            __props__['handlers'] = handlers
            __props__['instance_class'] = instance_class
            __props__['libraries'] = libraries
            __props__['noop_on_destroy'] = noop_on_destroy
            __props__['project'] = project
            if runtime is None:
                raise TypeError("Missing required property 'runtime'")
            __props__['runtime'] = runtime
            __props__['runtime_api_version'] = runtime_api_version
            __props__['service'] = service
            __props__['threadsafe'] = threadsafe
            __props__['version_id'] = version_id
            __props__['name'] = None
        super(StandardAppVersion, __self__).__init__(
            'gcp:appengine/standardAppVersion:StandardAppVersion',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None, delete_service_on_destroy=None, deployment=None, entrypoint=None, env_variables=None, handlers=None, instance_class=None, libraries=None, name=None, noop_on_destroy=None, project=None, runtime=None, runtime_api_version=None, service=None, threadsafe=None, version_id=None):
        """
        Get an existing StandardAppVersion resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.
        
        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[bool] delete_service_on_destroy: If set to `true`, the service will be deleted if it is the last version.    
        :param pulumi.Input[bool] noop_on_destroy: If set to `true`, the application version will not be deleted.
        :param pulumi.Input[str] project: The ID of the project in which the resource belongs.
               If it is not provided, the provider project is used.
        
        The **deployment** object supports the following:
        
          * `files` (`pulumi.Input[list]`)
        
            * `name` (`pulumi.Input[str]`)
            * `sha1Sum` (`pulumi.Input[str]`)
            * `sourceUrl` (`pulumi.Input[str]`)
        
          * `zip` (`pulumi.Input[dict]`)
        
            * `filesCount` (`pulumi.Input[float]`)
            * `sourceUrl` (`pulumi.Input[str]`)
        
        The **entrypoint** object supports the following:
        
          * `shell` (`pulumi.Input[str]`)
        
        The **handlers** object supports the following:
        
          * `authFailAction` (`pulumi.Input[str]`)
          * `login` (`pulumi.Input[str]`)
          * `redirectHttpResponseCode` (`pulumi.Input[str]`)
          * `script` (`pulumi.Input[dict]`)
        
            * `scriptPath` (`pulumi.Input[str]`)
        
          * `securityLevel` (`pulumi.Input[str]`)
          * `staticFiles` (`pulumi.Input[dict]`)
        
            * `applicationReadable` (`pulumi.Input[bool]`)
            * `expiration` (`pulumi.Input[str]`)
            * `httpHeaders` (`pulumi.Input[dict]`)
            * `mimeType` (`pulumi.Input[str]`)
            * `path` (`pulumi.Input[str]`)
            * `requireMatchingFile` (`pulumi.Input[bool]`)
            * `uploadPathRegex` (`pulumi.Input[str]`)
        
          * `urlRegex` (`pulumi.Input[str]`)
        
        The **libraries** object supports the following:
        
          * `name` (`pulumi.Input[str]`)
          * `version` (`pulumi.Input[str]`)

        > This content is derived from https://github.com/terraform-providers/terraform-provider-google/blob/master/website/docs/r/app_engine_standard_app_version.html.markdown.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()
        __props__["delete_service_on_destroy"] = delete_service_on_destroy
        __props__["deployment"] = deployment
        __props__["entrypoint"] = entrypoint
        __props__["env_variables"] = env_variables
        __props__["handlers"] = handlers
        __props__["instance_class"] = instance_class
        __props__["libraries"] = libraries
        __props__["name"] = name
        __props__["noop_on_destroy"] = noop_on_destroy
        __props__["project"] = project
        __props__["runtime"] = runtime
        __props__["runtime_api_version"] = runtime_api_version
        __props__["service"] = service
        __props__["threadsafe"] = threadsafe
        __props__["version_id"] = version_id
        return StandardAppVersion(resource_name, opts=opts, __props__=__props__)
    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

