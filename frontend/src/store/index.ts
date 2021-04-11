import { createStore, applyMiddleware, combineReducers } from 'redux';
import thunk from 'redux-thunk';

import { registrationReducer } from './registration/reducer';
import { sessionReducer } from './session/reducer';
import { resourceReducer } from './resource/reducer';
import { roleReducer } from './role/reducer';
import { userSessionMiddleware, getPreloadedUserSessionId } from './middleware/user_session';

const reducer = combineReducers({
    registrationResult: registrationReducer,
    userSession: sessionReducer,
    resourcesResult: resourceReducer,
    rolesResult: roleReducer,
});

const preloadedState = {
    userSession: getPreloadedUserSessionId(),
};

export const store = createStore(
    reducer,
    preloadedState,
    applyMiddleware(thunk, userSessionMiddleware),
);
