# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import pulumi
import pulumi.runtime
from .. import utilities, tables

class TargetHttpsProxy(pulumi.CustomResource):
    """
    Represents a TargetHttpsProxy resource, which is used by one or more
    global forwarding rule to route incoming HTTPS requests to a URL map.
    
    
    To get more information about TargetHttpsProxy, see:
    
    * [API documentation](https://cloud.google.com/compute/docs/reference/latest/targetHttpsProxies)
    * How-to Guides
        * [Official Documentation](https://cloud.google.com/compute/docs/load-balancing/http/target-proxies)
    
    <div class = "oics-button" style="float: right; margin: 0 0 -15px">
      <a href="https://console.cloud.google.com/cloudshell/open?cloudshell_git_repo=https%3A%2F%2Fgithub.com%2Fterraform-google-modules%2Fdocs-examples.git&cloudshell_working_dir=target_https_proxy_basic&cloudshell_image=gcr.io%2Fgraphite-cloud-shell-images%2Fterraform%3Alatest&open_in_editor=main.tf&cloudshell_print=.%2Fmotd&cloudshell_tutorial=.%2Ftutorial.md" target="_blank">
        <img alt="Open in Cloud Shell" src="//gstatic.com/cloudssh/images/open-btn.svg" style="max-height: 44px; margin: 32px auto; max-width: 100%;">
      </a>
    </div>
    """
    def __init__(__self__, __name__, __opts__=None, description=None, name=None, project=None, quic_override=None, ssl_certificates=None, ssl_policy=None, url_map=None):
        """Create a TargetHttpsProxy resource with the given unique name, props, and options."""
        if not __name__:
            raise TypeError('Missing resource name argument (for URN creation)')
        if not isinstance(__name__, str):
            raise TypeError('Expected resource name to be a string')
        if __opts__ and not isinstance(__opts__, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')

        __props__ = dict()

        __props__['description'] = description

        __props__['name'] = name

        __props__['project'] = project

        __props__['quic_override'] = quic_override

        if not ssl_certificates:
            raise TypeError('Missing required property ssl_certificates')
        __props__['ssl_certificates'] = ssl_certificates

        __props__['ssl_policy'] = ssl_policy

        if not url_map:
            raise TypeError('Missing required property url_map')
        __props__['url_map'] = url_map

        __props__['creation_timestamp'] = None
        __props__['proxy_id'] = None
        __props__['self_link'] = None

        super(TargetHttpsProxy, __self__).__init__(
            'gcp:compute/targetHttpsProxy:TargetHttpsProxy',
            __name__,
            __props__,
            __opts__)


    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

