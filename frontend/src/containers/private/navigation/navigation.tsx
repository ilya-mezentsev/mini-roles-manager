import { About, SharedAppNavigation, NavigationRoute } from '../../../components/shared';
import { Resources, SignOut } from '../connected';
import { Roles } from '../roles/roles';
import { SecondaryButton } from '../../../components/shared/navigation/navigation.types';
import { Account } from '../account/account';

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
        fallbackPath={'/resources'}
        size="large"
    />
);
