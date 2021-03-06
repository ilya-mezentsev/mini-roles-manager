import { useEffect } from 'react';
import { bindActionCreators } from 'redux';

import { EntrypointActions, EntrypointState, EntrypointProps } from './entrypoint.types';
import { Navigation as PrivateNavigation } from './navigation/navigation';
import { DispatchToPropsFn, StateToPropsFn } from '../../shared/types';
import { fetchRoles } from '../../store/role/actions';
import { fetchResources } from '../../store/resource/actions';
import { fetchInfo } from '../../store/account_info/actions';
import { fetchRolesVersion } from '../../store/roles_version/actions';

export const Entrypoint = (props: EntrypointProps) => {
    useEffect(() => {
        if (!props.resourcesResult.list) {
            props.loadResourcesAction();
        }

        if (!props.rolesVersionResult.list) {
            props.loadRolesVersionsAction();
        }

        if (!props.rolesResult.list) {
            props.loadRolesAction();
        }

        if (!props.accountInfoResult?.info?.login) {
            props.loadAccountInfo();
        }
        // eslint-disable-next-line
    }, []);

    return (
        <PrivateNavigation/>
    );
}

export const mapDispatchToProps: DispatchToPropsFn<EntrypointActions> = () => dispatch => ({
    loadResourcesAction: bindActionCreators(fetchResources, dispatch),
    loadRolesVersionsAction: bindActionCreators(fetchRolesVersion, dispatch),
    loadRolesAction: bindActionCreators(fetchRoles, dispatch),
    loadAccountInfo: bindActionCreators(fetchInfo, dispatch),
});

export const mapStateToProps: StateToPropsFn<EntrypointState> = () => state => ({
    rolesVersionResult: state.rolesVersionResult,
    rolesResult: state.rolesResult,
    resourcesResult: state.resourcesResult,
    accountInfoResult: state.accountInfoResult,
});
