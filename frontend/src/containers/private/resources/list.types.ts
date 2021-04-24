import { EditableResource } from '../../../services/api';
import { ResourcesResult } from '../../../store/resource/resource.types';
import { RolesResult } from '../../../store/role/role.types';

export interface ResourcesListActions {
    updateResourceAction: (resource: EditableResource) => void;
    cleanUpdateResourceErrorAction: () => void;

    deleteResourceAction: (resourceId: string) => void;
    cleanDeleteResourceErrorAction: () => void;
    cleanDeletedResourceIdAction: () => void;

    loadResourcesAction: () => void;
    cleanLoadResourcesError: () => void;

    loadRolesAction: () => void;
    cleanFetchRolesErrorAction: () => void;
}

export interface ResourcesListState {
    resourcesResult: ResourcesResult;
    rolesResult: RolesResult;
}

export type ResourcesListProps = ResourcesListActions & ResourcesListState;
