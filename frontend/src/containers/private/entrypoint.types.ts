import { RolesResult } from '../../store/role/role.types';
import { ResourcesResult } from '../../store/resource/resource.types';

export interface EntrypointActions {
    loadResourcesAction: () => void;
    loadRolesAction: () => void;
}

export interface EntrypointState {
    rolesResult: RolesResult;
    resourcesResult: ResourcesResult;
}

export type EntrypointProps = EntrypointActions & EntrypointState
