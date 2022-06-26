import produce from "immer";
import { Setting } from "../../api";
import { SettingsActionTypes } from "../action-types";
import { SettingsAction } from '../actions';

interface SettingsState {
  loading: boolean,
  settingsList: Setting[],
  error: string | null,
}

const initialState: SettingsState = {
  loading: false,
  settingsList: [],
  error: null
}

const reducer = produce((state: SettingsState = initialState, action: SettingsAction) => {
  switch (action.type) {
    case SettingsActionTypes.SETTINGS_REQUEST_SEND:
      state.loading = true;
      return
    case SettingsActionTypes.SETTINGS_SUCCESS:
      state.loading = false;
      state.settingsList = action.payload;
      return
    case SettingsActionTypes.SETTINGS_ERROR:
      state.loading = false;
      state.error = action.payload;
      return
    case SettingsActionTypes.SETTINGS_CREATE_UPDATE:
      for (let i=0; i<state.settingsList.length; i++) {
        /*if (state.settingsList[i].id === action.payload.id) {
          state.settingsList[i] = action.payload
        }*/
      }
      return state;
    default:
      return state;
  }
})

export default reducer;