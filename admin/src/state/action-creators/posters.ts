import { Dispatch } from "react";
import { Configuration, PostersApi, Poster } from "../../api";
import { GetTokenFromCookies } from "../../utils";
import { PostersActionTypes } from "../action-types";
import { Posters } from "../actions";

const POSTERS_ERROR = "Ошибка при получении данных"
const url = "http://localhost:3000"
const postersService = new PostersApi(new Configuration(), url)
export const getPosters = () => {
  return async (dispatch: Dispatch<Posters>) => {
    dispatch({
      type: PostersActionTypes.POSTERS_REQUEST_SEND,
    });
    postersService.postersGet().then((resp)=>{
      if (resp.status === 200) {
        dispatch({
            type: PostersActionTypes.POSTERS_SUCCESS,
            payload: resp.data,
        });
      } else {
          dispatch({
              type: PostersActionTypes.POSTERS_ERROR,
              payload: POSTERS_ERROR
          });
      }
    }).catch((e)=>{
      dispatch({
        type: PostersActionTypes.POSTERS_ERROR,
        payload: POSTERS_ERROR
      });
    })
  }
}

export const sortPosters = (posters: Poster[], dragIndex: number, hoverIndex: number) => {
  return async (dispatch: Dispatch<Posters>) => {
    let first  = {
      id: posters[dragIndex].id,
      order: hoverIndex,
    }
    let second = {
      id: posters[hoverIndex].id,
      order: dragIndex,
    }
    let token = GetTokenFromCookies()
    postersService.postersOrdersPost({first: first, second: second}, {headers: {"Authorization": token}}).then((resp)=>{
      if (resp.status === 200) {
        dispatch({
          type: PostersActionTypes.POSTERS_SORT,
          payload: {
            dragIndex: dragIndex,
            hoverIndex: hoverIndex
          },
        });
      }
    })
  }
}

export const addNewPoster = (file: any, title: string, description: string, callback:(error: string) => void) => {
  return async (dispatch: Dispatch<Posters>) => {
    if (!file) {
      callback("не приложили файл")
      return
    }
    if (!title || !description) {
      callback("не заполнили одно из полей")
      return
    }
    let token = GetTokenFromCookies()
    postersService.postersPut(file, title, description, {headers: {"Authorization": token}}).then((resp)=> {
      if (resp.status === 200) {
        dispatch({
          type: PostersActionTypes.POSTERS_NEW,
          payload: resp.data,
        });
        callback("")
      } else {
        callback("ошибка при добавлении")
      }
    })
    
  }
}

export const updatePoster = (poster: Poster, callback:(error: string) => void, file?: any) => {
  return async (dispatch: Dispatch<Posters>) => {
    let token = GetTokenFromCookies()
    if (poster.id?.length === 0 || poster.title?.length === 0 || poster.description?.length === 0 || poster.active === undefined || poster.order === undefined) {
      console.log('sdsd')
      callback("не заполнили одно из полей")
      return
    }
    postersService.postersPatch(poster.id!, poster.title!, poster.description!, poster.active!, poster.order!, file, poster.photo, {headers: {"Authorization": token}}).then((resp)=>{
      if (resp.status === 200) {
        dispatch({
          type: PostersActionTypes.POSTERS_UPDATE,
          payload: poster,
        });
        callback("")
      } else {
        callback("ошибка при редактировании")
      }
    })
  }
}

export const deletePoster = (id: string) => {
  return async (dispatch: Dispatch<Posters>) => {
    let token = GetTokenFromCookies()
    postersService.postersPosterIDDelete(id, {headers: {"Authorization": token}}).then((resp)=>{
      if (resp.status === 200) {
        dispatch({
          type: PostersActionTypes.POSTERS_DELETE,
          payload: id,
        });
      }
    })
  }
}