
export enum ACTIONS {
    SUCCESS_SIGN_IN = 'sign-in:success',
    FAILED_SIGN_IN = 'sign-in:error',
    FAILED_TO_PERFORM_SIGN_IN_ACTION = 'sign-in:unknown-error',
    CLEAN_SIGN_IN_ERROR = 'sign-in:clean-error',

    SUCCESS_LOGIN = 'login:success',
    FAILED_LOGIN = 'login:error',
    FAILED_TO_PERFORM_LOGIN_ACTION = 'login:unknown-error',

    SUCCESS_SIGN_OUT = 'sign-out:success',
    FAILED_SIGN_OUT = 'sign-out:error',
    FAILED_TO_PERFORM_SIGN_OUT_ACTION = 'sign-out:unknown-error',
    CLEAN_SIGN_OUT_ERROR = 'sign-out:clean-error',
}
