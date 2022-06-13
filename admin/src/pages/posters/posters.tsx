import React, { useCallback, useEffect, useState } from 'react';
import Card from '../../components/card';
import Spinner from '../../components/spinner/spinner';
import { useActions } from '../../hooks/use-actions';
import { useTypedSelector } from '../../hooks/use-typed-selector';
import {DndProvider} from 'react-dnd'
import { HTML5Backend } from 'react-dnd-html5-backend';
import { Popups } from '../../components/popups/popups';


export const Posters: React.FC = () =>  {
    const [opened, setOpened] = useState<number|null>(null);
    const [addForm, setAddFrom] = useState({
      title: "",
      file: null,
    })
    const [addFormError, setAddFormError] = useState("")
    const [updateForm, setUpdateForm] = useState({
      id: "",
      title: "",
      photo: "",
      order: 0,
      file: null,
    })
    const [updateFormError, setUpdateFormError] = useState("")
    const addPosterCallback = (error: string) => {
      if (error!==""){
        setAddFormError(error)
      } else {
        setOpened(null);
      }
    }
    const updatePosterCallback = (error: string) => {
      if (error!==""){
        setUpdateFormError(error)
      } else {
        setOpened(null);
      }
    }
    const popups = [
      {
        title: "Добавить постер",
        stateAction: setAddFrom,
        sendAction: () =>{ addNewPoster(addForm.file, addForm.title, addPosterCallback);},
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
        title: "Редактировать постер",
        stateAction: setUpdateForm,
        sendAction: ()=>updatePoster(updateForm, updatePosterCallback, updateForm.file),
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
            title: "Тайтл",
            type: "text", 
            placeholder: "Введите заголовок", 
            name: "title", 
            isError: false, 
            value:updateForm!.title,
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
    const posters = useTypedSelector(({ posters }) => {
      return posters
    })
    const { getPosters, sortPosters, addNewPoster, deletePoster, updatePoster } = useActions();
    
    useEffect(()=>{
      if (posters?.postersList.length === 0){
        getPosters()
      }
    },[getPosters, posters?.postersList])

    const moveCard = useCallback(async (dragIndex: number, hoverIndex: number) => {
      if (posters!.postersList.length>0){
        await sortPosters([...posters!.postersList],dragIndex, hoverIndex)
      }
    }, [sortPosters, posters])

    if (posters?.loading){
      return <Spinner />
    }

    return (
      <React.Fragment>
        <div className="card-header p-4 border-300 bg-soft">
          <div className="row g-3 justify-content-between align-items-center">
            <div className="col-12 col-md">
              <h1>Постеры</h1>
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
                {posters?.postersList.map((poster, index)=>{
                  return <Card key={poster.id} card={{ID: poster.id!, Type:0, Title: poster.title!, Text: "", Image: "http://localhost:8080/store/"+poster.photo!, Index: index, moveCard: moveCard, deleteClick: () => {deletePoster(poster.id!)}, editClick:()=>{setUpdateForm({id:poster.id!, title:poster.title!, photo:poster.photo!, order: !poster.order?0:poster.order!, file:null}); setUpdateFormError(""); setOpened(1)}}}/>
                })}
              </div>
          </ DndProvider>
        </div>
        <Popups Forms={popups} Opened={opened} Toggle={()=>{setOpened(null)}} />
      </React.Fragment>
    );

}