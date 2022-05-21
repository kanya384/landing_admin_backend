import { combineReducers } from "redux";
import authsReducer from './authsReducer';
import postersReducer from './postersReducer';
import hodReducer from './hodReducer';
import advantagesReducer from './advantagesReducer';
import plansReducer from './plansReducer';
import docsReducer from './docsReducer';
import videoReducer from './videoReducer';

const reducers = combineReducers({
    auths: authsReducer,
    posters: postersReducer,
    hod: hodReducer,
    advantages: advantagesReducer,
    plans: plansReducer,
    docs: docsReducer,
    videos: videoReducer,
})

export default reducers;
export type RootState = ReturnType<typeof reducers>;