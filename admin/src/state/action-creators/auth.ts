import { Dispatch } from 'redux';
import { AuthActionTypes } from '../action-types/auth';
import { Authorize } from '../actions/auth';
import {DefaultApi} from "../../api";

const apiSerivice = new DefaultApi();

export const sendAuth = (login: string, pass: string) => {
    return async (dispatch: Dispatch<Authorize>) => {
        dispatch({
            type: AuthActionTypes.AUTHENTICATE_REQUEST_SEND,
        });

        const result = await apiSerivice.loginPost({login, pass})
        if (result.status == 200) {
            dispatch({
                type: AuthActionTypes.AUTHENTICATE_SUCCESS,
                payload: {
                    token: result.data.token!,
                }
            });
        } else {
            dispatch({
                type: AuthActionTypes.AUTHENTICATE_ERROR,
                payload: result.data.error!
            });
        }
    }
}