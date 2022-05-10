import produce from "immer";
import { Plan } from "../../api";
import { PlansActionTypes } from "../action-types";
import { PlansAction } from "../actions";

interface PlansState {
  loading: boolean,
  plansList: Plan[],
  error: string | null,
}

const initialState: PlansState = {
  loading: false,
  plansList: [],
  error: null
}


const reducer = produce((state: PlansState = initialState, action: PlansAction) => {
  switch (action.type) {
    case PlansActionTypes.PLANS_REQUEST_SEND:
      state.loading = true
      return state;
    case PlansActionTypes.PLANS_SUCCESS:
      state.loading = false
      state.plansList = action.payload;
      return state
    case PlansActionTypes.PLANS_ERROR:
      state.loading = false;
      state.error = action.payload
      return state
    case PlansActionTypes.PLANS_UPDATE_PHOTO:
      for (let i=0; i<state.plansList.length; i++) {
        if (state.plansList[i].id === action.payload.id) {
          state.plansList[i] = action.payload
        }
      }
      return state;
    case PlansActionTypes.PLANS_UPDATE_ACTIVITY:
      return state;
    default:
      return state;
  }
})

export default reducer;