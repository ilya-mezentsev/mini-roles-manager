import { RoleStore } from './role.store';

import * as api from '../../services/api';
import { ErrorResponse, SuccessResponse } from '../../services/api/shared';
import { UnknownErrorCode, UnknownErrorDescription } from '../shared/const';

jest.mock('../../services/api');

describe('role tests', () => {
    let roleStore = new RoleStore();
    const d = {
        id: 'role-id',
        versionId: 'role-versionId',
        title: 'role-title',
        permissions: [],
        extends: [],
    };

    beforeEach(() => {
        roleStore = new RoleStore();
    });

    it('create role success', async () => {
        // @ts-ignore
        api.createRole = jest.fn().mockResolvedValue(new SuccessResponse(null));

        await roleStore.createRole(d);

        expect(api.createRole).toBeCalledWith(d);
        expect(roleStore.list).toEqual([d]);
    });

    it('create role parsed error', async () => {
        // @ts-ignore
        api.createRole = jest.fn().mockResolvedValue(new ErrorResponse('some-error'));

        await roleStore.createRole(d);

        expect(api.createRole).toBeCalledWith(d);
        expect(roleStore.list).toEqual([]);
        expect(roleStore.createRoleError).toEqual('some-error');
    });

    it('create role unknown error', async () => {
        // @ts-ignore
        api.createRole = jest.fn().mockRejectedValue('some-error');

        await roleStore.createRole(d);

        expect(api.createRole).toBeCalledWith(d);
        expect(roleStore.list).toEqual([]);
        expect(roleStore.createRoleError).toEqual({
            code: UnknownErrorCode,
            description: UnknownErrorDescription,
        });
    });

    it('fetch roles success', async () => {
        // @ts-ignore
        api.rolesList = jest.fn().mockResolvedValue(new SuccessResponse([d]));

        await roleStore.fetchRoles();

        expect(api.rolesList).toBeCalled();
        expect(roleStore.list).toEqual([d]);
    });

    it('fetch empty roles success', async () => {
        // @ts-ignore
        api.rolesList = jest.fn().mockResolvedValue(new SuccessResponse(null));

        await roleStore.fetchRoles();

        expect(api.rolesList).toBeCalled();
        expect(roleStore.list).toEqual([]);
    });

    it('fetch roles parsed error', async () => {
        // @ts-ignore
        api.rolesList = jest.fn().mockResolvedValue(new ErrorResponse('some-error'));

        await roleStore.fetchRoles();

        expect(api.rolesList).toBeCalled();
        expect(roleStore.list).toEqual([]);
        expect(roleStore.fetchRoleError).toEqual('some-error');
    });

    it('fetch roles unknown error', async () => {
        // @ts-ignore
        api.rolesList = jest.fn().mockRejectedValue('some-error');

        await roleStore.fetchRoles();

        expect(api.rolesList).toBeCalled();
        expect(roleStore.list).toEqual([]);
        expect(roleStore.fetchRoleError).toEqual({
            code: UnknownErrorCode,
            description: UnknownErrorDescription,
        });
    });

    it('update role success', async () => {
        // @ts-ignore
        api.updateRole = jest.fn().mockResolvedValue(new SuccessResponse(null));

        // @ts-ignore
        api.createRole = jest.fn().mockResolvedValue(new SuccessResponse(null));
        // add role to store
        await roleStore.createRole(d);

        const t = {
            ...d,
            title: 'another title',
        };

        // update exists role
        await roleStore.updateRole(t);

        expect(api.updateRole).toBeCalledWith(t);
        expect(roleStore.list).toEqual([t]);

        // update not exists role
        const i = {
            ...d,
            id: 'another-id',
        };

        await roleStore.updateRole(i);

        expect(api.updateRole).toBeCalledWith(i);
        expect(roleStore.list).toEqual([t]);
    });

    it('update role parsed error', async () => {
        // @ts-ignore
        api.updateRole = jest.fn().mockResolvedValue(new ErrorResponse('some-error'));

        await roleStore.updateRole(d);

        expect(api.updateRole).toBeCalledWith(d);
        expect(roleStore.list).toEqual([]);
        expect(roleStore.updateRoleError).toEqual('some-error');
    });

    it('update role unknown error', async () => {
        // @ts-ignore
        api.updateRole = jest.fn().mockRejectedValue('some-error');

        await roleStore.updateRole(d);

        expect(api.updateRole).toBeCalledWith(d);
        expect(roleStore.list).toEqual([]);
        expect(roleStore.updateRoleError).toEqual({
            code: UnknownErrorCode,
            description: UnknownErrorDescription,
        });
    });

    it('delete roles success', async () => {
        // @ts-ignore
        api.deleteRole = jest.fn().mockResolvedValue(new SuccessResponse(null));

        // @ts-ignore
        api.createRole = jest.fn().mockResolvedValue(new SuccessResponse(null));
        // add role to store
        await roleStore.createRole(d);
        expect(roleStore.list).toEqual([d]);

        await roleStore.deleteRole(d.versionId, d.id);

        expect(api.deleteRole).toBeCalledWith(d.versionId, d.id);
        expect(roleStore.list).toEqual([]);

        // add role to store (again)
        await roleStore.createRole(d);
        expect(roleStore.list).toEqual([d]);

        await roleStore.deleteRole('not-exists-version-id', d.id);
        expect(roleStore.list).toEqual([d]);

        await roleStore.deleteRole(d.versionId, 'not-exists-role-id');
        expect(roleStore.list).toEqual([d]);
    });

    it('delete role parsed error', async () => {
        // @ts-ignore
        api.deleteRole = jest.fn().mockResolvedValue(new ErrorResponse('some-error'));

        // @ts-ignore
        api.createRole = jest.fn().mockResolvedValue(new SuccessResponse(null));
        // add role to store
        await roleStore.createRole(d);
        expect(roleStore.list).toEqual([d]);

        await roleStore.deleteRole(d.versionId, d.id);

        expect(api.deleteRole).toBeCalledWith(d.versionId, d.id);
        expect(roleStore.list).toEqual([d]);
        expect(roleStore.deleteRoleError).toEqual('some-error');
    });

    it('delete role unknown error', async () => {
        // @ts-ignore
        api.deleteRole = jest.fn().mockRejectedValue('some-error');

        // @ts-ignore
        api.createRole = jest.fn().mockResolvedValue(new SuccessResponse(null));
        // add role to store
        await roleStore.createRole(d);
        expect(roleStore.list).toEqual([d]);

        await roleStore.deleteRole(d.versionId, d.id);

        expect(api.deleteRole).toBeCalledWith(d.versionId, d.id);
        expect(roleStore.list).toEqual([d]);
        expect(roleStore.deleteRoleError).toEqual({
            code: UnknownErrorCode,
            description: UnknownErrorDescription,
        });
    });

    it('clean role actions errors', () => {
        roleStore.createRoleError = 'some-error' as any;
        roleStore.fetchRoleError = 'some-error' as any;
        roleStore.updateRoleError = 'some-error' as any;
        roleStore.deleteRoleError = 'some-error' as any;

        roleStore.cleanRoleActionErrors();

        expect(roleStore.createRoleError).toBeNull();
        expect(roleStore.fetchRoleError).toBeNull();
        expect(roleStore.updateRoleError).toBeNull();
        expect(roleStore.deleteRoleError).toBeNull();
    });
});
