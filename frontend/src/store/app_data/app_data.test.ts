import { AppDataStore } from './app_data.store';

import * as api from '../../services/api';
import { ErrorResponse, SuccessResponse } from '../../services/api/shared';
import { UnknownErrorCode, UnknownErrorDescription } from '../shared/const';
import { RoleStore } from '../role/role.store';
import { ResourceStore } from '../resource/resource.store';
import { RolesVersionStore } from '../roles_version/roles_version.store';

jest.mock('../../services/api');

describe('app data tests', () => {
    let roleStore = new RoleStore();
    let resourceStore = new ResourceStore(roleStore);
    let rolesVersionStore = new RolesVersionStore(roleStore);

    let appDataStore = new AppDataStore(
        roleStore,
        resourceStore,
        rolesVersionStore,
    );

    beforeEach(() => {
        appDataStore.cleanImportResult();
    });

    it('import from file success', async () => {
        const d = {
            file: new File([], 'filename'),
        };
        // @ts-ignore
        api.importFromFile = jest.fn().mockResolvedValue(new SuccessResponse(null));

        await appDataStore.importFromFile(d);

        expect(api.resourcesList).toBeCalled();
        expect(api.rolesList).toBeCalled();
        expect(api.rolesVersionsList).toBeCalled();
        expect(api.importFromFile).toBeCalledWith(d);
        expect(appDataStore.importedOk).toBeTruthy();
    });

    it('import from file parsed error', async () => {
        const d = {
            file: new File([], 'filename'),
        };
        // @ts-ignore
        api.importFromFile = jest.fn().mockResolvedValue(new ErrorResponse('some-error'));

        await appDataStore.importFromFile(d);

        expect(api.importFromFile).toBeCalledWith(d);
        expect(appDataStore.importedOk).toBeFalsy();
        expect(appDataStore.importError).toEqual('some-error');
    });

    it('import from file validation error', async () => {
        const d = {
            file: new File([], 'filename'),
        };
        // @ts-ignore
        api.importFromFile = jest.fn().mockResolvedValue(new ErrorResponse({
            code: 'invalid-import-file',
            description: 'foo | bar',
        }));

        await appDataStore.importFromFile(d);

        expect(api.importFromFile).toBeCalledWith(d);
        expect(appDataStore.importedOk).toBeFalsy();
        expect(appDataStore.importError).toEqual({
            code: 'invalid-import-file',
            description: 'Validation error',
        });
        expect(appDataStore.validationErrors).toEqual(['foo', 'bar']);
    });

    it('import from file unknown error', async () => {
        const d = {
            file: new File([], 'filename'),
        };
        // @ts-ignore
        api.importFromFile = jest.fn().mockRejectedValue('some-error');

        await appDataStore.importFromFile(d);

        expect(api.importFromFile).toBeCalledWith(d);
        expect(appDataStore.importedOk).toBeFalsy();
        expect(appDataStore.importError).toEqual({
            code: UnknownErrorCode,
            description: UnknownErrorDescription,
        });
    });

    it('clean import result', () => {
        appDataStore.cleanImportResult();

        expect(appDataStore.importedOk).toBeFalsy();
        expect(appDataStore.importError).toBeNull();
        expect(appDataStore.validationErrors).toEqual([]);
    });
});
