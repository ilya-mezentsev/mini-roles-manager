import { About, SharedAppNavigation, NavigationRoute } from '../../../components/shared';
import {
    Resources,
    Roles,
    Account,
    SignOut,
} from '../connected';
import { SecondaryButton } from '../../../components/shared/navigation/navigation.types';

const routes: NavigationRoute[] = [
    {
        path: '/resources',
        name: 'Resources',
        component: () => <Resources/>
    },
    {
        path: '/roles',
        name: 'Roles',
        component: () => <Roles/>
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
