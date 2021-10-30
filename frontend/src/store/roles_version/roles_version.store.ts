import { makeAutoObservable } from 'mobx';

import * as log from '../../services/log';
import { APIError } from '../../services/api/shared';
import { RolesVersion } from '../../services/api';
import {
    createRolesVersion,
    rolesVersionsList,
    updateRolesVersion,
    deleteRolesVersion,
} from '../../services/api';
import { UnknownErrorCode, UnknownErrorDescription } from '../shared/const';
import { RoleStore } from '../role/role.store';

export class RolesVersionStore {
    public list: RolesVersion[] = [];
    public current: RolesVersion | null = null;
    public createRolesVersionError: APIError | null = null;
    public fetchRolesVersionError: APIError | null = null;
    public updateRolesVersionError: APIError | null = null;
    public deleteRolesVersionError: APIError | null = null;

    constructor(
        private readonly roleStore: RoleStore,
    ) {
        makeAutoObservable(this);
    }

    public setCurrentRolesVersion(rolesVersion: RolesVersion): void {
        this.current = rolesVersion;
    }

    public async createRolesVersion(rolesVersion: RolesVersion): Promise<void> {
        try {
            const response = await createRolesVersion(rolesVersion);

            if (response.isOk()) {
                this.list.push(rolesVersion);
            } else {
                this.createRolesVersionError = response.data() as APIError;
            }
        } catch (e) {
            log.error(`Unable to create roles version: ${e.toString()}`);

            this.createRolesVersionError = {
                code: UnknownErrorCode,
                description: UnknownErrorDescription,
            };
        }
    }

    public async fetchRolesVersions(): Promise<void> {
        try {
            const response = await rolesVersionsList();

            if (response.isOk()) {
                this.list = response.data() as RolesVersion[];
                this.setCurrentRolesVersion(this.list[0]);
            } else {
                this.fetchRolesVersionError = response.data() as APIError;
            }
        } catch (e) {
            log.error(`Unable to fetch roles versions: ${e.toString()}`);

            this.fetchRolesVersionError = {
                code: UnknownErrorCode,
                description: UnknownErrorDescription,
            };
        }
    }

    public async updateRolesVersion(rolesVersion: RolesVersion): Promise<void> {
        try {
            const response = await updateRolesVersion(rolesVersion);

            if (response.isOk()) {
                const updatedRolesVersionIndex = this.list.findIndex(rv => rv.id === rolesVersion.id);
                if (updatedRolesVersionIndex >= 0) {
                    this.list[updatedRolesVersionIndex] = {
                        ...this.list[updatedRolesVersionIndex],
                        title: rolesVersion.title,
                    };

                    if (this.current!.id === rolesVersion.id) {
                        this.setCurrentRolesVersion(rolesVersion);
                    }
                }
            } else {
                this.updateRolesVersionError = response.data() as APIError;
            }
        } catch (e) {
            log.error(`Unable to update roles version: ${e.toString()}`);

            this.updateRolesVersionError = {
                code: UnknownErrorCode,
                description: UnknownErrorDescription,
            };
        }
    }

    public async deleteRolesVersion(rolesVersionId: string): Promise<void> {
        try {
            const response = await deleteRolesVersion(rolesVersionId);

            if (response.isOk()) {
                this.list = this.list.filter(rv => rv.id !== rolesVersionId);
                if (this.current!.id === rolesVersionId) {
                    this.setCurrentRolesVersion(this.list[0]);
                }
            } else {
                this.deleteRolesVersionError = response.data() as APIError;
            }
        } catch (e) {
            log.error(`Unable to update roles version: ${e.toString()}`);

            this.deleteRolesVersionError = {
                code: UnknownErrorCode,
                description: UnknownErrorDescription,
            };
        } finally {
            // if some roles version was deleted we need to reload roles => so store will not contain irrelevant roles
            if (this.deleteRolesVersionError === null) {
                await this.roleStore.fetchRoles();
            }
        }
    }

    public cleanRolesVersionActionErrors(): void {
        this.createRolesVersionError = null;
        this.fetchRolesVersionError = null;
        this.updateRolesVersionError = null;
        this.deleteRolesVersionError = null;
    }
}
