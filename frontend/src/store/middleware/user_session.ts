import { Middleware } from 'redux';

const userSessionIdKey = 'userSessionId';

export const userSessionMiddleware: Middleware<unknown, {}, any> = storeAPI => next => action => {
    const result = next(action);
    const userSessionId = (storeAPI.getState() as any)?.userSession?.session?.id;
    if (userSessionId || userSessionId === '') {
        localStorage.setItem(userSessionIdKey, userSessionId);
    }

    return result;
};

export const getPreloadedUserSessionId = () => {
    return {
        session: {
            id: localStorage.getItem(userSessionIdKey) || '',
        }
    };
};
