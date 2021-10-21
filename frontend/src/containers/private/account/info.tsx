import { useEffect } from 'react';
import { bindActionCreators } from 'redux';
import {
    Card,
    CardContent,
    Typography,
} from '@material-ui/core';

import {
    DispatchToPropsFn,
    StateToPropsFn,
} from '../../../shared/types';
import {
    AccountInfoActions,
    AccountInfoProps,
    AccountInfoState,
} from './info.types';
import { cleanFetchInfoError } from '../../../store/account_info/actions';
import { Alert } from '../../../components/shared';
import { Role } from '../../../services/api';


export const Info = (props: AccountInfoProps) => {
    useEffect(() => {
        props.cleanFetchInfoErrorAction();

        return () => props.cleanFetchInfoErrorAction();
        // eslint-disable-next-line
    }, []);

    return (
        <>
            <h2>Account Info:</h2>
            <Card>
                <CardContent>
                    <Typography variant="body2" component="p">
                        Login: <b>{props.accountInfoResult?.info.login || 'Unknown'}</b>
                    </Typography>
                    <Typography variant="body2" component="p">
                        Created: <b>{makeCreated(props.accountInfoResult?.info.created)}</b>
                    </Typography>
                    <Typography variant="body2" component="p">
                        Roles count: <b>{uniqueRoleCount(props.rolesResult?.list || []) || 'Unknown'}</b>
                    </Typography>
                    <Typography variant="body2" component="p">
                        Resources count: <b>{props.resourcesResult?.list?.length ?? 'Unknown'}</b>
                    </Typography>
                </CardContent>
            </Card>

            <Alert
                shouldShow={!!props.accountInfoResult?.fetchInfoError}
                severity="error"
                message={props.accountInfoResult?.fetchInfoError?.description || 'Unknown error'}
                onCloseCb={() => props.cleanFetchInfoErrorAction()}
            />
        </>
    );
}

const makeCreated = (created?: string) => created ? (new Date(created)).toLocaleString() : 'Unknown';
// Count roles with unique role id
const uniqueRoleCount = (roles: Role[]) => (new Set(roles.map(r => r.id))).size;

export const mapStateToProps: StateToPropsFn<AccountInfoState> = () => state => ({
    accountInfoResult: state.accountInfoResult,
    rolesResult: state.rolesResult,
    resourcesResult: state.resourcesResult,
});

export const mapDispatchToProps: DispatchToPropsFn<AccountInfoActions> = () => dispatch => ({
    cleanFetchInfoErrorAction: bindActionCreators(cleanFetchInfoError, dispatch),
});
