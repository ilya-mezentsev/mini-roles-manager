import * as request from '../shared/request';

import { importFromFile } from './app_data';
import { SuccessResponse } from '../shared';

jest.mock('../shared/request');

describe('app_data tests', () => {
    it('import app data', async () => {
        const d = {
            file: new File([], 'name'),
        };
        const fd = new FormData();
        fd.append('app_data_file', d.file);

        // @ts-ignore
        request.POST = jest.fn().mockResolvedValue(null);

        const response = await importFromFile(d);

        expect(request.POST).toBeCalledWith('/app-data/import', fd, true);
        expect(response).toBeInstanceOf(SuccessResponse);
        expect(response.isOk()).toBeTruthy();
        expect(response.data()).toBeNull();
    });
});
