import React from 'react';
import { Snackbar } from '@material-ui/core'
import MuiAlert from '@material-ui/lab/Alert';
import { AlertProps } from './alert.types';

export const Alert = (props: AlertProps) => {
    const [open, setOpen] = React.useState(false);

    props.setOpenEmitter.on(props.setOpenEventName, () => setOpen(true));

    const handleClose = () => {
        setOpen(false);
        // fixme - если делать без тайм-аута, то на какой-то момент алерт станет зеленым
        setTimeout(() => props.onCloseCb(), 100);
    };

    return (
        <div>
            <Snackbar
                open={open && props.shouldShow}
                anchorOrigin={{
                    vertical: 'top',
                    horizontal: 'right',
                }}
            >
                <MuiAlert
                    elevation={6}
                    variant="filled"
                    onClose={handleClose}
                    severity={props.severity}
                >
                    {props.message}
                </MuiAlert>
            </Snackbar>
        </div>
    );
}
