import {
    APIError,
    APIResponse,
    APIResponseStatus,
    ErrorAPIResponse,
    ParsedAPIResponse, ResponseData,
    SuccessAPIResponse
} from './response.types';

export class SuccessResponse<T> implements ParsedAPIResponse<T> {
    constructor(
        private readonly _data: T,
    ) { }

    data(): T {
        return this._data;
    }

    isOk(): boolean {
        return true;
    }
}

export class ErrorResponse implements ParsedAPIResponse<APIError> {
    constructor(
        private readonly _data: APIError,
    ) { }

    data(): APIError {
        return this._data;
    }

    isOk(): boolean {
        return false;
    }
}

export type SuccessResponseConstructor<T> = (data: T) => SuccessResponse<T>;

export function errorResponseOrDefault(apiResponse: ErrorAPIResponse | null): ParsedAPIResponse<APIError | null> {
    if (apiResponse?.status === APIResponseStatus.ERROR) {
        return new ErrorResponse((apiResponse as ErrorAPIResponse).data);
    } else {
        return new SuccessResponse<null>(null);
    }
}

export function errorOrSuccessResponse<T>(apiResponse: APIResponse<T>): ParsedAPIResponse<T | APIError> {
    if (apiResponse?.status === APIResponseStatus.ERROR) {
        return new ErrorResponse((apiResponse as ErrorAPIResponse).data);
    } else {
        return new SuccessResponse<T>((apiResponse as SuccessAPIResponse<T>).data);
    }
}

export function parseResponse<T>(apiResponse: APIResponse<T>): ParsedAPIResponse<ResponseData<T>> {
    if (apiResponse === null) {
        return new SuccessResponse<null>(null);
    } else if (apiResponse?.status === APIResponseStatus.ERROR) {
        return new ErrorResponse((apiResponse as ErrorAPIResponse).data);
    } else {
        return new SuccessResponse<T>((apiResponse as SuccessAPIResponse<T>).data);
    }
}
