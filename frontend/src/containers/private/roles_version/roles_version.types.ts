import { RolesVersion } from '../../../services/api';
import { RolesVersionResult } from '../../../store/roles_version/roles_version.types';

export interface RolesVersionActions {
    createRolesVersionAction: (rv: RolesVersion) => void;
    cleanCreateRolesVersionErrorAction: () => void;
}

export interface RolesVersionState {
    rolesVersionResult: RolesVersionResult;
}

export type RolesVersionProps = RolesVersionActions & RolesVersionState;
