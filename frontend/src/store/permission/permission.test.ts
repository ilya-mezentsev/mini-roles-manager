import { PermissionStore } from './permission.store';

import * as api from '../../services/api';
import { ErrorResponse, SuccessResponse } from '../../services/api/shared';
import { UnknownErrorCode, UnknownErrorDescription } from '../shared/const';
import {Operation} from "../../services/api/shared/types";

jest.mock('../../services/api');

describe('fetch permission tests', () => {
    let permissionStore = new PermissionStore();
    const d = {
        rolesVersionId: 'rolesVersionId',
        roleId: 'roleId',
        resourceId: 'resourceId',
        operation: Operation.DELETE,
    };

    beforeEach(() => {
        permissionStore = new PermissionStore();
    });

    it('fetch permission success', async () => {
        // @ts-ignore
        api.fetchPermission = jest.fn().mockResolvedValue(new SuccessResponse('some-data'));

        await permissionStore.fetchPermission(d);

        expect(api.fetchPermission).toBeCalledWith(d);
        expect(permissionStore.permission).toEqual('some-data');
    });

    it('fetch permission parsed error', async () => {
        // @ts-ignore
        api.fetchPermission = jest.fn().mockResolvedValue(new ErrorResponse('some-error'));

        await permissionStore.fetchPermission(d);

        expect(api.fetchPermission).toBeCalledWith(d);
        expect(permissionStore.permission).toBeNull();
        expect(permissionStore.fetchPermissionError).toEqual('some-error');
    });

    it('fetch permission unknown error', async () => {
        // @ts-ignore
        api.fetchPermission = jest.fn().mockRejectedValue('some-error');

        await permissionStore.fetchPermission(d);

        expect(api.fetchPermission).toBeCalledWith(d);
        expect(permissionStore.permission).toBeNull();
        expect(permissionStore.fetchPermissionError).toEqual({
            code: UnknownErrorCode,
            description: UnknownErrorDescription,
        });
    });

    it('clean permission error', () => {
        permissionStore.fetchPermissionError = 'some-error' as any;
        permissionStore.permission = 'some-permission' as any;

        permissionStore.cleanFetchPermissionAction();

        expect(permissionStore.permission).toBeNull();
        expect(permissionStore.fetchPermissionError).toBeNull();
    });
});
