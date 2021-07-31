import { Effect, Operation } from '../shared/types';

export interface AccountCredentials {
    login: string;
    password: string;
}

export interface AccountSession {
    id: string;
}

export interface AccountInfo {
    apiKey: string;
    created: string;
    login: string;
}

export interface PermissionAccessRequest {
    rolesVersionId: string;
    roleId: string;
    resourceId: string;
    operation: Operation;
}

export interface PermissionAccessResponse {
    effect: Effect;
}
