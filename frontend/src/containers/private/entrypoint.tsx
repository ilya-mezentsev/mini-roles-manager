import { useEffect } from 'react';
import { bindActionCreators } from 'redux';

import { EntrypointActions, EntrypointState, EntrypointProps } from './entrypoint.types';
import { Navigation as PrivateNavigation } from './navigation/navigation';
import { DispatchToPropsFn, StateToPropsFn } from '../../shared/types';
import { fetchRoles } from '../../store/role/actions';
import { fetchResources } from '../../store/resource/actions';

export const Entrypoint = (props: EntrypointProps) => {
    useEffect(() => {
        if (!props.resourcesResult.list) {
            props.loadResourcesAction();
        }

        if (!props.rolesResult.list) {
            props.loadRolesAction();
        }
    }, []);

    return (
        <PrivateNavigation/>
    );
}

export const mapDispatchToProps: DispatchToPropsFn<EntrypointActions> = () => dispatch => ({
    loadResourcesAction: bindActionCreators(fetchResources, dispatch),
    loadRolesAction: bindActionCreators(fetchRoles, dispatch),
});

export const mapStateToProps: StateToPropsFn<EntrypointState> = () => state => ({
    rolesResult: state.rolesResult,
    resourcesResult: state.resourcesResult,
});
