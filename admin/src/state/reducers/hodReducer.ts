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
    case HodActionTypes.HOD_MONTHS_SUCCESS:
      state.loadingMonths = false;
      state.monthsList = action.payload;
      return state;
    case HodActionTypes.HOD_PHOTOS_SUCCESS:
      state.loadingPhotos = false;
      state.photosList = action.payload;
      return state;
    default:
      return state;
  }
})

export default reducer;