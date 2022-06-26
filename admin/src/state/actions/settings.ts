import { SettingsActionTypes } from "../action-types/settings";
import { Setting } from "../../api"

export interface SettingsRequestSend {
  type: SettingsActionTypes.SETTINGS_REQUEST_SEND,
}

export interface SettingsError {
  type: SettingsActionTypes.SETTINGS_ERROR,
  payload: string,
}

export interface SettingsSuccess {
  type: SettingsActionTypes.SETTINGS_SUCCESS,
  payload: Setting[],
}

export interface SettingsCreate {
  type: SettingsActionTypes.SETTINGS_CREATE,
  payload: Setting,
}

export interface SettingsUpdate {
  type: SettingsActionTypes.SETTINGS_UPDATE,
  payload: Setting,
}

export type SettingsAction =
  | SettingsRequestSend
  | SettingsError
  | SettingsSuccess
  | SettingsCreate
  | SettingsUpdate