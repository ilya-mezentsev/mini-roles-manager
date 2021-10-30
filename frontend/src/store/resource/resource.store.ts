import { makeAutoObservable } from 'mobx';

import {
    EditableResource,
    Resource as ResourceModel,
    createResource,
    resourcesList,
    updateResource,
    deleteResource,
} from '../../services/api';
import * as log from '../../services/log';
import { APIError } from '../../services/api/shared';
import { UnknownErrorCode, UnknownErrorDescription } from '../shared/const';
import { RoleStore } from '../role/role.store';

export class ResourceStore {
    public list: ResourceModel[] = [];
    public createResourceError: APIError | null = null;
    public fetchResourceError: APIError | null = null;
    public updateResourceError: APIError | null = null;
    public deleteResourceError: APIError | null = null;

    constructor(
        private readonly roleStore: RoleStore,
    ) {
        makeAutoObservable(this);
    }

    public async createResource(resource: EditableResource): Promise<void> {
        try {
            const response = await createResource(resource);

            if (response.isOk()) {
                await this.fetchResources();
            } else {
                this.createResourceError = response.data() as APIError;
            }
        } catch (e) {
            log.error(`Unable to create resource: ${e.toString()}`);

            this.createResourceError = {
                code: UnknownErrorCode,
                description: UnknownErrorDescription,
            };
        }
    }

    public async fetchResources(): Promise<void> {
        try {
            const response = await resourcesList();

            if (response.isOk()) {
                this.list = (response.data() || []) as ResourceModel[];
            } else {
                this.fetchResourceError = response.data() as APIError;
            }
        } catch (e) {
            log.error(`Unable to fetch resource: ${e.toString()}`);

            this.fetchResourceError = {
                code: UnknownErrorCode,
                description: UnknownErrorDescription,
            };
        }
    }

    public async updateResource(resource: EditableResource): Promise<void> {
        try {
            const response = await updateResource(resource);

            if (response.isOk()) {
                const updatedResourceIndex = this.list.findIndex(r => r.id === resource.id);
                if (updatedResourceIndex >= 0) {
                    this.list[updatedResourceIndex] = {
                        ...this.list[updatedResourceIndex],
                        title: resource.title,
                        linksTo: resource.linksTo,
                    };
                }
            } else {
                this.updateResourceError = response.data() as APIError;
            }
        } catch (e) {
            log.error(`Unable to update resource: ${e.toString()}`);

            this.updateResourceError = {
                code: UnknownErrorCode,
                description: UnknownErrorDescription,
            };
        }
    }

    public async deleteResource(resourceId: string): Promise<void> {
        try {
            const response = await deleteResource(resourceId);

            if (response.isOk()) {
                this.list = this.list.filter(r => r.id !== resourceId);
            } else {
                this.deleteResourceError = response.data() as APIError;
            }
        } catch (e) {
            log.error(`Unable to delete resource: ${e.toString()}`);

            this.deleteResourceError = {
                code: UnknownErrorCode,
                description: UnknownErrorDescription,
            };
        } finally {
            // if some resource was deleted we need to reload roles => so store will not contain irrelevant roles
            if (this.deleteResourceError === null) {
                await this.roleStore.fetchRoles();
            }
        }
    }

    public cleanResourceActionError(): void {
        this.createResourceError = null;
        this.fetchResourceError = null;
        this.updateResourceError = null;
        this.deleteResourceError = null;
    }
}
