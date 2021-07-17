import { Action } from 'redux';

import * as log from '../../services/log';

import { SessionActionResult, SessionResult } from './session.types';
import { ACTIONS } from './action_types';
import { UnknownErrorCode, UnknownErrorDescription } from '../shared/const';

export function sessionReducer(
    state = {},
    action: Action & { userSession?: SessionActionResult },
): SessionResult {
    switch (action.type) {
        case ACTIONS.SUCCESS_LOGIN:
        case ACTIONS.SUCCESS_SIGN_IN:
        case ACTIONS.FAILED_SIGN_IN:
        case ACTIONS.FAILED_SIGN_OUT:
            return action.userSession as SessionResult;

        case ACTIONS.SUCCESS_SIGN_OUT:
            return {
                ...state,
                session: {
                    id: '',
                }
            };

        case ACTIONS.FAILED_LOGIN:
        case ACTIONS.FAILED_TO_PERFORM_LOGIN_ACTION:
            log.error(`Failed to login: ${action.userSession?.error}`);
            return state;

        case ACTIONS.FAILED_TO_PERFORM_SIGN_OUT_ACTION:
        case ACTIONS.FAILED_TO_PERFORM_SIGN_IN_ACTION:
            log.error((action.userSession?.error as unknown as Error)?.toString() || 'Unknown error');
            return {
                ...state,
                error: {
                    code: UnknownErrorCode,
                    description: UnknownErrorDescription,
                },
            };

        case ACTIONS.CLEAN_SIGN_OUT_ERROR:
        case ACTIONS.CLEAN_SIGN_IN_ERROR:
            return {
                ...state,
                error: null,
            };

        default:
            return state;
    }
}
