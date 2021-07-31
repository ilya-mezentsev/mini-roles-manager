import { bindActionCreators } from 'redux';
import { Box } from '@material-ui/core';
import { Add } from '@material-ui/icons';
import EventEmitter from 'events';

import {
    DispatchToPropsFn,
    StateToPropsFn,
} from '../../../shared/types';
import {
    RolesVersionActions,
    RolesVersionState,
    RolesVersionProps,
} from './roles_version.types';
import {
    cleanCreateRolesVersionError,
    createRolesVersion,
} from '../../../store/roles_version/actions';
import { RolesVersionList } from '../connected';
import {
    EditRolesVersion as CreateRolesVersion,
} from '../../../components/private/roles_version';
import { Alert } from '../../../components/shared';

export const RolesVersion = (props: RolesVersionProps) => {
    const e = new EventEmitter();
    const openDialogueEventName = 'new-roles-version-dialogue:open';

    return (
        <>
            <Box>
                <h1>
                    Roles Versions
                    <Add
                        color="primary"
                        cursor="pointer"
                        fontSize="large"
                        titleAccess="Add new roles version"
                        onClick={() => e.emit(openDialogueEventName)}
                    />
                </h1>
            </Box>

            <Box>
                <RolesVersionList/>
            </Box>

            <CreateRolesVersion
                openDialogueEventName={openDialogueEventName}
                eventEmitter={e}
                save={rv => props.createRolesVersionAction(rv)}
            />

            <Alert
                shouldShow={!!props.rolesVersionResult?.createError}
                severity={'error'}
                message={props.rolesVersionResult?.createError?.description || 'Unknown error'}
                onCloseCb={() => props.cleanCreateRolesVersionErrorAction()}
            />
        </>
    );
};

export const mapDispatchToProps: DispatchToPropsFn<RolesVersionActions> = () => dispatch => ({
    createRolesVersionAction: bindActionCreators(createRolesVersion, dispatch),
    cleanCreateRolesVersionErrorAction: bindActionCreators(cleanCreateRolesVersionError, dispatch),
});

export const mapStateToProps: StateToPropsFn<RolesVersionState> = () => state => ({
    rolesVersionResult: state.rolesVersionResult,
});
