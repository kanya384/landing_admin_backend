import React, { useCallback, useEffect, useState } from 'react';
import Card from '../../components/card';
import Spinner from '../../components/spinner/spinner';
import { useActions } from '../../hooks/use-actions';
import { useTypedSelector } from '../../hooks/use-typed-selector';
import {DndProvider} from 'react-dnd'
import { HTML5Backend } from 'react-dnd-html5-backend';
import { Popups } from '../../components/popups/popups';


export const Actions: React.FC = () =>  {
    const [opened, setOpened] = useState<number|null>(null);
    const [addForm, setAddFrom] = useState({
      title: "",
      description: "",
      date: "",
      file: null,
    })
    const [addFormError, setAddFormError] = useState("")
    const [updateForm, setUpdateForm] = useState({
      id: "",
      title: "",
      description: "",
      date: "",
      photo: "",
      preview: "",
      order: 0,
      file: null,
    })
    const [updateFormError, setUpdateFormError] = useState("")
    const addActionCallback = (error: string) => {
      if (error!==""){
        setAddFormError(error)
      } else {
        setOpened(null);
      }
    }
    const updateActionCallback = (error: string) => {
      if (error!==""){
        setUpdateFormError(error)
      } else {
        setOpened(null);
      }
    }
    const popups = [
      {
        title: "Добавить акцию",
        stateAction: setAddFrom,
        sendAction: () =>{ addNewAction(addForm.file, addForm.title, addForm.description, addForm.date, addActionCallback);},
        error: addFormError,
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
            title: "Дата",
            type: "text", 
            placeholder: "Введите до какого числа действует акция", 
            name: "date", 
            isError: false, 
            value:addForm.date,
          },
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
        title: "Редактировать акцию",
        stateAction: setUpdateForm,
        sendAction: ()=>updateAction(updateForm, updateActionCallback, updateForm.file),
        error: updateFormError,
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
            title: "Дата",
            type: "text", 
            placeholder: "Введите до какого числа действует акция", 
            name: "date", 
            isError: false, 
            value:updateForm.date,
          },
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
          {
            title: "Индекс",
            type: "number", 
            placeholder: "Введите описание", 
            name: "order",
            isError: false, 
            value:updateForm!.order,
          },
        ],
      }
    ]
    //useEffect(()=>{//setPopups()}, [addForm])
    const actions = useTypedSelector(({ actions }) => {
      return actions
    })
    const { getActions, sortActions, addNewAction, deleteAction, updateAction } = useActions();
    
    useEffect(()=>{
      if (actions?.actionsList.length === 0){
        getActions()
      }
    },[getActions, actions?.actionsList])

    const moveCard = useCallback(async (dragIndex: number, hoverIndex: number) => {
      if (actions!.actionsList.length>0){
        await sortActions([...actions!.actionsList],dragIndex, hoverIndex)
      }
    }, [sortActions, actions])

    if (actions?.loading){
      return <Spinner />
    }

    return (
      <React.Fragment>
        <div className="card-header p-4 border-300 bg-soft">
          <div className="row g-3 justify-content-between align-items-center">
            <div className="col-12 col-md">
              <h1>Акции</h1>
            </div>
            <div className="col col-md-auto">
              <nav className="nav nav-underline justify-content-end border-0 doc-tab-nav align-items-center" role="tablist">
              <button className="btn btn-primary" onClick={() => {setAddFormError(""); setOpened(0)}}>Добавить</button>
              </nav>
            </div>
          </div>
        </div>
        <div className='p-4 code-to-copy'>
          <DndProvider backend={HTML5Backend}>
              <div className='row g-4'>
                {actions?.actionsList.map((action, index)=>{
                  return <Card key={action.id} card={{ID: action.id!, Type:0, Title: action.title!, Text: action.description!, Image: "/api/store/"+action.preview!, Index: index, moveCard: moveCard, deleteClick: () => {deleteAction(action.id!)}, editClick:()=>{setUpdateForm({id:action.id!, title:action.title!, description: action.description!, date: action.date!, preview: action.preview!, photo:action.photo!, order: !action.order?0:action.order!, file:null}); setUpdateFormError(""); setOpened(1)}}}/>
                })}
              </div>
          </ DndProvider>
        </div>
        <Popups Forms={popups} Opened={opened} Toggle={()=>{setOpened(null)}} />
      </React.Fragment>
    );

}