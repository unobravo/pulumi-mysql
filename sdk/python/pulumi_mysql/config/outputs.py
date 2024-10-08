# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = [
    'AwsConfigs',
    'AzureConfigs',
    'CustomTls',
]

@pulumi.output_type
class AwsConfigs(dict):
    def __init__(__self__, *,
                 access_key: Optional[str] = None,
                 profile: Optional[str] = None,
                 region: Optional[str] = None,
                 secret_key: Optional[str] = None):
        if access_key is not None:
            pulumi.set(__self__, "access_key", access_key)
        if profile is not None:
            pulumi.set(__self__, "profile", profile)
        if region is not None:
            pulumi.set(__self__, "region", region)
        if secret_key is not None:
            pulumi.set(__self__, "secret_key", secret_key)

    @property
    @pulumi.getter(name="accessKey")
    def access_key(self) -> Optional[str]:
        return pulumi.get(self, "access_key")

    @property
    @pulumi.getter
    def profile(self) -> Optional[str]:
        return pulumi.get(self, "profile")

    @property
    @pulumi.getter
    def region(self) -> Optional[str]:
        return pulumi.get(self, "region")

    @property
    @pulumi.getter(name="secretKey")
    def secret_key(self) -> Optional[str]:
        return pulumi.get(self, "secret_key")


@pulumi.output_type
class AzureConfigs(dict):
    def __init__(__self__, *,
                 client_id: Optional[str] = None,
                 client_secret: Optional[str] = None,
                 environment: Optional[str] = None,
                 tenant_id: Optional[str] = None):
        if client_id is not None:
            pulumi.set(__self__, "client_id", client_id)
        if client_secret is not None:
            pulumi.set(__self__, "client_secret", client_secret)
        if environment is not None:
            pulumi.set(__self__, "environment", environment)
        if tenant_id is not None:
            pulumi.set(__self__, "tenant_id", tenant_id)

    @property
    @pulumi.getter(name="clientId")
    def client_id(self) -> Optional[str]:
        return pulumi.get(self, "client_id")

    @property
    @pulumi.getter(name="clientSecret")
    def client_secret(self) -> Optional[str]:
        return pulumi.get(self, "client_secret")

    @property
    @pulumi.getter
    def environment(self) -> Optional[str]:
        return pulumi.get(self, "environment")

    @property
    @pulumi.getter(name="tenantId")
    def tenant_id(self) -> Optional[str]:
        return pulumi.get(self, "tenant_id")


@pulumi.output_type
class CustomTls(dict):
    def __init__(__self__, *,
                 ca_cert: str,
                 client_cert: str,
                 client_key: str,
                 config_key: Optional[str] = None):
        pulumi.set(__self__, "ca_cert", ca_cert)
        pulumi.set(__self__, "client_cert", client_cert)
        pulumi.set(__self__, "client_key", client_key)
        if config_key is not None:
            pulumi.set(__self__, "config_key", config_key)

    @property
    @pulumi.getter(name="caCert")
    def ca_cert(self) -> str:
        return pulumi.get(self, "ca_cert")

    @property
    @pulumi.getter(name="clientCert")
    def client_cert(self) -> str:
        return pulumi.get(self, "client_cert")

    @property
    @pulumi.getter(name="clientKey")
    def client_key(self) -> str:
        return pulumi.get(self, "client_key")

    @property
    @pulumi.getter(name="configKey")
    def config_key(self) -> Optional[str]:
        return pulumi.get(self, "config_key")


