import {
    Permission,
    Resource,
} from '../../../../services/api';
import { Effect, Operation } from '../../../../services/api/shared/types';

export interface PermissionsProps {
    resources: Resource[];
    existsPermissions: Permission[];
    onPermissionsUpdate: (updatedPermissions: Permission[]) => void;
}

export interface ResourcePermission {
    operation: Operation;
    effect: Effect;
}
