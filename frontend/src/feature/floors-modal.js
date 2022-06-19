import { useState } from "react"
import { RangeSlider } from "../components/range-slider"

export const FloorsModal = ({title, classes, liter, opened, close}) => {
  const [entrance, setEntrance] = useState(0)
  const [floor, setFloor] = useState(1)
  return (
    <div className={opened?`modal-full ${classes} open`:`modal-full ${classes}`}  >
      <div className="modal-full__close js-close-modal" onClick={()=>close()}>
        <div className="close-icon">
          <svg className="svg"><use xlinkHref="img/sprite.svg#close"></use></svg>
        </div>
      </div>
      <div className="modal-full__content">
        <div className="modal-flat__header">
          <div className="modal-flat-title">{title}</div>
          <div className="modal-flat-filter">
            <div className="floor-slider">
              <div className="floor-slider__title">Этаж</div>
              <RangeSlider value={floor} setValue={(floor) => setFloor(floor)} min={1} max={9} />
            </div>
            <ul className="entrance-list">
              <li className={entrance==0?"entrance-item active":"entrance-item"} onClick={()=>setEntrance(0)}>подьезд 1</li>
              <li className={entrance==1?"entrance-item active":"entrance-item"} onClick={()=>setEntrance(1)}>подьезд 2</li>
            </ul>
          </div>
        </div>
        <div className={"modal-flat-map liter_"+liter+"-floor_"+floor}><img src={`img/floors/liter${liter}/${floor>1?"2-9":1}/floor.png`} alt="" /></div>
        <div className="modal-flat-filter modal-flat-filter--mobile">
          <div className="floor-slider"><div className="floor-slider__title">Этаж</div></div>
          <ul className="entrance-list">
            <li className="entrance-item">подьезд 1</li>
            <li className="entrance-item active">подьезд 2</li>
          </ul>
        </div>
      </div>
    </div>
  )
}