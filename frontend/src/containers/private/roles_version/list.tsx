import { useState } from 'react';
import { observer } from 'mobx-react-lite';
import EventEmitter from 'events';
import _ from 'lodash';

import { RolesVersion } from '../../../services/api';
import {
    EditRolesVersion,
    RolesVersionList as RolesVersionListComponent
} from '../../../components/private/roles_version';
import { Prompter } from '../../../components/private/shared';
import { Alert } from '../../../components/shared';
import { rolesVersionStore, roleStore } from '../../../store';

export const RolesVersionList = observer(() => {
    const [deletingRolesVersion, setDeletingRolesVersion] = useState<RolesVersion | null>(null);
    const [editingRolesVersion, setEditingRolesVersion] = useState<RolesVersion | null>(null);

    const e = new EventEmitter();
    const openPrompterEventName = 'prompter:open';
    const openEditResourceEventName = 'edit-roles-version:open';

    const deleteRolesVersion = () => {
        if (deletingRolesVersion) {
            rolesVersionStore
                .deleteRolesVersion(deletingRolesVersion.id)
                .finally(() => setDeletingRolesVersion(null));
        }
    };

    const hasAnyError = () => {
        return (
            !!rolesVersionStore.fetchRolesVersionError ||
            !!rolesVersionStore.updateRolesVersionError ||
            !!rolesVersionStore.deleteRolesVersionError ||
            !!roleStore.fetchRoleError
        );
    };

    const cleanErrors = () => {
        rolesVersionStore.cleanRolesVersionActionErrors();
        roleStore.cleanRoleActionErrors();
    };

    const errorMessage: () => string = () => {
        return (
            rolesVersionStore.fetchRolesVersionError?.description ||
            rolesVersionStore.updateRolesVersionError?.description ||
            rolesVersionStore.deleteRolesVersionError?.description ||
            'Unknown error'
        );
    };

    return (
        <>
            <RolesVersionListComponent
                // here we need to copy array,
                // so RolesVersionListComponent is going to be re-rendered after list update
                rolesVersions={_.slice(rolesVersionStore.list)}
                tryEdit={rv => {
                    setEditingRolesVersion(rv);
                    e.emit(openEditResourceEventName);
                }}
                tryDelete={rv => {
                    setDeletingRolesVersion(rv);
                    e.emit(openPrompterEventName);
                }}
            />

            <EditRolesVersion
                openDialogueEventName={openEditResourceEventName}
                eventEmitter={e}
                save={rv => rolesVersionStore.updateRolesVersion(rv)}
                initialRolesVersion={editingRolesVersion}
            />

            <Prompter
                title={'Roles version deletion'}
                description={'Are you sure you want to delete this roles version? This will cause deletion of all roles linked to this version.'}
                onAgree={() => deleteRolesVersion()}
                onDisagree={() => setDeletingRolesVersion(null)}
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
