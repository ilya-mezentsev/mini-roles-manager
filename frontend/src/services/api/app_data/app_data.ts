import { ImportFile } from './app_data.types';
import {
    APIError,
    EmptyAPIResponse,
    ErrorAPIResponse,
    errorResponseOrDefault,
    ParsedAPIResponse,
    POST,
} from '../shared';

export async function importFromFile(d: ImportFile): Promise<ParsedAPIResponse<APIError | null>> {
    const fd = new FormData();
    fd.append('app_data_file', d.file);

    const response = await POST<ErrorAPIResponse | EmptyAPIResponse>('/app-data/import', fd, true);

    return errorResponseOrDefault(response);
}
