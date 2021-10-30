import Grid from '@material-ui/core/Grid';
import InfoIcon from '@material-ui/icons/Info';
import LockOpenIcon from '@material-ui/icons/LockOpen';
import CodeIcon from '@material-ui/icons/Code';
import SecurityIcon from '@material-ui/icons/Security';
import {
    Redirect,
    Route,
    Switch,
} from 'react-router-dom';

import { ListItemRoute } from '../../../components/shared/navigation/navigation.types';
import { SharedListNavigation } from '../../../components/shared';
import { Credentials as AccountCredentials } from './credentials';
import { Info as AccountInfo } from './info';
import { ApiKey as AccountApiKey } from './api_key';
import { CheckPermissions } from './check_permissions';

const fallbackPath = '/account/info';
const routes: (ListItemRoute & { component: () => JSX.Element })[] = [
    {
        path: fallbackPath,
        name: 'Info',
        component: () => <AccountInfo/>,
        iconComponent: () => <InfoIcon/>,
    },
    {
        path: '/account/credentials',
        name: 'Credentials',
        component: () => <AccountCredentials/>,
        iconComponent: () => <LockOpenIcon/>,
    },
    {
        path: '/account/api-key',
        name: 'API Key',
        component: () => <AccountApiKey/>,
        iconComponent: () => <CodeIcon/>,
    },
    {
        path: '/account/check-permissions',
        name: 'Check Permissions',
        component: () => <CheckPermissions/>,
        iconComponent: () => <SecurityIcon/>
    },
];

export const Account = () => (
    <>
        <Grid container spacing={3}>
            <Grid item xs={4}>
                <SharedListNavigation routes={routes}/>
            </Grid>

            <Grid item xs={8}>
                <Switch>
                    {
                        routes.map((r, index) => (
                            <Route path={r.path} key={`account_route_${index}`}>
                                { <r.component/> }
                            </Route>
                        ))
                    }

                    <Route path={"*"}>
                        <Redirect to={fallbackPath} />
                    </Route>
                </Switch>
            </Grid>
        </Grid>
    </>
);
