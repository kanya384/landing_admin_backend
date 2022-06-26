import { Title } from "../../api"
import { TitlesActionTypes } from "../action-types/titles"

export interface TitlesRequestSend {
  type: TitlesActionTypes.TITLES_REQUEST_SEND,
}

export interface TitlesError {
  type: TitlesActionTypes.TITLES_ERROR,
  payload: string,
}

export interface TitlesNew {
  type: TitlesActionTypes.TITLES_NEW,
  payload: Title,
}

export interface TitlesSuccess {
  type: TitlesActionTypes.TITLES_SUCCESS,
  payload: Title[],
}

export interface TitlesDelete {
  type: TitlesActionTypes.TITLES_DELETE,
  payload: string
}

export interface TitlesUpdate {
  type: TitlesActionTypes.TITLES_UPDATE,
  payload: Title,
}

export type TitlesAction =
  | TitlesRequestSend
  | TitlesError
  | TitlesSuccess
  | TitlesNew
  | TitlesUpdate
  | TitlesDelete