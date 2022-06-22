import React, { useCallback, useEffect, useState } from "react";
import { DndProvider } from "react-dnd";
import { HTML5Backend } from "react-dnd-html5-backend";
import { Trash } from "react-feather";
import Card from "../../components/card";
import { Popups } from "../../components/popups/popups";
import { useActions } from "../../hooks/use-actions";
import { useTypedSelector } from "../../hooks/use-typed-selector";
import { excludeExistingMonths, months } from "./months";

interface Active {
  year: null | number,
  month: null | number,
}
export const Hod: React.FC = () =>{
  const [yearForm, setYearForm] = useState({
    value: new Date().getFullYear(),
  })
  const [monthForm, setMonthForm] = useState({
    value: 3,
  })
  const [photosForm, setPhotosForm] = useState<{photos:any[]}>({
    photos: [],
  })
  const [error, setFormError] = useState("")
  const [opened, setOpened] = useState<number|null>(null);
  const [active, setActive] = useState<Active>({
    year: null,
    month: null,
  })
  const hod = useTypedSelector(({ hod }) => {
    return hod
  })
  const { getYears, getMonths, getPhotos, addYear, deleteYear, addMonth, deleteMonth, addPhotos, sortPhotos, deletePhotoHod } = useActions();
  const getMonthID = (monthIn: number) => {
    let monthRes: string = ""
    hod?.monthsList.map((month) => {
      if (month.value === monthIn) {
        monthRes = month.id!
      }
    })
    return monthRes
  }
  const getMonthName = (monthIn: number) => {
    let monthRes: string = ""
    hod?.monthsList.map((month) => {
      if (month.value === monthIn) {
        monthRes = month.name!
      }
    })
    return monthRes
  }

  const moveCard = useCallback(async (dragIndex: number, hoverIndex: number) => {
    if (hod!.photosList.length>0){
      await sortPhotos([...hod!.photosList],dragIndex, hoverIndex)
    }
  }, [sortPhotos, hod])

  const getYearID = (monthIn: number) => {
    let yearRes: string = ""
    hod?.yearsList.map((year) => {
      if (year.value === monthIn) {
        yearRes = year.id!
      }
    })
    return yearRes
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
      getPhotos(getMonthID(hod!.monthsList[0].value!))
    }

  },[hod, active, getYears, getMonths])

  const requestCallback = (error: string) => {
    if (error!==""){
      setFormError(error)
    } else {
      setOpened(null);
    }
  }
  const popups = [
    {
      title: "Добавить год",
      stateAction: setYearForm,
      sendAction: () =>{ addYear({value: yearForm.value}, requestCallback);},
      error: error,
      fields: [
        {
          title: "Год",
          type: "number", 
          placeholder: "Введите год", 
          name: "value", 
          isError: false, 
          value:yearForm.value,
        },
      ],
    },
    {
      title: `Добавить месяц (${active.year})`,
      stateAction: setMonthForm,
      sendAction: () =>{ addMonth({value: monthForm.value, year_id: getYearID(active.year!)}, requestCallback);},
      error: error,
      fields: [
        {
          title: "Месяц",
          type: "select", 
          placeholder: "Введите месяц", 
          name: "value", 
          isError: false, 
          value:monthForm.value,
          fields: excludeExistingMonths(months, hod?.monthsList!),
        },
      ],
    },
    {
      title: `Добавить фото (${getMonthName(active.month!)} ${active.year})`,
      stateAction: setPhotosForm,
      sendAction: () =>{ addPhotos(photosForm.photos, getMonthID(active.month!), requestCallback);},
      error: error,
      fields: [
        {
          title: "Перетащите фото",
          type: "drop", 
          placeholder: "", 
          name: "photos", 
          isError: false, 
          value:photosForm.photos,
        },
      ],
    },
  ]

  const yearClick = async (year: number) =>{
    await getMonths(getYearID(year))
    setActive((prevValue) => {
      return {
        year: year,
        month: hod?hod!.monthsList.length>0?hod!.monthsList[0].value!:null:null
      }
    })
  }
  const monthClick = (month: number) =>{
    setActive({
      year: active.year,
      month: month
    })
    getPhotos(getMonthID(month))
  }
  return(
    <React.Fragment>
      <div className="card-header p-4 border-300 bg-soft">
        <div className="row g-3 justify-content-between align-items-center">
          <div className="col-12 col-md">
            <h1>Ход строительства</h1>
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
                return <div className="btn-group p-2" role="group" key={year.id!} aria-label="Basic example">
                  <button className={active.year === year.value?"btn btn-primary":"btn tn-soft-primary"} onClick={()=>{yearClick(year.value!)}} type="button">{year.value}</button>
                  <button className={active.year === year.value?"btn btn btn-primary":"btn tn-soft-primary"} style={{padding: 2}}><Trash size={15} onClick={()=>{deleteYear(year.id!)}} /></button>
                </div>
              })}
          </div>
        </div>
        <div className="">
          <button className="btn btn-primary" onClick={() => {setOpened(0)}}>Добавить</button>
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
                <button className={active.month === month.value?"btn btn btn-primary":"btn tn-soft-primary"} onClick={()=>{monthClick(month.value!)}} type="button">{month.name}</button>
                <button className={active.month === month.value?"btn btn btn-primary":"btn tn-soft-primary"} style={{padding: 2}}><Trash size={15} onClick={()=>{deleteMonth(month.id!)}} /></button>
              </div>
              })}
            </div>
          </div>
          <div>
            <button className="btn btn-primary" onClick={() => {setOpened(1)}}>Добавить</button>
          </div>
      </div></ React.Fragment>:<div></div>}
      {hod && hod?.monthsList.length>0?<React.Fragment>
        <div style={{display: "flex", justifyContent: "space-between", padding: "1.5rem"}}>
            <h4>Фото</h4>
            <button className="btn btn-primary" onClick={() => {setPhotosForm({photos: []});setOpened(2)}}>Добавить</button>
        </div>
        <div className='p-4'>
          <DndProvider backend={HTML5Backend}>
            <div className='row g-4'>
              {hod.photosList?.map((photo, index)=>{
                return <Card key={photo.id} card={{ID: photo.id!, Type:1, Title: "", Text: "", Image: "/api/store/"+photo.image!, Index: index, moveCard: moveCard, deleteClick: () => {deletePhotoHod(photo.id!)}, editClick: ()=>{}}} />
              })}
            </div>
          </DndProvider>
        </div>
        </React.Fragment>:<div></div>}
      <Popups Forms={popups} Opened={opened} Toggle={()=>{setOpened(null)}} />
    </React.Fragment>
  )
}