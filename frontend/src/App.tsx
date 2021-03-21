import React from 'react';
import { Container } from '@material-ui/core';

import './App.css';
import { Navigation as PublicNavigation } from "./containers/public/navigation/navigation";

function App() {
    return (
        <Container>
            <PublicNavigation />
        </Container>
    );
}

export default App;
