import { ResourceStore } from './resource.store';

import * as api from '../../services/api';
import { ErrorResponse, SuccessResponse } from '../../services/api/shared';
import { UnknownErrorCode, UnknownErrorDescription } from '../shared/const';
import { RoleStore } from '../role/role.store';

jest.mock('../../services/api');

describe('resource tests', () => {
    let roleStore = new RoleStore();
    let resourceStore = new ResourceStore(roleStore);
    const d = {
        id: 'resource-id',
        title: 'resource-title',
        linksTo: [],
    };

    beforeEach(() => {
        roleStore = new RoleStore();
        resourceStore = new ResourceStore(roleStore);
    });

    it('create resource success', async () => {
        // @ts-ignore
        api.createResource = jest.fn().mockResolvedValue(new SuccessResponse(null));
        // @ts-ignore
        api.resourcesList = jest.fn().mockResolvedValue(new SuccessResponse(null));

        await resourceStore.createResource(d);

        expect(api.createResource).toBeCalledWith(d);
        expect(api.resourcesList).toBeCalled();
        expect(resourceStore.createResourceError).toBeNull();
    });

    it('create resource parsed error', async () => {
        // @ts-ignore
        api.createResource = jest.fn().mockResolvedValue(new ErrorResponse('some-error'));

        await resourceStore.createResource(d);

        expect(api.createResource).toBeCalledWith(d);
        expect(resourceStore.createResourceError).toEqual('some-error');
    });

    it('create resource unknown error', async () => {
        // @ts-ignore
        api.createResource = jest.fn().mockRejectedValue('some-error');

        await resourceStore.createResource(d);

        expect(api.createResource).toBeCalledWith(d);
        expect(resourceStore.createResourceError).toEqual({
            code: UnknownErrorCode,
            description: UnknownErrorDescription,
        });
    });

    it('fetch resources success', async () => {
        // @ts-ignore
        api.resourcesList = jest.fn().mockResolvedValue(new SuccessResponse([d]));

        await resourceStore.fetchResources();

        expect(api.resourcesList).toBeCalled();
        expect(resourceStore.list).toEqual([d]);

        // @ts-ignore
        api.resourcesList = jest.fn().mockResolvedValue(new SuccessResponse(null));

        await resourceStore.fetchResources();

        expect(api.resourcesList).toBeCalled();
        expect(resourceStore.list).toEqual([]);
    });

    it('fetch resources parsed error', async () => {
        // @ts-ignore
        api.resourcesList = jest.fn().mockResolvedValue(new ErrorResponse('some-error'));

        await resourceStore.fetchResources();

        expect(api.resourcesList).toBeCalled();
        expect(resourceStore.fetchResourceError).toEqual('some-error');
    });

    it('fetch resources unknown error', async () => {
        // @ts-ignore
        api.resourcesList = jest.fn().mockRejectedValue('some-error');

        await resourceStore.fetchResources();

        expect(api.resourcesList).toBeCalled();
        expect(resourceStore.fetchResourceError).toEqual({
            code: UnknownErrorCode,
            description: UnknownErrorDescription,
        });
    });

    it('update resource success', async () => {
        // @ts-ignore
        api.resourcesList = jest.fn().mockResolvedValue(new SuccessResponse([d]));

        // add resource to store
        await resourceStore.fetchResources();

        // @ts-ignore
        api.updateResource = jest.fn().mockResolvedValue(new SuccessResponse(null));

        const t = {
            ...d,
            title: 'some-new-title',
        };

        await resourceStore.updateResource(t);

        expect(api.updateResource).toBeCalledWith(t);
        expect(resourceStore.list).toEqual([t]);

        const n = {
            ...d,
            id: 'not-exists-id',
        };

        // update not exists resource
        await resourceStore.updateResource(n);

        expect(api.updateResource).toBeCalledWith(n);
        expect(resourceStore.list).toEqual([t]);
    });

    it('update resource parsed error', async () => {
        // @ts-ignore
        api.updateResource = jest.fn().mockResolvedValue(new ErrorResponse('some-error'));

        await resourceStore.updateResource(d);

        expect(api.updateResource).toBeCalledWith(d);
        expect(resourceStore.updateResourceError).toEqual('some-error');
    });

    it('update resource unknown error', async () => {
        // @ts-ignore
        api.updateResource = jest.fn().mockRejectedValue('some-error');

        await resourceStore.updateResource(d);

        expect(api.updateResource).toBeCalledWith(d);
        expect(resourceStore.updateResourceError).toEqual({
            code: UnknownErrorCode,
            description: UnknownErrorDescription,
        });
    });

    it('delete resource success', async () => {
        // @ts-ignore
        api.resourcesList = jest.fn().mockResolvedValue(new SuccessResponse([d]));

        // add resource to store
        await resourceStore.fetchResources();

        expect(resourceStore.list).toEqual([d]);

        // @ts-ignore
        api.deleteResource = jest.fn().mockResolvedValue(new SuccessResponse(null));

        await resourceStore.deleteResource(d.id);

        expect(api.deleteResource).toBeCalledWith(d.id);
        expect(resourceStore.list).toEqual([]);

        // delete not exists resource
        await resourceStore.deleteResource(d.id);

        expect(api.deleteResource).toBeCalledWith(d.id);
        expect(resourceStore.list).toEqual([]);
    });

    it('delete resource parsed error', async () => {
        // @ts-ignore
        api.deleteResource = jest.fn().mockResolvedValue(new ErrorResponse('some-error'));

        await resourceStore.deleteResource(d.id);

        expect(api.deleteResource).toBeCalledWith(d.id);
        expect(resourceStore.deleteResourceError).toEqual('some-error');
    });

    it('delete resource unknown error', async () => {
        // @ts-ignore
        api.deleteResource = jest.fn().mockRejectedValue('some-error');

        await resourceStore.deleteResource(d.id);

        expect(api.deleteResource).toBeCalledWith(d.id);
        expect(resourceStore.deleteResourceError).toEqual({
            code: UnknownErrorCode,
            description: UnknownErrorDescription,
        });
    });

    it('clean resource actions errors', () => {
        resourceStore.createResourceError = 'some-error' as any;
        resourceStore.fetchResourceError = 'some-error' as any;
        resourceStore.updateResourceError = 'some-error' as any;
        resourceStore.deleteResourceError = 'some-error' as any;

        resourceStore.cleanResourceActionError();

        expect(resourceStore.createResourceError).toBeNull();
        expect(resourceStore.fetchResourceError).toBeNull();
        expect(resourceStore.updateResourceError).toBeNull();
        expect(resourceStore.deleteResourceError).toBeNull();
    });
});
