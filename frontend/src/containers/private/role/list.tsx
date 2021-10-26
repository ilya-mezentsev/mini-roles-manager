import { useState } from 'react';
import { observer } from 'mobx-react-lite';
import EventEmitter from 'events';

import { Alert } from '../../../components/shared';
import { Prompter } from '../../../components/private/shared';
import { EditRole } from '../../../components/private/role';
import { RolesList as RolesListComponent } from '../../../components/private/role';
import { Role } from '../../../services/api';
import {
    rolesVersionStore,
    resourceStore,
    roleStore,
} from '../../../store';

export const RolesList = observer(() => {
    const [deletingRole, setDeletingRole] = useState<Role | null>(null);
    const [editingRole, setEditingRole] = useState<Role | null>(null);

    const e = new EventEmitter();
    const openPrompterEventName = 'prompter:open';
    const openEditResourceEventName = 'edit-role:open';

    const deleteRole = () => {
        if (deletingRole) {
            roleStore.deleteRole(
                rolesVersionStore.current!.id,
                deletingRole.id
            ).finally(() => setDeletingRole(null));
        }
    };

    const hasAnyError: () => boolean = () => {
        return (
            !!roleStore.fetchRoleError ||
            !!roleStore.updateRoleError ||
            !!roleStore.deleteRoleError
        );
    };

    const cleanErrors = () => roleStore.cleanRoleActionErrors();

    const errorMessage: () => string = () => {
        return (
            roleStore.fetchRoleError?.description ||
            roleStore.updateRoleError?.description ||
            roleStore.deleteRoleError?.description ||
            'Unknown error'
        );
    };

    return (
        <>
            <RolesListComponent
                roles={
                    roleStore.list
                        .filter(r => r.versionId === rolesVersionStore.current?.id)
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
                existRoles={roleStore.list}
                existsResources={resourceStore.list}
                roleVersionId={rolesVersionStore.current?.id || ''}
                save={r => roleStore.updateRole(r)}
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
                onCloseCb={() => cleanErrors()}
            />
        </>
    );
});
