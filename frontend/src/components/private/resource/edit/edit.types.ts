import EventEmitter from 'events';

import { EditableResource } from '../../../../services/api';

export interface EditResourceProps {
    openDialogueEventName: string;
    eventEmitter: EventEmitter;
    save: (r: EditableResource) => void;
    initialResource?: EditableResource | null;
}
