
export interface EditableResource {
    id: string;
    title?: string;
    linksTo?: string[];
}

export interface Resource extends EditableResource {
    permissions: Permission[];
}

enum Operation {
    CREATE = 'create',
    READ = 'read',
    UPDATE = 'update',
    DELETE = 'delete',
}

enum Effect {
    PERMIT = 'permit',
    DENY = 'deny',
}

interface Permission {
    id: string;
    operation: Operation;
    effect: Effect;
}
