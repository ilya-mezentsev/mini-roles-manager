export interface Route {
    path: string;
    name: string;
    component: () => JSX.Element
}

export interface NavigationProps {
    routes: Route[];
    fallbackPath: string;
    size?: 'small' | 'large';
}
