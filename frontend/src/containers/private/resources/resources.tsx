import { bindActionCreators } from 'redux';
import { Box } from '@material-ui/core';
import { Add } from '@material-ui/icons';
import EventEmitter from 'events';

import { DispatchToPropsFn, StateToPropsFn } from '../../../shared/types';
import { Alert } from '../../../components/shared/';
import {
    cleanCreateResourceError,
    createResource,
} from '../../../store/resource/actions';
import { ResourcesActions, ResourceState, ResourceProps } from './resources.types';
import { ResourcesList } from '../connected';
import { EditResource as CreateResource } from '../../../components/private/resource';

export const Resources = (props: ResourceProps) => {
    const e = new EventEmitter();
    const openDialogueEventName = 'new-resource-dialogue:open';

    return (
        <>
            <Box>
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
            </Box>
            <Box>
                <ResourcesList />
            </Box>

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
});

export const mapStateToProps: StateToPropsFn<ResourceState> = () => state => ({
    resourcesResult: state.resourcesResult,
});
