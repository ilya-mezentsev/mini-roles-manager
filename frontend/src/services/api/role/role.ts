import {
    APIError,
    EmptyAPIResponse,
    ErrorAPIResponse,
    ParsedAPIResponse,
    ResponseData,
    SuccessAPIResponse,
    errorOrSuccessResponse,
    errorResponseOrDefault,
    POST,
    GET,
    PATCH,
    DELETE,
} from '../shared';
import { Role } from './role.types';

export async function createRole(role: Role): Promise<ParsedAPIResponse<APIError | null>> {
    const response = await POST<ErrorAPIResponse | EmptyAPIResponse>('/role', { role });

    return errorResponseOrDefault(response);
}

export async function rolesList(): Promise<ParsedAPIResponse<ResponseData<Role[]>>> {
    const response = await GET<SuccessAPIResponse<Role[]> | ErrorAPIResponse>('/roles');

    return errorOrSuccessResponse(response);
}

export async function updateRole(role: Role): Promise<ParsedAPIResponse<APIError | null>> {
    const response = await PATCH<ErrorAPIResponse | EmptyAPIResponse>('/role', { role });

    return errorResponseOrDefault(response);
}

export async function deleteRole(roleId: string): Promise<ParsedAPIResponse<APIError | null>> {
    const response = await DELETE<ErrorAPIResponse | EmptyAPIResponse>(`/role/${roleId}`);

    return errorResponseOrDefault(response);
}
