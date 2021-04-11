import { Role } from '../../../services/api';
import { RoleState } from './roles.types';

export interface RolesListActions {
    cleanFetchRolesErrorAction: () => void;

    updateRoleAction: (role: Role) => void;
    cleanUpdateRoleErrorAction: () => void;

    deleteRoleAction: (roleId: string) => void;
    cleanDeleteRoleErrorAction: () => void;
}

export type RolesListState = RoleState;

export type RolesListProps = RolesListActions & RolesListState;
