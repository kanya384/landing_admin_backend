import { combineReducers } from "redux";
import authsReducer from './authsReducer';

const reducers = combineReducers({
    auths: authsReducer,
})

export default reducers;
export type RootState = ReturnType<typeof reducers>;