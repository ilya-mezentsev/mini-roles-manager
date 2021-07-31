import { RolesVersion } from './roles_version.types';
import {
    APIError,
    EmptyAPIResponse,
    ErrorAPIResponse,
    ParsedAPIResponse,
    ResponseData,
    SuccessAPIResponse,
    POST,
    GET,
    PATCH,
    DELETE,
    errorResponseOrDefault,
    errorOrSuccessResponse,
} from '../shared';

export async function createRolesVersion(rolesVersion: RolesVersion): Promise<ParsedAPIResponse<APIError | null>> {
    const response = await POST<ErrorAPIResponse | EmptyAPIResponse>('/roles-version', { rolesVersion });

    return errorResponseOrDefault(response);
}

export async function rolesVersionsList(): Promise<ParsedAPIResponse<ResponseData<RolesVersion[]>>> {
    const response = await GET<SuccessAPIResponse<RolesVersion[]> | ErrorAPIResponse>('/roles-versions');

    return errorOrSuccessResponse(response);
}

export async function updateRolesVersion(rolesVersion: RolesVersion): Promise<ParsedAPIResponse<APIError | null>> {
    const response = await PATCH<ErrorAPIResponse | EmptyAPIResponse>('/roles-version', { rolesVersion });

    return errorResponseOrDefault(response);
}

export async function deleteRolesVersion(rolesVersionId: string): Promise<ParsedAPIResponse<APIError | null>> {
    const response = await DELETE<ErrorAPIResponse | EmptyAPIResponse>(`/roles-version/${rolesVersionId}`);

    return errorResponseOrDefault(response);
}
