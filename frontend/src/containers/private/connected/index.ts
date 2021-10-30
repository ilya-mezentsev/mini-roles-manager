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
    RolesVersion as RolesVersionContainer,
    mapDispatchToProps as RolesVersionContainerMapDispatchToProps,
    mapStateToProps as RolesVersionContainerMapStateToProps,
} from '../roles_version/roles_version';

import {
    RolesVersionList as RolesVersionListContainer,
    mapDispatchToProps as RolesVersionListContainerMapDispatchToProps,
    mapStateToProps as RolesVersionListContainerMapStateToProps,
} from '../roles_version/list';

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
    Info as InfoContainer,
    mapDispatchToProps as InfoContainerMapDispatchToProps,
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
    Import as ImportContainer,
    mapDispatchToProps as ImportContainerMapDispatchToProps,
    mapStateToProps as ImportContainerMapStateToProps,
} from '../app_data/import';

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

export const RolesVersion = connect(
    RolesVersionContainerMapStateToProps(),
    RolesVersionContainerMapDispatchToProps(),
)(RolesVersionContainer);

export const RolesVersionList = connect(
    RolesVersionListContainerMapStateToProps(),
    RolesVersionListContainerMapDispatchToProps(),
)(RolesVersionListContainer);

export const Roles = connect(
    RolesContainerMapStateToProps(),
    RolesContainerMapDispatchToProps(),
)(RolesContainer);

export const RolesList = connect(
    RolesListContainerMapStateToProps(),
    RolesListContainerMapDispatchToProps(),
)(RolesListContainer);

export const AccountInfo = connect(
    InfoContainerMapStateToProps(),
    InfoContainerMapDispatchToProps(),
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

export const Import = connect(
    ImportContainerMapStateToProps(),
    ImportContainerMapDispatchToProps(),
)(ImportContainer);

export const SignOut = connect(
    SignOutContainerMapStateToProps(),
    SignOutMapDispatchToProps(),
)(SignOutContainer);
