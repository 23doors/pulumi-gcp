# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from .. import utilities, tables

class OrganizationPolicy(pulumi.CustomResource):
    boolean_policy: pulumi.Output[dict]
    """
    A boolean policy is a constraint that is either enforced or not. Structure is documented below.
    """
    constraint: pulumi.Output[str]
    """
    The name of the Constraint the Policy is configuring, for example, `serviceuser.services`. Check out the [complete list of available constraints](https://cloud.google.com/resource-manager/docs/organization-policy/understanding-constraints#available_constraints).
    """
    etag: pulumi.Output[str]
    """
    (Computed) The etag of the organization policy. `etag` is used for optimistic concurrency control as a way to help prevent simultaneous updates of a policy from overwriting each other.
    """
    list_policy: pulumi.Output[dict]
    """
    A policy that can define specific values that are allowed or denied for the given constraint. It can also be used to allow or deny all values. Structure is documented below.
    """
    project: pulumi.Output[str]
    """
    The project id of the project to set the policy for.
    """
    restore_policy: pulumi.Output[dict]
    """
    A restore policy is a constraint to restore the default policy. Structure is documented below.
    """
    update_time: pulumi.Output[str]
    """
    (Computed) The timestamp in RFC3339 UTC "Zulu" format, accurate to nanoseconds, representing when the variable was last updated. Example: "2016-10-09T12:33:37.578138407Z".
    """
    version: pulumi.Output[float]
    """
    Version of the Policy. Default version is 0.
    """
    def __init__(__self__, resource_name, opts=None, boolean_policy=None, constraint=None, list_policy=None, project=None, restore_policy=None, version=None, __name__=None, __opts__=None):
        """
        Allows management of Organization policies for a Google Project. For more information see
        [the official
        documentation](https://cloud.google.com/resource-manager/docs/organization-policy/overview) and
        [API](https://cloud.google.com/resource-manager/reference/rest/v1/projects/setOrgPolicy).
        
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[dict] boolean_policy: A boolean policy is a constraint that is either enforced or not. Structure is documented below.
        :param pulumi.Input[str] constraint: The name of the Constraint the Policy is configuring, for example, `serviceuser.services`. Check out the [complete list of available constraints](https://cloud.google.com/resource-manager/docs/organization-policy/understanding-constraints#available_constraints).
        :param pulumi.Input[dict] list_policy: A policy that can define specific values that are allowed or denied for the given constraint. It can also be used to allow or deny all values. Structure is documented below.
        :param pulumi.Input[str] project: The project id of the project to set the policy for.
        :param pulumi.Input[dict] restore_policy: A restore policy is a constraint to restore the default policy. Structure is documented below.
        :param pulumi.Input[float] version: Version of the Policy. Default version is 0.
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

        __props__['boolean_policy'] = boolean_policy

        if constraint is None:
            raise TypeError('Missing required property constraint')
        __props__['constraint'] = constraint

        __props__['list_policy'] = list_policy

        if project is None:
            raise TypeError('Missing required property project')
        __props__['project'] = project

        __props__['restore_policy'] = restore_policy

        __props__['version'] = version

        __props__['etag'] = None
        __props__['update_time'] = None

        super(OrganizationPolicy, __self__).__init__(
            'gcp:projects/organizationPolicy:OrganizationPolicy',
            resource_name,
            __props__,
            opts)


    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

