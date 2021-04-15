import { Dispatch } from 'redux';

import { fetchPermission as fetchPermissionAPI } from '../../services/api';
import { PermissionAccessRequest } from '../../services/api';
import { ACTIONS } from './action_types';

export function fetchPermission(request: PermissionAccessRequest): (dispatch: Dispatch) => Promise<void> {
    return async dispatch => {
        try {
            const response = await fetchPermissionAPI(request);

            if (response.isOk()) {
                dispatch({
                    type: ACTIONS.SUCCESS_FETCH_PERMISSION,
                    fetchPermissionResult: response.data(),
                });
            } else {
                dispatch({
                    type: ACTIONS.FAILED_FETCH_PERMISSION,
                    fetchPermissionResult: {
                        error: response.data(),
                    },
                });
            }
        } catch (e) {
            dispatch({
                type: ACTIONS.FAILED_TO_PERFORM_PERMISSION_FETCHING,
                fetchPermissionResult: {
                    error: e,
                },
            });
        }
    };
}

export function cleanFetchPermissionResult(): (dispatch: Dispatch) => void {
    return dispatch => {
        dispatch({
            type: ACTIONS.CLEAN_FETCH_PERMISSION_RESULT,
        });
    };
}
