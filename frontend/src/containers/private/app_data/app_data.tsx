import {
    Redirect,
    Route,
    Switch,
} from 'react-router';
import GetAppIcon from '@material-ui/icons/GetApp';
import PublishIcon from '@material-ui/icons/Publish';
import Grid from '@material-ui/core/Grid';

import { ListItemRoute } from '../../../components/shared/navigation/navigation.types';
import { SharedListNavigation } from '../../../components/shared';
import { Import } from '../connected';
import { Export } from './export';

const fallbackPath = '/app-data/export';
const routes: (ListItemRoute & { component: () => JSX.Element })[] = [
    {
        path: fallbackPath,
        name: 'Export',
        component: () => <Export/>,
        iconComponent: () => <GetAppIcon/>,
    },
    {
        path: '/app-data/import',
        name: 'Import',
        component: () => <Import/>,
        iconComponent: () => <PublishIcon/>
    },
];

export const AppData = () => (
    <>
        <Grid container spacing={3}>
            <Grid item xs={4}>
                <SharedListNavigation routes={routes}/>
            </Grid>

            <Grid item xs={8}>
                <Switch>
                    {
                        routes.map((r, index) => (
                            <Route path={r.path} key={`files_route_${index}`}>
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
