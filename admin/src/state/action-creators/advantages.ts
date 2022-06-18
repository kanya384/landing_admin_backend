import { Advantage, AdvantagesApi, Configuration } from "../../api"
import { AdvantageAction } from "../actions"
import { Dispatch } from "react";
import { AdvantagesActionTypes } from "../action-types";
import { GetTokenFromCookies } from "../../utils";

const ADVANTAGES_ERROR = "Ошибка при получении данных"
const advantagesService = new AdvantagesApi(new Configuration(), process.env.REACT_APP_BACKEND_URL)
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

export const getAdvantageById = (id: string) => {
    return async (dispatch: Dispatch<AdvantageAction>) => {
        dispatch({
            type: AdvantagesActionTypes.ADVANTAGE_ID_REQUEST_SEND,
        })
        let token = GetTokenFromCookies()
        advantagesService.advantagesIdGet(id, {headers: {"Authorization": token}}).then((resp) => {
            if (resp.status === 200) {
                dispatch({
                    type: AdvantagesActionTypes.ADVANTAGE_ID_SUCCESS,
                    payload: resp.data,
                })
            } else {
                dispatch({
                    type: AdvantagesActionTypes.ADVANTAGE_ID_ERROR,
                    payload: ADVANTAGES_ERROR
                })
            }
        }).catch((e) => {
            dispatch({
                type: AdvantagesActionTypes.ADVANTAGE_ID_ERROR,
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

export const deleteAdvantage = (id: string) => {
    return async (dispatch: Dispatch<AdvantageAction>) => {
        let token = GetTokenFromCookies()
        advantagesService.advantagesIdDelete(id, {headers: {"Authorization": token}}).then((resp)=> {
            if (resp.status === 200) {
                dispatch({
                    type: AdvantagesActionTypes.ADVANTAGES_DELETE,
                    payload: id
                })
            }
        })
    }
}

export const updateAdvantage = (advantage: Advantage, callback:(error: string) => void, file?: any) => {
    return async (dispatch: Dispatch<AdvantageAction>) => {
      let token = GetTokenFromCookies()
      if (advantage.id?.length === 0 || advantage.title?.length === 0 || advantage.description?.length === 0 || advantage.order === undefined) {
        callback("не все поля заполненны")
        return
      }
      advantagesService.advantagesPatch(advantage, {headers: {"Authorization": token}}).then((resp)=>{
        if (resp.status === 200) {
          dispatch({
            type: AdvantagesActionTypes.ADVANTAGES_UPDATE,
            payload: advantage,
          });
          callback("")
        } else {
          callback("ошибка при редактировании")
        }
      })
    }
}

export const addAdvantagePhoto = (filesList: any[], advantageID: string, callback:(error: string) => void) => {
    return async (dispatch: Dispatch<AdvantageAction>) => {
      console.log(filesList)
      if (filesList.length == 0) {
        callback("не приложили ни один файл")
        return
      }
      if (!advantageID) {
        callback("не указан месяц")
        return
      }
      let token = GetTokenFromCookies()
      filesList.forEach((file) => {
        console.log(file)
        advantagesService.advantagePhotoPut(file, advantageID, 0, {headers: {"Authorization": token}}).then((resp)=>{
          if (resp.status === 200) {
            dispatch({
                type: AdvantagesActionTypes.ADVANTAGE_PHOTO_NEW,
                payload: resp.data!,
            });
          }
        }).catch((e)=>{})
      })
    }
}

export const getAdvantagePhotosByID = (id: string) => {
    return async (dispatch: Dispatch<AdvantageAction>) => {
        dispatch({
            type: AdvantagesActionTypes.ADVANTAGE_PHOTO_REQUEST,
        })
        let token = GetTokenFromCookies()
        advantagesService.advantagePhotoIdGet(id, {headers: {"Authorization": token}}).then((resp)=>{
            if (resp.status === 200) {
                dispatch({
                    type: AdvantagesActionTypes.ADVANTAGE_PHOTO_SUCCESS,
                    payload: resp.data
                })
            }
        })
    }
}

export const deleteAdvantagePhotoByID = (id: string) => {
    return async (dispatch: Dispatch<AdvantageAction>) => {

        let token = GetTokenFromCookies()
        advantagesService.advantagePhotoIdDelete(id, {headers: {"Authorization": token}}).then((resp)=>{
            if (resp.status === 200) {
                dispatch({
                    type: AdvantagesActionTypes.ADVANTAGE_PHOTO_DELETE,
                    payload: id,
                })
            }
        })
    }
}