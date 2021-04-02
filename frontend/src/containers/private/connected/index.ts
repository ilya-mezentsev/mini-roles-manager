import { connect } from 'react-redux';

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
    SignOut as SignOutContainer,
    mapDispatchToProps as SignOutMapDispatchToProps,
    mapStateToProps as SignOutContainerMapStateToProps,
} from '../sign_out/sign_out';

export const Resources = connect(
    ResourcesContainerMapStateToProps(),
    ResourcesContainerMapDispatchToProps(),
)(ResourcesContainer);

export const ResourcesList = connect(
    ResourcesListContainerMapStateToProps(),
    ResourcesListMapDispatchToProps(),
)(ResourcesListContainer);

export const SignOut = connect(
    SignOutContainerMapStateToProps(),
    SignOutMapDispatchToProps(),
)(SignOutContainer);
