import { AccountInfoResult } from '../../../store/account_info/account_info.types';
import { AccountCredentials } from '../../../services/api';

export interface CredentialsActions {
    updateCredentialsAction: (c: AccountCredentials) => void;
    cleanUpdateCredentialsErrorAction: () => void;
}

export interface CredentialsState {
    accountInfoResult?: AccountInfoResult;
}

export type CredentialsProps = CredentialsActions & CredentialsState;
