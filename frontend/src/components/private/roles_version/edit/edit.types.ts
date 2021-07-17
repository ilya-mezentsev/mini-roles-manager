import EventEmitter from 'events';

import { RolesVersion } from '../../../../services/api';

export interface EditRolesVersionProps {
    openDialogueEventName: string;
    eventEmitter: EventEmitter;
    save: (rv: RolesVersion) => void;
    initialRolesVersion?: RolesVersion | null;
}
