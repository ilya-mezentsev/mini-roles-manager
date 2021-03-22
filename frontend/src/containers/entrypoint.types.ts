import { SessionResult } from '../store/session/session.types';

export interface EntrypointState {
    userSession?: SessionResult;
}

export interface EntrypointActions {
    login: () => void;
}

export type EntrypointProps = EntrypointState & EntrypointActions;
