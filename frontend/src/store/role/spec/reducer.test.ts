import { ACTIONS } from '../action_types';
import { actionToErrorType, roleReducer } from '../reducer';
import { Role } from '../../../services/api';
import { UnknownErrorCode, UnknownErrorDescription } from '../../shared/const';

describe('roles reducer tests', () => {
    it('reduce success role creation', () => {
        expect(roleReducer(undefined, {
            type: ACTIONS.SUCCESS_CREATE_ROLE,
            rolesResult: {
                role: {
                    id: 'role-1',
                    versionId: 'version-id',
                },
            },
        })).toEqual({
            list: [
                {
                    id: 'role-1',
                    versionId: 'version-id',
                },
            ],
        });
    });

    it('reduce success roles loading', () => {
        expect(roleReducer(undefined, {
            type: ACTIONS.SUCCESS_FETCH_ROLES,
            rolesResult: {
                list: [
                    {
                        id: 'role-1'
                    } as Role,
                ],
            }
        })).toEqual({
            list: [
                {
                    id: 'role-1'
                },
            ],
        });

        expect(roleReducer(undefined, {
            type: ACTIONS.SUCCESS_FETCH_ROLES,
            rolesResult: {
                list: null,
            }
        })).toEqual({
            list: [],
        });
    });

    it('reduce success role updating', () => {
        const initial = {
            list: [
                {
                    id: 'role-1',
                    title: 'title-1',
                    versionId: 'version-1',
                } as Role,

                {
                    id: 'role-1',
                    title: 'title-1',
                    versionId: 'version-2',
                } as Role,
            ],
        };

        expect(roleReducer(initial, {
            type: ACTIONS.SUCCESS_UPDATE_ROLE,
            rolesResult: {
                role: {
                    id: 'role-1',
                    title: 'title-2',
                    versionId: 'version-1',
                } as Role,
            },
        })).toEqual({
            list: [
                {
                    id: 'role-1',
                    title: 'title-2',
                    versionId: 'version-1',
                },
                {
                    id: 'role-1',
                    title: 'title-1',
                    versionId: 'version-2',
                },
            ],
        });

        expect(roleReducer(initial, {
            type: ACTIONS.SUCCESS_UPDATE_ROLE,
            rolesResult: {
                role: {
                    id: 'role-2',
                    title: 'title-2',
                } as Role,
            },
        })).toEqual({
            list: [
                {
                    id: 'role-1',
                    title: 'title-1',
                    versionId: 'version-1',
                },
                {
                    id: 'role-1',
                    title: 'title-1',
                    versionId: 'version-2',
                },
            ],
        });

        expect(roleReducer(undefined, {
            type: ACTIONS.SUCCESS_UPDATE_ROLE,
            rolesResult: {
                role: {
                    id: 'role-2',
                    title: 'title-2',
                } as Role,
            },
        })).toEqual({
            list: [],
        });
    });

    it('reduce success role deleting', () => {
        const initial = {
            list: [
                {
                    id: 'role-1',
                    title: 'title-1',
                    versionId: 'roles-version-1',
                } as Role,

                {
                    id: 'role-2',
                    title: 'title-1',
                    versionId: 'roles-version-1',
                    extends: ['role-1'],
                } as Role,

                {
                    id: 'role-1',
                    title: 'title-1',
                    versionId: 'roles-version-2',
                } as Role,

                {
                    id: 'role-2',
                    title: 'title-1',
                    versionId: 'roles-version-2',
                    extends: ['role-1'],
                } as Role,
            ],
        };

        expect(roleReducer(initial, {
            type: ACTIONS.SUCCESS_DELETE_ROLE,
            rolesResult: {
                roleId: 'role-1',
                rolesVersionId: 'roles-version-1',
            },
        })).toEqual({
            list: [
                {
                    id: 'role-2',
                    title: 'title-1',
                    versionId: 'roles-version-1',
                    extends: [],
                },
                {
                    id: 'role-1',
                    title: 'title-1',
                    versionId: 'roles-version-2',
                },
                {
                    id: 'role-2',
                    title: 'title-1',
                    versionId: 'roles-version-2',
                    extends: ['role-1'],
                },
            ],
        });

        expect(roleReducer(undefined, {
            type: ACTIONS.SUCCESS_DELETE_ROLE,
            rolesResult: {
                roleId: 'role-1',
                rolesVersionId: 'roles-version-1',
            },
        })).toEqual({
            list: [],
        });
    });

    it('reduce known error of failed actions', () => {
        const error = {
            code: 'some-code',
            description: 'some-description',
        };

        for (const actionType of [
            ACTIONS.FAILED_CREATE_ROLE,
            ACTIONS.FAILED_FETCH_ROLES,
            ACTIONS.FAILED_UPDATE_ROLE,
            ACTIONS.FAILED_DELETE_ROLE,
        ]) {
            expect(roleReducer(undefined, {
                type: actionType,
                rolesResult: { error },
            })).toEqual({
                [actionToErrorType[actionType]]: error,
                list: null,
            });
        }
    });

    it('reduce unknown error of failed actions', () => {
        const error = new Error('some-error');

        for (const actionType of [
            ACTIONS.FAILED_TO_PERFORM_ROLE_CREATION,
            ACTIONS.FAILED_TO_PERFORM_ROLES_FETCHING,
            ACTIONS.FAILED_TO_PERFORM_ROLE_UPDATING,
            ACTIONS.FAILED_TO_PERFORM_ROLE_DELETING,
        ]) {
            expect(roleReducer(undefined, {
                type: actionType,
                rolesResult: { error },
            })).toEqual({
                [actionToErrorType[actionType]]: {
                    code: UnknownErrorCode,
                    description: UnknownErrorDescription,
                },
                list: null,
            });
        }
    });

    it('reduce clean actions', () => {
        for (const actionType of [
            ACTIONS.CLEAN_CREATE_ROLE_ERROR,
            ACTIONS.CLEAN_FETCH_ROLES_ERROR,
            ACTIONS.CLEAN_UPDATE_ROLE_ERROR,
            ACTIONS.CLEAN_DELETE_ROLE_ERROR,
        ]) {
            expect(roleReducer(undefined, {
                type: actionType,
            })).toEqual({
                [actionToErrorType[actionType]]: null,
                list: null,
            });
        }
    });

    it('reduce unknown action', () => {
        expect(roleReducer(undefined, {
            type: 'foo',
        })).toEqual({
            list: null,
        });
    });
});
