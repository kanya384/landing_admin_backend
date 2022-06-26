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

export const addNewTitle = (title: Title, callback: (error: string) => void) => {
  return async (dispatch: Dispatch<TitlesAction>) => {
    if (!title.tagName || ! title.tagValue || !title.desktopTitle || !title.mobileTitle) {
      callback("не все поля заполненны")
      return
    }
    let token = GetTokenFromCookies()
    titlesService.titlesPut(title, {headers: {"Authorization": token}}).then((resp) => {
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

export const updateTitle = (title: Title, callback: (error: string) => void) => {
  return async (dispatch: Dispatch<TitlesAction>) => {
    if (!title.tagName || ! title.tagValue || !title.desktopTitle || !title.mobileTitle) {
      callback("не все поля заполненны")
      return
    }

    let token = GetTokenFromCookies()
    titlesService.titlesPatch(title, {headers: {"Authorization": token}}).then((resp) => {
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