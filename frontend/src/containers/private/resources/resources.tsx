import { observer } from 'mobx-react-lite';
import { Box } from '@material-ui/core';
import { Add } from '@material-ui/icons';
import EventEmitter from 'events';

import { Alert } from '../../../components/shared/';
import { ResourcesList } from './list';
import { EditResource as CreateResource } from '../../../components/private/resource';
import { resourceStore } from '../../../store';

export const Resources = observer(() => {
    const e = new EventEmitter();
    const openDialogueEventName = 'new-resource-dialogue:open';

    return (
        <>
            <Box>
                <h1>
                    Resources
                    <Add
                        color="primary"
                        cursor="pointer"
                        fontSize="large"
                        titleAccess="Add new resource"
                        onClick={() => e.emit(openDialogueEventName)}
                    />
                </h1>
            </Box>
            <Box>
                <ResourcesList />
            </Box>

            <CreateResource
                eventEmitter={e}
                openDialogueEventName={openDialogueEventName}
                save={r => resourceStore.createResource(r)}
            />

            <Alert
                shouldShow={!!resourceStore.createResourceError}
                severity={'error'}
                message={resourceStore.createResourceError?.description || 'Unknown error'}
                onCloseCb={() => resourceStore.cleanResourceActionError()}
            />
        </>
    );
});
