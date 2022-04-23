import type { FC } from 'react'
import { useRef } from 'react'
import { useDrag, useDrop } from 'react-dnd'
import "./card.css"
import type { Identifier, XYCoord } from 'dnd-core'
import { ItemTypes } from './itemTypes';

interface CardProps {
  ID: string,
  Title: string,
  Text: string,
  Image: string,
  Index: number,
  Button: string,
  moveCard: (dragIndex: number, hoverIndex: number) => void
}

interface DragItem {
  index: number
  id: string
  type: string
}

export const Card: FC<{card: CardProps}> = (props) => {
  const ref = useRef<HTMLDivElement>(null)
  const [{ handlerId }, drop] = useDrop<
    DragItem,
    void,
    { handlerId: Identifier | null }
  >({
    accept: ItemTypes.CARD,
    collect(monitor) {
      return {
        handlerId: monitor.getHandlerId(),
      }
    },
    hover(item: DragItem, monitor) {
      if (!ref.current) {
        return
      }
      const dragIndex = item.index
      const hoverIndex = props.card.Index

      // Don't replace items with themselves
      if (dragIndex === hoverIndex) {
        return
      }

      // Determine rectangle on screen
      const hoverBoundingRect = ref.current?.getBoundingClientRect()

      // Get vertical middle
      const hoverMiddleY =
        (hoverBoundingRect.bottom - hoverBoundingRect.top) / 3

      // Determine mouse position
      const clientOffset = monitor.getClientOffset()

      // Get pixels to the top
      const hoverClientY = (clientOffset as XYCoord).y - hoverBoundingRect.top

      // Only perform the move when the mouse has crossed half of the items height
      // When dragging downwards, only move when the cursor is below 50%
      // When dragging upwards, only move when the cursor is above 50%

      // Dragging downwards
      if (dragIndex < hoverIndex && hoverClientY < hoverMiddleY) {
        return
      }

      // Dragging upwards
      if (dragIndex > hoverIndex && hoverClientY > hoverMiddleY) {
        return
      }

      // Time to actually perform the action
      props.card.moveCard(dragIndex, hoverIndex)

      // Note: we're mutating the monitor item here!
      // Generally it's better to avoid mutations,
      // but it's good here for the sake of performance
      // to avoid expensive index searches.
      item.index = hoverIndex
    },
  })

  const [{ isDragging }, drag] = useDrag({
    type: ItemTypes.CARD,
    item: () => {
      let id = props.card.ID
      let index = props.card.Index
      console.log(id)
      return { id, index }
    },
    collect: (monitor: any) => ({
      isDragging: monitor.isDragging(),
    }),
  })

  const opacity = isDragging ? 0.5 : 1
  drag(drop(ref))
  return (
    <div ref={ref} style={{ opacity }} key={props.card.ID} data-handler-id={handlerId} className='col-sm-6 col-md-4 col-lg-3'>
      <div className="card" >
        <img className="card-img-top" src={props.card.Image} alt="..." />
        <div className="card-body">
          <h5 className="card-title">{props.card.Title}</h5>
          <p className="card-text">{props.card.Text}</p>
          <div className='row justify-content-around align-items-center'>
            <button className="btn btn-soft-primary col-5" style={{fontSize:12}}>Редактировать</button>
            <button className="btn btn-primary col-5" style={{fontSize:12}}>Удалить</button>
          </div>
        </div>
      </div>
    </div>
  )

}