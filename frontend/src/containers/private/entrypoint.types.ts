import { RolesResult } from '../../store/role/role.types';
import { ResourcesResult } from '../../store/resource/resource.types';
import { AccountInfoResult } from '../../store/account_info/account_info.types';

export interface EntrypointActions {
    loadResourcesAction: () => void;
    loadRolesAction: () => void;
    loadAccountInfo: () => void;
}

export interface EntrypointState {
    rolesResult: RolesResult;
    resourcesResult: ResourcesResult;
    accountInfoResult: AccountInfoResult;
}

export type EntrypointProps = EntrypointActions & EntrypointState
