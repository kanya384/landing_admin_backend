import { Video } from "../../api"
import { VideosActionTypes } from "../action-types"

export interface VideosRequestSend {
  type: VideosActionTypes.VIDEOS_REQUEST_SEND,
}

export interface VideosError {
  type: VideosActionTypes.VIDEOS_ERROR,
  payload: string,
}

export interface VideosNew {
  type: VideosActionTypes.VIDEOS_NEW,
  payload: Video,
}

export interface VideosSuccess {
  type: VideosActionTypes.VIDEOS_SUCCESS,
  payload: Video[],
}

export interface VideosDelete {
  type: VideosActionTypes.VIDEOS_DELETE,
  payload: string
}

export interface VideosUpdate {
  type: VideosActionTypes.VIDEOS_UPDATE,
  payload: Video,
}

export type VideosAction =
  | VideosRequestSend
  | VideosError
  | VideosSuccess
  | VideosNew
  | VideosUpdate
  | VideosDelete