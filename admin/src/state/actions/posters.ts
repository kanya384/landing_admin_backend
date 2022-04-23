import { Poster } from "../../api";
import { PostersActionTypes } from "../action-types";

export interface PostersRequestSend {
    type: PostersActionTypes.POSTERS_REQUEST_SEND,
}

export interface PostersSuccess {
    type: PostersActionTypes.POSTERS_SUCCESS,
    payload: Poster[],
}

export interface PostersError {
    type: PostersActionTypes.POSTERS_ERROR,
    payload: string,
}

export interface PostersSort {
    type: PostersActionTypes.POSTERS_SORT,
    payload: {
        dragIndex: number,
        hoverIndex: number
    },
}

export type Posters = 
 | PostersRequestSend
 | PostersSuccess
 | PostersError
 |PostersSort