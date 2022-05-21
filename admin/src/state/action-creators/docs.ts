import { Dispatch } from "react"
import { Configuration, DocsApi } from "../../api"
import { GetTokenFromCookies } from "../../utils"
import { DocsActionTypes } from "../action-types"
import { DocAction } from "../actions/docs"

const DOCS_ERROR = "Ошибка при получении данных"
const docsService = new DocsApi(new Configuration(), process.env.REACT_APP_BACKEND_URL)
export const getDocs = () => {
  return async (dispatch: Dispatch<DocAction>) => {
    dispatch({
      type: DocsActionTypes.DOCS_REQUEST_SEND,
    })
    let token = GetTokenFromCookies()
    docsService.docGet({headers: {"Authorization": token}}).then((resp) => {
      if (resp.status === 200) {
        dispatch({
          type: DocsActionTypes.DOCS_SUCCESS,
          payload: resp.data,
        })
      } else {
        dispatch({
          type: DocsActionTypes.DOCS_ERROR,
          payload: DOCS_ERROR,
        })
      }
    }).catch((e) => {
      dispatch({
        type: DocsActionTypes.DOCS_ERROR,
        payload: DOCS_ERROR,
      })
    })
  }
}

export const addNewDoc = (file: any, title: string,  callback:(error: string) => void) => {
  return async (dispatch: Dispatch<DocAction>) => {
    if (!file) {
      callback("не приложили файл")
      return
    }
    if (!title) {
      callback("не заполнен тайтл")
      return
    }
    let token = GetTokenFromCookies()
    docsService.docPut(file, title, {headers: {"Authorization": token}}).then((resp)=> {
      if (resp.status === 200) {
        dispatch({
          type: DocsActionTypes.DOCS_NEW,
          payload: resp.data,
        });
        callback("")
      } else {
        callback("ошибка при добавлении")
      }
    })
    
  }
}

export const editDoc = (file: any, id: string, title: string,  callback:(error: string) => void) => {
  return async (dispatch: Dispatch<DocAction>) => {
    if (!title || !id) {
      callback("не заполнен тайтл")
      return
    }
    let token = GetTokenFromCookies()
    if (typeof file === "string") {
      file = undefined
    }
    docsService.docPatch(id, title, file, {headers: {"Authorization": token}}).then((resp)=> {
      if (resp.status === 200) {
        dispatch({
          type: DocsActionTypes.DOCS_NEW,
          payload: resp.data,
        });
        callback("")
      } else {
        callback("ошибка при добавлении")
      }
    })
    
  }
}
