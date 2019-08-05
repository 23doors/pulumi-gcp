# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from .. import utilities, tables

class ClusterIAMPolicy(pulumi.CustomResource):
    cluster: pulumi.Output[str]
    """
    The name or relative resource id of the cluster to manage IAM policies for.
    """
    etag: pulumi.Output[str]
    """
    (Computed) The etag of the clusters's IAM policy.
    """
    policy_data: pulumi.Output[str]
    """
    The policy data generated by a `google_iam_policy` data source.
    """
    project: pulumi.Output[str]
    """
    The project in which the cluster belongs. If it
    is not provided, this provider will use the provider default.
    """
    region: pulumi.Output[str]
    """
    The region in which the cluster belongs. If it
    is not provided, this provider will use the provider default.
    """
    def __init__(__self__, resource_name, opts=None, cluster=None, policy_data=None, project=None, region=None, __name__=None, __opts__=None):
        """
        Three different resources help you manage IAM policies on dataproc clusters. Each of these resources serves a different use case:
        
        * `google_dataproc_cluster_iam_policy`: Authoritative. Sets the IAM policy for the cluster and replaces any existing policy already attached.
        * `google_dataproc_cluster_iam_binding`: Authoritative for a given role. Updates the IAM policy to grant a role to a list of members. Other roles within the IAM policy for the cluster are preserved.
        * `google_dataproc_cluster_iam_member`: Non-authoritative. Updates the IAM policy to grant a role to a new member. Other members for the role for the cluster are preserved.
        
        > **Note:** `google_dataproc_cluster_iam_policy` **cannot** be used in conjunction with `google_dataproc_cluster_iam_binding` and `google_dataproc_cluster_iam_member` or they will fight over what your policy should be. In addition, be careful not to accidentaly unset ownership of the cluster as `google_dataproc_cluster_iam_policy` replaces the entire policy.
        
        > **Note:** `google_dataproc_cluster_iam_binding` resources **can be** used in conjunction with `google_dataproc_cluster_iam_member` resources **only if** they do not grant privilege to the same role.
        
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] cluster: The name or relative resource id of the cluster to manage IAM policies for.
        :param pulumi.Input[str] policy_data: The policy data generated by a `google_iam_policy` data source.
        :param pulumi.Input[str] project: The project in which the cluster belongs. If it
               is not provided, this provider will use the provider default.
        :param pulumi.Input[str] region: The region in which the cluster belongs. If it
               is not provided, this provider will use the provider default.

        > This content is derived from https://github.com/terraform-providers/terraform-provider-google/blob/master/website/docs/r/dataproc_cluster_iam_policy.html.markdown.
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

        if cluster is None:
            raise TypeError("Missing required property 'cluster'")
        __props__['cluster'] = cluster

        if policy_data is None:
            raise TypeError("Missing required property 'policy_data'")
        __props__['policy_data'] = policy_data

        __props__['project'] = project

        __props__['region'] = region

        __props__['etag'] = None

        if opts is None:
            opts = pulumi.ResourceOptions()
        if opts.version is None:
            opts.version = utilities.get_version()
        super(ClusterIAMPolicy, __self__).__init__(
            'gcp:dataproc/clusterIAMPolicy:ClusterIAMPolicy',
            resource_name,
            __props__,
            opts)


    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

