import { Dispatch } from 'redux';
import {
    AccountCredentials,
    signIn as signInAPI,
    signOut as signOutAPI,
    login as loginAPI,
} from '../../services/api';
import { ACTIONS } from './action_types';

export function signIn(credentials: AccountCredentials): (dispatch: Dispatch) => Promise<void> {
    return async dispatch => {
        try {
            const response = await signInAPI(credentials);

            if (response.isOk()) {
                dispatch({
                    type: ACTIONS.SUCCESS_SIGN_IN,
                    userSession: {
                        session: response.data(),
                    },
                });
            } else {
                dispatch({
                    type: ACTIONS.FAILED_SIGN_IN,
                    userSession: {
                        error: response.data(),
                    },
                });
            }
        } catch (e) {
            dispatch({
                type: ACTIONS.FAILED_TO_PERFORM_SIGN_IN_ACTION,
                userSession: {
                    error: e,
                },
            });
        }
    };
}

export function cleanSignInError(): (dispatch: Dispatch) => void {
    return dispatch => {
        dispatch({
            type: ACTIONS.CLEAN_SIGN_IN_ERROR,
        });
    };
}

export function signOut(): (dispatch: Dispatch) => Promise<void> {
    return async dispatch => {
        try {
            const response = await signOutAPI();

            if (response.isOk()) {
                dispatch({
                    type: ACTIONS.SUCCESS_SIGN_OUT,
                });
            } else {
                dispatch({
                    type: ACTIONS.FAILED_SIGN_OUT,
                    userSession: {
                        error: response.data(),
                    },
                });
            }
        } catch (e) {
            dispatch({
                type: ACTIONS.FAILED_TO_PERFORM_SIGN_OUT_ACTION,
                userSession: {
                    error: e,
                },
            });
        }
    };
}

export function cleanSignOutError(): (dispatch: Dispatch) => void {
    return dispatch => {
        dispatch({
            type: ACTIONS.CLEAN_SIGN_OUT_ERROR,
        });
    };
}

export function login(): (dispatch: Dispatch) => Promise<void> {
    return async dispatch => {
        try {
            const response = await loginAPI();

            if (response.isOk()) {
                dispatch({
                    type: ACTIONS.SUCCESS_LOGIN,
                    userSession: {
                        session: response.data(),
                    },
                });
            } else {
                dispatch({
                    type: ACTIONS.FAILED_LOGIN,
                    userSession: {
                        error: response.data(),
                    },
                });
            }
        } catch (e) {
            dispatch({
                type: ACTIONS.FAILED_TO_PERFORM_LOGIN_ACTION,
                userSession: {
                    error: e,
                },
            });
        }
    };
}
