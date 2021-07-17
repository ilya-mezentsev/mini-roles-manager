
export enum ACTIONS {
    SUCCESS_CREATE_ROLES_VERSION = 'create-roles-version:success',
    FAILED_CREATE_ROLES_VERSION = 'create-roles-version:error',
    FAILED_TO_PERFORM_ROLES_VERSION_CREATION = 'create-roles-version:unknown-error',
    CLEAN_CREATE_ROLES_VERSION_ERROR = 'create-roles-version:clean-error',

    SUCCESS_FETCH_ROLES_VERSION = 'fetch-roles-version:success',
    FAILED_FETCH_ROLES_VERSION = 'fetch-roles-version:error',
    FAILED_TO_PERFORM_ROLES_VERSION_FETCHING = 'fetch-roles-version:unknown-error',
    CLEAN_FETCH_ROLES_VERSION_ERROR = 'fetch-roles-version:clean-error',

    SUCCESS_UPDATE_ROLES_VERSION = 'update-roles-version:success',
    FAILED_UPDATE_ROLES_VERSION = 'update-roles-version:error',
    FAILED_TO_PERFORM_ROLES_VERSION_UPDATING = 'update-roles-version:unknown-error',
    CLEAN_UPDATE_ROLES_VERSION_ERROR = 'update-roles-version:clean-error',

    SUCCESS_DELETE_ROLES_VERSION = 'delete-roles-version:success',
    FAILED_DELETE_ROLES_VERSION = 'delete-roles-version:error',
    FAILED_TO_PERFORM_ROLES_VERSION_DELETING = 'delete-roles-version:unknown-error',
    CLEAN_DELETE_ROLES_VERSION_ERROR = 'delete-roles-version:clean-error',

    SELECT_CURRENT_ROLES_VERSION = 'select-current-roles-version',
}
