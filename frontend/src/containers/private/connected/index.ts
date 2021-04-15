import { connect } from 'react-redux';

import {
    Entrypoint as EntrypointContainer,
    mapDispatchToProps as EntrypointContainerMapDispatchToProps,
    mapStateToProps as EntrypointContainerMapStateToProps,
} from '../entrypoint';

import {
    Resources as ResourcesContainer,
    mapDispatchToProps as ResourcesContainerMapDispatchToProps,
    mapStateToProps as ResourcesContainerMapStateToProps,
} from '../resources/resources';

import {
    ResourcesList as ResourcesListContainer,
    mapDispatchToProps as ResourcesListMapDispatchToProps,
    mapStateToProps as ResourcesListContainerMapStateToProps,
} from '../resources/list';

import {
    Roles as RolesContainer,
    mapDispatchToProps as RolesContainerMapDispatchToProps,
    mapStateToProps as RolesContainerMapStateToProps,
} from '../role/roles';

import {
    RolesList as RolesListContainer,
    mapDispatchToProps as RolesListContainerMapDispatchToProps,
    mapStateToProps as RolesListContainerMapStateToProps,
} from '../role/list';

import {
    Account as AccountContainer,
    mapDispatchToProps as AccountContainerMapDispatchToProps,
    mapStateToProps as AccountContainerMapStateToProps,
} from '../account/account';

import {
    Info as InfoContainer,
    mapStateToProps as InfoContainerMapStateToProps,
} from '../account/info';

import {
    ApiKey as ApiKeyContainer,
    mapStateToProps as ApiKeyContainerMapStateToProps,
} from '../account/api_key';

import {
    Credentials as CredentialsContainer,
    mapDispatchToProps as CredentialsContainerMapDispatchToProps,
    mapStateToProps as CredentialsContainerMapStateToProps,
} from '../account/credentials';

import {
    CheckPermissions as CheckPermissionsContainer,
    mapDispatchToProps as CheckPermissionsContainerMapDispatchToProps,
    mapStateToProps as CheckPermissionsContainerMapStateToProps,
} from '../account/check_permissions';

import {
    SignOut as SignOutContainer,
    mapDispatchToProps as SignOutMapDispatchToProps,
    mapStateToProps as SignOutContainerMapStateToProps,
} from '../sign_out/sign_out';

export const Entrypoint = connect(
    EntrypointContainerMapStateToProps(),
    EntrypointContainerMapDispatchToProps(),
)(EntrypointContainer);

export const Resources = connect(
    ResourcesContainerMapStateToProps(),
    ResourcesContainerMapDispatchToProps(),
)(ResourcesContainer);

export const ResourcesList = connect(
    ResourcesListContainerMapStateToProps(),
    ResourcesListMapDispatchToProps(),
)(ResourcesListContainer);

export const Roles = connect(
    RolesContainerMapStateToProps(),
    RolesContainerMapDispatchToProps(),
)(RolesContainer);

export const RolesList = connect(
    RolesListContainerMapStateToProps(),
    RolesListContainerMapDispatchToProps(),
)(RolesListContainer);

export const Account = connect(
    AccountContainerMapStateToProps(),
    AccountContainerMapDispatchToProps(),
)(AccountContainer);

export const AccountInfo = connect(
    InfoContainerMapStateToProps(),
)(InfoContainer);

export const AccountApiKey = connect(
    ApiKeyContainerMapStateToProps(),
)(ApiKeyContainer);

export const AccountCredentials = connect(
    CredentialsContainerMapStateToProps(),
    CredentialsContainerMapDispatchToProps(),
)(CredentialsContainer);

export const CheckPermissions = connect(
    CheckPermissionsContainerMapStateToProps(),
    CheckPermissionsContainerMapDispatchToProps(),
)(CheckPermissionsContainer);

export const SignOut = connect(
    SignOutContainerMapStateToProps(),
    SignOutMapDispatchToProps(),
)(SignOutContainer);
