# coding=utf-8
# *** WARNING: this file was generated by pulumigen. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from ... import _utilities
from . import outputs

__all__ = [
    'BoundObjectReference',
    'BoundObjectReferencePatch',
    'TokenRequestSpec',
    'TokenRequestSpecPatch',
    'TokenRequestStatus',
    'TokenRequestStatusPatch',
    'TokenReviewSpec',
    'TokenReviewSpecPatch',
    'TokenReviewStatus',
    'TokenReviewStatusPatch',
    'UserInfo',
    'UserInfoPatch',
]

@pulumi.output_type
class BoundObjectReference(dict):
    """
    BoundObjectReference is a reference to an object that a token is bound to.
    """
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "apiVersion":
            suggest = "api_version"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in BoundObjectReference. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        BoundObjectReference.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        BoundObjectReference.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 api_version: Optional[str] = None,
                 kind: Optional[str] = None,
                 name: Optional[str] = None,
                 uid: Optional[str] = None):
        """
        BoundObjectReference is a reference to an object that a token is bound to.
        :param str api_version: API version of the referent.
        :param str kind: Kind of the referent. Valid kinds are 'Pod' and 'Secret'.
        :param str name: Name of the referent.
        :param str uid: UID of the referent.
        """
        if api_version is not None:
            pulumi.set(__self__, "api_version", api_version)
        if kind is not None:
            pulumi.set(__self__, "kind", kind)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if uid is not None:
            pulumi.set(__self__, "uid", uid)

    @property
    @pulumi.getter(name="apiVersion")
    def api_version(self) -> Optional[str]:
        """
        API version of the referent.
        """
        return pulumi.get(self, "api_version")

    @property
    @pulumi.getter
    def kind(self) -> Optional[str]:
        """
        Kind of the referent. Valid kinds are 'Pod' and 'Secret'.
        """
        return pulumi.get(self, "kind")

    @property
    @pulumi.getter
    def name(self) -> Optional[str]:
        """
        Name of the referent.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def uid(self) -> Optional[str]:
        """
        UID of the referent.
        """
        return pulumi.get(self, "uid")


@pulumi.output_type
class BoundObjectReferencePatch(dict):
    """
    BoundObjectReference is a reference to an object that a token is bound to.
    """
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "apiVersion":
            suggest = "api_version"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in BoundObjectReferencePatch. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        BoundObjectReferencePatch.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        BoundObjectReferencePatch.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 api_version: Optional[str] = None,
                 kind: Optional[str] = None,
                 name: Optional[str] = None,
                 uid: Optional[str] = None):
        """
        BoundObjectReference is a reference to an object that a token is bound to.
        :param str api_version: API version of the referent.
        :param str kind: Kind of the referent. Valid kinds are 'Pod' and 'Secret'.
        :param str name: Name of the referent.
        :param str uid: UID of the referent.
        """
        if api_version is not None:
            pulumi.set(__self__, "api_version", api_version)
        if kind is not None:
            pulumi.set(__self__, "kind", kind)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if uid is not None:
            pulumi.set(__self__, "uid", uid)

    @property
    @pulumi.getter(name="apiVersion")
    def api_version(self) -> Optional[str]:
        """
        API version of the referent.
        """
        return pulumi.get(self, "api_version")

    @property
    @pulumi.getter
    def kind(self) -> Optional[str]:
        """
        Kind of the referent. Valid kinds are 'Pod' and 'Secret'.
        """
        return pulumi.get(self, "kind")

    @property
    @pulumi.getter
    def name(self) -> Optional[str]:
        """
        Name of the referent.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def uid(self) -> Optional[str]:
        """
        UID of the referent.
        """
        return pulumi.get(self, "uid")


