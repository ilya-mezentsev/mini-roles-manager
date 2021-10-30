import { observer } from 'mobx-react-lite';

import { PublicNavigation as PublicEntryPoint } from './public';
import { PrivateEntrypoint } from './private';
import { sessionStore } from '../store';

const Entrypoint = observer(() => {
    return (
        <>
            {
                !!sessionStore.session?.id
                    ? <PrivateEntrypoint />
                    : <PublicEntryPoint />
            }
        </>
    );
});

export default Entrypoint;
