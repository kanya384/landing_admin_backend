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

export interface SettingsCreateUpdate {
  type: SettingsActionTypes.SETTINGS_CREATE_UPDATE,
  payload: string,
}

export type SettingsAction =
  | SettingsRequestSend
  | SettingsError
  | SettingsSuccess
  | SettingsCreateUpdate