import { Resource } from '../../../../services/api';

export interface ListProps {
    resources: Resource[];

    tryEdit: (r: Resource) => void;
    tryDelete: (r: Resource) => void;
}
