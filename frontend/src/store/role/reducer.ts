import { Action } from 'redux';
import * as log from '../../services/log';

import { ACTIONS } from './action_types';
import { APIError } from '../../services/api/shared';
import { Role } from '../../services/api';
import { UnknownErrorCode, UnknownErrorDescription } from '../shared/const';
import {
    RoleActionResult,
    RoleErrorResult,
    RoleIdActionResult,
    RolesActionResult,
    RolesListResult,
    RolesResult,
} from './role.types';

const actionToLogMessage = {
    [ACTIONS.FAILED_TO_PERFORM_ROLE_CREATION]: 'Failed to create role',
    [ACTIONS.FAILED_TO_PERFORM_ROLES_FETCHING]: 'Failed to fetch roles',
    [ACTIONS.FAILED_TO_PERFORM_ROLE_UPDATING]: 'Failed to update role',
    [ACTIONS.FAILED_TO_PERFORM_ROLE_DELETING]: 'Failed to delete role',
};
export const actionToErrorType: {[key: string]: string} = {
    [ACTIONS.FAILED_CREATE_ROLE]: 'createError',
    [ACTIONS.FAILED_TO_PERFORM_ROLE_CREATION]: 'createError',
    [ACTIONS.CLEAN_CREATE_ROLE_ERROR]: 'createError',

    [ACTIONS.FAILED_FETCH_ROLES]: 'fetchError',
    [ACTIONS.FAILED_TO_PERFORM_ROLES_FETCHING]: 'fetchError',
    [ACTIONS.CLEAN_FETCH_ROLES_ERROR]: 'fetchError',

    [ACTIONS.FAILED_UPDATE_ROLE]: 'updateError',
    [ACTIONS.FAILED_TO_PERFORM_ROLE_UPDATING]: 'updateError',
    [ACTIONS.CLEAN_UPDATE_ROLE_ERROR]: 'updateError',

    [ACTIONS.FAILED_DELETE_ROLE]: 'deleteError',
    [ACTIONS.FAILED_TO_PERFORM_ROLE_DELETING]: 'deleteError',
    [ACTIONS.CLEAN_DELETE_ROLE_ERROR]: 'deleteError',
};

export function roleReducer(
    state: RolesResult = { list: null },
    action: Action<string> & { rolesResult?: RolesActionResult }
): RolesResult {
    switch (action.type) {
        case ACTIONS.SUCCESS_CREATE_ROLE:
            return {
                ...state,
                list: (state.list || []).concat((action.rolesResult as RoleActionResult).role as Role),
            };

        case ACTIONS.SUCCESS_FETCH_ROLES:
            return {
                list: (action.rolesResult as RolesListResult).list || [],
            };

        case ACTIONS.SUCCESS_UPDATE_ROLE:
            const newList = (state.list || []).slice();
            const updatedRole = (action.rolesResult as RoleActionResult).role;
            const updatedRoleIndex = newList.findIndex(r => r.id === updatedRole.id);
            if (updatedRoleIndex >= 0) {
                newList[updatedRoleIndex] = {
                    ...newList[updatedRoleIndex],
                    title: updatedRole.title,
                    permissions: updatedRole.permissions,
                    extends: updatedRole.extends,
                };
            }

            return {
                ...state,
                list: newList,
            };

        case ACTIONS.SUCCESS_DELETE_ROLE:
            const deletedRoleId = (action.rolesResult as RoleIdActionResult).roleId;
            return {
                ...state,
                list: (state.list || [])
                    .filter(r => r.id !== deletedRoleId)
                    .map(r => ({
                        ...r,
                        extends: r.extends?.filter(roleId => roleId !== deletedRoleId)
                    })),
            };

        case ACTIONS.FAILED_CREATE_ROLE:
        case ACTIONS.FAILED_FETCH_ROLES:
        case ACTIONS.FAILED_UPDATE_ROLE:
        case ACTIONS.FAILED_DELETE_ROLE:
            return {
                ...state,
                [actionToErrorType[action.type]]: (action.rolesResult as RoleErrorResult).error as APIError,
            };

        case ACTIONS.FAILED_TO_PERFORM_ROLE_CREATION:
        case ACTIONS.FAILED_TO_PERFORM_ROLES_FETCHING:
        case ACTIONS.FAILED_TO_PERFORM_ROLE_UPDATING:
        case ACTIONS.FAILED_TO_PERFORM_ROLE_DELETING:
            log.error(
                `${actionToLogMessage[action.type]}: ${(action.rolesResult as RoleErrorResult).error.toString()}`
            );
            return {
                ...state,
                [actionToErrorType[action.type]]: {
                    code: UnknownErrorCode,
                    description: UnknownErrorDescription,
                },
            };

        case ACTIONS.CLEAN_CREATE_ROLE_ERROR:
        case ACTIONS.CLEAN_FETCH_ROLES_ERROR:
        case ACTIONS.CLEAN_UPDATE_ROLE_ERROR:
        case ACTIONS.CLEAN_DELETE_ROLE_ERROR:
            return {
                ...state,
                [actionToErrorType[action.type]]: null,
            };

        default:
            return {
                ...state,
            };
    }
}
