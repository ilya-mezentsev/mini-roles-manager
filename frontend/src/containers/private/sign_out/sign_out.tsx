import { bindActionCreators } from 'redux';
import ExitToAppIcon from '@material-ui/icons/ExitToApp';

import { Alert } from '../../../components/shared';
import { DispatchToPropsFn, StateToPropsFn } from '../../../shared/types';
import { signOut, cleanSignOutError } from '../../../store/session/actions';
import { SignOutActions, SignOutProps, SignOutState } from './sign_out.types';

export const SignOut = (props: SignOutProps) => {
    return (
        <>
            <ExitToAppIcon
                color="inherit"
                cursor="pointer"
                onClick={() => props.signOutAction()}
            />

            <Alert
                shouldShow={!!props.userSession?.error}
                severity={'error'}
                message={props.userSession?.error?.description || 'Unknown error'}
                onCloseCb={() => props.cleanSignOutErrorAction()}
            />
        </>
    )
};

export const mapStateToProps: StateToPropsFn<SignOutState> = () => state => ({
    userSession: state.userSession,
});

export const mapDispatchToProps: DispatchToPropsFn<SignOutActions> = () => dispatch => ({
    signOutAction: bindActionCreators(signOut, dispatch),
    cleanSignOutErrorAction: bindActionCreators(cleanSignOutError, dispatch)
});