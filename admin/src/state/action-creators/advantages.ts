import { AdvantagesApi, Configuration } from "../../api"
import { AdvantageAction } from "../actions"
import { Dispatch } from "react";
import { AdvantagesActionTypes } from "../action-types";
import { GetTokenFromCookies } from "../../utils";

const ADVANTAGES_ERROR = "Ошибка при получении данных"
const url = "http://localhost:3000"
const advantagesService = new AdvantagesApi(new Configuration(), url)
export const getAdvantages = () => {
    return async (dispatch: Dispatch<AdvantageAction>) => {
        dispatch({
            type: AdvantagesActionTypes.ADVANTAGES_REQUEST_SEND,
        })
        let token = GetTokenFromCookies()
        advantagesService.advantagesGet({headers: {"Authorization": token}}).then((resp)=>{
            if (resp.status === 200) {
                dispatch({
                    type: AdvantagesActionTypes.ADVANTAGES_SUCCESS,
                    payload: resp.data!,
                });
              }
        }).catch((e)=>{
            dispatch({
                type: AdvantagesActionTypes.ADVANTAGES_ERROR,
                payload: ADVANTAGES_ERROR
            })
        })
    }
}

export const addAdvantage = (title: string, description: string, callback:(error: string) => void) => {
    return async (dispatch: Dispatch<AdvantageAction>) => {
        if (!title || !description) {
            callback("не все поля заполненны")
            return
        }
        let token = GetTokenFromCookies()
        advantagesService.advantagesPut({title, description}, {headers: {"Authorization": token}}).then((resp)=>{
            if (resp.status === 200) {
                dispatch({
                    type: AdvantagesActionTypes.ADVANTAGES_NEW,
                    payload: resp.data
                })
                callback("")
            } else {
                callback("ошибка при добавлении")
            }
        })
    }
}