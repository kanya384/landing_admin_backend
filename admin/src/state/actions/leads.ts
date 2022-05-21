import { Analytics, Lead } from "../../api";
import { LeadsActionTypes } from "../action-types";

export interface LeadsRequestSend {
  type: LeadsActionTypes.LEADS_REQUEST_SEND,
}

export interface LeadsSuccess {
  type: LeadsActionTypes.LEADS_SUCCESS,
  payload: Lead[],
}

export interface LeadsAnalyticsSuccess {
  type: LeadsActionTypes.LEADS_ANALYTICS_SUCCESS,
  payload: Analytics,
}

export interface LeadsError {
  type: LeadsActionTypes.LEADS_ERROR,
  payload: string,
}

export interface LeadsDelete {
  type: LeadsActionTypes.LEADS_DELETE,
  payload: string,
}

export type LeadsAction = 
  | LeadsRequestSend
  | LeadsSuccess
  | LeadsAnalyticsSuccess
  | LeadsError
  | LeadsDelete