@pulumi.output_type
class TokenRequestSpec(dict):
    """
    TokenRequestSpec contains client provided parameters of a token request.
    """
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "boundObjectRef":
            suggest = "bound_object_ref"
        elif key == "expirationSeconds":
            suggest = "expiration_seconds"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in TokenRequestSpec. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        TokenRequestSpec.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        TokenRequestSpec.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 audiences: Sequence[str],
                 bound_object_ref: Optional['outputs.BoundObjectReference'] = None,
                 expiration_seconds: Optional[int] = None):
        """
        TokenRequestSpec contains client provided parameters of a token request.
        :param Sequence[str] audiences: Audiences are the intendend audiences of the token. A recipient of a token must identitfy themself with an identifier in the list of audiences of the token, and otherwise should reject the token. A token issued for multiple audiences may be used to authenticate against any of the audiences listed but implies a high degree of trust between the target audiences.
        :param 'BoundObjectReferenceArgs' bound_object_ref: BoundObjectRef is a reference to an object that the token will be bound to. The token will only be valid for as long as the bound object exists. NOTE: The API server's TokenReview endpoint will validate the BoundObjectRef, but other audiences may not. Keep ExpirationSeconds small if you want prompt revocation.
        :param int expiration_seconds: ExpirationSeconds is the requested duration of validity of the request. The token issuer may return a token with a different validity duration so a client needs to check the 'expiration' field in a response.
        """
        pulumi.set(__self__, "audiences", audiences)
        if bound_object_ref is not None:
            pulumi.set(__self__, "bound_object_ref", bound_object_ref)
        if expiration_seconds is not None:
            pulumi.set(__self__, "expiration_seconds", expiration_seconds)

    @property
    @pulumi.getter
    def audiences(self) -> Sequence[str]:
        """
        Audiences are the intendend audiences of the token. A recipient of a token must identitfy themself with an identifier in the list of audiences of the token, and otherwise should reject the token. A token issued for multiple audiences may be used to authenticate against any of the audiences listed but implies a high degree of trust between the target audiences.
        """
        return pulumi.get(self, "audiences")

    @property
    @pulumi.getter(name="boundObjectRef")
    def bound_object_ref(self) -> Optional['outputs.BoundObjectReference']:
        """
        BoundObjectRef is a reference to an object that the token will be bound to. The token will only be valid for as long as the bound object exists. NOTE: The API server's TokenReview endpoint will validate the BoundObjectRef, but other audiences may not. Keep ExpirationSeconds small if you want prompt revocation.
        """
        return pulumi.get(self, "bound_object_ref")

    @property
    @pulumi.getter(name="expirationSeconds")
    def expiration_seconds(self) -> Optional[int]:
        """
        ExpirationSeconds is the requested duration of validity of the request. The token issuer may return a token with a different validity duration so a client needs to check the 'expiration' field in a response.
        """
        return pulumi.get(self, "expiration_seconds")


