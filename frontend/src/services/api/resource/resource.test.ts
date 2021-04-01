import * as request from '../shared/request';
import {
    createResource,
    resourcesList,
    updateResource,
    deleteResource,
} from './resource';
import { EditableResource } from './resource.types';
import { APIResponseStatus, ErrorResponse, SuccessResponse } from '../shared';

jest.mock('../shared/request');

describe('resource api tests', () => {
    it('create resource success', async () => {
        const d: EditableResource = {
            id: 'some-resource-id-1',
            title: 'SomeResourceTitle',
        };
        // @ts-ignore
        request.POST = jest.fn().mockResolvedValue(null);

        const response = await createResource(d);

        expect(request.POST).toBeCalledWith(
            '/resource',
            {
                resource: d
            },
        );
        expect(response).toBeInstanceOf(SuccessResponse);
        expect(response.isOk()).toBeTruthy();
        expect(response.data()).toBeNull();
    });

    it('create resource error', async () => {
        const d: EditableResource = {
            id: 'some-resource-id-1',
            title: 'SomeResourceTitle',
        };
        // @ts-ignore
        request.POST = jest.fn().mockResolvedValue({
            status: APIResponseStatus.ERROR,
            data: 'some-error',
        });

        const response = await createResource(d);

        expect(request.POST).toBeCalledWith(
            '/resource',
            {
                resource: d
            },
        );
        expect(response).toBeInstanceOf(ErrorResponse);
        expect(response.isOk()).toBeFalsy();
        expect(response.data()).toEqual('some-error');
    });

    it('resources list success', async () => {
        const d: EditableResource[] = [
            { id: 'some-resource-1' },
            { id: 'some-resource-2' },
        ];
        // @ts-ignore
        request.GET = jest.fn().mockResolvedValue({
            status: APIResponseStatus.OK,
            data: d,
        });

        const response = await resourcesList();

        expect(request.GET).toBeCalledWith('/resources');
        expect(response).toBeInstanceOf(SuccessResponse);
        expect(response.isOk()).toBeTruthy();
        expect(response.data()).toEqual(d);
    });

    it('resources list error', async () => {
        // @ts-ignore
        request.GET = jest.fn().mockResolvedValue({
            status: APIResponseStatus.ERROR,
            data: 'some-error',
        });

        const response = await resourcesList();

        expect(request.GET).toBeCalledWith('/resources');
        expect(response).toBeInstanceOf(ErrorResponse);
        expect(response.isOk()).toBeFalsy();
        expect(response.data()).toEqual('some-error');
    });

    it('update resource success', async () => {
        const d: EditableResource = {
            id: 'some-resource-id-1',
            title: 'SomeResourceTitle',
        };
        // @ts-ignore
        request.PATCH = jest.fn().mockResolvedValue(null);

        const response = await updateResource(d);

        expect(request.PATCH).toBeCalledWith(
            '/resource',
            {
                resource: d
            },
        );
        expect(response).toBeInstanceOf(SuccessResponse);
        expect(response.isOk()).toBeTruthy();
        expect(response.data()).toBeNull();
    });

    it('update resource error', async () => {
        const d: EditableResource = {
            id: 'some-resource-id-1',
            title: 'SomeResourceTitle',
        };
        // @ts-ignore
        request.PATCH = jest.fn().mockResolvedValue({
            status: APIResponseStatus.ERROR,
            data: 'some-error',
        });

        const response = await updateResource(d);

        expect(request.PATCH).toBeCalledWith(
            '/resource',
            {
                resource: d
            },
        );
        expect(response).toBeInstanceOf(ErrorResponse);
        expect(response.isOk()).toBeFalsy();
        expect(response.data()).toEqual('some-error');
    });

    it('delete resource success', async () => {
        const d = 'some-resource-1';
        // @ts-ignore
        request.DELETE = jest.fn().mockResolvedValue(null);

        const response = await deleteResource(d);

        expect(request.DELETE).toBeCalledWith(`/resource/${d}`);
        expect(response).toBeInstanceOf(SuccessResponse);
        expect(response.isOk()).toBeTruthy();
        expect(response.data()).toBeNull();
    });

    it('delete resource error', async () => {
        const d = 'some-resource-1';
        // @ts-ignore
        request.DELETE = jest.fn().mockResolvedValue({
            status: APIResponseStatus.ERROR,
            data: 'some-error',
        });

        const response = await deleteResource(d);

        expect(request.DELETE).toBeCalledWith(`/resource/${d}`);
        expect(response).toBeInstanceOf(ErrorResponse);
        expect(response.isOk()).toBeFalsy();
        expect(response.data()).toEqual('some-error');
    });
});
