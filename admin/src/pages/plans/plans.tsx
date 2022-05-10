import React, { useEffect, useState } from "react"
import { DndProvider } from "react-dnd"
import { HTML5Backend } from "react-dnd-html5-backend"
import { Popups } from "../../components/popups/popups"
import { useActions } from "../../hooks/use-actions"
import { useTypedSelector } from "../../hooks/use-typed-selector"
import { PlanCard } from "./plan-card"

export const Plans: React.FC = () => {
  const plans = useTypedSelector(({ plans }) => {
    return plans
  })

  const [opened, setOpened] = useState<number|null>(null);

  const [form, setForm] = useState({
    id: "",
    file: null,
  })

  const [error, setError] = useState("")

  const { getPlans, updatePlanPhoto } = useActions()

  useEffect(()=>{
    getPlans()
  }, [])

  const updatePlansCallback = (error: string) => {
    if (error!==""){
      setError(error)
    } else {
      setOpened(null);
    }
  }
  const popups = [
    {
      title: `Редактировать фото планировки`,
      stateAction: setForm,
      sendAction: () =>{updatePlanPhoto(form.id, form.file, updatePlansCallback)},
      error: error,
      fields: [
        {
          title: "Планировка",
          type: "file", 
          placeholder: "Файл не выбран", 
          name: "file", 
          isError: false, 
          value:"",
        },
      ],
    }, 
  ]

  const editPhotoPopupShow = (id: string) => {
    setForm({
      id: id,
      file: null,
    })
    setOpened(0)
  }
  return (
    <React.Fragment>
      <div className="card-header p-4 border-300 bg-soft">
        <div className="row g-3 justify-content-between align-items-center">
          <div className="col-12 col-md">
            <h1>Планировки</h1>
          </div>
        </div>
      </div>
      <div className='p-4'>
          <DndProvider backend={HTML5Backend}>
            <div className='row g-4'>
              {plans?.plansList?.map((plan)=>{
                return <PlanCard key={plan.id} plan={plan} btnClick={editPhotoPopupShow}/>
              })}
            </div>
          </DndProvider>
        </div>
        <Popups Forms={popups} Opened={opened} Toggle={()=>{setOpened(null)}} />
    </React.Fragment>
  )
}