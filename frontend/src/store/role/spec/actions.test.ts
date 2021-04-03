import configureMockStore from 'redux-mock-store';
import thunk from 'redux-thunk';

import * as api from '../../../services/api';
import { SuccessResponse, ErrorResponse } from '../../../services/api/shared';
import { EditableResource } from '../../../services/api';
import {
    createRole,
    cleanCreateRoleError,

    fetchRoles,
    cleanFetchRolesError,

    updateRole,
    cleanUpdateRoleError,

    deleteRole,
    cleanDeleteRoleError,
} from '../actions';
import { ACTIONS } from '../action_types';

jest.mock('../../../services/api');

describe('roles actions tests', () => {
    const middlewares = [thunk];
    const mockStore = configureMockStore(middlewares);

    it('create role success', async () => {
        const d: EditableResource = {
            id: 'foo',
            title: 'bar',
        };
        const store = mockStore({ rolesResult: null });
        // @ts-ignore
        api.createRole = jest.fn().mockResolvedValue(new SuccessResponse(null));

        // @ts-ignore
        await store.dispatch(createRole(d));

        expect(api.createRole).toBeCalledWith(d);
        expect(store.getActions()).toEqual([
            {
                type: ACTIONS.SUCCESS_CREATE_ROLE,
                rolesResult: {
                    role: d,
                },
            },
        ]);
    });

    it('create role parsed error', async () => {
        const d: EditableResource = {
            id: 'foo',
            title: 'bar',
        };
        const store = mockStore({ rolesResult: null });
        // @ts-ignore
        api.createRole = jest.fn().mockResolvedValue(new ErrorResponse('some-error'));

        // @ts-ignore
        await store.dispatch(createRole(d));

        expect(api.createRole).toBeCalledWith(d);
        expect(store.getActions()).toEqual([
            {
                type: ACTIONS.FAILED_CREATE_ROLE,
                rolesResult: {
                    error: 'some-error',
                },
            },
        ]);
    });

    it('create role unknown error', async () => {
        const d: EditableResource = {
            id: 'foo',
            title: 'bar',
        };
        const store = mockStore({ rolesResult: null });
        // @ts-ignore
        api.createRole = jest.fn().mockRejectedValue('some-error');

        // @ts-ignore
        await store.dispatch(createRole(d));

        expect(api.createRole).toBeCalledWith(d);
        expect(store.getActions()).toEqual([
            {
                type: ACTIONS.FAILED_TO_PERFORM_ROLE_CREATION,
                rolesResult: {
                    error: 'some-error',
                },
            },
        ]);
    });

    it('clean create role error', () => {
        const store = mockStore({ rolesResult: null });

        // @ts-ignore
        store.dispatch(cleanCreateRoleError());

        expect(store.getActions()).toEqual([
            {
                type: ACTIONS.CLEAN_CREATE_ROLE_ERROR,
            },
        ]);
    });

    it('load roles success', async () => {
        const store = mockStore({ rolesResult: null });
        // @ts-ignore
        api.rolesList = jest.fn().mockResolvedValue(new SuccessResponse('some-data'));

        // @ts-ignore
        await store.dispatch(fetchRoles());

        expect(api.rolesList).toBeCalled();
        expect(store.getActions()).toEqual([
            {
                type: ACTIONS.SUCCESS_FETCH_ROLES,
                rolesResult: {
                    list: 'some-data',
                },
            },
        ]);
    });

    it('load roles parsed error', async () => {
        const store = mockStore({ rolesResult: null });
        // @ts-ignore
        api.rolesList = jest.fn().mockResolvedValue(new ErrorResponse('some-error'));

        // @ts-ignore
        await store.dispatch(fetchRoles());

        expect(api.rolesList).toBeCalled();
        expect(store.getActions()).toEqual([
            {
                type: ACTIONS.FAILED_FETCH_ROLES,
                rolesResult: {
                    error: 'some-error',
                },
            },
        ]);
    });

    it('load roles unknown error', async () => {
        const store = mockStore({ rolesResult: null });
        // @ts-ignore
        api.rolesList = jest.fn().mockRejectedValue('some-error');

        // @ts-ignore
        await store.dispatch(fetchRoles());

        expect(api.rolesList).toBeCalled();
        expect(store.getActions()).toEqual([
            {
                type: ACTIONS.FAILED_TO_PERFORM_ROLES_FETCHING,
                rolesResult: {
                    error: 'some-error',
                },
            },
        ]);
    });

    it('clean load roles error', () => {
        const store = mockStore({ rolesResult: null });

        // @ts-ignore
        store.dispatch(cleanFetchRolesError());

        expect(store.getActions()).toEqual([
            {
                type: ACTIONS.CLEAN_FETCH_ROLES_ERROR,
            },
        ]);
    });

    it('update role success', async () => {
        const d: EditableResource = {
            id: 'foo',
            title: 'bar',
        };
        const store = mockStore({ rolesResult: null });
        // @ts-ignore
        api.updateRole = jest.fn().mockResolvedValue(new SuccessResponse(null));

        // @ts-ignore
        await store.dispatch(updateRole(d));

        expect(api.updateRole).toBeCalledWith(d);
        expect(store.getActions()).toEqual([
            {
                type: ACTIONS.SUCCESS_UPDATE_ROLE,
                rolesResult: {
                    role: d,
                },
            },
        ]);
    });

    it('update role parsed error', async () => {
        const d: EditableResource = {
            id: 'foo',
            title: 'bar',
        };
        const store = mockStore({ rolesResult: null });
        // @ts-ignore
        api.updateRole = jest.fn().mockResolvedValue(new ErrorResponse('some-error'));

        // @ts-ignore
        await store.dispatch(updateRole(d));

        expect(api.updateRole).toBeCalledWith(d);
        expect(store.getActions()).toEqual([
            {
                type: ACTIONS.FAILED_UPDATE_ROLE,
                rolesResult: {
                    error: 'some-error',
                },
            },
        ]);
    });

    it('update role unknown error', async () => {
        const d: EditableResource = {
            id: 'foo',
            title: 'bar',
        };
        const store = mockStore({ rolesResult: null });
        // @ts-ignore
        api.updateRole = jest.fn().mockRejectedValue('some-error');

        // @ts-ignore
        await store.dispatch(updateRole(d));

        expect(api.updateRole).toBeCalledWith(d);
        expect(store.getActions()).toEqual([
            {
                type: ACTIONS.FAILED_TO_PERFORM_ROLE_UPDATING,
                rolesResult: {
                    error: 'some-error',
                },
            },
        ]);
    });

    it('clean update role error', () => {
        const store = mockStore({ rolesResult: null });

        // @ts-ignore
        store.dispatch(cleanUpdateRoleError());

        expect(store.getActions()).toEqual([
            {
                type: ACTIONS.CLEAN_UPDATE_ROLE_ERROR,
            },
        ]);
    });

    it('delete role success', async () => {
        const d = 'role-id';
        const store = mockStore({ rolesResult: null });
        // @ts-ignore
        api.deleteRole = jest.fn().mockResolvedValue(new SuccessResponse(null));

        // @ts-ignore
        await store.dispatch(deleteRole(d));

        expect(api.deleteRole).toBeCalledWith(d);
        expect(store.getActions()).toEqual([
            {
                type: ACTIONS.SUCCESS_DELETE_ROLE,
                rolesResult: {
                    roleId: d,
                },
            },
        ]);
    });

    it('delete role parsed error', async () => {
        const d = 'role-id';
        const store = mockStore({ rolesResult: null });
        // @ts-ignore
        api.deleteRole = jest.fn().mockResolvedValue(new ErrorResponse('some-error'));

        // @ts-ignore
        await store.dispatch(deleteRole(d));

        expect(api.deleteRole).toBeCalledWith(d);
        expect(store.getActions()).toEqual([
            {
                type: ACTIONS.FAILED_DELETE_ROLE,
                rolesResult: {
                    error: 'some-error',
                },
            },
        ]);
    });

    it('delete role unknown error', async () => {
        const d = 'role-id';
        const store = mockStore({ rolesResult: null });
        // @ts-ignore
        api.deleteRole = jest.fn().mockRejectedValue('some-error');

        // @ts-ignore
        await store.dispatch(deleteRole(d));

        expect(api.deleteRole).toBeCalledWith(d);
        expect(store.getActions()).toEqual([
            {
                type: ACTIONS.FAILED_TO_PERFORM_ROLE_DELETING,
                rolesResult: {
                    error: 'some-error',
                },
            },
        ]);
    });

    it('clean delete role error', () => {
        const store = mockStore({ rolesResult: null });

        // @ts-ignore
        store.dispatch(cleanDeleteRoleError());

        expect(store.getActions()).toEqual([
            {
                type: ACTIONS.CLEAN_DELETE_ROLE_ERROR,
            },
        ]);
    });
});
