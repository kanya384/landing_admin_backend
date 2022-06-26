import { Dispatch } from "react"
import {  Configuration, Setting, SettingsApi } from "../../api"
import { GetTokenFromCookies } from "../../utils"
import { SettingsActionTypes } from "../action-types"
import { SettingsAction } from "../actions"

const SETTINGS_ERROR = "Ошибка при получении данных"
const settingsService = new SettingsApi(new Configuration(), process.env.REACT_APP_BACKEND_URL)

export const getSettings = () => {
  return async (dispatch: Dispatch<SettingsAction>) => {
    dispatch({
      type: SettingsActionTypes.SETTINGS_REQUEST_SEND
    })

    let token = GetTokenFromCookies()
    settingsService.settingsGet({headers: {"Authorization": token}}).then((resp)=>{
      if (resp.status === 200) {
        dispatch({
          type: SettingsActionTypes.SETTINGS_SUCCESS,
          payload: resp.data,
        })
      } else {
        dispatch({
          type: SettingsActionTypes.SETTINGS_ERROR,
          payload: SETTINGS_ERROR,
        })
      }
    }).catch((e) => {
      dispatch({
        type: SettingsActionTypes.SETTINGS_ERROR,
        payload: SETTINGS_ERROR,
      })
    })
  }
}

export const addUpdateSetting = (setting: Setting, callback: (error: string) => void) => {
  return async (dispatch: Dispatch<SettingsAction>) => {
    if (!setting.name || ! setting.value) {
      callback("не все поля заполненны")
      return
    }
    let token = GetTokenFromCookies()
    settingsService.settingsPut(setting, {headers: {"Authorization": token}}).then((resp) => {
      if (resp.status === 200) {
        dispatch({
          type: SettingsActionTypes.SETTINGS_CREATE_UPDATE,
          payload: resp.data,
        });
        callback("")
      } else {
        callback("ошибка при добавлении")
      }
    })
  }
}