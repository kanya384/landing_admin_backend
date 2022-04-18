import { AuthActionTypes } from "../action-types/auth";

export interface AuthenticateRequestSend {
    type: AuthActionTypes.AUTHENTICATE_REQUEST_SEND,
}

export interface AuthenticateSuccess {
    type: AuthActionTypes.AUTHENTICATE_SUCCESS,
    payload: {
        token: string,
    }
}

export interface AuthenticateError {
    type: AuthActionTypes.AUTHENTICATE_ERROR,
    payload: string,
}

export type Authorize = 
 | AuthenticateRequestSend
 | AuthenticateSuccess
 | AuthenticateError