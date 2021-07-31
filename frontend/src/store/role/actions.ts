import { Dispatch } from 'redux';

import { Role } from '../../services/api';
import {
    createRole as createRoleAPI,
    rolesList,
    updateRole as updateRoleAPI,
    deleteRole as deleteRoleAPI,
} from '../../services/api';
import { ACTIONS } from './action_types';

export function createRole(role: Role): (dispatch: Dispatch) => Promise<void> {
    return async dispatch => {
        try {
            const response = await createRoleAPI(role);

            if (response.isOk()) {
                dispatch({
                    type: ACTIONS.SUCCESS_CREATE_ROLE,
                    rolesResult: { role },
                });
            } else {
                dispatch({
                    type: ACTIONS.FAILED_CREATE_ROLE,
                    rolesResult: {
                        error: response.data(),
                    },
                });
            }
        } catch (e) {
            dispatch({
                type: ACTIONS.FAILED_TO_PERFORM_ROLE_CREATION,
                rolesResult: {
                    error: e,
                },
            });
        }
    };
}

export function cleanCreateRoleError(): (dispatch: Dispatch) => void {
    return dispatch => {
        dispatch({
            type: ACTIONS.CLEAN_CREATE_ROLE_ERROR,
        });
    };
}

export function fetchRoles(): (dispatch: Dispatch) => Promise<void> {
    return async dispatch => {
        try {
            const response = await rolesList();

            if (response.isOk()) {
                dispatch({
                    type: ACTIONS.SUCCESS_FETCH_ROLES,
                    rolesResult: {
                        list: response.data(),
                    },
                });
            } else {
                dispatch({
                    type: ACTIONS.FAILED_FETCH_ROLES,
                    rolesResult: {
                        error: response.data(),
                    },
                });
            }
        } catch (e) {
            dispatch({
                type: ACTIONS.FAILED_TO_PERFORM_ROLES_FETCHING,
                rolesResult: {
                    error: e,
                },
            });
        }
    };
}

export function cleanFetchRolesError(): (dispatch: Dispatch) => void {
    return dispatch => {
        dispatch({
            type: ACTIONS.CLEAN_FETCH_ROLES_ERROR,
        });
    };
}

export function updateRole(role: Role): (dispatch: Dispatch) => Promise<void> {
    return async dispatch => {
        try {
            const response = await updateRoleAPI(role);

            if (response.isOk()) {
                dispatch({
                    type: ACTIONS.SUCCESS_UPDATE_ROLE,
                    rolesResult: { role },
                });
            } else {
                dispatch({
                    type: ACTIONS.FAILED_UPDATE_ROLE,
                    rolesResult: {
                        error: response.data(),
                    },
                });
            }
        } catch (e) {
            dispatch({
                type: ACTIONS.FAILED_TO_PERFORM_ROLE_UPDATING,
                rolesResult: {
                    error: e,
                },
            });
        }
    };
}

export function cleanUpdateRoleError(): (dispatch: Dispatch) => void {
    return dispatch => {
        dispatch({
            type: ACTIONS.CLEAN_UPDATE_ROLE_ERROR,
        });
    };
}

export function deleteRole(rolesVersionId: string, roleId: string): (dispatch: Dispatch) => Promise<void> {
    return async dispatch => {
        try {
            const response = await deleteRoleAPI(rolesVersionId, roleId);

            if (response.isOk()) {
                dispatch({
                    type: ACTIONS.SUCCESS_DELETE_ROLE,
                    rolesResult: { roleId, rolesVersionId },
                });
            } else {
                dispatch({
                    type: ACTIONS.FAILED_DELETE_ROLE,
                    rolesResult: {
                        error: response.data(),
                    },
                });
            }
        } catch (e) {
            dispatch({
                type: ACTIONS.FAILED_TO_PERFORM_ROLE_DELETING,
                rolesResult: {
                    error: e,
                },
            });
        }
    };
}

export function cleanDeleteRoleError(): (dispatch: Dispatch) => void {
    return dispatch => {
        dispatch({
            type: ACTIONS.CLEAN_DELETE_ROLE_ERROR,
        });
    };
}
