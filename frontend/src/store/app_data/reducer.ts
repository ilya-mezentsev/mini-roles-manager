import { Action } from 'redux';
import * as log from '../../services/log';

import { ACTIONS } from './action_types';
import { AppDataError, AppDataResult } from './app_data.types';
import { APIError } from '../../services/api/shared';
import { UnknownErrorCode, UnknownErrorDescription } from '../shared/const';

export function appDataReducer(
    state: AppDataResult = {},
    action: Action<string> & { appDataResult?: AppDataError<APIError | Error> },
): AppDataResult<APIError> {
    switch (action.type) {
        case ACTIONS.SUCCESS_IMPORT_APP_DATA:
            return {
                importedOk: true,
            };

        case ACTIONS.FAILED_IMPORT_APP_DATA:
            let importFileValidationErrors: string[] = [];
            const importError = action.appDataResult?.importError as APIError;
            if (importError.code === 'invalid-import-file') {
                importFileValidationErrors = importError.description
                    .split('|')
                    .map(e => e.trim());

                importError.description = 'Validation error';
            }

            return {
                importedOk: false,
                appDataResult: {
                    importError,
                    importFileValidationErrors,
                },
            };

        case ACTIONS.FAILED_TO_PERFORM_APP_DATA_IMPORTING:
            log.error(
                `Unable to import app data: ${(action.appDataResult?.importError as Error).toString()}`
            );

            return {
                importedOk: false,
                appDataResult: {
                    importError: {
                        code: UnknownErrorCode,
                        description: UnknownErrorDescription,
                    },
                }
            };

        case ACTIONS.CLEAN_IMPORT_APP_DATA_RESULT:
        default:
            return {
                importedOk: false,
                appDataResult: {
                    importError: undefined,
                    importFileValidationErrors: undefined,
                },
            };
    }
}
