import {
    TextField,
    Box,
    Button,
} from '@material-ui/core';
import { bindActionCreators } from 'redux';

import { SignInActions, SignInProps, SignInState } from './sign_in.types';
import { cleanSignIn, signIn } from '../../../store/session/actions';
import { Alert } from '../../../components/shared';
import { EventEmitter } from 'events';
import { DispatchToPropsFn, StateToPropsFn } from '../../../shared/types';
import { APIError } from '../../../services/api/shared';

export const SignIn = (props: SignInProps) => {
    let login: string, password: string;
    const setOpenEventName = 'set:open';
    const e = new EventEmitter();

    const trySignIn = () => {
        props.cleanSignInAction();
        props.signInAction({ login, password });
        e.emit(setOpenEventName);
    };

    return (
        <Box>
            <h1>Sign-In</h1>

            <TextField
                label="Login"
                required
                fullWidth={true}
                margin="normal"
                onInput={e => login = (e.target as HTMLInputElement).value}
            />
            <TextField
                label="Password"
                required
                fullWidth={true}
                margin="normal"
                type="password"
                onInput={e => password = (e.target as HTMLInputElement).value}
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
                setOpenEmitter={e}
                setOpenEventName={setOpenEventName}
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
