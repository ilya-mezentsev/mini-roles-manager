import { SignUp, SignIn } from '../connected';
import { About, SharedNavigation } from '../../../components/shared';

const routes = [
    {
        path: '/sign-in',
        name: 'Sign-In',
        component: () => <SignIn/>,
    },
    {
        path: '/sign-up',
        name: 'Sign-Up',
        component: () => <SignUp/>,
    },
    {
        path: '/about',
        name: 'About',
        component: () => <About/>,
    },
];

export const Navigation = () => (
    <SharedNavigation
        routes={routes}
        fallbackPath={'/sign-in'}
    />
);
