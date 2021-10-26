import { useState } from 'react';
import { observer } from 'mobx-react-lite';
import { TextField } from '@material-ui/core';

import { Alert } from '../../../components/shared';
import { accountInfoStore } from '../../../store';

export const ApiKey = observer(() => {
    const [shouldShowCopiedNotification, setShouldShowCopiedNotification] = useState(false);
    const tryCopy = () => {
        if (accountInfoStore.info?.apiKey) {
            navigator.clipboard.writeText(accountInfoStore.info.apiKey)
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
                value={accountInfoStore.info?.apiKey}
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
});
