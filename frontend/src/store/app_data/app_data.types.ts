import {APIError} from '../../services/api/shared';

export interface AppDataError<E> {
    importError?: E;
    importFileValidationErrors?: string[];
}

export interface AppDataResult<E = APIError | Error> {
    appDataResult?: AppDataError<E>;
    importedOk?: boolean;
}
