import React, { useEffect, useState } from "react";
import { Popups } from "../../components/popups/popups";
import { useActions } from "../../hooks/use-actions";
import { useTypedSelector } from "../../hooks/use-typed-selector";
import { TableItem } from "./table-item";

export const Settings: React.FC = () => {
  const [opened, setOpened] = useState<number|null>(null);
  const [addForm, setAddFrom] = useState<{
    id: string,
    name: string,
    description: string,
    value: number,
  }>({
    id: "",
    name: "",
    description: "",
    value: 0,
  })
  const [updateForm, setUpdateForm] = useState<{
    id: string,
    name: string,
    description: string,
    value: number,
  }>({
    id: "",
    name: "",
    description: "",
    value: 0,
  })
  const [error, setError] = useState("")
  const titlesCallback = (error: string) => {
    if (error!==""){
      setError(error)
    } else {
      setOpened(null);
    }
  }
  const { getSettings, addUpdateSetting} = useActions()
  const popups = [
    {
      title: "Добавить параметр",
      stateAction: setAddFrom,
      sendAction: () =>{addUpdateSetting(addForm, titlesCallback)},
      error: error,
      fields: [
        {
            title: "Описание параметра",
            type: "text", 
            placeholder: "Введите описание", 
            name: "description", 
            isError: false, 
            value:addForm.description,
        },
        {
          title: "Имя параметра (анг)",
          type: "text", 
          placeholder: "Введите имя параметра", 
          name: "name", 
          isError: false, 
          value:addForm.name,
        },
        {
          title: "Значение параметра",
          type: "select", 
          placeholder: "Выберите значение параметра", 
          name: "value", 
          isError: false,
          value:addForm!.value,
          fields:[
            {
              name: "Да",
              value: 1, 
              type: 0,
            },
            {
              name: "Нет",
              value: 0,
              type: 0,
            },
          ],
        },
      ],
    },
    {
        title: "Редактировать параметр",
        stateAction: setAddFrom,
        sendAction: () =>{addUpdateSetting(updateForm, titlesCallback)},
        error: error,
        fields: [
          {
            title: "Заголовок для моб. версии",
            type: "text", 
            placeholder: "Введите заголовок", 
            name: "name", 
            isError: false, 
            value:addForm.name,
          },
          {
              title: "Заголовок для моб. версии",
              type: "text", 
              placeholder: "Введите заголовок", 
              name: "value", 
              isError: false, 
              value:addForm.value,
          },
        ],
      },
  ]
  const settings = useTypedSelector(({ settings }) => {
    return settings
  })

  
  useEffect(()=>{
    getSettings()
  }, [])
  return (
    <React.Fragment>
      <div className="card-header p-4 border-300 bg-soft">
        <div className="row g-3 justify-content-between align-items-center">
          <div className="col-12 col-md">
            <h1>Настройки</h1>
          </div>
          <div className="col col-md-auto">
              <nav className="nav nav-underline justify-content-end border-0 doc-tab-nav align-items-center" role="tablist">
              <button className="btn btn-primary" onClick={() => {setOpened(0)}}>Добавить</button>
              </nav>
            </div>
        </div>
      </div>
      <table className="table">
        <thead>
          <tr>
            <th scope="col">#</th>
            <th scope="col">Описание</th>
            <th scope="col">Параметр</th>
            <th scope="col">Значение</th>
          </tr>
        </thead>
        <tbody>
          {settings?.settingsList.map((setting, index) => {
            return <TableItem setting={setting} index={index+1} editAction={()=>{setUpdateForm({id: setting.id!, name: setting.name!, description:setting.description!, value: setting.value!}); setOpened(1)}} />
          })} 
        </tbody>
      </table>
      <Popups Forms={popups} Opened={opened} Toggle={()=>{setOpened(null)}} />
    </React.Fragment>
  )
}