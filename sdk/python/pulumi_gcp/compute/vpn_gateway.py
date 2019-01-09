# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import pulumi
import pulumi.runtime
from .. import utilities, tables

class VPNGateway(pulumi.CustomResource):
    """
    Represents a VPN gateway running in GCP. This virtual device is managed
    by Google, but used only by you.
    
    
    To get more information about VpnGateway, see:
    
    * [API documentation](https://cloud.google.com/compute/docs/reference/rest/v1/targetVpnGateways)
    
    <div class = "oics-button" style="float: right; margin: 0 0 -15px">
      <a href="https://console.cloud.google.com/cloudshell/open?cloudshell_git_repo=https%3A%2F%2Fgithub.com%2Fterraform-google-modules%2Fdocs-examples.git&cloudshell_working_dir=target_vpn_gateway_basic&cloudshell_image=gcr.io%2Fgraphite-cloud-shell-images%2Fterraform%3Alatest&open_in_editor=main.tf&cloudshell_print=.%2Fmotd&cloudshell_tutorial=.%2Ftutorial.md" target="_blank">
        <img alt="Open in Cloud Shell" src="//gstatic.com/cloudssh/images/open-btn.svg" style="max-height: 44px; margin: 32px auto; max-width: 100%;">
      </a>
    </div>
    """
    def __init__(__self__, __name__, __opts__=None, description=None, name=None, network=None, project=None, region=None):
        """Create a VPNGateway resource with the given unique name, props, and options."""
        if not __name__:
            raise TypeError('Missing resource name argument (for URN creation)')
        if not isinstance(__name__, str):
            raise TypeError('Expected resource name to be a string')
        if __opts__ and not isinstance(__opts__, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')

        __props__ = dict()

        __props__['description'] = description

        __props__['name'] = name

        if not network:
            raise TypeError('Missing required property network')
        __props__['network'] = network

        __props__['project'] = project

        __props__['region'] = region

        __props__['creation_timestamp'] = None
        __props__['self_link'] = None

        super(VPNGateway, __self__).__init__(
            'gcp:compute/vPNGateway:VPNGateway',
            __name__,
            __props__,
            __opts__)


    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

