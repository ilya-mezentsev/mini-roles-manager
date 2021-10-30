import { useEffect } from 'react';
import { observer } from 'mobx-react-lite';
import {
    Card,
    CardContent,
    Typography,
} from '@material-ui/core';

import { Alert } from '../../../components/shared';
import { Role } from '../../../services/api';
import {
    accountInfoStore,
    roleStore,
    resourceStore,
} from '../../../store';

export const Info = observer(() => {
    useEffect(() => {
        accountInfoStore.cleanAccountInfoActionErrors();

        return () => accountInfoStore.cleanAccountInfoActionErrors();
        // eslint-disable-next-line
    }, []);

    return (
        <>
            <h2>Account Info:</h2>
            <Card>
                <CardContent>
                    <Typography variant="body2" component="p">
                        Login: <b>{accountInfoStore.info?.login || 'Unknown'}</b>
                    </Typography>
                    <Typography variant="body2" component="p">
                        Created: <b>{makeCreated(accountInfoStore.info?.created)}</b>
                    </Typography>
                    <Typography variant="body2" component="p">
                        Roles count: <b>{uniqueRoleCount(roleStore.list)}</b>
                    </Typography>
                    <Typography variant="body2" component="p">
                        Resources count: <b>{resourceStore.list.length}</b>
                    </Typography>
                </CardContent>
            </Card>

            <Alert
                shouldShow={!!accountInfoStore.fetchInfoError}
                severity="error"
                message={accountInfoStore.fetchInfoError?.description || 'Unknown error'}
                onCloseCb={() => accountInfoStore.cleanAccountInfoActionErrors()}
            />
        </>
    );
});

const makeCreated = (created?: string) => created ? (new Date(created)).toLocaleString() : 'Unknown';
// Count roles with unique role id
const uniqueRoleCount = (roles: Role[]) => (new Set(roles.map(r => r.id))).size;
