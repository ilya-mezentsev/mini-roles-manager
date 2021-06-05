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
): Promise<T | null> {
    return await request({
        path: removeLeadingAndTrailingSlashes(path),
        method: RequestMethod.POST,
        body,
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
    });
}

export async function DELETE<T>(path: string): Promise<T | null> {
    return await request({
        path: removeLeadingAndTrailingSlashes(path),
        method: RequestMethod.DELETE,
    });
}

async function request<T>(params: RequestParams): Promise<T | null> {
    const options = {
        method: params.method,
        headers: {
            'Content-Type': 'application/json',
        },
        body: params.body ? JSON.stringify(params.body) : null,
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
