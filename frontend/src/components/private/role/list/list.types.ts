import { Role } from '../../../../services/api';

export interface ListProps {
    roles: Role[];

    tryEdit: (r: Role) => void;
    tryDelete: (r: Role) => void;
}
