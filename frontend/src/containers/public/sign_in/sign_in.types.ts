import { AccountCredentials } from '../../../services/api';
import { SessionResult } from '../../../store/session/session.types';

export interface SignInActions {
    signInAction: (credentials: AccountCredentials) => void;
    cleanSignInAction: () => void;
}

export interface SignInState {
    userSession?: SessionResult;
}

export type SignInProps = SignInActions & SignInState;
