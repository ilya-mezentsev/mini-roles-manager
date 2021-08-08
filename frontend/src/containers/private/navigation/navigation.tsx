import {
    About,
    SharedAppNavigation,
    NavigationRoute,
} from '../../../components/shared';
import {
    Resources,
    RolesVersion,
    Roles,
    Account,
    SignOut,
} from '../connected';
import { SecondaryButton } from '../../../components/shared/navigation/navigation.types';
import { AppData } from '../app_data/app_data';

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
