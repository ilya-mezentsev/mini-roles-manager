import { fetchPermissionReducer } from '../reducer';
import { ACTIONS } from '../action_types';
import { Effect } from '../../../services/api/shared/types';
import { UnknownErrorCode, UnknownErrorDescription } from '../../shared/const';

describe('permission reducer tests', () => {
    it('reduce success fetch permission action', () => {
        const fetchPermissionResult = {
            effect: Effect.PERMIT,
        };

        expect(fetchPermissionReducer(undefined, {
            type: ACTIONS.SUCCESS_FETCH_PERMISSION,
            fetchPermissionResult,
        })).toEqual(fetchPermissionResult);
    });

    it('reduce failed fetch permission action', () => {
        const fetchPermissionResult = {
            error: {
                code: 'some-code',
                description: 'some-description',
            },
        };

        expect(fetchPermissionReducer(undefined, {
            type: ACTIONS.FAILED_FETCH_PERMISSION,
            fetchPermissionResult,
        })).toEqual(fetchPermissionResult);
    });

    it('reduce unknown fetch permission error', () => {
        const fetchPermissionResult = {
            error: new Error('foo'),
        };

        expect(fetchPermissionReducer(undefined, {
            type: ACTIONS.FAILED_TO_PERFORM_PERMISSION_FETCHING,
            fetchPermissionResult,
        })).toEqual({
            error: {
                code: UnknownErrorCode,
                description: UnknownErrorDescription,
            }
        });
    });

    it('reduce clean fetch permission result action', () => {
        expect(fetchPermissionReducer(undefined, {
            type: ACTIONS.CLEAN_FETCH_PERMISSION_RESULT,
        })).toEqual({
            effect: null,
            error: null,
        });
    });

    it('reduce unknown action', () => {
        for (const initialState of [undefined, {x: 10}]) {
            expect(fetchPermissionReducer(initialState as any, {
                type: 'foo' as any
            })).toEqual(initialState === undefined ? null : initialState);
        }
    });
});
