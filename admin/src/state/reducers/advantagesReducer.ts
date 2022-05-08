import produce from "immer";
import { Advantage, AdvantagePhoto } from "../../api";
import { AdvantagesActionTypes } from "../action-types";
import { AdvantageAction } from "../actions";

interface AdvantagesState {
    loading: boolean;
    advantagesList: Advantage[],
    photosList: AdvantagePhoto[],
    error: string | null,
}

const initialState: AdvantagesState = {
    loading: false,
    advantagesList: [],
    photosList: [],
    error: null,
}

const sortAdvantages = (advantagesList: Advantage[], dragIndex: number, hoverIndex: number):Advantage[]  => {
    let tmpPoster =  advantagesList[dragIndex]
    advantagesList[dragIndex] = advantagesList[hoverIndex]
    advantagesList[hoverIndex] = tmpPoster
    return advantagesList
}

const sortPhotos = (photosList: AdvantagePhoto[], dragIndex: number, hoverIndex: number):AdvantagePhoto[]  => {
    let tmpPhoto =  photosList[dragIndex]
    photosList[dragIndex] = photosList[hoverIndex]
    photosList[hoverIndex] = tmpPhoto
    return photosList
}

const reducer = produce((state: AdvantagesState = initialState, action: AdvantageAction) => {
    switch (action.type) {
        case AdvantagesActionTypes.ADVANTAGES_REQUEST_SEND:
            state.loading = true;
            return state;
        case AdvantagesActionTypes.ADVANTAGES_SUCCESS:
            state.loading = false;
            state.advantagesList = action.payload;
            return state;
        case AdvantagesActionTypes.ADVANTAGES_ERROR:
            state.loading = false;
            state.error = action.payload;
            return state;
        case AdvantagesActionTypes.ADVANTAGES_NEW:
            state.advantagesList.unshift(action.payload)
            return state
        case AdvantagesActionTypes.ADVANTAGES_UPDATE:
            for (let i=0; i<state.advantagesList.length; i++) {
                if (state.advantagesList[i].id == action.payload.id) {
                    state.advantagesList[i] = action.payload
                }
            }
            return state;
        case AdvantagesActionTypes.ADVANTAGES_DELETE:
            state.advantagesList = state.advantagesList.filter((value) => value.id != action.payload)
            return state;
        case AdvantagesActionTypes.ADVANTAGES_SORT:
            state.advantagesList = sortAdvantages(state.advantagesList, action.payload.dragIndex, action.payload.hoverIndex)
            return state;

        /* Photos */

        case AdvantagesActionTypes.ADVANTAGE_PHOTO_REQUEST:
            state.loading = true;
            return state;
        case AdvantagesActionTypes.ADVANTAGE_PHOTO_NEW:
            state.photosList.unshift(action.payload)
            return state;
        case AdvantagesActionTypes.ADVANTAGE_PHOTO_DELETE:
            state.photosList = state.photosList.filter((value) => value.id != action.payload)
            return state;
        case AdvantagesActionTypes.ADVANTAGE_PHOTO_SORT:
            state.photosList = sortPhotos(state.photosList, action.payload.dragIndex, action.payload.hoverIndex);
            return state;
        default:
            return state;
    }
})
  

export default reducer;