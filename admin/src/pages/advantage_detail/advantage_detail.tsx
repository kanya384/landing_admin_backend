import React, { useCallback, useEffect, useState } from "react"
import { DndProvider } from "react-dnd"
import { HTML5Backend } from "react-dnd-html5-backend"
import { useParams } from "react-router-dom"
import Card from "../../components/card"
import { Popups } from "../../components/popups/popups"
import Spinner from "../../components/spinner/spinner"
import { useActions } from "../../hooks/use-actions"
import { useTypedSelector } from "../../hooks/use-typed-selector"

export const AdvantageDetail = () => {
  const { id } = useParams()
  const {getAdvantageById, getAdvantagePhotosByID, addAdvantagePhoto, deleteAdvantagePhotoByID} = useActions()
  const [opened, setOpened] = useState<number|null>(null);

  const advantages = useTypedSelector(({ advantages }) => {
    return advantages
  })

  const [error, setError] = useState("")

  const [photosForm, setPhotosForm] = useState<{photos:any[]}>({
    photos: [],
  })

  useEffect(()=>{
    getAdvantageById(id!)
    getAdvantagePhotosByID(id!)
  },[])

  const sendFormallback = (error: string) => {
    if (error!==""){
      setError(error)
    } else {
      setOpened(null);
    }
  }

  const popups = [
    {
      title: `Добавить фото`,
      stateAction: setPhotosForm,
      sendAction: () =>{ addAdvantagePhoto(photosForm.photos, id!, sendFormallback);},
      error: error,
      fields: [
        {
          title: "Перетащите фото",
          type: "drop", 
          placeholder: "", 
          name: "photos", 
          isError: false, 
          value:photosForm.photos,
        },
      ],
    },
  ]

  const moveCard = useCallback(async (dragIndex: number, hoverIndex: number) => {
    if (advantages!.photosList.length>0){
      //await sortPhotos([...advantages!.photosList],dragIndex, hoverIndex)
    }
  }, [/*sortPhotos,*/ advantages])

  if (advantages?.loading) {
    return <Spinner />
  }
  if (advantages?.errorDet) {
    return <div>Не найдено преимущество с таким id</div>
  }
  return(
    <React.Fragment>
      <div className="card-header p-4 border-300 bg-soft">
        <div className="row g-3 justify-content-between align-items-center">
          <div className="col-12 col-md">
            <h1>{advantages?.detail?.title}</h1>
          </div>
          <div className="col col-md-auto">
            <nav className="nav nav-underline justify-content-end border-0 doc-tab-nav align-items-center" role="tablist">
            {/*<button className="btn btn-primary" onClick={() => {}}>Редактировать</button>*/}
            </nav>
          </div>
        </div>
        <p className="card-text fs-1 mt-4">{advantages?.detail?.description}</p>
        <div className="dropdown-divider"></div>
        <div className="row g-3 justify-content-between align-items-center">
          <div className="col-12 col-md">
            <h3>Фото</h3>
          </div>
          <div className="col col-md-auto">
            <nav className="nav nav-underline justify-content-end border-0 doc-tab-nav align-items-center" role="tablist">
            <button className="btn btn-primary" onClick={() => {setOpened(0)}}>Добавить</button>
            </nav>
          </div>
        </div>
        <div className='mt-6'>
          <DndProvider backend={HTML5Backend}>
            <div className='row g-4'>
              {advantages!.photosList?.map((photo, index)=>{
                return <Card key={photo.id} card={{ID: photo.id!, Type:1, Title: "", Text: "", Image: "/api/store/"+photo.image!, Index: index, moveCard: moveCard, deleteClick: () => {deleteAdvantagePhotoByID(photo.id!)}, editClick: ()=>{}}} />
              })}
            </div>
          </DndProvider>
        </div>
      </div>
      <Popups Forms={popups} Opened={opened} Toggle={()=>{setOpened(null)}} />
    </React.Fragment>
  )
}