
export enum ACTIONS {
    SUCCESS_CREATE_RESOURCE = 'create-resource:success',
    FAILED_CREATE_RESOURCE = 'create-resource:error',
    FAILED_TO_PERFORM_RESOURCE_CREATION = 'create-resource:unknown-error',
    CLEAN_CREATE_RESOURCE_ERROR = 'create-resource:clean-error',

    SUCCESS_FETCH_RESOURCES = 'fetch-resources:success',
    FAILED_FETCH_RESOURCES = 'fetch-resources:error',
    FAILED_TO_PERFORM_RESOURCES_FETCHING = 'fetch-resources:unknown-error',
    CLEAN_FETCH_RESOURCES_ERROR = 'fetch-resources:clean-error',

    SUCCESS_UPDATE_RESOURCE = 'update-resource:success',
    FAILED_UPDATE_RESOURCE = 'update-resource:error',
    FAILED_TO_PERFORM_RESOURCE_UPDATING = 'update-resource:unknown-error',
    CLEAN_UPDATE_RESOURCE_ERROR = 'update-resource:clean-error',

    SUCCESS_DELETE_RESOURCE = 'delete-resource:success',
    FAILED_DELETE_RESOURCE = 'delete-resource:error',
    FAILED_TO_PERFORM_RESOURCE_DELETING = 'delete-resource:unknown-error',
    CLEAN_DELETE_RESOURCE_ERROR = 'delete-resource:clean-error',
}
