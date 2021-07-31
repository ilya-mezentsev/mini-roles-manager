import { useState, useEffect } from 'react';
import {
    Button,
    Dialog,
    DialogActions,
    DialogContent,
    DialogContentText,
    DialogTitle,
    TextField,
} from '@material-ui/core';

import { EditResourceProps } from './edit.types';

export const EditResource = (props: EditResourceProps) => {
    const { id: initialResourceId, title: initialResourceTitle } = props.initialResource || {
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

    const [resourceId, setResourceId] = useState('');
    const [resourceTitle, setResourceTitle] = useState('');

    useEffect(
        () => {
            setResourceId(initialResourceId || '');
            setResourceTitle(initialResourceTitle || '');
        },
        [initialResourceId, initialResourceTitle],
    );

    const handleClose = () => {
        setOpen(false);
        setResourceId('');
        setResourceTitle('');
    };

    const handleSave = () => {
        props.save({
            id: resourceId,
            title: resourceTitle,
        });
        handleClose();
    };

    return (
        <Dialog open={open} onClose={handleClose} aria-labelledby="form-dialog-title">
            <DialogTitle id="form-dialog-title">New Resource</DialogTitle>
            <DialogContent>
                <DialogContentText>
                    Enter new resource data below
                </DialogContentText>

                <TextField
                    margin="dense"
                    label="Resource Id"
                    fullWidth
                    disabled={!!initialResourceId}
                    value={resourceId}
                    onChange={e => setResourceId((e.target as HTMLInputElement).value)}
                />

                <TextField
                    margin="dense"
                    label="Resource Title"
                    fullWidth
                    value={resourceTitle}
                    onChange={e => setResourceTitle((e.target as HTMLInputElement).value)}
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
}
