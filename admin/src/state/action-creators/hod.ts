import { Dispatch } from "react"
import { Configuration, HodApi, HodPhoto, Month, Year } from "../../api"
import { GetTokenFromCookies } from "../../utils";
import { HodActionTypes } from "../action-types";
import { Hod } from "../actions";

const HOD_ERROR = "Ошибка при получении данных"
const hodService = new HodApi(new Configuration(), process.env.REACT_APP_BACKEND_URL)
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

export const addYear = (year: Year, callback?:(error: string) => void) => {
  return async (dispatch: Dispatch<Hod>) => {
    let token = GetTokenFromCookies()
    hodService.yearsPut(year, {headers: {"Authorization": token}}).then((resp)=>{
      if (resp.status === 200) {
        dispatch({
            type: HodActionTypes.HOD_YEARS_NEW,
            payload: resp.data!,
        });
        callback!("")
      }
    }).catch((e)=>{
      callback!("ошибка при добавлении")
    })
  }
}

export const deleteYear = (yearID: string) => {
  return async (dispatch: Dispatch<Hod>) => {
    let token = GetTokenFromCookies()
    hodService.yearsIdDelete(yearID, {headers: {"Authorization": token}}).then((resp)=>{
      if (resp.status === 200) {
        dispatch({
            type: HodActionTypes.HOD_YEARS_DELETE,
            payload: resp.data!.msg!,
        });
      }
    })
  }
}

export const getMonths = (yearID: string, callback?:(error: string) => void) =>{
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

export const addMonth = (month: Month, callback?:(error: string) => void) => {
  return async (dispatch: Dispatch<Hod>) => {
    let token = GetTokenFromCookies()
    hodService.monthsPut(month, {headers: {"Authorization": token}}).then((resp)=>{
      if (resp.status === 200) {
        dispatch({
            type: HodActionTypes.HOD_MONTHS_NEW,
            payload: resp.data!,
        });
        callback!("")
      }
    }).catch((e)=>{
      callback!("ошибка при добавлении")
    })
  }
}

export const deleteMonth = (monthID: string) => {
  return async (dispatch: Dispatch<Hod>) => {
    let token = GetTokenFromCookies()
    hodService.monthsIdDelete(monthID, {headers: {"Authorization": token}}).then((resp)=>{
      if (resp.status === 200) {
        dispatch({
            type: HodActionTypes.HOD_MONTHS_DELETE,
            payload: monthID,
        });
      }
    })
  }
}

export const getPhotos = (monthID: string, callback?:(error: string) => void) =>{
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

export const addPhotos = (filesList: any[], monthId: string, callback:(error: string) => void) => {
  return async (dispatch: Dispatch<Hod>) => {
    console.log(filesList)
    if (filesList.length == 0) {
      callback("не приложили ни один файл")
      return
    }
    if (!monthId) {
      callback("не указан месяц")
      return
    }
    let token = GetTokenFromCookies()
    filesList.forEach((file) => {
      console.log(file)
      hodService.photosPut(file, monthId, {headers: {"Authorization": token}}).then((resp)=>{
        if (resp.status === 200) {
          dispatch({
              type: HodActionTypes.HOD_PHOTOS_NEW,
              payload: resp.data!,
          });
        }
      }).catch((e)=>{})
    })
  }
}

export const deletePhotoHod = (photoID: string, callback?:(error: string) => void) => {
  return async (dispatch: Dispatch<Hod>) => {
    let token = GetTokenFromCookies()
    hodService.photosIdDelete(photoID, {headers: {"Authorization": token}}).then((resp)=>{
      if (resp.status === 200) {
        dispatch({
            type: HodActionTypes.HOD_PHOTOS_DELETE,
            payload: photoID,
        });
      }
    })
  }
}

export const sortPhotos = (photos: HodPhoto[], dragIndex: number, hoverIndex: number) => {
  return async (dispatch: Dispatch<Hod>) => {
    let first  = {
      id: photos[dragIndex].id,
      order: hoverIndex,
    }
    let second = {
      id: photos[hoverIndex].id,
      order: dragIndex,
    }
    let token = GetTokenFromCookies()
    hodService.photosOrdersPost({first: first, second: second}, {headers: {"Authorization": token}}).then((resp)=>{
      if (resp.status === 200) {
        dispatch({
          type: HodActionTypes.HOD_PHOTOS_SORT,
          payload: {
            dragIndex: dragIndex,
            hoverIndex: hoverIndex
          },
        });
      }
    })
  }
}