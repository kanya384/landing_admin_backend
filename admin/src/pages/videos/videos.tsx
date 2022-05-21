import React, { useEffect, useState } from "react";
import { DndProvider } from "react-dnd";
import { HTML5Backend } from "react-dnd-html5-backend";
import Card from "../../components/card";
import { Popups } from "../../components/popups/popups";
import { useActions } from "../../hooks/use-actions";
import { useTypedSelector } from "../../hooks/use-typed-selector";

export const Videos: React.FC = () => {
  const [opened, setOpened] = useState<number|null>(null)
  const [formError, setFormError] = useState("")
  const [addForm, setAddFrom] = useState({
    url: "",
    file: null
  })
  const [updateForm, setUpdateFrom] = useState({
    id: "",
    url: "",
    file: null
  })

  const videos = useTypedSelector(({ videos }) => {
    return videos
  })

  const {getVideos, addNewVideo, deleteVideo, updateVideo} = useActions()

  const formCallback = (error: string) => {
    if (error!==""){
      setFormError(error)
    } else {
      setOpened(null);
    }
  }
  const popups = [
    {
      title: "Добавить видео",
      stateAction: setAddFrom,
      sendAction: () =>{ addNewVideo(addForm.file, addForm.url, formCallback);},
      error: formError,
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
          title: "URL видео файла",
          type: "text", 
          placeholder: "Введите URL", 
          name: "url", 
          isError: false, 
          value:addForm.url,
        },
      ],
    },
    {
      title: "Редактировать видео",
      stateAction: setUpdateFrom,
      sendAction: () =>{ updateVideo(updateForm.file, updateForm.id, updateForm.url, formCallback);},
      error: formError,
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
          title: "URL видео файла",
          type: "text", 
          placeholder: "Введите URL",
          name: "url", 
          isError: false, 
          value:updateForm.url,
        },
      ],
    },
  ]

  useEffect(()=>{
    if (videos?.videoList.length === 0){
      getVideos()
    }
  },[getVideos, videos?.videoList])
  return (
    <React.Fragment>
      <div className="card-header p-4 border-300 bg-soft">
        <div className="row g-3 justify-content-between align-items-center">
          <div className="col-12 col-md">
            <h1>Видео</h1>
          </div>
          <div className="col col-md-auto">
            <nav className="nav nav-underline justify-content-end border-0 doc-tab-nav align-items-center" role="tablist">
            <button className="btn btn-primary" onClick={() => {setFormError(""); setOpened(0); setAddFrom({url:"", file:null})}}>Добавить</button>
            </nav>
          </div>
        </div>
      </div>
      <div className='p-4 code-to-copy'>
        <DndProvider backend={HTML5Backend}>
            <div className='row g-4 justify-content-center'>
              {videos?.videoList.map((video, index)=>{
                return <Card key={video.id} card={{ID: video.id!, Type:3, Title: "", Text: "", Image: "http://localhost:8080/store/"+video.preview, url:video.url, Index: index, moveCard: () =>{}, deleteClick: () => {deleteVideo(video.id!)}, editClick:()=>{setUpdateFrom({id:video.id!, url:video.url!, file:null}); setFormError(""); setOpened(1)}}}/>
              })}
            </div>
        </ DndProvider>
      </div>
      <Popups Forms={popups} Opened={opened} Toggle={()=>{setOpened(null)}} />
    </React.Fragment>)
}