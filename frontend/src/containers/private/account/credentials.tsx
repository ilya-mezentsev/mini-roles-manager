import { useState, useEffect } from 'react';
import { observer } from 'mobx-react-lite';
import {
    Button,
    TextField,
} from '@material-ui/core';

import { Alert } from '../../../components/shared';
import { accountInfoStore } from '../../../store';

export const Credentials = observer(() => {
    const [login, setLogin] = useState(accountInfoStore.info?.login || '');
    const [password, setPassword] = useState('');

    // useEffect(() => {
    //     setLogin(props.accountInfoResult?.info?.login || '');
    // }, [props.accountInfoResult?.info?.login]);

    // eslint-disable-next-line
    useEffect(() => () => accountInfoStore.cleanAccountInfoActionErrors(), []);

    const updateCredentials = () => {
        return accountInfoStore.updateCredentials({
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
                disabled={login === accountInfoStore.info?.login && !password}
                variant="contained"
                color="primary"
                onClick={updateCredentials}
            >
                Save
            </Button>

            <Alert
                shouldShow={!!accountInfoStore.credentials}
                severity="success"
                message="Credentials are updated successfully"
                onCloseCb={() => {
                    accountInfoStore.cleanAccountInfoActionErrors();
                    setPassword('');
                }}
            />

            <Alert
                shouldShow={!!accountInfoStore.updateCredentialsError}
                severity="error"
                message={accountInfoStore.updateCredentialsError?.description || 'Unknown error'}
                onCloseCb={() => accountInfoStore.cleanAccountInfoActionErrors()}
            />
        </>
    );
});
