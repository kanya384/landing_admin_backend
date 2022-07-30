import { Action } from "../../api";
import { ActionsActionTypes } from "../action-types";

export interface ActionsRequestSend {
    type: ActionsActionTypes.ACTIONS_REQUEST_SEND,
}

export interface ActionsSuccess {
    type: ActionsActionTypes.ACTIONS_SUCCESS,
    payload: Action[],
}

export interface ActionsError {
    type: ActionsActionTypes.ACTIONS_ERROR,
    payload: string,
}

export interface ActionsSort {
    type: ActionsActionTypes.ACTIONS_SORT,
    payload: {
        dragIndex: number,
        hoverIndex: number
    },
}
export interface ActionsNew {
    type: ActionsActionTypes.ACTIONS_NEW,
    payload: Action,
}

export interface ActionsUpdate {
    type: ActionsActionTypes.ACTIONS_UPDATE,
    payload: Action,
}

export interface ActionsDelete {
    type: ActionsActionTypes.ACTIONS_DELETE,
    payload: string,
}

export type Actions = 
 | ActionsRequestSend
 | ActionsSuccess
 | ActionsError
 | ActionsSort
 | ActionsNew
 | ActionsUpdate
 | ActionsDelete