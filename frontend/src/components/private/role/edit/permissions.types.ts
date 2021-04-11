import {
    Effect,
    Operation,
    Permission,
    Resource,
} from '../../../../services/api';

export interface PermissionsProps {
    resources: Resource[];
    existsPermissions: Permission[];
    onPermissionsUpdate: (updatedPermissions: Permission[]) => void;
}

export interface ResourcePermission {
    operation: Operation;
    effect: Effect;
}
