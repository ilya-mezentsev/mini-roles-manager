import * as log from '../../services/log';
import { RegistrationActionResult, RegistrationResult } from './registration.types';
import { ACTIONS } from './action_types';
import { Action } from 'redux';
import { UnknownErrorCode, UnknownErrorDescription } from '../shared/const';

export function registrationReducer(
    state = null,
    action: Action & { registrationResult?: RegistrationActionResult },
): RegistrationResult | null {
    switch (action.type) {
        case ACTIONS.SUCCESS_REGISTRATION:
        case ACTIONS.FAILED_TO_REGISTER_USER:
            return action.registrationResult as RegistrationResult;

        case ACTIONS.FAILED_TO_PERFORM_REGISTER_ACTION:
            log.error((action.registrationResult?.error as unknown as Error)?.toString() || 'Unknown error');
            return {
                error: {
                    code: UnknownErrorCode,
                    description: UnknownErrorDescription,
                },
            };

        case ACTIONS.CLEAN_REGISTRATION:
            return null;

        default:
            return state;
    }
}
