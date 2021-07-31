import { RolesVersion } from '../../services/api';
import { APIError } from '../../services/api/shared';

export interface RolesVersionActionResult {
    rolesVersion: RolesVersion;
}

export interface RolesVersionsListResult {
    list: RolesVersion[] | null;
}

export interface RolesVersionIdActionResult {
    rolesVersionId: string;
}

export interface RolesVersionErrorResult {
    error: APIError | Error;
}

export type RolesVersionsActionResult =
    RolesVersionActionResult |
    RolesVersionsListResult |
    RolesVersionIdActionResult |
    RolesVersionErrorResult;

export interface RolesVersionResult extends RolesVersionsListResult {
    currentRolesVersion: RolesVersion | null;
    createError?: APIError;
    fetchError?: APIError;
    updateError?: APIError;
    deleteError?: APIError;
}
