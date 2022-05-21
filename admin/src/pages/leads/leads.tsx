import React, { useEffect } from "react"
import { Line, LineChart, ResponsiveContainer, Tooltip, XAxis, YAxis } from "recharts"
import { useActions } from "../../hooks/use-actions";
import { useTypedSelector } from "../../hooks/use-typed-selector";
import { TableItem } from "./table-item";

export const Leads: React.FC = () => {
  const data = [
    {date: '21.05', лиды: 5, pv: 2400, amt: 2400}, 
    {date: '22.05', лиды: 10, pv: 2400, amt: 2400},
    {date: '23.05', лиды: 3, pv: 2400, amt: 2400},
    {date: '24.05', лиды: 10, pv: 2400, amt: 2400},
    {date: '25.05', лиды: 3, pv: 2400, amt: 2400},
    {date: '20.05', лиды: 20, pv: 2400, amt: 2400},
  ];
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
        </div>
      </div>
      {leads?.analytics?
      <div className="card-body">
        <div className="row align-items-center g-4">
          <div className="col-12 col-md-auto">
            <div className="d-flex align-items-center"><img src="assets/img/icons/illustrations/4.png" alt="" width="46" height="46" />
              <div className="ms-3">
                <h4 className="mb-0">{leads.analytics.today_count} шт</h4>
                <p className="text-800 fs--1 mb-0">Количество заявок за сегодня</p>
              </div>
            </div>
          </div>
          <div className="col-12 col-md-auto">
            <div className="d-flex align-items-center"><img src="assets/img/icons/illustrations/2.png" alt="" width="46" height="46" />
              <div className="ms-3">
                <h4 className="mb-0">{leads.analytics.month_count} шт</h4>
                <p className="text-800 fs--1 mb-0">Количество заявок в этом месяце</p>
              </div>
            </div>
          </div>
        </div>
        <hr className="bg-200 mb-6 mt-4"></hr>
        <div className="row flex-between-center mb-4 g-3">
          <div className="col-auto">
            <h3>График</h3>
            <p className="text-700 lh-sm mb-0">Распределение заявок по дням</p>
          </div>
        </div>
        <ResponsiveContainer width='100%' aspect={4.0/1}>
          <LineChart data={leads.analytics.chart_info}>
            <Line type="monotone" dataKey="count" stroke="#8884d8" />
            <XAxis dataKey="date" />
            <Tooltip />
          </LineChart>
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
             return <TableItem lead={lead} index={index+1} deleteAction={deleteLead}/>
           })}
        </tbody>
      </table>
      </div>
    </React.Fragment>
  )
}