import { Container } from '@material-ui/core';
import { BrowserRouter as Router } from 'react-router-dom';

import Entrypoint from './containers/entrypoint';

const App = () => (
    <Router>
        <Container>
            <Entrypoint/>
        </Container>
    </Router>
);

export default App;
