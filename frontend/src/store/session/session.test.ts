import { SessionStore } from './session.store';

import * as api from '../../services/api';
import { ErrorResponse, SuccessResponse } from '../../services/api/shared';
import { UnknownErrorCode, UnknownErrorDescription } from '../shared/const';

jest.mock('../../services/api');

describe('session tests', () => {
    let sessionStore = new SessionStore();
    const d = {
        login: 'some-login',
        password: 'some-password',
    };

    beforeEach(() => {
        sessionStore = new SessionStore();
    });

    it('sign-in success', async () => {
        // @ts-ignore
        api.signIn = jest.fn().mockResolvedValue(new SuccessResponse('some-session'));

        await sessionStore.signIn(d);

        expect(api.signIn).toBeCalledWith(d);
        expect(sessionStore.session).toEqual('some-session');
    });

    it('sign-in parsed error', async () => {
        // @ts-ignore
        api.signIn = jest.fn().mockResolvedValue(new ErrorResponse('some-error'));

        await sessionStore.signIn(d);

        expect(api.signIn).toBeCalledWith(d);
        expect(sessionStore.session).toBeNull();
        expect(sessionStore.signInError).toEqual('some-error');
    });

    it('sign-in unknown error', async () => {
        // @ts-ignore
        api.signIn = jest.fn().mockRejectedValue('some-error');

        await sessionStore.signIn(d);

        expect(api.signIn).toBeCalledWith(d);
        expect(sessionStore.session).toBeNull();
        expect(sessionStore.signInError).toEqual({
            code: UnknownErrorCode,
            description: UnknownErrorDescription,
        });
    });

    it('sign-out success', async () => {
        // @ts-ignore
        api.signOut = jest.fn().mockResolvedValue(new SuccessResponse(null));
        sessionStore.session = 'some-session' as any;

        await sessionStore.signOut();

        expect(api.signOut).toBeCalled();
        expect(sessionStore.session).toBeNull();
    });

    it('sign-out parsed error', async () => {
        // @ts-ignore
        api.signOut = jest.fn().mockResolvedValue(new ErrorResponse('some-error'));
        sessionStore.session = 'some-session' as any;

        await sessionStore.signOut();

        expect(api.signOut).toBeCalled();
        expect(sessionStore.session).toEqual('some-session');
        expect(sessionStore.signOutError).toEqual('some-error');
    });

    it('sign-out unknown error', async () => {
        // @ts-ignore
        api.signOut = jest.fn().mockRejectedValue('some-error');
        sessionStore.session = 'some-session' as any;

        await sessionStore.signOut();

        expect(api.signOut).toBeCalled();
        expect(sessionStore.session).toEqual('some-session');
        expect(sessionStore.signOutError).toEqual({
            code: UnknownErrorCode,
            description: UnknownErrorDescription,
        });
    });

    it('login success', async () => {
        // @ts-ignore
        api.login = jest.fn().mockResolvedValue(new SuccessResponse('some-session'));

        await sessionStore.login();

        expect(api.login).toBeCalled();
        expect(sessionStore.session).toEqual('some-session');
    });

    it('login parsed error', async () => {
        // @ts-ignore
        api.login = jest.fn().mockResolvedValue(new ErrorResponse('some-error'));

        await sessionStore.login();

        expect(api.login).toBeCalled();
        expect(sessionStore.session).toBeNull();
        expect(sessionStore.loginError).toEqual('some-error');
    });

    it('login unknown error', async () => {
        // @ts-ignore
        api.login = jest.fn().mockRejectedValue('some-error');

        await sessionStore.login();

        expect(api.login).toBeCalled();
        expect(sessionStore.session).toBeNull();
        expect(sessionStore.loginError).toEqual({
            code: UnknownErrorCode,
            description: UnknownErrorDescription,
        });
    });

    it('clean session actions errors', () => {
        sessionStore.loginError = 'some-error' as any;
        sessionStore.signOutError = 'some-error' as any;
        sessionStore.signInError = 'some-error' as any;

        sessionStore.cleanSessionActionErrors();

        expect(sessionStore.loginError).toBeNull();
        expect(sessionStore.signOutError).toBeNull();
        expect(sessionStore.signInError).toBeNull();
    });
});
