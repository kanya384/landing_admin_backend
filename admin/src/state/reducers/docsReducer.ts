import produce from "immer";
import { Doc } from "../../api";
import { DocsActionTypes } from "../action-types";
import { DocAction } from "../actions/docs";

interface DocsState {
  loading: boolean,
  docsList: Doc[],
  error: string | null
}

const initialState: DocsState = {
  loading: false,
  docsList: [],
  error: null
}

const reducer = produce((state: DocsState = initialState, action: DocAction) =>{
  switch (action.type) {
    case DocsActionTypes.DOCS_REQUEST_SEND:
      state.loading = true
      return state
    case DocsActionTypes.DOCS_SUCCESS:
      state.loading = false
      state.docsList = action.payload
      return state
    case DocsActionTypes.DOCS_ERROR:
      state.loading = false
      state.error = action.payload
      return state
    case DocsActionTypes.DOCS_NEW:
      state.docsList.unshift(action.payload)
      return state;
    case DocsActionTypes.DOCS_UPDATE:
      for (let i=0; i<state.docsList.length; i++) {
        if (state.docsList[i].id === action.payload.id) {
          state.docsList[i] = action.payload
        }
      }
      return state;
    case DocsActionTypes.DOCS_DELETE:
      state.docsList = state.docsList.filter((value) => value.id !== action.payload)
      return state;
    default:
      return state;
  }
})

export default reducer;