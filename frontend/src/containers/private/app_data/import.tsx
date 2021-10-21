import { useEffect } from 'react';
import { bindActionCreators } from 'redux';
import {
    Box,
    Button,
} from '@material-ui/core';
import { Alert as MaterialAlert } from '@material-ui/lab';

import {
    DispatchToPropsFn,
    StateToPropsFn,
} from '../../../shared/types';
import { Alert } from '../../../components/shared';
import {
    ImportActions,
    ImportProps,
    ImportState,
} from './import.types';
import {
    importFromFile,
    cleanAppDataResult,
} from '../../../store/app_data/actions';
import { fetchResources } from '../../../store/resource/actions';
import { fetchRolesVersion } from '../../../store/roles_version/actions';
import { fetchRoles } from '../../../store/role/actions';

export const Import = (props: ImportProps) => {
    useEffect(
        () => {
            if (props.appDataResult.importedOk) {
                props.loadRolesVersion();
                props.loadResources();
                props.loadRoles();
            }
        },
        [props.appDataResult.importedOk],
    );

    const handleFileSelected = (e: any): void => {
        props.importAppDataAction({
            file: e.target.files[0],
        });

        e.target.value = null;
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
                props.appDataResult.appDataResult?.importFileValidationErrors?.length &&
                props.appDataResult.appDataResult?.importFileValidationErrors.map((e, index) => (
                    <MaterialAlert
                        key={`validation_error_${index}`}
                        severity="error"
                    >
                        {e}
                    </MaterialAlert>
                ))
            }

            {/*todo importedOk does not work (state structure mistake?)*/}
            <Alert
                shouldShow={!!props.appDataResult.importedOk}
                severity={'success'}
                message={'Imported successfully'}
                onCloseCb={() => props.cleanAppDataResult()}
            />

            <Alert
                shouldShow={!!props.appDataResult.appDataResult?.importError}
                severity={'error'}
                message={props.appDataResult.appDataResult?.importError?.description || 'Unknown error'}
                onCloseCb={() => props.cleanAppDataResult()}
            />
        </>
    );
};

export const mapDispatchToProps: DispatchToPropsFn<ImportActions> = () => dispatch => ({
    importAppDataAction: bindActionCreators(importFromFile, dispatch),
    cleanAppDataResult: bindActionCreators(cleanAppDataResult, dispatch),

    loadRolesVersion: bindActionCreators(fetchRolesVersion, dispatch),
    loadResources: bindActionCreators(fetchResources, dispatch),
    loadRoles: bindActionCreators(fetchRoles, dispatch),
});

export const mapStateToProps: StateToPropsFn<ImportState> = () => state => ({
    appDataResult: state.appDataResult,
});
