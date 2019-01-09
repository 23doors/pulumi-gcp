# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import pulumi
import pulumi.runtime
from .. import utilities, tables

class GetFunctionResult(object):
    """
    A collection of values returned by getFunction.
    """
    def __init__(__self__, available_memory_mb=None, description=None, entry_point=None, environment_variables=None, event_triggers=None, https_trigger_url=None, labels=None, retry_on_failure=None, runtime=None, source_archive_bucket=None, source_archive_object=None, timeout=None, trigger_bucket=None, trigger_http=None, trigger_topic=None, id=None):
        if available_memory_mb and not isinstance(available_memory_mb, int):
            raise TypeError('Expected argument available_memory_mb to be a int')
        __self__.available_memory_mb = available_memory_mb
        """
        Available memory (in MB) to the function.
        """
        if description and not isinstance(description, str):
            raise TypeError('Expected argument description to be a str')
        __self__.description = description
        """
        Description of the function.
        """
        if entry_point and not isinstance(entry_point, str):
            raise TypeError('Expected argument entry_point to be a str')
        __self__.entry_point = entry_point
        """
        Name of a JavaScript function that will be executed when the Google Cloud Function is triggered.
        """
        if environment_variables and not isinstance(environment_variables, dict):
            raise TypeError('Expected argument environment_variables to be a dict')
        __self__.environment_variables = environment_variables
        if event_triggers and not isinstance(event_triggers, list):
            raise TypeError('Expected argument event_triggers to be a list')
        __self__.event_triggers = event_triggers
        """
        A source that fires events in response to a condition in another service. Structure is documented below.
        """
        if https_trigger_url and not isinstance(https_trigger_url, str):
            raise TypeError('Expected argument https_trigger_url to be a str')
        __self__.https_trigger_url = https_trigger_url
        """
        If function is triggered by HTTP, trigger URL is set here.
        """
        if labels and not isinstance(labels, dict):
            raise TypeError('Expected argument labels to be a dict')
        __self__.labels = labels
        """
        A map of labels applied to this function.
        """
        if retry_on_failure and not isinstance(retry_on_failure, bool):
            raise TypeError('Expected argument retry_on_failure to be a bool')
        __self__.retry_on_failure = retry_on_failure
        if runtime and not isinstance(runtime, str):
            raise TypeError('Expected argument runtime to be a str')
        __self__.runtime = runtime
        """
        The runtime in which the function is running.
        """
        if source_archive_bucket and not isinstance(source_archive_bucket, str):
            raise TypeError('Expected argument source_archive_bucket to be a str')
        __self__.source_archive_bucket = source_archive_bucket
        """
        The GCS bucket containing the zip archive which contains the function.
        """
        if source_archive_object and not isinstance(source_archive_object, str):
            raise TypeError('Expected argument source_archive_object to be a str')
        __self__.source_archive_object = source_archive_object
        """
        The source archive object (file) in archive bucket.
        """
        if timeout and not isinstance(timeout, int):
            raise TypeError('Expected argument timeout to be a int')
        __self__.timeout = timeout
        """
        Function execution timeout (in seconds).
        """
        if trigger_bucket and not isinstance(trigger_bucket, str):
            raise TypeError('Expected argument trigger_bucket to be a str')
        __self__.trigger_bucket = trigger_bucket
        """
        If function is triggered by bucket, bucket name is set here. Deprecated. Use `event_trigger` instead.
        """
        if trigger_http and not isinstance(trigger_http, bool):
            raise TypeError('Expected argument trigger_http to be a bool')
        __self__.trigger_http = trigger_http
        """
        If function is triggered by HTTP, this boolean is set.
        """
        if trigger_topic and not isinstance(trigger_topic, str):
            raise TypeError('Expected argument trigger_topic to be a str')
        __self__.trigger_topic = trigger_topic
        """
        If function is triggered by Pub/Sub topic, name of topic is set here. Deprecated. Use `event_trigger` instead.
        """
        if id and not isinstance(id, str):
            raise TypeError('Expected argument id to be a str')
        __self__.id = id
        """
        id is the provider-assigned unique ID for this managed resource.
        """

async def get_function(name=None, project=None, region=None):
    """
    Get information about a Google Cloud Function. For more information see
    the [official documentation](https://cloud.google.com/functions/docs/)
    and [API](https://cloud.google.com/functions/docs/apis).
    """
    __args__ = dict()

    __args__['name'] = name
    __args__['project'] = project
    __args__['region'] = region
    __ret__ = await pulumi.runtime.invoke('gcp:cloudfunctions/getFunction:getFunction', __args__)

    return GetFunctionResult(
        available_memory_mb=__ret__.get('availableMemoryMb'),
        description=__ret__.get('description'),
        entry_point=__ret__.get('entryPoint'),
        environment_variables=__ret__.get('environmentVariables'),
        event_triggers=__ret__.get('eventTriggers'),
        https_trigger_url=__ret__.get('httpsTriggerUrl'),
        labels=__ret__.get('labels'),
        retry_on_failure=__ret__.get('retryOnFailure'),
        runtime=__ret__.get('runtime'),
        source_archive_bucket=__ret__.get('sourceArchiveBucket'),
        source_archive_object=__ret__.get('sourceArchiveObject'),
        timeout=__ret__.get('timeout'),
        trigger_bucket=__ret__.get('triggerBucket'),
        trigger_http=__ret__.get('triggerHttp'),
        trigger_topic=__ret__.get('triggerTopic'),
        id=__ret__.get('id'))
