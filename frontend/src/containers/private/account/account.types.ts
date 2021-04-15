import { AccountInfoResult } from '../../../store/account_info/account_info.types';

export interface AccountActions {
    cleanFetchInfoErrorAction: () => void;
}

export interface AccountState {
    accountInfoResult?: AccountInfoResult;
}

export type AccountProps = AccountActions & AccountState;
