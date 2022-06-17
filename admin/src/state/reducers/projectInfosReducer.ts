import produce from 'immer';
import { ProjectInfo } from '../../api';
import { ProjectInfosActionTypes } from '../action-types';
import { ProjectInfos } from '../actions';


interface ProjectInfoState {
  loading: boolean;
  projectInfoList: ProjectInfo[],
  error: string | null,
}

const initialState: ProjectInfoState = {
  loading: false,
  projectInfoList: [],
  error: null,
}

const sortProjectInfos = (projectInfoList: ProjectInfo[], dragIndex: number, hoverIndex: number):ProjectInfo[]  => {
  let tmpProjectInfo =  projectInfoList[dragIndex]
  projectInfoList[dragIndex] = projectInfoList[hoverIndex]
  projectInfoList[hoverIndex] = tmpProjectInfo
  return projectInfoList
}



const reducer = produce((state: ProjectInfoState = initialState, action: ProjectInfos ) => {
  switch (action.type) {
    case ProjectInfosActionTypes.PROJECT_INFOS_REQUEST_SEND:
      state.loading = true;
      return state;
    case ProjectInfosActionTypes.PROJECT_INFOS_SUCCESS:
      state.loading = false;
      state.projectInfoList = action.payload;
      return state;
    case ProjectInfosActionTypes.PROJECT_INFOS_ERROR:
      state.loading = false;
      state.error = action.payload;
      return state;
    case ProjectInfosActionTypes.PROJECT_INFOS_NEW:
      state.projectInfoList.unshift(action.payload)
      return state;
    case ProjectInfosActionTypes.PROJECT_INFOS_UPDATE:
      for (let i=0; i<state.projectInfoList.length; i++) {
        if (state.projectInfoList[i].id === action.payload.id) {
          state.projectInfoList[i] = action.payload
        }
      }
      return state;
    case ProjectInfosActionTypes.PROJECT_INFOS_DELETE:
      state.projectInfoList = state.projectInfoList.filter((value) => value.id !== action.payload)
      return state;
    case ProjectInfosActionTypes.PROJECT_INFOS_SORT:
      state.projectInfoList = sortProjectInfos(state.projectInfoList, action.payload.dragIndex, action.payload.hoverIndex);
      return state;
    default:
      return state;
  }
})

export default reducer;