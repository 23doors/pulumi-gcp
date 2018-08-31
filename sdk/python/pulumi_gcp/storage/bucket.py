# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import pulumi
import pulumi.runtime

class Bucket(pulumi.CustomResource):
    """
    Creates a new bucket in Google cloud storage service (GCS).
    Once a bucket has been created, its location can't be changed.
    [ACLs](https://cloud.google.com/storage/docs/access-control/lists) can be applied using the `google_storage_bucket_acl` resource.
    For more information see
    [the official documentation](https://cloud.google.com/storage/docs/overview)
    and
    [API](https://cloud.google.com/storage/docs/json_api/v1/buckets).
    
    """
    def __init__(__self__, __name__, __opts__=None, cors=None, force_destroy=None, labels=None, lifecycle_rules=None, location=None, logging=None, name=None, project=None, storage_class=None, versioning=None, websites=None):
        """Create a Bucket resource with the given unique name, props, and options."""
        if not __name__:
            raise TypeError('Missing resource name argument (for URN creation)')
        if not isinstance(__name__, basestring):
            raise TypeError('Expected resource name to be a string')
        if __opts__ and not isinstance(__opts__, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')

        __props__ = dict()

        if cors and not isinstance(cors, list):
            raise TypeError('Expected property cors to be a list')
        __self__.cors = cors
        """
        The bucket's [Cross-Origin Resource Sharing (CORS)](https://www.w3.org/TR/cors/) configuration. Multiple blocks of this type are permitted. Structure is documented below.
        """
        __props__['cors'] = cors

        if force_destroy and not isinstance(force_destroy, bool):
            raise TypeError('Expected property force_destroy to be a bool')
        __self__.force_destroy = force_destroy
        """
        When deleting a bucket, this
        boolean option will delete all contained objects. If you try to delete a
        bucket that contains objects, Terraform will fail that run.
        """
        __props__['forceDestroy'] = force_destroy

        if labels and not isinstance(labels, dict):
            raise TypeError('Expected property labels to be a dict')
        __self__.labels = labels
        """
        A set of key/value label pairs to assign to the bucket.
        """
        __props__['labels'] = labels

        if lifecycle_rules and not isinstance(lifecycle_rules, list):
            raise TypeError('Expected property lifecycle_rules to be a list')
        __self__.lifecycle_rules = lifecycle_rules
        """
        The bucket's [Lifecycle Rules](https://cloud.google.com/storage/docs/lifecycle#configuration) configuration. Multiple blocks of this type are permitted. Structure is documented below.
        """
        __props__['lifecycleRules'] = lifecycle_rules

        if location and not isinstance(location, basestring):
            raise TypeError('Expected property location to be a basestring')
        __self__.location = location
        """
        The [GCS location](https://cloud.google.com/storage/docs/bucket-locations)
        """
        __props__['location'] = location

        if logging and not isinstance(logging, dict):
            raise TypeError('Expected property logging to be a dict')
        __self__.logging = logging
        """
        The bucket's [Access & Storage Logs](https://cloud.google.com/storage/docs/access-logs) configuration.
        """
        __props__['logging'] = logging

        if name and not isinstance(name, basestring):
            raise TypeError('Expected property name to be a basestring')
        __self__.name = name
        """
        The name of the bucket.
        """
        __props__['name'] = name

        if project and not isinstance(project, basestring):
            raise TypeError('Expected property project to be a basestring')
        __self__.project = project
        """
        The ID of the project in which the resource belongs. If it
        is not provided, the provider project is used.
        """
        __props__['project'] = project

        if storage_class and not isinstance(storage_class, basestring):
            raise TypeError('Expected property storage_class to be a basestring')
        __self__.storage_class = storage_class
        """
        The [Storage Class](https://cloud.google.com/storage/docs/storage-classes) of the new bucket. Supported values include: `MULTI_REGIONAL`, `REGIONAL`, `NEARLINE`, `COLDLINE`.
        """
        __props__['storageClass'] = storage_class

        if versioning and not isinstance(versioning, dict):
            raise TypeError('Expected property versioning to be a dict')
        __self__.versioning = versioning
        """
        The bucket's [Versioning](https://cloud.google.com/storage/docs/object-versioning) configuration.
        """
        __props__['versioning'] = versioning

        if websites and not isinstance(websites, list):
            raise TypeError('Expected property websites to be a list')
        __self__.websites = websites
        """
        Configuration if the bucket acts as a website. Structure is documented below.
        """
        __props__['websites'] = websites

        __self__.self_link = pulumi.runtime.UNKNOWN
        """
        The URI of the created resource.
        """
        __self__.url = pulumi.runtime.UNKNOWN
        """
        The base URL of the bucket, in the format `gs://<bucket-name>`.
        """

        super(Bucket, __self__).__init__(
            'gcp:storage/bucket:Bucket',
            __name__,
            __props__,
            __opts__)

    def set_outputs(self, outs):
        if 'cors' in outs:
            self.cors = outs['cors']
        if 'forceDestroy' in outs:
            self.force_destroy = outs['forceDestroy']
        if 'labels' in outs:
            self.labels = outs['labels']
        if 'lifecycleRules' in outs:
            self.lifecycle_rules = outs['lifecycleRules']
        if 'location' in outs:
            self.location = outs['location']
        if 'logging' in outs:
            self.logging = outs['logging']
        if 'name' in outs:
            self.name = outs['name']
        if 'project' in outs:
            self.project = outs['project']
        if 'selfLink' in outs:
            self.self_link = outs['selfLink']
        if 'storageClass' in outs:
            self.storage_class = outs['storageClass']
        if 'url' in outs:
            self.url = outs['url']
        if 'versioning' in outs:
            self.versioning = outs['versioning']
        if 'websites' in outs:
            self.websites = outs['websites']
