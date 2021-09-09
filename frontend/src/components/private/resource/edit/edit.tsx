import { useState, useEffect } from 'react';
import {
    Button,
    Dialog,
    DialogActions,
    DialogContent,
    DialogContentText,
    DialogTitle,
} from '@material-ui/core';

import { EditResourceProps } from './edit.types';
import { TextFields } from '../../shared';

export const EditResource = (props: EditResourceProps) => {
    const [open, setOpen] = useState(false);

    const openDialogue = () => setOpen(true);
    props.eventEmitter.on(props.openDialogueEventName, openDialogue);
    useEffect(() => {
        return () => {
            props.eventEmitter.off(props.openDialogueEventName, openDialogue);
        };
    });

    const [resourceId, setResourceId] = useState(props.initialResource?.id || '');
    const [resourceTitle, setResourceTitle] = useState(props.initialResource?.title || '');

    useEffect(() => {
        !resourceId && setResourceId(props.initialResource?.id || '');
        (!resourceTitle || props.initialResource?.title) && setResourceTitle(props.initialResource?.title || '');
        // eslint-disable-next-line
    }, [props.initialResource?.id, props.initialResource?.title]);

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

                <TextFields
                    fields={[
                        {
                            name: 'id',
                            value: resourceId,
                            label: 'Resource Id',
                            disabled: !!props.initialResource?.id,
                            onChange: newValue => setResourceId(newValue),
                        },
                        {
                            name: 'title',
                            value: resourceTitle,
                            label: 'Resource Title',
                            onChange: newValue => setResourceTitle(newValue),
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
}
