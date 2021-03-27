import { createStore, applyMiddleware, combineReducers } from 'redux';
import thunk from 'redux-thunk';

import { registrationReducer } from './registration/reducer';
import { sessionReducer } from './session/reducer';
import { resourceReducer } from './resource/reducer';

const reducer = combineReducers({
    registrationResult: registrationReducer,
    userSession: sessionReducer,
    resourcesResult: resourceReducer,
});

export const store = createStore(reducer, applyMiddleware(thunk));
