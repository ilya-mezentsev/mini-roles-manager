import { AccountInfoResult } from '../../../store/account_info/account_info.types';
import { RolesResult } from '../../../store/role/role.types';
import { ResourcesResult } from '../../../store/resource/resource.types';

export interface AccountInfoActions {
    cleanFetchInfoErrorAction: () => void;
}

export interface AccountInfoState {
    accountInfoResult?: AccountInfoResult;
    rolesResult: RolesResult;
    resourcesResult: ResourcesResult;
}

export type AccountInfoProps = AccountInfoActions & AccountInfoState;
