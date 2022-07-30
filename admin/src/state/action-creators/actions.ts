import { Dispatch } from "react";
import { Configuration, Action, ActionsApi } from "../../api";
import { GetTokenFromCookies } from "../../utils";
import { ActionsActionTypes } from "../action-types";
import { Actions } from "../actions";

const ACTIONS_ERROR = "Ошибка при получении данных"
const actionsService = new ActionsApi(new Configuration(), process.env.REACT_APP_BACKEND_URL)
export const getActions = () => {
  return async (dispatch: Dispatch<Actions>) => {
    dispatch({
      type: ActionsActionTypes.ACTIONS_REQUEST_SEND,
    });
    actionsService.actionsGet().then((resp)=>{
      if (resp.status === 200) {
        dispatch({
            type: ActionsActionTypes.ACTIONS_SUCCESS,
            payload: resp.data,
        });
      } else {
          dispatch({
              type: ActionsActionTypes.ACTIONS_ERROR,
              payload: ACTIONS_ERROR
          });
      }
    }).catch((e)=>{
      dispatch({
        type: ActionsActionTypes.ACTIONS_ERROR,
        payload: ACTIONS_ERROR
      });
    })
  }
}

export const sortActions = (actions: Action[], dragIndex: number, hoverIndex: number) => {
  return async (dispatch: Dispatch<Actions>) => {
    let first  = {
      id: actions[dragIndex].id,
      order: hoverIndex,
    }
    let second = {
      id: actions[hoverIndex].id,
      order: dragIndex,
    }
    let token = GetTokenFromCookies()
    actionsService.actionsOrdersPost({first: first, second: second}, {headers: {"Authorization": token}}).then((resp)=>{
      if (resp.status === 200) {
        dispatch({
          type: ActionsActionTypes.ACTIONS_SORT,
          payload: {
            dragIndex: dragIndex,
            hoverIndex: hoverIndex
          },
        });
      }
    })
  }
}

export const addNewAction = (file: any, title: string, description: string, date: string, callback:(error: string) => void) => {
  return async (dispatch: Dispatch<Actions>) => {
    if (!file) {
      callback("не приложили файл")
      return
    }
    if (!title) {
      callback("не все поля заполненны")
      return
    }
    let token = GetTokenFromCookies()
    actionsService.actionsPut(file, title, description, date, {headers: {"Authorization": token}}).then((resp)=> {
      if (resp.status === 200) {
        dispatch({
          type: ActionsActionTypes.ACTIONS_NEW,
          payload: resp.data,
        });
        callback("")
      } else {
        callback("ошибка при добавлении")
      }
    })
    
  }
}

export const updateAction = (action: Action, callback:(error: string) => void, file?: any) => {
  return async (dispatch: Dispatch<Actions>) => {
    let token = GetTokenFromCookies()
    if (action.id?.length === 0 || action.title?.length === 0 || action.order === undefined) {
      callback("не все поля заполненны")
      return
    }
    actionsService.actionsPatch(file, action.id!, action.title!, action.description!, action.date!, action.order!, action.photo, action.preview, {headers: {"Authorization": token}}).then((resp)=>{
      if (resp.status === 200) {
        dispatch({
          type: ActionsActionTypes.ACTIONS_UPDATE,
          payload: resp.data,
        });
        callback("")
      } else {
        callback("ошибка при редактировании")
      }
    })
  }
}

export const deleteAction = (id: string) => {
  return async (dispatch: Dispatch<Actions>) => {
    let token = GetTokenFromCookies()
    actionsService.actionsActionIDDelete(id, {headers: {"Authorization": token}}).then((resp)=>{
      if (resp.status === 200) {
        dispatch({
          type: ActionsActionTypes.ACTIONS_DELETE,
          payload: id,
        });
      }
    })
  }
}