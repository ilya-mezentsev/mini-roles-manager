import { SessionActionResult } from '../store/session/session.types';

export interface EntrypointState {
    userSession?: SessionActionResult;
}

export interface EntrypointActions {
    login: () => void;
}

export type EntrypointProps = EntrypointState & EntrypointActions;
