import * as request from '../shared/request';
import {
    login,
    signUp,
    signIn,
    signOut,
} from './account';
import { AccountCredentials } from './account.types';
import { APIResponseStatus, ErrorResponse, SuccessResponse } from '../shared';

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
});
