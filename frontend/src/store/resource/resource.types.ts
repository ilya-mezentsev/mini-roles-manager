import { APIError } from '../../services/api/shared';
import { Resource, EditableResource } from '../../services/api';

export interface ResourceActionResult {
    resource: EditableResource;
}

export interface ResourcesListResult {
    list: Resource[] | null;
}

export interface ResourceIdActionResult {
    resourceId: string;
}

export interface ResourceErrorResult {
    error: APIError | Error;
}

export type ResourcesActionResult = ResourceActionResult | ResourcesListResult | ResourceIdActionResult | ResourceErrorResult;

export interface ResourcesResult extends ResourcesListResult {
    deletedResourceId?: string;
    createError?: APIError;
    fetchError?: APIError;
    updateError?: APIError;
    deleteError?: APIError;
}
