import { Action } from 'redux';

import * as log from '../../services/log';
import {
    AccountInfoActionResult,
    AccountInfoResult,
    FetchAccountInfoActionErrorResult,
    FetchAccountInfoActionResult,
    UpdateCredentialsActionResult,
    UpdateCredentialsActionErrorResult,
} from './account_info.types';
import { ACTIONS } from './action_types';
import { APIError } from '../../services/api/shared';
import { UnknownErrorCode, UnknownErrorDescription } from '../shared/const';

export const actionToErrorKey = {
    [ACTIONS.CLEAN_FETCH_INFO_ERROR]: 'fetchInfoError',
    [ACTIONS.FAILED_TO_PERFORM_INFO_FETCHING]: 'fetchInfoError',

    [ACTIONS.CLEAN_UPDATE_CREDENTIALS_RESULT]: 'updateCredentialsError',
    [ACTIONS.FAILED_TO_PERFORM_CREDENTIALS_UPDATING]: 'updateCredentialsError',
};

export function accountInfoReducer(
    state = null,
    action: Action<ACTIONS> & { accountInfoResult?: AccountInfoActionResult }
): AccountInfoResult | null {
    const currentState = state || {
        info: {
            login: '',
            apiKey: '',
            created: '',
        }
    };

    switch (action.type) {
        case ACTIONS.SUCCESS_FETCH_INFO:
            return {
                ...currentState,
                info: (action.accountInfoResult as FetchAccountInfoActionResult).info,
            };

        case ACTIONS.FAILED_FETCH_INFO:
            return {
                ...currentState,
                fetchInfoError: (action.accountInfoResult as FetchAccountInfoActionErrorResult).fetchInfoError as APIError,
            };

        case ACTIONS.CLEAN_FETCH_INFO_ERROR:
            return {
                ...currentState,
                [actionToErrorKey[action.type]]: null,
            };

        case ACTIONS.SUCCESS_UPDATE_CREDENTIALS:
            return {
                ...currentState,
                info: {
                    ...currentState.info,
                    login: (action.accountInfoResult as UpdateCredentialsActionResult).credentials!.login,
                },
                credentials: (action.accountInfoResult as UpdateCredentialsActionResult).credentials,
            };

        case ACTIONS.FAILED_UPDATE_CREDENTIALS:
            return {
                ...currentState,
                updateCredentialsError: (action.accountInfoResult as UpdateCredentialsActionErrorResult).updateCredentialsError as APIError,
            };

        case ACTIONS.CLEAN_UPDATE_CREDENTIALS_RESULT:
            return {
                ...currentState,
                credentials: null,
                [actionToErrorKey[action.type]]: null,
            };

        case ACTIONS.FAILED_TO_PERFORM_CREDENTIALS_UPDATING:
        case ACTIONS.FAILED_TO_PERFORM_INFO_FETCHING:
            log.error(`Unable to perform action <${action.type}> - ${JSON.stringify(action.accountInfoResult)}`);

            return {
                ...currentState,
                [actionToErrorKey[action.type]]: {
                    code: UnknownErrorCode,
                    description: UnknownErrorDescription,
                },
            };

        default:
            return currentState;
    }
}
