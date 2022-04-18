import produce from 'immer';
import { AuthActionTypes } from '../action-types';
import { Authorize } from '../actions';

interface AuthState {
    loading: boolean;
    auth: boolean;
    error: string | null;
    token: string | null;
}

const initialState: AuthState = {
    loading: false,
    auth: false,
    error: null,
    token: null,
}

const reducer = produce((state: AuthState = initialState, action: Authorize) =>{
    switch (action.type) {
        case AuthActionTypes.AUTHENTICATE_REQUEST_SEND:
            state.loading = true
            state.auth = false
            return state;
        case AuthActionTypes.AUTHENTICATE_SUCCESS:
            const {token} = action.payload
            state.token = token
            state.auth = true
            return state
        case AuthActionTypes.AUTHENTICATE_ERROR:
            state.error = action.payload
            state.auth = false
            return state
        default:
            return state;
    }
})

export default reducer;