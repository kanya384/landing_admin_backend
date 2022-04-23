import React, { useCallback, useEffect, useState } from 'react';
import Card from '../../components/card';
import Spinner from '../../components/spinner/spinner';
import { useActions } from '../../hooks/use-actions';
import { useTypedSelector } from '../../hooks/use-typed-selector';
import {DndProvider} from 'react-dnd'
import { HTML5Backend } from 'react-dnd-html5-backend';
import { Popup } from '../../components/popup';


export const Posters: React.FC = () =>  {
    const [isOpen, setIsOpen] = useState(false);
    const toggle = () => setIsOpen(!isOpen);
    const posters = useTypedSelector(({ posters }) => {
      return posters
    })
    const { getPosters, sortPosters } = useActions();
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
        <Popup title={"Добавить постер"} isOpen={isOpen} toggle={toggle}>
          <div>Modal</div>
        </Popup>
        <div className="card-header p-4 border-300 bg-soft">
          <div className="row g-3 justify-content-between align-items-center">
            <div className="col-12 col-md">
              <h1>Постеры</h1>
            </div>
            <div className="col col-md-auto">
              <nav className="nav nav-underline justify-content-end border-0 doc-tab-nav align-items-center" role="tablist">
              <button className="btn btn-primary" onClick={()=>toggle()}>Добавить</button>
              </nav>
            </div>
          </div>
        </div>
        <div className='p-4 code-to-copy'>
          <DndProvider backend={HTML5Backend}>
              <div className='row g-4'>
                {posters?.postersList.map((poster, index)=>{
                  return <Card key={poster.id} card={{ID: poster.id!, Title: poster.title!, Text: poster.description!, Image: "http://localhost:8080/store/"+poster.photo!, Index: index, moveCard: moveCard,  Button: "Удалить"}}/>
                })}
              </div>
          </ DndProvider>
        </div>
      </React.Fragment>
    );

}