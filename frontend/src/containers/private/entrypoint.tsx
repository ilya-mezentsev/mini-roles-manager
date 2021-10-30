import { useEffect } from 'react';
import { observer } from 'mobx-react-lite';

import { Navigation as PrivateNavigation } from './navigation/navigation';
import {
    accountInfoStore,
    resourceStore,
    roleStore,
    rolesVersionStore,
} from '../../store';
import * as log from '../../services/log';

export const Entrypoint = observer(() => {
    useEffect(() => {
        const fetchTargets = [];

        if (!resourceStore.list.length) {
            fetchTargets.push(resourceStore.fetchResources());
        }

        if (!roleStore.list.length) {
            fetchTargets.push(roleStore.fetchRoles());
        }

        if (!rolesVersionStore.list.length) {
            fetchTargets.push(rolesVersionStore.fetchRolesVersions());
        }

        if (!accountInfoStore.info?.login) {
            fetchTargets.push(accountInfoStore.fetchInfo());
        }

        Promise.all(fetchTargets)
            .finally(() => log.info('All data is loaded'));
        // eslint-disable-next-line
    }, []);

    return (
        <PrivateNavigation/>
    );
});
