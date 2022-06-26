import { Configuration, Title, TitlesApi } from "../../api"
import { Dispatch } from "react";
import { TitlesAction } from "../actions";
import { TitlesActionTypes } from "../action-types";
import { GetTokenFromCookies } from "../../utils";

const TITLES_ERROR = "Ошибка при получении данных"
const titlesService = new TitlesApi(new Configuration(), process.env.REACT_APP_BACKEND_URL)

export const getTitles = () => {
  return async (dispatch: Dispatch<TitlesAction>) => {
    dispatch({
      type: TitlesActionTypes.TITLES_REQUEST_SEND
    })

    let token = GetTokenFromCookies()
    titlesService.titlesGet({headers: {"Authorization": token}}).then((resp)=>{
      if (resp.status === 200) {
        dispatch({
          type: TitlesActionTypes.TITLES_SUCCESS,
          payload: resp.data,
        })
      } else {
        dispatch({
          type: TitlesActionTypes.TITLES_ERROR,
          payload: TITLES_ERROR,
        })
      }
    }).catch((e) => {
      dispatch({
        type: TitlesActionTypes.TITLES_ERROR,
        payload: TITLES_ERROR,
      })
    })
  }
}

export const addNewTitle = (id: string | undefined, tagName: string, tagValue: string, desktopTitle: string, mobileTitle: string, callback: (error: string) => void) => {
  return async (dispatch: Dispatch<TitlesAction>) => {
    if (!tagName || ! tagValue || !desktopTitle || !mobileTitle) {
      callback("не все поля заполненны")
      return
    }
    let token = GetTokenFromCookies()
    titlesService.titlesPut(id, tagName, tagValue, desktopTitle, mobileTitle, {headers: {"Authorization": token}}).then((resp) => {
      if (resp.status === 200) {
        dispatch({
          type: TitlesActionTypes.TITLES_NEW,
          payload: resp.data,
        });
        callback("")
      } else {
        callback("ошибка при добавлении")
      }
    })
  }
}

export const updateTitle = (id: string | undefined, tagName: string, tagValue: string, desktopTitle: string, mobileTitle: string, callback: (error: string) => void) => {
  return async (dispatch: Dispatch<TitlesAction>) => {
    if (!tagName || ! tagValue || !desktopTitle || !mobileTitle) {
      callback("не все поля заполненны")
      return
    }

    let token = GetTokenFromCookies()
    titlesService.titlesPatch(id, tagName, tagValue, desktopTitle, mobileTitle, {headers: {"Authorization": token}}).then((resp) => {
      if (resp.status === 200) {
        dispatch({
          type: TitlesActionTypes.TITLES_UPDATE,
          payload: resp.data,
        });
        callback("")
      } else {
        callback("ошибка при добавлении")
      }
    })
  }
}

export const deleteTitles = (id: string) => {
  return async (dispatch: Dispatch<TitlesAction>) => {
    let token = GetTokenFromCookies()
    titlesService.titlesIdDelete(id,  {headers: {"Authorization": token}}).then((resp)=>{
      if (resp.status === 200) {
        dispatch({
          type: TitlesActionTypes.TITLES_DELETE,
          payload: id,
        });
      }
    })
  }
}