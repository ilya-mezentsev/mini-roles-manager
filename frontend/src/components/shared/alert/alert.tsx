import { useState, useEffect } from 'react';
import { Snackbar } from '@material-ui/core'
import MuiAlert from '@material-ui/lab/Alert';
import { AlertProps } from './alert.types';

export const Alert = (props: AlertProps) => {
    const transitionDuration = 200;
    const [open, setOpen] = useState(false);

    useEffect(
        () => setOpen(props.shouldShow),
        [props.shouldShow],
    );

    const handleClose = () => {
        setOpen(false);
        setTimeout(() => props.onCloseCb(), transitionDuration + 50);
    };

    return (
        <div>
            <Snackbar
                open={open && props.shouldShow}
                anchorOrigin={{
                    vertical: 'top',
                    horizontal: 'right',
                }}
                transitionDuration={transitionDuration}
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
