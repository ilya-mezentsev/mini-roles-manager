import { observer } from 'mobx-react-lite';
import { Box } from '@material-ui/core';
import { Add } from '@material-ui/icons';
import EventEmitter from 'events';

import { RolesVersionList } from './list';
import {
    EditRolesVersion as CreateRolesVersion,
} from '../../../components/private/roles_version';
import { Alert } from '../../../components/shared';
import { rolesVersionStore } from '../../../store';

export const RolesVersion = observer(() => {
    const e = new EventEmitter();
    const openDialogueEventName = 'new-roles-version-dialogue:open';

    return (
        <>
            <Box>
                <h1>
                    Roles Versions
                    <Add
                        color="primary"
                        cursor="pointer"
                        fontSize="large"
                        titleAccess="Add new roles version"
                        onClick={() => e.emit(openDialogueEventName)}
                    />
                </h1>
            </Box>

            <Box>
                <RolesVersionList/>
            </Box>

            <CreateRolesVersion
                openDialogueEventName={openDialogueEventName}
                eventEmitter={e}
                save={rv => rolesVersionStore.createRolesVersion(rv)}
            />

            <Alert
                shouldShow={!!rolesVersionStore.createRolesVersionError}
                severity={'error'}
                message={rolesVersionStore.createRolesVersionError?.description || 'Unknown error'}
                onCloseCb={() => rolesVersionStore.cleanRolesVersionActionErrors()}
            />
        </>
    );
});
