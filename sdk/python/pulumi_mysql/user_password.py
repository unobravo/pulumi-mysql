# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = ['UserPasswordArgs', 'UserPassword']

@pulumi.input_type
class UserPasswordArgs:
    def __init__(__self__, *,
                 user: pulumi.Input[str],
                 host: Optional[pulumi.Input[str]] = None,
                 plaintext_password: Optional[pulumi.Input[str]] = None,
                 retain_old_password: Optional[pulumi.Input[bool]] = None):
        """
        The set of arguments for constructing a UserPassword resource.
        :param pulumi.Input[str] user: The IAM user to associate with this access key.
        :param pulumi.Input[str] host: The source host of the user. Defaults to `localhost`.
        """
        pulumi.set(__self__, "user", user)
        if host is not None:
            pulumi.set(__self__, "host", host)
        if plaintext_password is not None:
            pulumi.set(__self__, "plaintext_password", plaintext_password)
        if retain_old_password is not None:
            pulumi.set(__self__, "retain_old_password", retain_old_password)

    @property
    @pulumi.getter
    def user(self) -> pulumi.Input[str]:
        """
        The IAM user to associate with this access key.
        """
        return pulumi.get(self, "user")

    @user.setter
    def user(self, value: pulumi.Input[str]):
        pulumi.set(self, "user", value)

    @property
    @pulumi.getter
    def host(self) -> Optional[pulumi.Input[str]]:
        """
        The source host of the user. Defaults to `localhost`.
        """
        return pulumi.get(self, "host")

    @host.setter
    def host(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "host", value)

    @property
    @pulumi.getter(name="plaintextPassword")
    def plaintext_password(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "plaintext_password")

    @plaintext_password.setter
    def plaintext_password(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "plaintext_password", value)

    @property
    @pulumi.getter(name="retainOldPassword")
    def retain_old_password(self) -> Optional[pulumi.Input[bool]]:
        return pulumi.get(self, "retain_old_password")

    @retain_old_password.setter
    def retain_old_password(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "retain_old_password", value)


@pulumi.input_type
class _UserPasswordState:
    def __init__(__self__, *,
                 host: Optional[pulumi.Input[str]] = None,
                 plaintext_password: Optional[pulumi.Input[str]] = None,
                 retain_old_password: Optional[pulumi.Input[bool]] = None,
                 user: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering UserPassword resources.
        :param pulumi.Input[str] host: The source host of the user. Defaults to `localhost`.
        :param pulumi.Input[str] user: The IAM user to associate with this access key.
        """
        if host is not None:
            pulumi.set(__self__, "host", host)
        if plaintext_password is not None:
            pulumi.set(__self__, "plaintext_password", plaintext_password)
        if retain_old_password is not None:
            pulumi.set(__self__, "retain_old_password", retain_old_password)
        if user is not None:
            pulumi.set(__self__, "user", user)

    @property
    @pulumi.getter
    def host(self) -> Optional[pulumi.Input[str]]:
        """
        The source host of the user. Defaults to `localhost`.
        """
        return pulumi.get(self, "host")

    @host.setter
    def host(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "host", value)

    @property
    @pulumi.getter(name="plaintextPassword")
    def plaintext_password(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "plaintext_password")

    @plaintext_password.setter
    def plaintext_password(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "plaintext_password", value)

    @property
    @pulumi.getter(name="retainOldPassword")
    def retain_old_password(self) -> Optional[pulumi.Input[bool]]:
        return pulumi.get(self, "retain_old_password")

    @retain_old_password.setter
    def retain_old_password(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "retain_old_password", value)

    @property
    @pulumi.getter
    def user(self) -> Optional[pulumi.Input[str]]:
        """
        The IAM user to associate with this access key.
        """
        return pulumi.get(self, "user")

    @user.setter
    def user(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "user", value)


class UserPassword(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 host: Optional[pulumi.Input[str]] = None,
                 plaintext_password: Optional[pulumi.Input[str]] = None,
                 retain_old_password: Optional[pulumi.Input[bool]] = None,
                 user: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Create a UserPassword resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] host: The source host of the user. Defaults to `localhost`.
        :param pulumi.Input[str] user: The IAM user to associate with this access key.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: UserPasswordArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Create a UserPassword resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param UserPasswordArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(UserPasswordArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 host: Optional[pulumi.Input[str]] = None,
                 plaintext_password: Optional[pulumi.Input[str]] = None,
                 retain_old_password: Optional[pulumi.Input[bool]] = None,
                 user: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = UserPasswordArgs.__new__(UserPasswordArgs)

            __props__.__dict__["host"] = host
            __props__.__dict__["plaintext_password"] = None if plaintext_password is None else pulumi.Output.secret(plaintext_password)
            __props__.__dict__["retain_old_password"] = retain_old_password
            if user is None and not opts.urn:
                raise TypeError("Missing required property 'user'")
            __props__.__dict__["user"] = user
        secret_opts = pulumi.ResourceOptions(additional_secret_outputs=["plaintextPassword"])
        opts = pulumi.ResourceOptions.merge(opts, secret_opts)
        super(UserPassword, __self__).__init__(
            'mysql:index/userPassword:UserPassword',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            host: Optional[pulumi.Input[str]] = None,
            plaintext_password: Optional[pulumi.Input[str]] = None,
            retain_old_password: Optional[pulumi.Input[bool]] = None,
            user: Optional[pulumi.Input[str]] = None) -> 'UserPassword':
        """
        Get an existing UserPassword resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] host: The source host of the user. Defaults to `localhost`.
        :param pulumi.Input[str] user: The IAM user to associate with this access key.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _UserPasswordState.__new__(_UserPasswordState)

        __props__.__dict__["host"] = host
        __props__.__dict__["plaintext_password"] = plaintext_password
        __props__.__dict__["retain_old_password"] = retain_old_password
        __props__.__dict__["user"] = user
        return UserPassword(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def host(self) -> pulumi.Output[Optional[str]]:
        """
        The source host of the user. Defaults to `localhost`.
        """
        return pulumi.get(self, "host")

    @property
    @pulumi.getter(name="plaintextPassword")
    def plaintext_password(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "plaintext_password")

    @property
    @pulumi.getter(name="retainOldPassword")
    def retain_old_password(self) -> pulumi.Output[Optional[bool]]:
        return pulumi.get(self, "retain_old_password")

    @property
    @pulumi.getter
    def user(self) -> pulumi.Output[str]:
        """
        The IAM user to associate with this access key.
        """
        return pulumi.get(self, "user")

