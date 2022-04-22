import { combineReducers } from "redux";
import authsReducer from './authsReducer';
import postersReducer from './postersReducer';

const reducers = combineReducers({
    auths: authsReducer,
    posters: postersReducer,
})

export default reducers;
export type RootState = ReturnType<typeof reducers>;