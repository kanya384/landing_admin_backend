import { Dispatch } from "react";
import { Configuration, PostersApi } from "../../api";
import { useApi } from "../../hooks/use-api";
import { PostersActionTypes } from "../action-types";
import { Posters } from "../actions";

const POSTERS_ERROR = "Ошибка при получении данных"
const url = "http://localhost:3000"
const postersService = new PostersApi(new Configuration(), url)
export const getPosters = () => {
  return async (dispatch: Dispatch<Posters>) => {
    dispatch({
      type: PostersActionTypes.POSTERS_REQUEST_SEND,
    });
    postersService.postersGet().then((resp)=>{
      if (resp.status === 200) {
        dispatch({
            type: PostersActionTypes.POSTERS_SUCCESS,
            payload: resp.data,
        });
      } else {
          dispatch({
              type: PostersActionTypes.POSTERS_ERROR,
              payload: POSTERS_ERROR
          });
      }
    }).catch((e)=>{
      dispatch({
        type: PostersActionTypes.POSTERS_ERROR,
        payload: POSTERS_ERROR
      });
    })
  }
}