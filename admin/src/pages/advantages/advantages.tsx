import React, { useEffect, useState } from "react"
import { DndProvider } from "react-dnd"
import { HTML5Backend } from "react-dnd-html5-backend"
import Card from "../../components/card"
import { Popups } from "../../components/popups/popups"
import { useActions } from "../../hooks/use-actions"
import { useTypedSelector } from "../../hooks/use-typed-selector"

export const Advantages: React.FC = () => {
  const [opened, setOpened] = useState<number|null>(null);
  const [addForm, setAddFrom] = useState({
    title: "",
    description: ""
  })
  const [addFormError, setAddFormError] = useState("")
  const advantages = useTypedSelector(({ advantages }) => {
    return advantages
  })
  const { getAdvantages, addAdvantage } = useActions();
  useEffect(()=>{
    getAdvantages()
  }, [])
  const addAdvantageCallback = (error: string) => {
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
      sendAction: () =>{addAdvantage(addForm.title, addForm.description, addAdvantageCallback)},
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
                return <Card key={advantage.id} card={{ID: advantage.id!, Type:2, Title: advantage.title!, Text: advantage.description!, Image: "http://localhost:8080/store/", date: advantage.updatedAt, Index: index, moveCard: ()=>{}, deleteClick: () => {}, editClick:()=>{}}}/>
              })}
        </ DndProvider>

      <Popups Forms={popups} Opened={opened} Toggle={()=>{setOpened(null)}} />
    </React.Fragment>
  )
}