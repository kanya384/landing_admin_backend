import React, { useEffect, useState } from "react"
import { DndProvider } from "react-dnd"
import { HTML5Backend } from "react-dnd-html5-backend"
import { useNavigate } from "react-router-dom"
import Card from "../../components/card"
import { Popups } from "../../components/popups/popups"
import { useActions } from "../../hooks/use-actions"
import { useTypedSelector } from "../../hooks/use-typed-selector"

export const Advantages: React.FC = () => {
  const [opened, setOpened] = useState<number|null>(null);
  const navigate = useNavigate();
  const [addForm, setAddFrom] = useState({
    title: "",
    description: ""
  })
  const [updateForm, setUpdateFrom] = useState({
    id: "",
    title: "",
    description: "",
    order: 0,
    updatedAt: "",

  })
  const [addFormError, setAddFormError] = useState("")
  const [updateFormError, setUpdateFormError] = useState("")
  const advantages = useTypedSelector(({ advantages }) => {
    return advantages
  })
  const { getAdvantages, addAdvantage, updateAdvantage, deleteAdvantage } = useActions();
  useEffect(()=>{
    getAdvantages()
  }, [])
  const sendFormallback = (error: string) => {
    if (error!==""){
      setAddFormError(error)
    } else {
      setOpened(null);
    }
  }
  const popups = [
    {
      title: "Добавить преимущество",
      stateAction: setAddFrom,
      sendAction: () =>{addAdvantage(addForm.title, addForm.description, sendFormallback)},
      error: addFormError,
      fields: [
        {
          title: "Тайтл",
          type: "text", 
          placeholder: "Введите заголовок", 
          name: "title", 
          isError: false, 
          value:addForm.title,
        },
        {
          title: "Описание",
          type: "textarea", 
          placeholder: "Введите описание", 
          name: "description", 
          isError: false, 
          value:addForm.description,
        },
      ],
    },
    {
      title: "Редактировать преимущество",
      stateAction: setUpdateFrom,
      sendAction: () =>{updateAdvantage(updateForm, sendFormallback)},
      error: updateFormError,
      fields: [
        {
          title: "Тайтл",
          type: "text", 
          placeholder: "Введите заголовок", 
          name: "title", 
          isError: false, 
          value:updateForm.title,
        },
        {
          title: "Описание",
          type: "textarea", 
          placeholder: "Введите описание", 
          name: "description", 
          isError: false, 
          value:updateForm.description,
        },
      ],
    },
  ]
  return(
    <React.Fragment>
      <div className="card-header p-4 border-300 bg-soft">
        <div className="row g-3 justify-content-between align-items-center">
          <div className="col-12 col-md">
            <h1>А так же вас ждет</h1>
          </div>
          <div className="col col-md-auto">
            <nav className="nav nav-underline justify-content-end border-0 doc-tab-nav align-items-center" role="tablist">
            <button className="btn btn-primary" onClick={() => {setAddFormError(""); setOpened(0)}}>Добавить</button>
            </nav>
          </div>
        </div>
      </div>
      <DndProvider backend={HTML5Backend}>
        {advantages?.advantagesList.map((advantage, index)=>{
          return <Card key={advantage.id} card={{ID: advantage.id!, Type:2, Title: advantage.title!, Text: advantage.description!, Image: "/api/store/", date: advantage.updatedAt, Index: index, click:()=>{ navigate(`${advantage.id}`, { replace: false })}, moveCard: ()=>{}, deleteClick: () => {deleteAdvantage(advantage.id!)}, editClick:()=>{setUpdateFrom({id: advantage.id!, title: advantage.title!, description: advantage.description!, order: !advantage.order?0:advantage.order!, updatedAt: advantage.updatedAt!}); setUpdateFormError(""); setOpened(1)}, }}/>
        })}
      </ DndProvider>
      <Popups Forms={popups} Opened={opened} Toggle={()=>{setOpened(null)}} />
    </React.Fragment>
  )
}