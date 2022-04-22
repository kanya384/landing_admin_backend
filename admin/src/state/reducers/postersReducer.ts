import produce from 'immer';
import { Poster } from '../../api';
import { PostersActionTypes } from '../action-types';
import { Posters } from '../actions';


interface PostersState {
  loading: boolean;
  postersList: Poster[],
  error: string | null,
}

const initialState: PostersState = {
  loading: false,
  postersList: [],
  error: null,
}

const reducer = produce((state: PostersState = initialState, action: Posters ) => {
  switch (action.type) {
    case PostersActionTypes.POSTERS_REQUEST_SEND:
      state.loading = true;
      return state;
    case PostersActionTypes.POSTERS_SUCCESS:
      state.loading = false;
      state.postersList = action.payload;
      return state;
    case PostersActionTypes.POSTERS_ERROR:
      state.loading = false;
      state.error = action.payload;
      return state;
    default:
      return state;
  }
})

export default reducer;