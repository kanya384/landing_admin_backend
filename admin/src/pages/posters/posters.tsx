import React, { useEffect } from 'react';
import Card from '../../components/card';
import Spinner from '../../components/spinner/spinner';
import { useActions } from '../../hooks/use-actions';
import { useTypedSelector } from '../../hooks/use-typed-selector';


export const Posters: React.FC = () =>  {
    const posters = useTypedSelector(({ posters }) => {
      return posters
    })
    const { getPosters } = useActions();
    useEffect(()=>{
      if (posters?.postersList.length == 0){
        getPosters()
      }
    },[])
    if (posters?.loading){
      return <Spinner />
    }
    return (
      <React.Fragment>
      <h1>Posters</h1>
      {posters?.postersList.map((poster)=>{
        return <Card card={{Title: poster.title!, Text: poster.description!, Image: "http://localhost:8080/store/"+poster.photo!, Button: "Удалить"}}/>
      })}
      </React.Fragment>
    );

}