@pulumi.output_type
class TokenRequestSpecPatch(dict):
    """
    TokenRequestSpec contains client provided parameters of a token request.
    """
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "boundObjectRef":
            suggest = "bound_object_ref"
        elif key == "expirationSeconds":
            suggest = "expiration_seconds"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in TokenRequestSpecPatch. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        TokenRequestSpecPatch.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        TokenRequestSpecPatch.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 audiences: Optional[Sequence[str]] = None,
                 bound_object_ref: Optional['outputs.BoundObjectReferencePatch'] = None,
                 expiration_seconds: Optional[int] = None):
        """
        TokenRequestSpec contains client provided parameters of a token request.
        :param Sequence[str] audiences: Audiences are the intendend audiences of the token. A recipient of a token must identitfy themself with an identifier in the list of audiences of the token, and otherwise should reject the token. A token issued for multiple audiences may be used to authenticate against any of the audiences listed but implies a high degree of trust between the target audiences.
        :param 'BoundObjectReferencePatchArgs' bound_object_ref: BoundObjectRef is a reference to an object that the token will be bound to. The token will only be valid for as long as the bound object exists. NOTE: The API server's TokenReview endpoint will validate the BoundObjectRef, but other audiences may not. Keep ExpirationSeconds small if you want prompt revocation.
        :param int expiration_seconds: ExpirationSeconds is the requested duration of validity of the request. The token issuer may return a token with a different validity duration so a client needs to check the 'expiration' field in a response.
        """
        if audiences is not None:
            pulumi.set(__self__, "audiences", audiences)
        if bound_object_ref is not None:
            pulumi.set(__self__, "bound_object_ref", bound_object_ref)
        if expiration_seconds is not None:
            pulumi.set(__self__, "expiration_seconds", expiration_seconds)

    @property
    @pulumi.getter
    def audiences(self) -> Optional[Sequence[str]]:
        """
        Audiences are the intendend audiences of the token. A recipient of a token must identitfy themself with an identifier in the list of audiences of the token, and otherwise should reject the token. A token issued for multiple audiences may be used to authenticate against any of the audiences listed but implies a high degree of trust between the target audiences.
        """
        return pulumi.get(self, "audiences")

    @property
    @pulumi.getter(name="boundObjectRef")
    def bound_object_ref(self) -> Optional['outputs.BoundObjectReferencePatch']:
        """
        BoundObjectRef is a reference to an object that the token will be bound to. The token will only be valid for as long as the bound object exists. NOTE: The API server's TokenReview endpoint will validate the BoundObjectRef, but other audiences may not. Keep ExpirationSeconds small if you want prompt revocation.
        """
        return pulumi.get(self, "bound_object_ref")

    @property
    @pulumi.getter(name="expirationSeconds")
    def expiration_seconds(self) -> Optional[int]:
        """
        ExpirationSeconds is the requested duration of validity of the request. The token issuer may return a token with a different validity duration so a client needs to check the 'expiration' field in a response.
        """
        return pulumi.get(self, "expiration_seconds")


@pulumi.output_type
class TokenRequestStatus(dict):
    """
    TokenRequestStatus is the result of a token request.
    """
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "expirationTimestamp":
            suggest = "expiration_timestamp"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in TokenRequestStatus. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        TokenRequestStatus.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        TokenRequestStatus.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 expiration_timestamp: str,
                 token: str):
        """
        TokenRequestStatus is the result of a token request.
        :param str expiration_timestamp: ExpirationTimestamp is the time of expiration of the returned token.
        :param str token: Token is the opaque bearer token.
        """
        pulumi.set(__self__, "expiration_timestamp", expiration_timestamp)
        pulumi.set(__self__, "token", token)

    @property
    @pulumi.getter(name="expirationTimestamp")
    def expiration_timestamp(self) -> str:
        """
        ExpirationTimestamp is the time of expiration of the returned token.
        """
        return pulumi.get(self, "expiration_timestamp")

    @property
    @pulumi.getter
    def token(self) -> str:
        """
        Token is the opaque bearer token.
        """
        return pulumi.get(self, "token")


@pulumi.output_type
class TokenRequestStatusPatch(dict):
    """
    TokenRequestStatus is the result of a token request.
    """
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "expirationTimestamp":
            suggest = "expiration_timestamp"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in TokenRequestStatusPatch. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        TokenRequestStatusPatch.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        TokenRequestStatusPatch.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 expiration_timestamp: Optional[str] = None,
                 token: Optional[str] = None):
        """
        TokenRequestStatus is the result of a token request.
        :param str expiration_timestamp: ExpirationTimestamp is the time of expiration of the returned token.
        :param str token: Token is the opaque bearer token.
        """
        if expiration_timestamp is not None:
            pulumi.set(__self__, "expiration_timestamp", expiration_timestamp)
        if token is not None:
            pulumi.set(__self__, "token", token)

    @property
    @pulumi.getter(name="expirationTimestamp")
    def expiration_timestamp(self) -> Optional[str]:
        """
        ExpirationTimestamp is the time of expiration of the returned token.
        """
        return pulumi.get(self, "expiration_timestamp")

    @property
    @pulumi.getter
    def token(self) -> Optional[str]:
        """
        Token is the opaque bearer token.
        """
        return pulumi.get(self, "token")


