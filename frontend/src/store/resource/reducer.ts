import { Action } from 'redux';
import * as log from '../../services/log';
import {
    ResourcesResult,
    ResourcesActionResult,
    ResourceActionResult,
    ResourceIdActionResult,
    ResourceErrorResult,
    ResourcesListResult
} from './resource.types';
import { ACTIONS } from './action_types';
import { APIError } from '../../services/api/shared';
import { Resource } from '../../services/api';
import { UnknownErrorCode, UnknownErrorDescription } from '../shared/const';

const actionToLogMessage = {
    [ACTIONS.FAILED_TO_PERFORM_RESOURCE_CREATION]: 'Failed to create resource',
    [ACTIONS.FAILED_TO_PERFORM_RESOURCES_FETCHING]: 'Failed to fetch resources',
    [ACTIONS.FAILED_TO_PERFORM_RESOURCE_UPDATING]: 'Failed to update resource',
    [ACTIONS.FAILED_TO_PERFORM_RESOURCE_DELETING]: 'Failed to delete resource',
};
export const actionToErrorType: {[key: string]: string} = {
    [ACTIONS.FAILED_CREATE_RESOURCE]: 'createError',
    [ACTIONS.FAILED_TO_PERFORM_RESOURCE_CREATION]: 'createError',
    [ACTIONS.CLEAN_CREATE_RESOURCE_ERROR]: 'createError',

    [ACTIONS.FAILED_FETCH_RESOURCES]: 'fetchError',
    [ACTIONS.FAILED_TO_PERFORM_RESOURCES_FETCHING]: 'fetchError',
    [ACTIONS.CLEAN_FETCH_RESOURCES_ERROR]: 'fetchError',

    [ACTIONS.FAILED_UPDATE_RESOURCE]: 'updateError',
    [ACTIONS.FAILED_TO_PERFORM_RESOURCE_UPDATING]: 'updateError',
    [ACTIONS.CLEAN_UPDATE_RESOURCE_ERROR]: 'updateError',

    [ACTIONS.FAILED_DELETE_RESOURCE]: 'deleteError',
    [ACTIONS.FAILED_TO_PERFORM_RESOURCE_DELETING]: 'deleteError',
    [ACTIONS.CLEAN_DELETE_RESOURCE_ERROR]: 'deleteError',
};

export function resourceReducer(
    state: ResourcesResult = { list: null },
    action: Action<string> & { resourcesResult?: ResourcesActionResult }
): ResourcesResult {
    switch (action.type) {
        case ACTIONS.SUCCESS_CREATE_RESOURCE:
            return {
                ...state,
                list: (state.list || []).concat((action.resourcesResult as ResourceActionResult).resource as Resource),
            };

        case ACTIONS.SUCCESS_FETCH_RESOURCES:
            return {
                list: (action.resourcesResult as ResourcesListResult).list || [],
            };

        case ACTIONS.SUCCESS_UPDATE_RESOURCE:
            const newList = (state.list || []).slice();
            const updatedResource = (action.resourcesResult as ResourceActionResult).resource;
            const updatedResourceIndex = newList.findIndex(r => r.id === updatedResource.id);
            if (updatedResourceIndex >= 0) {
                newList[updatedResourceIndex] = {
                    ...newList[updatedResourceIndex],
                    title: updatedResource.title,
                    linksTo: updatedResource.linksTo,
                };
            }

            return {
                ...state,
                list: newList,
            };

        case ACTIONS.SUCCESS_DELETE_RESOURCE:
            const deletedResourceId = (action.resourcesResult as ResourceIdActionResult).resourceId;
            return {
                ...state,
                list: (state.list || []).filter(r => r.id !== deletedResourceId),
                deletedResourceId,
            };

        case ACTIONS.CLEAN_DELETED_RESOURCE_ID:
            return {
                ...state,
                deletedResourceId: '',
            };

        case ACTIONS.FAILED_CREATE_RESOURCE:
        case ACTIONS.FAILED_FETCH_RESOURCES:
        case ACTIONS.FAILED_UPDATE_RESOURCE:
        case ACTIONS.FAILED_DELETE_RESOURCE:
            return {
                ...state,
                [actionToErrorType[action.type]]: (action.resourcesResult as ResourceErrorResult).error as APIError,
            };

        case ACTIONS.FAILED_TO_PERFORM_RESOURCE_CREATION:
        case ACTIONS.FAILED_TO_PERFORM_RESOURCES_FETCHING:
        case ACTIONS.FAILED_TO_PERFORM_RESOURCE_UPDATING:
        case ACTIONS.FAILED_TO_PERFORM_RESOURCE_DELETING:
            log.error(
                `${actionToLogMessage[action.type]}: ${(action.resourcesResult as ResourceErrorResult).error.toString()}`
            );
            return {
                ...state,
                [actionToErrorType[action.type]]: {
                    code: UnknownErrorCode,
                    description: UnknownErrorDescription,
                },
            };

        case ACTIONS.CLEAN_CREATE_RESOURCE_ERROR:
        case ACTIONS.CLEAN_FETCH_RESOURCES_ERROR:
        case ACTIONS.CLEAN_UPDATE_RESOURCE_ERROR:
        case ACTIONS.CLEAN_DELETE_RESOURCE_ERROR:
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
