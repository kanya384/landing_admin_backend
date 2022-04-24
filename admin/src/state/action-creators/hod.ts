import { Dispatch } from "react"
import { Configuration, HodApi } from "../../api"
import { GetTokenFromCookies } from "../../utils";
import { HodActionTypes } from "../action-types";
import { Hod } from "../actions";

const HOD_ERROR = "Ошибка при получении данных"
const url = "http://localhost:3000"
const hodService = new HodApi(new Configuration(), url)
export const getYears = () =>{
  return async (dispatch: Dispatch<Hod>) => {
    dispatch({
      type: HodActionTypes.HOD_YEARS_REQUEST_SEND
    })
    let token = GetTokenFromCookies()
    hodService.yearsGet({headers: {"Authorization": token}}).then((resp)=>{
      if (resp.status === 200) {
        dispatch({
            type: HodActionTypes.HOD_YEARS_SUCCESS,
            payload: resp.data!,
        });
      } else {
          dispatch({
              type: HodActionTypes.HOD_ERROR,
              payload: HOD_ERROR
          });
      }
    }).catch((e)=>{
      dispatch({
        type: HodActionTypes.HOD_ERROR,
        payload: HOD_ERROR
      });
    })
  }
}

export const getMonths = (yearID: string) =>{
  return async (dispatch: Dispatch<Hod>) => {
    dispatch({
      type: HodActionTypes.HOD_MONTHS_REQUEST_SEND
    })
    let token = GetTokenFromCookies()
    hodService.monthsIdGet(yearID, {headers: {"Authorization": token}}).then((resp)=>{
      if (resp.status === 200) {
        dispatch({
            type: HodActionTypes.HOD_MONTHS_SUCCESS,
            payload: resp.data!,
        });
      } else {
          dispatch({
              type: HodActionTypes.HOD_ERROR,
              payload: HOD_ERROR
          });
      }
    }).catch((e)=>{
      dispatch({
        type: HodActionTypes.HOD_ERROR,
        payload: HOD_ERROR
      });
    })
  }
}

export const getPhotos = (monthID: string) =>{
  return async (dispatch: Dispatch<Hod>) => {
    dispatch({
      type: HodActionTypes.HOD_PHOTOS_REQUEST_SEND
    })
    let token = GetTokenFromCookies()
    hodService.photosIdGet(monthID, {headers: {"Authorization": token}}).then((resp)=>{
      if (resp.status === 200) {
        dispatch({
            type: HodActionTypes.HOD_PHOTOS_SUCCESS,
            payload: resp.data!,
        });
      } else {
          dispatch({
              type: HodActionTypes.HOD_ERROR,
              payload: HOD_ERROR
          });
      }
    }).catch((e)=>{
      dispatch({
        type: HodActionTypes.HOD_ERROR,
        payload: HOD_ERROR
      });
    })
  }
}