import { Dispatch } from "react";
import { Configuration, ProjectInfoApi, ProjectInfo } from "../../api";
import { GetTokenFromCookies } from "../../utils";
import { ProjectInfosActionTypes } from "../action-types";
import { ProjectInfos } from "../actions";

const PROJECT_INFOS_ERROR = "Ошибка при получении данных"
const projectInfosService = new ProjectInfoApi(new Configuration(), process.env.REACT_APP_BACKEND_URL)
export const getProjectInfos = () => {
  return async (dispatch: Dispatch<ProjectInfos>) => {
    dispatch({
      type: ProjectInfosActionTypes.PROJECT_INFOS_REQUEST_SEND,
    });
    projectInfosService.projectInfoGet().then((resp)=>{
      if (resp.status === 200) {
        dispatch({
            type: ProjectInfosActionTypes.PROJECT_INFOS_SUCCESS,
            payload: resp.data,
        });
      } else {
          dispatch({
              type: ProjectInfosActionTypes.PROJECT_INFOS_ERROR,
              payload: PROJECT_INFOS_ERROR
          });
      }
    }).catch((e)=>{
      dispatch({
        type: ProjectInfosActionTypes.PROJECT_INFOS_ERROR,
        payload: PROJECT_INFOS_ERROR
      });
    })
  }
}

export const sortProjectInfos = (projectInfos: ProjectInfo[], dragIndex: number, hoverIndex: number) => {
  return async (dispatch: Dispatch<ProjectInfos>) => {
    let first  = {
      id: projectInfos[dragIndex].id,
      order: hoverIndex,
    }
    let second = {
      id: projectInfos[hoverIndex].id,
      order: dragIndex,
    }
    let token = GetTokenFromCookies()
    projectInfosService.projectInfoOrdersPost({first: first, second: second}, {headers: {"Authorization": token}}).then((resp)=>{
      if (resp.status === 200) {
        dispatch({
          type: ProjectInfosActionTypes.PROJECT_INFOS_SORT,
          payload: {
            dragIndex: dragIndex,
            hoverIndex: hoverIndex
          },
        });
      }
    })
  }
}

export const addNewProjectInfo = (file: any, title: string, anonce: string, description: string, callback:(error: string) => void) => {
  return async (dispatch: Dispatch<ProjectInfos>) => {
    if (!file) {
      callback("не приложили файл")
      return
    }
    if (!title) {
      callback("не все поля заполненны")
      return
    }
    let token = GetTokenFromCookies()
    projectInfosService.projectInfoPut(file, title, anonce, description, {headers: {"Authorization": token}}).then((resp)=> {
      if (resp.status === 200) {
        dispatch({
          type: ProjectInfosActionTypes.PROJECT_INFOS_NEW,
          payload: resp.data,
        });
        callback("")
      } else {
        callback("ошибка при добавлении")
      }
    })
    
  }
}

export const updateProjectInfo = (projectInfo: ProjectInfo, callback:(error: string) => void, file?: any) => {
  return async (dispatch: Dispatch<ProjectInfos>) => {
    let token = GetTokenFromCookies()
    if (projectInfo.id?.length === 0 || projectInfo.title?.length === 0 || projectInfo.description?.length === 0  || projectInfo.anonce?.length === 0 || projectInfo.order === undefined) {
      callback("не все поля заполненны")
      return
    }
    projectInfosService.projectInfoPatch(projectInfo.id!, projectInfo.title!, projectInfo.order!, projectInfo.anonce!, projectInfo.description!, file, projectInfo.photo!, {headers: {"Authorization": token}}).then((resp)=>{
      if (resp.status === 200) {
        dispatch({
          type: ProjectInfosActionTypes.PROJECT_INFOS_UPDATE,
          payload: resp.data,
        });
        callback("")
      } else {
        callback("ошибка при редактировании")
      }
    })
  }
}

export const deleteProjectInfo = (id: string) => {
  return async (dispatch: Dispatch<ProjectInfos>) => {
    let token = GetTokenFromCookies()
    projectInfosService.projectInfoPojectInfoIDDelete(id, {headers: {"Authorization": token}}).then((resp)=>{
      if (resp.status === 200) {
        dispatch({
          type: ProjectInfosActionTypes.PROJECT_INFOS_DELETE,
          payload: id,
        });
      }
    })
  }
}