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
                        Roles count: <b>{props.rolesResult?.list?.length ?? 'Unknown'}</b>
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

export const mapStateToProps: StateToPropsFn<AccountInfoState> = () => state => ({
    accountInfoResult: state.accountInfoResult,
    rolesResult: state.rolesResult,
    resourcesResult: state.resourcesResult,
});

export const mapDispatchToProps: DispatchToPropsFn<AccountInfoActions> = () => dispatch => ({
    cleanFetchInfoErrorAction: bindActionCreators(cleanFetchInfoError, dispatch),
});
