import * as request from '../shared/request';
import {
    createRole,
    rolesList,
    updateRole,
    deleteRole,
} from './role';
import { Role } from './role.types';
import { APIResponseStatus, ErrorResponse, SuccessResponse } from '../shared';

jest.mock('../shared/request');

describe('role api tests', () => {
    it('create role success', async () => {
        const d: Role = {
            id: 'role-id',
            versionId: 'version-id',
        };
        // @ts-ignore
        request.POST = jest.fn().mockResolvedValue(null);

        const response = await createRole(d);

        expect(request.POST).toBeCalledWith('/role', { role: d });
        expect(response).toBeInstanceOf(SuccessResponse);
        expect(response.isOk()).toBeTruthy();
        expect(response.data()).toBeNull();
    });

    it('create tole error', async () => {
        const d: Role = {
            id: 'role-id',
            versionId: 'version-id',
        };
        // @ts-ignore
        request.POST = jest.fn().mockResolvedValue({
            status: APIResponseStatus.ERROR,
            data: 'some-error',
        });

        const response = await createRole(d);

        expect(request.POST).toBeCalledWith('/role', { role: d });
        expect(response).toBeInstanceOf(ErrorResponse);
        expect(response.isOk()).toBeFalsy();
        expect(response.data()).toEqual('some-error');
    });

    it('roles list success', async () => {
        const d: Role[] = [
            { id: 'role-1', versionId: 'version-id' },
            { id: 'role-2', versionId: 'version-id' },
        ];
        // @ts-ignore
        request.GET = jest.fn().mockResolvedValue({
            status: APIResponseStatus.OK,
            data: d,
        });

        const response = await rolesList();

        expect(request.GET).toBeCalledWith('/roles');
        expect(response).toBeInstanceOf(SuccessResponse);
        expect(response.isOk()).toBeTruthy();
        expect(response.data()).toEqual(d);
    });

    it('roles list error', async () => {
        // @ts-ignore
        request.GET = jest.fn().mockResolvedValue({
            status: APIResponseStatus.ERROR,
            data: 'some-error',
        });

        const response = await rolesList();

        expect(request.GET).toBeCalledWith('/roles');
        expect(response).toBeInstanceOf(ErrorResponse);
        expect(response.isOk()).toBeFalsy();
        expect(response.data()).toEqual('some-error');
    });

    it('update role success', async () => {
        const d: Role = {
            id: 'role-id',
            versionId: 'version-id',
        };
        // @ts-ignore
        request.PATCH = jest.fn().mockResolvedValue(null);

        const response = await updateRole(d);

        expect(request.PATCH).toBeCalledWith('/role', { role: d });
        expect(response).toBeInstanceOf(SuccessResponse);
        expect(response.isOk()).toBeTruthy();
        expect(response.data()).toBeNull();
    });

    it('update role error', async () => {
        const d: Role = {
            id: 'role-id',
            versionId: 'version-id',
        };
        // @ts-ignore
        request.PATCH = jest.fn().mockResolvedValue({
            status: APIResponseStatus.ERROR,
            data: 'some-error',
        });

        const response = await updateRole(d);

        expect(request.PATCH).toBeCalledWith('/role', { role: d });
        expect(response).toBeInstanceOf(ErrorResponse);
        expect(response.isOk()).toBeFalsy();
        expect(response.data()).toEqual('some-error');
    });

    it('delete role success', async () => {
        const roleId = 'role-id';
        const rolesVersionId = 'roles-version-id';
        // @ts-ignore
        request.DELETE = jest.fn().mockResolvedValue(null);

        const response = await deleteRole(rolesVersionId, roleId);

        expect(request.DELETE).toBeCalledWith(`/role/${rolesVersionId}/${roleId}`);
        expect(response).toBeInstanceOf(SuccessResponse);
        expect(response.isOk()).toBeTruthy();
        expect(response.data()).toBeNull();
    });

    it('delete role error', async () => {
        const roleId = 'role-id';
        const rolesVersionId = 'roles-version-id';
        // @ts-ignore
        request.DELETE = jest.fn().mockResolvedValue({
            status: APIResponseStatus.ERROR,
            data: 'some-error',
        });

        const response = await deleteRole(rolesVersionId, roleId);

        expect(request.DELETE).toBeCalledWith(`/role/${rolesVersionId}/${roleId}`);
        expect(response).toBeInstanceOf(ErrorResponse);
        expect(response.isOk()).toBeFalsy();
        expect(response.data()).toEqual('some-error');
    });
});
