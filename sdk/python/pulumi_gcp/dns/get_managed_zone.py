# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from .. import utilities, tables

class GetManagedZoneResult:
    """
    A collection of values returned by getManagedZone.
    """
    def __init__(__self__, description=None, dns_name=None, name=None, name_servers=None, project=None, id=None):
        if description and not isinstance(description, str):
            raise TypeError("Expected argument 'description' to be a str")
        __self__.description = description
        """
        A textual description field.
        """
        if dns_name and not isinstance(dns_name, str):
            raise TypeError("Expected argument 'dns_name' to be a str")
        __self__.dns_name = dns_name
        """
        The fully qualified DNS name of this zone, e.g. `example.com.`.
        """
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        __self__.name = name
        if name_servers and not isinstance(name_servers, list):
            raise TypeError("Expected argument 'name_servers' to be a list")
        __self__.name_servers = name_servers
        """
        The list of nameservers that will be authoritative for this
        domain. Use NS records to redirect from your DNS provider to these names,
        thus making Google Cloud DNS authoritative for this zone.
        """
        if project and not isinstance(project, str):
            raise TypeError("Expected argument 'project' to be a str")
        __self__.project = project
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        __self__.id = id
        """
        id is the provider-assigned unique ID for this managed resource.
        """

    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return self

    __iter__ = __await__

def get_managed_zone(name=None,project=None,opts=None):
    """
    Provides access to a zone's attributes within Google Cloud DNS.
    For more information see
    [the official documentation](https://cloud.google.com/dns/zones/)
    and
    [API](https://cloud.google.com/dns/api/v1/managedZones).

    > This content is derived from https://github.com/terraform-providers/terraform-provider-google/blob/master/website/docs/d/dns_managed_zone.html.markdown.
    """
    __args__ = dict()

    __args__['name'] = name
    __args__['project'] = project
    if opts is None:
        opts = pulumi.ResourceOptions()
    if opts.version is None:
        opts.version = utilities.get_version()
    __ret__ = pulumi.runtime.invoke('gcp:dns/getManagedZone:getManagedZone', __args__, opts=opts).value

    return GetManagedZoneResult(
        description=__ret__.get('description'),
        dns_name=__ret__.get('dnsName'),
        name=__ret__.get('name'),
        name_servers=__ret__.get('nameServers'),
        project=__ret__.get('project'),
        id=__ret__.get('id'))
