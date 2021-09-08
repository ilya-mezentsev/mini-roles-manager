import {
    useState,
    useEffect,
    useRef,
} from 'react';
import { bindActionCreators } from 'redux';
import EventEmitter from 'events';

import { DispatchToPropsFn, StateToPropsFn } from '../../../shared/types';
import {
    RolesVersionListActions,
    RolesVersionListState,
    RolesVersionListProps,
} from './list.types';
import {
    cleanDeleteRolesVersionError,
    cleanFetchRolesVersionError,
    cleanUpdateRolesVersionError, deleteRolesVersion,
    updateRolesVersion
} from '../../../store/roles_version/actions';
import { RolesVersion } from '../../../services/api';
import {
    EditRolesVersion,
    RolesVersionList as RolesVersionListComponent
} from '../../../components/private/roles_version';
import { Prompter } from '../../../components/private/shared';
import { Alert } from '../../../components/shared';
import { cleanFetchRolesError, fetchRoles } from '../../../store/role/actions';

export const RolesVersionList = (props: RolesVersionListProps) => {
    const rolesVersionsListRef = useRef(props.rolesVersionResult.list);
    useEffect(() => {
        const currentCount = rolesVersionsListRef.current?.length || 0;
        const updatedCount = props.rolesVersionResult.list?.length || 0;

        // if some roles version was deleted we need to reload roles => so store will not contain irrelevant roles
        if (currentCount > updatedCount) {
            props.loadRolesAction();
        }

        rolesVersionsListRef.current = props.rolesVersionResult.list;

        // eslint-disable-next-line
    }, [props.rolesVersionResult.list]);

    const [deletingRolesVersion, setDeletingRolesVersion] = useState<RolesVersion | null>(null);
    const [editingRolesVersion, setEditingRolesVersion] = useState<RolesVersion | null>(null);

    const e = new EventEmitter();
    const openPrompterEventName = 'prompter:open';
    const openEditResourceEventName = 'edit-roles-version:open';

    const deleteRolesVersion = () => {
        if (deletingRolesVersion) {
            props.deleteRolesVersionAction(deletingRolesVersion.id);
            setDeletingRolesVersion(null);
        }
    };

    const hasAnyError = () => {
        return (
            !!props.rolesVersionResult?.fetchError ||
            !!props.rolesVersionResult?.updateError ||
            !!props.rolesVersionResult?.deleteError ||
            !!props.rolesResult?.fetchError
        );
    };

    const cleanErrors = () => {
        props.cleanFetchRolesVersionsErrorAction();
        props.cleanUpdateRolesVersionErrorAction();
        props.cleanDeleteRolesVersionErrorAction();
        props.cleanFetchRolesErrorAction();
    };

    const errorMessage: () => string = () => {
        return (
            props.rolesVersionResult?.fetchError?.description ||
            props.rolesVersionResult?.updateError?.description ||
            props.rolesVersionResult?.deleteError?.description ||
            'Unknown error'
        );
    };

    return (
        <>
            <RolesVersionListComponent
                rolesVersions={props.rolesVersionResult?.list || []}
                tryEdit={rv => {
                    setEditingRolesVersion(rv);
                    e.emit(openEditResourceEventName);
                }}
                tryDelete={rv => {
                    setDeletingRolesVersion(rv);
                    e.emit(openPrompterEventName);
                }}
            />

            <EditRolesVersion
                openDialogueEventName={openEditResourceEventName}
                eventEmitter={e}
                save={rv => props.updateRolesVersionAction(rv)}
                initialRolesVersion={editingRolesVersion}
            />

            <Prompter
                title={'Roles version deletion'}
                description={'Are you sure you want to delete this roles version? This will cause deletion of all roles linked to this version.'}
                onAgree={() => deleteRolesVersion()}
                onDisagree={() => setDeletingRolesVersion(null)}
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
};

export const mapDispatchToProps: DispatchToPropsFn<RolesVersionListActions> = () => dispatch => ({
    cleanFetchRolesVersionsErrorAction: bindActionCreators(cleanFetchRolesVersionError, dispatch),

    updateRolesVersionAction: bindActionCreators(updateRolesVersion, dispatch),
    cleanUpdateRolesVersionErrorAction: bindActionCreators(cleanUpdateRolesVersionError, dispatch),

    deleteRolesVersionAction: bindActionCreators(deleteRolesVersion, dispatch),
    cleanDeleteRolesVersionErrorAction: bindActionCreators(cleanDeleteRolesVersionError, dispatch),

    loadRolesAction: bindActionCreators(fetchRoles, dispatch),
    cleanFetchRolesErrorAction: bindActionCreators(cleanFetchRolesError, dispatch),
});

export const mapStateToProps: StateToPropsFn<RolesVersionListState> = () => state => ({
    rolesVersionResult: state.rolesVersionResult,
    rolesResult: state.rolesResult,
});
