import produce from 'immer';
import { AuthActionTypes } from '../action-types';
import { Authorize } from '../actions';

interface AuthState {
    loading: boolean;
    auth: boolean | null;
    error: string | null;
    token: string | null;
}

const initialState: AuthState = {
    loading: true,
    auth: null,
    error: null,
    token: null,
}

const reducer = produce((state: AuthState = initialState, action: Authorize) =>{
    switch (action.type) {
        case AuthActionTypes.AUTHENTICATE_REQUEST_SEND:
            state.loading = true
            return state;
        case AuthActionTypes.AUTHENTICATE_SUCCESS:
            const {token} = action.payload
            state.token = token
            state.auth = true
            state.loading = false
            return state
        case AuthActionTypes.AUTHENTICATE_ERROR:
            state.error = action.payload
            state.auth = false
            state.loading = false
            return state
        default:
            return state;
    }
})

export default reducer;