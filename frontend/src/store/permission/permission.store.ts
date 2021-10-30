import { makeAutoObservable } from 'mobx';

import * as log from '../../services/log';
import {
    fetchPermission,
    PermissionAccessRequest,
} from '../../services/api';
import { APIError } from '../../services/api/shared';
import { UnknownErrorCode, UnknownErrorDescription } from '../shared/const';
import { Effect } from '../../services/api/shared/types';

export class PermissionStore {
    public permission: {effect: Effect} | null = null;
    public fetchPermissionError: APIError | null = null;

    constructor() {
        makeAutoObservable(this);
    }

    public async fetchPermission(request: PermissionAccessRequest): Promise<void> {
        try {
            const response = await fetchPermission(request);

            if (response.isOk()) {
                this.permission = response.data() as {effect: Effect};
            } else {
                this.fetchPermissionError = response.data() as APIError;
            }
        } catch (e) {
            log.error(`Unable to fetch permission: ${e.toString()}`);

            this.fetchPermissionError = {
                code: UnknownErrorCode,
                description: UnknownErrorDescription,
            };
        }
    }

    public cleanFetchPermissionAction(): void {
        this.permission = null;
        this.fetchPermissionError = null;
    }
}
