import { APIError } from '../../services/api/shared';

export interface RegistrationResult {
    error?: APIError | Error;
}
