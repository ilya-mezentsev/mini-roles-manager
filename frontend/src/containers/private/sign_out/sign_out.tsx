import { observer } from 'mobx-react-lite';
import ExitToAppIcon from '@material-ui/icons/ExitToApp';

import { Alert } from '../../../components/shared';
import { sessionStore } from '../../../store';

export const SignOut = observer(() => {
    return (
        <>
            <ExitToAppIcon
                color="inherit"
                cursor="pointer"
                onClick={() => sessionStore.signOut()}
            />

            <Alert
                shouldShow={!!sessionStore.signOutError}
                severity={'error'}
                message={sessionStore.signOutError?.description || 'Unknown error'}
                onCloseCb={() => sessionStore.cleanSessionActionErrors()}
            />
        </>
    )
});
