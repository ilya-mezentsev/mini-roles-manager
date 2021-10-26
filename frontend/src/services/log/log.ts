
export function warning(message: string): void {
    console.warn(`[WARNING]: ${message}`);
}

export function error(message: string): void {
    console.error(`[ERROR]: ${message}`);
}

export function info(message: string): void {
    console.info(`[INFO]: ${message}`);
}
