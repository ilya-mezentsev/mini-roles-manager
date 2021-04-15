import { PermissionAccessRequest } from '../../../services/api';
import { ResourcesResult } from '../../../store/resource/resource.types';
import { RolesResult } from '../../../store/role/role.types';
import { FetchPermissionResult } from '../../../store/permission/permissions.types';

export interface CheckPermissionsActions {
    fetchPermissionAction: (request: PermissionAccessRequest) => void;
    cleanFetchPermissionResult: () => void;
}

export interface CheckPermissionsState {
    fetchPermissionResult: FetchPermissionResult;
    rolesResult: RolesResult;
    resourcesResult: ResourcesResult;
}

export type CheckPermissionsProps = CheckPermissionsActions & CheckPermissionsState;
