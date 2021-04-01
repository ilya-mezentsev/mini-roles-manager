import { useEffect } from 'react';
import { bindActionCreators } from 'redux';
import { Container } from '@material-ui/core';
import { Add } from '@material-ui/icons';
import EventEmitter from 'events';

import { DispatchToPropsFn, StateToPropsFn } from '../../../shared/types';
import { Alert } from '../../../components/shared/';
import {
    cleanCreateResourceError,
    createResource,
    loadResources,
} from '../../../store/resource/actions';
import { ResourcesActions, ResourceState, ResourceProps } from './resources.types';
import { ResourcesList } from '../connected';
import { EditResource as CreateResource } from '../../../components/private/resource';

export const Resources = (props: ResourceProps) => {
    useEffect(() => {
        if (!props.resourcesResult.list) {
            props.loadResourcesAction();
        }
    }, []);

    const e = new EventEmitter();
    const openDialogueEventName = 'new-resource-dialogue:open';

    return (
        <>
            <Container>
                <h1>
                    Resources
                    <Add
                        color="primary"
                        cursor="pointer"
                        fontSize="large"
                        titleAccess="Add new resource"
                        onClick={() => e.emit(openDialogueEventName)}
                    />
                </h1>
            </Container>
            <Container>
                <ResourcesList />
            </Container>

            <CreateResource
                eventEmitter={e}
                openDialogueEventName={openDialogueEventName}
                save={r => props.createResourceAction(r)}
            />

            <Alert
                shouldShow={!!props.resourcesResult?.createError}
                severity={'error'}
                message={props.resourcesResult?.createError?.description || 'Unknown error'}
                onCloseCb={() => props.cleanCreateResourceErrorAction()}
            />
        </>
    );
}

export const mapDispatchToProps: DispatchToPropsFn<ResourcesActions> = () => dispatch => ({
    createResourceAction: bindActionCreators(createResource, dispatch),
    cleanCreateResourceErrorAction: bindActionCreators(cleanCreateResourceError, dispatch),

    loadResourcesAction: bindActionCreators(loadResources, dispatch),
});

export const mapStateToProps: StateToPropsFn<ResourceState> = () => state => ({
    resourcesResult: state.resourcesResult,
});
