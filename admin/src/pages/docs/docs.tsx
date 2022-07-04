import React, { useEffect, useState } from "react";
import { Popups } from "../../components/popups/popups";
import { useActions } from "../../hooks/use-actions";
import { useTypedSelector } from "../../hooks/use-typed-selector";
import { TableItem } from "./table-item";

export const Doc: React.FC = () => {
  const [opened, setOpened] = useState<number|null>(null);
  const [addForm, setAddFrom] = useState({
    title: "",
    file: null,
  })
  const [updateForm, setUpdateForm] = useState<{
    id: string,
    title: string,
    file: null | any
  }>({
    id: "",
    title: "",
    file: null,
  })
  const [error, setError] = useState("")
  const docsCallback = (error: string) => {
    if (error!==""){
      setError(error)
    } else {
      setOpened(null);
    }
  }
  const { getDocs, addNewDoc, editDoc, deleteDoc } = useActions()
  const popups = [
    {
      title: "Добавить документ",
      stateAction: setAddFrom,
      sendAction: () =>{addNewDoc(addForm.file, addForm.title, docsCallback)},
      error: error,
      fields: [
        {
          title: "Файл",
          type: "file", 
          placeholder: "Файл не выбран", 
          name: "file", 
          isError: false, 
          value:"",
        },
        {
          title: "Тайтл",
          type: "text", 
          placeholder: "Введите заголовок", 
          name: "title", 
          isError: false, 
          value:addForm.title,
        },
      ],
    },
    {
      title: "Редактировать документ",
      stateAction: setUpdateForm,
      sendAction: ()=>{editDoc(updateForm.file, updateForm.id, updateForm.title, docsCallback)},
      error: error,
      fields: [
        {
          title: "Обложка",
          type: "file", 
          placeholder: "Файл не выбран", 
          name: "file", 
          isError: false, 
          value:"",
        },
        {
          title: "Тайтл",
          type: "text", 
          placeholder: "Введите заголовок", 
          name: "title", 
          isError: false, 
          value:updateForm!.title,
        },
      ],
    }
  ]
  const docs = useTypedSelector(({ docs }) => {
    return docs
  })

  
  useEffect(()=>{
    getDocs()
  }, [])
  return (
    <React.Fragment>
      <div className="card-header p-4 border-300 bg-soft">
        <div className="row g-3 justify-content-between align-items-center">
          <div className="col-12 col-md">
            <h1>Документы</h1>
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
            <th scope="col">Название</th>
            <th scope="col">Действие</th>
          </tr>
        </thead>
        <tbody>
          {docs?.docsList.map((doc, index) => {
            return <TableItem doc={doc} index={index+1} editAction={()=>{setUpdateForm({id: doc.id!, title: doc.title!, file: doc.file!}); setOpened(1)}} deleteAction={()=>{deleteDoc(doc.id!)}} />
          })} 
        </tbody>
      </table>
      <Popups Forms={popups} Opened={opened} Toggle={()=>{setOpened(null)}} />
    </React.Fragment>
  )
}