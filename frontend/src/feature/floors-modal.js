import { useState } from "react"
import { Link } from "react-router-dom"
import { RangeSlider } from "../components/range-slider"
import { flatIDs } from "./flat-ids"

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
        <div className={"modal-flat-map"}>
          <div id={"lit"+liter+"__flat-"+floor} className={floor>1?"lit"+liter+"__flat-1__2-9":"lit"+liter+"__flat-1__"+floor}>
            <img src={`img/floors/liter${liter}/${floor>1?"2-9":1}/floor.png`} alt="" />
            {flatIDs[`liter${liter}`] && flatIDs[`liter${liter}`][`floor${floor}`] && flatIDs[`liter${liter}`][`floor${floor}`][`entrance${entrance+1}`]?.map((id, index) => {
              return  <Link to={"/plans/"+id}><div class={`flat-room flat-room_${index+1}`} ><div class="flat-room_before"></div></div></Link>
            })}
          </div>
        </div>
        <div className="modal-flat-filter modal-flat-filter--mobile">
          <div className="floor-slider">
            <div className="floor-slider__title">Этаж</div>
            <RangeSlider value={floor} setValue={(floor) => setFloor(floor)} min={1} max={9} />
          </div>
          <ul className="entrance-list">
            <li className="entrance-item">подьезд 1</li>
            <li className="entrance-item active">подьезд 2</li>
          </ul>
        </div>
      </div>
    </div>
  )
}
