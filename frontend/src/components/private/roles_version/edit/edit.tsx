import { useState, useEffect } from 'react';
import {
    Button,
    Dialog,
    DialogActions,
    DialogContent,
    DialogContentText,
    DialogTitle,
} from '@material-ui/core';

import { EditRolesVersionProps } from './edit.types';
import { TextFields } from '../../shared';

export const EditRolesVersion = (props: EditRolesVersionProps) => {
    const [open, setOpen] = useState(false);

    const openDialogue = () => setOpen(true);
    props.eventEmitter.on(props.openDialogueEventName, openDialogue);
    useEffect(() => {
        return () => {
            props.eventEmitter.off(props.openDialogueEventName, openDialogue);
        };
    });

    const [rolesVersionId, setRolesVersionId] = useState(props.initialRolesVersion?.id || '');
    const [rolesVersionTitle, setRolesVersionTitle] = useState(props.initialRolesVersion?.title || '');

    useEffect(() => {
        !rolesVersionId && setRolesVersionId(props.initialRolesVersion?.id || '');
        (!rolesVersionTitle || props.initialRolesVersion?.title) && setRolesVersionTitle(props.initialRolesVersion?.title || '');
        // eslint-disable-next-line
    }, [props.initialRolesVersion?.id, props.initialRolesVersion?.title]);

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
            <DialogTitle id="form-dialog-title">New Roles Version</DialogTitle>
            <DialogContent>
                <DialogContentText>
                    Enter new roles version data below
                </DialogContentText>

                <TextFields
                    fields={[
                        {
                            name: 'id',
                            value: rolesVersionId,
                            label: 'Roles Version Id',
                            disabled: !!props.initialRolesVersion?.id,
                            onChange: newValue => setRolesVersionId(newValue),
                        },
                        {
                            name: 'title',
                            value: rolesVersionTitle,
                            label: 'Roles Version Title',
                            onChange: newValue => setRolesVersionTitle(newValue),
                        },
                    ]}
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
