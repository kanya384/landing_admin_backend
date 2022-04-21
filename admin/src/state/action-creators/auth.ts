import { Dispatch } from 'redux';
import { AuthActionTypes } from '../action-types/auth';
import { Authorize } from '../actions/auth';
import {Configuration, DefaultApi} from "../../api";
import { AxiosResponse } from 'axios';
import { GetTokenFromCookies } from '../../utils';


const AUTH_ERROR = "Проверьте логин или пароль"
const ERR_FIELD_EMPTY = "Заполните необходимые поля"

const apiSerivice = new DefaultApi(new Configuration(), "http://localhost:3000/");

export const sendAuth = (login: string, pass: string) => {
    return async (dispatch: Dispatch<Authorize>) => {
        if (login === "" || pass === "") {
            dispatch({
                type: AuthActionTypes.AUTHENTICATE_ERROR,
                payload: ERR_FIELD_EMPTY
            });
            return
        }
        dispatch({
            type: AuthActionTypes.AUTHENTICATE_REQUEST_SEND,
        });

        apiSerivice.loginPost({login, pass}).then((resp)=>{
            if (resp.status === 200) {
                dispatch({
                    type: AuthActionTypes.AUTHENTICATE_SUCCESS,
                    payload: {
                        token: resp.data.token!,
                    }
                });
            } else {
                dispatch({
                    type: AuthActionTypes.AUTHENTICATE_ERROR,
                    payload: AUTH_ERROR
                });
            }
        }).catch((e)=>{
            dispatch({
                type: AuthActionTypes.AUTHENTICATE_ERROR,
                payload: AUTH_ERROR
            });
        })
        
    }
}

export const checkAuth = () => {
    return async (dispatch: Dispatch<Authorize>) => {
        dispatch({
            type: AuthActionTypes.AUTHENTICATE_REQUEST_SEND,
        });

        let token = GetTokenFromCookies()
        apiSerivice.pingGet({headers: {"Authorization": token}}).then((resp:AxiosResponse)=>{
            if (resp.status!==200){
                dispatch({
                    type: AuthActionTypes.AUTHENTICATE_ERROR,
                    payload: AUTH_ERROR,
                });
                return
            }
            dispatch({
                type: AuthActionTypes.AUTHENTICATE_SUCCESS,
                payload: {
                    token: token,
                }
            });
        }).catch((e) => {
            dispatch({
                type: AuthActionTypes.AUTHENTICATE_ERROR,
                payload: "",
            }); 
        })
    }
}