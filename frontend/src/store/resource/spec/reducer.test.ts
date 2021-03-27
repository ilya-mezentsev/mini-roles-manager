import { ACTIONS } from '../action_types';
import { actionToErrorType, resourceReducer } from '../reducer';
import { Resource } from '../../../services/api';
import {UnknownErrorCode, UnknownErrorDescription} from "../../shared/const";

describe('resources reducer tests', () => {
    it('reduce success resource creation', () => {
        expect(resourceReducer(undefined, {
            type: ACTIONS.SUCCESS_CREATE_RESOURCE,
            resourcesResult: {
                resource: {
                    id: 'resource-1'
                },
            },
        })).toEqual({
            list: [
                {
                    id: 'resource-1'
                },
            ],
        });
    });

    it('reduce success resources loading', () => {
        expect(resourceReducer(undefined, {
            type: ACTIONS.SUCCESS_FETCH_RESOURCES,
            resourcesResult: {
                list: [
                    {
                        id: 'resource-1'
                    } as Resource,
                ],
            }
        })).toEqual({
            list: [
                {
                    id: 'resource-1'
                },
            ],
        });

        expect(resourceReducer(undefined, {
            type: ACTIONS.SUCCESS_FETCH_RESOURCES,
            resourcesResult: {
                list: null,
            }
        })).toEqual({
            list: [],
        });
    });

    it('reduce success resource updating', () => {
        const initial = {
            list: [
                {
                    id: 'resource-1',
                    title: 'title-1',
                } as Resource,
            ],
        };

        expect(resourceReducer(initial, {
            type: ACTIONS.SUCCESS_UPDATE_RESOURCE,
            resourcesResult: {
                resource: {
                    id: 'resource-1',
                    title: 'title-2',
                } as Resource,
            },
        })).toEqual({
            list: [
                {
                    id: 'resource-1',
                    title: 'title-2',
                } as Resource,
            ],
        });

        expect(resourceReducer(initial, {
            type: ACTIONS.SUCCESS_UPDATE_RESOURCE,
            resourcesResult: {
                resource: {
                    id: 'resource-2',
                    title: 'title-2',
                } as Resource,
            },
        })).toEqual({
            list: [
                {
                    id: 'resource-1',
                    title: 'title-1',
                } as Resource,
            ],
        });

        expect(resourceReducer(undefined, {
            type: ACTIONS.SUCCESS_UPDATE_RESOURCE,
            resourcesResult: {
                resource: {
                    id: 'resource-2',
                    title: 'title-2',
                } as Resource,
            },
        })).toEqual({
            list: [],
        });
    });

    it('reduce success resource deleting', () => {
        const initial = {
            list: [
                {
                    id: 'resource-1',
                    title: 'title-1',
                } as Resource,
            ],
        };

        expect(resourceReducer(initial, {
            type: ACTIONS.SUCCESS_DELETE_RESOURCE,
            resourcesResult: {
                resourceId: 'resource-1',
            },
        })).toEqual({
            list: [],
        });

        expect(resourceReducer(undefined, {
            type: ACTIONS.SUCCESS_DELETE_RESOURCE,
            resourcesResult: {
                resourceId: 'resource-1',
            },
        })).toEqual({
            list: [],
        });
    });

    it('reduce known error of failed actions', () => {
        const error = {
            code: 'some-code',
            description: 'some-description',
        };

        for (const actionType of [
            ACTIONS.FAILED_CREATE_RESOURCE,
            ACTIONS.FAILED_FETCH_RESOURCES,
            ACTIONS.FAILED_UPDATE_RESOURCE,
            ACTIONS.FAILED_DELETE_RESOURCE,
        ]) {
            expect(resourceReducer(undefined, {
                type: actionType,
                resourcesResult: { error },
            })).toEqual({
                [actionToErrorType[actionType]]: error,
                list: null,
            });
        }
    });

    it('reduce unknown error of failed actions', () => {
        const error = new Error('some-error');

        for (const actionType of [
            ACTIONS.FAILED_TO_PERFORM_RESOURCE_CREATION,
            ACTIONS.FAILED_TO_PERFORM_RESOURCES_FETCHING,
            ACTIONS.FAILED_TO_PERFORM_RESOURCE_UPDATING,
            ACTIONS.FAILED_TO_PERFORM_RESOURCE_DELETING,
        ]) {
            expect(resourceReducer(undefined, {
                type: actionType,
                resourcesResult: { error },
            })).toEqual({
                [actionToErrorType[actionType]]: {
                    code: UnknownErrorCode,
                    description: UnknownErrorDescription,
                },
                list: null,
            });
        }
    });

    it('reduce clean actions', () => {
        for (const actionType of [
            ACTIONS.CLEAN_CREATE_RESOURCE_ERROR,
            ACTIONS.CLEAN_FETCH_RESOURCES_ERROR,
            ACTIONS.CLEAN_UPDATE_RESOURCE_ERROR,
            ACTIONS.CLEAN_DELETE_RESOURCE_ERROR,
        ]) {
            expect(resourceReducer(undefined, {
                type: actionType,
            })).toEqual({
                [actionToErrorType[actionType]]: null,
                list: null,
            });
        }
    });

    it('reduce unknown action', () => {
        expect(resourceReducer(undefined, {
            type: 'foo',
        })).toEqual({
            list: null,
        });
    });
});
