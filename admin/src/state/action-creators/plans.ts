import { Dispatch } from "react"
import { Configuration, PlansApi } from "../../api"
import { GetTokenFromCookies } from "../../utils"
import { PlansActionTypes } from "../action-types"
import { PlansAction } from "../actions"

const PLANS_ERROR = "Ошибка при получении данных"
const url = "http://localhost:3000"
const plansService = new PlansApi(new Configuration(), url)

export const getPlans = () => {
  return async (dispatch: Dispatch<PlansAction>) => {
    dispatch({
      type: PlansActionTypes.PLANS_REQUEST_SEND
    })
    let token = GetTokenFromCookies()
    plansService.plansGet({headers: {"Authorization": token}}).then((resp) => {
      if (resp.status === 200) {
        dispatch({
          type: PlansActionTypes.PLANS_SUCCESS,
          payload: resp.data
        })
      }
    }).catch((e)=>{
      dispatch({
        type: PlansActionTypes.PLANS_ERROR,
        payload: PLANS_ERROR,
      })
    })
  }
}

export const updatePlanPhoto = (id:string, file: any, callback:(error: string) => void) => {
  return async (dispatch: Dispatch<PlansAction>) => {
    let token = GetTokenFromCookies()
    plansService.plansPatch(file, id, {headers: {"Authorization": token}}).then((resp) => {
      if (resp.status === 200) {
        dispatch({
          type: PlansActionTypes.PLANS_UPDATE_PHOTO,
          payload: resp.data
        })
        callback("")
      } else {
        callback("ошибка при обновлении данных")
      }
    }).catch((e) => {
      callback("ошибка при обновлении данных")
    })
  }
}