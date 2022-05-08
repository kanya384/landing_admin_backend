import { AdvantagesApi, Configuration } from "../../api"
import { AdvantageAction } from "../actions"
import { Dispatch } from "react";
import { AdvantagesActionTypes } from "../action-types";

const ADVANTAGES_ERROR = "Ошибка при получении данных"
const url = "http://localhost:3000"
const advantagesService = new AdvantagesApi(new Configuration(), url)
export const getADvantages = () => {
    return async (dispatch: Dispatch<AdvantageAction>) => {
        dispatch({
            type: AdvantagesActionTypes.ADVANTAGES_REQUEST_SEND,
        })
        advantagesService.advantagesGet()
    }
}