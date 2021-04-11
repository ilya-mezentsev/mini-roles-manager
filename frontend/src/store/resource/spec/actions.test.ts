import configureMockStore from 'redux-mock-store';
import thunk from 'redux-thunk';

import * as api from '../../../services/api';
import { SuccessResponse, ErrorResponse } from '../../../services/api/shared';
import { EditableResource } from '../../../services/api';
import {
    createResource,
    cleanCreateResourceError,

    fetchResources,
    cleanLoadResourcesError,

    updateResource,
    cleanUpdateResourceError,

    deleteResource,
    cleanDeleteResourceError,
} from '../actions';
import { ACTIONS } from '../action_types';

jest.mock('../../../services/api');

describe('resources actions tests', () => {
    const middlewares = [thunk];
    const mockStore = configureMockStore(middlewares);

    it('create resource success', async () => {
        const d: EditableResource = {
            id: 'foo',
            title: 'bar',
        };
        const store = mockStore({ resourcesResult: null });
        // @ts-ignore
        api.createResource = jest.fn().mockResolvedValue(new SuccessResponse(null));

        // @ts-ignore
        await store.dispatch(createResource(d));

        expect(api.createResource).toBeCalledWith(d);
        expect(store.getActions()).toEqual([
            {
                type: ACTIONS.SUCCESS_CREATE_RESOURCE,
                resourcesResult: {
                    resource: d,
                },
            },
        ]);
    });

    it('create resource parsed error', async () => {
        const d: EditableResource = {
            id: 'foo',
            title: 'bar',
        };
        const store = mockStore({ resourcesResult: null });
        // @ts-ignore
        api.createResource = jest.fn().mockResolvedValue(new ErrorResponse('some-error'));

        // @ts-ignore
        await store.dispatch(createResource(d));

        expect(api.createResource).toBeCalledWith(d);
        expect(store.getActions()).toEqual([
            {
                type: ACTIONS.FAILED_CREATE_RESOURCE,
                resourcesResult: {
                    error: 'some-error',
                },
            },
        ]);
    });

    it('create resource unknown error', async () => {
        const d: EditableResource = {
            id: 'foo',
            title: 'bar',
        };
        const store = mockStore({ resourcesResult: null });
        // @ts-ignore
        api.createResource = jest.fn().mockRejectedValue('some-error');

        // @ts-ignore
        await store.dispatch(createResource(d));

        expect(api.createResource).toBeCalledWith(d);
        expect(store.getActions()).toEqual([
            {
                type: ACTIONS.FAILED_TO_PERFORM_RESOURCE_CREATION,
                resourcesResult: {
                    error: 'some-error',
                },
            },
        ]);
    });

    it('clean create resource error', () => {
        const store = mockStore({ resourcesResult: null });

        // @ts-ignore
        store.dispatch(cleanCreateResourceError());

        expect(store.getActions()).toEqual([
            {
                type: ACTIONS.CLEAN_CREATE_RESOURCE_ERROR,
            },
        ]);
    });

    it('load resources success', async () => {
        const store = mockStore({ resourcesResult: null });
        // @ts-ignore
        api.resourcesList = jest.fn().mockResolvedValue(new SuccessResponse('some-data'));

        // @ts-ignore
        await store.dispatch(fetchResources());

        expect(api.resourcesList).toBeCalled();
        expect(store.getActions()).toEqual([
            {
                type: ACTIONS.SUCCESS_FETCH_RESOURCES,
                resourcesResult: {
                    list: 'some-data',
                },
            },
        ]);
    });

    it('load resources parsed error', async () => {
        const store = mockStore({ resourcesResult: null });
        // @ts-ignore
        api.resourcesList = jest.fn().mockResolvedValue(new ErrorResponse('some-error'));

        // @ts-ignore
        await store.dispatch(fetchResources());

        expect(api.resourcesList).toBeCalled();
        expect(store.getActions()).toEqual([
            {
                type: ACTIONS.FAILED_FETCH_RESOURCES,
                resourcesResult: {
                    error: 'some-error',
                },
            },
        ]);
    });

    it('load resources unknown error', async () => {
        const store = mockStore({ resourcesResult: null });
        // @ts-ignore
        api.resourcesList = jest.fn().mockRejectedValue('some-error');

        // @ts-ignore
        await store.dispatch(fetchResources());

        expect(api.resourcesList).toBeCalled();
        expect(store.getActions()).toEqual([
            {
                type: ACTIONS.FAILED_TO_PERFORM_RESOURCES_FETCHING,
                resourcesResult: {
                    error: 'some-error',
                },
            },
        ]);
    });

    it('clean load resources error', () => {
        const store = mockStore({ resourcesResult: null });

        // @ts-ignore
        store.dispatch(cleanLoadResourcesError());

        expect(store.getActions()).toEqual([
            {
                type: ACTIONS.CLEAN_FETCH_RESOURCES_ERROR,
            },
        ]);
    });

    it('update resource success', async () => {
        const d: EditableResource = {
            id: 'foo',
            title: 'bar',
        };
        const store = mockStore({ resourcesResult: null });
        // @ts-ignore
        api.updateResource = jest.fn().mockResolvedValue(new SuccessResponse(null));

        // @ts-ignore
        await store.dispatch(updateResource(d));

        expect(api.updateResource).toBeCalledWith(d);
        expect(store.getActions()).toEqual([
            {
                type: ACTIONS.SUCCESS_UPDATE_RESOURCE,
                resourcesResult: {
                    resource: d,
                },
            },
        ]);
    });

    it('update resource parsed error', async () => {
        const d: EditableResource = {
            id: 'foo',
            title: 'bar',
        };
        const store = mockStore({ resourcesResult: null });
        // @ts-ignore
        api.updateResource = jest.fn().mockResolvedValue(new ErrorResponse('some-error'));

        // @ts-ignore
        await store.dispatch(updateResource(d));

        expect(api.updateResource).toBeCalledWith(d);
        expect(store.getActions()).toEqual([
            {
                type: ACTIONS.FAILED_UPDATE_RESOURCE,
                resourcesResult: {
                    error: 'some-error',
                },
            },
        ]);
    });

    it('update resource unknown error', async () => {
        const d: EditableResource = {
            id: 'foo',
            title: 'bar',
        };
        const store = mockStore({ resourcesResult: null });
        // @ts-ignore
        api.updateResource = jest.fn().mockRejectedValue('some-error');

        // @ts-ignore
        await store.dispatch(updateResource(d));

        expect(api.updateResource).toBeCalledWith(d);
        expect(store.getActions()).toEqual([
            {
                type: ACTIONS.FAILED_TO_PERFORM_RESOURCE_UPDATING,
                resourcesResult: {
                    error: 'some-error',
                },
            },
        ]);
    });

    it('clean update resource error', () => {
        const store = mockStore({ resourcesResult: null });

        // @ts-ignore
        store.dispatch(cleanUpdateResourceError());

        expect(store.getActions()).toEqual([
            {
                type: ACTIONS.CLEAN_UPDATE_RESOURCE_ERROR,
            },
        ]);
    });

    it('delete resource success', async () => {
        const d = 'resource-id';
        const store = mockStore({ resourcesResult: null });
        // @ts-ignore
        api.deleteResource = jest.fn().mockResolvedValue(new SuccessResponse(null));

        // @ts-ignore
        await store.dispatch(deleteResource(d));

        expect(api.deleteResource).toBeCalledWith(d);
        expect(store.getActions()).toEqual([
            {
                type: ACTIONS.SUCCESS_DELETE_RESOURCE,
                resourcesResult: {
                    resourceId: d,
                },
            },
        ]);
    });

    it('delete resource parsed error', async () => {
        const d = 'resource-id';
        const store = mockStore({ resourcesResult: null });
        // @ts-ignore
        api.deleteResource = jest.fn().mockResolvedValue(new ErrorResponse('some-error'));

        // @ts-ignore
        await store.dispatch(deleteResource(d));

        expect(api.deleteResource).toBeCalledWith(d);
        expect(store.getActions()).toEqual([
            {
                type: ACTIONS.FAILED_DELETE_RESOURCE,
                resourcesResult: {
                    error: 'some-error',
                },
            },
        ]);
    });

    it('delete resource unknown error', async () => {
        const d = 'resource-id';
        const store = mockStore({ resourcesResult: null });
        // @ts-ignore
        api.deleteResource = jest.fn().mockRejectedValue('some-error');

        // @ts-ignore
        await store.dispatch(deleteResource(d));

        expect(api.deleteResource).toBeCalledWith(d);
        expect(store.getActions()).toEqual([
            {
                type: ACTIONS.FAILED_TO_PERFORM_RESOURCE_DELETING,
                resourcesResult: {
                    error: 'some-error',
                },
            },
        ]);
    });

    it('clean delete resource error', () => {
        const store = mockStore({ resourcesResult: null });

        // @ts-ignore
        store.dispatch(cleanDeleteResourceError());

        expect(store.getActions()).toEqual([
            {
                type: ACTIONS.CLEAN_DELETE_RESOURCE_ERROR,
            },
        ]);
    });
});
