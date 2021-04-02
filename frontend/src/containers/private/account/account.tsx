import Grid from '@material-ui/core/Grid';
import InfoIcon from '@material-ui/icons/Info';
import LockOpenIcon from '@material-ui/icons/LockOpen';
import CodeIcon from '@material-ui/icons/Code';
import {
    Redirect,
    Route,
    Switch,
} from 'react-router-dom';
import { Info } from './info';
import { Credentials } from './credentials';
import { APIKey } from './api_key';
import { ListItemRoute } from '../../../components/shared/navigation/navigation.types';
import { SharedListNavigation } from '../../../components/shared';

const fallbackPath = '/account/info';
const routes: (ListItemRoute & { component: () => JSX.Element })[] = [
    {
        path: fallbackPath,
        name: 'Info',
        component: () => <Info/>,
        iconComponent: () => <InfoIcon/>,
    },
    {
        path: '/account/credentials',
        name: 'Credentials',
        component: () => <Credentials/>,
        iconComponent: () => <LockOpenIcon/>,
    },
    {
        path: '/account/api-key',
        name: 'API Key',
        component: () => <APIKey/>,
        iconComponent: () => <CodeIcon/>,
    },
];

export const Account = () => {
    return (
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
    );
};
