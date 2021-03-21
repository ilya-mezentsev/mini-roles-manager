import {
    TextField,
    Box,
    Button,
} from '@material-ui/core';

export const SignIn = () => {
    let login: string, password: string;

    return (
        <Box>
            <h1>Sign-In</h1>

            <TextField
                id="standard-basic"
                label="Login"
                required
                fullWidth={true}
                margin="normal"
                onInput={e => login = (e.target as HTMLInputElement).value}
            />
            <TextField
                id="standard-basic"
                label="Password"
                required
                fullWidth={true}
                margin="normal"
                type="password"
                onInput={e => password = (e.target as HTMLInputElement).value}
            />

            <Button
                variant="contained"
                color="primary"
                onClick={() => trySignIn(login, password)}
            >Sign-In</Button>
        </Box>
    )
}

function trySignIn(login: string, password: string): void {
    if (!login || !password) {
        alert('No login of password provided');
    } else {
        alert(`Got login - ${login}, password - ${password}`);
    }
}
