import {
    Role,
    RolesVersion,
} from '../../../services/api';
import { RolesResult } from '../../../store/role/role.types';
import { ResourcesResult } from '../../../store/resource/resource.types';
import { RolesVersionResult } from '../../../store/roles_version/roles_version.types';

export interface RoleActions {
    createRoleAction: (role: Role) => void;
    cleanCreateRoleErrorAction: () => void;

    selectCurrentRolesVersionAction: (rv: RolesVersion) => void;
}

export interface RoleState {
    rolesVersionResult: RolesVersionResult;
    rolesResult: RolesResult;
    resourcesResult: ResourcesResult;
}

export type RoleProps = RoleActions & RoleState;
