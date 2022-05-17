import { Configuration, Video, VideoApi } from "../../api"
import { Dispatch } from "react";
import { VideosAction } from "../actions";
import { VideosActionTypes } from "../action-types";
import { GetTokenFromCookies } from "../../utils";

const VIDEOS_ERROR = "Ошибка при получении данных"
const url = "http://localhost:3000"
const videosService = new VideoApi(new Configuration(), url)

export const getVideos = () => {
  return async (dispatch: Dispatch<VideosAction>) => {
    dispatch({
      type: VideosActionTypes.VIDEOS_REQUEST_SEND
    })

    videosService.videoGet().then((resp)=>{
      if (resp.status === 200) {
        dispatch({
          type: VideosActionTypes.VIDEOS_SUCCESS,
          payload: resp.data,
        })
      } else {
        dispatch({
          type: VideosActionTypes.VIDEOS_ERROR,
          payload: VIDEOS_ERROR,
        })
      }
    }).catch((e) => {
      dispatch({
        type: VideosActionTypes.VIDEOS_ERROR,
        payload: VIDEOS_ERROR,
      })
    })
  }
}

export const addNewVideo = (file: any, url: string, callback: (error: string) => void) => {
  return async (dispatch: Dispatch<VideosAction>) => {
    if (!file) {
      callback("не приложили файл")
      return
    }
    if (!url) {
      callback("не заполненн url")
      return
    }
    let token = GetTokenFromCookies()
    videosService.videoPut(file, url, {headers: {"Authorization": token}}).then((resp) => {
      if (resp.status === 200) {
        dispatch({
          type: VideosActionTypes.VIDEOS_NEW,
          payload: resp.data,
        });
        callback("")
      } else {
        callback("ошибка при добавлении")
      }
    })
  }
}

export const updateVideo = (file: any, id: string, url: string, callback: (error: string) => void) => {
  return async (dispatch: Dispatch<VideosAction>) => {
    if (!file) {
      callback("не приложили файл")
      return
    }
    if (!url) {
      callback("не заполненн url")
      return
    }
    let token = GetTokenFromCookies()
    videosService.videoPatch(file, id, url, {headers: {"Authorization": token}}).then((resp) => {
      if (resp.status === 200) {
        dispatch({
          type: VideosActionTypes.VIDEOS_NEW,
          payload: resp.data,
        });
        callback("")
      } else {
        callback("ошибка при добавлении")
      }
    })
  }
}

export const deleteVideo = (id: string) => {
  return async (dispatch: Dispatch<VideosAction>) => {
    let token = GetTokenFromCookies()
    videosService.videoIdDelete(id,  {headers: {"Authorization": token}}).then((resp)=>{
      if (resp.status === 200) {
        dispatch({
          type: VideosActionTypes.VIDEOS_DELETE,
          payload: id,
        });
      }
    })
  }
}