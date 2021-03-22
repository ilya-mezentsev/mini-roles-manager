import { About, SharedNavigation, NavigationRoute } from '../../../components/shared';
import { Resources } from '../resources/resources';
import { Roles } from '../roles/roles';

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
        path: '/about',
        name: 'About',
        component: () => <About/>
    },
];

export const Navigation = () => (
    <SharedNavigation
        routes={routes}
        fallbackPath={'/resources'}
    />
);
