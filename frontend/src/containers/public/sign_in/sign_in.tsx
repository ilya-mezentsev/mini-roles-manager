import { useState, useEffect } from 'react';
import {
    TextField,
    Box,
    Button,
} from '@material-ui/core';
import { bindActionCreators } from 'redux';

import { SignInActions, SignInProps, SignInState } from './sign_in.types';
import { cleanSignIn, signIn } from '../../../store/session/actions';
import { Alert } from '../../../components/shared';
import { DispatchToPropsFn, StateToPropsFn } from '../../../shared/types';
import { APIError } from '../../../services/api/shared';

export const SignIn = (props: SignInProps) => {
    const [login, setLogin] = useState('');
    const [password, setPassword] = useState('');

    useEffect(() => () => props.cleanSignInAction(), [])

    const trySignIn = () => {
        props.cleanSignInAction();
        props.signInAction({ login, password });
    };

    return (
        <Box>
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
                shouldShow={!!props.userSession?.error}
                severity={'error'}
                message={(props.userSession?.error as APIError)?.description || 'Unknown error'}
                onCloseCb={() => props.cleanSignInAction()}
            />
        </Box>
    )
}

export const mapDispatchToProps: DispatchToPropsFn<SignInActions> = () => dispatch => ({
    signInAction: bindActionCreators(signIn, dispatch),
    cleanSignInAction: bindActionCreators(cleanSignIn, dispatch),
});

export const mapStateToProps: StateToPropsFn<SignInState> = () => state => ({
    userSession: state.userSession,
});
