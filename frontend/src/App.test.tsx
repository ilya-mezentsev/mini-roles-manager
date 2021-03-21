import React from 'react';
import { render, screen } from '@testing-library/react';
import App from './App';

test('renders learn react link', async () => {
    render(<App />);
    const linkElements = await screen.findAllByText(/Sign-In/i);
    for (const linkElement of linkElements) {
        expect(linkElement).toBeInTheDocument();
    }
});
