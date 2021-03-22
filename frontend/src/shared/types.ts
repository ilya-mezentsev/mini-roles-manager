import { Dispatch } from 'redux';

export type StateToPropsFn<T> = () => (state: T) => T
export type DispatchToPropsFn<T> = () => (dispatch: Dispatch) => T;
