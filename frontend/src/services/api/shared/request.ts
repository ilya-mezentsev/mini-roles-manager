import { removeLeadingAndTrailingSlashes } from './helpers';
import { RequestMethod, RequestParams } from './request.types';

export async function GET<T>(path: string): Promise<T | null> {
    return await request({
        path: removeLeadingAndTrailingSlashes(path),
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
        `/api/${params.path}`,
        options,
    );
    const resHasText = !!res.body;
    if (resHasText) {
        return await res.json();
    } else if (
        !res.ok &&
        !resHasText
    ) {
        throw Error(`Response status (${res.status}) is unsuccessful and no body is present`);
    }

    return null;
}
