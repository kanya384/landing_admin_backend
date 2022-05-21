import produce from "immer";
import { Analytics, Lead } from "../../api";
import { LeadsActionTypes } from "../action-types";
import { LeadsAction } from "../actions/leads";

interface LeadsState {
  loading: boolean,
  leadsList: Lead[],
  analytics: null | Analytics,
  error: string | null
}

const initialState: LeadsState = {
  loading: false,
  leadsList: [],
  analytics: null,
  error: null
}

const reducer = produce((state: LeadsState = initialState, action: LeadsAction) =>{
  switch (action.type) {
    case LeadsActionTypes.LEADS_REQUEST_SEND:
      state.loading = true
      return state
    case LeadsActionTypes.LEADS_SUCCESS:
      state.loading = false
      state.leadsList = action.payload
      return state
    case LeadsActionTypes.LEADS_ERROR:
      state.loading = false
      state.error = action.payload
      return state
    case LeadsActionTypes.LEADS_ANALYTICS_SUCCESS:
      state.analytics = action.payload
      return state;
    case LeadsActionTypes.LEADS_DELETE:
      state.leadsList = state.leadsList.filter((value) => value.id !== action.payload)
      return state;
    default:
      return state;
  }
})

export default reducer;