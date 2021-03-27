import { useHistory } from 'react-router-dom';
import {
    Box,
    Button,
    TextField,
} from '@material-ui/core';
import { bindActionCreators } from 'redux';
import { signUp, cleanSignUp } from '../../../store/registration/actions';
import { SignUpActions, SignUpProps, SignUpState } from './sign_up.types';
import { Alert } from '../../../components/shared/';
import { EventEmitter } from 'events';
import { DispatchToPropsFn, StateToPropsFn } from '../../../shared/types';
import { APIError } from '../../../services/api/shared';

export const SignUp = (props: SignUpProps) => {
    let login = '';
    let password = '';
    const setOpenEventName = 'set:open';
    const e = new EventEmitter();
    const history = useHistory();

    const onCLoseAlert = () => {
        if (!props.registrationResult?.error) {
            history.push('/sign-in');
        }

        props.cleanSignUpAction();
    };

    const alertMessage = () => {
        if (props.registrationResult?.error) {
            return (props.registrationResult.error as APIError).description || 'Unknown error';
        } else {
            return 'Registration performed successfully';
        }
    };

    const trySignUp = () => {
        props.cleanSignUpAction();
        props.signUpAction({ login, password });
        e.emit(setOpenEventName);
    };

    return (
        <Box>
            <h1>Sign-Up</h1>

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
                type="password"
                margin="normal"
                onInput={e => password = (e.target as HTMLInputElement).value}
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
                setOpenEmitter={e}
                setOpenEventName={setOpenEventName}
            />
        </Box>
    )
}

export const mapDispatchToProps: DispatchToPropsFn<SignUpActions> = () => dispatch => ({
    signUpAction: bindActionCreators(signUp, dispatch),
    cleanSignUpAction: bindActionCreators(cleanSignUp, dispatch),
});

export const mapStateToProps: StateToPropsFn<SignUpState> = () => state => ({
    registrationResult: state.registrationResult,
});
