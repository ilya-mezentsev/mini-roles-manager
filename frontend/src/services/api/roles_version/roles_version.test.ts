import * as request from '../shared/request';
import {
    createRolesVersion,
    rolesVersionsList,
    updateRolesVersion,
    deleteRolesVersion,
} from './roles_version';
import { RolesVersion } from './roles_version.types';
import { APIResponseStatus, ErrorResponse, SuccessResponse } from '../shared';

jest.mock('../shared/request');

describe('roles version api tests', () => {
    it('create role success', async () => {
        const d: RolesVersion = {
            id: 'some-id',
            title: 'some-title',
        };
        // @ts-ignore
        request.POST = jest.fn().mockResolvedValue(null);

        const response = await createRolesVersion(d);

        expect(request.POST).toBeCalledWith('/roles-version', { rolesVersion: d });
        expect(response).toBeInstanceOf(SuccessResponse);
        expect(response.isOk()).toBeTruthy();
        expect(response.data()).toBeNull();
    });

    it('create role error', async () => {
        const d: RolesVersion = {
            id: 'some-id',
            title: 'some-title',
        };
        // @ts-ignore
        request.POST = jest.fn().mockResolvedValue({
            status: APIResponseStatus.ERROR,
            data: 'some-error',
        });

        const response = await createRolesVersion(d);

        expect(request.POST).toBeCalledWith('/roles-version', { rolesVersion: d });
        expect(response).toBeInstanceOf(ErrorResponse);
        expect(response.isOk()).toBeFalsy();
        expect(response.data()).toEqual('some-error');
    });

    it('roles versions list success', async () => {
        const d: RolesVersion[] = [
            { id: 'id-1', title: 'title-1' },
            { id: 'id-2', title: 'title-2' },
        ];
        // @ts-ignore
        request.GET = jest.fn().mockResolvedValue({
            status: APIResponseStatus.OK,
            data: d,
        });

        const response = await rolesVersionsList();

        expect(request.GET).toBeCalledWith('/roles-versions');
        expect(response).toBeInstanceOf(SuccessResponse);
        expect(response.isOk()).toBeTruthy();
        expect(response.data()).toEqual(d);
    });

    it('roles versions list error', async () => {
        // @ts-ignore
        request.GET = jest.fn().mockResolvedValue({
            status: APIResponseStatus.ERROR,
            data: 'some-error',
        });

        const response = await rolesVersionsList();

        expect(request.GET).toBeCalledWith('/roles-versions');
        expect(response).toBeInstanceOf(ErrorResponse);
        expect(response.isOk()).toBeFalsy();
        expect(response.data()).toEqual('some-error');
    });

    it('update roles version success', async () => {
        const d: RolesVersion = {
            id: 'some-id',
            title: 'some-title',
        };
        // @ts-ignore
        request.PATCH = jest.fn().mockResolvedValue(null);

        const response = await updateRolesVersion(d);

        expect(request.PATCH).toBeCalledWith('/roles-version', { rolesVersion: d });
        expect(response).toBeInstanceOf(SuccessResponse);
        expect(response.isOk()).toBeTruthy();
        expect(response.data()).toBeNull();
    });

    it('update roles version error', async () => {
        const d: RolesVersion = {
            id: 'some-id',
            title: 'some-title',
        };
        // @ts-ignore
        request.PATCH = jest.fn().mockResolvedValue({
            status: APIResponseStatus.ERROR,
            data: 'some-error',
        });

        const response = await updateRolesVersion(d);

        expect(request.PATCH).toBeCalledWith('/roles-version', { rolesVersion: d });
        expect(response).toBeInstanceOf(ErrorResponse);
        expect(response.isOk()).toBeFalsy();
        expect(response.data()).toEqual('some-error');
    });

    it('delete roles version success', async () => {
        const d = 'roles-version-id-1';
        // @ts-ignore
        request.DELETE = jest.fn().mockResolvedValue(null);

        const response = await deleteRolesVersion(d);

        expect(request.DELETE).toBeCalledWith(`/roles-version/${d}`);
        expect(response).toBeInstanceOf(SuccessResponse);
        expect(response.isOk()).toBeTruthy();
        expect(response.data()).toBeNull();
    });

    it('delete roles version error', async () => {
        const d = 'roles-version-id-1';
        // @ts-ignore
        request.DELETE = jest.fn().mockResolvedValue({
            status: APIResponseStatus.ERROR,
            data: 'some-error',
        });

        const response = await deleteRolesVersion(d);

        expect(request.DELETE).toBeCalledWith(`/roles-version/${d}`);
        expect(response).toBeInstanceOf(ErrorResponse);
        expect(response.isOk()).toBeFalsy();
        expect(response.data()).toEqual('some-error');
    });
});
