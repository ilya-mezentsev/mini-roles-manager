import { useState, useEffect } from 'react';
import {
    TextField,
    Button, Container,
} from '@material-ui/core';
import { observer } from 'mobx-react-lite';

import { Alert } from '../../../components/shared';
import { sessionStore } from '../../../store';

export const SignIn = observer(() => {
    const [login, setLogin] = useState('');
    const [password, setPassword] = useState('');

    // eslint-disable-next-line
    useEffect(() => () => sessionStore.cleanSessionActionErrors(), [])

    const trySignIn = () => {
        sessionStore.cleanSessionActionErrors();
        return sessionStore.signIn({ login, password });
    };

    return (
        <Container maxWidth="sm">
            <h1>Sign-In</h1>

            <TextField
                label="Login"
                required
                fullWidth={true}
                margin="normal"
                value={login}
                onChange={e => setLogin((e.target as HTMLInputElement).value)}
            />
            <TextField
                label="Password"
                required
                fullWidth={true}
                margin="normal"
                type="password"
                value={password}
                onChange={e => setPassword((e.target as HTMLInputElement).value)}
            />

            <Button
                variant="contained"
                color="primary"
                onClick={() => trySignIn()}
            >
                Sign-In
            </Button>

            <Alert
                shouldShow={!!sessionStore.signInError}
                severity={'error'}
                message={sessionStore.signInError?.description || 'Unknown error'}
                onCloseCb={() => sessionStore.cleanSessionActionErrors()}
            />
        </Container>
    );
});
