import { useState, useEffect } from 'react';
import { bindActionCreators } from 'redux';
import {
    Button,
    TextField,
} from '@material-ui/core';

import { Alert } from '../../../components/shared';
import { DispatchToPropsFn, StateToPropsFn } from '../../../shared/types';
import { CredentialsProps, CredentialsState, CredentialsActions } from './credentials.types';
import {
    cleanUpdateCredentialsResult,
    updateCredentials,
} from '../../../store/account_info/actions';

export const Credentials = (props: CredentialsProps) => {
    const [login, setLogin] = useState('');
    const [password, setPassword] = useState('');

    useEffect(() => {
        setLogin(props.accountInfoResult?.info?.login || '');
    }, [props.accountInfoResult?.info?.login]);

    // eslint-disable-next-line
    useEffect(() => () => props.cleanUpdateCredentialsErrorAction(), []);

    const updateCredentials = () => {
        props.updateCredentialsAction({
            login,
            password: password || '',
        });
    };

    return (
        <>
            <h2>Account credentials:</h2>
            <TextField
                label="New Login"
                margin="dense"
                fullWidth
                value={login}
                onChange={e => setLogin((e.target as HTMLInputElement).value)}
            />
            <TextField
                label="New Password"
                required
                fullWidth
                margin="normal"
                type="password"
                value={password}
                onChange={e => setPassword((e.target as HTMLInputElement).value)}
            />

            <Button
                disabled={login === props.accountInfoResult?.info?.login && !password}
                variant="contained"
                color="primary"
                onClick={updateCredentials}
            >
                Save
            </Button>

            <Alert
                shouldShow={!!props.accountInfoResult?.credentials}
                severity="success"
                message="Credentials are updated successfully"
                onCloseCb={() => {
                    props.cleanUpdateCredentialsErrorAction();
                    setPassword('');
                }}
            />

            <Alert
                shouldShow={!!props.accountInfoResult?.updateCredentialsError}
                severity="error"
                message={props.accountInfoResult?.updateCredentialsError?.description || 'Unknown error'}
                onCloseCb={() => props.cleanUpdateCredentialsErrorAction()}
            />
        </>
    );
}

export const mapDispatchToProps: DispatchToPropsFn<CredentialsActions> = () => dispatch => ({
    updateCredentialsAction: bindActionCreators(updateCredentials, dispatch),
    cleanUpdateCredentialsErrorAction: bindActionCreators(cleanUpdateCredentialsResult, dispatch),
});

export const mapStateToProps: StateToPropsFn<CredentialsState> = () => state => ({
    accountInfoResult: state.accountInfoResult,
});
