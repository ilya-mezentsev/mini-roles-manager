import { Color } from '@material-ui/lab';
import EventEmitter from 'events';

export interface AlertProps {
    message: string;
    severity: Color;
    shouldShow: boolean;
    onCloseCb: () => void;

    setOpenEventName: string;
    setOpenEmitter: EventEmitter;
}
