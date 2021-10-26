import { makeAutoObservable } from 'mobx';

import * as log from '../../services/log';
import {
    AccountCredentials,
    AccountInfo,
    fetchInfo,
    updateCredentials,
} from '../../services/api';
import { APIError } from '../../services/api/shared';
import { UnknownErrorCode, UnknownErrorDescription } from '../shared/const';

export class AccountInfoStore {
    public info: AccountInfo | null = null;
    public credentials: AccountCredentials | null = null;
    public fetchInfoError: APIError | null = null;
    public updateCredentialsError: APIError | null = null;

    constructor() {
        makeAutoObservable(this);
    }

    public async fetchInfo(): Promise<void> {
        try {
            const response = await fetchInfo();

            if (response.isOk()) {
                this.info = response.data() as AccountInfo;
            } else {
                this.fetchInfoError = response.data() as APIError;
            }
        } catch (e) {
            log.error(`Unable to fetch account info: ${e.toString()}`);

            this.fetchInfoError = {
                code: UnknownErrorCode,
                description: UnknownErrorDescription,
            };
        }
    }

    public async updateCredentials(credentials: AccountCredentials): Promise<void> {
        try {
            const response = await updateCredentials(credentials);

            if (response.isOk()) {
                this.info!.login = credentials.login;
                this.credentials = credentials;
            } else {
                this.updateCredentialsError = response.data() as APIError;
            }
        } catch (e) {
            log.error(`Unable to update credentials: ${e.toString()}`);

            this.updateCredentialsError = {
                code: UnknownErrorCode,
                description: UnknownErrorDescription,
            };
        }
    }

    public cleanAccountInfoActionErrors(): void {
        this.fetchInfoError = null;
        this.updateCredentialsError = null;
    }
}
