import produce from "immer";
import { Title } from "../../api";
import { TitlesActionTypes } from "../action-types";
import { TitlesAction } from '../actions';

interface TitlesState {
  loading: boolean,
  titleList: Title[],
  error: string | null,
}

const initialState: TitlesState = {
  loading: false,
  titleList: [],
  error: null
}

const reducer = produce((state: TitlesState = initialState, action: TitlesAction) => {
  switch (action.type) {
    case TitlesActionTypes.TITLES_REQUEST_SEND:
      state.loading = true;
      return
    case TitlesActionTypes.TITLES_SUCCESS:
      state.loading = false;
      state.titleList = action.payload;
      return
    case TitlesActionTypes.TITLES_ERROR:
      state.loading = false;
      state.error = action.payload;
      return
    case TitlesActionTypes.TITLES_NEW:
      state.titleList.unshift(action.payload)
      return state;
    case TitlesActionTypes.TITLES_UPDATE:
      for (let i=0; i<state.titleList.length; i++) {
        if (state.titleList[i].id === action.payload.id) {
          state.titleList[i] = action.payload
        }
      }
      return state;
    case TitlesActionTypes.TITLES_DELETE:
      state.titleList = state.titleList.filter((value) => value.id !== action.payload)
      return state;
    default:
      return state;
  }
})

export default reducer;