@pulumi.output_type
class TokenReviewSpec(dict):
    """
    TokenReviewSpec is a description of the token authentication request.
    """
    def __init__(__self__, *,
                 audiences: Optional[Sequence[str]] = None,
                 token: Optional[str] = None):
        """
        TokenReviewSpec is a description of the token authentication request.
        :param Sequence[str] audiences: Audiences is a list of the identifiers that the resource server presented with the token identifies as. Audience-aware token authenticators will verify that the token was intended for at least one of the audiences in this list. If no audiences are provided, the audience will default to the audience of the Kubernetes apiserver.
        :param str token: Token is the opaque bearer token.
        """
        if audiences is not None:
            pulumi.set(__self__, "audiences", audiences)
        if token is not None:
            pulumi.set(__self__, "token", token)

    @property
    @pulumi.getter
    def audiences(self) -> Optional[Sequence[str]]:
        """
        Audiences is a list of the identifiers that the resource server presented with the token identifies as. Audience-aware token authenticators will verify that the token was intended for at least one of the audiences in this list. If no audiences are provided, the audience will default to the audience of the Kubernetes apiserver.
        """
        return pulumi.get(self, "audiences")

    @property
    @pulumi.getter
    def token(self) -> Optional[str]:
        """
        Token is the opaque bearer token.
        """
        return pulumi.get(self, "token")


@pulumi.output_type
class TokenReviewSpecPatch(dict):
    """
    TokenReviewSpec is a description of the token authentication request.
    """
    def __init__(__self__, *,
                 audiences: Optional[Sequence[str]] = None,
                 token: Optional[str] = None):
        """
        TokenReviewSpec is a description of the token authentication request.
        :param Sequence[str] audiences: Audiences is a list of the identifiers that the resource server presented with the token identifies as. Audience-aware token authenticators will verify that the token was intended for at least one of the audiences in this list. If no audiences are provided, the audience will default to the audience of the Kubernetes apiserver.
        :param str token: Token is the opaque bearer token.
        """
        if audiences is not None:
            pulumi.set(__self__, "audiences", audiences)
        if token is not None:
            pulumi.set(__self__, "token", token)

    @property
    @pulumi.getter
    def audiences(self) -> Optional[Sequence[str]]:
        """
        Audiences is a list of the identifiers that the resource server presented with the token identifies as. Audience-aware token authenticators will verify that the token was intended for at least one of the audiences in this list. If no audiences are provided, the audience will default to the audience of the Kubernetes apiserver.
        """
        return pulumi.get(self, "audiences")

    @property
    @pulumi.getter
    def token(self) -> Optional[str]:
        """
        Token is the opaque bearer token.
        """
        return pulumi.get(self, "token")


