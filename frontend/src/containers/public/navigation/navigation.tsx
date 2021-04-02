import { SignUp, SignIn } from '../connected';
import { About, SharedAppNavigation } from '../../../components/shared';

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
    <SharedAppNavigation
        routes={routes}
        fallbackPath={'/sign-in'}
    />
);
