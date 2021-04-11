import { accountInfoReducer, actionToErrorKey } from '../reducer';
import { ACTIONS } from '../action_types';
import { UnknownErrorCode, UnknownErrorDescription } from '../../shared/const';

describe('account info reducer tests', () => {
    const defaultInfo = {
        login: '',
        apiKey: '',
        created: '',
    };

    it('reduce success fetch info', () => {
        const info = {
            login: 'some-login',
            apiKey: 'some-api-key',
            created: '2021.04.15',
        };

        expect(accountInfoReducer(undefined, {
            type: ACTIONS.SUCCESS_FETCH_INFO,
            accountInfoResult: { info },
        })).toEqual({ info });
    });

    it('reduce failed fetch info', () => {
        const error = {
            code: 'some-code',
            description: 'some-description',
        };

        expect(accountInfoReducer(undefined, {
            type: ACTIONS.FAILED_FETCH_INFO,
            accountInfoResult: {
                fetchInfoError: error,
            },
        })).toEqual({ fetchInfoError: error, info: defaultInfo });
    });

    it('reduce clean fetch info error', () => {
        expect(accountInfoReducer(undefined, {
            type: ACTIONS.CLEAN_FETCH_INFO_ERROR,
        })).toEqual({ fetchInfoError: null, info: defaultInfo });
    });

    it('reduce success update credentials', () => {
        const initialState = {
            info: {
                login: 'some-login',
                apiKey: 'some-api-key',
                created: '2021.04.15',
            },
        };
        const newCredentials = {
            login: 'new-login',
            password: 'new-password',
        };

        expect(accountInfoReducer(initialState as any, {
            type: ACTIONS.SUCCESS_UPDATE_CREDENTIALS,
            accountInfoResult: {
                credentials: newCredentials,
            },
        })).toEqual({
            info: {
                ...initialState.info,
                login: 'new-login'
            },
            credentials: newCredentials,
        });
    });

    it('reduce failed update credentials', () => {
        const error = {
            code: 'some-code',
            description: 'some-description',
        };

        expect(accountInfoReducer(undefined, {
            type: ACTIONS.FAILED_UPDATE_CREDENTIALS,
            accountInfoResult: {
                updateCredentialsError: error,
            },
        })).toEqual({ updateCredentialsError: error, info: defaultInfo });
    });

    it('reduce clean update credentials error', () => {
        expect(accountInfoReducer(undefined, {
            type: ACTIONS.CLEAN_UPDATE_CREDENTIALS_RESULT,
        })).toEqual({ updateCredentialsError: null, credentials: null, info: defaultInfo });
    });

    it('reduce failed to perform action', () => {
        for (const unknownErrorAction of [
            ACTIONS.FAILED_TO_PERFORM_INFO_FETCHING,
            ACTIONS.FAILED_TO_PERFORM_CREDENTIALS_UPDATING,
        ]) {
            expect(accountInfoReducer(undefined, {
                type: unknownErrorAction,
            })).toEqual({
                // @ts-ignore
                [actionToErrorKey[unknownErrorAction]]: {
                    code: UnknownErrorCode,
                    description: UnknownErrorDescription,
                },
                info: defaultInfo,
            });
        }
    });

    it('reduce unknown action', () => {
        expect(accountInfoReducer(undefined, {
            type: 'foo' as any,
        })).toEqual({ info: defaultInfo });
    });
});
