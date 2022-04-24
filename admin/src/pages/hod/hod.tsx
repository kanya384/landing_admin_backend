import React, { useEffect, useState } from "react";
import { DndProvider } from "react-dnd";
import { HTML5Backend } from "react-dnd-html5-backend";
import { Trash } from "react-feather";
import { useActions } from "../../hooks/use-actions";
import { useTypedSelector } from "../../hooks/use-typed-selector";

interface Active {
  year: null | number,
  month: null | number,
}
export const Hod: React.FC = () =>{
  const [active, setActive] = useState<Active>({
    year: null,
    month: null,
  })
  const hod = useTypedSelector(({ hod }) => {
    return hod
  })
  const { getYears, getMonths, getPhotos } = useActions();
  const getMonthID = (monthIn: number) => {
    let monthRes: string = ""
    hod?.monthsList.map((month) => {
      if (month.value === monthIn) {
        monthRes = month.id!
      }
    })
    return monthRes
  }
  useEffect(()=>{
    if (hod?.yearsList.length === 0){
      getYears()
    } else if (active.year === null) {
      setActive({
        year: hod!.yearsList[0].value!,
        month: active.month,
      })
      getMonths(hod!.yearsList[0].id!)
    } else if (active.month === null && hod!.monthsList.length > 0) {
      setActive({
        year: hod!.yearsList[0].value!,
        month: hod!.monthsList[0].value!,
      })
    } else if (active.month !== null && hod!.photosList.length === 0) {
      getPhotos(getMonthID(active.month))
    }

  },[hod, active, getYears, getMonths])
  return(
    <React.Fragment>
      <div className="card-header p-4 border-300 bg-soft">
        <div className="row g-3 justify-content-between align-items-center">
          <div className="col-12 col-md">
            <h1>Ход стороительства</h1>
          </div>
        </div>
      </div>
      <div style={{display: "flex", justifyContent: "space-between", padding: "1.5rem", borderBottom: "0.2px solid grey"}}>
        <div className="">
            <div className="mb-3">
              <h4>Год</h4>
            </div>
            <div>
              {hod?.yearsList.map((year)=>{
                return <div className="btn-group p-2" role="group" aria-label="Basic example">
                  <button className={active.year === year.value?"btn btn-primary":"btn btn-phoenix-primary"} type="button">{year.value}</button>
                  <button className="btn btn-soft-danger" style={{padding: 10}}><Trash size={15}  /></button>
                </div>
              })}
          </div>
        </div>
        <div className="">
          <button className="btn btn-primary" onClick={() => {}}>Добавить</button>
        </div>
      </div>
      {hod && hod?.yearsList.length>0?<React.Fragment>
        <div style={{display: "flex", justifyContent: "space-between", padding: "1.5rem", borderBottom: "0.2px solid grey"}}>
          <div className="">
              <div className="mb-3">
                <h4>Месяц</h4>
              </div>
              <div>
              {hod?.monthsList.map((month)=>{
                return <div className="btn-group p-2" role="group" aria-label="Basic example">
                <button className={active.month === month.value?"btn btn btn-primary":"btn btn-phoenix-primary"} type="button">{month.name}</button>
                <button className="btn btn-soft-danger" style={{padding: 10}}><Trash size={15}  /></button>
              </div>
              })}
            </div>
          </div>
          <div>
            <button className="btn btn-primary" onClick={() => {}}>Добавить</button>
          </div>
      </div></ React.Fragment>:<div></div>}
      {hod && hod?.monthsList.length>0?<React.Fragment>
        <div style={{display: "flex", justifyContent: "space-between", padding: "1.5rem"}}>
            <h4>Фото</h4>
            <button className="btn btn-primary" onClick={() => {}}>Добавить</button>
        </div>
        <div className='p-4'>
          <DndProvider backend={HTML5Backend}>
            <div className='row g-4'>
              
            </div>
          </DndProvider>
        </div>
        </React.Fragment>:<div></div>}
    </React.Fragment>
  )
}