import { Role } from '../../../services/api';
import { RolesResult } from '../../../store/role/role.types';
import { ResourcesResult } from '../../../store/resource/resource.types';

export interface RoleActions {
    createRoleAction: (role: Role) => void;
    cleanCreateRoleErrorAction: () => void;
}

export interface RoleState {
    rolesResult: RolesResult;
    resourcesResult: ResourcesResult;
}

export type RoleProps = RoleActions & RoleState;
