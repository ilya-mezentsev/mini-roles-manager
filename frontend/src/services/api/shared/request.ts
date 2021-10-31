import { removeLeadingAndTrailingSlashes, makeQueryParams } from './helpers';
import { RequestMethod, RequestParams } from './request.types';

export async function GET<T>(
    path: string,
    params: {[key: string]: string} | null = null,
): Promise<T | null> {
    return await request({
        path: removeLeadingAndTrailingSlashes(path) + makeQueryParams(params),
        method: RequestMethod.GET,
    });
}

export async function POST<T>(
    path: string,
    body: any,
    isFile: boolean = false,
): Promise<T | null> {
    return await request({
        path: removeLeadingAndTrailingSlashes(path),
        method: RequestMethod.POST,
        body,
        shouldEncode: !isFile,
        shouldAddContentType: !isFile,
    });
}

export async function PATCH<T>(
    path: string,
    body: any
): Promise<T | null> {
    return await request({
        path: removeLeadingAndTrailingSlashes(path),
        method: RequestMethod.PATCH,
        body,
        shouldEncode: true,
        shouldAddContentType: true,
    });
}

export async function DELETE<T>(path: string): Promise<T | null> {
    return await request({
        path: removeLeadingAndTrailingSlashes(path),
        method: RequestMethod.DELETE,
    });
}

async function request<T>(params: RequestParams): Promise<T | null> {
    const options: any = {
        method: params.method,
        body: params.body
            ? params.shouldEncode
                ? JSON.stringify(params.body)
                : params.body
            : null,
    }
    if (params.shouldAddContentType) {
        options['headers'] = {
            'Content-Type': 'application/json',
        };
    }

    const res = await fetch(
        `/api/web-app/${params.path}`,
        options,
    );
    const responseText = await res.text();
    if (responseText) {
        return JSON.parse(responseText);
    } else if (!res.ok) {
        throw Error(`Response status (${res.status}) is unsuccessful and no body is present`);
    }

    return null;
}
