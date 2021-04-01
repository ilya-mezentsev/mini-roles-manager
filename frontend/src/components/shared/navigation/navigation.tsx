import {
    AppBar,
    Button,
    Container,
    Toolbar,
} from '@material-ui/core';
import {
    Link,
    Redirect,
    Route,
    Switch,
} from 'react-router-dom';
import { NavigationProps } from './navigation.types';

export const Navigation = (props: NavigationProps) => {
    let size: 'sm' | 'xl';
    if (props.size === 'large') {
        size = 'xl';
    } else {
        size = 'sm';
    }

    return (
        <>
            <AppBar position="static">
                <Toolbar>
                    {
                        props.routes.map((route, index) =>
                            <Button
                                key={`route_${index}`}
                                color="inherit"
                                component={Link}
                                to={route.path}
                            >
                                {route.name}
                            </Button>
                        )
                    }
                </Toolbar>
            </AppBar>

            <Container maxWidth={size}>
                <Switch>
                    {
                        props.routes.map((route, index) =>
                            <Route
                                key={`route_${index}`}
                                path={route.path}
                            >
                                {<route.component/>}
                            </Route>
                        )
                    }
                    <Route path="*">
                        <Redirect to={props.fallbackPath} />
                    </Route>
                </Switch>
            </Container>
        </>
    );
}
