import {
    BrowserRouter as Router,
    Switch,
    Route,
    Redirect,
    Link
} from 'react-router-dom';
import {
    AppBar,
    Button,
    Toolbar,
    Container,
} from '@material-ui/core';

import { SignIn } from '../sign_in/sign_in';
import { SignUp } from '../sign_up/sign_up';
import { About } from '../../shared/about/about';

export const Navigation = () => {
    return (
        <Router>
            <div>
                <AppBar position="static">
                    <Toolbar>
                        <Button color="inherit" component={Link} to="/sign-in">
                            Sign-In
                        </Button>

                        <Button color="inherit" component={Link} to="/sign-up">
                            Sign-Up
                        </Button>

                        <Button color="inherit" component={Link} to="/about">
                            About
                        </Button>
                    </Toolbar>
                </AppBar>
            </div>

            <Container maxWidth="sm">
                <Switch>
                    <Route path="/sign-in">
                        <SignIn />
                    </Route>
                    <Route path="/sign-up">
                        <SignUp />
                    </Route>
                    <Route path="/about">
                        <About />
                    </Route>
                    <Route path="*">
                        <Redirect to="/sign-in" />
                    </Route>
                </Switch>
            </Container>
        </Router>
    );
}
