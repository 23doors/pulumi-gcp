# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import utilities, tables

class RegionInstanceGroupManager(pulumi.CustomResource):
    auto_healing_policies: pulumi.Output[dict]
    """
    The autohealing policies for this managed instance
    group. You can specify only one value. Structure is documented below. For more information, see the [official documentation](https://cloud.google.com/compute/docs/instance-groups/creating-groups-of-managed-instances#monitoring_groups).
    
      * `healthCheck` (`str`)
      * `initialDelaySec` (`float`)
    """
    base_instance_name: pulumi.Output[str]
    """
    The base instance name to use for
    instances in this group. The value must be a valid
    [RFC1035](https://www.ietf.org/rfc/rfc1035.txt) name. Supported characters
    are lowercase letters, numbers, and hyphens (-). Instances are named by
    appending a hyphen and a random four-character string to the base instance
    name.
    """
    description: pulumi.Output[str]
    """
    An optional textual description of the instance
    group manager.
    """
    distribution_policy_zones: pulumi.Output[list]
    """
    The distribution policy for this managed instance
    group. You can specify one or more values. For more information, see the [official documentation](https://cloud.google.com/compute/docs/instance-groups/distributing-instances-with-regional-instance-groups#selectingzones).
    - - -
    """
    fingerprint: pulumi.Output[str]
    """
    The fingerprint of the instance group manager.
    """
    instance_group: pulumi.Output[str]
    """
    The full URL of the instance group created by the manager.
    """
    name: pulumi.Output[str]
    """
    The name of the instance group manager. Must be 1-63
    characters long and comply with
    [RFC1035](https://www.ietf.org/rfc/rfc1035.txt). Supported characters
    include lowercase letters, numbers, and hyphens.
    """
    named_ports: pulumi.Output[list]
    """
    The named port configuration. See the section below
    for details on configuration.
    
      * `name` (`str`) - The name of the instance group manager. Must be 1-63
        characters long and comply with
        [RFC1035](https://www.ietf.org/rfc/rfc1035.txt). Supported characters
        include lowercase letters, numbers, and hyphens.
      * `port` (`float`)
    """
    project: pulumi.Output[str]
    """
    The ID of the project in which the resource belongs. If it
    is not provided, the provider project is used.
    """
    region: pulumi.Output[str]
    """
    The region where the managed instance group resides.
    """
    self_link: pulumi.Output[str]
    """
    The URL of the created resource.
    """
    target_pools: pulumi.Output[list]
    """
    The full URL of all target pools to which new
    instances in the group are added. Updating the target pools attribute does
    not affect existing instances.
    """
    target_size: pulumi.Output[float]
    """
    The target number of running instances for this managed
    instance group. This value should always be explicitly set unless this resource is attached to
    an autoscaler, in which case it should never be set. Defaults to `0`.
    """
    update_policy: pulumi.Output[dict]
    """
    The update policy for this managed instance group. Structure is documented below. For more information, see the [official documentation](https://cloud.google.com/compute/docs/instance-groups/updating-managed-instance-groups) and [API](https://cloud.google.com/compute/docs/reference/rest/beta/regionInstanceGroupManagers/patch)
    
      * `instanceRedistributionType` (`str`)
      * `maxSurgeFixed` (`float`)
      * `maxSurgePercent` (`float`)
      * `maxUnavailableFixed` (`float`)
      * `maxUnavailablePercent` (`float`)
      * `minReadySec` (`float`)
      * `minimalAction` (`str`)
      * `type` (`str`)
    """
    versions: pulumi.Output[list]
    """
    Application versions managed by this instance group. Each
    version deals with a specific instance template, allowing canary release scenarios.
    Structure is documented below.
    
      * `instanceTemplate` (`str`)
      * `name` (`str`) - The name of the instance group manager. Must be 1-63
        characters long and comply with
        [RFC1035](https://www.ietf.org/rfc/rfc1035.txt). Supported characters
        include lowercase letters, numbers, and hyphens.
      * `target_size` (`dict`) - The target number of running instances for this managed
        instance group. This value should always be explicitly set unless this resource is attached to
        an autoscaler, in which case it should never be set. Defaults to `0`.
    
        * `fixed` (`float`)
        * `percent` (`float`)
    """
    wait_for_instances: pulumi.Output[bool]
    """
    Whether to wait for all instances to be created/updated before
    returning. Note that if this is set to true and the operation does not succeed, this provider will
    continue trying until it times out.
    """
    def __init__(__self__, resource_name, opts=None, auto_healing_policies=None, base_instance_name=None, description=None, distribution_policy_zones=None, name=None, named_ports=None, project=None, region=None, target_pools=None, target_size=None, update_policy=None, versions=None, wait_for_instances=None, __props__=None, __name__=None, __opts__=None):
        """
        The Google Compute Engine Regional Instance Group Manager API creates and manages pools
        of homogeneous Compute Engine virtual machine instances from a common instance
        template. For more information, see [the official documentation](https://cloud.google.com/compute/docs/instance-groups/distributing-instances-with-regional-instance-groups)
        and [API](https://cloud.google.com/compute/docs/reference/latest/regionInstanceGroupManagers)
        
        > **Note:** Use [compute.InstanceGroupManager](https://www.terraform.io/docs/providers/google/r/compute_instance_group_manager.html) to create a single-zone instance group manager.
        
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[dict] auto_healing_policies: The autohealing policies for this managed instance
               group. You can specify only one value. Structure is documented below. For more information, see the [official documentation](https://cloud.google.com/compute/docs/instance-groups/creating-groups-of-managed-instances#monitoring_groups).
        :param pulumi.Input[str] base_instance_name: The base instance name to use for
               instances in this group. The value must be a valid
               [RFC1035](https://www.ietf.org/rfc/rfc1035.txt) name. Supported characters
               are lowercase letters, numbers, and hyphens (-). Instances are named by
               appending a hyphen and a random four-character string to the base instance
               name.
        :param pulumi.Input[str] description: An optional textual description of the instance
               group manager.
        :param pulumi.Input[list] distribution_policy_zones: The distribution policy for this managed instance
               group. You can specify one or more values. For more information, see the [official documentation](https://cloud.google.com/compute/docs/instance-groups/distributing-instances-with-regional-instance-groups#selectingzones).
               - - -
        :param pulumi.Input[str] name: The name of the instance group manager. Must be 1-63
               characters long and comply with
               [RFC1035](https://www.ietf.org/rfc/rfc1035.txt). Supported characters
               include lowercase letters, numbers, and hyphens.
        :param pulumi.Input[list] named_ports: The named port configuration. See the section below
               for details on configuration.
        :param pulumi.Input[str] project: The ID of the project in which the resource belongs. If it
               is not provided, the provider project is used.
        :param pulumi.Input[str] region: The region where the managed instance group resides.
        :param pulumi.Input[list] target_pools: The full URL of all target pools to which new
               instances in the group are added. Updating the target pools attribute does
               not affect existing instances.
        :param pulumi.Input[float] target_size: The target number of running instances for this managed
               instance group. This value should always be explicitly set unless this resource is attached to
               an autoscaler, in which case it should never be set. Defaults to `0`.
        :param pulumi.Input[dict] update_policy: The update policy for this managed instance group. Structure is documented below. For more information, see the [official documentation](https://cloud.google.com/compute/docs/instance-groups/updating-managed-instance-groups) and [API](https://cloud.google.com/compute/docs/reference/rest/beta/regionInstanceGroupManagers/patch)
        :param pulumi.Input[list] versions: Application versions managed by this instance group. Each
               version deals with a specific instance template, allowing canary release scenarios.
               Structure is documented below.
        :param pulumi.Input[bool] wait_for_instances: Whether to wait for all instances to be created/updated before
               returning. Note that if this is set to true and the operation does not succeed, this provider will
               continue trying until it times out.
        
        The **auto_healing_policies** object supports the following:
        
          * `healthCheck` (`pulumi.Input[str]`)
          * `initialDelaySec` (`pulumi.Input[float]`)
        
        The **named_ports** object supports the following:
        
          * `name` (`pulumi.Input[str]`) - The name of the instance group manager. Must be 1-63
            characters long and comply with
            [RFC1035](https://www.ietf.org/rfc/rfc1035.txt). Supported characters
            include lowercase letters, numbers, and hyphens.
          * `port` (`pulumi.Input[float]`)
        
        The **update_policy** object supports the following:
        
          * `instanceRedistributionType` (`pulumi.Input[str]`)
          * `maxSurgeFixed` (`pulumi.Input[float]`)
          * `maxSurgePercent` (`pulumi.Input[float]`)
          * `maxUnavailableFixed` (`pulumi.Input[float]`)
          * `maxUnavailablePercent` (`pulumi.Input[float]`)
          * `minReadySec` (`pulumi.Input[float]`)
          * `minimalAction` (`pulumi.Input[str]`)
          * `type` (`pulumi.Input[str]`)
        
        The **versions** object supports the following:
        
          * `instanceTemplate` (`pulumi.Input[str]`)
          * `name` (`pulumi.Input[str]`) - The name of the instance group manager. Must be 1-63
            characters long and comply with
            [RFC1035](https://www.ietf.org/rfc/rfc1035.txt). Supported characters
            include lowercase letters, numbers, and hyphens.
          * `target_size` (`pulumi.Input[dict]`) - The target number of running instances for this managed
            instance group. This value should always be explicitly set unless this resource is attached to
            an autoscaler, in which case it should never be set. Defaults to `0`.
        
            * `fixed` (`pulumi.Input[float]`)
            * `percent` (`pulumi.Input[float]`)

        > This content is derived from https://github.com/terraform-providers/terraform-provider-google/blob/master/website/docs/r/compute_region_instance_group_manager.html.markdown.
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

            __props__['auto_healing_policies'] = auto_healing_policies
            if base_instance_name is None:
                raise TypeError("Missing required property 'base_instance_name'")
            __props__['base_instance_name'] = base_instance_name
            __props__['description'] = description
            __props__['distribution_policy_zones'] = distribution_policy_zones
            __props__['name'] = name
            __props__['named_ports'] = named_ports
            __props__['project'] = project
            if region is None:
                raise TypeError("Missing required property 'region'")
            __props__['region'] = region
            __props__['target_pools'] = target_pools
            __props__['target_size'] = target_size
            __props__['update_policy'] = update_policy
            if versions is None:
                raise TypeError("Missing required property 'versions'")
            __props__['versions'] = versions
            __props__['wait_for_instances'] = wait_for_instances
            __props__['fingerprint'] = None
            __props__['instance_group'] = None
            __props__['self_link'] = None
        super(RegionInstanceGroupManager, __self__).__init__(
            'gcp:compute/regionInstanceGroupManager:RegionInstanceGroupManager',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None, auto_healing_policies=None, base_instance_name=None, description=None, distribution_policy_zones=None, fingerprint=None, instance_group=None, name=None, named_ports=None, project=None, region=None, self_link=None, target_pools=None, target_size=None, update_policy=None, versions=None, wait_for_instances=None):
        """
        Get an existing RegionInstanceGroupManager resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.
        
        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[dict] auto_healing_policies: The autohealing policies for this managed instance
               group. You can specify only one value. Structure is documented below. For more information, see the [official documentation](https://cloud.google.com/compute/docs/instance-groups/creating-groups-of-managed-instances#monitoring_groups).
        :param pulumi.Input[str] base_instance_name: The base instance name to use for
               instances in this group. The value must be a valid
               [RFC1035](https://www.ietf.org/rfc/rfc1035.txt) name. Supported characters
               are lowercase letters, numbers, and hyphens (-). Instances are named by
               appending a hyphen and a random four-character string to the base instance
               name.
        :param pulumi.Input[str] description: An optional textual description of the instance
               group manager.
        :param pulumi.Input[list] distribution_policy_zones: The distribution policy for this managed instance
               group. You can specify one or more values. For more information, see the [official documentation](https://cloud.google.com/compute/docs/instance-groups/distributing-instances-with-regional-instance-groups#selectingzones).
               - - -
        :param pulumi.Input[str] fingerprint: The fingerprint of the instance group manager.
        :param pulumi.Input[str] instance_group: The full URL of the instance group created by the manager.
        :param pulumi.Input[str] name: The name of the instance group manager. Must be 1-63
               characters long and comply with
               [RFC1035](https://www.ietf.org/rfc/rfc1035.txt). Supported characters
               include lowercase letters, numbers, and hyphens.
        :param pulumi.Input[list] named_ports: The named port configuration. See the section below
               for details on configuration.
        :param pulumi.Input[str] project: The ID of the project in which the resource belongs. If it
               is not provided, the provider project is used.
        :param pulumi.Input[str] region: The region where the managed instance group resides.
        :param pulumi.Input[str] self_link: The URL of the created resource.
        :param pulumi.Input[list] target_pools: The full URL of all target pools to which new
               instances in the group are added. Updating the target pools attribute does
               not affect existing instances.
        :param pulumi.Input[float] target_size: The target number of running instances for this managed
               instance group. This value should always be explicitly set unless this resource is attached to
               an autoscaler, in which case it should never be set. Defaults to `0`.
        :param pulumi.Input[dict] update_policy: The update policy for this managed instance group. Structure is documented below. For more information, see the [official documentation](https://cloud.google.com/compute/docs/instance-groups/updating-managed-instance-groups) and [API](https://cloud.google.com/compute/docs/reference/rest/beta/regionInstanceGroupManagers/patch)
        :param pulumi.Input[list] versions: Application versions managed by this instance group. Each
               version deals with a specific instance template, allowing canary release scenarios.
               Structure is documented below.
        :param pulumi.Input[bool] wait_for_instances: Whether to wait for all instances to be created/updated before
               returning. Note that if this is set to true and the operation does not succeed, this provider will
               continue trying until it times out.
        
        The **auto_healing_policies** object supports the following:
        
          * `healthCheck` (`pulumi.Input[str]`)
          * `initialDelaySec` (`pulumi.Input[float]`)
        
        The **named_ports** object supports the following:
        
          * `name` (`pulumi.Input[str]`) - The name of the instance group manager. Must be 1-63
            characters long and comply with
            [RFC1035](https://www.ietf.org/rfc/rfc1035.txt). Supported characters
            include lowercase letters, numbers, and hyphens.
          * `port` (`pulumi.Input[float]`)
        
        The **update_policy** object supports the following:
        
          * `instanceRedistributionType` (`pulumi.Input[str]`)
          * `maxSurgeFixed` (`pulumi.Input[float]`)
          * `maxSurgePercent` (`pulumi.Input[float]`)
          * `maxUnavailableFixed` (`pulumi.Input[float]`)
          * `maxUnavailablePercent` (`pulumi.Input[float]`)
          * `minReadySec` (`pulumi.Input[float]`)
          * `minimalAction` (`pulumi.Input[str]`)
          * `type` (`pulumi.Input[str]`)
        
        The **versions** object supports the following:
        
          * `instanceTemplate` (`pulumi.Input[str]`)
          * `name` (`pulumi.Input[str]`) - The name of the instance group manager. Must be 1-63
            characters long and comply with
            [RFC1035](https://www.ietf.org/rfc/rfc1035.txt). Supported characters
            include lowercase letters, numbers, and hyphens.
          * `target_size` (`pulumi.Input[dict]`) - The target number of running instances for this managed
            instance group. This value should always be explicitly set unless this resource is attached to
            an autoscaler, in which case it should never be set. Defaults to `0`.
        
            * `fixed` (`pulumi.Input[float]`)
            * `percent` (`pulumi.Input[float]`)

        > This content is derived from https://github.com/terraform-providers/terraform-provider-google/blob/master/website/docs/r/compute_region_instance_group_manager.html.markdown.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()
        __props__["auto_healing_policies"] = auto_healing_policies
        __props__["base_instance_name"] = base_instance_name
        __props__["description"] = description
        __props__["distribution_policy_zones"] = distribution_policy_zones
        __props__["fingerprint"] = fingerprint
        __props__["instance_group"] = instance_group
        __props__["name"] = name
        __props__["named_ports"] = named_ports
        __props__["project"] = project
        __props__["region"] = region
        __props__["self_link"] = self_link
        __props__["target_pools"] = target_pools
        __props__["target_size"] = target_size
        __props__["update_policy"] = update_policy
        __props__["versions"] = versions
        __props__["wait_for_instances"] = wait_for_instances
        return RegionInstanceGroupManager(resource_name, opts=opts, __props__=__props__)
    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

