import { Plan } from "../../api";
import { PlansActionTypes } from "../action-types/plans";

export interface PlansRequestSend {
  type: PlansActionTypes.PLANS_REQUEST_SEND,
}

export interface PlansError {
  type: PlansActionTypes.PLANS_ERROR,
  payload: string,
}

export interface PlansSuccess {
  type: PlansActionTypes.PLANS_SUCCESS,
  payload: Plan[],
}

export interface PlansUpdatePhoto {
  type: PlansActionTypes.PLANS_UPDATE_PHOTO,
  payload: Plan,
}

export interface PlansUpdateActivity {
  type: PlansActionTypes.PLANS_UPDATE_ACTIVITY,
  payload: {
    id: string,
    status: boolean,
  }
}

export type PlansAction =
  | PlansRequestSend
  | PlansError
  | PlansSuccess
  | PlansUpdatePhoto
  | PlansUpdateActivity