import { ACTIONS } from '../action_types';
import { sessionReducer } from '../reducer';
import { SessionResult } from '../session.types';
import { UnknownErrorCode, UnknownErrorDescription } from '../../shared/const';

describe('session reducer tests', () => {
    it('reduce set session result action', () => {
        const userSession: SessionResult = {
            session: {
                id: 'some-id',
            },
        };

        for (const actionType of [
            ACTIONS.SUCCESS_LOGIN,
            ACTIONS.SUCCESS_SIGN_IN,
            ACTIONS.FAILED_SIGN_IN,
        ]) {
            expect(sessionReducer(undefined, {
                type: actionType,
                userSession,
            })).toEqual(userSession);
        }
    });

    it('reduce failed login actions', () => {
        const userSession: SessionResult = {
            session: {
                id: 'some-id',
            },
        };

        for (const actionType of [
            ACTIONS.FAILED_LOGIN,
            ACTIONS.FAILED_TO_PERFORM_LOGIN_ACTION,
        ]) {
            expect(sessionReducer(undefined, {
                type: actionType,
                userSession,
            })).toBeNull();
        }
    });

    it('reduce failed to perform sign-in action', () => {
        const userSession: SessionResult = {
            error: new Error('foo'),
        };

        expect(sessionReducer(undefined, {
            type: ACTIONS.FAILED_TO_PERFORM_SIGN_IN_ACTION,
            userSession,
        })).toEqual({
            error: {
                code: UnknownErrorCode,
                description: UnknownErrorDescription,
            },
        });

        expect(sessionReducer(undefined, {
            type: ACTIONS.FAILED_TO_PERFORM_SIGN_IN_ACTION,
            userSession: { error: null as any },
        })).toEqual({
            error: {
                code: UnknownErrorCode,
                description: UnknownErrorDescription,
            },
        });
    });

    it('reduce unknown action', () => {
        expect(sessionReducer(undefined, {
            type: 'foo',
            userSession: {},
        })).toBeNull();
    });
});
