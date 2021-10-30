import { makeAutoObservable } from 'mobx';

import {
    ImportFile,
    importFromFile,
} from '../../services/api/';
import { APIError } from '../../services/api/shared';
import * as log from '../../services/log';
import { UnknownErrorCode, UnknownErrorDescription } from '../shared/const';
import { RoleStore } from '../role/role.store';
import { ResourceStore } from '../resource/resource.store';
import { RolesVersionStore } from '../roles_version/roles_version.store';

export class AppDataStore {
    public importedOk: boolean = false;
    public importError: APIError | null = null;
    public validationErrors: string[] = [];

    constructor(
        private readonly roleStore: RoleStore,
        private readonly resourceStore: ResourceStore,
        private readonly rolesVersionStore: RolesVersionStore
    ) {
        makeAutoObservable(this);
    }

    public async importFromFile(d: ImportFile): Promise<void> {
        try {
            const response = await importFromFile(d);

            if (response.isOk()) {
                this.importedOk = true;
            } else {
                this.processImportError(response.data() as APIError);
            }
        } catch (e) {
            log.error(`Unable to import app data: ${e.toString()}`);

            this.importError = {
                code: UnknownErrorCode,
                description: UnknownErrorDescription,
            };
        } finally {
            if (this.importedOk) {
                await Promise.all([
                    this.roleStore.fetchRoles(),
                    this.rolesVersionStore.fetchRolesVersions(),
                    this.resourceStore.fetchResources(),
                ]);
            }
        }
    }

    private processImportError(error: APIError): void {
        if (error.code === 'invalid-import-file') {
            this.validationErrors = error.description.split('|').map(e => e.trim());
            error.description = 'Validation error';
        }

        this.importError = error;
    }

    public cleanImportResult(): void {
        this.importError = null;
        this.importedOk = false;
        this.validationErrors = [];
    }
}
