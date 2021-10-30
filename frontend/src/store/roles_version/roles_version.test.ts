import { RolesVersionStore } from './roles_version.store';

import * as api from '../../services/api';
import { ErrorResponse, SuccessResponse } from '../../services/api/shared';
import { UnknownErrorCode, UnknownErrorDescription } from '../shared/const';
import { RoleStore } from '../role/role.store';

jest.mock('../../services/api');

describe('roles version tests', () => {
    let roleStore = new RoleStore();
    let rolesVersionStore = new RolesVersionStore(roleStore);
    const d = {
        id: 'roles-version-id-1',
        title: 'roles-version-title-1',
    };
    const t = {
        id: 'roles-version-id-2',
        title: 'roles-version-title-2',
    };

    beforeEach(() => {
        roleStore = new RoleStore();
        rolesVersionStore = new RolesVersionStore(roleStore);
    });

    it('create roles version success', async () => {
        // @ts-ignore
        api.createRolesVersion = jest.fn().mockResolvedValue(new SuccessResponse(null));

        await rolesVersionStore.createRolesVersion(d);

        expect(api.createRolesVersion).toBeCalledWith(d);
        expect(rolesVersionStore.list).toEqual([d]);
    });

    it('create roles version parsed error', async () => {
        // @ts-ignore
        api.createRolesVersion = jest.fn().mockResolvedValue(new ErrorResponse('some-error'));

        await rolesVersionStore.createRolesVersion(d);

        expect(api.createRolesVersion).toBeCalledWith(d);
        expect(rolesVersionStore.list).toEqual([]);
        expect(rolesVersionStore.createRolesVersionError).toEqual('some-error');
    });

    it('create roles version unknown error', async () => {
        // @ts-ignore
        api.createRolesVersion = jest.fn().mockRejectedValue('some-error');

        await rolesVersionStore.createRolesVersion(d);

        expect(api.createRolesVersion).toBeCalledWith(d);
        expect(rolesVersionStore.list).toEqual([]);
        expect(rolesVersionStore.createRolesVersionError).toEqual({
            code: UnknownErrorCode,
            description: UnknownErrorDescription,
        });
    });

    it('fetch roles versions success', async () => {
        // @ts-ignore
        api.rolesVersionsList = jest.fn().mockResolvedValue(new SuccessResponse([d, t]));

        await rolesVersionStore.fetchRolesVersions();

        expect(api.rolesVersionsList).toBeCalled();
        expect(rolesVersionStore.list).toEqual([d, t]);
        expect(rolesVersionStore.current).toEqual(d);
    });

    it('fetch roles versions parsed error', async () => {
        // @ts-ignore
        api.rolesVersionsList = jest.fn().mockResolvedValue(new ErrorResponse('some-error'));

        await rolesVersionStore.fetchRolesVersions();

        expect(api.rolesVersionsList).toBeCalled();
        expect(rolesVersionStore.list).toEqual([]);
        expect(rolesVersionStore.current).toEqual(null);
        expect(rolesVersionStore.fetchRolesVersionError).toEqual('some-error');
    });

    it('fetch roles versions unknown error', async () => {
        // @ts-ignore
        api.rolesVersionsList = jest.fn().mockRejectedValue('some-error');

        await rolesVersionStore.fetchRolesVersions();

        expect(api.rolesVersionsList).toBeCalled();
        expect(rolesVersionStore.list).toEqual([]);
        expect(rolesVersionStore.current).toEqual(null);
        expect(rolesVersionStore.fetchRolesVersionError).toEqual({
            code: UnknownErrorCode,
            description: UnknownErrorDescription,
        });
    });

    it('update roles version success', async () => {
        // @ts-ignore
        api.updateRolesVersion = jest.fn().mockResolvedValue(new SuccessResponse(null));

        // @ts-ignore
        api.rolesVersionsList = jest.fn().mockResolvedValue(new SuccessResponse([d, t]));

        // add roles versions to store
        await rolesVersionStore.fetchRolesVersions();
        expect(rolesVersionStore.current).toEqual(d);

        const i = {
            ...d,
            title: 'some-new-title',
        };

        await rolesVersionStore.updateRolesVersion(i);

        expect(api.updateRolesVersion).toBeCalledWith(i);
        expect(rolesVersionStore.list).toEqual([i, t]);
        expect(rolesVersionStore.current).toEqual(i);

        const j = {
            ...t,
            title: 'another-title',
        };

        await rolesVersionStore.updateRolesVersion(j);

        expect(api.updateRolesVersion).toBeCalledWith(j);
        expect(rolesVersionStore.list).toEqual([i, j]);
        expect(rolesVersionStore.current).toEqual(i);

        const n = {
            id: 'not-exists-id',
            title: 'some-title',
        };

        await rolesVersionStore.updateRolesVersion(n);

        expect(api.updateRolesVersion).toBeCalledWith(n);
        expect(rolesVersionStore.list).toEqual([i, j]);
        expect(rolesVersionStore.current).toEqual(i);
    });

    it('update roles version parsed error', async () => {
        // @ts-ignore
        api.updateRolesVersion = jest.fn().mockResolvedValue(new ErrorResponse('some-error'));

        await rolesVersionStore.updateRolesVersion(d);

        expect(api.updateRolesVersion).toBeCalledWith(d);
        expect(rolesVersionStore.updateRolesVersionError).toEqual('some-error');
    });

    it('update roles version unknown error', async () => {
        // @ts-ignore
        api.updateRolesVersion = jest.fn().mockRejectedValue('some-error');

        await rolesVersionStore.updateRolesVersion(d);

        expect(api.updateRolesVersion).toBeCalledWith(d);
        expect(rolesVersionStore.updateRolesVersionError).toEqual({
            code: UnknownErrorCode,
            description: UnknownErrorDescription,
        });
    });

    it('delete roles version success', async () => {
        // @ts-ignore
        api.deleteRolesVersion = jest.fn().mockResolvedValue(new SuccessResponse(null));

        // @ts-ignore
        api.rolesVersionsList = jest.fn().mockResolvedValue(new SuccessResponse([d, t]));

        // add roles versions to store
        await rolesVersionStore.fetchRolesVersions();
        expect(rolesVersionStore.current).toEqual(d);

        await rolesVersionStore.deleteRolesVersion(d.id);

        expect(api.rolesList).toBeCalled();
        expect(api.deleteRolesVersion).toBeCalledWith(d.id);
        expect(rolesVersionStore.list).toEqual([t]);
        expect(rolesVersionStore.current).toEqual(t);

        // add roles versions to store
        await rolesVersionStore.fetchRolesVersions();
        expect(rolesVersionStore.current).toEqual(d);

        await rolesVersionStore.deleteRolesVersion(t.id);

        expect(api.rolesList).toBeCalled();
        expect(api.deleteRolesVersion).toBeCalledWith(t.id);
        expect(rolesVersionStore.list).toEqual([d]);
        expect(rolesVersionStore.current).toEqual(d);

        // add roles versions to store
        await rolesVersionStore.fetchRolesVersions();
        expect(rolesVersionStore.current).toEqual(d);

        const n = 'not-exists-id';

        await rolesVersionStore.deleteRolesVersion(n);

        expect(api.rolesList).toBeCalled();
        expect(api.deleteRolesVersion).toBeCalledWith(n);
        expect(rolesVersionStore.list).toEqual([d, t]);
    });

    it('delete roles version parsed error', async () => {
        // @ts-ignore
        api.deleteRolesVersion = jest.fn().mockResolvedValue(new ErrorResponse('some-error'));

        await rolesVersionStore.deleteRolesVersion(d.id);

        expect(api.deleteRolesVersion).toBeCalledWith(d.id);
        expect(rolesVersionStore.deleteRolesVersionError).toEqual('some-error');
    });

    it('delete roles version unknown error', async () => {
        // @ts-ignore
        api.deleteRolesVersion = jest.fn().mockRejectedValue('some-error');

        await rolesVersionStore.deleteRolesVersion(d.id);

        expect(api.deleteRolesVersion).toBeCalledWith(d.id);
        expect(rolesVersionStore.deleteRolesVersionError).toEqual({
            code: UnknownErrorCode,
            description: UnknownErrorDescription,
        });
    });

    it('clean roles version actions errors', () => {
        rolesVersionStore.createRolesVersionError = 'some-error' as any;
        rolesVersionStore.fetchRolesVersionError = 'some-error' as any;
        rolesVersionStore.updateRolesVersionError = 'some-error' as any;
        rolesVersionStore.deleteRolesVersionError = 'some-error' as any;

        rolesVersionStore.cleanRolesVersionActionErrors();

        expect(rolesVersionStore.createRolesVersionError).toBeNull();
        expect(rolesVersionStore.fetchRolesVersionError).toBeNull();
        expect(rolesVersionStore.updateRolesVersionError).toBeNull();
        expect(rolesVersionStore.deleteRolesVersionError).toBeNull();
    });
});
