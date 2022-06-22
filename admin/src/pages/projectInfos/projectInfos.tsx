import React, { useCallback, useEffect, useState } from 'react';
import Card from '../../components/card';
import Spinner from '../../components/spinner/spinner';
import { useActions } from '../../hooks/use-actions';
import { useTypedSelector } from '../../hooks/use-typed-selector';
import {DndProvider} from 'react-dnd'
import { HTML5Backend } from 'react-dnd-html5-backend';
import { Popups } from '../../components/popups/popups';


export const ProjectInfos: React.FC = () =>  {
    const [opened, setOpened] = useState<number|null>(null);
    const [addForm, setAddFrom] = useState({
      title: "",
      anonce: "",
      description: "",
      file: null,
    })
    const [addFormError, setAddFormError] = useState("")
    const [updateForm, setUpdateForm] = useState({
      id: "",
      title: "",
      anonce: "",
      description: "",
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
        title: "Добавить информацию о проекте",
        stateAction: setAddFrom,
        sendAction: () =>{ addNewProjectInfo(addForm.file, addForm.title, addForm.anonce, addForm.description, addPosterCallback);},
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
          {
            title: "Анонс",
            type: "textarea", 
            placeholder: "Введите короткое описание", 
            name: "anonce", 
            isError: false, 
            value:addForm.anonce,
          },
          {
            title: "Описание",
            type: "textarea", 
            placeholder: "Введите полное описание", 
            name: "description", 
            isError: false, 
            value:addForm.description,
          },
        ],
      },
      {
        title: "Редактировать постер",
        stateAction: setUpdateForm,
        sendAction: ()=>updateProjectInfo(updateForm, updatePosterCallback, updateForm.file),
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
            title: "Анонс",
            type: "textarea", 
            placeholder: "Введите короткое описание", 
            name: "anonce", 
            isError: false, 
            value:updateForm.anonce,
          },
          {
            title: "Описание",
            type: "textarea", 
            placeholder: "Введите полное описание", 
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
    const projectInfos = useTypedSelector(({ projectInfos }) => {
      return projectInfos
    })
    const { getProjectInfos, sortProjectInfos, addNewProjectInfo, deleteProjectInfo, updateProjectInfo } = useActions();
    
    useEffect(()=>{
      if (projectInfos?.projectInfoList.length === 0){
        getProjectInfos()
      }
    },[getProjectInfos, projectInfos?.projectInfoList])

    const moveCard = useCallback(async (dragIndex: number, hoverIndex: number) => {
      if (projectInfos!.projectInfoList.length>0){
        await sortProjectInfos([...projectInfos!.projectInfoList],dragIndex, hoverIndex)
      }
    }, [sortProjectInfos, projectInfos])

    if (projectInfos?.loading){
      return <Spinner />
    }

    return (
      <React.Fragment>
        <div className="card-header p-4 border-300 bg-soft">
          <div className="row g-3 justify-content-between align-items-center">
            <div className="col-12 col-md">
              <h1>О проекте</h1>
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
                {projectInfos?.projectInfoList.map((projectInfo, index)=>{
                  return <Card key={projectInfo.id} card={{ID: projectInfo.id!, Type:0, Title: projectInfo.title!, Text: projectInfo.anonce!, Image: "/api/store/"+projectInfo.photo!, Index: index, moveCard: moveCard, deleteClick: () => {deleteProjectInfo(projectInfo.id!)}, editClick:()=>{setUpdateForm({id:projectInfo.id!, title:projectInfo.title!, anonce:projectInfo.anonce!, description:projectInfo.description!,  photo:projectInfo.photo!, order: !projectInfo.order?0:projectInfo.order!, file:null}); setUpdateFormError(""); setOpened(1)}}}/>
                })}
              </div>
          </ DndProvider>
        </div>
        <Popups Forms={popups} Opened={opened} Toggle={()=>{setOpened(null)}} />
      </React.Fragment>
    );

}