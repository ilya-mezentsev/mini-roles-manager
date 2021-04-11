import { useState } from 'react';
import _ from 'lodash';
import {
    Input,
    InputLabel,
    MenuItem,
    Select,
    Table,
    TableBody,
    TableCell,
    TableContainer,
    TableHead,
    TableRow,
    TextField,
} from '@material-ui/core';
import { Add, Delete } from '@material-ui/icons';
import { Autocomplete } from '@material-ui/lab';

import { Effect, Operation, Permission } from '../../../../services/api';
import { PermissionsProps, ResourcePermission } from './permissions.types';

const operations = [
    Operation.CREATE,
    Operation.READ,
    Operation.UPDATE,
    Operation.DELETE,
];
const effects = [
    Effect.PERMIT,
    Effect.DENY,
];

export const Permissions = (props: PermissionsProps) => {
    const existsPermissions = props.existsPermissions || [];
    const [existsRolePermissions, setExistsRolePermissions] = useState(new Map<string, Permission[]>(
        props.resources
            .filter(r => _.intersection(r.permissions, existsPermissions).length > 0)
            .map(r => [
                r.id,
                r.permissions.filter(p => existsPermissions.find(ep => ep.id === p.id)),
            ])
    ));

    const [newResourceCreatePermission, setNewResourceCreatePermission] = useState<ResourcePermission>({ operation: Operation.CREATE, effect: Effect.DENY });
    const [newResourceReadPermission, setNewResourceReadPermission] = useState<ResourcePermission>({ operation: Operation.READ, effect: Effect.DENY });
    const [newResourceUpdatePermission, setNewResourceUpdatePermission] = useState<ResourcePermission>({ operation: Operation.UPDATE, effect: Effect.DENY });
    const [newResourceDeletePermission, setNewResourceDeletePermission] = useState<ResourcePermission>({ operation: Operation.DELETE, effect: Effect.DENY });
    const newResourcePermissions = [
        { permission: newResourceCreatePermission, setPermission: setNewResourceCreatePermission },
        { permission: newResourceReadPermission, setPermission: setNewResourceReadPermission },
        { permission: newResourceUpdatePermission, setPermission: setNewResourceUpdatePermission },
        { permission: newResourceDeletePermission, setPermission: setNewResourceDeletePermission },
    ];

    const [newResourcePermissionsId, setNewResourcePermissionsId] = useState('');
    const onNewResourceIdChanged = (value: string | null) => {
        setNewResourcePermissionsId(value || '');

        resetNewResourcePermissions();
    };
    const resetNewResourcePermissions = () => {
        for (const newResourcePermission of newResourcePermissions) {
            newResourcePermission.setPermission({
                ...newResourcePermission.permission,
                effect: Effect.DENY,
            });
        }
    };
    const cleanNewResourcePermissions = () => {
        setNewResourcePermissionsId('');
        resetNewResourcePermissions();
    };

    const chosenResource = () => props.resources.find(r => r.id === newResourcePermissionsId);
    const chosenExistsResourceId = () => !!chosenResource();

    const updatePermissions = (permissions: Map<string, Permission[]>) => {
        props.onPermissionsUpdate(Array.from(permissions.values()).flat());
    };

    const addPermission = (resourceId: string, permissions: Permission[]) => {
        setExistsRolePermissions(prev => {
            const updatedPermissions = new Map([
                ...Array.from(prev.entries()),
                [
                    resourceId,
                    permissions,
                ],
            ]);
            updatePermissions(updatedPermissions);

            return updatedPermissions;
        });
    };
    const addNewPermission = () => {
        addPermission(
            newResourcePermissionsId,
            newResourcePermissions.map(newResourcePermission => ({
                id: props.resources
                    .find(r => r.id === newResourcePermissionsId)!.permissions!
                    .find(p => (
                        p.effect === newResourcePermission.permission.effect &&
                        p.operation === newResourcePermission.permission.operation
                    ))!.id,
                effect: newResourcePermission.permission.effect,
                operation: newResourcePermission.permission.operation,
            })),
        );
    };
    const updatePermission = (resourceId: string, updatedPermission: Permission) => {
        setExistsRolePermissions(prev => {
            const currentResourcePermissions = prev.get(resourceId)!;
            const updatedPermissionIndex = currentResourcePermissions.findIndex(p => p.id === updatedPermission.id);
            if (updatedPermissionIndex >= 0) {
                currentResourcePermissions[updatedPermissionIndex] = {
                    ...updatedPermission,
                    id: props.resources
                        .find(r => r.id === resourceId)!.permissions!
                        .find(p => (
                            p.effect === updatedPermission.effect &&
                            p.operation === updatedPermission.operation
                        ))!.id,
                };
            }

            const updatedPermissions = new Map(Array.from(prev.entries())).set(resourceId, currentResourcePermissions);
            updatePermissions(updatedPermissions);

            return updatedPermissions;
        });
    };
    const deletePermissions = (resourceId: string) => {
        setExistsRolePermissions(prev => {
            const newExistsRolePermissions = new Map(Array.from(prev.entries()));
            newExistsRolePermissions.delete(resourceId);
            updatePermissions(newExistsRolePermissions);

            return newExistsRolePermissions;
        });
    };

    return (
        <TableContainer>
            <Table stickyHeader aria-label="sticky table">
                <TableHead>
                    <TableRow>
                        <TableCell width="25%">
                            Resource
                        </TableCell>
                        {
                            operations.map((o, index) => (
                                <TableCell
                                    align="center"
                                    key={`operation_name_${index}`}
                                >
                                    {o}
                                </TableCell>
                            ))
                        }
                        <TableCell>
                            Action
                        </TableCell>
                    </TableRow>
                </TableHead>
                <TableBody>
                    <TableRow hover tabIndex={-1}>
                        <TableCell width="25%">
                            <Autocomplete
                                id="new-resource-permissions-input"
                                disabled={props.resources.length < 1}
                                options={props.resources.map(r => r.id).filter(resourceId => !Array.from(existsRolePermissions.keys()).includes(resourceId))}
                                value={newResourcePermissionsId}
                                onChange={(_, newValue) => onNewResourceIdChanged(newValue)}
                                fullWidth
                                getOptionLabel={option => option}
                                renderInput={(params) => (
                                    <TextField {...params} label="Resource Id" variant="outlined" />
                                )}
                            />
                        </TableCell>

                        {
                            newResourcePermissions.map((p, index) => (
                                <TableCell key={`new_resource_permission_${index}`}>
                                    <InputLabel>Effect</InputLabel>
                                    <Select
                                        key={`new_resource_permission_effect_${index}`}
                                        disabled={!chosenExistsResourceId()}
                                        fullWidth
                                        value={p.permission.effect}
                                        onChange={e => p.setPermission({
                                            ...p.permission,
                                            effect: (e.target as HTMLSelectElement).value as Effect,
                                        })}
                                        input={<Input />}
                                    >
                                        {effects.map(e => (
                                            <MenuItem key={`operation_${p.permission.operation}_${e}`} value={e}>
                                                {e}
                                            </MenuItem>
                                        ))}
                                    </Select>
                                </TableCell>
                            ))
                        }

                        <TableCell>
                            <Add
                                color="primary"
                                cursor="pointer"
                                titleAccess="Add permission"
                                onClick={() => {
                                    addNewPermission();
                                    cleanNewResourcePermissions();
                                }}
                            />
                        </TableCell>
                    </TableRow>

                    {
                        Array.from(existsRolePermissions.entries()).map(([resourceId, resourcePermissions], index) => (
                            <TableRow key={`role_permissions_${index}`}>
                                <TableCell width="25%">
                                    <TextField
                                        disabled
                                        fullWidth
                                        value={resourceId}
                                    />
                                </TableCell>

                                {
                                    resourcePermissions.map((p, index) => (
                                        <TableCell key={`role_permission_${resourceId}_${p.operation}`}>
                                            <InputLabel>Effect</InputLabel>
                                            <Select
                                                key={`new_resource_permission_effect_${index}`}
                                                fullWidth
                                                value={p.effect}
                                                onChange={e => {
                                                    p.effect = (e.target as HTMLSelectElement).value as Effect;
                                                    updatePermission(resourceId, p);
                                                }}
                                                input={<Input />}
                                            >
                                                {effects.map(e => (
                                                    <MenuItem key={`operation_create_${e}`} value={e}>
                                                        {e}
                                                    </MenuItem>
                                                ))}
                                            </Select>
                                        </TableCell>
                                    ))
                                }

                                <TableCell>
                                    <Delete
                                        color="primary"
                                        cursor="pointer"
                                        titleAccess="Delete permission"
                                        onClick={() => deletePermissions(resourceId)}
                                    />
                                </TableCell>
                            </TableRow>
                        ))
                    }
                </TableBody>
            </Table>
        </TableContainer>
    );
}
