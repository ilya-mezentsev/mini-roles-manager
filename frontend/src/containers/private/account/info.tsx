import {
    Card,
    CardContent,
    Typography,
} from '@material-ui/core';

import { StateToPropsFn } from '../../../shared/types';
import { AccountInfoProps, AccountInfoState } from './info.types';

export const Info = (props: AccountInfoProps) => (
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
    </>
)

const makeCreated = (created?: string) => created ? (new Date(created)).toLocaleString() : 'Unknown';

export const mapStateToProps: StateToPropsFn<AccountInfoState> = () => state => ({
    accountInfoResult: state.accountInfoResult,
    rolesResult: state.rolesResult,
    resourcesResult: state.resourcesResult,
});
