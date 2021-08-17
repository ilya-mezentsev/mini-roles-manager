import { useState } from 'react';
import { TextField } from '@material-ui/core';

import { Alert } from '../../../components/shared';
import { StateToPropsFn } from '../../../shared/types';
import { ApiKeyProps, ApiKeyState } from './api_key.types';

export const ApiKey = (props: ApiKeyProps) => {
    const [shouldShowCopiedNotification, setShouldShowCopiedNotification] = useState(false);
    const tryCopy = () => {
        if (props.accountInfoResult?.info?.apiKey) {
            navigator.clipboard.writeText(props.accountInfoResult.info.apiKey)
                .then(() => notifyCopied());
        }
    };
    const notifyCopied = () => setShouldShowCopiedNotification(true);

    return (
        <>
            <h2>API Key:</h2>
            <TextField
                margin="dense"
                fullWidth
                disabled
                value={props.accountInfoResult?.info?.apiKey}
                onClick={tryCopy}
            />

            <Alert
                message="Copied to clipboard"
                severity="info"
                shouldShow={shouldShowCopiedNotification}
                onCloseCb={() => setShouldShowCopiedNotification(false)}
            />
        </>
    );
}

export const mapStateToProps: StateToPropsFn<ApiKeyState> = () => state => ({
    accountInfoResult: state.accountInfoResult,
});
