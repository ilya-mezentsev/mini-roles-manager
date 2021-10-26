import { useState } from 'react';
import { observer } from 'mobx-react-lite';
import EventEmitter from 'events';

import { Alert } from '../../../components/shared';
import { Prompter } from '../../../components/private/shared';
import {
    EditResource,
    ResourcesList as ResourcesListComponent,
} from '../../../components/private/resource';
import { Resource } from '../../../services/api';
import { resourceStore, roleStore } from '../../../store';

export const ResourcesList = observer(() => {
    const [deletingResource, setDeletingResource] = useState<Resource | null>(null);
    const [editingResource, setEditingResource] = useState<Resource | null>(null);

    const e = new EventEmitter();
    const openPrompterEventName = 'prompter:open';
    const openEditResourceEventName = 'edit-resource:open';

    const deleteResource = () => {
        if (deletingResource) {
            resourceStore
                .deleteResource(deletingResource.id)
                .finally(() => setDeletingResource(null));
        }
    };

    const hasAnyError: () => boolean = () => {
        return (
            !!resourceStore.fetchResourceError ||
            !!resourceStore.updateResourceError ||
            !!resourceStore.deleteResourceError ||
            !!roleStore.fetchRoleError
        );
    };

    const cleanErrors = () => {
        resourceStore.cleanResourceActionError();
        roleStore.cleanRoleActionErrors();
    };

    const errorMessage: () => string = () => {
        return (
            resourceStore.fetchResourceError?.description ||
            resourceStore.updateResourceError?.description ||
            resourceStore.deleteResourceError?.description ||
            roleStore.fetchRoleError?.description ||
            'Unknown error'
        );
    };

    return (
        <>
            <ResourcesListComponent
                resources={resourceStore.list || []}
                tryEdit={r => {
                    setEditingResource(r);
                    e.emit(openEditResourceEventName);
                }}
                tryDelete={r => {
                    setDeletingResource(r);
                    e.emit(openPrompterEventName);
                }}
            />

            <EditResource
                openDialogueEventName={openEditResourceEventName}
                eventEmitter={e}
                save={r => resourceStore.updateResource(r)}
                initialResource={editingResource}
            />

            <Prompter
                title="Resource deletion"
                description="Are you sure you want to delete this resource?"
                onAgree={() => deleteResource()}
                onDisagree={() => setDeletingResource(null)}
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
