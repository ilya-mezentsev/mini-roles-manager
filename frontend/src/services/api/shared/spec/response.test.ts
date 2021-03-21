import {
    ErrorResponse,
    parseResponse,
    errorResponseOrDefault,
    SuccessResponse, errorOrSuccessResponse,
} from '../response';
import { APIError, APIResponseStatus } from '../response.types';

describe('shared response tests', () => {
    it('errorResponseOrDefault default', () => {
        const response = errorResponseOrDefault(null);

        expect(response).toBeInstanceOf(SuccessResponse);
        expect(response.data()).toBeNull();
        expect(response.isOk()).toBeTruthy();
    });

    it('errorResponseOrDefault error', () => {
        const d: APIError = {
            code: 'some-code',
            description: 'Some description',
        };
        const response = errorResponseOrDefault({
            status: APIResponseStatus.ERROR,
            data: d,
        });

        expect(response).toBeInstanceOf(ErrorResponse);
        expect(response.isOk()).toBeFalsy();
        expect(response.data()).toEqual(d);
    });

    it('parseResponse success', () => {
        const response = parseResponse(
            {
                status: APIResponseStatus.OK,
                data: 'foo',
            },
        );

        expect(response).toBeInstanceOf(SuccessResponse);
        expect(response.data()).toEqual('foo');
        expect(response.isOk()).toBeTruthy();
    });

    it('parseResponse default', () => {
        const response = parseResponse(null);

        expect(response).toBeInstanceOf(SuccessResponse);
        expect(response.data()).toEqual(null);
        expect(response.isOk()).toBeTruthy();
    });

    it('parseResponse error', () => {
        const d: APIError = {
            code: 'some-code',
            description: 'Some description',
        };
        const response = parseResponse(
            {
                status: APIResponseStatus.ERROR,
                data: d,
            },
        );

        expect(response).toBeInstanceOf(ErrorResponse);
        expect(response.isOk()).toBeFalsy();
        expect(response.data()).toEqual(d);
    });

    it('errorOrSuccessResponse success', () => {
        const response = errorOrSuccessResponse(
            {
                status: APIResponseStatus.OK,
                data: 'foo',
            },
        );

        expect(response).toBeInstanceOf(SuccessResponse);
        expect(response.data()).toEqual('foo');
        expect(response.isOk()).toBeTruthy();
    });

    it('errorOrSuccessResponse error', () => {
        const d: APIError = {
            code: 'some-code',
            description: 'Some description',
        };
        const response = parseResponse(
            {
                status: APIResponseStatus.ERROR,
                data: d,
            },
        );

        expect(response).toBeInstanceOf(ErrorResponse);
        expect(response.isOk()).toBeFalsy();
        expect(response.data()).toEqual(d);
    });
});
