import { Dispatch } from 'redux';

import { ACTIONS } from './action_types';
import {
    ImportFile,
    importFromFile as importFromFileAPI,
} from '../../services/api/';

export function importFromFile(d: ImportFile): (dispatch: Dispatch) => Promise<void> {
    return async dispatch => {
        try {
            const response = await importFromFileAPI(d);

            if (response.isOk()) {
                dispatch({
                    type: ACTIONS.SUCCESS_IMPORT_APP_DATA,
                });
            } else {
                dispatch({
                    type: ACTIONS.FAILED_IMPORT_APP_DATA,
                    appDataResult: {
                        importError: response.data(),
                    },
                });
            }
        } catch (e) {
            dispatch({
                type: ACTIONS.FAILED_TO_PERFORM_APP_DATA_IMPORTING,
                appDataResult: {
                    importError: e,
                },
            });
        }
    };
}

export function cleanAppDataResult(): (dispatch: Dispatch) => void {
    return dispatch => {
        dispatch({
            type: ACTIONS.CLEAN_IMPORT_APP_DATA_RESULT,
        });
    };
}
