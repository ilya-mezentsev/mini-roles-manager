import { useState, useEffect, useRef } from 'react';
import { bindActionCreators } from 'redux';
import * as _ from 'lodash';
import EventEmitter from 'events';

import { Alert } from '../../../components/shared';
import { Prompter } from '../../../components/private/shared';
import {
    EditResource,
    ResourcesList as ResourcesListComponent,
} from '../../../components/private/resource';
import { DispatchToPropsFn, StateToPropsFn } from '../../../shared/types';
import {
    cleanDeletedResourceId,
    cleanDeleteResourceError,
    cleanLoadResourcesError,
    cleanUpdateResourceError,
    deleteResource,
    fetchResources,
    updateResource,
} from '../../../store/resource/actions';
import { fetchRoles, cleanFetchRolesError } from '../../../store/role/actions';
import { ResourcesListActions, ResourcesListState, ResourcesListProps } from './list.types';
import { Resource } from '../../../services/api';

export const ResourcesList = (props: ResourcesListProps) => {
    const resourcesListRef = useRef(props.resourcesResult.list);
    useEffect(() => {
        const differentResources = _.differenceWith(
            props.resourcesResult.list,
            resourcesListRef.current as any,
            _.isEqual,
        );

        if (differentResources.length > 0) {
            const gotNewResource = differentResources.some(d => !d.permissions);

            if (gotNewResource) {
                props.loadResourcesAction();
            }

            resourcesListRef.current = props.resourcesResult.list;
        }
        // eslint-disable-next-line
    }, [props.resourcesResult.list]);

    useEffect(() => {
        if (props.resourcesResult?.deletedResourceId) {
            props.loadRolesAction();
            props.cleanDeletedResourceIdAction();
        }
        // eslint-disable-next-line
    }, [props.resourcesResult?.deletedResourceId]);

    const [deletingResource, setDeletingResource] = useState<Resource | null>(null);
    const [editingResource, setEditingResource] = useState<Resource | null>(null);

    const e = new EventEmitter();
    const openPrompterEventName = 'prompter:open';
    const openEditResourceEventName = 'edit-resource:open';

    const deleteResource = () => {
        if (deletingResource) {
            props.deleteResourceAction(deletingResource.id);
            setDeletingResource(null);
        }
    };

    const hasAnyError: () => boolean = () => {
        return (
            !!props.resourcesResult?.fetchError ||
            !!props.resourcesResult?.updateError ||
            !!props.resourcesResult?.deleteError ||
            !!props.rolesResult?.fetchError
        );
    };

    const cleanErrors = () => {
        props.cleanLoadResourcesError();
        props.cleanUpdateResourceErrorAction();
        props.cleanDeleteResourceErrorAction();
        props.cleanFetchRolesErrorAction();
    };

    const errorMessage: () => string = () => {
        return (
            props.resourcesResult?.fetchError?.description ||
            props.resourcesResult?.updateError?.description ||
            props.resourcesResult?.deleteError?.description ||
            props.rolesResult?.fetchError?.description ||
            'Unknown error'
        );
    };

    return (
        <>
            <ResourcesListComponent
                resources={props.resourcesResult.list || []}
                tryEdit={r => {
                    setEditingResource(r);
                    e.emit(openEditResourceEventName);
                }}
                tryDelete={r => {
                    setDeletingResource(r);
                    e.emit(openPrompterEventName);
                }}
            />

            <EditResource
                openDialogueEventName={openEditResourceEventName}
                eventEmitter={e}
                save={r => props.updateResourceAction(r)}
                initialResource={editingResource}
            />

            <Prompter
                title="Resource deletion"
                description="Are you sure you want to delete this resource?"
                onAgree={() => deleteResource()}
                onDisagree={() => setDeletingResource(null)}
                openDialogueEventName={openPrompterEventName}
                eventEmitter={e}
            />

            <Alert
                shouldShow={hasAnyError()}
                severity={'error'}
                message={errorMessage()}
                onCloseCb={() => cleanErrors()}
            />
        </>
    );
}

export const mapDispatchToProps: DispatchToPropsFn<ResourcesListActions> = () => dispatch => ({
    updateResourceAction: bindActionCreators(updateResource, dispatch),
    cleanUpdateResourceErrorAction: bindActionCreators(cleanUpdateResourceError, dispatch),

    deleteResourceAction: bindActionCreators(deleteResource, dispatch),
    cleanDeleteResourceErrorAction: bindActionCreators(cleanDeleteResourceError, dispatch),
    cleanDeletedResourceIdAction: bindActionCreators(cleanDeletedResourceId, dispatch),

    loadResourcesAction: bindActionCreators(fetchResources, dispatch),
    cleanLoadResourcesError: bindActionCreators(cleanLoadResourcesError, dispatch),

    loadRolesAction: bindActionCreators(fetchRoles, dispatch),
    cleanFetchRolesErrorAction: bindActionCreators(cleanFetchRolesError, dispatch),
});

export const mapStateToProps: StateToPropsFn<ResourcesListState> = () => state => ({
    resourcesResult: state.resourcesResult,
    rolesResult: state.rolesResult,
});
