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
} from '../shared';
import { AccountCredentials, AccountSession } from './account.types';

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

    return errorResponseOrDefault(response)
}
