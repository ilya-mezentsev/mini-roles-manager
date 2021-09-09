import { useEffect, useState } from 'react';
import {
    Button,
    Dialog,
    DialogActions,
    DialogContent,
    DialogContentText,
    DialogTitle,
} from '@material-ui/core';

import { PrompterProps } from './prompter.types';

export const Prompter = (props: PrompterProps) => {
    const [open, setOpen] = useState(false);

    const openDialogue = () => setOpen(true);
    props.eventEmitter.on(props.openDialogueEventName, openDialogue);
    useEffect(() => {
        return () => {
            props.eventEmitter.off(props.openDialogueEventName, openDialogue);
        };
    });

    const handleClose = () => {
        setOpen(false);
    }

    const handleAgreed = () => {
        handleClose();
        props.onAgree();
    }

    const handleDisagreed = () => {
        handleClose();
        props.onDisagree();
    }

    return (
        <Dialog open={open} onClose={handleClose}>
            <DialogTitle id="form-dialog-title">{props.title}</DialogTitle>
            <DialogContent>
                <DialogContentText>
                    {props.description}
                </DialogContentText>
            </DialogContent>
            <DialogActions>
                <Button onClick={handleDisagreed} color="primary">
                    Disagree
                </Button>
                <Button onClick={handleAgreed} color="primary">
                    Agree
                </Button>
            </DialogActions>
        </Dialog>
    );
}
