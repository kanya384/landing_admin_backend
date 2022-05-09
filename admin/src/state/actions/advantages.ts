import { Advantage, AdvantagePhoto } from "../../api";
import { AdvantagesActionTypes } from "../action-types";

export interface AdvantagesRequestSend{
    type: AdvantagesActionTypes.ADVANTAGES_REQUEST_SEND
}

export interface AdvantagesRequestSuccess{
    type: AdvantagesActionTypes.ADVANTAGES_SUCCESS,
    payload: Advantage[]
}

export interface AdvantagesRequestError {
    type: AdvantagesActionTypes.ADVANTAGES_ERROR,
    payload: string,
}

export interface AdvantagesSort {
    type: AdvantagesActionTypes.ADVANTAGES_SORT,
    payload: {
        dragIndex: number,
        hoverIndex: number
    },
}

export interface AdvantageNew {
    type: AdvantagesActionTypes.ADVANTAGES_NEW,
    payload: Advantage
}

export interface AdvantageUpdate{
    type: AdvantagesActionTypes.ADVANTAGES_UPDATE,
    payload: Advantage,
}

export interface AdvantageDelete{
    type: AdvantagesActionTypes.ADVANTAGES_DELETE,
    payload: string,
}

export interface AdvantagePhotoRequest {
    type: AdvantagesActionTypes.ADVANTAGE_PHOTO_REQUEST,
}

export interface AdvantagePhotoNew{
    type: AdvantagesActionTypes.ADVANTAGE_PHOTO_NEW,
    payload: AdvantagePhoto,
}

export interface AdvantagePhotoDelete{
    type: AdvantagesActionTypes.ADVANTAGE_PHOTO_DELETE,
    payload: string,
}

export interface AdvantagePhotoSort{
    type: AdvantagesActionTypes.ADVANTAGE_PHOTO_SORT,
    payload: {
        dragIndex: number,
        hoverIndex: number
    },
}

export interface AdvantageIDRequestSend{
    type: AdvantagesActionTypes.ADVANTAGE_ID_REQUEST_SEND,
}

export interface AdvantageIDSuccess{
    type: AdvantagesActionTypes.ADVANTAGE_ID_SUCCESS,
    payload: Advantage,
}

export interface AdvantageIDError {
    type: AdvantagesActionTypes.ADVANTAGE_ID_ERROR,
    payload: string,
}

export interface AdvantagePhotosGet {
    type: AdvantagesActionTypes.ADVANTAGE_PHOTO_SUCCESS,
    payload: AdvantagePhoto[],
}


export type AdvantageAction = 
    | AdvantagesRequestSend
    | AdvantagesRequestSuccess
    | AdvantagesRequestError
    | AdvantagesSort
    | AdvantageNew
    | AdvantageUpdate
    | AdvantageDelete
    | AdvantagePhotoRequest
    | AdvantagePhotoNew
    | AdvantagePhotoDelete
    | AdvantagePhotoSort
    | AdvantagePhotosGet
    | AdvantageIDRequestSend
    | AdvantageIDSuccess
    | AdvantageIDError