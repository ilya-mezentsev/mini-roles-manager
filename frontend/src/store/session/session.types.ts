import { APIError } from '../../services/api/shared';
import { AccountSession } from '../../services/api';

export interface SessionResult {
    session?: AccountSession;
    error?: APIError | Error | null;
}
