import EventEmitter from 'events';

export interface PrompterProps {
    title: string;
    description: string;
    onAgree: () => void;
    onDisagree: () => void;
    openDialogueEventName: string;
    eventEmitter: EventEmitter;
}
