import { EditableResource } from '../../../services/api';
import { ResourcesResult } from '../../../store/resource/resource.types';

export interface ResourcesActions {
    createResourceAction: (resource: EditableResource) => void;
    cleanCreateResourceErrorAction: () => void;
}

export interface ResourceState {
    resourcesResult: ResourcesResult;
}

export type ResourceProps = ResourcesActions & ResourceState;
