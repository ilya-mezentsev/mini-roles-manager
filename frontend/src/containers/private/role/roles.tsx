import { observer } from 'mobx-react-lite';
import {
    Box,
    Input,
    InputLabel,
    MenuItem,
    Select,
} from '@material-ui/core';
import { Add } from '@material-ui/icons';
import EventEmitter from 'events';

import { Alert } from '../../../components/shared';
import { EditRole as CreateRole } from '../../../components/private/role';
import { RolesList } from './list';
import {
    rolesVersionStore,
    resourceStore,
    roleStore,
} from '../../../store';

export const Roles = observer(() => {
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
                <InputLabel>Roles version</InputLabel>
                <Select
                    margin="dense"
                    fullWidth
                    value={rolesVersionStore.current?.id || ''}
                    onChange={e => rolesVersionStore.setCurrentRolesVersion(
                        rolesVersionStore.list!.find(rv => rv.id === (e.target as HTMLSelectElement).value)!
                    )}
                    input={<Input />}
                >
                    {rolesVersionStore.list.map(rv => (
                        <MenuItem key={`operation_${rv.id}`} value={rv.id}>
                            {rv.id}
                        </MenuItem>
                    ))}
                </Select>
            </Box>

            <Box>
                <RolesList/>
            </Box>

            <CreateRole
                eventEmitter={e}
                openDialogueEventName={openDialogueEventName}
                existRoles={roleStore.list}
                existsResources={resourceStore.list}
                roleVersionId={rolesVersionStore.current?.id || ''}
                save={r => roleStore.createRole(r)}
            />

            <Alert
                shouldShow={!!roleStore.createRoleError}
                severity={'error'}
                message={roleStore.createRoleError?.description || 'Unknown error'}
                onCloseCb={() => roleStore.cleanRoleActionErrors()}
            />
        </>
    )
});
