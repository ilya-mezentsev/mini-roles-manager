import {
    Redirect,
    Route,
    Switch,
} from 'react-router';
import LiveHelpIcon from '@material-ui/icons/LiveHelp';
import PersonIcon from '@material-ui/icons/Person';
import NatureIcon from '@material-ui/icons/Nature';
import CodeIcon from '@material-ui/icons/Code';
import MergeTypeIcon from '@material-ui/icons/MergeType';
import Grid from '@material-ui/core/Grid';

import { ListItemRoute } from '../navigation/navigation.types';
import { SharedListNavigation } from '../../../components/shared';
import { InfoDesc } from './info.desc';
import { ResourceDesc } from './resource.desc';
import { RolesVersionDesc } from './roles_version.desc';
import { RoleDesc } from './role.desc';
import { ApiDesc } from './api.desc';


const fallbackPath = '/about/info';
const routes: (ListItemRoute & { component: () => JSX.Element })[] = [
    {
        path: fallbackPath,
        name: 'Info',
        component: () => <InfoDesc/>,
        iconComponent: () => <LiveHelpIcon/>,
    },
    {
        path: '/about/resource',
        name: 'Resource',
        component: () => <ResourceDesc/>,
        iconComponent: () => <NatureIcon/>,
    },
    {
        path: '/about/roles-version',
        name: 'Roles Version',
        component: () => <RolesVersionDesc/>,
        iconComponent: () => <MergeTypeIcon/>,
    },
    {
        path: '/about/role',
        name: 'Role',
        component: () => <RoleDesc/>,
        iconComponent: () => <PersonIcon/>,
    },
    {
        path: '/about/api',
        name: 'API',
        component: () => <ApiDesc/>,
        iconComponent: () => <CodeIcon/>,
    },
];

export const About = () => (
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
