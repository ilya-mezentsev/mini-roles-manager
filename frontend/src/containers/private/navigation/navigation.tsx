import {
    About,
    SharedAppNavigation,
    NavigationRoute,
} from '../../../components/shared';
import { SecondaryButton } from '../../../components/shared/navigation/navigation.types';

import { AppData } from '../app_data/app_data';
import { Account } from '../account/account';
import { Resources } from '../resources/resources';
import { RolesVersion } from '../roles_version/roles_version';
import { Roles } from '../role/roles';
import { SignOut } from '../sign_out/sign_out';

const routes: NavigationRoute[] = [
    {
        path: '/resources',
        name: 'Resources',
        component: () => <Resources/>
    },
    {
        path: '/roles-versions',
        name: 'Roles Versions',
        component: () => <RolesVersion/>
    },
    {
        path: '/roles',
        name: 'Roles',
        component: () => <Roles/>
    },
    {
        path: '/app-data',
        name: 'App data',
        component: () => <AppData/>,
    },
    {
        path: '/account',
        name: 'Account',
        component: () => <Account/>
    },
    {
        path: '/about',
        name: 'About',
        component: () => <About/>
    },
];

const secondaryButtons: SecondaryButton[] = [
    {
        component: () => <SignOut/>
    }
];

export const Navigation = () => (
    <SharedAppNavigation
        routes={routes}
        secondaryButtons={secondaryButtons}
        fallbackPath="/resources"
    />
);
