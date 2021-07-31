import EventEmitter from 'events';
import { Resource, Role } from '../../../../services/api';

export interface EditRoleProps {
    openDialogueEventName: string;
    eventEmitter: EventEmitter;
    existRoles: Role[];
    existsResources: Resource[];
    roleVersionId: string;
    save: (r: Role) => void;
    initialRole?: Role | null;
}
