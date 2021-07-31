import { useState } from 'react';
import { bindActionCreators } from 'redux';
import EventEmitter from 'events';

import { Alert } from '../../../components/shared';
import { Prompter } from '../../../components/private/shared/prompter';
import { EditRole } from '../../../components/private/role';
import { RolesListActions, RolesListState, RolesListProps } from './list.types';
import { DispatchToPropsFn, StateToPropsFn } from '../../../shared/types';
import { RolesList as RolesListComponent } from '../../../components/private/role';
import {
    cleanDeleteRoleError,
    cleanFetchRolesError,
    cleanUpdateRoleError,
    deleteRole,
    updateRole,
} from '../../../store/role/actions';
import { Role } from '../../../services/api';

export const RolesList = (props: RolesListProps) => {
    const [deletingRole, setDeletingRole] = useState<Role | null>(null);
    const [editingRole, setEditingRole] = useState<Role | null>(null);

    const e = new EventEmitter();
    const openPrompterEventName = 'prompter:open';
    const openEditResourceEventName = 'edit-role:open';

    const deleteRole = () => {
        if (deletingRole) {
            props.deleteRoleAction(
                props.rolesVersionResult.currentRolesVersion!.id,
                deletingRole.id,
            );
            setDeletingRole(null);
        }
    };

    const hasAnyError: () => boolean = () => {
        return (
            !!props.rolesResult?.fetchError ||
            !!props.rolesResult?.updateError ||
            !!props.rolesResult?.deleteError
        );
    };

    const cleanError = () => {
        if (props.rolesResult?.fetchError) {
            props.cleanFetchRolesErrorAction();
        } else if (props.rolesResult?.updateError) {
            props.cleanUpdateRoleErrorAction();
        } else if (props.rolesResult?.deleteError) {
            props.cleanDeleteRoleErrorAction();
        }
    };

    const errorMessage: () => string = () => {
        return (
            props.rolesResult?.fetchError?.description ||
            props.rolesResult?.updateError?.description ||
            props.rolesResult?.deleteError?.description ||
            'Unknown error'
        );
    };

    return (
        <>
            <RolesListComponent
                roles={
                    (props.rolesResult.list || [])
                        .filter(r => r.versionId === props.rolesVersionResult.currentRolesVersion?.id)
                }
                tryEdit={r => {
                    setEditingRole(r);
                    e.emit(openEditResourceEventName);
                }}
                tryDelete={r => {
                    setDeletingRole(r);
                    e.emit(openPrompterEventName);
                }}
            />

            <EditRole
                eventEmitter={e}
                openDialogueEventName={openEditResourceEventName}
                existRoles={props.rolesResult.list || []}
                existsResources={props.resourcesResult.list || []}
                roleVersionId={props.rolesVersionResult.currentRolesVersion?.id || ''}
                save={r => props.updateRoleAction(r)}
                initialRole={editingRole}
            />

            <Prompter
                title="Role deletion"
                description="Are you sure you want to delete this role?"
                onAgree={() => deleteRole()}
                onDisagree={() => setDeletingRole(null)}
                openDialogueEventName={openPrompterEventName}
                eventEmitter={e}
            />

            <Alert
                shouldShow={hasAnyError()}
                severity={'error'}
                message={errorMessage()}
                onCloseCb={() => cleanError()}
            />
        </>
    );
};

export const mapDispatchToProps: DispatchToPropsFn<RolesListActions> = () => dispatch => ({
    cleanFetchRolesErrorAction: bindActionCreators(cleanFetchRolesError, dispatch),

    updateRoleAction: bindActionCreators(updateRole, dispatch),
    cleanUpdateRoleErrorAction: bindActionCreators(cleanUpdateRoleError, dispatch),

    deleteRoleAction: bindActionCreators(deleteRole, dispatch),
    cleanDeleteRoleErrorAction: bindActionCreators(cleanDeleteRoleError, dispatch),
});

export const mapStateToProps: StateToPropsFn<RolesListState> = () => state => ({
    rolesVersionResult: state.rolesVersionResult,
    rolesResult: state.rolesResult,
    resourcesResult: state.resourcesResult,
});
