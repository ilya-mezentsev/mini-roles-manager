import { Effect } from '../../services/api/shared/types';
import { APIError } from '../../services/api/shared';

export interface FetchPermissionActionResult<T = Effect, E = APIError | Error> {
    effect?: T;
    error?: E;
}

export type FetchPermissionResult = FetchPermissionActionResult<Effect | null, APIError | null>
