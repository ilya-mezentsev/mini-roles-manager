import { bindActionCreators } from 'redux';
import { Box } from '@material-ui/core';
import { Add } from '@material-ui/icons';
import EventEmitter from 'events';

import { Alert } from '../../../components/shared';
import { EditRole as CreateRole } from '../../../components/private/role';
import { DispatchToPropsFn, StateToPropsFn } from '../../../shared/types';
import {
    cleanCreateRoleError,
    createRole,
    fetchRoles,
} from '../../../store/role/actions';
import { RoleActions, RoleState, RoleProps } from './roles.types';
import { RolesList } from '../connected';

export const Roles = (props: RoleProps) => {
    const e = new EventEmitter();
    const openDialogueEventName = 'new-role-dialogue:open';

    return (
        <>
            <Box>
                <h1>
                    Roles
                    <Add
                        color="primary"
                        cursor="pointer"
                        fontSize="large"
                        titleAccess="Add new role"
                        onClick={() => e.emit(openDialogueEventName)}
                    />
                </h1>
            </Box>

            <Box>
                <RolesList/>
            </Box>

            <CreateRole
                eventEmitter={e}
                openDialogueEventName={openDialogueEventName}
                existRoles={props.rolesResult.list || []}
                existsResources={props.resourcesResult.list || []}
                save={r => props.createRoleAction(r)}
            />

            <Alert
                shouldShow={!!props.rolesResult?.createError}
                severity={'error'}
                message={props.rolesResult?.createError?.description || 'Unknown error'}
                onCloseCb={() => props.cleanCreateRoleErrorAction()}
            />
        </>
    )
};

export const mapDispatchToProps: DispatchToPropsFn<RoleActions> = () => dispatch => ({
    createRoleAction: bindActionCreators(createRole, dispatch),
    cleanCreateRoleErrorAction: bindActionCreators(cleanCreateRoleError, dispatch),

    loadRolesAction: bindActionCreators(fetchRoles, dispatch),
});

export const mapStateToProps: StateToPropsFn<RoleState> = () => state => ({
    rolesResult: state.rolesResult,
    resourcesResult: state.resourcesResult,
});
