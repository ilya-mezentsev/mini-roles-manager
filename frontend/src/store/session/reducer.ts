import { Action } from 'redux';

import * as log from '../../services/log';

import { SessionResult } from './session.types';
import { ACTIONS } from './action_types';
import { UnknownErrorCode, UnknownErrorDescription } from '../shared/const';

export function sessionReducer(
    state = null,
    action: Action & { userSession: SessionResult },
): SessionResult | null {
    switch (action.type) {
        case ACTIONS.SUCCESS_LOGIN:
        case ACTIONS.SUCCESS_SIGN_IN:
        case ACTIONS.FAILED_SIGN_IN:
            return action.userSession;

        case ACTIONS.FAILED_LOGIN:
        case ACTIONS.FAILED_TO_PERFORM_LOGIN_ACTION:
            log.error(`Failed to login: ${action.userSession.error}`);
            return state;

        case ACTIONS.FAILED_TO_PERFORM_SIGN_IN_ACTION:
            log.error((action.userSession.error as unknown as Error)?.toString() || 'Unknown error');
            return {
                ...(state || {}),
                error: {
                    code: UnknownErrorCode,
                    description: UnknownErrorDescription,
                },
            }

        case ACTIONS.CLEAN_SIGN_IN:
            return {
                ...(state || {}),
                error: null,
            };

        default:
            return state;
    }
}
