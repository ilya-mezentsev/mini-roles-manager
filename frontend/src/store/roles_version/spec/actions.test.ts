import configureMockStore from 'redux-mock-store';
import thunk from 'redux-thunk';

import * as api from '../../../services/api';
import { SuccessResponse, ErrorResponse } from '../../../services/api/shared';

import { RolesVersion } from '../../../services/api';
import {
    createRolesVersion,
    cleanCreateRolesVersionError,

    fetchRolesVersion,
    cleanFetchRolesVersionError,

    updateRolesVersion,
    cleanUpdateRolesVersionError,

    deleteRolesVersion,
    cleanDeleteRolesVersionError,

    selectCurrentRolesVersion,
} from '../actions';
import { ACTIONS } from '../action_types';

jest.mock('../../../services/api');

describe('roles version actions tests', () => {
    const middlewares = [thunk];
    const mockStore = configureMockStore(middlewares);

    it('create roles version success', async () => {
        const d: RolesVersion = {
            id: 'roles-version-id',
            title: 'roles-version-title',
        };
        const store = mockStore({ rolesVersionResult: null });
        // @ts-ignore
        api.createRolesVersion = jest.fn().mockResolvedValue(new SuccessResponse(null));

        // @ts-ignore
        await store.dispatch(createRolesVersion(d));

        expect(api.createRolesVersion).toBeCalledWith(d);
        expect(store.getActions()).toEqual([
            {
                type: ACTIONS.SUCCESS_CREATE_ROLES_VERSION,
                rolesVersionResult: {
                    rolesVersion: d,
                },
            },
        ]);
    });

    it('create roles version parsed error', async () => {
        const d: RolesVersion = {
            id: 'roles-version-id',
            title: 'roles-version-title',
        };
        const store = mockStore({ rolesVersionResult: null });
        // @ts-ignore
        api.createRolesVersion = jest.fn().mockResolvedValue(new ErrorResponse('some-error'));

        // @ts-ignore
        await store.dispatch(createRolesVersion(d));

        expect(api.createRolesVersion).toBeCalledWith(d);
        expect(store.getActions()).toEqual([
            {
                type: ACTIONS.FAILED_CREATE_ROLES_VERSION,
                rolesVersionResult: {
                    error: 'some-error',
                },
            },
        ]);
    });

    it('create roles version unknown error', async () => {
        const d: RolesVersion = {
            id: 'roles-version-id',
            title: 'roles-version-title',
        };
        const store = mockStore({ rolesVersionResult: null });
        // @ts-ignore
        api.createRolesVersion = jest.fn().mockRejectedValue('some-error');

        // @ts-ignore
        await store.dispatch(createRolesVersion(d));

        expect(api.createRolesVersion).toBeCalledWith(d);
        expect(store.getActions()).toEqual([
            {
                type: ACTIONS.FAILED_TO_PERFORM_ROLES_VERSION_CREATION,
                rolesVersionResult: {
                    error: 'some-error',
                },
            },
        ]);
    });

    it('clean create roles version error', () => {
        const store = mockStore({ rolesResult: null });

        // @ts-ignore
        store.dispatch(cleanCreateRolesVersionError());

        expect(store.getActions()).toEqual([
            {
                type: ACTIONS.CLEAN_CREATE_ROLES_VERSION_ERROR,
            },
        ]);
    });

    it('fetch roles version success', async () => {
        const d: RolesVersion[] = [
            { id: 'id-1', title: 'title-1' },
            { id: 'id-2', title: 'title-2' },
        ];
        const store = mockStore({ rolesVersionResult: null });
        // @ts-ignore
        api.rolesVersionsList = jest.fn().mockResolvedValue(new SuccessResponse(d));

        // @ts-ignore
        await store.dispatch(fetchRolesVersion());

        expect(api.rolesVersionsList).toBeCalled();
        expect(store.getActions()).toEqual([
            {
                type: ACTIONS.SUCCESS_FETCH_ROLES_VERSION,
                rolesVersionResult: {
                    list: d,
                },
            },
        ]);
    });

    it('fetch roles version parsed error', async () => {
        const store = mockStore({ rolesVersionResult: null });
        // @ts-ignore
        api.rolesVersionsList = jest.fn().mockResolvedValue(new ErrorResponse('some-error'));

        // @ts-ignore
        await store.dispatch(fetchRolesVersion());

        expect(api.rolesVersionsList).toBeCalled();
        expect(store.getActions()).toEqual([
            {
                type: ACTIONS.FAILED_FETCH_ROLES_VERSION,
                rolesVersionResult: {
                    error: 'some-error',
                },
            },
        ]);
    });

    it('fetch roles version unknown error', async () => {
        const store = mockStore({ rolesVersionResult: null });
        // @ts-ignore
        api.rolesVersionsList = jest.fn().mockRejectedValue('some-error');

        // @ts-ignore
        await store.dispatch(fetchRolesVersion());

        expect(api.rolesVersionsList).toBeCalled();
        expect(store.getActions()).toEqual([
            {
                type: ACTIONS.FAILED_TO_PERFORM_ROLES_VERSION_FETCHING,
                rolesVersionResult: {
                    error: 'some-error',
                },
            },
        ]);
    });

    it('clean fetch roles version error', () => {
        const store = mockStore({ rolesVersionResult: null });

        // @ts-ignore
        store.dispatch(cleanFetchRolesVersionError());

        expect(store.getActions()).toEqual([
            {
                type: ACTIONS.CLEAN_FETCH_ROLES_VERSION_ERROR,
            },
        ]);
    });

    it('update roles version success', async () => {
        const d: RolesVersion = {
            id: 'roles-version-id',
            title: 'roles-version-title',
        };
        const store = mockStore({ rolesVersionResult: null });
        // @ts-ignore
        api.updateRolesVersion = jest.fn().mockResolvedValue(new SuccessResponse(null));

        // @ts-ignore
        await store.dispatch(updateRolesVersion(d));

        expect(api.updateRolesVersion).toBeCalledWith(d);
        expect(store.getActions()).toEqual([
            {
                type: ACTIONS.SUCCESS_UPDATE_ROLES_VERSION,
                rolesVersionResult: {
                    rolesVersion: d,
                },
            },
        ]);
    });

    it('update roles version parsed error', async () => {
        const d: RolesVersion = {
            id: 'roles-version-id',
            title: 'roles-version-title',
        };
        const store = mockStore({ rolesVersionResult: null });
        // @ts-ignore
        api.updateRolesVersion = jest.fn().mockResolvedValue(new ErrorResponse('some-error'));

        // @ts-ignore
        await store.dispatch(updateRolesVersion(d));

        expect(api.updateRolesVersion).toBeCalledWith(d);
        expect(store.getActions()).toEqual([
            {
                type: ACTIONS.FAILED_UPDATE_ROLES_VERSION,
                rolesVersionResult: {
                    error: 'some-error',
                },
            },
        ]);
    });

    it('update roles version unknown error', async () => {
        const d: RolesVersion = {
            id: 'roles-version-id',
            title: 'roles-version-title',
        };
        const store = mockStore({ rolesVersionResult: null });
        // @ts-ignore
        api.updateRolesVersion = jest.fn().mockRejectedValue('some-error');

        // @ts-ignore
        await store.dispatch(updateRolesVersion(d));

        expect(api.updateRolesVersion).toBeCalledWith(d);
        expect(store.getActions()).toEqual([
            {
                type: ACTIONS.FAILED_TO_PERFORM_ROLES_VERSION_UPDATING,
                rolesVersionResult: {
                    error: 'some-error',
                },
            },
        ]);
    });

    it('clean update roles version error', async () => {
        const store = mockStore({ rolesVersionResult: null });

        // @ts-ignore
        store.dispatch(cleanUpdateRolesVersionError());

        expect(store.getActions()).toEqual([
            {
                type: ACTIONS.CLEAN_UPDATE_ROLES_VERSION_ERROR,
            },
        ]);
    });

    it('delete roles version success', async () => {
        const d = 'roles-version-id';
        const store = mockStore({ rolesVersionResult: null });
        // @ts-ignore
        api.deleteRolesVersion = jest.fn().mockResolvedValue(new SuccessResponse(null));

        // @ts-ignore
        await store.dispatch(deleteRolesVersion(d));

        expect(api.deleteRolesVersion).toBeCalledWith(d);
        expect(store.getActions()).toEqual([
            {
                type: ACTIONS.SUCCESS_DELETE_ROLES_VERSION,
                rolesVersionResult: {
                    rolesVersionId: d,
                },
            },
        ]);
    });

    it('delete roles version parsed error', async () => {
        const d = 'roles-version-id';
        const store = mockStore({ rolesVersionResult: null });
        // @ts-ignore
        api.deleteRolesVersion = jest.fn().mockResolvedValue(new ErrorResponse('some-error'));

        // @ts-ignore
        await store.dispatch(deleteRolesVersion(d));

        expect(api.deleteRolesVersion).toBeCalledWith(d);
        expect(store.getActions()).toEqual([
            {
                type: ACTIONS.FAILED_DELETE_ROLES_VERSION,
                rolesVersionResult: {
                    error: 'some-error',
                },
            },
        ]);
    });

    it('delete roles version unknown error', async () => {
        const d = 'roles-version-id';
        const store = mockStore({ rolesVersionResult: null });
        // @ts-ignore
        api.deleteRolesVersion = jest.fn().mockRejectedValue('some-error');

        // @ts-ignore
        await store.dispatch(deleteRolesVersion(d));

        expect(api.deleteRolesVersion).toBeCalledWith(d);
        expect(store.getActions()).toEqual([
            {
                type: ACTIONS.FAILED_TO_PERFORM_ROLES_VERSION_DELETING,
                rolesVersionResult: {
                    error: 'some-error',
                },
            },
        ]);
    });

    it('clean roles version delete error', () => {
        const store = mockStore({ rolesVersionResult: null });

        // @ts-ignore
        store.dispatch(cleanDeleteRolesVersionError());

        expect(store.getActions()).toEqual([
            {
                type: ACTIONS.CLEAN_DELETE_ROLES_VERSION_ERROR,
            },
        ]);
    });

    it('select current roles version', () => {
        const d: RolesVersion = {
            id: 'roles-version-id',
            title: 'roles-version-title',
        };
        const store = mockStore({ rolesVersionResult: null });

        // @ts-ignore
        store.dispatch(selectCurrentRolesVersion(d));

        expect(store.getActions()).toEqual([
            {
                type: ACTIONS.SELECT_CURRENT_ROLES_VERSION,
                rolesVersionResult: {
                    rolesVersion: d,
                },
            },
        ]);
    });
});
