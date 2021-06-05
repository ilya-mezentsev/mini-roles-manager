import { removeLeadingAndTrailingSlashes, makeQueryParams } from '../helpers';

describe('shared helpers tests', () => {
    it('remove slashes', () => {
        const _ = removeLeadingAndTrailingSlashes;

        expect(_('/path')).toEqual('path');
        expect(_('path/')).toEqual('path');
        expect(_('////path//////')).toEqual('path');
    });

    it('make query params', () => {
        expect(makeQueryParams({
            foo: 'bar',
            baz: 'xyz',
        })).toEqual('?foo=bar&baz=xyz');
    });

    it('make query params from empty params', () => {
        expect(makeQueryParams()).toEqual('');
    });
});
