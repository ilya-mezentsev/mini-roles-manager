import { useState, useEffect } from 'react';
import { useHistory } from 'react-router-dom';
import {
    Button, Container,
    TextField,
} from '@material-ui/core';
import { bindActionCreators } from 'redux';
import { signUp, cleanSignUp } from '../../../store/registration/actions';
import { SignUpActions, SignUpProps, SignUpState } from './sign_up.types';
import { Alert } from '../../../components/shared/';
import { DispatchToPropsFn, StateToPropsFn } from '../../../shared/types';

export const SignUp = (props: SignUpProps) => {
    const [login, setLogin] = useState('');
    const [password, setPassword] = useState('');
    const history = useHistory();

    useEffect(() => () => props.cleanSignUpAction(), []);

    const onCLoseAlert = () => {
        props.cleanSignUpAction();

        if (!props.registrationResult?.error) {
            history.push('/sign-in');
        }
    };

    const alertMessage = () => {
        if (props.registrationResult?.error) {
            return props.registrationResult.error.description || 'Unknown error';
        } else {
            return 'Registration performed successfully';
        }
    };

    const trySignUp = () => {
        props.cleanSignUpAction();
        props.signUpAction({ login, password });
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
                shouldShow={!!props.registrationResult}
                severity={!props.registrationResult?.error ? 'success' : 'error'}
                message={alertMessage()}
                onCloseCb={() => onCLoseAlert()}
            />
        </Container>
    )
}

export const mapDispatchToProps: DispatchToPropsFn<SignUpActions> = () => dispatch => ({
    signUpAction: bindActionCreators(signUp, dispatch),
    cleanSignUpAction: bindActionCreators(cleanSignUp, dispatch),
});

export const mapStateToProps: StateToPropsFn<SignUpState> = () => state => ({
    registrationResult: state.registrationResult,
});
