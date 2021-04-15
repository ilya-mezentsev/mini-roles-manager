import { Dispatch } from 'redux';

import { ACTIONS } from './action_types';
import {
    AccountCredentials,
    fetchInfo as fetchInfoAPI,
    updateCredentials as updateCredentialsAPI,
} from '../../services/api';

export function fetchInfo(): (dispatch: Dispatch) => Promise<void> {
    return async dispatch => {
        try {
            const response = await fetchInfoAPI();

            if (response.isOk()) {
                dispatch({
                    type: ACTIONS.SUCCESS_FETCH_INFO,
                    accountInfoResult: {
                        info: response.data(),
                    },
                });
            } else {
                dispatch({
                    type: ACTIONS.FAILED_FETCH_INFO,
                    accountInfoResult: {
                        fetchInfoError: response.data(),
                    },
                });
            }
        } catch (e) {
            dispatch({
                type: ACTIONS.FAILED_TO_PERFORM_INFO_FETCHING,
                accountInfoResult: {
                    fetchInfoError: e,
                },
            });
        }
    };
}

export function cleanFetchInfoError(): (dispatch: Dispatch) => void {
    return dispatch => {
        dispatch({
            type: ACTIONS.CLEAN_FETCH_INFO_ERROR,
        });
    };
}

export function updateCredentials(credentials: AccountCredentials): (dispatch: Dispatch) => Promise<void> {
    return async dispatch => {
        try {
            const response = await updateCredentialsAPI(credentials);

            if (response.isOk()) {
                dispatch({
                    type: ACTIONS.SUCCESS_UPDATE_CREDENTIALS,
                    accountInfoResult: { credentials },
                });
            } else {
                dispatch({
                    type: ACTIONS.FAILED_UPDATE_CREDENTIALS,
                    accountInfoResult: {
                        updateCredentialsError: response.data(),
                    },
                });
            }
        } catch (e) {
            dispatch({
                type: ACTIONS.FAILED_TO_PERFORM_CREDENTIALS_UPDATING,
                accountInfoResult: {
                    updateCredentialsError: e,
                },
            });
        }
    };
}

export function cleanUpdateCredentialsResult(): (dispatch: Dispatch) => void {
    return dispatch => {
        dispatch({
            type: ACTIONS.CLEAN_UPDATE_CREDENTIALS_RESULT,
        });
    };
}
