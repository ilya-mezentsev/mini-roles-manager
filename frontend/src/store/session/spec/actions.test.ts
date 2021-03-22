import configureMockStore from 'redux-mock-store';
import thunk from 'redux-thunk';

import * as api from '../../../services/api';
import { SuccessResponse, ErrorResponse } from '../../../services/api/shared';
import { AccountCredentials } from '../../../services/api';
import { login, signIn, cleanSignIn } from '../actions';
import { ACTIONS } from '../action_types';

jest.mock('../../../services/api');

describe('session actions tests', () => {
    const middlewares = [thunk];
    const mockStore = configureMockStore(middlewares);

    it('sign-in success', async () => {
        const d: AccountCredentials = { login: 'login', password: 'password' };
        const store = mockStore({ userSession: null });
        // @ts-ignore
        api.signIn = jest.fn().mockResolvedValue(new SuccessResponse('some-data'));

        // @ts-ignore
        await store.dispatch(signIn(d));

        expect(api.signIn).toBeCalledWith(d);
        expect(store.getActions()).toEqual([
            {
                type: ACTIONS.SUCCESS_SIGN_IN,
                userSession: {
                    session: 'some-data',
                },
            },
        ]);
    });

    it('sign-in parsed error', async () => {
        const d: AccountCredentials = { login: 'login', password: 'password' };
        const store = mockStore({ userSession: null });
        // @ts-ignore
        api.signIn = jest.fn().mockResolvedValue(new ErrorResponse('some-error'));

        // @ts-ignore
        await store.dispatch(signIn(d));

        expect(api.signIn).toBeCalledWith(d);
        expect(store.getActions()).toEqual([
            {
                type: ACTIONS.FAILED_SIGN_IN,
                userSession: {
                    error: 'some-error',
                },
            },
        ]);
    });

    it('sign-in unknown error', async () => {
        const d: AccountCredentials = { login: 'login', password: 'password' };
        const store = mockStore({ userSession: null });
        // @ts-ignore
        api.signIn = jest.fn().mockRejectedValue('some-error');

        // @ts-ignore
        await store.dispatch(signIn(d));

        expect(api.signIn).toBeCalledWith(d);
        expect(store.getActions()).toEqual([
            {
                type: ACTIONS.FAILED_TO_PERFORM_SIGN_IN_ACTION,
                userSession: {
                    error: 'some-error',
                },
            },
        ]);
    });

    it('clean sign-in', () => {
        const store = mockStore({ userSession: null });

        // @ts-ignore
        store.dispatch(cleanSignIn());

        expect(store.getActions()).toEqual([
            {
                type: ACTIONS.CLEAN_SIGN_IN,
            },
        ]);
    });

    it('login success', async () => {
        const store = mockStore({ userSession: null });
        // @ts-ignore
        api.login = jest.fn().mockResolvedValue(new SuccessResponse('some-data'));

        // @ts-ignore
        await store.dispatch(login());

        expect(api.login).toBeCalled();
        expect(store.getActions()).toEqual([
            {
                type: ACTIONS.SUCCESS_LOGIN,
                userSession: {
                    session: 'some-data',
                },
            },
        ]);
    });

    it('login parsed error', async () => {
        const store = mockStore({ userSession: null });
        // @ts-ignore
        api.login = jest.fn().mockResolvedValue(new ErrorResponse('some-error'));

        // @ts-ignore
        await store.dispatch(login());

        expect(api.login).toBeCalled();
        expect(store.getActions()).toEqual([
            {
                type: ACTIONS.FAILED_LOGIN,
                userSession: {
                    error: 'some-error',
                },
            },
        ]);
    });

    it('login unknown error', async () => {
        const store = mockStore({ userSession: null });
        // @ts-ignore
        api.login = jest.fn().mockRejectedValue('some-error');

        // @ts-ignore
        await store.dispatch(login());

        expect(api.login).toBeCalled();
        expect(store.getActions()).toEqual([
            {
                type: ACTIONS.FAILED_TO_PERFORM_LOGIN_ACTION,
                userSession: {
                    error: 'some-error',
                },
            },
        ]);
    });
});
