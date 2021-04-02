import { APIError } from '../../services/api/shared';
import { AccountSession } from '../../services/api';

interface _SessionResult<T> {
    session?: AccountSession;
    error?: T | null;
}

export interface SessionActionResult extends _SessionResult<APIError | Error> {}

export interface SessionResult extends _SessionResult<APIError> {}
