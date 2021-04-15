import {
    APIError,
    APIResponse,
    EmptyAPIResponse,
    ErrorAPIResponse,
    ParsedAPIResponse,
    ResponseData,
    SuccessAPIResponse,
    parseResponse,
    errorResponseOrDefault,
    errorOrSuccessResponse,
    POST,
    GET,
    DELETE,
    PATCH,
} from '../shared';
import {
    AccountCredentials,
    AccountInfo,
    AccountSession,
    PermissionAccessRequest,
    PermissionAccessResponse,
} from './account.types';

export async function signUp(credentials: AccountCredentials): Promise<ParsedAPIResponse<APIError | null>> {
    const response = await POST<ErrorAPIResponse | EmptyAPIResponse>(
        '/registration/user',
        { credentials },
    );

    return errorResponseOrDefault(response);
}

export async function login(): Promise<ParsedAPIResponse<ResponseData<AccountSession>>> {
    const response = await GET<APIResponse<AccountSession>>('/session');

    return parseResponse(response)
}

export async function signIn(credentials: AccountCredentials): Promise<ParsedAPIResponse<AccountSession | APIError>> {
    const response = await POST<SuccessAPIResponse<AccountSession> | ErrorAPIResponse>(
        '/session',
        { credentials },
    );

    return errorOrSuccessResponse(response);
}

export async function signOut(): Promise<ParsedAPIResponse<APIError | null>> {
    const response = await DELETE<ErrorAPIResponse | EmptyAPIResponse>('/session');

    return errorResponseOrDefault(response);
}

export async function fetchInfo(): Promise<ParsedAPIResponse<ResponseData<AccountInfo>>> {
    const response = await GET<APIResponse<AccountInfo>>('/account/info');

    return parseResponse(response);
}

export async function updateCredentials(credentials: AccountCredentials): Promise<ParsedAPIResponse<APIError | null>> {
    const response = await PATCH<ErrorAPIResponse | EmptyAPIResponse>(
        '/account/credentials',
        { credentials },
    );

    return errorResponseOrDefault(response);
}

export async function fetchPermission(request: PermissionAccessRequest): Promise<ParsedAPIResponse<ResponseData<PermissionAccessResponse>>> {
    const response = await POST<APIResponse<PermissionAccessResponse>>(
        '/check-permissions',
        { ...request },
    );

    return parseResponse(response);
}
