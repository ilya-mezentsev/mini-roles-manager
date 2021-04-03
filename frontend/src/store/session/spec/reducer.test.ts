import { ACTIONS } from '../action_types';
import { sessionReducer } from '../reducer';
import { SessionActionResult } from '../session.types';
import { UnknownErrorCode, UnknownErrorDescription } from '../../shared/const';

describe('session reducer tests', () => {
    it('reduce set session result action', () => {
        const userSession: SessionActionResult = {
            session: {
                id: 'some-id',
            },
        };

        for (const actionType of [
            ACTIONS.SUCCESS_LOGIN,
            ACTIONS.SUCCESS_SIGN_IN,
            ACTIONS.FAILED_SIGN_IN,
            ACTIONS.FAILED_SIGN_OUT,
        ]) {
            expect(sessionReducer(undefined, {
                type: actionType,
                userSession,
            })).toEqual(userSession);
        }
    });

    it('reduce failed login actions', () => {
        const userSession: SessionActionResult = {
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
            })).toEqual({});
        }
    });

    it('reduce success sign-out action', () => {
        const userSession: SessionActionResult = {
            session: {
                id: 'some-id',
            },
        };

        expect(sessionReducer(undefined, {
            type: ACTIONS.SUCCESS_SIGN_OUT,
            userSession,
        })).toEqual({
            session: {
                id: '',
            }
        });
    });

    it('reduce failed to perform sign-in/out action', () => {
        for (const actionType of [
            ACTIONS.FAILED_TO_PERFORM_SIGN_IN_ACTION,
            ACTIONS.FAILED_TO_PERFORM_SIGN_OUT_ACTION,
        ]) {
            for (const error of [
                new Error('foo'),
                null,
            ]) {
                for (const state of [undefined, {x: 10}]) {
                    expect(sessionReducer(state, {
                        type: actionType,
                        userSession: { error },
                    })).toEqual({
                        ...(state || {}),
                        error: {
                            code: UnknownErrorCode,
                            description: UnknownErrorDescription,
                        },
                    });
                }
            }
        }
    });

    it('reduce clean actions', () => {
        for (const actionType of [
            ACTIONS.CLEAN_SIGN_IN_ERROR,
            ACTIONS.CLEAN_SIGN_OUT_ERROR,
        ]) {
            for (const state of [undefined, {x: 10}]) {
                expect(sessionReducer(state, {
                    type: actionType,
                })).toEqual({
                    ...(state || {}),
                    error: null,
                });
            }
        }
    });

    it('reduce unknown action', () => {
        expect(sessionReducer(undefined, {
            type: 'foo',
            userSession: {},
        })).toEqual({});
    });
});
