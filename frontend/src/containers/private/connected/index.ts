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

export const Resources = connect(
    ResourcesContainerMapStateToProps(),
    ResourcesContainerMapDispatchToProps(),
)(ResourcesContainer);

export const ResourcesList = connect(
    ResourcesListContainerMapStateToProps(),
    ResourcesListMapDispatchToProps(),
)(ResourcesListContainer);
