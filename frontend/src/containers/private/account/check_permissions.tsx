import { useEffect, useState } from 'react';
import { bindActionCreators } from 'redux';
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
import {
    CheckPermissionsProps,
    CheckPermissionsActions,
    CheckPermissionsState,
} from './check_permissions.types';
import { DispatchToPropsFn, StateToPropsFn } from '../../../shared/types';
import { cleanFetchPermissionResult, fetchPermission } from '../../../store/permission/actions';
import { Operation } from '../../../services/api/shared/types';

const operations = [
    Operation.CREATE,
    Operation.READ,
    Operation.UPDATE,
    Operation.DELETE,
];

export const CheckPermissions = (props: CheckPermissionsProps) => {
    const [rolesVersionId, setRolesVersionId] = useState('');
    const [roleId, setRoleId] = useState('');
    const [resourceId, setResourceId] = useState('');
    const [operation, setOperation] = useState(Operation.CREATE);

    // eslint-disable-next-line
    useEffect(() => () => props.cleanFetchPermissionResult(), []);

    return (
        <>
            <h3>Check permissions:</h3>

            <Autocomplete
                disabled={(props.resourcesResult.list?.length || 0) < 1}
                options={(props.resourcesResult.list || []).map(r => r.id)}
                value={resourceId}
                onChange={(_, newValue) => setResourceId(newValue || '')}
                fullWidth
                getOptionLabel={option => option}
                renderInput={(params) => (
                    <TextField {...params} label="Resource Id" variant="outlined" margin="dense" />
                )}
            />

            <Autocomplete
                disabled={(props.rolesResult.list?.length || 0) < 1}
                options={
                    (props.rolesResult.list || [])
                        .filter(r => r.versionId === props.rolesVersionResult.currentRolesVersion?.id)
                        .map(r => r.id)
                }
                value={roleId}
                onChange={(_, newValue) => setRoleId(newValue || '')}
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
                value={rolesVersionId || props.rolesVersionResult.currentRolesVersion?.id || ''}
                onChange={e => setRolesVersionId((e.target as HTMLSelectElement).value)}
                input={<Input />}
            >
                {(props.rolesVersionResult.list || []).map(rv => (
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
                        props.cleanFetchPermissionResult();
                        props.fetchPermissionAction({
                            roleId,
                            resourceId,
                            operation,
                            rolesVersionId: rolesVersionId || props.rolesVersionResult.currentRolesVersion?.id || '',
                        });
                    }}
                >
                    Check
                </Button>
            </Box>

            <Alert
                message={`Effect is ${(props.fetchPermissionResult?.effect || 'Unknown').toUpperCase()}`}
                severity="info"
                shouldShow={!!props.fetchPermissionResult?.effect}
                onCloseCb={() => props.cleanFetchPermissionResult()}
            />

            <Alert
                message={props.fetchPermissionResult?.error?.description || 'Unknown error'}
                severity="error"
                shouldShow={!!props.fetchPermissionResult?.error}
                onCloseCb={() => props.cleanFetchPermissionResult()}
            />
        </>
    );
}

export const mapDispatchToProps: DispatchToPropsFn<CheckPermissionsActions> = () => dispatch => ({
    fetchPermissionAction: bindActionCreators(fetchPermission, dispatch),
    cleanFetchPermissionResult: bindActionCreators(cleanFetchPermissionResult, dispatch),
});

export const mapStateToProps: StateToPropsFn<CheckPermissionsState> = () => state => ({
    fetchPermissionResult: state.fetchPermissionResult,
    rolesVersionResult: state.rolesVersionResult,
    resourcesResult: state.resourcesResult,
    rolesResult: state.rolesResult,
});
