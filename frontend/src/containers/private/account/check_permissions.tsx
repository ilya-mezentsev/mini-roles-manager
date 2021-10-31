import { useEffect, useState } from 'react';
import { observer } from 'mobx-react-lite';
import {
    Button,
    Input,
    InputLabel,
    MenuItem,
    Select,
    TextField,
    Box,
} from '@material-ui/core';
import { Autocomplete } from '@material-ui/lab';

import { Alert } from '../../../components/shared';
import { Operation } from '../../../services/api/shared/types';
import {
    resourceStore,
    roleStore,
    rolesVersionStore,
    permissionStore,
} from '../../../store';

const operations = [
    Operation.CREATE,
    Operation.READ,
    Operation.UPDATE,
    Operation.DELETE,
];

export const CheckPermissions = observer(() => {
    const [rolesVersionId, setRolesVersionId] = useState('');
    const [roleId, setRoleId] = useState('');
    const [resourceId, setResourceId] = useState('');
    const [operation, setOperation] = useState(Operation.CREATE);

    // eslint-disable-next-line
    useEffect(() => () => permissionStore.cleanFetchPermissionAction(), []);

    return (
        <>
            <h3>Check permissions:</h3>

            <Autocomplete
                disabled={resourceStore.list.length < 1}
                options={resourceStore.list.map(r => r.id)}
                value={resourceId}
                onChange={(_, newValue) => setResourceId((newValue as string) || '')}
                fullWidth
                getOptionLabel={option => option}
                renderInput={(params) => (
                    <TextField {...params} label="Resource Id" variant="outlined" margin="dense" />
                )}
            />

            <Autocomplete
                disabled={roleStore.list.length < 1}
                options={
                    Array.from(new Set(
                        roleStore.list.map(r => r.id),
                    ))
                }
                value={roleId}
                onChange={(_, newValue) => setRoleId((newValue as string) || '')}
                fullWidth
                getOptionLabel={option => option}
                renderInput={(params) => (
                    <TextField {...params} label="Role Id" variant="outlined" margin="dense" />
                )}
            />

            <InputLabel>Roles version</InputLabel>
            <Select
                margin="dense"
                fullWidth
                value={rolesVersionId || rolesVersionStore.current?.id || ''}
                onChange={e => setRolesVersionId((e.target as HTMLSelectElement).value)}
                input={<Input />}
            >
                {rolesVersionStore.list.map(rv => (
                    <MenuItem key={`operation_${rv.id}`} value={rv.id}>
                        {rv.id}
                    </MenuItem>
                ))}
            </Select>

            <InputLabel>Effect</InputLabel>
            <Select
                margin="dense"
                fullWidth
                value={operation}
                onChange={e => setOperation((e.target as HTMLSelectElement).value as Operation)}
                input={<Input />}
            >
                {operations.map(o => (
                    <MenuItem key={`operation_${o}`} value={o}>
                        {o}
                    </MenuItem>
                ))}
            </Select>

            <Box mt={3}>
                <Button
                    disabled={!roleId || !resourceId || !operation}
                    variant="contained"
                    color="primary"
                    onClick={() => {
                        permissionStore.cleanFetchPermissionAction();
                        return permissionStore.fetchPermission({
                            roleId,
                            resourceId,
                            operation,
                            rolesVersionId: rolesVersionId || rolesVersionStore.current?.id || '',
                        });
                    }}
                >
                    Check
                </Button>
            </Box>

            <Alert
                message={`Effect is ${(permissionStore.permission?.effect || 'Unknown').toUpperCase()}`}
                severity="info"
                shouldShow={!!permissionStore.permission?.effect}
                onCloseCb={() => permissionStore.cleanFetchPermissionAction()}
            />

            <Alert
                message={permissionStore.fetchPermissionError?.description || 'Unknown error'}
                severity="error"
                shouldShow={!!permissionStore.fetchPermissionError}
                onCloseCb={() => permissionStore.cleanFetchPermissionAction()}
            />
        </>
    );
});
