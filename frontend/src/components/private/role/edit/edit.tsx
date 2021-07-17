import { useEffect, useState } from 'react';
import * as _ from 'lodash';
import {
    Button,
    Chip,
    Dialog,
    DialogActions,
    DialogContent,
    DialogContentText,
    DialogTitle,
    InputLabel,
    Input,
    MenuItem,
    Select,
    TextField,
} from '@material-ui/core';

import { Permission } from '../../../../services/api';
import { EditRoleProps } from './edit.types';
import { Permissions } from './permissions';

export const EditRole = (props: EditRoleProps) => {
    const [open, setOpen] = useState(false);

    const initialRoleId = props.initialRole?.id || '';
    const initialTitle = props.initialRole?.title || '';
    let permissions = props?.initialRole?.permissions || [];
    const initialExtends = props.initialRole?.extends || [];

    const openDialogue = () => setOpen(true);
    props.eventEmitter.on(props.openDialogueEventName, openDialogue);
    useEffect(() => {
        return () => {
            props.eventEmitter.off(props.openDialogueEventName, openDialogue);
        };
    });

    const [roleId, setRoleId] = useState('');
    const [roleTitle, setRoleTitle] = useState('');
    const [extends_, setExtends] = useState<string[]>([]);
    const canExtendsFrom = props.existRoles.filter(r => r.id !== initialRoleId && r.versionId === props.roleVersionId);

    useEffect(() => {
        (!roleId || initialRoleId) && setRoleId(initialRoleId || '');
        (!roleTitle || initialTitle) && setRoleTitle(initialTitle || '');

        (!extends_ || !_.isEmpty(initialExtends)) && setExtends(initialExtends || []);
        // eslint-disable-next-line
    }, [initialRoleId, initialTitle, initialExtends]);

    const onPermissionsUpdate = (updatedPermissions: Permission[]) => {
        permissions = updatedPermissions.map(p => p.id);
    }

    const handleClose = () => {
        setOpen(false);
        setRoleId('');
        setRoleTitle('');
        setExtends([]);
        permissions = [];
    };

    const handleSave = () => {
        props.save({
            id: roleId,
            versionId: props.roleVersionId,
            title: roleTitle,
            permissions,
            extends: extends_,
        });
        handleClose();
    };

    return (
        <Dialog
            fullWidth={true}
            maxWidth="lg"
            open={open}
            onClose={handleClose}
            aria-labelledby="form-dialog-title"
        >
            <DialogTitle id="form-dialog-title">New Role</DialogTitle>
            <DialogContent>
                <DialogContentText>
                    Enter new role data below
                </DialogContentText>

                <TextField
                    margin="dense"
                    label="Role Id"
                    fullWidth
                    disabled={!!initialRoleId}
                    value={roleId}
                    onChange={e => setRoleId((e.target as HTMLInputElement).value)}
                />

                <TextField
                    margin="dense"
                    label="Role Title"
                    fullWidth
                    value={roleTitle}
                    onChange={e => setRoleTitle((e.target as HTMLInputElement).value)}
                />

                <InputLabel id="select-role-extends">Extends from</InputLabel>
                <Select
                    labelId="select-role-extends"
                    disabled={canExtendsFrom.length < 1}
                    fullWidth
                    multiple
                    value={extends_}
                    onChange={e => setExtends((e.target as HTMLSelectElement).value as unknown as string[])}
                    input={<Input />}
                    renderValue={(selected: unknown) => (
                        <div>
                            {(selected as string[]).map((value) => (
                                <Chip key={value} label={value} />
                            ))}
                        </div>
                    )}
                >
                    {canExtendsFrom.map((r, index) => (
                        <MenuItem key={`role_item_${index}`} value={r.id}>
                            {r.id}
                        </MenuItem>
                    ))}
                </Select>

                <Permissions
                    resources={props.existsResources}
                    existsPermissions={
                        (permissions || [])
                            .map(
                                permissionId => props.existsResources
                                    .find(r => r.permissions.find(p => p.id === permissionId))!.permissions
                                    .find(p => p.id === permissionId)
                            ) as Permission[]
                    }
                    onPermissionsUpdate={onPermissionsUpdate}
                />
            </DialogContent>
            <DialogActions>
                <Button onClick={handleClose} color="primary">
                    Cancel
                </Button>
                <Button onClick={handleSave} color="primary">
                    Save
                </Button>
            </DialogActions>
        </Dialog>
    );
};
