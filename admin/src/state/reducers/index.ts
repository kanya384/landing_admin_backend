import { combineReducers } from "redux";
import authsReducer from './authsReducer';
import postersReducer from './postersReducer';
import hodReducer from './hodReducer';

const reducers = combineReducers({
    auths: authsReducer,
    posters: postersReducer,
    hod: hodReducer,
})

export default reducers;
export type RootState = ReturnType<typeof reducers>;