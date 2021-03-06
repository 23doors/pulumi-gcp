# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import utilities, tables

class ObjectAccessControl(pulumi.CustomResource):
    bucket: pulumi.Output[str]
    domain: pulumi.Output[str]
    email: pulumi.Output[str]
    entity: pulumi.Output[str]
    entity_id: pulumi.Output[str]
    generation: pulumi.Output[float]
    object: pulumi.Output[str]
    project_team: pulumi.Output[dict]
    role: pulumi.Output[str]
    def __init__(__self__, resource_name, opts=None, bucket=None, entity=None, object=None, role=None, __props__=None, __name__=None, __opts__=None):
        """
        The ObjectAccessControls resources represent the Access Control Lists
        (ACLs) for objects within Google Cloud Storage. ACLs let you specify
        who has access to your data and to what extent.
        
        There are two roles that can be assigned to an entity:
        
        READERs can get an object, though the acl property will not be revealed.
        OWNERs are READERs, and they can get the acl property, update an object,
        and call all objectAccessControls methods on the object. The owner of an
        object is always an OWNER.
        For more information, see Access Control, with the caveat that this API
        uses READER and OWNER instead of READ and FULL_CONTROL.
        
        
        To get more information about ObjectAccessControl, see:
        
        * [API documentation](https://cloud.google.com/storage/docs/json_api/v1/objectAccessControls)
        * How-to Guides
            * [Official Documentation](https://cloud.google.com/storage/docs/access-control/create-manage-lists)
        
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.

        > This content is derived from https://github.com/terraform-providers/terraform-provider-google/blob/master/website/docs/r/storage_object_access_control.html.markdown.
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

            if bucket is None:
                raise TypeError("Missing required property 'bucket'")
            __props__['bucket'] = bucket
            if entity is None:
                raise TypeError("Missing required property 'entity'")
            __props__['entity'] = entity
            if object is None:
                raise TypeError("Missing required property 'object'")
            __props__['object'] = object
            if role is None:
                raise TypeError("Missing required property 'role'")
            __props__['role'] = role
            __props__['domain'] = None
            __props__['email'] = None
            __props__['entity_id'] = None
            __props__['generation'] = None
            __props__['project_team'] = None
        super(ObjectAccessControl, __self__).__init__(
            'gcp:storage/objectAccessControl:ObjectAccessControl',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None, bucket=None, domain=None, email=None, entity=None, entity_id=None, generation=None, object=None, project_team=None, role=None):
        """
        Get an existing ObjectAccessControl resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.
        
        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        
        The **project_team** object supports the following:
        
          * `projectNumber` (`pulumi.Input[str]`)
          * `team` (`pulumi.Input[str]`)

        > This content is derived from https://github.com/terraform-providers/terraform-provider-google/blob/master/website/docs/r/storage_object_access_control.html.markdown.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()
        __props__["bucket"] = bucket
        __props__["domain"] = domain
        __props__["email"] = email
        __props__["entity"] = entity
        __props__["entity_id"] = entity_id
        __props__["generation"] = generation
        __props__["object"] = object
        __props__["project_team"] = project_team
        __props__["role"] = role
        return ObjectAccessControl(resource_name, opts=opts, __props__=__props__)
    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

