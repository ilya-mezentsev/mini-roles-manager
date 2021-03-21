export enum APIResponseStatus {
    OK = 'ok',
    ERROR = 'error'
}

export type APIResponse<T> = SuccessAPIResponse<T> | ErrorAPIResponse | EmptyAPIResponse;

export type SuccessAPIResponse<T> = {
    status: APIResponseStatus.OK;
    data: T;
}

export type ErrorAPIResponse = {
    status: APIResponseStatus.ERROR;
    data: APIError;
}

export interface APIError {
    code: string;
    description: string;
}

export type EmptyAPIResponse = null;

export type ResponseData<T> = T | APIError | null

export interface ParsedAPIResponse<T> {
    isOk(): boolean;
    data(): T;
}
