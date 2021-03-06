# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import utilities, tables

class FhirStore(pulumi.CustomResource):
    dataset: pulumi.Output[str]
    disable_referential_integrity: pulumi.Output[bool]
    disable_resource_versioning: pulumi.Output[bool]
    enable_history_import: pulumi.Output[bool]
    enable_update_create: pulumi.Output[bool]
    labels: pulumi.Output[dict]
    name: pulumi.Output[str]
    notification_config: pulumi.Output[dict]
    self_link: pulumi.Output[str]
    version: pulumi.Output[str]
    def __init__(__self__, resource_name, opts=None, dataset=None, disable_referential_integrity=None, disable_resource_versioning=None, enable_history_import=None, enable_update_create=None, labels=None, name=None, notification_config=None, version=None, __props__=None, __name__=None, __opts__=None):
        """
        Create a FhirStore resource with the given unique name, props, and options.
        
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        
        The **notification_config** object supports the following:
        
          * `pubsubTopic` (`pulumi.Input[str]`)

        > This content is derived from https://github.com/terraform-providers/terraform-provider-google/blob/master/website/docs/r/healthcare_fhir_store.html.markdown.
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

            if dataset is None:
                raise TypeError("Missing required property 'dataset'")
            __props__['dataset'] = dataset
            __props__['disable_referential_integrity'] = disable_referential_integrity
            __props__['disable_resource_versioning'] = disable_resource_versioning
            __props__['enable_history_import'] = enable_history_import
            __props__['enable_update_create'] = enable_update_create
            __props__['labels'] = labels
            __props__['name'] = name
            __props__['notification_config'] = notification_config
            __props__['version'] = version
            __props__['self_link'] = None
        super(FhirStore, __self__).__init__(
            'gcp:healthcare/fhirStore:FhirStore',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None, dataset=None, disable_referential_integrity=None, disable_resource_versioning=None, enable_history_import=None, enable_update_create=None, labels=None, name=None, notification_config=None, self_link=None, version=None):
        """
        Get an existing FhirStore resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.
        
        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        
        The **notification_config** object supports the following:
        
          * `pubsubTopic` (`pulumi.Input[str]`)

        > This content is derived from https://github.com/terraform-providers/terraform-provider-google/blob/master/website/docs/r/healthcare_fhir_store.html.markdown.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()
        __props__["dataset"] = dataset
        __props__["disable_referential_integrity"] = disable_referential_integrity
        __props__["disable_resource_versioning"] = disable_resource_versioning
        __props__["enable_history_import"] = enable_history_import
        __props__["enable_update_create"] = enable_update_create
        __props__["labels"] = labels
        __props__["name"] = name
        __props__["notification_config"] = notification_config
        __props__["self_link"] = self_link
        __props__["version"] = version
        return FhirStore(resource_name, opts=opts, __props__=__props__)
    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