@pulumi.output_type
class TokenReviewStatus(dict):
    """
    TokenReviewStatus is the result of the token authentication request.
    """
    def __init__(__self__, *,
                 audiences: Optional[Sequence[str]] = None,
                 authenticated: Optional[bool] = None,
                 error: Optional[str] = None,
                 user: Optional['outputs.UserInfo'] = None):
        """
        TokenReviewStatus is the result of the token authentication request.
        :param Sequence[str] audiences: Audiences are audience identifiers chosen by the authenticator that are compatible with both the TokenReview and token. An identifier is any identifier in the intersection of the TokenReviewSpec audiences and the token's audiences. A client of the TokenReview API that sets the spec.audiences field should validate that a compatible audience identifier is returned in the status.audiences field to ensure that the TokenReview server is audience aware. If a TokenReview returns an empty status.audience field where status.authenticated is "true", the token is valid against the audience of the Kubernetes API server.
        :param bool authenticated: Authenticated indicates that the token was associated with a known user.
        :param str error: Error indicates that the token couldn't be checked
        :param 'UserInfoArgs' user: User is the UserInfo associated with the provided token.
        """
        if audiences is not None:
            pulumi.set(__self__, "audiences", audiences)
        if authenticated is not None:
            pulumi.set(__self__, "authenticated", authenticated)
        if error is not None:
            pulumi.set(__self__, "error", error)
        if user is not None:
            pulumi.set(__self__, "user", user)

    @property
    @pulumi.getter
    def audiences(self) -> Optional[Sequence[str]]:
        """
        Audiences are audience identifiers chosen by the authenticator that are compatible with both the TokenReview and token. An identifier is any identifier in the intersection of the TokenReviewSpec audiences and the token's audiences. A client of the TokenReview API that sets the spec.audiences field should validate that a compatible audience identifier is returned in the status.audiences field to ensure that the TokenReview server is audience aware. If a TokenReview returns an empty status.audience field where status.authenticated is "true", the token is valid against the audience of the Kubernetes API server.
        """
        return pulumi.get(self, "audiences")

    @property
    @pulumi.getter
    def authenticated(self) -> Optional[bool]:
        """
        Authenticated indicates that the token was associated with a known user.
        """
        return pulumi.get(self, "authenticated")

    @property
    @pulumi.getter
    def error(self) -> Optional[str]:
        """
        Error indicates that the token couldn't be checked
        """
        return pulumi.get(self, "error")

    @property
    @pulumi.getter
    def user(self) -> Optional['outputs.UserInfo']:
        """
        User is the UserInfo associated with the provided token.
        """
        return pulumi.get(self, "user")


@pulumi.output_type
class TokenReviewStatusPatch(dict):
    """
    TokenReviewStatus is the result of the token authentication request.
    """
    def __init__(__self__, *,
                 audiences: Optional[Sequence[str]] = None,
                 authenticated: Optional[bool] = None,
                 error: Optional[str] = None,
                 user: Optional['outputs.UserInfoPatch'] = None):
        """
        TokenReviewStatus is the result of the token authentication request.
        :param Sequence[str] audiences: Audiences are audience identifiers chosen by the authenticator that are compatible with both the TokenReview and token. An identifier is any identifier in the intersection of the TokenReviewSpec audiences and the token's audiences. A client of the TokenReview API that sets the spec.audiences field should validate that a compatible audience identifier is returned in the status.audiences field to ensure that the TokenReview server is audience aware. If a TokenReview returns an empty status.audience field where status.authenticated is "true", the token is valid against the audience of the Kubernetes API server.
        :param bool authenticated: Authenticated indicates that the token was associated with a known user.
        :param str error: Error indicates that the token couldn't be checked
        :param 'UserInfoPatchArgs' user: User is the UserInfo associated with the provided token.
        """
        if audiences is not None:
            pulumi.set(__self__, "audiences", audiences)
        if authenticated is not None:
            pulumi.set(__self__, "authenticated", authenticated)
        if error is not None:
            pulumi.set(__self__, "error", error)
        if user is not None:
            pulumi.set(__self__, "user", user)

    @property
    @pulumi.getter
    def audiences(self) -> Optional[Sequence[str]]:
        """
        Audiences are audience identifiers chosen by the authenticator that are compatible with both the TokenReview and token. An identifier is any identifier in the intersection of the TokenReviewSpec audiences and the token's audiences. A client of the TokenReview API that sets the spec.audiences field should validate that a compatible audience identifier is returned in the status.audiences field to ensure that the TokenReview server is audience aware. If a TokenReview returns an empty status.audience field where status.authenticated is "true", the token is valid against the audience of the Kubernetes API server.
        """
        return pulumi.get(self, "audiences")

    @property
    @pulumi.getter
    def authenticated(self) -> Optional[bool]:
        """
        Authenticated indicates that the token was associated with a known user.
        """
        return pulumi.get(self, "authenticated")

    @property
    @pulumi.getter
    def error(self) -> Optional[str]:
        """
        Error indicates that the token couldn't be checked
        """
        return pulumi.get(self, "error")

    @property
    @pulumi.getter
    def user(self) -> Optional['outputs.UserInfoPatch']:
        """
        User is the UserInfo associated with the provided token.
        """
        return pulumi.get(self, "user")


