import { AccountInfoResult } from '../../../store/account_info/account_info.types';

export interface ApiKeyState {
    accountInfoResult?: AccountInfoResult;
}

export type ApiKeyProps = ApiKeyState;
