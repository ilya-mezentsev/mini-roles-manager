import { useEffect } from 'react';
import { observer } from 'mobx-react-lite';
import {
    Box,
    Button,
} from '@material-ui/core';
import { Alert as MaterialAlert } from '@material-ui/lab';

import { Alert } from '../../../components/shared';
import { appDataStore } from '../../../store';

export const Import = observer(() => {
    // eslint-disable-next-line
    useEffect(() => () => appDataStore.cleanImportResult(), []);

    const handleFileSelected = (e: any): void => {
        appDataStore.importFromFile({
            file: e.target.files[0],
        }).finally(() => e.target.value = null);
    }

    return (
        <>
            <h2>To import application data from file click button below:</h2>
            <MaterialAlert severity="warning">Attention - all exists data will be replaced!</MaterialAlert>

            <Box mt={3} mb={3}>
                <input
                    accept="application/json"
                    style={{ display: 'none' }}
                    id="raised-button-file"
                    type="file"
                    onChange={handleFileSelected}
                />
                <label htmlFor="raised-button-file">
                    <Button variant="contained" component="span">
                        Import file
                    </Button>
                </label>
            </Box>

            {
                appDataStore.validationErrors.length > 0 &&
                appDataStore.validationErrors.map((e, index) => (
                    <MaterialAlert
                        key={`validation_error_${index}`}
                        severity="error"
                    >
                        {e}
                    </MaterialAlert>
                ))
            }

            <Alert
                shouldShow={appDataStore.importedOk}
                severity={'success'}
                message={'Imported successfully'}
                onCloseCb={() => appDataStore.cleanImportResult()}
            />

            <Alert
                shouldShow={!!appDataStore.importError}
                severity={'error'}
                message={appDataStore.importError?.description || 'Unknown error'}
                onCloseCb={() => appDataStore.cleanImportResult()}
            />
        </>
    );
});
