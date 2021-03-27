
import { ACTIONS } from '../action_types';
import { registrationReducer } from '../reducer';
import { RegistrationResult } from '../registration.types';
import { UnknownErrorCode, UnknownErrorDescription } from '../../shared/const';

describe('registration reducer tests', () => {
    it('reduce success registration action', () => {
        const registrationResult: RegistrationResult = { error: undefined };

        expect(registrationReducer(undefined, {
            type: ACTIONS.SUCCESS_REGISTRATION,
            registrationResult,
        })).toEqual(registrationResult);
    });

    it('reduce error registration action', () => {
        const registrationResult: RegistrationResult = {
            error: {
                code: 'some-code',
                description: 'Some description'
            }
        };

        expect(registrationReducer(undefined, {
            type: ACTIONS.FAILED_TO_REGISTER_USER,
            registrationResult,
        })).toEqual(registrationResult);
    });

    it('reduce unknown registration error', () => {
        expect(registrationReducer(undefined, {
            type: ACTIONS.FAILED_TO_PERFORM_REGISTER_ACTION,
            registrationResult: {
                error: new Error('foo'),
            },
        })).toEqual({
            error: {
                code: UnknownErrorCode,
                description: UnknownErrorDescription,
            },
        });

        expect(registrationReducer(undefined, {
            type: ACTIONS.FAILED_TO_PERFORM_REGISTER_ACTION,
            registrationResult: {
                // hack for coverage
                error: null as unknown as Error,
            },
        })).toEqual({
            error: {
                code: UnknownErrorCode,
                description: UnknownErrorDescription,
            },
        });
    });

    it('reduce clean registration action', () => {
        expect(registrationReducer(undefined, {
            type: ACTIONS.CLEAN_REGISTRATION,
        })).toBeNull();
    });

    it('reduce unknown action', () => {
        expect(registrationReducer(undefined, {
            type: 'bar',
        })).toBeNull();
    });
});
