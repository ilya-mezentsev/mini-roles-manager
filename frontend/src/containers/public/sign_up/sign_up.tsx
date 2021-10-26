import { useState, useEffect } from 'react';
import {
    Button, Container,
    TextField,
} from '@material-ui/core';
import { observer } from 'mobx-react-lite';

import { Alert } from '../../../components/shared/';
import { registrationStore } from '../../../store';

export const SignUp = observer(() => {
    const [login, setLogin] = useState('');
    const [password, setPassword] = useState('');

    // eslint-disable-next-line
    useEffect(() => () => registrationStore.cleanRegistrationAction(), []);

    const onCLoseAlert = () => {
        registrationStore.cleanRegistrationAction();
    };

    const alertMessage = () => {
        if (registrationStore.registrationError) {
            return registrationStore.registrationError.description || 'Unknown error';
        } else {
            return 'Registration performed successfully';
        }
    };

    const trySignUp = () => {
        registrationStore.cleanRegistrationAction();
        return registrationStore.signUp({ login, password });
    };

    return (
        <Container maxWidth="sm">
            <h1>Sign-Up</h1>

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
                type="password"
                margin="normal"
                value={password}
                onChange={e => setPassword((e.target as HTMLInputElement).value)}
            />

            <Button
                variant="contained"
                color="primary"
                onClick={() => trySignUp()}
            >
                Sign-Up
            </Button>

            <Alert
                shouldShow={!!registrationStore.registrationError || registrationStore.registeredOk}
                severity={!registrationStore.registrationError ? 'success' : 'error'}
                message={alertMessage()}
                onCloseCb={() => onCLoseAlert()}
            />
        </Container>
    )
});

