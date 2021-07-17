import { RolesResult } from '../../store/role/role.types';
import { ResourcesResult } from '../../store/resource/resource.types';
import { AccountInfoResult } from '../../store/account_info/account_info.types';
import { RolesVersionResult } from '../../store/roles_version/roles_version.types';

export interface EntrypointActions {
    loadResourcesAction: () => void;
    loadRolesVersionsAction: () => void;
    loadRolesAction: () => void;
    loadAccountInfo: () => void;
}

export interface EntrypointState {
    rolesVersionResult: RolesVersionResult;
    rolesResult: RolesResult;
    resourcesResult: ResourcesResult;
    accountInfoResult: AccountInfoResult;
}

export type EntrypointProps = EntrypointActions & EntrypointState
