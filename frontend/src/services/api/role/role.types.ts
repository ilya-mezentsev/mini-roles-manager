import { Effect, Operation } from '../shared/types';


export interface Role {
    id: string;
    versionId: string;
    title?: string;
    permissions?: string[];
    extends?: string[];
}

export interface Permission {
    id: string;
    operation: Operation;
    effect: Effect;
}
