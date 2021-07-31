import { Action } from 'redux';
import * as log from '../../services/log';

import { ACTIONS } from './action_types';
import {
    RolesVersionActionResult,
    RolesVersionErrorResult,
    RolesVersionIdActionResult,
    RolesVersionResult,
    RolesVersionsActionResult,
    RolesVersionsListResult,
} from './roles_version.types';
import { APIError } from '../../services/api/shared';
import {
    UnknownErrorCode,
    UnknownErrorDescription,
} from '../shared/const';

const actionToLogMessage = {
    [ACTIONS.FAILED_TO_PERFORM_ROLES_VERSION_CREATION]: 'Failed to create role',
    [ACTIONS.FAILED_TO_PERFORM_ROLES_VERSION_FETCHING]: 'Failed to fetch roles',
    [ACTIONS.FAILED_TO_PERFORM_ROLES_VERSION_UPDATING]: 'Failed to update role',
    [ACTIONS.FAILED_TO_PERFORM_ROLES_VERSION_DELETING]: 'Failed to delete role',
};
export const actionToErrorType: {[key: string]: string} = {
    [ACTIONS.FAILED_CREATE_ROLES_VERSION]: 'createError',
    [ACTIONS.FAILED_TO_PERFORM_ROLES_VERSION_CREATION]: 'createError',
    [ACTIONS.CLEAN_CREATE_ROLES_VERSION_ERROR]: 'createError',

    [ACTIONS.FAILED_FETCH_ROLES_VERSION]: 'fetchError',
    [ACTIONS.FAILED_TO_PERFORM_ROLES_VERSION_FETCHING]: 'fetchError',
    [ACTIONS.CLEAN_FETCH_ROLES_VERSION_ERROR]: 'fetchError',

    [ACTIONS.FAILED_UPDATE_ROLES_VERSION]: 'updateError',
    [ACTIONS.FAILED_TO_PERFORM_ROLES_VERSION_UPDATING]: 'updateError',
    [ACTIONS.CLEAN_UPDATE_ROLES_VERSION_ERROR]: 'updateError',

    [ACTIONS.FAILED_DELETE_ROLES_VERSION]: 'deleteError',
    [ACTIONS.FAILED_TO_PERFORM_ROLES_VERSION_DELETING]: 'deleteError',
    [ACTIONS.CLEAN_DELETE_ROLES_VERSION_ERROR]: 'deleteError',
};

export function rolesVersionReducer(
    state: RolesVersionResult = { list: null, currentRolesVersion: null },
    action: Action<string> & { rolesVersionResult?: RolesVersionsActionResult },
): RolesVersionResult {
    switch (action.type) {
        case ACTIONS.SUCCESS_CREATE_ROLES_VERSION:
            return {
                ...state,
                list: (state.list || []).concat((action.rolesVersionResult as RolesVersionActionResult).rolesVersion),
            };

        case ACTIONS.SUCCESS_FETCH_ROLES_VERSION:
            const list = (action.rolesVersionResult as RolesVersionsListResult).list || [];

            return {
                currentRolesVersion: list[0] || null,
                list,
            };

        case ACTIONS.SUCCESS_UPDATE_ROLES_VERSION:
            const newListAfterUpdating = (state.list || []).slice();
            const updatedRolesVersion = (action.rolesVersionResult as RolesVersionActionResult).rolesVersion;
            const updatedRolesVersionIndex = newListAfterUpdating.findIndex(rv => rv.id === updatedRolesVersion.id);
            if (updatedRolesVersionIndex >= 0) {
                newListAfterUpdating[updatedRolesVersionIndex] = {
                    ...newListAfterUpdating[updatedRolesVersionIndex],
                    title: updatedRolesVersion.title,
                };
            }

            return {
                ...state,
                list: newListAfterUpdating,
            };

        case ACTIONS.SUCCESS_DELETE_ROLES_VERSION:
            const deletedRolesVersionId = (action.rolesVersionResult as RolesVersionIdActionResult).rolesVersionId;
            const newListAfterDeleting = (state.list || []).filter(rv => rv.id !== deletedRolesVersionId);

            return {
                ...state,
                currentRolesVersion: state.currentRolesVersion?.id === deletedRolesVersionId
                    ? newListAfterDeleting[0]
                    : state.currentRolesVersion,
                list: newListAfterDeleting,
            };

        case ACTIONS.SELECT_CURRENT_ROLES_VERSION:
            return {
                ...state,
                currentRolesVersion: (action.rolesVersionResult as RolesVersionActionResult).rolesVersion,
            };

        case ACTIONS.FAILED_CREATE_ROLES_VERSION:
        case ACTIONS.FAILED_FETCH_ROLES_VERSION:
        case ACTIONS.FAILED_UPDATE_ROLES_VERSION:
        case ACTIONS.FAILED_DELETE_ROLES_VERSION:
            return {
                ...state,
                [actionToErrorType[action.type]]: (action.rolesVersionResult as RolesVersionErrorResult).error as APIError,
            };

        case ACTIONS.FAILED_TO_PERFORM_ROLES_VERSION_CREATION:
        case ACTIONS.FAILED_TO_PERFORM_ROLES_VERSION_FETCHING:
        case ACTIONS.FAILED_TO_PERFORM_ROLES_VERSION_UPDATING:
        case ACTIONS.FAILED_TO_PERFORM_ROLES_VERSION_DELETING:
            log.error(
                `${actionToLogMessage[action.type]}: ${(action.rolesVersionResult as RolesVersionErrorResult).error.toString()}`
            );
            return {
                ...state,
                [actionToErrorType[action.type]]: {
                    code: UnknownErrorCode,
                    description: UnknownErrorDescription,
                },
            };

        case ACTIONS.CLEAN_CREATE_ROLES_VERSION_ERROR:
        case ACTIONS.CLEAN_FETCH_ROLES_VERSION_ERROR:
        case ACTIONS.CLEAN_UPDATE_ROLES_VERSION_ERROR:
        case ACTIONS.CLEAN_DELETE_ROLES_VERSION_ERROR:
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
