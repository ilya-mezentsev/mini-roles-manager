export function removeLeadingAndTrailingSlashes(path: string): string {
    while (path.startsWith('/')) {
        path = path.substr(1)
    }
    while (path.endsWith('/')) {
        path = path.substr(0, path.length - 1);
    }

    return path;
}


export function makeQueryParams(params: {[key: string]: string} | null = null): string {
    if (params) {
        return '?' + Object.entries(params)
            .map(([key, value]) => `${key}=${value}`)
            .join('&');
    } else {
        return '';
    }
}
