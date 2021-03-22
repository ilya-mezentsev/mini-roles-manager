import configureMockStore from 'redux-mock-store';
import thunk from 'redux-thunk';

import * as api from '../../../services/api';
import { SuccessResponse, ErrorResponse } from '../../../services/api/shared';
import { AccountCredentials } from '../../../services/api';
import { signUp, cleanSignUp } from '../actions';
import { ACTIONS } from '../action_types';

jest.mock('../../../services/api');

describe('registration actions tests', () => {
    const middlewares = [thunk];
    const mockStore = configureMockStore(middlewares);

    it('sign-up success', async () => {
        const d: AccountCredentials = {login: 'login', password: 'password'};
        const store = mockStore({ registrationResult: null });
        // @ts-ignore
        api.signUp = jest.fn().mockResolvedValue(new SuccessResponse(null));

        // @ts-ignore
        await store.dispatch(signUp(d));

        expect(api.signUp).toBeCalledWith(d);
        expect(store.getActions()).toEqual([
            {
                type: ACTIONS.SUCCESS_REGISTRATION,
                registrationResult: {
                    failed: false,
                },
            },
        ]);
    });

    it('sign-up parsed error', async () => {
        const d: AccountCredentials = {login: 'login', password: 'password'};
        const store = mockStore({ registrationResult: null });
        // @ts-ignore
        api.signUp = jest.fn().mockResolvedValue(new ErrorResponse('some-error'));

        // @ts-ignore
        await store.dispatch(signUp(d));

        expect(api.signUp).toBeCalledWith(d);
        expect(store.getActions()).toEqual([
            {
                type: ACTIONS.FAILED_TO_REGISTER_USER,
                registrationResult: {
                    failed: true,
                    error: 'some-error',
                },
            },
        ]);
    });

    it('sign-up unknown error', async () => {
        const d: AccountCredentials = {login: 'login', password: 'password'};
        const store = mockStore({ registrationResult: null });
        // @ts-ignore
        api.signUp = jest.fn().mockRejectedValue('some-error');

        // @ts-ignore
        await store.dispatch(signUp(d));

        expect(api.signUp).toBeCalledWith(d);
        expect(store.getActions()).toEqual([
            {
                type: ACTIONS.FAILED_TO_PERFORM_REGISTER_ACTION,
                registrationResult: {
                    failed: true,
                    error: 'some-error',
                },
            },
        ]);
    });

    it('clean sign-up', () => {
        const store = mockStore({ registrationResult: null });

        // @ts-ignore
        store.dispatch(cleanSignUp());

        expect(store.getActions()).toEqual([
            {
                type: ACTIONS.CLEAN_REGISTRATION,
            },
        ]);
    });
});
