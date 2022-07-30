import produce from 'immer';
import { Action } from '../../api';
import { ActionsActionTypes } from '../action-types';
import { Actions } from '../actions';


interface ActionsState {
  loading: boolean;
  actionsList: Action[],
  error: string | null,
}

const initialState: ActionsState = {
  loading: false,
  actionsList: [],
  error: null,
}

const sortActions = (actionsList: Action[], dragIndex: number, hoverIndex: number):Action[]  => {
  let tmpAction =  actionsList[dragIndex]
  actionsList[dragIndex] = actionsList[hoverIndex]
  actionsList[hoverIndex] = tmpAction
  return actionsList
}



const reducer = produce((state: ActionsState = initialState, action: Actions ) => {
  switch (action.type) {
    case ActionsActionTypes.ACTIONS_REQUEST_SEND:
      state.loading = true;
      return state;
    case ActionsActionTypes.ACTIONS_SUCCESS:
      state.loading = false;
      state.actionsList = action.payload;
      return state;
    case ActionsActionTypes.ACTIONS_ERROR:
      state.loading = false;
      state.error = action.payload;
      return state;
    case ActionsActionTypes.ACTIONS_NEW:
      state.actionsList.unshift(action.payload)
      return state;
    case ActionsActionTypes.ACTIONS_UPDATE:
      for (let i=0; i<state.actionsList.length; i++) {
        if (state.actionsList[i].id === action.payload.id) {
          state.actionsList[i] = action.payload
        }
      }
      return state;
    case ActionsActionTypes.ACTIONS_DELETE:
      state.actionsList = state.actionsList.filter((value) => value.id !== action.payload)
      return state;
    case ActionsActionTypes.ACTIONS_SORT:
      state.actionsList = sortActions(state.actionsList, action.payload.dragIndex, action.payload.hoverIndex);
      return state;
    default:
      return state;
  }
})

export default reducer;