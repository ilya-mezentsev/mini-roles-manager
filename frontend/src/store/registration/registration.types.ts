import { APIError } from '../../services/api/shared';

export interface RegistrationActionResult {
    error?: APIError | Error;
}

export interface RegistrationResult {
    error?: APIError;
}
