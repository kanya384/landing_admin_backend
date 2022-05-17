import produce from "immer";
import { Video } from "../../api";
import { VideosActionTypes } from "../action-types";
import { VideosAction } from '../actions';

interface VideosState {
  loading: boolean,
  videoList: Video[],
  error: string | null,
}

const initialState: VideosState = {
  loading: false,
  videoList: [],
  error: null
}

const reducer = produce((state: VideosState = initialState, action: VideosAction) => {
  switch (action.type) {
    case VideosActionTypes.VIDEOS_REQUEST_SEND:
      state.loading = true;
      return
    case VideosActionTypes.VIDEOS_SUCCESS:
      state.loading = false;
      state.videoList = action.payload;
      return
    case VideosActionTypes.VIDEOS_ERROR:
      state.loading = false;
      state.error = action.payload;
      return
    case VideosActionTypes.VIDEOS_NEW:
      state.videoList.unshift(action.payload)
      return state;
    case VideosActionTypes.VIDEOS_UPDATE:
      for (let i=0; i<state.videoList.length; i++) {
        if (state.videoList[i].id === action.payload.id) {
          state.videoList[i] = action.payload
        }
      }
      return state;
    case VideosActionTypes.VIDEOS_DELETE:
      state.videoList = state.videoList.filter((value) => value.id !== action.payload)
      return state;
    default:
      return state;
  }
})

export default reducer;