import { SessionResult } from '../../../store/session/session.types';

export interface SignOutState {
    userSession?: SessionResult;
}

export interface SignOutActions {
    signOutAction: () => void;
    cleanSignOutErrorAction: () => void;
}

export type SignOutProps = SignOutState & SignOutActions;
