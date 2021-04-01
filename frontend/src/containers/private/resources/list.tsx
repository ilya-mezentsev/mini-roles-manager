import { useState, useEffect, useRef } from 'react';
import { bindActionCreators } from 'redux';
import * as _ from 'lodash';
import EventEmitter from 'events';

import { Alert } from '../../../components/shared';
import { Prompter } from '../../../components/private/shared/prompter';
import {
    EditResource,
    ResourcesList as ResourcesListComponent,
} from '../../../components/private/resource';
import { DispatchToPropsFn, StateToPropsFn } from '../../../shared/types';
import {
    cleanDeleteResourceError,
    cleanLoadResourcesError,
    cleanUpdateResourceError,
    deleteResource,
    loadResources,
    updateResource,
} from '../../../store/resource/actions';
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
    }, [props.resourcesResult.list]);

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
            !!props.resourcesResult?.deleteError
        );
    };

    const cleanError = () => {
        if (props.resourcesResult?.fetchError) {
            props.cleanLoadResourcesError();
        } else if (props.resourcesResult?.updateError) {
            props.cleanUpdateResourceErrorAction();
        } else if (props.resourcesResult?.deleteError) {
            props.cleanDeleteResourceErrorAction();
        }
    };

    const errorMessage: () => string = () => {
        return (
            props.resourcesResult?.fetchError?.description ||
            props.resourcesResult?.updateError?.description ||
            props.resourcesResult?.deleteError?.description ||
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
                initialResourceId={editingResource?.id}
                initialResourceTitle={editingResource?.title}
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
                onCloseCb={() => cleanError()}
            />
        </>
    );
}

export const mapDispatchToProps: DispatchToPropsFn<ResourcesListActions> = () => dispatch => ({
    updateResourceAction: bindActionCreators(updateResource, dispatch),
    cleanUpdateResourceErrorAction: bindActionCreators(cleanUpdateResourceError, dispatch),

    deleteResourceAction: bindActionCreators(deleteResource, dispatch),
    cleanDeleteResourceErrorAction: bindActionCreators(cleanDeleteResourceError, dispatch),

    loadResourcesAction: bindActionCreators(loadResources, dispatch),
    cleanLoadResourcesError: bindActionCreators(cleanLoadResourcesError, dispatch),
});

export const mapStateToProps: StateToPropsFn<ResourcesListState> = () => state => ({
    resourcesResult: state.resourcesResult,
});