@pulumi.output_type
class UserInfo(dict):
    """
    UserInfo holds the information about the user needed to implement the user.Info interface.
    """
    def __init__(__self__, *,
                 extra: Optional[Mapping[str, Sequence[str]]] = None,
                 groups: Optional[Sequence[str]] = None,
                 uid: Optional[str] = None,
                 username: Optional[str] = None):
        """
        UserInfo holds the information about the user needed to implement the user.Info interface.
        :param Mapping[str, Sequence[str]] extra: Any additional information provided by the authenticator.
        :param Sequence[str] groups: The names of groups this user is a part of.
        :param str uid: A unique value that identifies this user across time. If this user is deleted and another user by the same name is added, they will have different UIDs.
        :param str username: The name that uniquely identifies this user among all active users.
        """
        if extra is not None:
            pulumi.set(__self__, "extra", extra)
        if groups is not None:
            pulumi.set(__self__, "groups", groups)
        if uid is not None:
            pulumi.set(__self__, "uid", uid)
        if username is not None:
            pulumi.set(__self__, "username", username)

    @property
    @pulumi.getter
    def extra(self) -> Optional[Mapping[str, Sequence[str]]]:
        """
        Any additional information provided by the authenticator.
        """
        return pulumi.get(self, "extra")

    @property
    @pulumi.getter
    def groups(self) -> Optional[Sequence[str]]:
        """
        The names of groups this user is a part of.
        """
        return pulumi.get(self, "groups")

    @property
    @pulumi.getter
    def uid(self) -> Optional[str]:
        """
        A unique value that identifies this user across time. If this user is deleted and another user by the same name is added, they will have different UIDs.
        """
        return pulumi.get(self, "uid")

    @property
    @pulumi.getter
    def username(self) -> Optional[str]:
        """
        The name that uniquely identifies this user among all active users.
        """
        return pulumi.get(self, "username")


@pulumi.output_type
class UserInfoPatch(dict):
    """
    UserInfo holds the information about the user needed to implement the user.Info interface.
    """
    def __init__(__self__, *,
                 extra: Optional[Mapping[str, Sequence[str]]] = None,
                 groups: Optional[Sequence[str]] = None,
                 uid: Optional[str] = None,
                 username: Optional[str] = None):
        """
        UserInfo holds the information about the user needed to implement the user.Info interface.
        :param Mapping[str, Sequence[str]] extra: Any additional information provided by the authenticator.
        :param Sequence[str] groups: The names of groups this user is a part of.
        :param str uid: A unique value that identifies this user across time. If this user is deleted and another user by the same name is added, they will have different UIDs.
        :param str username: The name that uniquely identifies this user among all active users.
        """
        if extra is not None:
            pulumi.set(__self__, "extra", extra)
        if groups is not None:
            pulumi.set(__self__, "groups", groups)
        if uid is not None:
            pulumi.set(__self__, "uid", uid)
        if username is not None:
            pulumi.set(__self__, "username", username)

    @property
    @pulumi.getter
    def extra(self) -> Optional[Mapping[str, Sequence[str]]]:
        """
        Any additional information provided by the authenticator.
        """
        return pulumi.get(self, "extra")

    @property
    @pulumi.getter
    def groups(self) -> Optional[Sequence[str]]:
        """
        The names of groups this user is a part of.
        """
        return pulumi.get(self, "groups")

    @property
    @pulumi.getter
    def uid(self) -> Optional[str]:
        """
        A unique value that identifies this user across time. If this user is deleted and another user by the same name is added, they will have different UIDs.
        """
        return pulumi.get(self, "uid")

    @property
    @pulumi.getter
    def username(self) -> Optional[str]:
        """
        The name that uniquely identifies this user among all active users.
        """
        return pulumi.get(self, "username")


