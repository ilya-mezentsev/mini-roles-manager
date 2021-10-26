import { makeAutoObservable } from 'mobx';

import * as log from '../../services/log';
import { Role } from '../../services/api';
import {
    createRole,
    rolesList,
    updateRole,
    deleteRole,
} from '../../services/api';
import { APIError } from '../../services/api/shared';
import { UnknownErrorCode, UnknownErrorDescription } from '../shared/const';

export class RoleStore {
    public list: Role[] = [];
    public createRoleError: APIError | null = null;
    public fetchRoleError: APIError | null = null;
    public updateRoleError: APIError | null = null;
    public deleteRoleError: APIError | null = null;

    constructor() {
        makeAutoObservable(this);
    }

    public async createRole(role: Role): Promise<void> {
        try {
            const response = await createRole(role);

            if (response.isOk()) {
                this.list.push(role);
            } else {
                this.createRoleError = response.data() as APIError;
            }
        } catch (e) {
            log.error(`Unable to create role: ${e.toString()}`);

            this.createRoleError = {
                code: UnknownErrorCode,
                description: UnknownErrorDescription,
            };
        }
    }

    public async fetchRoles(): Promise<void> {
        try {
            const response = await rolesList();

            if (response.isOk()) {
                this.list = (response.data() || []) as Role[];
            } else {
                this.fetchRoleError = response.data() as APIError;
            }
        } catch (e) {
            log.error(`Unable to fetch roles: ${e.toString()}`);

            this.fetchRoleError = {
                code: UnknownErrorCode,
                description: UnknownErrorDescription,
            };
        }
    }

    public async updateRole(role: Role): Promise<void> {
        try {
            const response = await updateRole(role);

            if (response.isOk()) {
                const updatedRoleIndex = this.list
                    .findIndex(r => r.id === role.id && r.versionId === role.versionId);
                if (updatedRoleIndex >= 0) {
                    this.list[updatedRoleIndex] = {
                        ...this.list[updatedRoleIndex],
                        title: role.title,
                        permissions: role.permissions,
                        extends: role.extends,
                    };
                }
            } else {
                this.updateRoleError = response.data() as APIError;
            }
        } catch (e) {
            log.error(`Unable to update role: ${e.toString()}`);

            this.updateRoleError = {
                code: UnknownErrorCode,
                description: UnknownErrorDescription,
            };
        }
    }

    public async deleteRole(rolesVersionId: string, roleId: string): Promise<void> {
        try {
            const response = await deleteRole(rolesVersionId, roleId);

            if (response.isOk()) {
                this.list = this.list.filter(r => r.id !== roleId || r.versionId !== rolesVersionId);
            } else {
                this.deleteRoleError = response.data() as APIError;
            }
        } catch (e) {
            log.error(`Unable to delete role: ${e.toString()}`);

            this.deleteRoleError = {
                code: UnknownErrorCode,
                description: UnknownErrorDescription,
            };
        }
    }

    public cleanRoleActionErrors(): void {
        this.createRoleError = null;
        this.fetchRoleError = null;
        this.updateRoleError = null;
        this.deleteRoleError = null;
    }
}
