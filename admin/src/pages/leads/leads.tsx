import React, { useEffect } from "react"
import { Area, AreaChart, CartesianGrid, Line, LineChart, ResponsiveContainer, Tooltip, XAxis, YAxis } from "recharts"
import { useActions } from "../../hooks/use-actions";
import { useTypedSelector } from "../../hooks/use-typed-selector";
import { TableItem } from "./table-item";

export const Leads: React.FC = () => {
  const { getLeads, getLeadsAnalytics, deleteLead } = useActions()
  const leads = useTypedSelector(({ leads }) => {
    return leads
  })

  useEffect(()=>{
    if (leads?.leadsList.length == 0) {
      getLeads()
      getLeadsAnalytics()
    }
  }, [])
  return (
    <React.Fragment>
      <div className="card-header p-4 border-300 bg-soft">
        <div className="row g-3 justify-content-between align-items-center">
          <div className="col-12 col-md">
            <h1>Панель управления</h1>
          </div>
          <div className="col col-md-auto">
            <nav className="nav nav-underline justify-content-end border-0 doc-tab-nav align-items-center" role="tablist">
            <a className="btn btn-primary" target="_blank" href={"http://localhost:3000/?administarate=y"} style={{width:"100%"}}>Режим редактирования</a>
            </nav>
          </div>
        </div>
      </div>
      {leads?.analytics?
      <div className="card-body">
        <div className="row align-items-center g-4">
          <div className="col-12 col-md-auto">
            <div className="d-flex align-items-center"><img src="assets/img/icons/illustrations/4.png" alt="" width="46" height="46" />
              <div className="ms-3">
                <h4 className="mb-0">{leads.analytics.today_count?leads.analytics.today_count:0} шт</h4>
                <p className="text-800 fs--1 mb-0">Количество заявок за сегодня</p>
              </div>
            </div>
          </div>
          <div className="col-12 col-md-auto">
            <div className="d-flex align-items-center"><img src="assets/img/icons/illustrations/2.png" alt="" width="46" height="46" />
              <div className="ms-3">
                <h4 className="mb-0">{leads.analytics.month_count?leads.analytics.month_count:0} шт</h4>
                <p className="text-800 fs--1 mb-0">Количество заявок в этом месяце</p>
              </div>
            </div>
          </div>
        </div>
        <hr className="bg-200 mb-6 mt-4"></hr>
        <div className="row flex-between-center mb-4 g-3">
          <div className="col-auto">
            <h3>График</h3>
            <p className="text-700 lh-sm mb-0">Распределение заявок в этом месяце по дням</p>
          </div>
        </div>
        <ResponsiveContainer width='100%' aspect={4.0/1}>
          <AreaChart data={leads.analytics.chart_info}>
            <CartesianGrid strokeDasharray="3 3" />
            <Area type="monotone" dataKey="count" stroke="#82ca9d" fill="#82ca9d" />
            <XAxis dataKey="date" />
            <Tooltip />
          </AreaChart>
        </ResponsiveContainer>
      </div>:""}
      <div className="card-body">
      <hr className="bg-200 mb-6 "></hr>
        <div className="row flex-between-center mb-4 g-3">
          <div className="col-auto">
            <h3>Заявки</h3>
            <p className="text-700 lh-sm mb-0">Список заявок (последние 50 шт.)</p>
          </div>
        </div>
        <table className="table">
        <thead>
          <tr>
            <th scope="col">#</th>
            <th scope="col">Дата</th>
            <th scope="col">Номер телефона</th>
            <th scope="col">Имя</th>
            <th scope="col">Email</th>
            <th scope="col">Текст</th>
            <th scope="col">Действия</th>
          </tr>
        </thead>
        <tbody>
           {leads?.leadsList.map((lead, index) => {
             return <TableItem lead={lead} index={index+1} deleteAction={() => deleteLead(lead.id!)}/>
           })}
        </tbody>
      </table>
      </div>
    </React.Fragment>
  )
}