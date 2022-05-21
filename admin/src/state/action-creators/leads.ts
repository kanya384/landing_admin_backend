import { Dispatch } from "react";
import { Configuration, DayLeadsInfo, LeadsApi } from "../../api";
import { GetTokenFromCookies } from "../../utils";
import { LeadsActionTypes } from "../action-types";
import { LeadsAction } from "../actions";

const LEADS_ERROR = "Ошибка при получении данных"
const leadsService = new LeadsApi(new Configuration(), process.env.REACT_APP_BACKEND_URL)

export const getLeads = () => {
  return async (dispatch: Dispatch<LeadsAction>) => {
    dispatch({
      type: LeadsActionTypes.LEADS_REQUEST_SEND,
    })
    let token = GetTokenFromCookies()
    leadsService.leadGet({headers: {"Authorization": token}}).then((resp) => {
      if (resp.status === 200) {
        dispatch({
          type: LeadsActionTypes.LEADS_SUCCESS,
          payload: resp.data,
        })
      } else {
        dispatch({
          type: LeadsActionTypes.LEADS_ERROR,
          payload: LEADS_ERROR,
        })
      }
    }).catch((e) => {
      dispatch({
        type: LeadsActionTypes.LEADS_ERROR,
        payload: LEADS_ERROR,
      })
    })
  }
}

export const getLeadsAnalytics = () => {
  return async (dispatch: Dispatch<LeadsAction>) => {
    let token = GetTokenFromCookies()
    leadsService.leadAnalyticsGet({headers: {"Authorization": token}}).then((resp) => {
      if (resp.status === 200) {
        let newChartInfo = <DayLeadsInfo[]>[]
        resp.data.chart_info?.map((info) => {
          newChartInfo.push({
            count: info.count?info.count:0,
            date: info.date?.split('.').reverse().join('.')
          })
        })
        resp.data.chart_info = newChartInfo
        dispatch({
          type: LeadsActionTypes.LEADS_ANALYTICS_SUCCESS,
          payload: resp.data,
        })
      } 
    }).catch((e) => {
      
    })
  }
}

export const deleteLead = (id: string) => {
  return async (dispatch: Dispatch<LeadsAction>) => {
      let token = GetTokenFromCookies()
      leadsService.leadIdDelete(id, {headers: {"Authorization": token}}).then((resp)=> {
          if (resp.status === 200) {
              dispatch({
                  type: LeadsActionTypes.LEADS_DELETE,
                  payload: id
              })
          }
      })
  }
}