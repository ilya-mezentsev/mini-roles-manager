import { useState, useEffect } from 'react';
import {
    Button,
    Dialog,
    DialogActions,
    DialogContent,
    DialogContentText,
    DialogTitle,
    TextField
} from '@material-ui/core';

import { EditRolesVersionProps } from './edit.types';

export const EditRolesVersion = (props: EditRolesVersionProps) => {
    const { id: initialRolesVersionId, title: initialRolesVersionTitle } = props.initialRolesVersion || {
        id: '',
        title: '',
    };

    const [open, setOpen] = useState(false);

    const openDialogue = () => setOpen(true);
    props.eventEmitter.on(props.openDialogueEventName, openDialogue);
    useEffect(() => {
        return () => {
            props.eventEmitter.off(props.openDialogueEventName, openDialogue);
        };
    });

    const [rolesVersionId, setRolesVersionId] = useState('');
    const [rolesVersionTitle, setRolesVersionTitle] = useState('');

    useEffect(
        () => {
            setRolesVersionId(initialRolesVersionId || '');
            setRolesVersionTitle(initialRolesVersionTitle || '');
        },
        [initialRolesVersionId, initialRolesVersionTitle],
    );

    const handleClose = () => {
        setOpen(false);
        setRolesVersionId('');
        setRolesVersionTitle('');
    };

    const handleSave = () => {
        props.save({
            id: rolesVersionId,
            title: rolesVersionTitle,
        });
        handleClose();
    };

    return (
        <Dialog open={open} onClose={handleClose} aria-labelledby="form-dialog-title">
            <DialogTitle id="form-dialog-title">New RolesVersion</DialogTitle>
            <DialogContent>
                <DialogContentText>
                    Enter new roles version data below
                </DialogContentText>

                <TextField
                    margin="dense"
                    label="Roles Version Id"
                    fullWidth
                    disabled={!!initialRolesVersionId}
                    value={rolesVersionId}
                    onChange={e => setRolesVersionId((e.target as HTMLInputElement).value)}
                />

                <TextField
                    margin="dense"
                    label="Roles Version Title"
                    fullWidth
                    value={rolesVersionTitle}
                    onChange={e => setRolesVersionTitle((e.target as HTMLInputElement).value)}
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
