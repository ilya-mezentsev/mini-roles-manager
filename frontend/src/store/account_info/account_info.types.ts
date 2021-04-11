import { AccountCredentials, AccountInfo } from '../../services/api';
import { APIError } from '../../services/api/shared';

type ActionError = APIError | Error;

export interface FetchAccountInfoActionResult {
    info: AccountInfo;
}

export interface FetchAccountInfoActionErrorResult<E = ActionError> {
    fetchInfoError?: E;
}

export interface UpdateCredentialsActionResult {
    credentials?: AccountCredentials | null;
}

export interface UpdateCredentialsActionErrorResult<E = ActionError> {
    updateCredentialsError?: E;
}

export type AccountInfoActionResult =
    FetchAccountInfoActionResult | FetchAccountInfoActionErrorResult |
    UpdateCredentialsActionResult | UpdateCredentialsActionErrorResult;

export type AccountInfoResult =
    FetchAccountInfoActionResult &
    UpdateCredentialsActionResult &
    FetchAccountInfoActionErrorResult<APIError> &
    UpdateCredentialsActionErrorResult<APIError>
