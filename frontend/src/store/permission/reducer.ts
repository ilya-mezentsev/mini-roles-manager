import { Action } from 'redux';

import {
    FetchPermissionResult,
    FetchPermissionActionResult,
} from './permissions.types';
import { ACTIONS } from './action_types';
import { UnknownErrorCode, UnknownErrorDescription } from '../shared/const';
import { APIError } from '../../services/api/shared';
import * as log from '../../services/log';

export function fetchPermissionReducer(
    state = null,
    action: Action<ACTIONS> & { fetchPermissionResult?: FetchPermissionActionResult },
): FetchPermissionResult | null {
    const currentState = state || {};

    switch (action.type) {
        case ACTIONS.SUCCESS_FETCH_PERMISSION:
            return {
                ...currentState,
                effect: action.fetchPermissionResult!.effect,
            };

        case ACTIONS.FAILED_FETCH_PERMISSION:
            return {
                ...currentState,
                error: action.fetchPermissionResult!.error as APIError,
            };

        case ACTIONS.FAILED_TO_PERFORM_PERMISSION_FETCHING:
            log.error(`Failed to fetch permission: ${(action.fetchPermissionResult!.error!.toString())}`);

            return {
                ...currentState,
                error: {
                    code: UnknownErrorCode,
                    description: UnknownErrorDescription,
                },
            };

        case ACTIONS.CLEAN_FETCH_PERMISSION_RESULT:
            return {
                ...currentState,
                effect: null,
                error: null,
            };

        default:
            return state;
    }
}
