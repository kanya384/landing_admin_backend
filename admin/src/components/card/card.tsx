import React from 'react';
interface Card {
  Title: string,
  Text: string,
  Image: string,
  Button: string,
}
export const Card: React.FC<{card: Card}> = (props) => {
  return (
    <div className="card" style={{width:"20rem"}}>
      <img className="card-img-top" src={props.card.Image} alt="..." />
      <div className="card-body">
        <h5 className="card-title">{props.card.Title}</h5>
        <p className="card-text">{props.card.Text}</p>
        <button className="btn btn-primary">{props.card.Button}</button>
      </div>
    </div>
  )
}