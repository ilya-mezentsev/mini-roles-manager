
export enum ACTIONS {
    SUCCESS_FETCH_INFO = 'fetch-info:success',
    FAILED_FETCH_INFO = 'fetch-info:error',
    FAILED_TO_PERFORM_INFO_FETCHING = 'fetch-info:unknown-error',
    CLEAN_FETCH_INFO_ERROR = 'fetch-info:clean-error',

    SUCCESS_UPDATE_CREDENTIALS = 'update-credentials:success',
    FAILED_UPDATE_CREDENTIALS = 'update-credentials:error',
    FAILED_TO_PERFORM_CREDENTIALS_UPDATING = 'update-credentials:unknown-error',
    CLEAN_UPDATE_CREDENTIALS_RESULT = 'update-credentials:clean',
}
