import configureMockStore from 'redux-mock-store';
import thunk from 'redux-thunk';

import * as api from '../../../services/api';
import { PermissionAccessRequest } from '../../../services/api';
import { Operation } from '../../../services/api/shared/types';
import { SuccessResponse, ErrorResponse } from '../../../services/api/shared';
import { fetchPermission, cleanFetchPermissionResult } from '../actions';
import { ACTIONS } from '../action_types';

jest.mock('../../../services/api');

describe('permission actions rests', () => {
    const middlewares = [thunk];
    const mockStore = configureMockStore(middlewares);

    it('fetch permission success', async () => {
        const d: PermissionAccessRequest = {
            roleId: 'role-1',
            resourceId: 'resource-1',
            operation: Operation.CREATE,
        };
        const store = mockStore({ fetchPermissionResult: null });
        // @ts-ignore
        api.fetchPermission = jest.fn().mockResolvedValue(new SuccessResponse('some-data'));

        // @ts-ignore
        await store.dispatch(fetchPermission(d));

        expect(api.fetchPermission).toBeCalledWith(d);
        expect(store.getActions()).toEqual([
            {
                type: ACTIONS.SUCCESS_FETCH_PERMISSION,
                fetchPermissionResult: 'some-data',
            }
        ]);
    });

    it('fetch permission parsed error', async () => {
        const d: PermissionAccessRequest = {
            roleId: 'role-1',
            resourceId: 'resource-1',
            operation: Operation.CREATE,
        };
        const store = mockStore({ fetchPermissionResult: null });
        // @ts-ignore
        api.fetchPermission = jest.fn().mockResolvedValue(new ErrorResponse('some-error'));

        // @ts-ignore
        await store.dispatch(fetchPermission(d));

        expect(api.fetchPermission).toBeCalledWith(d);
        expect(store.getActions()).toEqual([
            {
                type: ACTIONS.FAILED_FETCH_PERMISSION,
                fetchPermissionResult: {
                    error: 'some-error'
                },
            }
        ]);
    });

    it('fetch permission unknown error', async () => {
        const d: PermissionAccessRequest = {
            roleId: 'role-1',
            resourceId: 'resource-1',
            operation: Operation.CREATE,
        };
        const store = mockStore({ fetchPermissionResult: null });
        // @ts-ignore
        api.fetchPermission = jest.fn().mockRejectedValue('some-error');

        // @ts-ignore
        await store.dispatch(fetchPermission(d));

        expect(api.fetchPermission).toBeCalledWith(d);
        expect(store.getActions()).toEqual([
            {
                type: ACTIONS.FAILED_TO_PERFORM_PERMISSION_FETCHING,
                fetchPermissionResult: {
                    error: 'some-error'
                },
            }
        ]);
    });

    it('clean fetch permission result', async () => {
        const store = mockStore({ fetchPermissionResult: null });

        // @ts-ignore
        store.dispatch(cleanFetchPermissionResult());

        expect(store.getActions()).toEqual([
            {
                type: ACTIONS.CLEAN_FETCH_PERMISSION_RESULT,
            }
        ]);
    });
});
