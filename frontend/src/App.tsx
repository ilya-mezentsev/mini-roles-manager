import { Container } from '@material-ui/core';
import { BrowserRouter as Router } from 'react-router-dom';
import { Provider } from 'react-redux';

import { store } from './store';
import Entrypoint from './containers/entrypoint';

const App = () => (
    <Provider store={store}>
        <Router>
            <Container>
                <Entrypoint/>
            </Container>
        </Router>
    </Provider>
);

export default App;
