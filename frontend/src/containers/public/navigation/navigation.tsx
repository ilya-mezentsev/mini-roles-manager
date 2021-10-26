import { About, SharedAppNavigation } from '../../../components/shared';
import { SignIn } from '../sign_in/sign_in';
import { SignUp } from '../sign_up/sign_up';

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
