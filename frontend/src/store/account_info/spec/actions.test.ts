import configureMockStore from 'redux-mock-store';
import thunk from 'redux-thunk';

import * as api from '../../../services/api';
import { AccountCredentials } from '../../../services/api';
import { SuccessResponse, ErrorResponse } from '../../../services/api/shared';
import {
    fetchInfo,
    cleanFetchInfoError,
    updateCredentials,
    cleanUpdateCredentialsResult,
} from '../actions';
import { ACTIONS } from '../action_types';

jest.mock('../../../services/api');

describe('account info actions tests', () => {
    const middlewares = [thunk];
    const mockStore = configureMockStore(middlewares);

    it('fetch info success', async () => {
        const store = mockStore({ accountInfoResult: null });
        // @ts-ignore
        api.fetchInfo = jest.fn().mockResolvedValue(new SuccessResponse('some-data'));

        // @ts-ignore
        await store.dispatch(fetchInfo());

        expect(api.fetchInfo).toBeCalled();
        expect(store.getActions()).toEqual([
            {
                type: ACTIONS.SUCCESS_FETCH_INFO,
                accountInfoResult: {
                    info: 'some-data'
                },
            },
        ]);
    });

    it('fetch info parsed error', async () => {
        const store = mockStore({ accountInfoResult: null });
        // @ts-ignore
        api.fetchInfo = jest.fn().mockResolvedValue(new ErrorResponse('some-error'));

        // @ts-ignore
        await store.dispatch(fetchInfo());

        expect(api.fetchInfo).toBeCalled();
        expect(store.getActions()).toEqual([
            {
                type: ACTIONS.FAILED_FETCH_INFO,
                accountInfoResult: {
                    fetchInfoError: 'some-error'
                },
            },
        ]);
    });

    it('fetch info unknown error', async () => {
        const store = mockStore({ accountInfoResult: null });
        // @ts-ignore
        api.fetchInfo = jest.fn().mockRejectedValue('some-error');

        // @ts-ignore
        await store.dispatch(fetchInfo());

        expect(api.fetchInfo).toBeCalled();
        expect(store.getActions()).toEqual([
            {
                type: ACTIONS.FAILED_TO_PERFORM_INFO_FETCHING,
                accountInfoResult: {
                    fetchInfoError: 'some-error'
                },
            },
        ]);
    });

    it('clean fetch info error', () => {
        const store = mockStore({ accountInfoResult: null });

        // @ts-ignore
        store.dispatch(cleanFetchInfoError());

        expect(store.getActions()).toEqual([
            {
                type: ACTIONS.CLEAN_FETCH_INFO_ERROR,
            },
        ]);
    });

    it('update credentials success', async () => {
        const d: AccountCredentials = {
            login: 'some-login',
            password: 'some-password',
        };
        const store = mockStore({ accountInfoResult: null });
        // @ts-ignore
        api.updateCredentials = jest.fn().mockResolvedValue(new SuccessResponse(null));

        // @ts-ignore
        await store.dispatch(updateCredentials(d));

        expect(api.updateCredentials).toBeCalledWith(d);
        expect(store.getActions()).toEqual([
            {
                type: ACTIONS.SUCCESS_UPDATE_CREDENTIALS,
                accountInfoResult: {
                    credentials: d,
                },
            },
        ])
    });

    it('update credentials parsed error', async () => {
        const d: AccountCredentials = {
            login: 'some-login',
            password: 'some-password',
        };
        const store = mockStore({ accountInfoResult: null });
        // @ts-ignore
        api.updateCredentials = jest.fn().mockResolvedValue(new ErrorResponse('some-error'));

        // @ts-ignore
        await store.dispatch(updateCredentials(d));

        expect(api.updateCredentials).toBeCalledWith(d);
        expect(store.getActions()).toEqual([
            {
                type: ACTIONS.FAILED_UPDATE_CREDENTIALS,
                accountInfoResult: {
                    updateCredentialsError: 'some-error',
                },
            },
        ]);
    });

    it('update credentials unknown error', async () => {
        const d: AccountCredentials = {
            login: 'some-login',
            password: 'some-password',
        };
        const store = mockStore({ accountInfoResult: null });
        // @ts-ignore
        api.updateCredentials = jest.fn().mockRejectedValue('some-error');

        // @ts-ignore
        await store.dispatch(updateCredentials(d));

        expect(api.updateCredentials).toBeCalledWith(d);
        expect(store.getActions()).toEqual([
            {
                type: ACTIONS.FAILED_TO_PERFORM_CREDENTIALS_UPDATING,
                accountInfoResult: {
                    updateCredentialsError: 'some-error',
                },
            },
        ]);
    });

    it('clean update credentials result', () => {
        const store = mockStore({ accountInfoResult: null });

        // @ts-ignore
        store.dispatch(cleanUpdateCredentialsResult());

        expect(store.getActions()).toEqual([
            {
                type: ACTIONS.CLEAN_UPDATE_CREDENTIALS_RESULT,
            },
        ]);
    });
});
