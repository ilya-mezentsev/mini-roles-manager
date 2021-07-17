import { APIError } from '../../services/api/shared';
import { Role } from '../../services/api';

export interface RoleActionResult {
    role: Role;
}

export interface RolesListResult {
    list: Role[] | null;
}

export interface RoleIdActionResult {
    roleId: string;
    rolesVersionId: string;
}

export interface RoleErrorResult {
    error: APIError | Error;
}

export type RolesActionResult = RoleActionResult | RolesListResult | RoleIdActionResult | RoleErrorResult;

export interface RolesResult extends RolesListResult {
    createError?: APIError;
    fetchError?: APIError;
    updateError?: APIError;
    deleteError?: APIError;
}
