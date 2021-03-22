import { AccountCredentials, signUp as signUpAPI } from '../../services/api';
import { ACTIONS } from './action_types';
import { APIError } from '../../services/api/shared';
import { Dispatch } from 'redux';

export function signUp(credentials: AccountCredentials): (dispatch: Dispatch) => Promise<void> {
    return async (dispatch: Dispatch) => {
        try {
            const response = await signUpAPI(credentials);

            if (response.isOk()) {
                dispatch({
                    type: ACTIONS.SUCCESS_REGISTRATION,
                    registrationResult: {
                        failed: false,
                    },
                });
            } else {
                dispatch({
                    type: ACTIONS.FAILED_TO_REGISTER_USER,
                    registrationResult: {
                        failed: true,
                        error: response.data() as APIError,
                    },
                });
            }
        } catch (e) {
            dispatch({
                type: ACTIONS.FAILED_TO_PERFORM_REGISTER_ACTION,
                registrationResult: {
                    failed: true,
                    error: e,
                },
            });
        }
    };
}

export function cleanSignUp(): (dispatch: Dispatch) => void {
    return (dispatch: Dispatch) => {
        dispatch({
            type: ACTIONS.CLEAN_REGISTRATION,
        });
    }
}
