import { removeLeadingAndTrailingSlashes } from '../helpers';

describe('shared helpers tests', () => {
    it('remove slashes', () => {
        const _ = removeLeadingAndTrailingSlashes;

        expect(_('/path')).toEqual('path');
        expect(_('path/')).toEqual('path');
        expect(_('////path//////')).toEqual('path');
    });
});
