import { Resource, EditableResource } from './resource.types';
import {
    APIError,
    ParsedAPIResponse,
    ResponseData,
    EmptyAPIResponse,
    ErrorAPIResponse,
    SuccessAPIResponse,
    errorResponseOrDefault,
    errorOrSuccessResponse,
    GET,
    POST,
    PATCH,
    DELETE,
} from '../shared';

export async function createResource(resource: EditableResource): Promise<ParsedAPIResponse<APIError | null>> {
    const response = await POST<ErrorAPIResponse | EmptyAPIResponse>('/resource', {
        resource,
    });

    return errorResponseOrDefault(response);
}

export async function resourcesList(): Promise<ParsedAPIResponse<ResponseData<Resource[]>>> {
    const response = await GET<SuccessAPIResponse<Resource[]> | ErrorAPIResponse>('/resources');

    return errorOrSuccessResponse(response);
}

export async function updateResource(resource: EditableResource): Promise<ParsedAPIResponse<APIError | null>> {
    const response = await PATCH<ErrorAPIResponse | EmptyAPIResponse>('/resource', {
        resource,
    });

    return errorResponseOrDefault(response);
}

export async function deleteResource(resourceId: string): Promise<ParsedAPIResponse<APIError | null>> {
    const response = await DELETE<ErrorAPIResponse | EmptyAPIResponse>(`/resource/${resourceId}`);

    return errorResponseOrDefault(response);
}
