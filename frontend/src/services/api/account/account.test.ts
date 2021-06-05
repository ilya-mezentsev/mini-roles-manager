import * as request from '../shared/request';
import {
    fetchInfo,
    fetchPermission,
    login,
    signIn,
    signOut,
    signUp,
    updateCredentials,
} from './account';
import { AccountCredentials, PermissionAccessRequest } from './account.types';
import {
    APIResponseStatus,
    ErrorResponse,
    SuccessResponse,
} from '../shared';
import { Operation } from '../shared/types';

jest.mock('../shared/request');

describe('account api tests', () => {
    it('register account success', async () => {
        const d: AccountCredentials = {login: 'some-login', password: 'some-password'};
        // @ts-ignore
        request.POST = jest.fn().mockResolvedValue(null);

        const response = await signUp(d);

        expect(request.POST).toBeCalledWith('/registration/user', { credentials: d });
        expect(response).toBeInstanceOf(SuccessResponse);
        expect(response.isOk()).toBeTruthy();
        expect(response.data()).toBeNull();
    });

    it('register account error', async () => {
        const d: AccountCredentials = {login: 'some-login', password: 'some-password'};
        // @ts-ignore
        request.POST = jest.fn().mockResolvedValue({
            status: APIResponseStatus.ERROR,
            data: 'some-error',
        });

        const response = await signUp(d);

        expect(request.POST).toBeCalledWith('/registration/user', { credentials: d });
        expect(response).toBeInstanceOf(ErrorResponse);
        expect(response.isOk()).toBeFalsy();
        expect(response.data()).toEqual('some-error');
    });

    it('login empty response', async () => {
        // @ts-ignore
        request.GET = jest.fn().mockResolvedValue(null);

        const response = await login();

        expect(request.GET).toBeCalledWith('/session');
        expect(response).toBeInstanceOf(SuccessResponse);
        expect(response.isOk()).toBeTruthy();
        expect(response.data()).toBeNull();
    });

    it('login success response', async () => {
        // @ts-ignore
        request.GET = jest.fn().mockResolvedValue({
            status: APIResponseStatus.OK,
            data: {
                id: 'some-account-id',
            },
        });

        const response = await login();

        expect(request.GET).toBeCalledWith('/session');
        expect(response).toBeInstanceOf(SuccessResponse);
        expect(response.isOk()).toBeTruthy();
        expect(response.data()).toEqual({
            id: 'some-account-id',
        });
    });

    it('login error response', async () => {
        // @ts-ignore
        request.GET = jest.fn().mockResolvedValue({
            status: APIResponseStatus.ERROR,
            data: 'some-error'
        });

        const response = await login();

        expect(request.GET).toBeCalledWith('/session');
        expect(response).toBeInstanceOf(ErrorResponse);
        expect(response.isOk()).toBeFalsy();
        expect(response.data()).toEqual('some-error');
    });

    it('sign-in success', async () => {
        const d: AccountCredentials = {login: 'some-login', password: 'some-password'};
        // @ts-ignore
        request.POST = jest.fn().mockResolvedValue({
            status: APIResponseStatus.OK,
            data: {
                id: 'some-account-id',
            },
        });

        const response = await signIn(d);

        expect(request.POST).toBeCalledWith('/session', { credentials: d });
        expect(response).toBeInstanceOf(SuccessResponse);
        expect(response.isOk()).toBeTruthy();
        expect(response.data()).toEqual({
            id: 'some-account-id',
        });
    });

    it('sign-in error', async () => {
        const d: AccountCredentials = {login: 'some-login', password: 'some-password'};
        // @ts-ignore
        request.POST = jest.fn().mockResolvedValue({
            status: APIResponseStatus.ERROR,
            data: 'some-error',
        });

        const response = await signIn(d);

        expect(request.POST).toBeCalledWith('/session', { credentials: d });
        expect(response).toBeInstanceOf(ErrorResponse);
        expect(response.isOk()).toBeFalsy();
        expect(response.data()).toEqual('some-error');
    });

    it('sign-out success', async () => {
        // @ts-ignore
        request.DELETE = jest.fn().mockResolvedValue(null);

        const response = await signOut();

        expect(request.DELETE).toBeCalledWith('/session');
        expect(response).toBeInstanceOf(SuccessResponse);
        expect(response.isOk()).toBeTruthy();
        expect(response.data()).toBeNull();
    });

    it('fetch info success', async () => {
        // @ts-ignore
        request.GET = jest.fn().mockResolvedValue({
            status: APIResponseStatus.OK,
            data: 'some-data',
        });

        const response = await fetchInfo();

        expect(request.GET).toBeCalledWith('/account/info');
        expect(response).toBeInstanceOf(SuccessResponse);
        expect(response.isOk()).toBeTruthy();
        expect(response.data()).toEqual('some-data');
    });

    it('update password success', async () => {
        const d: AccountCredentials = {login: 'new-login', password: 'some-password'};
        // @ts-ignore
        request.PATCH = jest.fn().mockResolvedValue(null);

        const response = await updateCredentials(d);

        expect(request.PATCH).toBeCalledWith('/account/credentials', { credentials: d });
        expect(response).toBeInstanceOf(SuccessResponse);
        expect(response.isOk()).toBeTruthy();
        expect(response.data()).toBeNull();
    });

    it('check permissions success', async () => {
        const d: PermissionAccessRequest = {
            roleId: 'role-1',
            operation: Operation.CREATE,
            resourceId: 'resource-1',
        }
        // @ts-ignore
        request.GET = jest.fn().mockResolvedValue({
            status: APIResponseStatus.OK,
            data: 'some-data',
        });

        const response = await fetchPermission(d);

        expect(request.GET).toBeCalledWith('/permissions', d);
        expect(response).toBeInstanceOf(SuccessResponse);
        expect(response.isOk()).toBeTruthy();
        expect(response.data()).toEqual('some-data');
    });
});
