export interface Route {
    path: string;
    name: string;
    component: () => JSX.Element;
    iconComponent?: () => JSX.Element;
}

export interface ListItemRoute {
    path: string;
    name: string;
    iconComponent?: () => JSX.Element;
}

export interface SecondaryButton {
    component: () => JSX.Element;
}

export interface NavigationProps<T> {
    routes: T[];
    fallbackPath?: string;
    secondaryButtons?: SecondaryButton[];
}
