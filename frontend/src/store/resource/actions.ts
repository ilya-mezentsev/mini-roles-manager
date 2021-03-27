import { Dispatch } from 'redux';

import {
    EditableResource,
    createResource as createResourceAPI,
    resourcesList,
    updateResource as updateResourceAPI,
    deleteResource as deleteResourceAPI,
} from '../../services/api';
import { ACTIONS } from './action_types';

export function createResource(resource: EditableResource): (dispatch: Dispatch) => Promise<void> {
    return async dispatch => {
        try {
            const response = await createResourceAPI(resource);

            if (response.isOk()) {
                dispatch({
                    type: ACTIONS.SUCCESS_CREATE_RESOURCE,
                    resourcesResult: { resource },
                });
            } else {
                dispatch({
                    type: ACTIONS.FAILED_CREATE_RESOURCE,
                    resourcesResult: {
                        error: response.data(),
                    },
                });
            }
        } catch (e) {
            dispatch({
                type: ACTIONS.FAILED_TO_PERFORM_RESOURCE_CREATION,
                resourcesResult: {
                    error: e,
                },
            });
        }
    };
}

export function cleanCreateResourceError(): (dispatch: Dispatch) => void {
    return dispatch => {
        dispatch({
            type: ACTIONS.CLEAN_CREATE_RESOURCE_ERROR,
        });
    };
}

export function loadResources(): (dispatch: Dispatch) => Promise<void> {
    return async dispatch => {
        try {
            const response = await resourcesList();

            if (response.isOk()) {
                dispatch({
                    type: ACTIONS.SUCCESS_FETCH_RESOURCES,
                    resourcesResult: {
                        list: response.data(),
                    },
                });
            } else {
                dispatch({
                    type: ACTIONS.FAILED_FETCH_RESOURCES,
                    resourcesResult: {
                        error: response.data(),
                    },
                });
            }
        } catch (e) {
            dispatch({
                type: ACTIONS.FAILED_TO_PERFORM_RESOURCES_FETCHING,
                resourcesResult: {
                    error: e,
                },
            });
        }
    };
}

export function cleanLoadResourcesError(): (dispatch: Dispatch) => void {
    return dispatch => {
        dispatch({
            type: ACTIONS.CLEAN_FETCH_RESOURCES_ERROR,
        });
    };
}

export function updateResource(resource: EditableResource): (dispatch: Dispatch) => Promise<void> {
    return async dispatch => {
        try {
            const response = await updateResourceAPI(resource);

            if (response.isOk()) {
                dispatch({
                    type: ACTIONS.SUCCESS_UPDATE_RESOURCE,
                    resourcesResult: { resource },
                });
            } else {
                dispatch({
                    type: ACTIONS.FAILED_UPDATE_RESOURCE,
                    resourcesResult: {
                        error: response.data(),
                    },
                });
            }
        } catch (e) {
            dispatch({
                type: ACTIONS.FAILED_TO_PERFORM_RESOURCE_UPDATING,
                resourcesResult: {
                    error: e,
                },
            });
        }
    };
}

export function cleanUpdateResourceError(): (dispatch: Dispatch) => void {
    return dispatch => {
        dispatch({
            type: ACTIONS.CLEAN_UPDATE_RESOURCE_ERROR,
        });
    };
}

export function deleteResource(resourceId: string): (dispatch: Dispatch) => Promise<void> {
    return async dispatch => {
        try {
            const response = await deleteResourceAPI(resourceId);

            if (response.isOk()) {
                dispatch({
                    type: ACTIONS.SUCCESS_DELETE_RESOURCE,
                    resourcesResult: { resourceId },
                });
            } else {
                dispatch({
                    type: ACTIONS.FAILED_DELETE_RESOURCE,
                    resourcesResult: {
                        error: response.data(),
                    },
                });
            }
        } catch (e) {
            dispatch({
                type: ACTIONS.FAILED_TO_PERFORM_RESOURCE_DELETING,
                resourcesResult: {
                    error: e,
                },
            });
        }
    };
}

export function cleanDeleteResourceError(): (dispatch: Dispatch) => void {
    return dispatch => {
        dispatch({
            type: ACTIONS.CLEAN_DELETE_RESOURCE_ERROR,
        });
    };
}
