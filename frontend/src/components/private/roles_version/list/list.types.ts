import { RolesVersion } from '../../../../services/api';

export interface RolesVersionListProps {
    rolesVersions: RolesVersion[];

    tryEdit: (r: RolesVersion) => void;
    tryDelete: (r: RolesVersion) => void;
}
