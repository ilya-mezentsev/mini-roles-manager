import {
    Box,
    Button,
    TextField,
} from "@material-ui/core";

export const SignUp = () => {
    let login: string, password1: string, password2: string;

    return (
        <Box>
            <h1>Sign-Up</h1>

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
                type="password"
                margin="normal"
                onInput={e => password1 = (e.target as HTMLInputElement).value}
            />
            <TextField
                id="standard-basic"
                label="Password Again"
                required
                fullWidth={true}
                type="password"
                margin="normal"
                onInput={e => password2 = (e.target as HTMLInputElement).value}
            />

            <Button
                variant="contained"
                color="primary"
                onClick={() => trySignUp(login, password1, password2)}
            >Sign-Up</Button>
        </Box>
    )
}

function trySignUp(login: string, password1: string, password2: string): void {
    if (![login, password1, password1].every(s => !!s)) {
        alert('Login of one of passwords are not provided');
    } else if (password1 !== password2) {
        alert('Passwords are not match');
    } else {
        alert(`Got login - ${login}, password - ${password1}`);
    }
}
