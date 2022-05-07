import { HodPhoto, Month, Year } from "../../api";
import { HodActionTypes } from "../action-types";

export interface HodError {
  type: HodActionTypes.HOD_ERROR,
  payload: string,
}

/* years */
export interface HodYearsRequestSend {
  type: HodActionTypes.HOD_YEARS_REQUEST_SEND,
}

export interface HodYearsRequestSuccess {
  type: HodActionTypes.HOD_YEARS_SUCCESS,
  payload: Year[],
}

export interface HodYearsNew {
  type: HodActionTypes.HOD_YEARS_NEW,
  payload: Year,
}

export interface HodYearsUpdate {
  type: HodActionTypes.HOD_YEARS_UPDATE,
  payload: Year,
}

export interface HodYearsDelete {
  type: HodActionTypes.HOD_YEARS_DELETE,
  payload: string,
}



/* months */

export interface HodMonthsRequestSend {
  type: HodActionTypes.HOD_MONTHS_REQUEST_SEND,
}

export interface HodMonthsRequestSuccess {
  type: HodActionTypes.HOD_MONTHS_SUCCESS,
  payload: Month[],
}

export interface HodMonthsNew {
  type: HodActionTypes.HOD_MONTHS_NEW,
  payload: Month,
}

export interface HodMonthsUpdate {
  type: HodActionTypes.HOD_MONTHS_UPDATE,
  payload: Month,
}

export interface HodMonthsDelete {
  type: HodActionTypes.HOD_MONTHS_DELETE,
  payload: string,
}


/* photos */

export interface HodPhotosRequestSend {
  type: HodActionTypes.HOD_PHOTOS_REQUEST_SEND,
}

export interface HodPhotosRequestSuccess {
  type: HodActionTypes.HOD_PHOTOS_SUCCESS,
  payload: HodPhoto[],
}

export interface HodPhotosSort {
  type: HodActionTypes.HOD_PHOTOS_SORT,
  payload: {
    dragIndex: number,
    hoverIndex: number
  },
}

export interface HodPhotosNew {
  type: HodActionTypes.HOD_PHOTOS_NEW,
  payload: HodPhoto,
}

export interface HodPhotosUpdate {
  type: HodActionTypes.HOD_PHOTOS_UPDATE,
  payload: HodPhoto,
}

export interface HodPhotosDelete {
  type: HodActionTypes.HOD_PHOTOS_DELETE,
  payload: string,
}

export type Hod = 
  | HodError
  | HodYearsRequestSend
  | HodYearsRequestSuccess
  | HodYearsNew
  | HodYearsUpdate
  | HodYearsDelete
  | HodMonthsRequestSend
  | HodMonthsRequestSuccess
  | HodMonthsNew
  | HodMonthsUpdate
  | HodMonthsDelete
  | HodPhotosRequestSend
  | HodPhotosRequestSuccess
  | HodPhotosSort
  | HodPhotosNew
  | HodPhotosUpdate
  | HodPhotosDelete
