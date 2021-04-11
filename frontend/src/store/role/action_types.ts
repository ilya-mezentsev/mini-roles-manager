
export enum ACTIONS {
    SUCCESS_CREATE_ROLE = 'create-role:success',
    FAILED_CREATE_ROLE = 'create-role:error',
    FAILED_TO_PERFORM_ROLE_CREATION = 'create-role:unknown-error',
    CLEAN_CREATE_ROLE_ERROR = 'create-role:clean-error',

    SUCCESS_FETCH_ROLES = 'fetch-roles:success',
    FAILED_FETCH_ROLES = 'fetch-roles:error',
    FAILED_TO_PERFORM_ROLES_FETCHING = 'fetch-roles:unknown-error',
    CLEAN_FETCH_ROLES_ERROR = 'fetch-roles:clean-error',

    SUCCESS_UPDATE_ROLE = 'update-role:success',
    FAILED_UPDATE_ROLE = 'update-role:error',
    FAILED_TO_PERFORM_ROLE_UPDATING = 'update-role:unknown-error',
    CLEAN_UPDATE_ROLE_ERROR = 'update-role:clean-error',

    SUCCESS_DELETE_ROLE = 'delete-role:success',
    FAILED_DELETE_ROLE = 'delete-role:error',
    FAILED_TO_PERFORM_ROLE_DELETING = 'delete-role:unknown-error',
    CLEAN_DELETE_ROLE_ERROR = 'delete-role:clean-error',
}
