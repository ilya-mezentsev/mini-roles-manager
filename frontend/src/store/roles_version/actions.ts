import { Dispatch } from 'redux';

import { RolesVersion } from '../../services/api';
import {
    createRolesVersion as createRolesVersionAPI,
    rolesVersionsList,
    updateRolesVersion as updateRolesVersionAPI,
    deleteRolesVersion as deleteRolesVersionAPI,
} from '../../services/api';
import { ACTIONS } from './action_types';

export function createRolesVersion(rolesVersion: RolesVersion): (dispatch: Dispatch) => Promise<void> {
    return async dispatch => {
        try {
            const response = await createRolesVersionAPI(rolesVersion);

            if (response.isOk()) {
                dispatch({
                    type: ACTIONS.SUCCESS_CREATE_ROLES_VERSION,
                    rolesVersionResult: { rolesVersion },
                });
            } else {
                dispatch({
                    type: ACTIONS.FAILED_CREATE_ROLES_VERSION,
                    rolesVersionResult: {
                        error: response.data(),
                    },
                });
            }
        } catch (e) {
            dispatch({
                type: ACTIONS.FAILED_TO_PERFORM_ROLES_VERSION_CREATION,
                rolesVersionResult: {
                    error: e,
                },
            });
        }
    };
}

export function cleanCreateRolesVersionError(): (dispatch: Dispatch) => void {
    return dispatch => {
        dispatch({
            type: ACTIONS.CLEAN_CREATE_ROLES_VERSION_ERROR,
        });
    };
}

export function fetchRolesVersion(): (dispatch: Dispatch) => Promise<void> {
    return async dispatch => {
        try {
            const response = await rolesVersionsList();

            if (response.isOk()) {
                dispatch({
                    type: ACTIONS.SUCCESS_FETCH_ROLES_VERSION,
                    rolesVersionResult: {
                        list: response.data(),
                    },
                });
            } else {
                dispatch({
                    type: ACTIONS.FAILED_FETCH_ROLES_VERSION,
                    rolesVersionResult: {
                        error: response.data(),
                    },
                });
            }
        } catch (e) {
            dispatch({
                type: ACTIONS.FAILED_TO_PERFORM_ROLES_VERSION_FETCHING,
                rolesVersionResult: {
                    error: e,
                },
            });
        }
    };
}

export function cleanFetchRolesVersionError(): (dispatch: Dispatch) => void {
    return dispatch => {
        dispatch({
            type: ACTIONS.CLEAN_FETCH_ROLES_VERSION_ERROR,
        });
    };
}

export function updateRolesVersion(rolesVersion: RolesVersion): (dispatch: Dispatch) => Promise<void> {
    return async dispatch => {
        try {
            const response = await updateRolesVersionAPI(rolesVersion);

            if (response.isOk()) {
                dispatch({
                    type: ACTIONS.SUCCESS_UPDATE_ROLES_VERSION,
                    rolesVersionResult: { rolesVersion },
                });
            } else {
                dispatch({
                    type: ACTIONS.FAILED_UPDATE_ROLES_VERSION,
                    rolesVersionResult: {
                        error: response.data(),
                    },
                });
            }
        } catch (e) {
            dispatch({
                type: ACTIONS.FAILED_TO_PERFORM_ROLES_VERSION_UPDATING,
                rolesVersionResult: {
                    error: e,
                },
            });
        }
    };
}

export function cleanUpdateRolesVersionError(): (dispatch: Dispatch) => void {
    return dispatch => {
        dispatch({
            type: ACTIONS.CLEAN_UPDATE_ROLES_VERSION_ERROR,
        });
    };
}

export function deleteRolesVersion(rolesVersionId: string): (dispatch: Dispatch) => Promise<void> {
    return async dispatch => {
        try {
            const response = await deleteRolesVersionAPI(rolesVersionId);

            if (response.isOk()) {
                dispatch({
                    type: ACTIONS.SUCCESS_DELETE_ROLES_VERSION,
                    rolesVersionResult: { rolesVersionId },
                });
            } else {
                dispatch({
                    type: ACTIONS.FAILED_DELETE_ROLES_VERSION,
                    rolesVersionResult: {
                        error: response.data(),
                    },
                });
            }
        } catch (e) {
            dispatch({
                type: ACTIONS.FAILED_TO_PERFORM_ROLES_VERSION_DELETING,
                rolesVersionResult: {
                    error: e,
                },
            });
        }
    };
}

export function cleanDeleteRolesVersionError(): (dispatch: Dispatch) => void {
    return dispatch => {
        dispatch({
            type: ACTIONS.CLEAN_DELETE_ROLES_VERSION_ERROR,
        });
    };
}

export function selectCurrentRolesVersion(rolesVersion: RolesVersion): (dispatch: Dispatch) => void {
    return dispatch => {
        dispatch({
            type: ACTIONS.SELECT_CURRENT_ROLES_VERSION,
            rolesVersionResult: { rolesVersion },
        });
    };
}
