# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import utilities, tables

class GetCaCertsResult:
    """
    A collection of values returned by getCaCerts.
    """
    def __init__(__self__, active_version=None, certs=None, instance=None, project=None, id=None):
        if active_version and not isinstance(active_version, str):
            raise TypeError("Expected argument 'active_version' to be a str")
        __self__.active_version = active_version
        if certs and not isinstance(certs, list):
            raise TypeError("Expected argument 'certs' to be a list")
        __self__.certs = certs
        if instance and not isinstance(instance, str):
            raise TypeError("Expected argument 'instance' to be a str")
        __self__.instance = instance
        if project and not isinstance(project, str):
            raise TypeError("Expected argument 'project' to be a str")
        __self__.project = project
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        __self__.id = id
        """
        id is the provider-assigned unique ID for this managed resource.
        """
class AwaitableGetCaCertsResult(GetCaCertsResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetCaCertsResult(
            active_version=self.active_version,
            certs=self.certs,
            instance=self.instance,
            project=self.project,
            id=self.id)

def get_ca_certs(instance=None,project=None,opts=None):
    """
    Use this data source to access information about an existing resource.
    
    """
    __args__ = dict()

    __args__['instance'] = instance
    __args__['project'] = project
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = utilities.get_version()
    __ret__ = pulumi.runtime.invoke('gcp:sql/getCaCerts:getCaCerts', __args__, opts=opts).value

    return AwaitableGetCaCertsResult(
        active_version=__ret__.get('activeVersion'),
        certs=__ret__.get('certs'),
        instance=__ret__.get('instance'),
        project=__ret__.get('project'),
        id=__ret__.get('id'))
