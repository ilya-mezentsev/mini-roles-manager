import { createStore, applyMiddleware, combineReducers } from 'redux';
import thunk from 'redux-thunk';

import { registrationReducer } from './registration/reducer';
import { sessionReducer } from './session/reducer';

const reducer = combineReducers({
    registrationResult: registrationReducer,
    userSession: sessionReducer,
});

export const store = createStore(reducer, applyMiddleware(thunk));
