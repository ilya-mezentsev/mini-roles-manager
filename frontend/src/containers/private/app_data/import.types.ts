import { ImportFile } from '../../../services/api';
import { AppDataResult } from '../../../store/app_data/app_data.types';
import { APIError } from '../../../services/api/shared';

export interface ImportActions {
    importAppDataAction: (d: ImportFile) => void;
    cleanAppDataResult: () => void;

    loadRolesVersion: () => void;
    loadResources: () => void;
    loadRoles: () => void;
}

export interface ImportState {
    appDataResult: AppDataResult<APIError>;
}

export type ImportProps = ImportActions & ImportState;
