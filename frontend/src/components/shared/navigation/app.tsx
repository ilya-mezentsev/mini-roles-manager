import {
    AppBar,
    Box,
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
import { NavigationProps, Route as NavigationRoute } from './navigation.types';

export const Navigation = (props: NavigationProps<NavigationRoute>) => {
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
                    <Box style={{ flex: 1 }}>
                        {
                            props.routes.map((route, index) => (
                                <Button
                                    key={`route_${index}`}
                                    color="inherit"
                                    component={Link}
                                    to={route.path}
                                >
                                    {route.name}
                                </Button>
                            ))
                        }
                    </Box>
                    <Box>
                        {
                            (props.secondaryButtons || []).map((b, index) => (
                                <b.component
                                    key={`secondary_button_${index}`}
                                />
                            ))
                        }
                    </Box>
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
                    {
                        props.fallbackPath &&
                        <Route path="*">
                            <Redirect to={props.fallbackPath} />
                        </Route>
                    }
                </Switch>
            </Container>
        </>
    );
}
