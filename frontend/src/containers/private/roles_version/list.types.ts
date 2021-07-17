import { RolesVersion } from '../../../services/api';
import { RolesVersionState } from './roles_version.types';
import { RolesResult } from '../../../store/role/role.types';

export interface RolesVersionListActions {
    cleanFetchRolesVersionsErrorAction: () => void;

    updateRolesVersionAction: (rv: RolesVersion) => void;
    cleanUpdateRolesVersionErrorAction: () => void;

    deleteRolesVersionAction: (rolesVersionId: string) => void;
    cleanDeleteRolesVersionErrorAction: () => void;

    loadRolesAction: () => void;
    cleanFetchRolesErrorAction: () => void;
}

export interface RolesVersionListState extends RolesVersionState {
    rolesResult: RolesResult;
}

export type RolesVersionListProps = RolesVersionListActions & RolesVersionListState;
