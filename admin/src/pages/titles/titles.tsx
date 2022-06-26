import React, { useEffect, useState } from "react";
import { Popups } from "../../components/popups/popups";
import { useActions } from "../../hooks/use-actions";
import { useTypedSelector } from "../../hooks/use-typed-selector";
import { TableItem } from "./table-item";

export const Titles: React.FC = () => {
  const [opened, setOpened] = useState<number|null>(null);
  const [addForm, setAddFrom] = useState({
    id: "",
    tagName: "",
    tagValue: "",
    desktopTitle: "",
    mobileTitle: "",
  })
  const [updateForm, setUpdateForm] = useState<{
    id: string,
    tagName: string,
    tagValue: string,
    desktopTitle: string,
    mobileTitle: string,
  }>({
    id: "",
    tagName: "",
    tagValue: "",
    desktopTitle: "",
    mobileTitle: "",
  })
  const [error, setError] = useState("")
  const titlesCallback = (error: string) => {
    if (error!==""){
      setError(error)
    } else {
      setOpened(null);
    }
  }
  const { getTitles, addNewTitle, updateTitle, deleteTitles } = useActions()
  const popups = [
    {
      title: "Добавить тайтл",
      stateAction: setAddFrom,
      sendAction: () =>{addNewTitle(addForm, titlesCallback)},
      error: error,
      fields: [
        {
          title: "Заголовок для десктоп версии",
          type: "text", 
          placeholder: "Введите заголовок", 
          name: "desktopTitle", 
          isError: false, 
          value:addForm.desktopTitle,
        },
        {
          title: "Заголовок для моб. версии",
          type: "text", 
          placeholder: "Введите заголовок", 
          name: "mobileTitle", 
          isError: false, 
          value:addForm.mobileTitle,
        },
        {
          title: "Тег",
          type: "select", 
          placeholder: "Выберите метку подмены", 
          name: "tagName", 
          isError: false,
          value:addForm!.tagName,
          fields:[
            {
              name: "utm_medium",
              value: "utm_medium", 
              type: 1,
            },
            {
              name: "utm_content",
              value: "utm_content",
              type: 1,
            },
            {
              name: "utm_term",
              value: "utm_term",
              type: 1,
            },
            {
              name: "utm_campaign",
              value: "utm_campaign",
              type: 1,
            },
            {
              name: "utm_source",
              value: "utm_source",
              type: 1,
            }
          ],
        },
        {
          title: "Значение подмены",
          type: "text", 
          placeholder: "Введите значение подмены", 
          name: "tagValue", 
          isError: false, 
          value:addForm!.tagValue,
        },
      ],
    },
    {
      title: "Редактировать документ",
      stateAction: setUpdateForm,
      sendAction: ()=>{updateTitle(updateForm, titlesCallback)},
      error: error,
      fields: [
        {
          title: "Заголовок для десктоп версии",
          type: "text", 
          placeholder: "Введите заголовок", 
          name: "desktopTitle", 
          isError: false, 
          value:updateForm.desktopTitle,
        },
        {
          title: "Заголовок для моб. версии",
          type: "text", 
          placeholder: "Введите заголовок", 
          name: "mobileTitle", 
          isError: false, 
          value:updateForm.mobileTitle,
        },
        {
          title: "Тег",
          type: "select", 
          placeholder: "Выберите метку подмены", 
          name: "tagName", 
          isError: false,
          value:updateForm!.tagName,
          fields:[
            {
              name: "utm_medium",
              value: "utm_medium", 
              type: 1,
            },
            {
              name: "utm_content",
              value: "utm_content",
              type: 1,
            },
            {
              name: "utm_term",
              value: "utm_term",
              type: 1,
            },
            {
              name: "utm_campaign",
              value: "utm_campaign",
              type: 1,
            },
            {
              name: "utm_source",
              value: "utm_source",
              type: 1,
            }
          ],
        },
        {
          title: "Значение подмены",
          type: "text", 
          placeholder: "Введите значение подмены", 
          name: "tagValue", 
          isError: false, 
          value:updateForm!.tagValue,
        },
      ],
    }
  ]
  const titles = useTypedSelector(({ titles }) => {
    return titles
  })

  
  useEffect(()=>{
    getTitles()
  }, [])
  return (
    <React.Fragment>
      <div className="card-header p-4 border-300 bg-soft">
        <div className="row g-3 justify-content-between align-items-center">
          <div className="col-12 col-md">
            <h1>Подмены</h1>
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
            <th scope="col">Подмена для десктоп версии</th>
            <th scope="col">Подмена для моб. версии</th>
            <th scope="col">Метка подмены</th>
            <th scope="col">Значение метки</th>
            <th scope="col">Действия</th>
          </tr>
        </thead>
        <tbody>
          {titles?.titleList.map((title, index) => {
            return <TableItem title={title} index={index+1} editAction={()=>{setUpdateForm({id: title.id!, tagName: title.tagName!, tagValue: title.tagValue!, desktopTitle: title.desktopTitle!, mobileTitle: title.mobileTitle!}); setOpened(1)}} deleteAction={()=>{deleteTitles(title.id!)}} />
          })} 
        </tbody>
      </table>
      <Popups Forms={popups} Opened={opened} Toggle={()=>{setOpened(null)}} />
    </React.Fragment>
  )
}