import { ACTIONS } from '../action_types';
import { actionToErrorType, rolesVersionReducer } from '../reducer';
import { RolesVersion } from '../../../services/api';
import { UnknownErrorCode, UnknownErrorDescription } from '../../shared/const';

describe('roles version reducer tests', () => {
    it('reduce success roles version creation', () => {
        expect(rolesVersionReducer(undefined, {
            type: ACTIONS.SUCCESS_CREATE_ROLES_VERSION,
            rolesVersionResult: {
                rolesVersion: {
                    id: 'roles-version-id',
                    title: 'roles-version-title',
                },
            },
        })).toEqual({
            currentRolesVersion: null,
            list: [
                {
                    id: 'roles-version-id',
                    title: 'roles-version-title',
                },
            ],
        });
    });

    it('reduce success fetch roles versions', () => {
        const d: RolesVersion[] = [
            { id: 'id-1', title: 'title-1' },
            { id: 'id-2', title: 'title-2' },
        ];

        expect(rolesVersionReducer(undefined, {
            type: ACTIONS.SUCCESS_FETCH_ROLES_VERSION,
            rolesVersionResult: {
                list: d,
            },
        })).toEqual({
            currentRolesVersion: d[0],
            list: d,
        });

        expect(rolesVersionReducer(undefined, {
            type: ACTIONS.SUCCESS_FETCH_ROLES_VERSION,
            rolesVersionResult: {
                list: null,
            },
        })).toEqual({
            currentRolesVersion: null,
            list: [],
        });
    });

    it('reduce success update exists roles version', () => {
        const d: RolesVersion[] = [
            { id: 'id-1', title: 'title-1' },
            { id: 'id-2', title: 'title-2' },
        ];
        const updated: RolesVersion = {
            ...d[1],
            title: 'updated-title',
        };

        expect(rolesVersionReducer({ list: d, currentRolesVersion: d[0] }, {
            type: ACTIONS.SUCCESS_UPDATE_ROLES_VERSION,
            rolesVersionResult: {
                rolesVersion: updated,
            },
        })).toEqual({
            currentRolesVersion: d[0],
            list: [
                d[0],
                updated,
            ],
        });
    });

    it('reduce success update not exists roles version', () => {
        const d: RolesVersion[] = [
            { id: 'id-1', title: 'title-1' },
            { id: 'id-2', title: 'title-2' },
        ];
        const updated: RolesVersion = {
            ...d[1],
            title: 'updated-title',
        };

        expect(rolesVersionReducer(undefined, {
            type: ACTIONS.SUCCESS_UPDATE_ROLES_VERSION,
            rolesVersionResult: {
                rolesVersion: updated,
            },
        })).toEqual({
            currentRolesVersion: null,
            list: [ ],
        });
    });

    it('reduce success delete roles version', () => {
        const d: RolesVersion[] = [
            { id: 'id-1', title: 'title-1' },
            { id: 'id-2', title: 'title-2' },
        ];
        const deletedId = 'id-1';

        expect(rolesVersionReducer({ list: d, currentRolesVersion: null }, {
            type: ACTIONS.SUCCESS_DELETE_ROLES_VERSION,
            rolesVersionResult: {
                rolesVersionId: deletedId,
            },
        })).toEqual({
            currentRolesVersion: null,
            list: d.filter(r => r.id !== deletedId),
        });

        expect(rolesVersionReducer(undefined, {
            type: ACTIONS.SUCCESS_DELETE_ROLES_VERSION,
            rolesVersionResult: {
                rolesVersionId: deletedId,
            },
        })).toEqual({
            currentRolesVersion: null,
            list: [],
        });
    });

    it('reduce success delete roles version (current version)', () => {
        const d: RolesVersion[] = [
            { id: 'id-1', title: 'title-1' },
            { id: 'id-2', title: 'title-2' },
        ];
        const deletedId = 'id-1';

        expect(rolesVersionReducer({ list: d, currentRolesVersion: d.find(r => r.id === deletedId)! }, {
            type: ACTIONS.SUCCESS_DELETE_ROLES_VERSION,
            rolesVersionResult: {
                rolesVersionId: deletedId,
            },
        })).toEqual({
            currentRolesVersion: { id: 'id-2', title: 'title-2' },
            list: d.filter(r => r.id !== deletedId),
        });
    });

    it('reduce select roles version', () => {
        const d: RolesVersion = { id: 'id-1', title: 'title-1' };

        expect(rolesVersionReducer(undefined, {
            type: ACTIONS.SELECT_CURRENT_ROLES_VERSION,
            rolesVersionResult: {
                rolesVersion: d,
            },
        })).toEqual({
            currentRolesVersion: d,
            list: null,
        });
    });

    it('reduce known error of failed actions', () => {
        const error = {
            code: 'some-code',
            description: 'some-description',
        };

        for (const actionType of [
            ACTIONS.FAILED_CREATE_ROLES_VERSION,
            ACTIONS.FAILED_FETCH_ROLES_VERSION,
            ACTIONS.FAILED_UPDATE_ROLES_VERSION,
            ACTIONS.FAILED_DELETE_ROLES_VERSION,
        ]) {
            expect(rolesVersionReducer(undefined, {
                type: actionType,
                rolesVersionResult: { error },
            })).toEqual({
                [actionToErrorType[actionType]]: error,
                list: null,
                currentRolesVersion: null,
            });
        }
    });

    it('reduce unknown error of failed actions', () => {
        const error = new Error('some-error');

        for (const actionType of [
            ACTIONS.FAILED_TO_PERFORM_ROLES_VERSION_CREATION,
            ACTIONS.FAILED_TO_PERFORM_ROLES_VERSION_FETCHING,
            ACTIONS.FAILED_TO_PERFORM_ROLES_VERSION_UPDATING,
            ACTIONS.FAILED_TO_PERFORM_ROLES_VERSION_DELETING,
        ]) {
            expect(rolesVersionReducer(undefined, {
                type: actionType,
                rolesVersionResult: { error },
            })).toEqual({
                [actionToErrorType[actionType]]: {
                    code: UnknownErrorCode,
                    description: UnknownErrorDescription,
                },
                list: null,
                currentRolesVersion: null,
            });
        }
    });

    it('reduce clean actions', () => {
        for (const actionType of [
            ACTIONS.CLEAN_CREATE_ROLES_VERSION_ERROR,
            ACTIONS.CLEAN_FETCH_ROLES_VERSION_ERROR,
            ACTIONS.CLEAN_UPDATE_ROLES_VERSION_ERROR,
            ACTIONS.CLEAN_DELETE_ROLES_VERSION_ERROR,
        ]) {
            expect(rolesVersionReducer(undefined, {
                type: actionType,
            })).toEqual({
                [actionToErrorType[actionType]]: null,
                list: null,
                currentRolesVersion: null,
            });
        }
    });

    it('reduce unknown action', () => {
        expect(rolesVersionReducer(undefined, {
            type: 'foo',
        })).toEqual({
            list: null,
            currentRolesVersion: null,
        });
    });
});
