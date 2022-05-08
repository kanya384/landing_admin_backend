import produce from 'immer';
import { HodPhoto, Month, Year } from '../../api';
import { HodActionTypes } from '../action-types';
import { Hod } from '../actions';


interface HodState {
  loadingYears: boolean
  yearsList: Year[]
  loadingMonths: boolean
  monthsList: Month[]
  loadingPhotos: boolean
  photosList: HodPhoto[],
  error: string | null,
}

const initialState: HodState = {
  loadingYears: false,
  yearsList: [],
  loadingMonths: false,
  monthsList: [],
  loadingPhotos: false,
  photosList: [],
  error: null,
}

const sortPhotos = (photosList: HodPhoto[], dragIndex: number, hoverIndex: number):HodPhoto[]  => {
  let tmpPhoto =  photosList[dragIndex]
  photosList[dragIndex] = photosList[hoverIndex]
  photosList[hoverIndex] = tmpPhoto
  return photosList
}



const reducer = produce((state: HodState = initialState, action: Hod ) => {
  switch (action.type) {
    case HodActionTypes.HOD_ERROR:
      state.loadingYears = false;
      state.loadingMonths = false;
      state.loadingPhotos = false;
      return state;
    case HodActionTypes.HOD_YEARS_SUCCESS:
      state.loadingYears = false;
      state.yearsList = action.payload;
      return state;
    case HodActionTypes.HOD_YEARS_NEW:
      state.yearsList.unshift(action.payload)
      return state;
    case HodActionTypes.HOD_YEARS_DELETE:
      state.yearsList = state.yearsList.filter((value) => value.id !== action.payload)
      return state;
    case HodActionTypes.HOD_MONTHS_SUCCESS:
      state.loadingMonths = false;
      state.monthsList = action.payload;
      return state;
    case HodActionTypes.HOD_MONTHS_NEW:
      state.monthsList.unshift(action.payload)
      return state;
    case HodActionTypes.HOD_MONTHS_DELETE:
      state.monthsList = state.monthsList.filter((value) => value.id !== action.payload)
      return state;
    case HodActionTypes.HOD_PHOTOS_SUCCESS:
      state.loadingPhotos = false;
      state.photosList = action.payload;
      return state;
    case HodActionTypes.HOD_PHOTOS_NEW:
      state.photosList.unshift(action.payload)
      return state;
    case HodActionTypes.HOD_PHOTOS_DELETE:
      state.photosList = state.photosList.filter((value) => value.id !== action.payload)
      return state;
    case HodActionTypes.HOD_PHOTOS_SORT:
      state.photosList = sortPhotos(state.photosList, action.payload.dragIndex, action.payload.hoverIndex);
      return state;
    default:
      return state;
  }
})

export default reducer;