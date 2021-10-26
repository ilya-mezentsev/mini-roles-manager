import { makeAutoObservable } from 'mobx';

import * as log from '../../services/log';
import {
    AccountCredentials,
    signUp,
} from '../../services/api';
import { UnknownErrorCode, UnknownErrorDescription } from '../shared/const';
import { APIError } from '../../services/api/shared';

export class RegistrationStore {
    public registeredOk: boolean = false;
    public registrationError: APIError | null = null;

    constructor() {
        makeAutoObservable(this);
    }

    public async signUp(credentials: AccountCredentials): Promise<void> {
        try {
            const response = await signUp(credentials);

            if (response.isOk()) {
                this.registeredOk = true;
            } else {
                this.registrationError = response.data() as APIError;
            }
        } catch (e) {
            log.error(`Unable to perform registration: ${e.toString()}`);

            this.registrationError = {
                code: UnknownErrorCode,
                description: UnknownErrorDescription,
            };
        }
    }

    public cleanRegistrationAction(): void {
        this.registeredOk = false;
        this.registrationError = null;
    }
}
