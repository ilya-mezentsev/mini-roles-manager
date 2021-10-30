
export enum RequestMethod {
    GET = 'GET',
    POST = 'POST',
    PATCH = 'PATCH',
    DELETE = 'DELETE'
}

export interface RequestParams {
    path: string;
    method: RequestMethod;
    body?: any;
    shouldEncode?: boolean;
    shouldAddContentType?: boolean;
}
