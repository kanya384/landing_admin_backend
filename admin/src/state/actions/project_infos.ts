import { Poster } from "../../api";
import { ProjectInfosActionTypes } from "../action-types";


export interface ProjectInfosRequestSend {
    type: ProjectInfosActionTypes.PROJECT_INFOS_REQUEST_SEND,
}

export interface ProjectInfosSuccess {
    type: ProjectInfosActionTypes.PROJECT_INFOS_SUCCESS,
    payload: Poster[],
}

export interface ProjectInfosError {
    type: ProjectInfosActionTypes.PROJECT_INFOS_ERROR,
    payload: string,
}

export interface ProjectInfosSort {
    type: ProjectInfosActionTypes.PROJECT_INFOS_SORT,
    payload: {
        dragIndex: number,
        hoverIndex: number
    },
}
export interface ProjectInfosNew {
    type: ProjectInfosActionTypes.PROJECT_INFOS_NEW,
    payload: Poster,
}

export interface ProjectInfosUpdate {
    type: ProjectInfosActionTypes.PROJECT_INFOS_UPDATE,
    payload: Poster,
}

export interface ProjectInfosDelete {
    type: ProjectInfosActionTypes.PROJECT_INFOS_DELETE,
    payload: string,
}

export type ProjectInfos = 
 | ProjectInfosRequestSend
 | ProjectInfosSuccess
 | ProjectInfosError
 | ProjectInfosSort
 | ProjectInfosNew
 | ProjectInfosUpdate
 | ProjectInfosDelete