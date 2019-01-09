# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import pulumi
import pulumi.runtime
from .. import utilities, tables

class GetRuleResult(object):
    """
    A collection of values returned by getRule.
    """
    def __init__(__self__, included_permissions=None, stage=None, title=None, id=None):
        if included_permissions and not isinstance(included_permissions, list):
            raise TypeError('Expected argument included_permissions to be a list')
        __self__.included_permissions = included_permissions
        """
        specifies the list of one or more permissions to include in the custom role, such as - `iam.roles.get`
        """
        if stage and not isinstance(stage, str):
            raise TypeError('Expected argument stage to be a str')
        __self__.stage = stage
        """
        indicates the stage of a role in the launch lifecycle, such as `GA`, `BETA` or `ALPHA`.
        """
        if title and not isinstance(title, str):
            raise TypeError('Expected argument title to be a str')
        __self__.title = title
        """
        is a friendly title for the role, such as "Role Viewer"
        """
        if id and not isinstance(id, str):
            raise TypeError('Expected argument id to be a str')
        __self__.id = id
        """
        id is the provider-assigned unique ID for this managed resource.
        """

async def get_rule(name=None):
    """
    Use this data source to get information about a Google IAM Role.
    
    ```hcl
    data "google_iam_role" "roleinfo" {
      name = "roles/compute.viewer"
    }
    
    output "the_role_permissions" {
      value = "${data.google_iam_role.roleinfo.included_permissions}"
    }
    
    ```
    """
    __args__ = dict()

    __args__['name'] = name
    __ret__ = await pulumi.runtime.invoke('gcp:iam/getRule:getRule', __args__)

    return GetRuleResult(
        included_permissions=__ret__.get('includedPermissions'),
        stage=__ret__.get('stage'),
        title=__ret__.get('title'),
        id=__ret__.get('id'))
