export function removeLeadingAndTrailingSlashes(path: string): string {
    while (path.startsWith('/')) {
        path = path.substr(1)
    }
    while (path.endsWith('/')) {
        path = path.substr(0, path.length - 1);
    }

    return path;
}
