import { EditableResource } from '../../../services/api';
import { ResourceState } from './resources.types';

export interface ResourcesListActions {
    updateResourceAction: (resource: EditableResource) => void;
    cleanUpdateResourceErrorAction: () => void;

    deleteResourceAction: (resourceId: string) => void;
    cleanDeleteResourceErrorAction: () => void;

    loadResourcesAction: () => void;
    cleanLoadResourcesError: () => void;
}

export type ResourcesListState = ResourceState;

export type ResourcesListProps = ResourcesListActions & ResourcesListState;
