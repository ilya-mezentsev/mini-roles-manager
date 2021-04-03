
export enum Operation {
    CREATE = 'create',
    READ = 'read',
    UPDATE = 'update',
    DELETE = 'delete',
}

export enum Effect {
    PERMIT = 'permit',
    DENY = 'deny',
}

export interface Role {
    id: string;
    title?: string;
    permissions?: string[];
    extends?: string[];
}

export interface Permission {
    id: string;
    operation: Operation;
    effect: Effect;
}
