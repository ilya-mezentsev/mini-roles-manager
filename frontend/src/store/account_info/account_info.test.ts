import { AccountInfoStore } from './account_info.store';

import * as api from '../../services/api';
import { ErrorResponse, SuccessResponse } from '../../services/api/shared';
import { UnknownErrorCode, UnknownErrorDescription } from '../shared/const';

jest.mock('../../services/api');

describe('account info tests', () => {
    let accountInfoStore = new AccountInfoStore();

    beforeEach(() => {
        accountInfoStore = new AccountInfoStore();
    });

    it('fetch info success', async () => {
        // @ts-ignore
        api.fetchInfo = jest.fn().mockResolvedValue(new SuccessResponse('some-info'));

        await accountInfoStore.fetchInfo();

        expect(api.fetchInfo).toBeCalled();
        expect(accountInfoStore.info).toEqual('some-info');
    });

    it('fetch info parsed error', async () => {
        // @ts-ignore
        api.fetchInfo = jest.fn().mockResolvedValue(new ErrorResponse('some-error'));

        await accountInfoStore.fetchInfo();

        expect(api.fetchInfo).toBeCalled();
        expect(accountInfoStore.info).toBeNull();
        expect(accountInfoStore.fetchInfoError).toEqual('some-error');
    });

    it('fetch info unknown error', async () => {
        // @ts-ignore
        api.fetchInfo = jest.fn().mockRejectedValue('some-error');

        await accountInfoStore.fetchInfo();

        expect(api.fetchInfo).toBeCalled();
        expect(accountInfoStore.info).toBeNull();
        expect(accountInfoStore.fetchInfoError).toEqual({
            code: UnknownErrorCode,
            description: UnknownErrorDescription,
        });
    });

    it('update credentials success', async () => {
        const d = {
            login: 'some-login',
            password: 'some-password',
        };
        // @ts-ignore
        api.updateCredentials = jest.fn().mockResolvedValue(new SuccessResponse(null));

        // @ts-ignore
        api.fetchInfo = jest.fn().mockResolvedValue(new SuccessResponse({
            login: 'foo',
        }));

        // we need to fetch info first to fill account info prop
        await accountInfoStore.fetchInfo();

        expect(accountInfoStore.info?.login).toEqual('foo');

        await accountInfoStore.updateCredentials(d);

        expect(api.updateCredentials).toBeCalledWith(d);
        expect(accountInfoStore.credentials).toEqual(d);
        expect(accountInfoStore.info?.login).toEqual('some-login');
    });

    it('update credentials parsed error', async () => {
        const d = {
            login: 'some-login',
            password: 'some-password',
        };
        // @ts-ignore
        api.updateCredentials = jest.fn().mockResolvedValue(new ErrorResponse('some-error'));

        await accountInfoStore.updateCredentials(d);

        expect(accountInfoStore.credentials).toBeNull();
        expect(accountInfoStore.updateCredentialsError).toEqual('some-error');
    });

    it('update credentials unknown error', async () => {
        const d = {
            login: 'some-login',
            password: 'some-password',
        };
        // @ts-ignore
        api.updateCredentials = jest.fn().mockRejectedValue('some-error');

        await accountInfoStore.updateCredentials(d);

        expect(accountInfoStore.credentials).toBeNull();
        expect(accountInfoStore.updateCredentialsError).toEqual({
            code: UnknownErrorCode,
            description: UnknownErrorDescription,
        });
    });

    it('clean account info actions errors', () => {
        accountInfoStore.fetchInfoError = 'foo' as any;
        accountInfoStore.updateCredentialsError = 'foo' as any;

        accountInfoStore.cleanAccountInfoActionErrors();

        expect(accountInfoStore.fetchInfoError).toBeNull();
        expect(accountInfoStore.updateCredentialsError).toBeNull();
    });
});
