import { makeAutoObservable } from 'mobx';

import * as log from '../../services/log';
import {
    AccountCredentials,
    AccountSession,
    signIn,
    signOut,
    login,
} from '../../services/api';
import { UnknownErrorCode, UnknownErrorDescription } from '../shared/const';
import { APIError } from '../../services/api/shared';

export class SessionStore {
    public session: AccountSession | null = null;
    public signInError: APIError | null = null;
    public signOutError: APIError | null = null;
    public loginError: APIError | null = null;

    constructor() {
        makeAutoObservable(this);
    }

    public async signIn(credentials: AccountCredentials): Promise<void> {
        try {
            const response = await signIn(credentials);

            if (response.isOk()) {
                this.session = response.data() as AccountSession;
            } else {
                this.signInError = response.data() as APIError;
            }
        } catch (e) {
            log.error(`Unable to sign-in: ${e.toString()}`);

            this.signInError = {
                code: UnknownErrorCode,
                description: UnknownErrorDescription,
            };
        }
    }

    public async signOut(): Promise<void> {
        try {
            const response = await signOut();

            if (response.isOk()) {
                this.session = null;
            } else {
                this.signOutError = response.data() as APIError;
            }
        } catch (e) {
            log.error(`Unable to sign-out: ${e.toString()}`);

            this.signOutError = {
                code: UnknownErrorCode,
                description: UnknownErrorDescription,
            };
        }
    }

    public async login(): Promise<void> {
        try {
            const response = await login();

            if (response.isOk()) {
                this.session = response.data() as AccountSession;
            } else {
                this.loginError = response.data() as APIError;
            }
        } catch (e) {
            log.error(`Unable to login: ${e.toString()}`);

            this.loginError = {
                code: UnknownErrorCode,
                description: UnknownErrorDescription,
            };
        }
    }

    public cleanSessionActionErrors(): void {
        this.signInError = null;
        this.signOutError = null;
        this.loginError = null;
    }
}
