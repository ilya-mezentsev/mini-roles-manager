import { Effect, Operation } from '../shared/types';

export interface EditableResource {
    id: string;
    title?: string;
    linksTo?: string[];
}

export interface Resource extends EditableResource {
    permissions: Permission[];
}

interface Permission {
    id: string;
    operation: Operation;
    effect: Effect;
}
